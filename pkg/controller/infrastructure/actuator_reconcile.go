package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/metal-pod/gardener-extension-provider-metal/pkg/metal"
	metalclient "github.com/metal-pod/gardener-extension-provider-metal/pkg/metal/client"

	metalapi "github.com/metal-pod/gardener-extension-provider-metal/pkg/apis/metal"
	metalgo "github.com/metal-pod/metal-go"
	metalfirewall "github.com/metal-pod/metal-go/api/client/firewall"

	extensionscontroller "github.com/gardener/gardener-extensions/pkg/controller"
	controllererrors "github.com/gardener/gardener-extensions/pkg/controller/error"

	v1alpha1constants "github.com/gardener/gardener/pkg/apis/core/v1alpha1/constants"

	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/gardener/gardener/pkg/utils/secrets"

	"github.com/coreos/container-linux-config-transpiler/config/types"
)

const (
	firewallPolicyControllerName = "firewall-policy-controller"
	droptailerClientName         = "droptailer"
)

func (a *actuator) reconcile(ctx context.Context, infrastructure *extensionsv1alpha1.Infrastructure, cluster *extensionscontroller.Cluster) error {
	infrastructureConfig, infrastructureStatus, err := a.decodeInfrastructure(infrastructure)
	if err != nil {
		return err
	}

	mclient, err := metalclient.NewClient(ctx, a.client, &infrastructure.Spec.SecretRef)
	if err != nil {
		return err
	}

	clusterID := cluster.Shoot.GetUID()
	clusterTag := fmt.Sprintf("%s=%s", metal.ShootAnnotationClusterID, clusterID)

	nodeCIDR, err := a.ensureNodeNetwork(ctx, string(clusterID), mclient, infrastructure, infrastructureConfig, infrastructureStatus, cluster)
	if err != nil {
		return &controllererrors.RequeueAfterError{
			Cause:        err,
			RequeueAfter: 30 * time.Second,
		}
	}

	firewallStatus := infrastructureStatus.Firewall
	if firewallStatus.Succeeded {
		// verify that the firewall is still there and correctly reconciled
		machineID := decodeMachineID(firewallStatus.MachineID)

		resp, err := mclient.FirewallFind(&metalgo.FirewallFindRequest{
			MachineFindRequest: metalgo.MachineFindRequest{
				ID:                &machineID,
				AllocationProject: &infrastructureConfig.ProjectID,
				Tags:              []string{clusterTag},
			},
		})
		if err != nil {
			return &controllererrors.RequeueAfterError{
				Cause:        err,
				RequeueAfter: 30 * time.Second,
			}
		}

		if len(resp.Firewalls) > 0 {
			fw := resp.Firewalls[0]
			// TODO: in the future we probably want to allow resizing a firewall or updating the machine image
			// this could go here...
			if *fw.Size.ID != infrastructureConfig.Firewall.Size {
				a.logger.Error(fmt.Errorf("firewall size has changed"), "firewall spec has changed, which is not supported", "clusterid", clusterID, "machineid", machineID)
			}
			if *fw.Allocation.Image.ID != infrastructureConfig.Firewall.Image {
				a.logger.Error(fmt.Errorf("firewall image has changed"), "firewall spec has changed, which is not supported", "clusterid", clusterID, "machineid", machineID)
			}
			return nil
		}

		a.logger.Error(err, "firewall does not exist anymore, creating new firewall", "clusterid", clusterID, "machineid", machineID)

		firewallStatus.MachineID = ""
		firewallStatus.Succeeded = false
		err = a.updateProviderStatus(ctx, infrastructure, infrastructureConfig, firewallStatus)
		if err != nil {
			return err
		}
	}

	if firewallStatus.MachineID != "" {
		// firewall was created, waiting for completion
		machineID := decodeMachineID(firewallStatus.MachineID)

		resp, err := mclient.FirewallGet(machineID)
		if err != nil {
			switch e := err.(type) {
			case *metalfirewall.FindFirewallDefault:
				if e.Code() >= 500 {
					return &controllererrors.RequeueAfterError{
						Cause:        e,
						RequeueAfter: 5 * time.Second,
					}
				}
			default:
				return e
			}
		}

		allocation := resp.Firewall.Allocation
		if allocation == nil {
			return fmt.Errorf("firewall %q was created but has no allocation", machineID)
		}

		firewallStatus.Succeeded = *resp.Firewall.Allocation.Succeeded
		return a.updateProviderStatus(ctx, infrastructure, infrastructureConfig, firewallStatus)
	}

	// we need to create a firewall
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	// Example values:
	// cluster.Shoot.Status.TechnicalID  "shoot--dev--johndoe-metal"
	clusterName := cluster.Shoot.Status.TechnicalID
	name := clusterName + "-firewall-" + uuid.String()[:5]

	// find private network
	privateNetwork, err := metalclient.GetPrivateNetworkFromNodeNetwork(mclient, infrastructureConfig.ProjectID, nodeCIDR)
	if err != nil {
		return err
	}

	kubeconfig, err := a.createFirewallPolicyControllerKubeconfig(ctx, infrastructure, cluster)
	if err != nil {
		return err
	}

	firewallUserData, err := a.renderFirewallUserData(kubeconfig)
	if err != nil {
		return err
	}

	// assemble firewall allocation request
	var networks []metalgo.MachineAllocationNetwork
	network := metalgo.MachineAllocationNetwork{
		NetworkID:   *privateNetwork.ID,
		Autoacquire: true,
	}
	networks = append(networks, network)
	for _, n := range infrastructureConfig.Firewall.Networks {
		network := metalgo.MachineAllocationNetwork{
			NetworkID:   n,
			Autoacquire: true,
		}
		networks = append(networks, network)
	}

	createRequest := &metalgo.FirewallCreateRequest{
		MachineCreateRequest: metalgo.MachineCreateRequest{
			Description:   name + " created by Gardener",
			Name:          name,
			Hostname:      name,
			Size:          infrastructureConfig.Firewall.Size,
			Project:       infrastructureConfig.ProjectID,
			Partition:     infrastructureConfig.PartitionID,
			Image:         infrastructureConfig.Firewall.Image,
			SSHPublicKeys: []string{string(infrastructure.Spec.SSHPublicKey)},
			Networks:      networks,
			UserData:      firewallUserData,
			Tags:          []string{clusterTag},
		},
	}

	a.logger.Info("create firewall", "name", createRequest.Name)

	fcr, err := mclient.FirewallCreate(createRequest)
	if err != nil {
		a.logger.Error(err, "failed to create firewall", "infrastructure", infrastructure.Name)
		return &controllererrors.RequeueAfterError{
			Cause:        err,
			RequeueAfter: 30 * time.Second,
		}
	}

	machineID := encodeMachineID(*fcr.Firewall.Partition.ID, *fcr.Firewall.ID)

	allocation := fcr.Firewall.Allocation
	if allocation == nil {
		return fmt.Errorf("firewall %q was created but has no allocation", machineID)
	}

	firewallStatus.MachineID = machineID
	firewallStatus.Succeeded = true

	return a.updateProviderStatus(ctx, infrastructure, infrastructureConfig, firewallStatus)
}

func (a *actuator) ensureNodeNetwork(ctx context.Context, clusterID string, mclient *metalgo.Driver, infrastructure *extensionsv1alpha1.Infrastructure, infrastructureConfig *metalapi.InfrastructureConfig, infrastructureStatus *metalapi.InfrastructureStatus, cluster *extensionscontroller.Cluster) (string, error) {
	if cluster.Shoot.Spec.Networking.Nodes != nil {
		return *cluster.Shoot.Spec.Networking.Nodes, nil
	}
	if infrastructure.Status.NodesCIDR != nil {
		resp, err := mclient.NetworkFind(&metalgo.NetworkFindRequest{
			ProjectID:   &infrastructureConfig.ProjectID,
			PartitionID: &infrastructureConfig.PartitionID,
			Labels:      map[string]string{metal.ShootAnnotationClusterID: clusterID},
		})
		if err != nil {
			return "", err
		}

		if len(resp.Networks) != 0 {
			return *infrastructure.Status.NodesCIDR, nil
		}

		return "", fmt.Errorf("node network disappeared from cloud provider: %s", *infrastructure.Status.NodesCIDR)
	}

	resp, err := mclient.NetworkAllocate(&metalgo.NetworkAllocateRequest{
		ProjectID:   infrastructureConfig.ProjectID,
		PartitionID: infrastructureConfig.PartitionID,
		Name:        cluster.Shoot.GetName(),
		Description: clusterID,
		Labels:      map[string]string{metal.ShootAnnotationClusterID: clusterID},
	})
	if err != nil {
		return "", err
	}

	nodeCIDR := resp.Network.Prefixes[0]

	infrastructure.Status.NodesCIDR = &nodeCIDR
	err = a.updateProviderStatus(ctx, infrastructure, infrastructureConfig, infrastructureStatus.Firewall)
	if err != nil {
		return "", err
	}

	return nodeCIDR, nil
}

func (a *actuator) createFirewallPolicyControllerKubeconfig(ctx context.Context, infrastructure *extensionsv1alpha1.Infrastructure, cluster *extensionscontroller.Cluster) (string, error) {
	apiServerURL := fmt.Sprintf("api.%s", *cluster.Shoot.Spec.DNS.Domain)
	infrastructureSecrets := &secrets.Secrets{
		CertificateSecretConfigs: map[string]*secrets.CertificateSecretConfig{
			v1alpha1constants.SecretNameCACluster: {
				Name:       v1alpha1constants.SecretNameCACluster,
				CommonName: "kubernetes",
				CertType:   secrets.CACert,
			},
		},
		SecretConfigsFunc: func(cas map[string]*secrets.Certificate, clusterName string) []secrets.ConfigInterface {
			return []secrets.ConfigInterface{
				&secrets.ControlPlaneSecretConfig{
					CertificateSecretConfig: &secrets.CertificateSecretConfig{
						Name:         firewallPolicyControllerName,
						CommonName:   fmt.Sprintf("system:%s", firewallPolicyControllerName),
						Organization: []string{firewallPolicyControllerName},
						CertType:     secrets.ClientCert,
						SigningCA:    cas[v1alpha1constants.SecretNameCACluster],
					},
					KubeConfigRequest: &secrets.KubeConfigRequest{
						ClusterName:  clusterName,
						APIServerURL: apiServerURL,
					},
				},
			}
		},
	}

	secret, err := infrastructureSecrets.Deploy(ctx, a.clientset, a.gardenerClientset, infrastructure.Namespace)
	if err != nil {
		return "", err
	}

	kubeconfig, ok := secret[firewallPolicyControllerName].Data["kubeconfig"]
	if !ok {
		return "", fmt.Errorf("kubeconfig not part of generated firewall policy controller secret")
	}

	return string(kubeconfig), nil
}

func (a *actuator) renderFirewallUserData(kubeconfig string) (string, error) {
	cfg := types.Config{}
	cfg.Systemd = types.Systemd{}

	enabled := true
	fpcUnit := types.SystemdUnit{
		Name:    fmt.Sprintf("%s.service", firewallPolicyControllerName),
		Enable:  enabled,
		Enabled: &enabled,
	}
	dcUnit := types.SystemdUnit{
		Name:    fmt.Sprintf("%s.service", droptailerClientName),
		Enable:  enabled,
		Enabled: &enabled,
	}

	cfg.Systemd.Units = append(cfg.Systemd.Units, fpcUnit, dcUnit)

	cfg.Storage = types.Storage{}

	mode := 0600
	id := 0
	ignitionFile := types.File{
		Path:       "/etc/firewall-policy-controller/.kubeconfig",
		Filesystem: "root",
		Mode:       &mode,
		User: &types.FileUser{
			Id: &id,
		},
		Group: &types.FileGroup{
			Id: &id,
		},
		Contents: types.FileContents{
			Inline: string(kubeconfig),
		},
	}
	cfg.Storage.Files = append(cfg.Storage.Files, ignitionFile)

	outCfg, report := types.Convert(cfg, "", nil)
	if report.IsFatal() {
		return "", fmt.Errorf("could not transpile ignition config: %s", report.String())
	}

	userData, err := json.Marshal(outCfg)
	if err != nil {
		return "", err
	}

	return string(userData), nil
}

package metal

import "path/filepath"

const (
	// Name is the name of the Metal provider.
	Name = "provider-metal"

	// MachineControllerManagerImageName is the name of the MachineControllerManager image.
	MachineControllerManagerImageName = "machine-controller-manager"
	// MCMProviderMetalImageName is the name of the metal provider plugin image.
	MCMProviderMetalImageName = "machine-controller-manager-provider-metal"
	// CCMImageName is the name of the cloud controller manager image.
	CCMImageName = "metalccm"
	// GroupRolebindingControllerImageName is the name of the GroupRolebindingController image
	GroupRolebindingControllerImageName = "group-rolebinding-controller"
	// AccountingExporterImageName is the name of the accounting exporter image
	AccountingExporterImageName = "accounting-exporter"
	// AuthNWebhookImageName is the name of the AuthN Webhook configured with the shoot kube-apiserver
	AuthNWebhookImageName = "authn-webhook"
	// SplunkAuditWebhookImageName is the name of the splunk audit Webhook configured with the shoot kube-apiserver
	SplunkAuditWebhookImageName = "splunk-audit-webhook"
	// DroptailerImageName is the name of the Droptailer to deploy to the shoot.
	DroptailerImageName = "droptailer"
	// MetallbSpeakerImageName is the name of the metallb speaker to deploy to the shoot.
	MetallbSpeakerImageName = "metallb-speaker"
	// MetallbControllerImageName is the name of the metallb controller to deploy to the shoot.
	MetallbControllerImageName = "metallb-controller"
	// LimitValidatingWebhookImageName is the name of the limit validating webhook to deploy to the seed's shoot namespace.
	LimitValidatingWebhookImageName = "limit-validating-webhook"
	// CSIControllerImageName is the name of the csi lvm controller to deploy to the shoot.
	CSIControllerImageName = "csi-lvm-controller"
	// CSIProvisionerImageName is the name of the csi lvm provisioner to deploy to the shoot.
	CSIProvisionerImageName = "csi-lvm-provisioner"

	// CSIPluginImageName is the name of the csi-driver-lvm plugin to deploy to the shoot.
	CSIPluginImageName = "csi-lvmplugin"
	// CSIPluginProvisionerImageName is the name of the csi-driver-lvm provisioner to deploy to the shoot.
	CSIPluginProvisionerImageName = "csi-lvmplugin-provisioner"
	// CSIExternalAttacherImageName is the name of the csi-attacher image to deploy to the shoot.
	CSIExternalAttacherImageName = "csi-attacher"
	// CSIExternalProvisionerImageName is the name of the csi-provisioner image to deploy to the shoot.
	CSIExternalProvisionerImageName = "csi-provisioner"
	// CSIExternalNodeDriverRegistrarImageName is the name of the csi node-driver to deploy to the shoot.
	CSIExternalNodeDriverRegistrarImageName = "csi-node-driver-registrar"
	// CSIExternalResizerImageName is the name of the csi-resizer image to deploy to the shoot.
	CSIExternalResizerImageName = "csi-resizer"
	// CSIExternalLivenessImageName is the name of the csi-driver-lvm provisioner to deploy to the shoot.
	CSIExternalLivenessImageName = "csi-livenessprobe"

	// APIKey is a constant for the key in a cloud provider secret.
	APIKey = "metalAPIKey"
	// APIHMac is a constant for the hmac in a cloud provider secret.
	APIHMac = "metalAPIHMac"

	// CloudProviderConfigName is the name of the configmap containing the cloud provider config.
	CloudProviderConfigName = "cloud-provider-config"
	// MachineControllerManagerName is a constant for the name of the machine-controller-manager.
	MachineControllerManagerName = "machine-controller-manager"

	AuthNWebHookConfigName               = "authn-webhook-config"
	AuthNWebHookCertName                 = "authn-webhook-cert"
	ShootExtensionTypeTokenIssuer        = "tokenissuer"
	DroptailerNamespace                  = "firewall"
	DroptailerClientSecretName           = "droptailer-client"
	DroptailerServerSecretName           = "droptailer-server"
	CloudControllerManagerDeploymentName = "cloud-controller-manager"
	CloudControllerManagerServerName     = "cloud-controller-manager-server"
	GroupRolebindingControllerName       = "group-rolebinding-controller"
	LimitValidatingWebhookDeploymentName = "limit-validating-webhook"
	LimitValidatingWebhookServerName     = "limit-validating-webhook-server"
	AccountingExporterName               = "accounting-exporter"
	AuthNWebhookDeploymentName           = "kube-jwt-authn-webhook"
	AuthNWebhookServerName               = "kube-jwt-authn-webhook-server"
	SplunkAuditWebhookDeploymentName     = "splunk-audit-webhook"
	SplunkAuditWebhookServerName         = "splunk-audit-webhook-server"
	SplunkAuditWebHookConfigName         = "splunk-audit-webhook-config"
	SplunkAuditWebHookCertName           = "splunk-audit-webhook-cert"
)

var (
	// ChartsPath is the path to the charts
	ChartsPath = filepath.Join("controllers", Name, "charts")
	// InternalChartsPath is the path to the internal charts
	InternalChartsPath = filepath.Join(ChartsPath, "internal")
)

// Credentials stores Metal credentials.
type Credentials struct {
	MetalAPIKey  string
	MetalAPIHMac string
}

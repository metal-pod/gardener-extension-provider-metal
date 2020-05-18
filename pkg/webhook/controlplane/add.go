package controlplane

import (
	extensionswebhook "github.com/gardener/gardener/extensions/pkg/webhook"
	"github.com/gardener/gardener/extensions/pkg/webhook/controlplane"
	"github.com/gardener/gardener/extensions/pkg/webhook/controlplane/genericmutator"
	"github.com/metal-stack/gardener-extension-provider-metal/pkg/apis/config"
	"github.com/metal-stack/gardener-extension-provider-metal/pkg/metal"

	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var (
	// DefaultAddOptions are the default AddOptions for AddToManager.
	DefaultAddOptions = AddOptions{}
	logger            = log.Log.WithName("metal-controlplane-webhook")
)

// AddOptions are options to apply when adding the metal infrastructure controller to the manager.
type AddOptions struct {
	// Controller are the controller.Options.
	ControllerConfig config.ControllerConfiguration
}

func AddToManagerWithOptions(mgr manager.Manager, opts AddOptions) (*extensionswebhook.Webhook, error) {
	logger.Info("Adding webhook to manager")
	fciCodec := controlplane.NewFileContentInlineCodec()
	return controlplane.Add(mgr, controlplane.AddArgs{
		Kind:     controlplane.KindShoot,
		Provider: metal.Type,
		Types:    []runtime.Object{&appsv1.Deployment{}, &extensionsv1alpha1.OperatingSystemConfig{}},
		Mutator: genericmutator.NewMutator(NewEnsurer(logger, opts.ControllerConfig), controlplane.NewUnitSerializer(),
			controlplane.NewKubeletConfigCodec(fciCodec), fciCodec, logger),
	})
}

// AddToManager creates a webhook and adds it to the manager.
func AddToManager(mgr manager.Manager) (*extensionswebhook.Webhook, error) {
	return AddToManagerWithOptions(mgr, DefaultAddOptions)
}

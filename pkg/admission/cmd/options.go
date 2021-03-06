package cmd

import (
	extensionswebhook "github.com/gardener/gardener/extensions/pkg/webhook"
	webhookcmd "github.com/gardener/gardener/extensions/pkg/webhook/cmd"

	"github.com/metal-stack/gardener-extension-provider-metal/pkg/admission/validator"
)

// GardenWebhookSwitchOptions are the webhookcmd.SwitchOptions for the admission webhooks.
func GardenWebhookSwitchOptions() *webhookcmd.SwitchOptions {
	return webhookcmd.NewSwitchOptions(
		webhookcmd.Switch(extensionswebhook.ValidatorName, validator.New),
	)
}

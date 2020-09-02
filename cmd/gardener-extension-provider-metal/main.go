package main

import (
	"github.com/gardener/gardener/extensions/pkg/controller"
	"github.com/metal-stack/gardener-extension-provider-metal/cmd/gardener-extension-provider-metal/app"

	controllercmd "github.com/gardener/gardener/extensions/pkg/controller/cmd"
	"github.com/gardener/gardener/extensions/pkg/log"
	runtimelog "sigs.k8s.io/controller-runtime/pkg/log"
)

func main() {
	runtimelog.SetLogger(log.ZapLogger(false))
	cmd := app.NewControllerManagerCommand(controller.SetupSignalHandlerContext())

	if err := cmd.Execute(); err != nil {
		controllercmd.LogErrAndExit(err, "error executing the main controller command")
	}
}

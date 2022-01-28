package main

import (
	"github.com/chaunceyjiang/sample-k8s-scheduler/pkg/plugins"
	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	command := app.NewSchedulerCommand(
		app.WithPlugin(plugins.Name, plugins.New),
	)
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

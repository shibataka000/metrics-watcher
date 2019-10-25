package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/shibataka000/metrics-watcher/pkg/kubernetes"
	"github.com/shibataka000/metrics-watcher/pkg/metricswatcher"
	"github.com/urfave/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func action(c *cli.Context) error {
	clientConfig, err := kubernetes.NewClientConfig(c.String("kubeconfig"), c.String("context"))
	if err != nil {
		return err
	}
	log.Printf("%v\n", clientConfig)

	config, err := metricswatcher.NewConfig(c.String("namespace"), c.String("selector"))
	if err != nil {
		return err
	}
	watcher, err := metricswatcher.NewMetricsWatcher(clientConfig, config)
	if err != nil {
		return err
	}

	err = watcher.Start()

	return err
}

func main() {
	app := cli.NewApp()
	app.Name = "metricswatcher"
	app.Usage = "Output some information to know HorizontalPodAutoscaler internal behavior"
	app.UsageText = "metricswatcher deployment"
	app.Version = "v0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "kubeconfig",
			Value: filepath.Join(os.Getenv("HOME"), ".kube", "config"),
			Usage: "Path to the kubeconfig file to use for CLI requests.",
		},
		cli.StringFlag{
			Name:  "context",
			Value: "",
			Usage: "The name of the kubeconfig context to use",
		},
		cli.StringFlag{
			Name:  "namespace, n",
			Value: "default",
			Usage: "If present, the namespace scope for this CLI request",
		},
		cli.StringFlag{
			Name:  "selector, l",
			Value: "",
			Usage: "Selector (label query) to filter on, supports '='.(e.g. -l key1=value1,key2=value2)",
		},
	}
	app.Action = action
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

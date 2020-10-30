package main

import (
	"log"
	"os"
	"time"

	"github.com/emicklei/gdep/bigquery"
	"github.com/urfave/cli/v2"
)

var version = time.Now().String()

func main() {
	if err := newApp().Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Version = version
	app.EnableBashCompletion = true
	app.Name = "gdep"
	app.Usage = `Google Cloud Platform resource dependencies tool

	see https://github.com/emicklei/gdep for documentation.
`
	// override -v
	cli.VersionFlag = &cli.BoolFlag{
		Name:  "print-version, V",
		Usage: "print only the version",
	}
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "v",
			Usage: "verbose logging",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:  "bigquery",
			Usage: "gdep bigquery PROJECT(.|:)DATASET.VIEW,...",
			Action: func(c *cli.Context) error {
				defer logBegin(c)()
				return logEnd(c, bigquery.ExportViewDepencyGraph(c))
			},
		},
	}
	return app
}

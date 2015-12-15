package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/untoldwind/gotrack/server/commands"
	"github.com/untoldwind/gotrack/server/config"
)

func main() {
	app := cli.NewApp()
	app.Name = "gotrack"
	app.Usage = "Collect connection traffic information in linux based router"
	app.Version = config.Version()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config-dir",
			Value: "/etc/gotrack.d",
			Usage: "config directory",
		},
		cli.StringFlag{
			Name:  "log-file",
			Value: "",
			Usage: "Log to file instead stdout",
		},
		cli.StringFlag{
			Name:  "log-format",
			Value: "text",
			Usage: "Log format to use (test, json, logstash)",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Enable debug logging",
		},
	}
	app.Commands = []cli.Command{
		commands.ServerCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Errorf("Failed to run command: %s", err.Error())
	}
}

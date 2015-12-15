package commands

import (
	"github.com/codegangsta/cli"
	"github.com/untoldwind/gotrack/server/conntrack"
)

var ServerCommand = cli.Command{
	Name:   "server",
	Usage:  "Start server",
	Action: runWithContext(serverCommand),
}

func serverCommand(ctx *cli.Context, runCtx *runContext) {
	_, err := conntrack.NewProvider(runCtx.config.Provider, runCtx.logger)
	if err != nil {
		runCtx.logger.ErrorErr(err)
		return
	}

}

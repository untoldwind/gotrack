package commands

import (
	"github.com/codegangsta/cli"
	"github.com/untoldwind/gotrack/server/conntrack"
	"github.com/untoldwind/gotrack/server/http"
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

	server := http.NewServer(runCtx.config.Server, runCtx.logger)

	if err := server.Start(); err != nil {
		runCtx.logger.ErrorErr(err)
		return
	}
	defer server.Stop()

	runCtx.handleSignals()
}

package commands

import (
	"github.com/codegangsta/cli"
	"github.com/untoldwind/gotrack/server/conntrack"
	"github.com/untoldwind/gotrack/server/http"
	"github.com/untoldwind/gotrack/server/store"
)

var ServerCommand = cli.Command{
	Name:   "server",
	Usage:  "Start server",
	Action: runWithContext(serverCommand),
}

func serverCommand(ctx *cli.Context, runCtx *runContext) {
	provider, err := conntrack.NewProvider(runCtx.config.Provider, runCtx.logger)
	if err != nil {
		runCtx.logger.ErrorErr(err)
		return
	}

	store, err := store.NewStore(runCtx.config.Store, provider, runCtx.logger)
	if err != nil {
		runCtx.logger.ErrorErr(err)
		return
	}
	defer store.Stop()

	server := http.NewServer(runCtx.config.Server, store, runCtx.logger)

	if err := server.Start(); err != nil {
		runCtx.logger.ErrorErr(err)
		return
	}
	defer server.Stop()

	runCtx.handleSignals()
}

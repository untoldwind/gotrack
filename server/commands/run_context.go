package commands

import (
	"github.com/codegangsta/cli"
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/logging"
	"os"
	"os/signal"
	"syscall"
)

type runContext struct {
	config *config.Config
	logger logging.Logger
}

func newRunContext(ctx *cli.Context) (*runContext, error) {
	logger := logging.NewLogrusLogger(ctx)

	config, err := config.NewConfig(ctx.GlobalString("config-dir"), logger)
	if err != nil {
		logger.Errorf("Read config failed: %s", err.Error())
		return nil, err
	}

	return &runContext{
		config: config,
		logger: logger,
	}, nil
}

func (c *runContext) close() {
}

func (c *runContext) handleSignals() int {
	signalCh := make(chan os.Signal, 4)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)

	for {
		sig := <-signalCh

		c.logger.Infof("Caught signal: %v", sig)

		shutdown := false
		if sig == os.Interrupt || sig == syscall.SIGTERM {
			shutdown = true
		}

		if shutdown {
			return 0
		}
	}
	panic("Universe is in unknown state")
}

func runWithContext(command func(*cli.Context, *runContext)) func(*cli.Context) {
	return func(ctx *cli.Context) {
		runCtx, err := newRunContext(ctx)

		if err != nil {
			return
		}
		defer runCtx.close()

		command(ctx, runCtx)
	}
}

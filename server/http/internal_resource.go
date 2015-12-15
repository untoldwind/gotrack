package http

import (
	"github.com/untoldwind/gotrack/server/config"
	"github.com/untoldwind/gotrack/server/logging"
	"github.com/untoldwind/routing"
	"net/http"
	"runtime"
)

type internalResource struct {
	logger logging.Logger
}

type StatusVO struct {
	Version string `json:"version"`
}

type HealthVO struct {
	MaxProcs     int
	NumGoroutine int
	Memory       runtime.MemStats
}

func InternalRoutes(parent logging.Logger) routing.Matcher {
	logger := parent.WithContext(map[string]interface{}{"resource": "certificates"})
	resource := &internalResource{
		logger: logger,
	}
	return routing.PrefixSeq("/internal",
		routing.PrefixSeq("/status",
			routing.EndSeq(
				routing.GETFunc(wrap(resource.logger, resource.Status)),
				SendError(logger, MethodNotAllowed()),
			),
		),
		routing.PrefixSeq("/health",
			routing.EndSeq(
				routing.GETFunc(wrap(resource.logger, resource.Health)),
				SendError(logger, MethodNotAllowed()),
			),
		),
		routing.PrefixSeq("/gc",
			routing.EndSeq(
				routing.PUTFunc(wrap(resource.logger, resource.TriggerGC)),
				SendError(logger, MethodNotAllowed()),
			),
		),
	)
}

func (r *internalResource) Status(req *http.Request) (interface{}, error) {
	return &StatusVO{
		Version: config.Version(),
	}, nil
}

func (r *internalResource) Health(req *http.Request) (interface{}, error) {
	health := HealthVO{
		MaxProcs:     runtime.GOMAXPROCS(0),
		NumGoroutine: runtime.NumGoroutine(),
	}

	runtime.ReadMemStats(&health.Memory)

	return health, nil
}

func (r *internalResource) TriggerGC(req *http.Request) (interface{}, error) {
	runtime.GC()

	return nil, nil
}

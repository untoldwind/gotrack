package http

import (
	"github.com/untoldwind/gotrack/server/logging"
	"github.com/untoldwind/gotrack/server/store"
	"github.com/untoldwind/routing"
	"net/http"
)

type totalsResource struct {
	store  store.Store
	logger logging.Logger
}

func TotalsRoutes(store store.Store, parent logging.Logger) routing.Matcher {
	logger := parent.WithContext(map[string]interface{}{"resource": "data"})
	resource := &totalsResource{
		store:  store,
		logger: logger,
	}
	return routing.PrefixSeq("/totals",
		routing.PrefixSeq("/span",
			routing.EndSeq(
				routing.GETFunc(wrap(resource.logger, resource.GetTotalsSpan)),
				SendError(logger, MethodNotAllowed()),
			),
		),
		routing.PrefixSeq("/rates",
			routing.EndSeq(
				routing.GETFunc(wrap(resource.logger, resource.GetTotalsRates)),
				SendError(logger, MethodNotAllowed()),
			),
		),
	)
}

func (r *totalsResource) GetTotalsSpan(req *http.Request) (interface{}, error) {
	return r.store.TotalsSpan(), nil
}

func (r *totalsResource) GetTotalsRates(req *http.Request) (interface{}, error) {
	return r.store.TotalsRates(), nil
}

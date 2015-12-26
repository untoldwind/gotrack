package http

import (
	"github.com/untoldwind/gotrack/server/logging"
	"github.com/untoldwind/gotrack/server/store"
	"github.com/untoldwind/routing"
	"net/http"
)

type devicesResource struct {
	store  store.Store
	logger logging.Logger
}

func DeviceRoutes(store store.Store, parent logging.Logger) routing.Matcher {
	logger := parent.WithContext(map[string]interface{}{"resource": "devices"})
	resource := &devicesResource{
		store:  store,
		logger: logger,
	}
	return routing.PrefixSeq("/devices",
		routing.StringPart(func(deviceIp string) routing.Matcher {
			return routing.Sequence(
				routing.PrefixSeq("/span",
					routing.EndSeq(
						routing.GETFunc(wrap(resource.logger, resource.GetDeviceSpan(deviceIp))),
						SendError(logger, MethodNotAllowed()),
					),
				),
				routing.EndSeq(
					routing.GETFunc(wrap(resource.logger, resource.GetDeviceDetails(deviceIp))),
					SendError(logger, MethodNotAllowed()),
				),
			)
		}),
		routing.EndSeq(
			routing.GETFunc(wrap(resource.logger, resource.GetDevices)),
			SendError(logger, MethodNotAllowed()),
		),
	)
}

func (r *devicesResource) GetDevices(req *http.Request) (interface{}, error) {
	return r.store.Devices(), nil
}

func (r *devicesResource) GetDeviceDetails(deviceIp string) func(req *http.Request) (interface{}, error) {
	return func(req *http.Request) (interface{}, error) {
		deviceDetails := r.store.DeviceDetails(deviceIp)
		if deviceDetails == nil {
			return nil, NotFound()
		}
		return deviceDetails, nil
	}
}

func (r *devicesResource) GetDeviceSpan(deviceIp string) func(req *http.Request) (interface{}, error) {
	return func(req *http.Request) (interface{}, error) {
		return r.store.DeviceSpan(deviceIp), nil
	}
}

package http

import (
	"encoding/json"
	"github.com/untoldwind/gotrack/server/logging"
	"net/http"
	"strconv"
	"time"
)

func queryParamUint(req *http.Request, key string, defaultValue uint64) (uint64, error) {
	value := req.FormValue(key)
	if value == "" {
		return defaultValue, nil
	}
	return strconv.ParseUint(value, 10, 64)
}

func queryParamBool(req *http.Request, key string, defaultValue bool) (bool, error) {
	value := req.FormValue(key)
	if value == "" {
		return defaultValue, nil
	}
	return strconv.ParseBool(value)
}

func wrap(logger logging.Logger, handler func(req *http.Request) (interface{}, error)) func(resp http.ResponseWriter, req *http.Request) {
	f := func(resp http.ResponseWriter, req *http.Request) {
		logger := logger.WithContext(map[string]interface{}{"url": req.URL, "method": req.Method})
		start := time.Now()
		defer func() {
			logger.Debugf("http: Request (%v)", time.Now().Sub(start))
		}()
		obj, err := handler(req)
		if err != nil {
			encodeError(logger, resp, req, err)
			return
		}
		if obj != nil {
			var buf []byte
			buf, err = json.Marshal(obj)
			if err != nil {
				encodeError(logger, resp, req, err)
				return
			}
			resp.Header().Set("Content-Type", "application/json")
			resp.Write(buf)
		} else {
			resp.WriteHeader(http.StatusNoContent)
		}
	}
	return f
}

func wrapCreate(logger logging.Logger, handler func(req *http.Request) (interface{}, string, error)) func(resp http.ResponseWriter, req *http.Request) {
	f := func(resp http.ResponseWriter, req *http.Request) {
		logger := logger.WithContext(map[string]interface{}{"url": req.URL, "method": req.Method})
		start := time.Now()
		defer func() {
			logger.Debugf("http: Request (%v)", time.Now().Sub(start))
		}()
		obj, location, err := handler(req)
		if err != nil {
			encodeError(logger, resp, req, err)
			return
		}
		if obj != nil {
			var buf []byte
			buf, err = json.Marshal(obj)
			if err != nil {
				encodeError(logger, resp, req, err)
				return
			}
			resp.Header().Set("Location", location)
			resp.Header().Set("Content-Type", "application/json")
			resp.WriteHeader(http.StatusCreated)
			resp.Write(buf)
		} else {
			resp.WriteHeader(http.StatusNoContent)
		}
	}
	return f
}

func encodeError(logger logging.Logger, resp http.ResponseWriter, req *http.Request, err error) {
	logger.ErrorErr(err)
	var httpError *HTTPError
	switch err.(type) {
	case *HTTPError:
		httpError = err.(*HTTPError)
	default:
		httpError = &HTTPError{
			Code:    500,
			Message: err.Error(),
		}
	}
	resp.WriteHeader(httpError.Code)
	if errJson, errMarshal := json.Marshal(httpError); errMarshal == nil {
		resp.Write(errJson)
	} else {
		resp.Write([]byte(httpError.Message))
	}
}

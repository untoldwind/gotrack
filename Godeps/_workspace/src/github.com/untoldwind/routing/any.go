package routing

import "net/http"

func Any(handler http.Handler) Matcher {
	return func(remainingPath string, resp http.ResponseWriter, req *http.Request) bool {
		handler.ServeHTTP(resp, req)
		return true
	}
}

func AnyFunc(handler func(http.ResponseWriter, *http.Request)) Matcher {
	return Any(http.HandlerFunc(handler))
}

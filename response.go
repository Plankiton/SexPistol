package Sex

import (
	"fmt"
	"net/http"
)

// Response to make complete response with Cookies, Headers, and all http.ResponseWrite another features
type Response struct {
	http.ResponseWriter
	Headers *http.Header
	Body    []byte
	Status  int
}

func (self *Response) Header() http.Header {
	if self.Headers == nil {
		self.Headers = &http.Header{
			"Content-Type":   {"text/pain; charset=UTF-8"},
			"Content-Length": {"0"},
		}
	}

	return *self.Headers
}

// Response constructor function
func NewResponse() *Response {
	return new(Response)
}

// Function to run route func SexPistol
func runRoute(route_func interface{}, response Response, r Request) error {
	if route_func, ok := route_func.(func(http.ResponseWriter, *http.Request)); ok {
		route_func(response.ResponseWriter, &r.Request)
	}

	if route_func, ok := route_func.(func(Request) ([]byte, int)); ok {
		res, status := route_func(r)
		response.WriteHeader(status)
		response.Write(res)
	}

	if route_func, ok := route_func.(func(Request) []byte); ok {
		res := route_func(r)
		response.Write(res)
	}

	if route_func, ok := route_func.(func(Request) (string, int)); ok {
		res, status := route_func(r)
		response.WriteHeader(status)
		response.Write([]byte(res))
	}

	if route_func, ok := route_func.(func(Request) string); ok {
		res := route_func(r)
		response.Write([]byte(res))
	}

	if route_func, ok := route_func.(func(Request) (*Response, int)); ok {
		res, status := route_func(r)
		response.WriteHeader(status)
		response.Write(res.Body)
	}

	if route_func, ok := route_func.(func(Request) *Response); ok {
		res := route_func(r)
		response.Write(res.Body)
	}

	if route_func, ok := route_func.(func(Request) (Json, int)); ok {
		res, status := route_func(r)
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(status)
		response.Write(Jsonify(res))
	}

	if route_func, ok := route_func.(func(Request) Json); ok {
		res := route_func(r)
		response.Header().Set("Content-Type", "application/json")
		response.Write(Jsonify(res))
	}

	response_log_message := Fmt("%s %s %s %d: %s", r.Method, r.URL.Path, r.URL.RawQuery, response.Status, StatusText(response.Status))
	if response.Status >= 400 {
		RawLog(LogLevelError, false, response_log_message)
		return fmt.Errorf(response_log_message)
	} else if response.Status >= 300 {
		RawLog(LogLevelWarn, false, response_log_message)
	} else if response.Status >= 200 {
		RawLog(LogLevelInfo, false, response_log_message)
	}
	return nil
}

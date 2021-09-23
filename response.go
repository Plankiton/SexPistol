package Sex

import (
	"fmt"
	"net/http"
	"time"
)

// Response to make complete response with Cookies, Headers, and all http.ResponseWrite another features
type Response struct {
	http.ResponseWriter
	Headers *http.Header
	Body    []byte
	Status  int
	err     error
}

// Error provides Pistol last error
func (r *Response) Error() error { return r.err }

// SetErr sets Pistol last error
func (r *Response) SetErr(err error) { r.err = err }

// WriteHeader sets Response status code
func (r *Response) WriteHeader(status int) {
	r.SetStatus(status)
}

// SetBody sets Response body
func (r *Response) SetBody(v []byte) *Response {
	r.Body = v
	return r
}

// SetStatus sets Response status code
func (r *Response) SetStatus(code int) *Response {
	r.Status = 200
	return r
}

// SetCookie sets Response cookies
func (r *Response) SetCookie(key string, value string, expires time.Duration) *Response {
	cookie := &http.Cookie{
		Name:    key,
		Value:   value,
		Expires: time.Now().Add(expires),
	}
	http.SetCookie(r, cookie)

	return r
}

// Header returns response headers setter
func (r *Response) Header() http.Header {
	if r.Headers == nil {
		r.Headers = &http.Header{
			"Content-Type":   {"text/pain; charset=UTF-8"},
			"Content-Length": {"0"},
		}
	}

	return *r.Headers
}

// NewResponse provides new Response
func NewResponse() *Response {
	return new(Response)
}

// runRoute run sex route function
func runRoute(route_func interface{}, response Response, r Request) error {
	if route_func, ok := route_func.(func(Response, Request)); ok {
		route_func(response, r)
	}

	if route_func, ok := route_func.(func(*Response, *Request)); ok {
		route_func(&response, &r)
	}

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
		response.Write(Byteify(res))
	}

	if route_func, ok := route_func.(func(Request) Json); ok {
		res := route_func(r)
		response.Header().Set("Content-Type", "application/json")
		response.Write(Byteify(res))
	}

	response_log_message := Fmt("%s %s %s %d: %s", r.Method, r.URL.Path, r.URL.RawQuery, response.Status, StatusText(response.Status))
	if response.Status >= 400 {
		RawLog(LogLevelError, false, response_log_message)
		return fmt.Errorf(string(response.Body))
	} else if response.Status >= 300 {
		RawLog(LogLevelWarn, false, response_log_message)
	} else if response.Status >= 200 {
		RawLog(LogLevelInfo, false, response_log_message)
	}
	return nil
}

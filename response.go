package Sex

import "net/http"


// Response to make complete response with Cookies, Headers, and all http.ResponseWrite another features
type Response struct {
    http.ResponseWriter
    Headers *http.Header
    Body    []byte
    Status  int
}

func (self *Response) Header() http.Header {
    if self.Headers == nil {
        self.Headers = &http.Header {
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
func RunRouteFunc(route_func interface {}, w http.ResponseWriter, r Request) {
    if route_func, ok := route_func.(func(http.ResponseWriter, *http.Request)); ok {
        route_func(w, &r.Request)
    }

    if route_func, ok := route_func.(func(Request)([]byte, int)); ok {
        res, status := route_func(r)
        w.WriteHeader(status)
        w.Write(res)
    }

    if route_func, ok := route_func.(func(Request)([]byte)); ok {
        res := route_func(r)
        w.Write(res)
    }

    if route_func, ok := route_func.(func(Request)(string, int)); ok {
        res, status := route_func(r)
        w.WriteHeader(status)
        w.Write([]byte(res))
    }

    if route_func, ok := route_func.(func(Request)(string)); ok {
        res := route_func(r)
        w.Write([]byte(res))
    }

    if route_func, ok := route_func.(func(Request)(*Response, int)); ok {
        res, status := route_func(r)
        w.WriteHeader(status)
        w.Write(res.Body)
    }

    if route_func, ok := route_func.(func(Request)(*Response)); ok {
        res := route_func(r)
        w.Write(res.Body)
    }

    if route_func, ok := route_func.(func(Request)(Json, int)); ok {
        res, status := route_func(r)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(status)
        w.Write(Jsonify(res))
    }

    if route_func, ok := route_func.(func(Request)(Json)); ok {
        res := route_func(r)
        w.Header().Set("Content-Type", "application/json")
        w.Write(Jsonify(res))
    }
}

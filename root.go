package Sex

import (
	"os"
	re "regexp"

	"net/http"
)

// Function to make Sex.Pistol a http.Handler
func (pistol *Pistol) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if pistol == nil || pistol.ServeMux == nil {
        pistol = NewPistol()
    }

    pistol.ServeMux.ServeHTTP(w, r)
}

// root endpoint to run all Sex endpoints
func (pistol *Pistol) root(w http.ResponseWriter, r *http.Request) {
    body := Request {}

    path := r.URL.Path
    path = fixPath(r.URL.Path)

    root_path := fixPath(pistol.RootPath)
    if root_path == "/" {
        root_path = ""
    }

    response_log_message := Fmt("%s %s %s", r.Method, path, r.URL.RawQuery)
    for path_pattern, methods := range pistol.Routes {

        path_regex := re.MustCompile("^"+root_path+path_pattern+`{1}`)

        route_conf := pistol.RouteConfs[path_pattern]

        iroute_func := methods
        if methods, ok := methods.(Prop); ok {
            iroute_func = methods[r.Method]
        }
        isMatching := path_regex.MatchString(path)
        if os.Getenv("SEX_DEBUG") == "true" {
            RawLog("\033[32;1m[info] \033[00m", false, Fmt(`"%s" is "%s" ? %v`,path_pattern, path, isMatching))
        }
        if isMatching {

            if methods != nil {

                if iroute_func != nil {
                    body.Conf = route_conf.(Prop)
                    body.PathVars, _ = GetPathVars(route_conf.(Prop)["path-template"].(string), path)

                    body.Request = r
                    body.Writer = new(Response)
                    body.Writer.ResponseWriter = w

                    status_code := 200
                    if route_func, ok := iroute_func.(func(http.ResponseWriter, *http.Request)); ok {

                        route_func(body.Writer, r)
                        status_code = body.Writer.Status
                        w.WriteHeader(status_code)

                    } else
                    if route_func, ok := iroute_func.(func(Request)([]byte, int)); ok {

                        res, status := route_func(body)
                        if status != 0 {
                            status_code = status
                        }

                        w.WriteHeader(status)
                        w.Write(res)

                    } else
                    if route_func, ok := iroute_func.(func(Request)([]byte)); ok {

                        res := route_func(body)
                        w.Write(res)

                    } else
                    if route_func, ok := iroute_func.(func(Request)(string, int)); ok {

                        res, status := route_func(body)
                        if status != 0 {
                            status_code = status
                            w.WriteHeader(status)
                        }

                        w.Write([]byte(res))

                    } else
                    if route_func, ok := iroute_func.(func(Request)(string)); ok {

                        res := route_func(body)
                        w.Write([]byte(res))

                    } else
                    if route_func, ok := iroute_func.(func(Request)(*Response, int)); ok {

                        res, status := route_func(body)
                        if status == 0 {
                            if res.Status != 0 {
                                status = res.Status
                            } else {
                                status = 200
                            }
                        }
                        status_code = status

                        w.WriteHeader(status)
                        w.Write(res.Body)

                    } else
                    if route_func, ok := iroute_func.(func(Request)(*Response)); ok {

                        res := route_func(body)
                        w.Write(res.Body)

                    } else
                    if route_func, ok := iroute_func.(func(Request)(Json, int)); ok {

                        res, status := route_func(body)
                        if status != 0 {
                            status = 200
                            status_code = status
                            w.WriteHeader(status)
                        }

                        w.Header().Set("Content-Type", "application/json")
                        w.Write(Jsonify(res))

                    } else
                    if route_func, ok := iroute_func.(func(Request)(Json)); ok {

                        res := route_func(body)
                        w.Header().Set("Content-Type", "application/json")
                        w.Write(Jsonify(res))

                    }

                    response_log_message += Fmt("%d: %s", status_code, http.StatusText(status_code))
                    if status_code >= 400 {
                        RawLog("\033[31;1m[erro] \033[00m", false, response_log_message)
                    } else
                    if status_code >= 300 {
                        RawLog("\033[33;1m[warn] \033[00m", false, response_log_message)
                    } else {
                        RawLog("\033[32;1m[info] \033[00m", false, response_log_message)
                    }

                    return

                }
            }

            status_code := 405
            response_log_message += Fmt("%d: %s", status_code, http.StatusText(status_code))
            RawLog("\033[31;1m[erro] \033[00m", false, response_log_message)
            w.WriteHeader(status_code)
            return
        }
    }

    status_code := 404
    response_log_message += Fmt("%d: %s", status_code, http.StatusText(status_code))
    RawLog("\033[31;1m[erro] \033[00m", false, response_log_message)
}

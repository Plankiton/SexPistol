package Sex
import (
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

        if path_regex.MatchString(path) {

            if methods != nil {

                if iroute_func != nil {
                    body.Conf = route_conf.(Prop)
                    body.PathVars, _ = GetPathVars(route_conf.(Prop)["path-template"].(string), path)

                    body.Request = r
                    body.Writer = new(Response)
                    body.Writer.ResponseWriter = w

                    sc := 200
                    if route_func, ok := iroute_func.(func(http.ResponseWriter, *http.Request)); ok {

                        route_func(w, r)

                    } else
                    if route_func, ok := iroute_func.(func(Request)([]byte, int)); ok {

                        res, status := route_func(body)
                        if status == 0 {
                            status = 200
                        }
                        sc = status

                        w.WriteHeader(status)
                        w.Write(res)

                    } else
                    if route_func, ok := iroute_func.(func(Request)([]byte)); ok {

                        res := route_func(body)
                        status := 200
                        sc = status

                        w.WriteHeader(status)
                        w.Write(res)

                    } else
                    if route_func, ok := iroute_func.(func(Request)(string, int)); ok {

                        res, status := route_func(body)
                        if status == 0 {
                            status = 200
                        }
                        sc = status

                        w.WriteHeader(status)
                        w.Write([]byte(res))

                        return
                    } else
                    if route_func, ok := iroute_func.(func(Request)(string)); ok {

                        res := route_func(body)
                        status := 200
                        sc = status

                        w.WriteHeader(status)
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
                        sc = status

                        w.WriteHeader(status)
                        w.Write(res.Body)

                    } else
                    if route_func, ok := iroute_func.(func(Request)(*Response)); ok {

                        res := route_func(body)
                        if res.Status == 0 {
                            res.Status = 200
                        }
                        sc = res.Status

                        w.WriteHeader(res.Status)
                        w.Write(res.Body)

                    } else
                    if route_func, ok := iroute_func.(func(Request)(Json, int)); ok {

                        res, status := route_func(body)
                        if status == 0 {
                            status = 200
                        }
                        sc = status

                        w.Header().Set("Content-Type", "application/json")
                        w.WriteHeader(status)
                        w.Write(Jsonify(res))

                    } else
                    if route_func, ok := iroute_func.(func(Request)(Json)); ok {

                        res := route_func(body)
                        status := 200
                        sc = status

                        w.Header().Set("Content-Type", "application/json")
                        w.WriteHeader(status)
                        w.Write(Jsonify(res))

                    }

                    response_log_message += Fmt(" -> %d: %s", sc, http.StatusText(sc))
                    if sc >= 400 {
                        Err(response_log_message)
                    } else
                    if sc >= 300 {
                        War(response_log_message)
                    } else {
                        Log(response_log_message)
                    }

                    return

                }
            }

            response_log_message += Fmt(" -> %d: %s", 405, http.StatusText(405))
            Err(response_log_message)
            w.WriteHeader(405)
            return
        }
    }

    response_log_message += Fmt(" -> %d: %s", 404, http.StatusText(404))
    Err(response_log_message)
    w.WriteHeader(404)
}

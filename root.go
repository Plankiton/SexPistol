package Sex
import (
    str "strings"
    re "regexp"

    "net/http"

    "strconv"
    "bytes"

    "github.com/Showmax/go-fqdn"
    "github.com/rs/cors"
)

func (router *Pistol) RootRoute(w http.ResponseWriter, r *http.Request) {
    body := Request {}

    path := r.URL.Path
    if path != "/" {
        path = fixPath(r.URL.Path)
    }

    Log(r.Method, path, r.URL.RawQuery)

    for path_pattern, methods := range router.Routes {

        path_regex := re.MustCompile(path_pattern)

        route_conf := router.RouteConfs[path_pattern]
        route_func := methods[r.Method]

        if path_regex.MatchString(path) {

            if methods != nil {

                if route_func != nil {
                    body.Conf = route_conf
                    body.PathVars, _ = GetPathVars(route_conf["path-template"].(string), path)

                    body.Request = r
                    body.Writer = new(Response)
                    body.Writer.ResponseWriter = w

                    sb := ""
                    sc := 200
                    if isRawFunc(route_func) {

                        res, status := route_func.(func(Request)([]byte, int))(body)
                        if status == 0 {
                            status = 200
                        }
                        sc = status
                        sb = string(res)

                        w.WriteHeader(status)
                        w.Write(res)

                    } else
                    if isStrFunc(route_func) {

                        res, status := route_func.(func(Request)(string, int))(body)
                        if status == 0 {
                            status = 200
                        }
                        sc = status
                        sb = res.Body

                        w.WriteHeader(status)
                        w.Write([]byte(res))

                    } else
                    if isResFunc(route_func) {

                        res, status := route_func.(func(Request)(*Response, int))(body)
                        if status == 0 {
                            if res.Status != 0 {
                                status = res.Status
                            } else {
                                status = 200
                            }
                        }
                        sc = status
                        sb = string(res.Body)

                        w.WriteHeader(status)
                        w.Write(res.Body)

                    } else
                    if isPureResFunc(route_func) {

                        res := route_func.(func(Request)(*Response))(body)
                        if res.Status == 0 {
                            res.Status = 200
                        }
                        sc = res.Status
                        sb = string(res.Body)

                        w.WriteHeader(res.Status)
                        w.Write(res.Body)

                    } else {

                        res, status := route_func.(func(Request)(interface{}, int))(body)
                        if status == 0 {
                            status = 200
                        }
                        sc = status
                        sb = string(Jsonify(res))

                        w.Header().Set("Content-Type", "application/json")
                        w.WriteHeader(status)
                        w.Write(Jsonify(res))
                    }

                    msg := Fmt("%d -> %s, %s", sc, http.StatusText(sc), sb)
                    if sc != 200 {
                        Err(msg)
                    }

                    return
                }

                Err("Route not found")
                w.WriteHeader(404)
                w.Write(Jsonify(Bullet {
                    Message: "Route not found",
                    Type:    "Error",
                }))
                return
            }

            Err("Method not allowed")
            w.WriteHeader(405)
            w.Write(Jsonify(Bullet {
                Message: "Method not allowed",
                Type:    "Error",
            }))
            return
        }
    }

    Err("Route not found")
    w.WriteHeader(404)
    w.Write(Jsonify(Bullet {
        Message: "Route not found",
        Type:    "Error",
    }))
}

func (router *Pistol) Run(a...interface{}) error {
    port := 8000
    path := "/"

    if a != nil {
        for _, v := range a {
            if ValidateData(v, GenericString) {
                path = v.(string)
            }

            if ValidateData(v, GenericInt) {
                port = v.(int)
            }
        }
    }

    router.Mux = http.NewServeMux()

    router.RootPath = path
    router.Mux.HandleFunc(path, router.RootRoute)

    host, err := fqdn.FqdnHostname()
    if err != nil {
        host = "localhost"
    }

    Log(Fmt("Running Sex Pistol server at %s:%d%s", host, port, path))
    handler := cors.Default().Handler(router.Mux)
    return http.ListenAndServe(Fmt(":%d", port), handler)
}

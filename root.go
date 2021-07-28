package Sex
import (
    re "regexp"

    "net/http"
)

func (router *Pistol) ROOT(w http.ResponseWriter, r *http.Request) {
    body := Request {}

    path := r.URL.Path
    path = fixPath(r.URL.Path)

    Log(r.Method, path, r.URL.RawQuery)

    for path_pattern, methods := range router.Routes {

        path_regex := re.MustCompile("^"+fixPath(router.RootPath)+path_pattern+`{1}`)

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
                    if isRawFuncNoStatus(route_func) {

                        res := route_func.(func(Request)([]byte))(body)
                        status := 200
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
                        sb = res

                        w.WriteHeader(status)
                        w.Write([]byte(res))

                        return
                    } else
                    if isStrFuncNoStatus(route_func) {

                        res := route_func.(func(Request)(string))(body)
                        status := 200
                        sc = status
                        sb = res

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
                    if isResFuncNoStatus(route_func) {

                        res := route_func.(func(Request)(*Response))(body)
                        if res.Status == 0 {
                            res.Status = 200
                        }
                        sc = res.Status
                        sb = string(res.Body)

                        w.WriteHeader(res.Status)
                        w.Write(res.Body)

                    } else
                    if isJsonFunc(route_func) {

                        res, status := route_func.(func(Request)(Json, int))(body)
                        if status == 0 {
                            status = 200
                        }
                        sc = status
                        sb = string(Jsonify(res))

                        w.Header().Set("Content-Type", "application/json")
                        w.WriteHeader(status)
                        w.Write(Jsonify(res))

                    } else
                    if isJsonFuncNoStatus(route_func) {

                        res := route_func.(func(Request)(Json))(body)
                        status := 200
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

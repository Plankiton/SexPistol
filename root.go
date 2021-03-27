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

    end := "\n\t-> Body: <"+r.Header.Get("Content-Type")+">"
    if str.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
        l, _ := strconv.Atoi(r.Header.Get("Content-Lenght"))
        r.ParseMultipartForm(int64(l))
    } else
    if str.HasPrefix(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
        r.ParseForm()
    } else {
        raw_body := new(bytes.Buffer)
        raw_body.ReadFrom(r.Body)
        end = "\n\t-> Body: "+ raw_body.String()
        if raw_body.Len() == 0 {
            end = ""
        }

        if FromJson(raw_body.Bytes(), &body.Body) != nil {
            body.Body = new(bytes.Buffer)
            body.Body.(*bytes.Buffer).ReadFrom(raw_body)
        }
    }

    path := r.URL.Path
    if path != "/" {
        path = fixPath(r.URL.Path)
    }

    Log(r.Method, path, r.URL.RawQuery, end)

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

                    sc := 200
                    if isRawFunc(route_func) {

                        res, status := route_func.(func(Request)([]byte, int))(body)
                        sc = status

                        w.WriteHeader(status)
                        w.Write(res)

                    } else
                    if isStrFunc(route_func) {

                        res, status := route_func.(func(Request)(string, int))(body)
                        sc = status

                        w.WriteHeader(status)
                        w.Write([]byte(res))

                    } else
                    if isResFunc(route_func) {

                        res, status := route_func.(func(Request)(*Response, int))(body)
                        sc = status

                        w.WriteHeader(status)
                        w.Write(res.Body.([]byte))

                    } else
                    if isFunc(route_func) {

                        res, status := route_func.(func(Request)(interface{}, int))(body)
                        sc = status

                        w.WriteHeader(status)
                        w.Write(Jsonify(res))

                    } else
                    if isPureResFunc(route_func) {

                        res := route_func.(func(Request)(*Response))(body)
                        sc = res.Status

                        w.WriteHeader(res.Status)
                        w.Write(Jsonify(res.Body))

                    }

                    msg := Fmt("%d -> %s", sc, http.StatusText(sc))
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

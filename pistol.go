package Sex

import (
    "github.com/Showmax/go-fqdn"
    "github.com/rs/cors"
    "net/http"
    "strings"
    "log"
    "os"
)

type Pistol struct {
    *http.ServeMux
    RootPath        string
    RouteConfs      Dict
    Routes          Dict
    RawRoutes       []string
    Auth            bool
}

func NewPistol() *Pistol {
    pistol := new(Pistol)
    pistol.ServeMux = http.NewServeMux()
    pistol.AddRaw("/", pistol.ROOT)
    if logger == nil {
        logger = log.New(os.Stderr, "\r\n", log.LstdFlags)
    }

    return pistol
}

func (pistol *Pistol) Add(path string, route interface {}, methods ...string) *Pistol {
    if f, ok := route.(httpRawFunc); ok {
        return pistol.AddRaw(path, f, methods...)
    }

    path = fixPath(path)
    root_path := fixPath(pistol.RootPath)
    if (path != root_path) {
        path = root_path + path
    }

    path_pattern := GetPathPattern(path)

    if len(pistol.RouteConfs) == 0 {
        pistol.RouteConfs = make(Dict)
    }
    if len(pistol.Routes) == 0 {
        pistol.Routes = make(Dict)
    }

    if _, exist := pistol.RouteConfs[path_pattern]; !exist {
        pistol.RouteConfs[path_pattern] = make(Prop)
    }
    conf := Prop{}
    conf["path-template"] = path
    pistol.RouteConfs[path_pattern] = conf

    if methods != nil {
        if _, exist := pistol.Routes[path_pattern]; !exist {
            pistol.Routes[path_pattern] = make(Prop)
        }
        for _, method := range methods {
            method = strings.ToUpper(method)
            pistol.Routes[path_pattern].(Prop)[method] = route
        }
    } else {
        pistol.Routes[path_pattern] = route
    }

    return pistol
}

func (pistol *Pistol) AddRaw(path string, f func(http.ResponseWriter, *http.Request), methods...string) (*Pistol) {
    if pistol.ServeMux == nil {
        pistol.ServeMux = http.NewServeMux()
    }

    path = fixPath(path)
    pistol.HandleFunc(path, func(w http.ResponseWriter, r *http.Request){
        run_request := true

        if methods != nil {
            run_request = false
            for _, m := range methods {
                if strings.ToUpper(m) != r.Method {
                    run_request = true
                    break
                }
            }
        }

        if run_request {
            Log(r.Method, path, r.URL.RawQuery)
            f(w, r)
        }
    })

    pistol.RawRoutes = append(pistol.RawRoutes, path)
    return pistol
}

func (pistol *Pistol) Run(a ...interface{}) error {
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

    pistol.RootPath = path
    host, err := fqdn.FqdnHostname()
    if err != nil {
        host = "localhost"
    }

    Log(Fmt("Running Sex Pistol server at %s:%d%s", host, port, path))
    if GetEnv("SEX_DEBUG", "false") != "false" {
        for path, methods := range pistol.Routes {
            methods_str := ""
            if methods, ok := methods.(Prop); ok {
                for method := range methods {
                    methods_str += Fmt("%s ", method)
                }
            } else {
                methods_str = "ALL METHODS"
            }

            Log(Fmt("%s <- %s", pistol.RouteConfs[path].(Prop)["path-template"], methods_str))
        }

        for _, path := range pistol.RawRoutes {
            Log(Fmt("%s <- %s", path, "ALL METHODS"))
        }
    }

    handler := cors.Default().Handler(pistol)
    return http.ListenAndServe(Fmt(":%d", port), handler)
}

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
    router := new(Pistol)
    router.ServeMux = http.NewServeMux()
    if logger == nil {
        logger = log.New(os.Stderr, "\r\n", log.LstdFlags)
    }

    return router
}

func (router *Pistol) Add(path string, route interface {}, methods ...string) *Pistol {
    if methods == nil {
        methods = []string{"GET"}
    }

    path = fixPath(path)
    root_path := fixPath(router.RootPath)
    if (path != root_path) {
        path = root_path + path
    }

    path_pattern := GetPathPattern(path)

    if len(router.RouteConfs) == 0 {
        router.RouteConfs = make(Dict)
    }
    if len(router.Routes) == 0 {
        router.Routes = make(Dict)
    }

    if _, exist := router.Routes[path_pattern]; !exist {
        router.Routes[path_pattern] = make(Prop)
    }
    if _, exist := router.RouteConfs[path_pattern]; !exist {
        router.RouteConfs[path_pattern] = make(Prop)
    }

    conf := Prop{}
    conf["path-template"] = path
    router.RouteConfs[path_pattern] = conf

    for _, method := range methods {
        method = strings.ToUpper(method)
        router.Routes[path_pattern][method] = route
    }

    return router
}

func (router *Pistol) AddRaw(path string, f func(http.ResponseWriter, *http.Request)) (*Pistol) {
    if router.ServeMux == nil {
        router.ServeMux = http.NewServeMux()
    }

    path = fixPath(path)
    router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request){
        Log(r.Method, path, r.URL.RawQuery)
        f(w, r)
    })

    router.RawRoutes = append(router.RawRoutes, path)
    return router
}

func (router *Pistol) Run(a ...interface{}) error {
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

    router.RootPath = path
    host, err := fqdn.FqdnHostname()
    if err != nil {
        host = "localhost"
    }

    Log(Fmt("Running Sex Pistol server at %s:%d%s", host, port, path))
    if GetEnv("SEX_DEBUG", "false") != "false" {
        for path, methods := range router.Routes {
            methods_str := ""
            for method := range methods {
                methods_str += Fmt("%s ", method)
            }

            Log(Fmt("%s <- %s", router.RouteConfs[path]["path-template"], methods_str))
        }

        for _, path := range router.RawRoutes {
            Log(Fmt("%s <- %s", path, "ALL METHODS"))
        }
    }

    router.AddRaw("/", router.ROOT)
    handler := cors.Default().Handler(router)
    return http.ListenAndServe(Fmt(":%d", port), handler)
}

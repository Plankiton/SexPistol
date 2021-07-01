package Sex

import (
    "net/http"
    "strings"
    "log"
    "os"
)

type Pistol struct {
    RootPath        string
    RouteConfs      Dict
    Routes          Dict
    Auth            bool
    Mux             *http.ServeMux
}

func NewPistol() *Pistol {
    router := new(Pistol)
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

package Sex

import (
    "net/http"
    "strings"
)

type Pistol struct {
    RootPath        string
    RouteConfs      RouteConfDict
    Routes          RouteDict
    Auth            bool
    Mux             *http.ServeMux
}

func (router *Pistol) Add(path string, route interface {}, methods ...string) *Pistol {
    if methods == nil {
        methods = []string{"GET"}
    }

    path = fixPath(path)
    path = router.RootPath + path
    path_pattern := GetPathPattern(path)

    if len(router.RouteConfs) == 0 {
        router.RouteConfs = make(RouteConfDict)
    }
    if len(router.Routes) == 0 {
        router.Routes = make(RouteDict)
    }

    if _, exist := router.Routes[path_pattern]; !exist {
        router.Routes[path_pattern] = make(Route)
    }
    if _, exist := router.RouteConfs[path_pattern]; !exist {
        router.RouteConfs[path_pattern] = make(RouteConf)
    }

    conf := RouteConf{}
    conf["path-template"] = path
    router.RouteConfs[path_pattern] = conf

    for _, method := range methods {
        method = strings.ToUpper(method)
        router.Routes[path_pattern][method] = route
    }

    return router
}

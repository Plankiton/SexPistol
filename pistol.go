package Sex

import (
    "net/http"
    "strings"
)

type Pistol struct {
    RootPath        string
    RouteConfs      routeConfDict
    Routes          routeDict
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
        router.RouteConfs = make(routeConfDict)
    }
    if len(router.Routes) == 0 {
        router.Routes = make(routeDict)
    }

    if _, exist := router.Routes[path_pattern]; !exist {
        router.Routes[path_pattern] = make(route_t)
    }
    if _, exist := router.RouteConfs[path_pattern]; !exist {
        router.RouteConfs[path_pattern] = make(routeConf)
    }

    conf := routeConf{}
    conf["path-template"] = path
    router.RouteConfs[path_pattern] = conf

    for _, method := range methods {
        method = strings.ToUpper(method)
        router.Routes[path_pattern][method] = route
    }

    return router
}

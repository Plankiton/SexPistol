package api

import (
    "encoding/json"
    "net/http"
    "strings"
    "bytes"
    "fmt"

    "gorm.io/gorm"
)

type RouteFunc func(r *http.Request) (Response, int)

type Route map[string] RouteFunc
type RouteDict map[string] Route

type RouteConf map[string] interface{}
type RouteConfDict map[string] RouteConf

type API struct {
    RootPath string
    RouteConfs RouteConfDict
    Routes   RouteDict
    Database *gorm.DB
}

func (router *API) Add(method string, path string, conf RouteConf, route RouteFunc) *API {
    method = strings.ToUpper(method)
    if path[len(path)-1] != '/' {
        path += "/"
    }

    path = router.RootPath + path

    if len(router.RouteConfs) == 0 {
        router.RouteConfs = make(RouteConfDict)
    }
    if len(router.Routes) == 0 {
        router.Routes = make(RouteDict)
    }

    if _, exist := router.Routes[path]; !exist {
        router.Routes[path] = make(Route)
    }

    router.RouteConfs[path] = conf
    router.Routes[path][method] = route
    return router
}

func (router *API) RootRoute(w http.ResponseWriter, r *http.Request) {
    body := Request {}
    raw_body := new(bytes.Buffer)
    raw_body.ReadFrom(r.Body)

    json.NewDecoder(r.Body).Decode(&body)
    path := r.URL.Path
    if path[len(path)-1] != '/' {
        path += "/"
    }

    Log(r.Method, path, r.URL.RawQuery, "\n\t-> Body: ", raw_body)
    if router.Routes[path] != nil{
        if router.Routes[path][r.Method] != nil{

            if router.RouteConfs[path]["need-auth"] == true {
                token := Token { ID: body.Token }
                if !token.Verify() {
                    Err("Authentication fail, permission denied")
                    w.WriteHeader(405)
                    json.NewEncoder(w).Encode(Response {
                        Message: "Authentication fail, permission denied",
                        Type:    "Error",
                    })
                    return
                }

                Log("Authentication sucessfull")
            }

            //             Route        [   "/"    ][ "GET"  ](r)
            res, status := router.Routes[path][r.Method](r)
            w.WriteHeader(status)
            json.NewEncoder(w).Encode(
                res,
            )
            return

        }

        Err("Method not allowed")
        w.WriteHeader(405)
        json.NewEncoder(w).Encode(Response {
            Message: "Method not allowed",
            Type:    "Error",
        })
        return
    }

    Err("Route not found")
    w.WriteHeader(404)
    json.NewEncoder(w).Encode(Response {
        Message: "Route not found",
        Type:    "Error",
    })
}

func (router *API) Run(path string, port uint) {
    router.RootPath = path
    http.HandleFunc(path, router.RootRoute)
    Err(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

package api

import (
    "encoding/json"
    "net/http"
    "fmt"

    "strings"
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

func (router *API) Add(method string, path string, route RouteFunc, conf RouteConf) *API {
    method = strings.ToUpper(method)
    if path[0] == '/' {
        path = path[1:]
    }
    if path[len(path)-1] != '/' {
        path += "/"
    }

    router.RouteConfs[router.RootPath + path] = conf
    router.Routes[router.RootPath + path][method] = route
    return router
}

func (router *API) RootRoute(w http.ResponseWriter, r *http.Request) {
    body := Request {}
    json.NewDecoder(r.Body).Decode(&body)

    path := r.URL.Path
    if path[len(path)-1] != '/' {
        path += "/"
    }

    Log(r.Method, path, r.URL.RawQuery, "\n\t-> Body: ", GetPrototype(body))
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
            res, status := router.Routes[r.URL.Path][r.Method](r)
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

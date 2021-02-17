package api

import (
    "encoding/json"
    "net/http"
    "strings"
    "bytes"
    "fmt"

    "github.com/gorilla/mux"
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
    path_pattern, _ := mux.NewRouter().HandleFunc(path, func(w http.ResponseWriter, r *http.Request){}).GetPathRegexp()
    path_regex := ReCompile(path_pattern)
    print("\n\n",path_pattern," == ", path, " -> ",  path_regex.MatchString(path), "\n\n")

    if len(router.RouteConfs) == 0 {
        router.RouteConfs = make(RouteConfDict)
    }
    if len(router.Routes) == 0 {
        router.Routes = make(RouteDict)
    }

    if _, exist := router.Routes[path_pattern]; !exist {
        router.Routes[path_pattern] = make(Route)
    }

    router.RouteConfs[path_pattern] = conf
    router.Routes[path_pattern][method] = route
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

    end := "\n\t-> Body: "+raw_body.String()
    if raw_body.Len() == 0 {
        end = ""
    }

    Log(r.Method, path, r.URL.RawQuery, end)

    for path_pattern, methods := range router.Routes {

        path_regex := ReCompile(path_pattern)

        route_conf := router.RouteConfs[path_pattern]
        route_func := methods[r.Method]


        if path_regex.MatchString(path) {

            if methods != nil{

                if route_func != nil {
                    if route_conf["need-auth"] == true {
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

                    res, status := route_func(r)

                    w.WriteHeader(status)
                    json.NewEncoder(w).Encode(
                        res,
                    )
                    return
                }

                Err("Route not found")
                w.WriteHeader(404)
                json.NewEncoder(w).Encode(Response {
                    Message: "Route not found",
                    Type:    "Error",
                })
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

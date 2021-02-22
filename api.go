package api

import (
	"encoding/json"
	"net/http"
	"strings"
  "strconv"
  "bytes"
	"fmt"

	"gorm.io/gorm"
)

type Response struct {
    Message   string             `json:"message,omitempty"`
    Type      string             `json:"type,omitempty"`
    Data      interface{}        `json:"data,omitempty"`
}

type Request  struct {
    Token     string               `json:"auth,omitempty"`
    Data      interface{}          `json:"data,omitempty"`
    PathVars  map[string]string
    Conf      RouteConf
}

type Route map[string] RouteFunc
type RouteDict map[string] Route

type RouteConf map[string] interface{}
type RouteConfDict map[string] RouteConf
type RouteFunc func(r Request) (Response, int)

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

    if conf == nil {
        conf = RouteConf {}
    }

    conf["path-template"] = path
    router.RouteConfs[path_pattern] = conf

    router.Routes[path_pattern][method] = route

    Log("Adding", path, "route to", router.RootPath, "router")
    return router
}

func (router *API) RootRoute(w http.ResponseWriter, r *http.Request) {
    body := Request {}

    raw_body := new(bytes.Buffer)
    var parse_err error

    end := ""
    if r.Header.Get("Content-Type") == "application/json" {
        end = "\n\t-> Body: "+ raw_body.String()
        raw_body.ReadFrom(r.Body)
        if raw_body.Len() == 0 {
            end = ""
        }

        parse_err = json.Unmarshal(raw_body.Bytes(), &body)
    }

    if strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
        l, _ := strconv.Atoi(r.Header.Get("Content-Lenght"))
        r.ParseMultipartForm(int64(l))
        _, f, err := r.FormFile("data")
        body.Data = r.MultipartForm

        if err == nil {
            end = "\n\t-> Body: <File:"+r.FormValue("description")+" <- \""+f.Filename+"\">"
        } else {
            end = ""
        }

    }

    path := r.URL.Path
    if path[len(path)-1] != '/' {
        path += "/"
    }

    Log(r.Method, path, r.URL.RawQuery, end)

    if parse_err != nil {
            Err("Bad request, json parsing error")
            w.WriteHeader(400)
            json.NewEncoder(w).Encode(Response {
                Message: "Bad request, json parsing error",
                Type:    "Error",
            })
            return
    }

    for path_pattern, methods := range router.Routes {

        path_regex := ReCompile(path_pattern)

        route_conf := router.RouteConfs[path_pattern]
        route_func := methods[r.Method]

        if path_regex.MatchString(path) {

            if methods != nil{

                if route_func != nil {
                    if route_conf["need-auth"] == false {
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

                    body.Conf = route_conf
                    body.PathVars, _ = GetPathVars(route_conf["path-template"].(string), path)
                    res, status := route_func(body)

                    if status == 0 {status = 200}

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

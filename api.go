package sex

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

  "time"
	"gorm.io/gorm"
)

type Response struct {
    Message   string             `json:"message,omitempty"`
    Type      string             `json:"type,omitempty"`
    Data      interface{}        `json:"data,omitempty"`
    Cookies   []*http.Cookie     `json:"-"`
}

type Request  struct {
    Data      interface{}          `json:"data,omitempty"`
    Token     string
    PathVars  map[string]string
    Conf      RouteConf
}

func (self *Response) AddCookie(key string, value string, expires time.Duration) {
    self.Cookies = append(self.Cookies, &http.Cookie {
        Name: key,
        Value: value,
        Expires: time.Now().Add(expires),
    })
}

func (self *Response) SetCookies(w http.ResponseWriter) {
    for _, cookie := range self.Cookies {
        http.SetCookie(w, cookie)
    }
}

type Route map[string] interface{}
type RouteDict map[string] Route

type RouteConf map[string] interface{}
type RouteConfDict map[string] RouteConf

type RouteFunc func(r Request) (Response, int)
type RawRouteFunc func(r Request) ([]byte, int)

type Pistol struct {
    RootPath        string
    RouteConfs      RouteConfDict
    Routes          RouteDict
    Auth            bool
    Database        *gorm.DB
}

func (router *Pistol) Add(method string, path string, conf RouteConf, route interface {}) *Pistol {
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

    Log("Adding", method, path, "route to", router.RootPath, "router")
    return router
}

func (router *Pistol) RootRoute(w http.ResponseWriter, r *http.Request) {
    body := Request {}

    end := ""

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

    } else if strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {

        raw_body := new(bytes.Buffer)
        raw_body.ReadFrom(r.Body)
        end = "\n\t-> Body: "+ raw_body.String()
        if raw_body.Len() == 0 {
            end = ""
        }

        if json.Unmarshal(raw_body.Bytes(), &body.Data) != nil {
            Err("Bad request, json parsing error")
            w.WriteHeader(400)
            json.NewEncoder(w).Encode(Response {
                Message: "Bad request, json parsing error",
                Type:    "Error",
            })
            return
        }

    } else {
        body.Data = new(bytes.Buffer)
        body.Data.(*bytes.Buffer).ReadFrom(r.Body)
    }


    path := r.URL.Path
    if path[len(path)-1] != '/' {
        path += "/"
    }

    ip := r.Header.Get("x-forwarded-for")
    if ip == "" {
        ip = strings.Split(r.RemoteAddr, ":")[0]
    }
    Log(r.Method, path, r.URL.RawQuery, end)

    for path_pattern, methods := range router.Routes {

        path_regex := ReCompile(path_pattern)

        route_conf := router.RouteConfs[path_pattern]
        route_func := methods[r.Method]

        if path_regex.MatchString(path) {

            if methods != nil{

                if route_func != nil {
                    if router.Auth {
                        if _, e := route_conf["need-auth"];
                        !e || route_conf["need-auth"] != false {
                            auth_token := r.Header.Get("Authorization")

                            token := Token { ID: auth_token }
                            body.Token = auth_token
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
                    }

                    body.Conf = route_conf
                    body.PathVars, _ = GetPathVars(route_conf["path-template"].(string), path)

                    body.Conf["headers"] = r.Header
                    r.ParseForm()
                    body.Conf["form"] = r.Form
                    body.Conf["query"] = r.URL.Query()
                    body.Conf["raw_response_writer"] = w
                    body.Conf["raw_request"] = r

                    if IsFunc(route_func) {
                        res, status := route_func.(func(Request)(Response,int))(body)

                        if status == 0 {status = 200}

                        w.WriteHeader(status)
                        json.NewEncoder(w).Encode(
                            res,
                        )

                        res.SetCookies(w)
                        return
                    }

                    if IsRawFunc(route_func) {
                        res, status := route_func.(func(Request)([]byte,int))(body)
                        if status == 0 {status = 200}
                        w.WriteHeader(status)
                        w.Write(res)
                        return
                    }

                    Err("Invalid route for ", path)
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

func (router *Pistol) Run(path string, port uint) {
    router.RootPath = path
    http.HandleFunc(path, router.RootRoute)
    Err(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

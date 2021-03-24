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
  "github.com/rs/cors"
)

type Response struct {
    Message   string             `json:"message,omitempty"`
    Type      string             `json:"type,omitempty"`
    Data      interface{}        `json:"data,omitempty"`
}

type Request  struct {
    Data      interface{}          `json:"data,omitempty"`
    Token     string
    PathVars  map[string]string
    Conf      RouteConf
    Reader    *http.Request
    Writer    http.ResponseWriter
}

func (self *Response) SetCookie(key string, value string, expires time.Duration, request Request) {
    cookie := &http.Cookie {
        Name: key,
        Value: value,
        Expires: time.Now().Add(expires),
    }
    http.SetCookie(request.Writer, cookie)
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
    Mux             *http.ServeMux
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
    body.Writer = w
    body.Reader = r

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

    }
    raw_body := new(bytes.Buffer)
    raw_body.ReadFrom(r.Body)
    end = "\n\t-> Body: "+ raw_body.String()
    if raw_body.Len() == 0 {
        end = ""
    }

    if json.Unmarshal(raw_body.Bytes(), &body.Data) != nil {
        body.Data = new(bytes.Buffer)
        body.Data.(*bytes.Buffer).ReadFrom(raw_body)
    }


    path := r.URL.Path
    if path[len(path)-1] != '/' {
        path += "/"
    }

    auth_token := r.Header.Get("Authorization")
    body.Token = auth_token
    token := Token { ID: auth_token }
    token_info := ""
    if auth_token != "" {
        token_info = ToLabel(token.ID, "AuthToken")
    }

    Log(r.Method, path, r.URL.RawQuery, token_info, end)
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
                            if !token.Verify() {
                                Err("Authentication fail, permission denied")
                                w.WriteHeader(403)
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

                    if IsFunc(route_func) {
                        res, status := route_func.(func(Request)(Response,int))(body)

                        if status == 0 {status = 200}

                        w.WriteHeader(status)
                        json.NewEncoder(w).Encode(
                            res,
                        )
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

func (router *Pistol) SignDB(con_str string, createDB func (string) (*gorm.DB, error), models ...interface{}) (*gorm.DB, error) {
    db, err := createDB(con_str)
    router.Database = db

    if err != nil {
        Die("Error on creation of tables on database")
    }

    if models != nil {
        db.Migrator().CurrentDatabase()
        db.AutoMigrate(models...)
    }

    _database = db
    return router.Database, err
}

func (router *Pistol) Run(path string, port uint) {
    router.Mux = http.NewServeMux()

    router.RootPath = path
    router.Mux.HandleFunc(path, router.RootRoute)

    handler := cors.Default().Handler(router.Mux)
    Err(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}

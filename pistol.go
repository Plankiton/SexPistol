package Sex

import (
    "github.com/Showmax/go-fqdn"
    "github.com/rs/cors"
    "net/http"
    "strings"
)

// Pistol is the Sex HTTP handler, who are used to setup http server
// Example:
//    router := Sex.NewPistol().
//       Add("/", func(Sex.Request) string {
//          return "Hello World"
//       }).
//       Run()
type Pistol struct {
    *http.ServeMux
    RootPath        string
    RouteConfs      Dict
    Routes          Dict
    RawRoutes       []string
    Auth            bool
}

// Function thats create a Sex.Pistol and create the init configurations
// Example:
//    router := Sex.NewPistol()
func NewPistol() *Pistol {
    pistol := new(Pistol)
    pistol.ServeMux = http.NewServeMux()
    pistol.AddRaw("/", pistol.root)

    return pistol
}

// Function to Add endpoints to the Sex.Pistol Server
// path are the endpoint location
// route is a void interface thats need to be on next format list:
//     - func (http.ResponseWriter, *http.Request)
//     - func (Sex.Request) Sex.Response
//                                         (res, status)
//     - func (Sex.Request) string   // Or (string,   int)
//     - func (Sex.Request) []byte   // Or ([]byte,   int)
//     - func (Sex.Request) Sex.Json // Or (Sex.Json, int)
// methods are a list of accepted HTTP methods to endpoint
// Example:
//       router.Add("/", func(Sex.Request) string {
//          return "Hello World"
//       }, "POST")
//       router.Add("/ok", func(Sex.Request) Sex.Json, int {
//          return map[stirng]bool{
//             "ok": true,
//          }, 404
//       })
func (pistol *Pistol) Add(path string, route interface {}, methods ...string) *Pistol {
    if f, ok := route.(httpRawFunc); ok {
        return pistol.AddRaw(path, f, methods...)
    }

    path = fixPath(path)
    root_path := fixPath(pistol.RootPath)
    if (path != root_path) {
        path = root_path + path
    }

    path_pattern := GetPathPattern(path)

    if len(pistol.RouteConfs) == 0 {
        pistol.RouteConfs = make(Dict)
    }
    if len(pistol.Routes) == 0 {
        pistol.Routes = make(Dict)
    }

    if _, exist := pistol.RouteConfs[path_pattern]; !exist {
        pistol.RouteConfs[path_pattern] = make(Prop)
    }
    conf := Prop{}
    conf["path-template"] = path
    pistol.RouteConfs[path_pattern] = conf

    if methods != nil {
        if _, exist := pistol.Routes[path_pattern]; !exist {
            pistol.Routes[path_pattern] = make(Prop)
        }
        for _, method := range methods {
            method = strings.ToUpper(method)
            pistol.Routes[path_pattern].(Prop)[method] = route
        }
    } else {
        pistol.Routes[path_pattern] = route
    }

    return pistol
}

// Function to Add golang raw http endpoints to the Sex.Pistol Server
// Example:
//       router.AddRaw("/", func(w http.ResponseWriter, r *http.Request) {
//          w.Write([]byte("Hello World"))
//       })
func (pistol *Pistol) AddRaw(path string, f func(http.ResponseWriter, *http.Request), methods...string) (*Pistol) {
    if pistol.ServeMux == nil {
        pistol.ServeMux = http.NewServeMux()
    }

    path = fixPath(path)
    pistol.HandleFunc(path, func(w http.ResponseWriter, r *http.Request){
        run_request := true

        if methods != nil {
            run_request = false
            for _, m := range methods {
                if strings.ToUpper(m) != r.Method {
                    run_request = true
                    break
                }
            }
        }

        if run_request {
            Log(r.Method, path, r.URL.RawQuery)
            f(w, r)
        }
    })

    pistol.RawRoutes = append(pistol.RawRoutes, path)
    return pistol
}

// Function to execute de Sex.Pistol server
// Example:
//    pistol.Run(5000)        // Will run server on port 5000
//    pistol.Run("/joao")     // will run server on path "/joao"
//    pistol.Run("/joao", 80) // will run server on path "/joao" and port 80
//    pistol.Run(80, "/joao") // will run server on path "/joao" and port 80
//
// If you run a Sex Pistol server with $SEX_DEBUG setted as "true" thats function will to log list all Sex endpoints of router
func (pistol *Pistol) Run(a ...interface{}) error {
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

    pistol.RootPath = path
    host, err := fqdn.FqdnHostname()
    if err != nil {
        host = "localhost"
    }

    Log(Fmt("Running Sex Pistol server at %s:%d%s", host, port, path))
    if GetEnv("SEX_DEBUG", "false") != "false" {
        for path, methods := range pistol.Routes {
            methods_str := ""
            if methods, ok := methods.(Prop); ok {
                for method := range methods {
                    methods_str += Fmt("%s ", method)
                }
            } else {
                methods_str = "ALL METHODS"
            }

            Log(Fmt("%s <- %s", pistol.RouteConfs[path].(Prop)["path-template"], methods_str))
        }

        for _, path := range pistol.RawRoutes {
            Log(Fmt("%s <- %s", path, "ALL METHODS"))
        }
    }

    handler := cors.Default().Handler(pistol)
    return http.ListenAndServe(Fmt(":%d", port), handler)
}

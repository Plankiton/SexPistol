package Sex

import (
	"net/http"
	"strings"

	"github.com/Showmax/go-fqdn"
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
	rootPath string

	config Config
	routes Dict

	plugins []Plugin

	err error
}

type Plugin interface {
	Name() string
	Init(*Pistol) (*Pistol, error)
	Root(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request)
}

// Get Pistol running path
func (pistol *Pistol) GetPath() string { return pistol.rootPath }

// Get Pistol route list
func (pistol *Pistol) GetRoutes() Dict { return pistol.routes }

// Get Pistol last error
func (pistol *Pistol) Error() error { return pistol.err }

// Set Pistol last error
func (pistol *Pistol) SetErr(err error) { pistol.err = err }

// Function thats create a Sex.Pistol and create the init configurations
// Example:
//    router := Sex.NewPistol()
func NewPistol() *Pistol {
	pistol := new(Pistol)
	for _, plugin := range pistol.plugins {
		p, err := plugin.Init(pistol)
		if err == nil {
			pistol = p
		}
	}

	pistol.config = make(Config)
	pistol.routes = make(Dict)
	pistol.ServeMux = http.NewServeMux()
	pistol.HandleFunc("/", pistol.root)

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
func (pistol *Pistol) Add(path string, route interface{}, methods ...string) *Pistol {
	for i, method := range methods {
		methods[i] = strings.ToUpper(method)
	}

	path = fixPath(path)
	root_path := fixPath(pistol.rootPath)
	if path != root_path {
		path = fixPath(root_path + path)
	}

	path_pattern := GetPathPattern(path)

	conf := Prop{
		"path-template":   []string{path},
		"methods-allowed": methods,
	}

	pistol.config[path_pattern] = conf
	pistol.routes[path_pattern] = route

	return pistol
}

// Add plugin to Sex Pistol
func (pistol *Pistol) AddPlugin(plugin Plugin) *Pistol {
	pistol.plugins = append(pistol.plugins, plugin)
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
			if v, ok := v.(string); ok {
				path = v
			}

			if v, ok := v.(int); ok {
				port = v
			}
		}
	}

	pistol.rootPath = path
	host, err := fqdn.FqdnHostname()
	if err != nil {
		host = "localhost"
	}

	msg := Fmt("Running Sex Pistol server at %s:%d%s", host, port, path)
	RawLog(LogLevelInfo, false, msg)
	if GetEnv("SEX_DEBUG", "false") != "false" {
		for path := range pistol.routes {
			methods_str := "ALL METHODS"
			if methods := pistol.config[path].Values("methods-allowed"); len(methods) > 0 {
				methods_str = strings.Join(methods, ", ")
			}

			msg := Fmt("%s <- %s", pistol.config[path].Get("path-template"), methods_str)
			RawLog(LogLevelInfo, false, msg)
		}
	}

	return http.ListenAndServe(Fmt(":%d", port), pistol)
}

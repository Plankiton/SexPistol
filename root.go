package Sex

import (
	"os"
	re "regexp"

	"net/http"
)

// Function to make Sex.Pistol a http.Handler
func (pistol *Pistol) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if pistol == nil || pistol.ServeMux == nil {
		pistol = NewPistol()
	}

	for _, plugin := range pistol.plugins {
		w, r = plugin.Root(w, r)
	}

	pistol.ServeMux.ServeHTTP(w, r)
}

// root endpoint to run all Sex endpoints
func (pistol *Pistol) root(w http.ResponseWriter, r *http.Request) {
	request := NewRequest()
	request.Request = *r

	path := r.URL.Path
	path = fixPath(r.URL.Path)

	root_path := fixPath(pistol.rootPath)
	if root_path == "/" {
		root_path = ""
	}

	response_log_message := Fmt("%s %s %s ", r.Method, path, r.URL.RawQuery)
	for path_pattern, route := range pistol.routes {
		path_regex := re.MustCompile("^" + root_path + path_pattern + `{1}`)
		conf := pistol.config[path_pattern]

		isMatching := path_regex.MatchString(path)
		if os.Getenv("SEX_DEBUG") == "true" {
			RawLog(LogLevelInfo, false, Fmt(`"%s" is "%s" ? %v`, path_pattern, path, isMatching))
		}

		if isMatching {
			request.Conf = conf
			request.PathVars, _ = GetPathVars(conf.Get("path-template"), path)
			methods_allowed := conf.Values("methods-allowed")

			if len(methods_allowed) > 0 && IndexOf(request.Method, methods_allowed) < 0 {
				st := StatusMethodNotAllowed
				msg := StatusText(st)
				http.Error(w, msg, st)
				RawLog(LogLevelError, false,
					response_log_message+Fmt("%d: %s", st, msg))
				return
			}

			response := NewResponse()
			response.ResponseWriter = w
			if err := runRoute(route, *response, *request); err != nil {
				st := response.Status
				msg := StatusText(st)
				if len(response.Body) == 0 {
					msg = string(response.Body)
				}
				RawLog(LogLevelError, false, response_log_message+Fmt("%d: %s", st, msg))
				http.Error(w, msg, st)
			}

			return
		}
	}

	st := StatusNotFound
	msg := StatusText(st)
	http.Error(w, msg, st)
	RawLog(LogLevelError, false,
		response_log_message+Fmt("%d: %s", st, msg))
}

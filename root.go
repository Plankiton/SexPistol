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

    path := r.URL.Path
    path = fixPath(r.URL.Path)

    root_path := fixPath(pistol.rootPath)
    if root_path == "/" {
        root_path = ""
    }

    response_log_message := Fmt("%s %s %s", r.Method, path, r.URL.RawQuery)
    for path_pattern, methods := range pistol.routes {

        path_regex := re.MustCompile("^"+root_path+path_pattern+`{1}`)

        route_conf := pistol.routeConfs[path_pattern]

        iroute_func := methods
        if methods, ok := methods.(Prop); ok {
            iroute_func = methods[r.Method]
        }
        isMatching := path_regex.MatchString(path)
        if os.Getenv("SEX_DEBUG") == "true" {
            RawLog(LogLevelInfo, false, Fmt(`"%s" is "%s" ? %v`,path_pattern, path, isMatching))
        }
        if isMatching {

            if methods != nil {

                if iroute_func != nil {
                    request.Conf = route_conf.(Prop)
                    request.PathVars, _ = GetPathVars(route_conf.(Prop)["path-template"].(string), path)
                    request.Request = r

                    response := new(Response)
                    response.ResponseWriter = w

                    if err := runRouteFunc(iroute_func, response, request); err != nil {
                        http.Error(response, err.Error(), response.Status)
                    }

                    response_log_message += Fmt("%d: %s", status_code, http.StatusText(response.Status))
                    if response.Status >= 400 {
                        RawLog(LogLevelError, false, response_log_message)
                    } else
                    if response.Status >= 300 {
                        RawLog(LogLevelWarn, false, response_log_message)
                    } else {
                        RawLog(LogLevelInfo, false, response_log_message)
                    }

                    return

                }
            }

            status_code := 405
            response_log_message += Fmt("%d: %s", status_code, http.StatusText(status_code))
            RawLog(LogLevelError, false, response_log_message)
            w.WriteHeader(status_code)
            return
        }
    }

    status_code := 404
    response_log_message += Fmt("%d: %s", status_code, http.StatusText(status_code))
    RawLog(LogLevelError, false, response_log_message)
}

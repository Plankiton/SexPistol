package Sex

import (
	re "regexp"
	str "strings"

	"bytes"
	"errors"
	"net/http"
)

// Request is the request sent by client (*http.Request) with inproviments like path variables and Pistol Route configurations
// Example:
//    router.Add("/hello/{name}", func (r Sex.Request) string {
//        name := r.PathVars["name"]
//        return "Hello "+ name
//    }
type Request struct {
	http.Request
	PathVars Prop
	Conf     Prop
}

// NewRequest create new Request
func NewRequest() *Request {
	return new(Request)
}

// JSON provides marshalled Json body to a variable
// Example:
//      var data map[string]interface{} // Can be Structs too
//      r.JsonBody(&data)
func (r *Request) Json(v interface{}) error {
	encoded := new(bytes.Buffer)
	encoded.ReadFrom(r.Body)
	return Jsonify(encoded.Bytes(), v)
}

// Raw provides byte array body to a variable
// Example:
//      var data []byte
//      r.Raw(&data)
func (r *Request) Raw(b *[]byte) error {
	body := new(bytes.Buffer)
	_, err := body.ReadFrom(r.Body)
	*b = body.Bytes()

	return err
}

// GetPathPattern provides regex pattern of a Sex path template
// Example:
//    Sex.GetPathPattern("/hello/{name}")
func GetPathPattern(t string) string {
	var_patt := re.MustCompile(`\{(\w{1,}):{0,1}(.{0,})\}`)

	path_tmplt := str.Split(t, "/")

	path_pattern := "/"
	for i := 0; i < len(path_tmplt); i++ {
		if path_tmplt[i] == "" {
			continue
		}

		values := var_patt.FindStringSubmatch(path_tmplt[i])
		if len(values) >= 2 {
			if values[2] == "" {
				path_pattern += `[a-zA-Z0-9_\ ]{1,}`
			} else {
				path_pattern += values[2]
			}
		} else {
			path_pattern += path_tmplt[i]
		}

		path_pattern += "/"
	}
	path_pattern = fixPath(path_pattern)
	path_pattern += "$"
	if path_pattern == "^$" {
		path_pattern = "^/$"
	}

	return path_pattern
}

// GetPathVars provides path variables a Sex path template
// Example:
//    Sex.GetPathVars("/hello/{name}", "/hello/joao")
func GetPathVars(t string, p string) (Prop, error) {
	var_patt := re.MustCompile(`\{(\w{1,}):{0,1}(.{0,})\}`)

	path_tmplt := str.Split(fixPath(t), "/")
	path := str.Split(fixPath(p), "/")

	if len(path) != len(path_tmplt) {
		return Prop{}, errors.New("Path don't match with the path template")
	}

	path_vars_values := []map[string]string{}
	for i := 0; i < len(path); i++ {

		values := var_patt.FindStringSubmatch(path_tmplt[i])
		tmpl_patt := re.MustCompile("")

		if len(values) == 3 {
			tmpl_patt = re.MustCompile(values[2])

			if tmpl_patt.MatchString(path[i]) {
				path_vars_values = append(path_vars_values, map[string]string{
					"name":  values[1],
					"value": path[i],
				})
			} else {
				return Prop{}, errors.New(
					Fmt("Variable \"%s\" need that \"%s\"  match to \"%s\"", values[1], path[i], values[2]),
				)
			}

			continue
		}

		path_vars_values = append(path_vars_values, map[string]string{
			"value": path[i],
		})

	}

	path_vars := Prop{}
	for _, v := range path_vars_values {
		if _, exist := v["name"]; exist {
			path_vars.Add(v["name"], v["value"])
		}
	}

	return path_vars, nil
}

package Sex
import (
    str "strings"
    re "regexp"

    "net/http"
    "errors"
    "bytes"
    "time"
)

type Response struct {
    http.ResponseWriter
    Body   []byte
    Status int
}

type Request  struct {
    *http.Request
    PathVars    map[string]string
    Conf        Prop
    Writer      *Response
}

func (self *Request) JsonBody(v interface{}) error {
    encoded := new(bytes.Buffer)
    encoded.ReadFrom(self.Body)
    return FromJson(encoded.Bytes(), v)
}

func (self *Request) RawBody(b *[]byte) error {
    body := new(bytes.Buffer)
    _, err := body.ReadFrom(self.Body)
    *b = body.Bytes()

    return err
}

func (self *Request) MkResponse() *Response {
    return self.Writer
}

func (self *Response) SetBody(v []byte) *Response {
    self.Body = v
    return self
}

func (self *Response) SetStatus(code int) *Response {
    self.Status = 200
    return self
}

func (self *Response) SetCookie(key string, value string, expires time.Duration) *Response {
    cookie := &http.Cookie {
        Name: key,
        Value: value,
        Expires: time.Now().Add(expires),
    }
    http.SetCookie(self, cookie)

    return self
}

func GetPathPattern(t string) string {
    var_patt := re.MustCompile(`\{(\w{1,}):{0,1}(.{0,})\}`)

    path_tmplt := str.Split(t, "/")

    path_pattern := "/"
    for i := 0; i < len(path_tmplt); i++ {
        if path_tmplt[i] == "" {
            continue
        }

        path_tmplt[i] += "/"
        values := var_patt.FindStringSubmatch(path_tmplt[i])
        if len(values)>=2 {
            if values[2] == "" {
                path_pattern += `[a-zA-Z0-9_\ ]{1,}`
            } else {
                path_pattern += values[2]
            }
        } else {
            path_pattern += path_tmplt[i]
        }
    }
    path_pattern = fixPath(path_pattern)
    path_pattern += "$"
    if path_pattern == "^$" {
        path_pattern = "^/$"
    }

    return path_pattern
}

func GetPathVars(t string, p string) (map[string]string, error) {
    var_patt := re.MustCompile(`\{(\w{1,}):{0,1}(.{0,})\}`)

    path_tmplt := str.Split(t, "/")
    path := str.Split(p, "/")

    if len(path) != len(path_tmplt) {
        return map[string]string {}, errors.New("Path don't match with the path template")
    }

    path_vars_values := []map[string]string{}
    for i := 0; i < len(path); i++ {

        values := var_patt.FindStringSubmatch(path_tmplt[i])

        tmpl_patt := re.MustCompile("")

        if len(values)==3 {
            tmpl_patt = re.MustCompile(values[2])

            if tmpl_patt.MatchString(path[i]){
                path_vars_values = append(path_vars_values, map[string]string {
                    "name": values[1],
                    "value": path[i],
                })
            } else {
                return map[string]string {}, errors.New(
                    Fmt("Variable \"%s\" need that \"%s\"  match to \"%s\"", values[1], path[i], values[2]),
                )
            }

            continue
        }

        path_vars_values = append(path_vars_values, map[string]string{
            "value": path[i],
        })

    }

    path_vars := map[string]string{}
    for _, v := range path_vars_values {
        if _, exist := v["name"]; exist {
            path_vars[v["name"]] = v["value"]
        }
    }

    return path_vars, nil
}

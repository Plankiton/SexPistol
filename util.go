package sex

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strings"

	"mime/multipart"

	"golang.org/x/crypto/bcrypt"
)

type List [] interface{}
type Dict map[interface{}] interface{}

func (self Dict) ToStrMap() map[string]interface{} {
    m := map[string]interface{}{}
    for v, k := range self {
        m[v.(string)] = k
    }
    return m
}

func MapTo(m map[string]interface{}, v interface{}) {
    encoded, _ := json.Marshal(m)
    json.Unmarshal(encoded, v)
}

func ToHash(s string) string {
    h := sha1.New()
    io.WriteString(h, s)
    return fmt.Sprintf("%x", h.Sum(nil))
}
func ToPassHash(s string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
    return string(hash), err
}

func CheckPass(p []byte, s string) (error) {
    err := bcrypt.CompareHashAndPassword(p, []byte(s))
    return err
}

func GetEnv(key string, def string) string {
    val, ok := os.LookupEnv(key)
    if !ok {
        return def
    }
    return val
}

func ToLabel(ID interface{}, Type string) string {
    if (reflect.TypeOf(ID).Kind() == reflect.Int) {
        return fmt.Sprintf("<%s:%d>", Type, ID)
    }
    return fmt.Sprintf("<%s:%v>", Type, ID)
}

func GetPrototype(model interface{}) string {
    return fmt.Sprintf("%+v", model)
}

func GetModelType(model interface{}) string {
    t := reflect.TypeOf(model)
    type_raw_text := t.String()
    type_raw_list := ReCompile(`[a-zA-Z0-9]{1,}\.`).Split(type_raw_text, -1)

    return strings.Join(type_raw_list, "")
}

func TypeParse(t string) string {
    t_re := ReCompile(`([A-Z][a-z_0-9]{1,})?`)
    list_match := t_re.FindAllStringSubmatch(t, -1)

    type_out := ""
    for i, v := range list_match {
        type_out += strings.ToLower(v[0])

        if i < len(list_match)-1 {
            type_out += "_"
        }
    }
    type_out += "s"

    return type_out
}

func ReCompile(pattern string) *regexp.Regexp {
    return regexp.MustCompile(pattern)
}

func SuperPut (v...interface{}) {
    fmt.Println("\n--------------------------------")
    fmt.Println(v...)
    fmt.Print("--------------------------------\n")
}

func GetPathPattern(t string) string {
    var_patt := ReCompile(`\{(\w{1,}):{0,1}(.{0,})\}`)

    path_tmplt := strings.Split(t, "/")

    path_pattern := "^/"
    for i := 0; i < len(path_tmplt); i++ {
        if path_tmplt[i] == "" {
            continue
        }

        values := var_patt.FindStringSubmatch(path_tmplt[i])
        if len(values)>=2 {
            if values[2] == "" {
                path_pattern += "[a-zA-Z0-9_]{1,}"
            } else {
                path_pattern += values[2]
            }
        } else {
            path_pattern += path_tmplt[i]
        }

        path_pattern += "/"
    }
    path_pattern += "$"

    return path_pattern
}

func GetPathVars(t string, p string) (map[string]string, error) {
    var_patt := ReCompile(`\{(\w{1,}):{0,1}(.{0,})\}`)

    path_tmplt := strings.Split(t, "/")
    path := strings.Split(p, "/")

    if len(path) != len(path_tmplt) {
        return map[string]string {}, errors.New("Path don't match with the path template")
    }

    path_vars_values := []map[string]string{}
    for i := 0; i < len(path); i++ {

        values := var_patt.FindStringSubmatch(path_tmplt[i])

        tmpl_patt := ReCompile("")

        if len(values)==3 {
            tmpl_patt = ReCompile(values[2])

            if tmpl_patt.MatchString(path[i]){
                path_vars_values = append(path_vars_values, map[string]string {
                    "name": values[1],
                    "value": path[i],
                })
            } else {
                return map[string]string {}, errors.New(
                    fmt.Sprintf("Variable \"%s\" need that \"%s\"  match to \"%s\"", values[1], path[i], values[2]),
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

func IsRawFunc(f interface{}) bool {
    var gf RawRouteFunc
    return reflect.TypeOf(f).AssignableTo(reflect.TypeOf(gf))
}

func IsFunc(f interface{}) bool {
    var gf RouteFunc
    return reflect.TypeOf(f).AssignableTo(reflect.TypeOf(gf))
}

func GenericInterface () reflect.Type {
    var i interface{}
    return reflect.TypeOf(&i).Elem()
}

func GenericString () reflect.Type {
    var i string
    return reflect.TypeOf(i)
}

func GenericJsonObj() reflect.Type {
    return reflect.MapOf(GenericString(), GenericInterface())
}

func GenericJsonArray() reflect.Type {
    return reflect.ArrayOf(-1, reflect.MapOf(GenericString(), GenericInterface().Elem()))
}

func GenericBuff() reflect.Type {
    var i bytes.Buffer
    return reflect.TypeOf(&i)
}

func GenericForm() reflect.Type {
    var i multipart.Form
    return reflect.TypeOf(&i)
}

func ValidateData(data interface{}, t func()reflect.Type) bool {
    if data == nil {
        return false
    }

    if reflect.TypeOf(data).Kind() == t().Kind() {
        return true
    }

    return false
}

package Sex

import (
    str "strings"
    re "regexp"

    "mime/multipart"
    "encoding/json"
    "net/url"

    "reflect"
    "bytes"
    "fmt"
    "os"
)

func fixPath(path string) string {
    end_path := len(path)-1
    if path[end_path] == '/' {
        path = path[:end_path]
    }

    return path
}

func Copy(m interface{}, v interface{}) error {
    encoded, err := json.Marshal(m)
    if err != nil {
        return err
    }

    return json.Unmarshal(encoded, v)
}

func FromJson(encoded []byte, v interface{}) error {
    return json.Unmarshal(encoded, v)
}

func Jsonify(v interface{}) []byte {
    res, _ := json.Marshal(v)
    return res
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
        return Fmt("<%s:%d>", Type, ID)
    }

    return Fmt("<%s:%v>", str.Replace(Type, "*", "", -1), ID)
}

func TypeParse(t string) string {
    t_re := re.MustCompile(`([A-Z][a-z_0-9]{1,})?`)
    list_match := t_re.FindAllStringSubmatch(t, -1)

    type_out := ""
    for i, v := range list_match {
        type_out += str.ToLower(v[0])

        if i < len(list_match)-1 {
            type_out += "_"
        }
    }
    type_out += "s"

    return type_out
}

func SuperPut (v...interface{}) {
    fmt.Println("\n--------------------------------")
    fmt.Println(v...)
    fmt.Print("--------------------------------\n")
}


func isRawFunc(f interface{}) bool {
    var gf RawRouteFunc
    return reflect.TypeOf(f).AssignableTo(reflect.TypeOf(gf))
}

func isStrFunc(f interface{}) bool {
    var gf StrRouteFunc
    return reflect.TypeOf(f).AssignableTo(reflect.TypeOf(gf))
}

func isResFunc(f interface{}) bool {
    var gf ResRouteFunc
    return reflect.TypeOf(f).AssignableTo(reflect.TypeOf(gf))
}

func isPureResFunc(f interface{}) bool {
    var gf PureResRouteFunc
    return reflect.TypeOf(f).AssignableTo(reflect.TypeOf(gf))
}

func isFunc(f interface{}) bool {
    var gf RouteFunc
    return reflect.TypeOf(f).AssignableTo(reflect.TypeOf(gf))
}

func GenericInterface () reflect.Type {
    var i interface{}
    return reflect.TypeOf(i)
}

func GenericInt () reflect.Type {
    var i int
    return reflect.TypeOf(i)
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
    var i url.Values
    return reflect.TypeOf(&i)
}

func GenericMultipartForm() reflect.Type {
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

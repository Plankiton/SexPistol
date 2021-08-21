package Sex

import (
    str "strings"

    "mime/multipart"
    "encoding/json"
    "net/url"

    "reflect"
    "bytes"
    "os"
)

func fixPath(path string) string {
    return "/"+str.Trim(str.Trim(path, " "), "/")
}

// Sex utility function to make copy of map or struct to another map or struct
// Required: Destine need to be a pointer
// Example:
//    var m struct { Name string `json:"name"` }
//    j := map[string]interface{}{
//       "name": "Joao",
//    }
//    Sex.Copy(j, &m)
func Copy(source interface{}, destine interface{}) error {
    encoded, err := json.Marshal(source)
    if err != nil {
        return err
    }

    return json.Unmarshal(encoded, destine)
}

// Sex utility function to make merge of map or struct and another map or struct
// Required: Destine need to be a pointer
// Example:
//    var m := struct { Name string `json:"name"` } {
//        Name: "Joao",
//    }
//    j := map[string]interface{}{
//       "idade": "Joao",
//       "name": nil,
//    }
//    Sex.Copy(m, &j)
//
// Merge rules:
//    If the field on source dont exists on destine it will be created (just if destine are map)
//    If the field on source exists on destine but are dont seted it will be seted
//    If the field on source exists on destine but are seted it will not be seted
//    If override are seted as true, the field on destine will be overrided by source
func Merge(source interface{}, destine interface{}, override ...bool) (map[string]interface{}, error) {
    final := map[string]interface{}{}

    dst := map[string]interface{}{}
    src := map[string]interface{}{}

    ok := false
    if src, ok = source.(map[string]interface{}); !ok {
        err := Copy(source, &src)
        if err != nil {
            return final, err
        }
    }

    if dst, ok = destine.(map[string]interface{}); !ok {
        err := Copy(destine, &dst)
        if err != nil {
            return final, err
        }
    }

    for k, v := range src {
        final[k] = v
    }

    for k, v := range dst {
        if _, exists := final[k]; exists && (override != nil && override[0]) {
            continue
        }

        final[k] = v
    }

    return final, nil
}

// Function thats parse a byte list on a json and write on a variable
// Required: v needs to be a pointer
func FromJson(encoded []byte, v interface{}) error {
    return json.Unmarshal(encoded, v)
}

// Function thats parse a byte list on a json and write on a variable
func Jsonify(v interface{}) []byte {
    res, _ := json.Marshal(v)
    return res
}

// Function thats get a environment var or default value if var does not exist
func GetEnv(key string, def string) string {
    val, ok := os.LookupEnv(key)
    if !ok {
        return def
    }
    return val
}

func isRawFunc(f interface{}) bool {
    var gf rawRouteFunc
    return reflect.TypeOf(f).AssignableTo(reflect.TypeOf(gf))
}

func isRawFuncNoStatus(f interface{}) bool {
    var gf rawRouteFuncNoStatus
    return reflect.TypeOf(f).AssignableTo(reflect.TypeOf(gf))
}

func isStrFunc(f interface{}) bool {
    var gf strRouteFunc
    return reflect.TypeOf(f).AssignableTo(reflect.TypeOf(gf))
}

func isStrFuncNoStatus(f interface{}) bool {
    var gf strRouteFuncNoStatus
    return reflect.TypeOf(f).AssignableTo(reflect.TypeOf(gf))
}

func isResFunc(f interface{}) bool {
    var gf resRouteFunc
    return reflect.TypeOf(f).AssignableTo(reflect.TypeOf(gf))
}

func isResFuncNoStatus(f interface{}) bool {
    var gf resRouteFuncNoStatus
    return reflect.TypeOf(f).AssignableTo(reflect.TypeOf(gf))
}

func isJsonFunc(f interface{}) bool {
    var gf jsonRouteFunc
    return reflect.TypeOf(f).AssignableTo(reflect.TypeOf(gf))
}

func isJsonFuncNoStatus(f interface{}) bool {
    var gf jsonRouteFuncNoStatus
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

package Sex

import (
    str "strings"

    "encoding/json"
    "os"
)

func fixPath(path string) string {
    return "/"+str.Trim(str.Trim(path, " "), "/")
}

// Sex utility function to make copy of map or struct to another map or struct
// Required: Destine need to be a pointer
// Example:
//    var m struct { Name string `json:"name"` }
//    j := Dict{
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
//    j := Dict{
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
func Merge(source interface{}, destine interface{}, override ...bool) (Dict, error) {
    final := Dict{}

    dst := Dict{}
    src := Dict{}

    ok := false
    if src, ok = source.(Dict); !ok {
        err := Copy(source, &src)
        if err != nil {
            return final, err
        }
    }

    if dst, ok = destine.(Dict); !ok {
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

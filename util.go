package Sex

import (
	"reflect"
	str "strings"

	"encoding/json"
	"os"
)

func fixPath(path string) string {
	return "/" + str.Trim(str.Trim(path, " "), "/")
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
func Merge(source interface{}, destine interface{}, override ...bool) (interface{}, error) {
	final := Dict{}

	if override != nil && override[0] {
		if err := Copy(destine, &final); err != nil {
			return nil, err
		}
		if err := Copy(source, &final); err != nil {
			return nil, err
		}
	} else {
		if err := Copy(source, &final); err != nil {
			return nil, err
		}
	}

	if err := Copy(final, destine); err != nil {
		return nil, err
	}

	return destine, nil
}

// FromJSON function thats parse a byte list on a json and write on a variable
// Required: v needs to be a pointer
func FromJSON(encoded []byte, v interface{}) error {
	return json.Unmarshal(encoded, v)
}

// Jsonify function thats parse a byte list on a json and write on a variable
func Jsonify(v interface{}) []byte {
	res, _ := json.Marshal(v)
	return res
}

// GetEnv function thats get a environment var or default value if var does not exist
func GetEnv(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}

// IndexOf function to index first ocurrence of string on string slice
func IndexOf(i interface{}, l interface{}) int {
	if typ := reflect.TypeOf(l).Kind(); typ == reflect.Slice || typ == reflect.Array || typ == reflect.String {
		list := reflect.ValueOf(l)
		item := reflect.ValueOf(i)

		for i := 0; i < list.Len(); i++ {
			value := list.Index(i)
			if reflect.DeepEqual(value.Interface(), item.Interface()) {
				return i
			}
		}
	}

	return -1
}

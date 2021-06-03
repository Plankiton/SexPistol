package SexHtml
import (
    "text/template"
    "errors"
    "bytes"
)

func Render(input interface{}, data interface{}) ([]byte, error) {
    var err error

    tmpl := template.New("SexTemplate")
    buffer := new(bytes.Buffer)

    if input, ok := input.(string); ok {
        tmpl, err = tmpl.Parse(input)
        if err == nil {
            err = tmpl.Execute(buffer, data)
            return buffer.Bytes(), err
        }
    }

    if err == nil {
        err = errors.New("Invalid input type")
    }

    return []byte{}, err
}

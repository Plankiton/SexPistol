package SexHtml
import (
    "text/template"
    "errors"
    "bytes"
    "os"
)

// Function to simplify text/html templating
func Render(input interface{}, data interface{}) ([]byte, error) {
    var err error

    tmpl := template.New("SexTemplate")
    buffer := new(bytes.Buffer)

    if text, ok := input.(string); ok {
        tmpl, err = tmpl.Parse(text)
        if err == nil {
            err = tmpl.Execute(buffer, data)
            return buffer.Bytes(), err
        }
    } else
    if file, ok := input.(*os.File); ok {
        size, err := file.Seek(0, os.SEEK_END)

        file.Seek(0, os.SEEK_SET)
        if err == nil {
            content := make([]byte, size)
            _, err = file.Read(content)

            if err == nil {
                tmpl, err = tmpl.Parse(string(content))
                if err == nil {
                    err = tmpl.Execute(buffer, data)
                    return buffer.Bytes(), err
                }
            }
        }
    }

    if err == nil {
        err = errors.New("Invalid input type")
    }

    return []byte{}, err
}

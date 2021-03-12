package sex

import (
	b64 "encoding/base64"
	mp "mime/multipart"
	"time"

	"bytes"
	"fmt"
	"os"
)

type File struct {
    Model
    AltText   string  `json:"alt_text,omitempty"`
    Path      string  `json:"-"`
    Filename  string  `json:"-" gorm:"index"`
    Mime      string  `json:"-"`
}

func (model *File) Render() string {
    if model.Path == "" {
        root := GetEnv("DB_FILE_ROOT", ".")
        model.Path = root + "/uploads/"
    }
    if model.Path[len(model.Path)-1] != '/' {
        model.Path += "/"
    }

    file, _ := os.Open(model.Path + model.Filename)
    f := new(bytes.Buffer)
    f.ReadFrom(file)

    buff := make([]byte, b64.RawStdEncoding.EncodedLen(f.Len()))
    b64.RawStdEncoding.Encode(buff, f.Bytes())
    return "data:" + model.Mime + ";base64," + string(buff)
}

func (model *File) Create() bool {
    if (model.ModelType == "") {
        model.ModelType = GetModelType(model)
    }

    if ModelCreate(model) == nil {
        ID := model.ID
        ModelType := model.ModelType
        Log("Created", ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *File) Delete() bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelCreate(model) == nil {
        Log("Deleted", ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *File) Save() bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelSave(model) == nil {
        Log("Updated", ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *File) Update(columns Dict) bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelUpdate(model, columns) == nil {
        Log("Updated", ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model * File) Load(form *mp.Form) {
    model.AltText = form.Value["description"][0]
    buff := bytes.NewBuffer(nil)
    for _, h := range form.File["data"] {
        file, _ := h.Open()
        model.Filename = h.Filename
        model.Mime = h.Header.Get("Content-Type")

        buff.ReadFrom(file)
        file.Close()
    }
    model.Create()
    _database.First(model)

    if model.Path == "" {
        root := GetEnv("DB_FILE_ROOT", ".")
        model.Path = root + "/uploads/"
    }
    if model.Path[len(model.Path)-1] != '/' {
        model.Path += "/"
    }

    if _, e := os.Stat(model.Path); os.IsNotExist(e) {
        os.MkdirAll(model.Path, 0700)
    }

    model.Filename = ToHash( fmt.Sprintf("%s;%s;%s;%s", model.AltText, model.Mime, model.Filename, time.Now().String()) )
    file, err := os.Open(model.Path + model.Filename)
    if err != nil {
        file, err = os.Create(model.Path + model.Filename)
    }

    file.Write(buff.Bytes())
    file.Close()

    model.Create()
    model.Save()
}

func (model * File) LoadString(s string) {
    model.Mime = "text/plain"
    model.Filename = "file"

    if model.Path == "" {
        root := GetEnv("DB_FILE_ROOT", ".")
        model.Path = root + "/uploads/"
    }
    if model.Path[len(model.Path)-1] != '/' {
        model.Path += "/"
    }

    if _, e := os.Stat(model.Path); os.IsNotExist(e) {
        os.MkdirAll(model.Path, 0700)
    }

    model.Filename = ToHash( fmt.Sprintf("%s;%s;%s;%s", model.AltText, model.Mime, model.Filename, time.Now().String()) )
    file, err := os.Open(model.Path + model.Filename)
    if err != nil {
        file, err = os.Create(model.Path + model.Filename)
    }

    file.Write([]byte(s))
    file.Close()

    model.Create()
}

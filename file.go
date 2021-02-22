package api

import (
	b64 "encoding/base64"
	mp "mime/multipart"

	"bytes"
  "fmt"
	"os"
)

type File struct {
    Model
    AltText   string  `json:"alt_text,omitempty"`
    Path      string  `json:"-"`
    Filename  string  `json:"-"`
    Mime      string  `json:"-"`
}

func (model *File) Render() string {
    file, _ := os.Open(model.Path)
    f := new(bytes.Buffer)
    f.ReadFrom(file)

    buff := make([]byte, b64.RawStdEncoding.EncodedLen(f.Len()))
    b64.RawStdEncoding.Encode(buff, f.Bytes())
    return "data:" + model.Mime + ";base64," + string(buff)
}

func (model *File) Create() {
    model.ModelType = GetModelType(model)

    _database.Create(model)

    e := _database.First(model)
    if e.Error == nil {
        ID := model.ID
        ModelType := model.ModelType
        Log("Created", ToLabel(ID, ModelType))
    }
}

func (model *File) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Delete(model)
        Log("Deleted", ToLabel(ID, ModelType))
    }
}

func (model *File) Save() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Save(model)
        Log("Updated", ToLabel(ID, ModelType))
    }
}

func (model *File) Update(columns Dict) {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.First(model).Updates(columns.ToStrMap())
        Log("Updated", ToLabel(ID, ModelType))
    }
}

func (model * File) Load(form *mp.Form) {
    model.AltText = form.Value["description"][0]
    buff := bytes.NewBuffer(nil)
    for _, h := range form.File["data"] {
        file, _ := h.Open()
        model.Mime = h.Header.Get("Content-Type")
        buff.ReadFrom(file)
        file.Close()
    }
    model.Create()
    _database.First(model)

    if model.Path == "" {
        root := GetEnv("$DB_FILE_ROOT", ".")
        model.Path = root + "/uploads/"
    }
    if model.Path[len(model.Path)-1] != '/' {
        model.Path += "/"
    }

    if _, e := os.Stat(model.Path); os.IsNotExist(e) {
        os.MkdirAll(model.Path, 0700)
    }

    model.Filename = ToHash( fmt.Sprintf("%d;%s;%s;%s", model.ID, model.AltText, model.Mime, model.Path) )
    file, err := os.Open(model.Path + model.Filename)
    if _, e := os.Stat(model.Path + model.Filename); os.IsNotExist(e) {
        file, err = os.Create(model.Path + model.Filename)
    }

    if err != nil {
        SuperPut(err)
    }

    SuperPut(file, buff.Bytes())
    file.Write(buff.Bytes())
    file.Close()

    model.Save()
}

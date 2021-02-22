package api

import (
    mp "mime/multipart"
    "bytes"
)

type File struct {
    Model
    Data      []byte  `json:"-"`
    AltText   string  `json:"alt_text,omitempty"`
    Mime      string  `json:"-"`
}

func (model *File) Render() string {
    return "data:" + model.Mime + "," + string(model.Data)
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

    model.Data = buff.Bytes()
}

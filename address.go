package api

import (
    "encoding/json"
    "net/http"
    "fmt"
)

type Address struct {
    Model
    Street       string `json:"street,omitempty"`
    State        string `json:"state,omitempty"`
    Number       string `json:"number,omitempty" gorm:"NOT NULL"`
    Code         string `json:"cep,omitempty"`
    City         string `json:"city,omitempty"`
    Neigh        string `json:"neighborhood,omitempty"`
    Compl        string `json:"complement,omitempty" gorm:"default:NULL"`
}

func (model *Address) FromPostalCode(cep string) *Address {
    model.Code = cep

    r_addr, err := http.Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", model.Code))
    if err == nil && r_addr.StatusCode == 200 {
        json.NewDecoder(r_addr.Body).Decode(&model)
    }

    return model
}

func (model *Address) Create() {
    model.ModelType = GetModelType(model)

    _database.Create(model)

    e := _database.First(model)
    if e.Error == nil {
        ID := model.ID
        ModelType := model.ModelType
        Log("Created", ToLabel(ID, ModelType))
    }
}

func (model *Address) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Delete(model)
        Log("Deleted", ToLabel(ID, ModelType))
    }
}

func (model *Address) Save() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Save(model)
        Log("Updated", ToLabel(ID, ModelType))
    }
}

func (model *Address) Update(columns Dict) {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.First(model).Updates(columns.ToStrMap())
        Log("Updated", ToLabel(ID, ModelType))
    }
}

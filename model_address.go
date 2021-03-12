package sex

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

func (model *Address) Create() bool {
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

func (model *Address) Delete() bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelCreate(model) == nil {
        Log("Deleted", ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *Address) Save() bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelSave(model) == nil {
        Log("Updated", ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *Address) Update(columns Dict) bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelUpdate(model, columns) == nil {
        Log("Updated", ToLabel(ID, ModelType))
        return true
    }

    return false
}

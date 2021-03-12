package sex

import "time"

type Event struct {
    Model
    Name       string       `json:"name,omitempty" gorm:"index"`
    Type       string       `json:"type,omitempty" gorm:"index"`

    AddrId     uint         `json:"-" gorm:"index"`
    CoverId    uint         `json:"-" gorm:"index"`

    BeginAt    time.Time    `json:"begin,omitempty" gorm:"index"`
    EndAt      time.Time    `json:"end,omitempty" gorm:"index"`
}

func (model *Event) Create() bool {
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

func (model *Event) Delete() bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelCreate(model) == nil {
        Log("Deleted", ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *Event) Save() bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelSave(model) == nil {
        Log("Updated", ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *Event) Update(columns Dict) bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelUpdate(model, columns) == nil {
        Log("Updated", ToLabel(ID, ModelType))
        return true
    }

    return false
}

package api

import "time"

type Event struct {
    Model
    Name       string       `json:"name,omitempty" gorm:"index"`
    Type       string       `json:"type,omitempty" gorm:"index"`

    AddrId     uint         `json:",empty" gorm:"index"`
    CoverId    uint         `json:",empty"`

    BeginAt    time.Time    `json:"begin,omitempty" gorm:"index"`
    EndAt      time.Time    `json:"end,omitempty" gorm:"index"`
}

func (model *Event) Create() {
    if (model.ModelType == "") {
        model.ModelType = GetModelType(model)
    }

    _database.Create(model)

    e := _database.First(model)
    if e.Error == nil {

        ID := model.ID
        ModelType := model.ModelType
        Log("Created", ToLabel(ID, ModelType))
    }
}

func (model *Event) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Delete(model)
        Log("Deleted", ToLabel(ID, ModelType))
    }
}

func (model *Event) Save() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(&Event{}, "id = ?", model.ID)
    if e.Error == nil {
        _database.Save(model)
        Log("Updated", ToLabel(ID, ModelType))
    }
}

func (model *Event) Update(columns Dict) {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(&Event{}, "id = ?", model.ID)
    if e.Error == nil {
        _database.First(model).Updates(columns.ToStrMap())
        Log("Updated", ToLabel(ID, ModelType))
    }
}

package api

import (
    "time"
    "gorm.io/gorm"
)

type Token struct {
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`

    ModelType string

    ID        string `json:"auth,omitempty" gorm:"PrimaryKey"`
    UserId    uint   `json:",empty"`
}

func (model *Token) Verify() bool {
    if _database.Where("id = ?", model.ID).First(model).Error == nil {
        return true
    }

    return false
}

func (model *Token) Create() {
    model.ModelType = GetModelType(model)

    _database.Create(model)

    e := _database.First(model)
    if e.Error == nil {


        ID := model.ID
        ModelType := model.ModelType
        Log("Created", ToLabel(ID, ModelType))
    }
}

func (model *Token) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Delete(model)
        Log("Deleted", ToLabel(ID, ModelType))
    }
}

func (model *Token) Save() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Save(model)
        Log("Updated", ToLabel(ID, ModelType))
    }
}

func (model *Token) Update(columns Dict) {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.First(model).Updates(columns.ToStrMap())
        Log("Updated", ToLabel(ID, ModelType))
    }
}
package SexDatabase

import (
    "time"
    "gorm.io/gorm"
)

type ModelSkel interface {
    TableName() string
    GetID() (interface{}, error)
    SetID(interface{}) error
    GetDB() (interface{}, error)
    SetDB(interface{}) error
}

type MinimalModel struct {
    CreatedAt time.Time      `json:"-" gorm:"index"`
    UpdatedAt time.Time      `json:"-" gorm:"index"`

    DB        *gorm.DB       `json:"-" gorm:"-"`
}

type Model struct {
    MinimalModel
    ID        uint      `json:"id,omitempty" gorm:"primaryKey,auto_increment,NOT NULL"`
}

func (model *Model) TableName() string {
    return "models"
}

func (model *Model) SetID(id uint) error {
    model.ID = id

    ModelSave(model)
    return nil
}

func (model *Model) GetID() (uint, error) {
    return model.ID, nil
}

func (model *Model) SetDB(db *gorm.DB) error {
    model.DB = db
    return nil
}

func (model *Model) GetDB() (*gorm.DB, error) {
    return model.DB, nil
}

func (model *Model) Query(query ...interface{}) *gorm.DB {
    db := model.GetDB()
    return db.Table(model.TableName()).Select(query...)
}

func Create(model ModelSkel) error {
    e := model.GetDB().Create(model)
    return e.Error
}

func Delete(model ModelSkel) error {
    e := model.GetDB().Delete(model)
    return e.Error
}

func Save(model ModelSkel) error {
    e := model.GetDB().Save(model)
    return e.Error
}

func Update(model ModelSkel, columns Dict) error {
    e := model.GetDB().First(model).Updates(columns.ToStrMap())
    return e.Error
}

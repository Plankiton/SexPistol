package sex

import (
    "time"
    "gorm.io/gorm"
)

type IModel interface {
    Create ()
    Delete ()
    Update (Columns Dict)
}

type ModelNoID struct {
    CreatedAt time.Time      `json:"-" gorm:"index"`
    UpdatedAt time.Time      `json:"-" gorm:"index"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
    ModelType string         `json:"class,omitempty"`
}

type Model struct {
    ModelNoID
    ID        uint      `json:"id,omitempty" gorm:"primaryKey,auto_increment,NOT NULL"`
}

func ModelCreate(model IModel) error {
    e := _database.Create(model)
    return e.Error
}

func ModelDelete(model IModel) error {
    e := _database.Delete(model)
    return e.Error
}

func ModelSave(model IModel) error {
    e := _database.Save(model)
    return e.Error
}

func ModelUpdate(model IModel, columns Dict) error {
    e := _database.First(model).Updates(columns.ToStrMap())
    return e.Error
}

package sex

import (
    "time"
)

type IModel interface {
    Create () bool
    Delete () bool
    Save   () bool
    Update (Columns Dict) bool
}


type ModelNoID struct {
    CreatedAt time.Time      `json:"created_at" gorm:"index"`
    UpdatedAt time.Time      `json:"updated_at" gorm:"index"`
    ModelType string         `json:"-"`
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

package SexDatabase

import (
    "time"
    "gorm.io/gorm"
)

type ModelSkel interface {
    TableName() string

    GetID() (interface{}, error)
    SetID(interface{}) error

    New() error
    Del() error
    Upd() error
}

type Model struct {
    ID        uint      `json:"id,omitempty" gorm:"primaryKey,auto_increment,NOT NULL"`
    CreatedAt time.Time `json:"-" gorm:"index"`
    UpdatedAt time.Time `json:"-" gorm:"index"`

    DB        *gorm.DB  `json:"-" gorm:"-"`
}

func (model Model) New() error {
    return nil
}
func (model Model) Del() error {
    return nil
}
func (model Model) Upd() error {
    return nil
}

func (model Model) SetID(id interface{}) error {
    model.ID = id.(uint)
    return nil
}

func (model Model) GetID() (interface{}, error) {
    return model.ID, nil
}

package SexDatabase

import (
    "time"
)

type ModelSkel interface {
    TableName() string

    GetID() (interface{}, error)
    SetID(interface{}) error

    New() error
    Del() error
    Upd() error
}

type MinimalModel struct {
}

type MetaModel struct {
    MinimalModel
    CreatedAt time.Time `json:"-" gorm:"index"`
    UpdatedAt time.Time `json:"-" gorm:"index"`
}

type Model struct {
    MetaModel

    ID        uint `json:"id,omitempty" gorm:"primaryKey,auto_increment,NOT NULL"`
}

func (model MinimalModel) New() error {
    return nil
}
func (model MinimalModel) Del() error {
    return nil
}
func (model MinimalModel) Upd() error {
    return nil
}

func (model MinimalModel) SetID(id interface{}) error {
    return nil
}

func (model MinimalModel) GetID() (interface{}, error) {
    return nil, nil
}

func (model Model) SetID(id interface{}) error {
    model.ID = id.(uint)
    return nil
}

func (model Model) GetID() (interface{}, error) {
    return model.ID, nil
}

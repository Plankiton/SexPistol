package SexDB

import (
    "time"
)

// Database model skeleton interface
type ModelSkel interface {
    TableName() string

    GetID() (interface{}, error)
    SetID(interface{}) error

    New() error
    Del() error
    Upd() error
}

// MinimalModel template
type MinimalModel struct {
}

// MetaModel template
type MetaModel struct {
    MinimalModel
    CreatedAt time.Time `json:"-" gorm:"index"`
    UpdatedAt time.Time `json:"-" gorm:"index"`
}

// Default Model template
type Model struct {
    MetaModel

    ID        uint `json:"id,omitempty" gorm:"primaryKey,auto_increment,NOT NULL"`
}

// Function executed before to create new data on database
// If it returns a error, data are not created
func (model MinimalModel) New() error {
    return nil
}

// Function executed before to delete data from database
// If it returns a error, data are not deleted
func (model MinimalModel) Del() error {
    return nil
}

// Function executed before to update data on database
// If it returns a error, data are not saved
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

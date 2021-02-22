package api

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

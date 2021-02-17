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
    CreatedAt time.Time `gorm:"index"`
    UpdatedAt time.Time `gorm:"index"`
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Model struct {
    gorm.Model
    ModelType string
}

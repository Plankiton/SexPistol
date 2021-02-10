package api

import (
    "time"
)

type Model struct {
    ID         string    `json:"id,omitempty" gorm:"primaryKey"`
    CreateAt time.Time   `json:"created_at,omitempty"`
    UpdateAt time.Time   `json:"updated_at,omitempty"`
}

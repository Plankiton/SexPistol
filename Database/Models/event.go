package SexDB

import (
    "time"
    "github.com/plankiton/SexPistol/Database"
)

type Event struct {
    SexDB.Model
    Name       string       `json:"name,omitempty" gorm:"index"`
    Type       string       `json:"type,omitempty" gorm:"index"`

    AddrId     uint         `json:"-" gorm:"index"`
    CoverId    uint         `json:"-" gorm:"index"`

    BeginAt    time.Time    `json:"begin,omitempty" gorm:"index"`
    EndAt      time.Time    `json:"end,omitempty" gorm:"index"`
}

func (m Event) TableName() string {
    return "events"
}

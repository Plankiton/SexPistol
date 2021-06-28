package SexModel

import (
    "time"
    "github.com/Plankiton/SexPistol/Cartridge"
)

type Event struct {
    SexDB.Model
    Name       string       `json:"name,omitempty"`
    Type       string       `json:"type,omitempty"`

    AddrId     uint         `json:"-"`
    Addr       Address      `json:"address,omitempty" gorm:"-"`

    BeginAt    time.Time    `json:"begin,omitempty"`
    EndAt      time.Time    `json:"end,omitempty"`
}

func (m Event) TableName() string {
    return "events"
}

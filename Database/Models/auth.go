package SexDB
import "github.com/plankiton/SexPistol/Database"

type Token struct {
    SexDB.Model
    ID        string `json:"Token,omitempty" gorm:"PrimaryKey, NOT NULL"`

    UserId    uint   `json:"-" gorm:"index, NOT NULL"`
}

func (m Token) TableName() string {
    return "tokens"
}

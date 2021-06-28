package SexModel
import "github.com/plankiton/SexPistol/Database"

type Token struct {
    SexDB.Model
    ID        string `json:"token,omitempty" gorm:"PrimaryKey, NOT NULL"`

    UserId    uint   `json:"-" gorm:"index, NOT NULL"`
    User      User   `json:"user,omitempty" gorm:"-"`
}

func (m Token) TableName() string {
    return "tokens"
}

package SexModel
import "github.com/plankiton/SexPistol/Database"

type Group struct {
    SexDB.Model
    Name      string  `json:"name,omitempty" gorm:"index"`
}

func (m Group) TableName() string {
    return "groups"
}

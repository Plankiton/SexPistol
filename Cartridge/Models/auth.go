package SexModel
import "github.com/Plankiton/SexPistol/Cartridge"

// Sex User Token Template
type Token struct {
    SexDB.Model
    ID        string `json:"token,omitempty" gorm:"PrimaryKey, NOT NULL"`

    UserId    uint   `json:"-" gorm:"index,NOT NULL"`
    // User      User   `json:"user,omitempty" gorm:"foreignKey:UserId"`
}

func (m Token) TableName() string {
    return "tokens"
}

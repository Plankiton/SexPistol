package SexModel
import "github.com/Plankiton/SexPistol/Cartridge"

type Role struct {
    SexDB.Model
    Name    string  `json:"name,omitempty" gorm:"unique"`
    Desc    string  `json:"desc,omitempty" gorm:"unique"`

    Users   []*User `json:"roles,omitempty" gorm:"many2many:user_roles"`
}

type LinkRoleUser struct {
    Roles     []*Role `json:"roles,omitempty" gorm:"many2many:user_roles"`
}

func (m Role) TableName() string {
    return "roles"
}

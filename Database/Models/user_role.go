package SexDB
import "github.com/plankiton/SexPistol/Database"

type UserRole struct {
    SexDB.Model
    UserId  uint  `json:"-" gorm:"index"`
    RoleId  uint  `json:"-" gorm:"index"`
}

func (m UserRole) TableName() string {
    return "link_user_roles"
}

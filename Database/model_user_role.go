package SexDatabase

type UserRole struct {
    Model
    UserId  uint  `json:"-" gorm:"index"`
    RoleId  uint  `json:"-" gorm:"index"`
}

func (m UserRole) TableName() string {
    return "link_user_roles"
}

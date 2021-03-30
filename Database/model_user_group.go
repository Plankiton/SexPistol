package SexDatabase

type UserGroup struct {
    Model
    UserId   uint  `json:"-" gorm:"index"`
    GroupId  uint  `json:"-" gorm:"index"`
}

func (m UserGroup) TableName() string {
    return "link_user_groups"
}

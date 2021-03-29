package SexDatabase

type Group struct {
    Model
    Name      string  `json:"name,omitempty" gorm:"index"`
}

func (m Group) TableName() string {
    return "groups"
}

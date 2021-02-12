package api

type Role struct {
    Model
    Name    string  `json:"name,omitempty" gorm:"unique"`

    Users []interface {}
}

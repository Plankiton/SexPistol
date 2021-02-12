package api

type Auth struct {
    Model

    ID        uint `json:"auth,omitempty" gorm:"PrimaryKey"`
    UserId    uint `json:",empty"`
}

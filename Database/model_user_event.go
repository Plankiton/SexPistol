package SexDatabase

type UserEvent struct {
    Model
    UserId   uint  `json:"-" gorm:"index"`
    EventId  uint  `json:"-" gorm:"index"`
}

func (m UserEvent) TableName() string {
    return "link_user_events"
}

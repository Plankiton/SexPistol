package sex

type UserEvent struct {
    Model
    UserId   uint  `json:"-" gorm:"index"`
    EventId  uint  `json:"-" gorm:"index"`
}

func (m UserEvent) TableName() string {
    return "user_events"
}

package SexDatabase

type Token struct {
    MinimalModel
    ID        string `json:"Token,omitempty" gorm:"PrimaryKey, NOT NULL"`

    UserId    uint   `json:"-" gorm:"index, NOT NULL"`
}

func (m Token) TableName() string {
    return "tokens"
}

func (model *Token) Verify() bool {
    if model.ID == "" {
        return false
    }

    if _database.First(model, "id = ?", model.ID).Error == nil {
        return true
    }

    return false
}

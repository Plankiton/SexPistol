package sex

type Token struct {
    ModelNoID
    ID        string `json:"Token,omitempty" gorm:"PrimaryKey, NOT NULL"`

    UserId    uint   `json:"-" gorm:"index, NOT NULL"`
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

func (model *Token) Create() bool {
    if (model.ModelType == "") {
        model.ModelType = GetModelType(model)
    }

    if ModelCreate(model) == nil {
        ID := model.ID
        ModelType := model.ModelType
        Log("Created", ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *Token) Delete() bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelCreate(model) == nil {
        Log("Deleted", ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *Token) Save() bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelSave(model) == nil {
        Log("Updated", ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *Token) Update(columns Dict) bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelUpdate(model, columns) == nil {
        Log("Updated", ToLabel(ID, ModelType))
        return true
    }

    return false
}

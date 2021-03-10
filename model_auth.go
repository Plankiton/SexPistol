package sex

type Token struct {
    ModelNoID
    ID        string `json:"Token,omitempty" gorm:"PrimaryKey"`
    ModelType string

    UserId    uint   `json:"-" gorm:"index"`
}

func (model *Token) Verify() bool {
    if _database.First(model, "id = ?", model.ID).Error == nil {
        return true
    }

    return false
}

func (model *Token) Create() {
    if (model.ModelType == "") {
        model.ModelType = GetModelType(model)
    }

    if ModelCreate(model) == nil {
        ID := model.ID
        ModelType := model.ModelType
        Log("Created", ToLabel(ID, ModelType))
    }
}

func (model *Token) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    if ModelCreate(model) == nil {
        Log("Deleted", ToLabel(ID, ModelType))
    }
}

func (model *Token) Save() {
    ID := model.ID
    ModelType := model.ModelType

    if ModelSave(model) == nil {
        Log("Updated", ToLabel(ID, ModelType))
    }
}

func (model *Token) Update(columns Dict) {
    ID := model.ID
    ModelType := model.ModelType

    if ModelUpdate(model, columns) == nil {
        Log("Updated", ToLabel(ID, ModelType))
    }
}

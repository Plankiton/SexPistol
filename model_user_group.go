package sex

type UserGroup struct {
    Model
    UserId   uint  `json:"-" gorm:"index"`
    GroupId  uint  `json:"-" gorm:"index"`
}

func (model *UserGroup) Create() bool {
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

func (model *UserGroup) Delete() bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelCreate(model) == nil {
        Log("Deleted", ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *UserGroup) Save() bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelSave(model) == nil {
        Log("Updated", ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *UserGroup) Update(columns Dict) bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelUpdate(model, columns) == nil {
        Log("Updated", ToLabel(ID, ModelType))
        return true
    }

    return false
}

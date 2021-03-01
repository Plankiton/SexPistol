package api

type UserRole struct {
    Model
    UserId  uint  `json:"-" gorm:"index"`
    RoleId  uint  `json:"-" gorm:"index"`
}

func (model *UserRole) Create() {
    if (model.ModelType == "") {
        model.ModelType = GetModelType(model)
    }

    if ModelCreate(model) == nil {
        ID := model.ID
        ModelType := model.ModelType
        Log("Created", ToLabel(ID, ModelType))
    }
}

func (model *UserRole) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    if ModelCreate(model) == nil {
        Log("Deleted", ToLabel(ID, ModelType))
    }
}

func (model *UserRole) Save() {
    ID := model.ID
    ModelType := model.ModelType

    if ModelSave(model) == nil {
        Log("Updated", ToLabel(ID, ModelType))
    }
}

func (model *UserRole) Update(columns Dict) {
    ID := model.ID
    ModelType := model.ModelType

    if ModelUpdate(model, columns) == nil {
        Log("Updated", ToLabel(ID, ModelType))
    }
}

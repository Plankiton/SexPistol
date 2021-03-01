package api

type UserEvent struct {
    Model
    UserId   uint  `json:"-" gorm:"index"`
    EventId  uint  `json:"-" gorm:"index"`
}

func (model *UserEvent) Create() {
    if (model.ModelType == "") {
        model.ModelType = GetModelType(model)
    }

    if ModelCreate(model) == nil {
        ID := model.ID
        ModelType := model.ModelType
        Log("Created", ToLabel(ID, ModelType))
    }
}

func (model *UserEvent) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    if ModelCreate(model) == nil {
        Log("Deleted", ToLabel(ID, ModelType))
    }
}

func (model *UserEvent) Save() {
    ID := model.ID
    ModelType := model.ModelType

    if ModelSave(model) == nil {
        Log("Updated", ToLabel(ID, ModelType))
    }
}

func (model *UserEvent) Update(columns Dict) {
    ID := model.ID
    ModelType := model.ModelType

    if ModelUpdate(model, columns) == nil {
        Log("Updated", ToLabel(ID, ModelType))
    }
}

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

    _database.Create(model)

    e := _database.First(model)
    if e.Error == nil {

        ID := model.ID
        ModelType := model.ModelType
        Log("Created", ToLabel(ID, ModelType))
    }
}

func (model *UserEvent) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Delete(model)
        Log("Deleted", ToLabel(ID, ModelType))
    }
}

func (model *UserEvent) Save() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(&User{}, "id = ?", model.ID)
    if e.Error == nil {
        _database.Save(model)
        Log("Updated", ToLabel(ID, ModelType))
    }
}

func (model *UserEvent) Update(columns Dict) {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(&User{}, "id = ?", model.ID)
    if e.Error == nil {
        _database.First(model).Updates(columns.ToStrMap())
        Log("Updated", ToLabel(ID, ModelType))
    }
}

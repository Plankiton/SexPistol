package api

type UserGroup struct {
    Model
    UserId   uint  `json:"-" gorm:"index"`
    GroupId  uint  `json:"-" gorm:"index"`
}

func (model *UserGroup) Create() {
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

func (model *UserGroup) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Delete(model)
        Log("Deleted", ToLabel(ID, ModelType))
    }
}

func (model *UserGroup) Save() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Save(model)
        Log("Updated", ToLabel(ID, ModelType))
    }
}

func (model *UserGroup) Update(columns Dict) {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.First(model).Updates(columns.ToStrMap())
        Log("Updated", ToLabel(ID, ModelType))
    }
}

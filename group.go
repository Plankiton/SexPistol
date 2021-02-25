package api

type Group struct {
    Model
    Name      string  `json:"name,omitempty" gorm:"index"`
}

func (model *Group) Create() {
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

func (model *Group) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Delete(model)
        Log("Deleted", ToLabel(ID, ModelType))
    }
}

func (model *Group) Save() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(&Group{}, "id = ?", model.ID)
    if e.Error == nil {
        _database.Save(model)
        Log("Updated", ToLabel(ID, ModelType))
    }
}

func (model *Group) Update(columns Dict) {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(&Group{}, "id = ?", model.ID)
    if e.Error == nil {
        _database.First(model).Updates(columns.ToStrMap())
        Log("Updated", ToLabel(ID, ModelType))
    }
}

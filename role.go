package api

type Role struct {
    Model
    Name    string  `json:"name,omitempty" gorm:"unique"`
}

func (model *Role) Create() {
    model.ModelType = GetModelType(model)

    _database.Create(model)

    e := _database.First(model)
    if e.Error == nil {

        ID := model.ID
        ModelType := model.ModelType
        Log("Created", ToLabel(ID, ModelType))
    }
}

func (model *Role) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Delete(model)
        Log("Deleted", ToLabel(ID, ModelType))
    }
}

func (model *Role) Save() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(&Role{}, "id = ?", model.ID)
    if e.Error == nil {
        _database.Save(model)
        Log("Updated", ToLabel(ID, ModelType))
    }
}

func (model *Role) Update(columns Dict) {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(&Role{}, "id = ?", model.ID)
    if e.Error == nil {
        _database.First(model).Updates(columns.ToStrMap())
        Log("Updated", ToLabel(ID, ModelType))
    }
}

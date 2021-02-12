package api

type Auth struct {
    Model

    ID        uint `json:"auth,omitempty" gorm:"PrimaryKey"`
    UserId    uint `json:",empty"`
}

func (model *Auth) Create() {
    ID := model.ID
    ModelType := model.ModelType

    _database.Create(model)

    e := _database.First(model, ID)
    if e.Error == nil {
        Log("Created", ToLabel(ID, ModelType))
    }
}

func (model *Auth) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model, ID)
    if e.Error == nil {
        _database.Delete(model)
        Log("Deleted", ToLabel(ID, ModelType))
    }
}

func (model *Auth) Update(columns Dict) {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model, ID)
    if e.Error == nil {
        _database.First(model, ID).Updates(columns.ToStrMap())
        Log("Updated", ToLabel(ID, ModelType))
    }
}

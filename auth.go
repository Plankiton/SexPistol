package api

type Auth struct {
    Model

    ID        uint `json:"auth,omitempty" gorm:"PrimaryKey"`
    RoleId    uint `json:",empty"`
}

func (model *Auth) Verify(token uint) bool {
    if _database.First(model, token).Error == nil {
        return true
    }

    return false
}

func (model *Auth) Create() {
    ID := model.ID
    ModelType := model.ModelType

    _database.Create(model)

    e := _database.First(model)
    if e.Error == nil {
        Log("Created", ToLabel(ID, ModelType))
    }
}

func (model *Auth) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Delete(model)
        Log("Deleted", ToLabel(ID, ModelType))
    }
}

func (model *Auth) Save() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Save(model)
        Log("Updated", ToLabel(ID, ModelType))
    }
}


func (model *Auth) Update(columns Dict) {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.First(model).Updates(columns.ToStrMap())
        Log("Updated", ToLabel(ID, ModelType))
    }
}

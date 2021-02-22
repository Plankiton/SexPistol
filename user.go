package api

import (
    "time"
)

type User struct {
    Model

    Phone      string    `json:"phone,omitempty" gorm:"index,default:null"`
    Email      string    `json:"email,omitempty" gorm:"index,default:null"`
    Name       string    `json:"name,omitempty" gorm:"index"`
    Born       time.Time `json:"born_date,omitempty" gorm:"index"`
    Genre      string    `json:"genre,omitempty" gorm:"default:'M'"`
    PassHash   string    `json:"-"`
}

func (model *User) CheckPass(s string) bool {
    byteHash := []byte(model.PassHash)
    err := CheckPass(byteHash, s)
    if err != nil {
        return false
    }
    return true
}

func (model *User) SetPass(s string) (string, error) {
    hash, err := ToPassHash(s)
    if err != nil {
        return "", nil
    }

    model.PassHash = hash
    return model.PassHash, nil
}

func (model *User) Create() {
    model.ModelType = GetModelType(model)

    _database.Create(model)

    e := _database.First(model)
    if e.Error == nil {


        ID := model.ID
        ModelType := model.ModelType
        Log("Created", ToLabel(ID, ModelType))
    }
}

func (model *User) Delete() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Delete(model)
        Log("Deleted", ToLabel(ID, ModelType))
    }
}

func (model *User) Save() {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.Save(model)
        Log("Updated", ToLabel(ID, ModelType))
    }
}

func (model *User) Update(columns Dict) {
    ID := model.ID
    ModelType := model.ModelType

    e := _database.First(model)
    if e.Error == nil {
        _database.First(model).Updates(columns.ToStrMap())
        Log("Updated", ToLabel(ID, ModelType))
    }
}

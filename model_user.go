package sex

import (
    "time"
)

type User struct {
    Model

    Phone      string    `json:"phone,omitempty" gorm:"unique,default:null"`
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

func (model *User) Create() bool {
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

func (model *User) Delete() bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelCreate(model) == nil {
        Log("Deleted", ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *User) Save() bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelSave(model) == nil {
        Log("Updated", ToLabel(ID, ModelType))
        return true
    }

    return false
}

func (model *User) Update(columns Dict) bool {
    ID := model.ID
    ModelType := model.ModelType

    if ModelUpdate(model, columns) == nil {
        Log("Updated", ToLabel(ID, ModelType))
        return true
    }

    return false
}

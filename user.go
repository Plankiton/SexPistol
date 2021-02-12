package api

import (
    "gorm.io/gorm"
)

type Person struct {
    gorm.Model
    ModelType  string    `json:"model_type,omitempty" gorm:"default:'User'"`

    Document   string    `json:"doc,omitempty" gorm:"uniqueIndex"`
    Phone      string    `json:"phone,omitempty" gorm:"index,default:null"`
    Name       string    `json:"name,omitempty" gorm:"index"`

    PassHash   string    `json:",empty"`
}

func (model *Person) CheckPass(s string) bool {
    byteHash := []byte(model.PassHash)
    err := CheckPass(byteHash, s)
    if err != nil {
        return false
    }
    return true
}

func (model *Person) SetPass(s string) (string, error) {
    hash, err := ToPassHash(s)
    if err != nil {
        return "", nil
    }

    model.PassHash = hash
    return model.PassHash, nil
}


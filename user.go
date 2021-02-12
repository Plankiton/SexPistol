package api

import (
    "gorm.io/gorm"
)

type Person struct {
    gorm.Model
    ModelType  string `gorm:"default:'User'"`

    Document   string    `json:"doc,omitempty" gorm:"uniqueIndex"`
    Phone      string    `json:"phone,omitempty" gorm:"index,default:null"`
    Name       string    `json:"name,omitempty" gorm:"index"`

    PassHash   string    `json:",empty"`
}

func (self *Person) CheckPass(s string) bool {
    byteHash := []byte(self.PassHash)
    err := CheckPass(byteHash, s)
    if err != nil {
        return false
    }
    return true
}

func (self *Person) SetPass(s string) (string, error) {
    hash, err := ToPassHash(s)
    if err != nil {
        return "", nil
    }

    self.PassHash = hash
    return self.PassHash, nil
}

func (self *Person) AddTo() {
    e := _database.First(self)

    if e.Error != nil {
        _database.Create(self)
        Log("Created <", self.ModelType, " ", self.Name," :", self.ID,"> !!")
    }
}

func (self *Person) DelFrom() {
    e := _database.First(self)

    if e.Error == nil {
        _database.Delete(self)
        Log("Deleted <", self.ModelType, " ", self.Name," :", self.ID,"> !!")
    }
}

func (self *Person) SetOn() {
    e := _database.First(self)

    if e.Error == nil {
        _database.Delete(self)
        Log("Updated <", self.ModelType, " ", self.Name," :", self.ID,"> !!")
    }
}

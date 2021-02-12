package api

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

var _database * gorm.DB
func CreateDB(con_string string) (*gorm.DB, error) {
    dsn := getEnv("DB_URI", con_string)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return db, err
    }

    _database = db
    _database.Migrator().CurrentDatabase()
    return _database, err
}

func Add(model interface {}) {
    e := _database.First(model)

    if e.Error != nil {
        _database.Create(model)
        Log("Created <", model.(struct{ ModelType interface{} }).ModelType, " ", model.(struct{ Name interface{} }).Name," :", model.(struct{ ID interface{} }).ID,"> !!")
    }
}

func Del(model interface{}) {
    e := _database.First(model)

    if e.Error == nil {
        _database.Delete(model)
        Log("Deleted <", model.(struct{ ModelType interface{} }).ModelType, " ", model.(struct{ Name interface{} }).Name," :", model.(struct{ ID interface{} }).ID,"> !!")
    }
}

func Set(model interface{}) {
    e := _database.First(model)

    if e.Error == nil {
        _database.Delete(model)
        Log("Updated <", model.(struct{ ModelType interface{} }).ModelType, " ", model.(struct{ Name interface{} }).Name," :", model.(struct{ ID interface{} }).ID,"> !!")
    }
}

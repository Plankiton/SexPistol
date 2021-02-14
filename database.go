package api

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "gorm.io/driver/postgres"
)

var _database * gorm.DB
func Postgres(con_string string) (*gorm.DB, error) {
    dsn := GetEnv("DB_URI", con_string)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true,
    })
    if err != nil {
        return db, err
    }

    _database = db
    _database.Migrator().CurrentDatabase()
    return _database, err
}

func Sqlite(con_string string) (*gorm.DB, error) {
    dsn := GetEnv("DB_URI", con_string)
    db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true,
    })
    if err != nil {
        return db, err
    }

    _database = db
    _database.Migrator().CurrentDatabase()
    return _database, err
}

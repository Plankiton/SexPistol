package database

import (
    "os"
    "log"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/driver/sqlite"
    "gorm.io/driver/postgres"
)

type Database struct {
    gorm.DB
}

func Postgres(dsn string) gorm.Dialector {
    return postgres.Open(dsn)
}

func Sqlite(dsn string) gorm.Dialector {
    return sqlite.Open(dsn)
}

func (db *Database) SetLogLevel(s string) {
    db.Config.logger = Logger(s)
}

func (db *Database) AddModels(m ...IModel) {
    db.AutoMigrate(m...)
}

func Create(con_string string, db_type func(string)(gorm.Dialector)) (*Database, error) {
    logger := Logger()

    dsn := GetEnv("DB_URI", con_string)
    gorm_db, err := gorm.Open(db_type(dsn), &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true,
        Logger: logger,
    })

    db := &Database {
        DB: gorm_db,
    }
    return db, err
}

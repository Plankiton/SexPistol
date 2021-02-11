package api

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

func CreateDB(con_string string) (*gorm.DB, error) {
    dsn := getEnv("DB_URI", con_string)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return db, err
    }

    return db, err
}

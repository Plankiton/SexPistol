package SexDatabase

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "gorm.io/driver/postgres"

    "os"
)

type DatabaseSkel interface {
    SetLogLevel()
    Create(ModelSkel) error
    Delete(ModelSkel) error
    Save(ModelSkel) error
    Update(ModelSkel, map[string]interface{}) error
}

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
    db.Config.Logger = Logger(s)
}

func (db *Database) AddModels(m ...ModelSkel) {
    for _, model := range m {
        db.AutoMigrate(model)
    }
}

func (db *Database) Create(model ModelSkel) error {
    e := model.New()
    if e == nil {
        e := db.DB.Create(model).Error
        return e
    }

    return e
}

func (db *Database) Delete(model ModelSkel) error {
    e := db.DB.Delete(model).Error
    return e
}

func (db *Database) Save(model ModelSkel) error {
    e := db.DB.Save(model).Error
    return e
}

func (db *Database) Update(model ModelSkel, columns map[string]interface{}) error {
    e := db.DB.First(model).Updates(columns).Error
    return e
}

func Open(con_string string, db_type func(string)(gorm.Dialector)) (*Database, error) {
    logger := Logger()

    dsn := os.Getenv("DB_URI")
    if dsn == "" {
        dsn = con_string
    }

    gorm_db, err := gorm.Open(db_type(dsn), &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true,
        Logger: logger,
    })

    db := &Database {
        DB: *gorm_db,
    }
    return db, err
}

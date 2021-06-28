package SexDB

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "gorm.io/driver/postgres"

    "os"
)

type CartridgeSkel interface {
    SetLogLevel()
    Add(ModelSkel) error
    Del(ModelSkel) error
    Sav(ModelSkel) error
    Set(ModelSkel, map[string]interface{}) error
}

type Cartridge struct {
    gorm.DB
}

func ToDB(db *gorm.DB) Cartridge {
    return Cartridge {
        DB: *db,
    }
}

func Postgres(dsn string) gorm.Dialector {
    return postgres.Open(dsn)
}

func Sqlite(dsn string) gorm.Dialector {
    return sqlite.Open(dsn)
}

func (db *Cartridge) SetLogLevel(s string) {
    db.Config.Logger = Logger(s)
}

func (db *Cartridge) AddModels(m ...ModelSkel) {
    for _, model := range m {
        db.AutoMigrate(model)
    }
}

func (db *Cartridge) Add(model ModelSkel) error {
    e := model.New()
    if e == nil {
        e := db.Create(model).Error
        return e
    }

    return e
}

func (db *Cartridge) Del(model ModelSkel) error {
    e := db.Delete(model).Error
    return e
}

func (db *Cartridge) Sav(model ModelSkel) error {
    e := db.Save(model).Error
    return e
}

func (db *Cartridge) Set(model ModelSkel, columns map[string]interface{}) error {
    e := db.First(model).Updates(columns).Error
    return e
}

func Open(con_string string, db_type func(string)(gorm.Dialector)) (*Cartridge, error) {
    logger := Logger()

    dsn := os.Getenv("DB_URI")
    if dsn == "" {
        dsn = con_string
    }

    gorm_db, err := gorm.Open(db_type(dsn), &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true,
        Logger: logger,
    })

    db := &Cartridge {
        DB: *gorm_db,
    }
    return db, err
}

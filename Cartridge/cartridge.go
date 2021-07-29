package SexDB

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "gorm.io/driver/postgres"

    "os"
)

// Cartridge Skeleton interface
type CartridgeSkel interface {
    SetLogLevel()
    Add(ModelSkel) error
    Del(ModelSkel) error
    Sav(ModelSkel) error
    Set(ModelSkel, map[string]interface{}) error
}

// Sex Cartridge type
type Cartridge struct {
    gorm.DB
}

// Converte *gorm.DB in a SexDB.Cartridge
func ToDB(db *gorm.DB) *Cartridge {
    return &Cartridge {
        DB: *db,
    }
}

// Postgres Sql Dialector
func Postgres(dsn string) gorm.Dialector {
    return postgres.Open(dsn)
}

// Sqlite Sql Dialector
func Sqlite(dsn string) gorm.Dialector {
    return sqlite.Open(dsn)
}

// Database Log Level Setting
// Log Level can be "error", "warn", "info"
func (db *Cartridge) SetLogLevel(s string) {
    db.Config.Logger = Logger(s)
}

// Add models to auto migrate
func (db *Cartridge) AddModels(m ...ModelSkel) {
    for _, model := range m {
        db.AutoMigrate(model)
    }
}

// Alias to create model data of database
func (db *Cartridge) Add(model ModelSkel) error {
    e := model.New()
    if e == nil {
        e := db.Create(model).Error
        return e
    }

    return e
}

// Alias to delete model data of database
func (db *Cartridge) Del(model ModelSkel) error {
    e := model.Del()
    if e == nil {
        e := db.Delete(model).Error
        return e
    }
    return e
}

// Alias to save model data to database
func (db *Cartridge) Sav(model ModelSkel) error {
    e := model.Upd()
    if e == nil {
        e := db.Save(model).Error
        return e
    }
    return e
}

// Alias to update expecific columns on database
func (db *Cartridge) Set(model ModelSkel, columns map[string]interface{}) error {
    e := model.Upd()
    if e == nil {
        e := db.First(model).Updates(columns).Error
        return e
    }
    return e
}

// Database connector
// Required: con_string need to be complatible with Database Driver format
// Example:
//    // import "gorm.io/driver/mysql"
//    db := SexDB.Open("test.db", SexDB.Sqlite)
//    db := SexDB.Open("host=server user=user password=pass dbname=test port=port", SexDB.Postgres)
//    db := SexDB.Open("user:pass@server:port/test", mysql.Open)
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

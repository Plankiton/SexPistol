package api

import (
    "os"
    "log"
    "time"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/driver/sqlite"
    "gorm.io/driver/postgres"
)

func Logger () logger.Interface {
    return logger.New(
        log.New(os.Stderr, "\r\n", log.LstdFlags), // io writer
        logger.Config{
            SlowThreshold: time.Second,   // Slow SQL threshold
            LogLevel:      logger.Error, // Log level
            Colorful:      true,
        },
    )
}

var _database * gorm.DB
func Postgres(con_string string) (*gorm.DB, error) {
    dsn := GetEnv("DB_URI", con_string)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true,
        Logger: Logger(),
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
        Logger: Logger(),
    })
    if err != nil {
        return db, err
    }

    _database = db
    _database.Migrator().CurrentDatabase()
    return _database, err
}

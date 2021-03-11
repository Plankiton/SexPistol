package sex

import (
    "os"
    "log"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/driver/sqlite"
    "gorm.io/driver/postgres"
)

func Logger () logger.Interface {
    return logger.New(
        log.New(os.Stderr, "\r\n", log.LstdFlags),
        logger.Config{
            SlowThreshold: 0,
            LogLevel:      logger.Error,
            Colorful:      true,
        },
    )
}

func (router *Pistol) SignDB(con_str string, createDB func (string) (*gorm.DB, error), models ...interface{}) (*gorm.DB, error) {
    db, err := createDB(con_str)
    router.Database = db

    if err != nil {
        Die("Error on creation of tables on database")
    }

    if models != nil {
        db.Migrator().CurrentDatabase()
        db.AutoMigrate(models...)
    }

    _database = db
    return router.Database, err
}

var _database * gorm.DB
func Postgres(con_string string) (*gorm.DB, error) {
    dsn := GetEnv("DB_URI", con_string)
    return gorm.Open(postgres.Open(dsn), &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true,
        Logger: Logger(),
    })
}

func Sqlite(con_string string) (*gorm.DB, error) {
    dsn := GetEnv("DB_URI", con_string)
    return gorm.Open(sqlite.Open(dsn), &gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true,
        Logger: Logger(),
    })
}

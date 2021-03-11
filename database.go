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

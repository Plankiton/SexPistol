package SexDB
import (
    "gorm.io/gorm/logger"

    "log"
    "os"
)

func Logger (level...string) logger.Interface {
    log_level := logger.Error
    if len(level) > 0 {
        levels :=  map[string]logger.LogLevel{
            "error": logger.Error,
            "warn": logger.Warn,
            "info": logger.Info,
        }

        if level, ok := levels[level[0]]; ok {
            log_level = level
        }
    }
    return logger.New(
        log.New(os.Stderr, "\r\n", log.LstdFlags),
        logger.Config{
            SlowThreshold: 0,
            LogLevel:      log_level,
            Colorful:      true,
        },
    )
}

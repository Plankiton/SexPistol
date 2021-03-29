package SexDatabase
import "gorm.io/gorm/logger"

func Logger (level...string) logger.Interface {
    log_level := logger.Error
    if len(level) > 0 {
        log_level, ok := map[string]interface{}{
            "error": logger.Error,
            "warn": logger.Warn,
            "info": logger.Info,
        }[level]
        if !ok {
            logger.Error
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

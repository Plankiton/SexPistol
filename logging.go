package Sex

import (
    "os"
    "log"
    "fmt"
)

func Fmt(s string, v...interface {}) string {
    return fmt.Sprintf(s, v...)
}

var logger *log.Logger
func Logger() *log.Logger {
    return logger
}

func Log (args ...interface {}) {
    fmt_args := []interface {}{"\033[32;1m[info] \033[00m"}
    fmt_args = append(fmt_args, args...)

    logger.Println(fmt_args...)
}

func Err (args ...interface {}) {
    if logger == nil {
        logger = log.New(os.Stderr, "\r\n", log.LstdFlags)
    }
    fmt_args := []interface {}{"\033[31;1m[erro] \033[00m"}
    fmt_args = append(fmt_args, args...)

    logger.Println(fmt_args...)
}

func War (args ...interface {}) {
    if logger == nil {
        logger = log.New(os.Stderr, "\r\n", log.LstdFlags)
    }
    fmt_args := []interface {}{"\033[33;1m[warn] \033[00m"}
    fmt_args = append(fmt_args, args...)

    logger.Println(fmt_args...)
}

func Die (args ...interface {}) {
    if logger == nil {
        logger = log.New(os.Stderr, "\r\n", log.LstdFlags)
    }
    fmt_args := []interface {}{"\033[31;1m[erro] \033[00m"}
    fmt_args = append(fmt_args, args...)

    logger.Println(fmt_args...)
    os.Exit(1)
}

func Logf (args ...interface {}) {
    if logger == nil {
        logger = log.New(os.Stderr, "\r\n", log.LstdFlags)
    }

    fmt_args := []interface {}{"\033[32;1m[info] \033[00m"}
    fmt_args = append(fmt_args, Fmt(args[0].(string), args[1:]...))

    logger.Println(fmt_args...)
}

func Errf (args ...interface {}) {
    if logger == nil {
        logger = log.New(os.Stderr, "\r\n", log.LstdFlags)
    }

    fmt_args := []interface {}{"\033[31;1m[erro] \033[00m"}
    fmt_args = append(fmt_args, Fmt(args[0].(string), args[1:]...))

    logger.Println(fmt_args...)
}

func Warf (args ...interface {}) {
    if logger == nil {
        logger = log.New(os.Stderr, "\r\n", log.LstdFlags)
    }

    fmt_args := []interface {}{"\033[33;1m[warn] \033[00m"}
    fmt_args = append(fmt_args, Fmt(args[0].(string), args[1:]...))

    logger.Println(fmt_args...)
}

func Dief (args ...interface {}) {
    if logger == nil {
        logger = log.New(os.Stderr, "\r\n", log.LstdFlags)
    }

    fmt_args := []interface {}{"\033[31;1m[erro] \033[00m"}
    fmt_args = append(fmt_args, Fmt(args[0].(string), args[1:]...))

    logger.Println(fmt_args...)
    os.Exit(1)
}

package Sex

import (
    "os"
    "log"
    "fmt"
)

// Geting formated string
func Fmt(s string, v...interface {}) string {
    return fmt.Sprintf(s, v...)
}

var logger *log.Logger = log.New(os.Stderr, "\r\n", log.LstdFlags)

// Get the logger from SexPistol
func Logger() *log.Logger {
    return logger
}

// Logging information logs with Sex.Logger()
func Log (args ...interface {}) {
    fmt_args := []interface {}{"\033[32;1m[info] \033[00m"}
    fmt_args = append(fmt_args, args...)

    logger.Println(fmt_args...)
}

// Logging error logs with Sex.Logger()
func Err (args ...interface {}) {
    fmt_args := []interface {}{"\033[31;1m[erro] \033[00m"}
    fmt_args = append(fmt_args, args...)

    logger.Println(fmt_args...)
}

// Logging warning logs with Sex.Logger()
func War (args ...interface {}) {
    fmt_args := []interface {}{"\033[33;1m[warn] \033[00m"}
    fmt_args = append(fmt_args, args...)

    logger.Println(fmt_args...)
}

// Logging error logs with Sex.Logger() and killing the application
func Die (args ...interface {}) {
    fmt_args := []interface {}{"\033[31;1m[erro] \033[00m"}
    fmt_args = append(fmt_args, args...)

    logger.Println(fmt_args...)
    os.Exit(1)
}

// Logging information formated logs with Sex.Logger()
// Example:
//    Logf("%s %+v", "joao", []string{"joao", "maria"})
//    Logf("%.2f", 409.845)
func Logf (args ...interface {}) {
    fmt_args := []interface {}{"\033[32;1m[info] \033[00m"}
    fmt_args = append(fmt_args, Fmt(args[0].(string), args[1:]...))

    logger.Println(fmt_args...)
}

// Logging error formated logs with Sex.Logger()
// Example:
//    Errf("%s %+v", "joao", []string{"joao", "maria"})
//    Errf("%.2f", 409.845)
func Errf (args ...interface {}) {
    fmt_args := []interface {}{"\033[31;1m[erro] \033[00m"}
    fmt_args = append(fmt_args, Fmt(args[0].(string), args[1:]...))

    logger.Println(fmt_args...)
}

// Logging warning formated logs with Sex.Logger()
// Example:
//    Warf("%s %+v", "joao", []string{"joao", "maria"})
//    Warf("%.2f", 409.845)
func Warf (args ...interface {}) {
    fmt_args := []interface {}{"\033[33;1m[warn] \033[00m"}
    fmt_args = append(fmt_args, Fmt(args[0].(string), args[1:]...))

    logger.Println(fmt_args...)
}

// Logging error logs with Sex.Logger() and killing the application
// Example:
//    Dief("%s %+v", "joao", []string{"joao", "maria"})
//    Dief("%.2f", 409.845)
func Dief (args ...interface {}) {
    fmt_args := []interface {}{"\033[31;1m[erro] \033[00m"}
    fmt_args = append(fmt_args, Fmt(args[0].(string), args[1:]...))

    logger.Println(fmt_args...)
    os.Exit(1)
}

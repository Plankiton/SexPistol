package Sex

import (
    "os"
    "log"
    "fmt"
    "runtime"
)

// Geting formated string
func Fmt(s string, v...interface {}) string {
    return fmt.Sprintf(s, v...)
}

var logger *log.Logger = log.New(os.Stderr, "\r\n", log.LstdFlags | log.Lmicroseconds)

// Get the logger from SexPistol
func Logger() *log.Logger {
    return logger
}

// Logging a raw log
func RawLog(typ string, args...interface{}) {
    _, file, line, _ := runtime.Caller(1)
    caller := Fmt("%s:%d", file, line)

    fmt_args := []interface {}{caller, Fmt("%s", typ)}
    fmt_args = append(fmt_args, fmt.Sprint(args...))

    logger.Println(fmt_args...)
}

// Logging information logs with Sex.Logger()
func Log (args ...interface {}) {
    RawLog("\033[32;1m[info] \033[00m", args...)
}

// Logging error logs with Sex.Logger()
func Err (args ...interface {}) {
    RawLog("\033[31;1m[erro] \033[00m", args...)
}

// Logging warning logs with Sex.Logger()
func War (args ...interface {}) {
    RawLog("\033[33;1m[warn] \033[00m", args...)
}

// Logging error logs with Sex.Logger() and killing the application
func Die (args ...interface {}) {
    RawLog("\033[31;1m[erro] \033[00m", args...)
    os.Exit(1)
}

// Logging information formated logs with Sex.Logger()
// Example:
//    Logf("%s %+v", "joao", []string{"joao", "maria"})
//    Logf("%.2f", 409.845)
func Logf (args ...interface {}) {
    RawLog("\033[32;1m[info] \033[00m", Fmt(args[0].(string), args[1:]...))
}

// Logging error formated logs with Sex.Logger()
// Example:
//    Errf("%s %+v", "joao", []string{"joao", "maria"})
//    Errf("%.2f", 409.845)
func Errf (args ...interface {}) {
    RawLog("\033[31;1m[erro] \033[00m", Fmt(args[0].(string), args[1:]...))
}

// Logging warning formated logs with Sex.Logger()
// Example:
//    Warf("%s %+v", "joao", []string{"joao", "maria"})
//    Warf("%.2f", 409.845)
func Warf (args ...interface {}) {
    RawLog("\033[33;1m[warn] \033[00m", Fmt(args[0].(string), args[1:]...))
}

// Logging error logs with Sex.Logger() and killing the application
// Example:
//    Dief("%s %+v", "joao", []string{"joao", "maria"})
//    Dief("%.2f", 409.845)
func Dief (args ...interface {}) {
    RawLog("\033[31;1m[erro] \033[00m", Fmt(args[0].(string), args[1:]...))
    os.Exit(1)
}

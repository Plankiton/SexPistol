package Sex

import (
    "fmt"
    "log"
    "os"
    "runtime"
)

const (
   LogLevelInfo = "\033[32;1m[info] \033[00m"
   LogLevelWarn = "\033[33;1m[warn] \033[00m"
   LogLevelError = "\033[31;1m[fail] \033[00m"
)

// Geting formated string
func Fmt(s string, v...interface {}) string {
    return fmt.Sprintf(s, v...)
}

var logger *Logger = NewLogger()
type Logger struct {
    log.Logger
}

// Get SexPistol logger
func NewLogger() *Logger {
    l := new(Logger)
    l.Logger = *log.New(os.Stderr, "\r\n", log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
    return l
}

// Set SexPistol logger
func UseLogger(l *Logger) {
    logger = l
}

// Logging a raw log
func RawLog(typ string, useCaller bool, args...interface{}) {
    caller := ""
    if useCaller {
        _, file, line, _ := runtime.Caller(2)

        workingdir, _ := os.Getwd()
        if len(workingdir) < len(file) {
            file = file[len(workingdir)+1:]
        }

        caller = Fmt("%s:%d ", file, line)
    }

    fmt_args := []interface {}{caller, Fmt("%s", typ)}
    fmt_args = append(fmt_args, fmt.Sprint(args...))

    logger.Print(fmt_args...)
}

// Logging information logs with Sex.Logger()
func Log (args ...interface {}) {
    RawLog(LogLevelInfo, true, args...)
}

// Logging error logs with Sex.Logger()
func Err (args ...interface {}) {
    RawLog(LogLevelError, true, args...)
}

// Logging warning logs with Sex.Logger()
func War (args ...interface {}) {
    RawLog(LogLevelWarn, true, args...)
}

// Logging error logs with Sex.Logger() and killing the application
func Die (args ...interface {}) {
    RawLog(LogLevelError, true, args...)
    os.Exit(1)
}

// Logging information formated logs with Sex.Logger()
// Example:
//    Logf("%s %+v", "joao", []string{"joao", "maria"})
//    Logf("%.2f", 409.845)
func Logf (args ...interface {}) {
    RawLog(LogLevelInfo, true, Fmt(args[0].(string), args[1:]...))
}

// Logging error formated logs with Sex.Logger()
// Example:
//    Errf("%s %+v", "joao", []string{"joao", "maria"})
//    Errf("%.2f", 409.845)
func Errf (args ...interface {}) {
    RawLog(LogLevelError, true, Fmt(args[0].(string), args[1:]...))
}

// Logging warning formated logs with Sex.Logger()
// Example:
//    Warf("%s %+v", "joao", []string{"joao", "maria"})
//    Warf("%.2f", 409.845)
func Warf (args ...interface {}) {
    RawLog(LogLevelWarn, true, Fmt(args[0].(string), args[1:]...))
}

// Logging error logs with Sex.Logger() and killing the application
// Example:
//    Dief("%s %+v", "joao", []string{"joao", "maria"})
//    Dief("%.2f", 409.845)
func Dief (args ...interface {}) {
    RawLog(LogLevelError, true, Fmt(args[0].(string), args[1:]...))
    os.Exit(1)
}

// Debuging stdout display
func Debug (v...interface{}) {
    fmt.Println("\n--------------------------------")
    Log(v...)
    fmt.Print("\n")
}

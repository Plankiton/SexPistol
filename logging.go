package Sex

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

const (
	LogLevelInfo  = "\033[32;1m[info] \033[00m"
	LogLevelWarn  = "\033[33;1m[warn] \033[00m"
	LogLevelError = "\033[31;1m[fail] \033[00m"
)

// Fmt provides formated string
func Fmt(s string, v ...interface{}) string {
	return fmt.Sprintf(s, v...)
}

var logger *Logger = &Logger{Logger: *log.New(os.Stderr, "\r\n", log.LstdFlags|log.Lmicroseconds)}

type Logger struct {
	log.Logger
}

// NewLogger provides new SexPistol logger
func NewLogger() *Logger {
	l := new(Logger)
	l.Logger = *log.New(os.Stderr, "\r\n", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	return l
}

// UseLogger sets SexPistol logger
func UseLogger(l *Logger) {
	logger = l
}

// RawLog logs a raw log
func RawLog(typ string, useCaller bool, args ...interface{}) {
	caller := ""
	if useCaller {
		_, file, line, _ := runtime.Caller(2)

		workingdir, _ := os.Getwd()
		if len(workingdir) < len(file) {
			file = file[len(workingdir)+1:]
		}

		caller = Fmt("%s:%d ", file, line)
	}

	fmt_args := []interface{}{caller, Fmt("%s", typ)}
	fmt_args = append(fmt_args, fmt.Sprint(args...))

	logger.Print(fmt_args...)
}

// Log logs information logs with Sex.Logger()
func Log(args ...interface{}) {
	RawLog(LogLevelInfo, true, args...)
}

// Err logs error logs with Sex.Logger()
func Err(args ...interface{}) {
	RawLog(LogLevelError, true, args...)
}

// War logs warning logs with Sex.Logger()
func War(args ...interface{}) {
	RawLog(LogLevelWarn, true, args...)
}

// Die logs error logs with Sex.Logger() and kill application
func Die(args ...interface{}) {
	RawLog(LogLevelError, true, args...)
	os.Exit(1)
}

// Logf logs information formated logs with Sex.Logger()
// Example:
//    Logf("%s %+v", "joao", []string{"joao", "maria"})
//    Logf("%.2f", 409.845)
func Logf(args ...interface{}) {
	RawLog(LogLevelInfo, true, Fmt(args[0].(string), args[1:]...))
}

// Errf logs error formated with Sex.Logger()
// Example:
//    Errf("%s %+v", "joao", []string{"joao", "maria"})
//    Errf("%.2f", 409.845)
func Errf(args ...interface{}) {
	RawLog(LogLevelError, true, Fmt(args[0].(string), args[1:]...))
}

// Warf logs warning formated with Sex.Logger()
// Example:
//    Warf("%s %+v", "joao", []string{"joao", "maria"})
//    Warf("%.2f", 409.845)
func Warf(args ...interface{}) {
	RawLog(LogLevelWarn, true, Fmt(args[0].(string), args[1:]...))
}

// Dief logs error with Sex.Logger() and killing the application
// Example:
//    Dief("%s %+v", "joao", []string{"joao", "maria"})
//    Dief("%.2f", 409.845)
func Dief(args ...interface{}) {
	RawLog(LogLevelError, true, Fmt(args[0].(string), args[1:]...))
	os.Exit(1)
}

// Debuging stdout display
func Debug(v ...interface{}) {
	fmt.Println("----")
	logger.Println(v...)
	fmt.Print("----\n\n")
}

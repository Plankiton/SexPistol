package sex

import (
    "os"
    "log"
)

func Log (args ...interface {}) {
    fmt_args := []interface {}{"\033[32;1m[info] \033[00m"}
    fmt_args = append(fmt_args, args...)

    log.Println(fmt_args...)
}

func Err (args ...interface {}) {
    fmt_args := []interface {}{"\033[31;1m[erro] \033[00m"}
    fmt_args = append(fmt_args, args...)

    log.Println(fmt_args...)
}

func War (args ...interface {}) {
    fmt_args := []interface {}{"\033[33;1m[warn] \033[00m"}
    fmt_args = append(fmt_args, args...)

    log.Println(fmt_args...)
}

func Die (args ...interface {}) {
    fmt_args := []interface {}{"\033[31;1m[erro] \033[00m"}
    fmt_args = append(fmt_args, args...)

    log.Println(fmt_args...)
    os.Exit(1)
}

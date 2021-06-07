package Sex

import (
)

type CartridgeSkel interface {
    Add(ModelSkel) error
    Del(ModelSkel) error
    Sav(ModelSkel) error

    Select(...interface {}) CartridgeSkel
    Update(...interface {}) CartridgeSkel
    Insert(...interface{})  CartridgeSkel
    Delete(...interface{})  CartridgeSkel

    Where(...interface {})  CartridgeSkel
    Join(...interface{})    CartridgeSkel
    Like(...interface{})    CartridgeSkel
    In(...interface{})      CartridgeSkel
    OrderBy(...interface{}) CartridgeSkel
    GroupBy(...interface{}) CartridgeSkel

    Run()    CartridgeSkel
    Commit() CartridgeSkel
}

type Cartridge struct {
    Query    []interface {}
    Err      error

}

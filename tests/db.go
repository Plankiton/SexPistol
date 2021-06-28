package main
import (
    "github.com/Plankiton/SexPistol"
    "github.com/Plankiton/SexPistol/Cartridge"
    "github.com/Plankiton/SexPistol/Cartridge/Models"
)

type User struct {
    SexModel.User
}

func (u *User) New() error {
    u.Name += "_JOAOAOJOJODI"

    return nil
}

func main() {
    db, _ := SexDB.Open("j.db", SexDB.Sqlite)
    db.AddModels(&User{})


    u := User{}
    u.Name = "joao"

    Sex.SuperPut(u.Name)
    db.Create(&u)
    Sex.SuperPut(u.Name)
}

package main
import (
    "github.com/Plankiton/SexPistol"
    "github.com/Plankiton/SexPistol/Database"
    "github.com/Plankiton/SexPistol/Database/Models"
)

type User struct {
    SexModel.User
}

func (u *User) New() error {
    u.Name += "_JOAOAOJOJODI"

    return nil
}

func main() {
    db, _ := SexDB.Open("j.db", SexDatabase.Sqlite)
    db.AddModels(&User{})


    u := User{}
    u.Name = "joao"

    Sex.SuperPut(u.Name)
    db.Create(&u)
    Sex.SuperPut(u.Name)
}

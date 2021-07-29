package main

import "github.com/Plankiton/SexPistol"

func main() {
    j := map[string] string {
        "joao": "maria",
    }
    m := map[string] interface{} {
        "maria": true,
        "joao": false,
    }
    Sex.Log(m)

    Sex.Merge(j, &m)
    Sex.Log(m)

    Sex.Merge(j, &m, true)
    Sex.Log(m)
}

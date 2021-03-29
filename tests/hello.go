package main
import "github.com/Plankiton/SexPistol"

func main() {
    router := new (Sex.Pistol)
    router.
    Add(`/hello/{name}`, func (r Sex.Request) (string, int) {
        return Sex.Fmt("Hello %s", r.PathVars["name"]), 200
    }).
    Add("/api", func (r Sex.Request) (Sex.Json, int) {
        return Sex.Bullet {
            Message: "Joao eh gay",
        }, 200
    })

    Sex.Err(router.Run())
}

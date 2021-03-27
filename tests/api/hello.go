package main
import "github.com/Plankiton/SexPistol"

func main() {
    router := new (Sex.Pistol)
    router.
    Add("/hello/{name:.{0,}}", func (r Sex.Request) (string, int) {
        return Sex.Fmt("Hello %s", r.PathVars["name"]), 200
    }).
    Add("/api", func (r Sex.Request) (interface{}, int) {
        return Sex.Bullet {
            Message: "Joao eh gay",
        }, 200
    })

    Sex.Err(router.Run())
}

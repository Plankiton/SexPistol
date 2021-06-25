package main
import "github.com/Plankiton/SexPistol"

func main() {
    router := Sex.NewPistol()
    router.
    Add(`/hello/{name}`, func (r Sex.Request) (string, int) {
        name := r.PathVars["name"]
        Sex.Logf("Hello %s", name)
        return Sex.Fmt("Hello %s", name), 200
    }).
    Add("/api", func (r Sex.Request) (Sex.Json, int) {
        return Sex.Bullet {
            Message: "Joao eh gay",
        }, 200
    })

    Sex.Err(router.Run())
}

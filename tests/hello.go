package main
import (
    "net/http"
    "github.com/Plankiton/SexPistol"
)

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
    }).
    Add("/joao/logo", func (w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Joao eh gay"))
        w.WriteHeader(300)
    })

    Sex.Err(router.Run())
}

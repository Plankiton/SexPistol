package main
import "github.com/Plankiton/SexPistol"

func main() {
    router := new (Sex.Pistol)
    router.
    Add("/{name:.{0,}}", func (r Sex.Request) ([]byte, int) {
        return []byte("Hello "+r.PathVars["name"]+"!"), 200
    })

    Sex.Err(router.Run())
}

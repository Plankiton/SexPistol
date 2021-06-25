<a href="https://plankiton.github.io/SexPistol"><img src="assets/Icon.png" alt="Sex Pistol Icon" align="left" style="max-height: 700px;max-width: 30%"/></a>


# Sex Pistol

----

A Sex web micro framework for GoLang.

# Get Started

-----

## Install

It is easy, just take that command on your favorite terminal:

```shell
$ go get github.com/Plankiton/SexPistol
```

## First Code

Its so easy like that:

```go
package main
import "github.com/Plankiton/SexPistol"
func main() {
    Sex.NewPistol().
    Add("/{name}", func (r Sex.Request) string {
        return Sex.Fmt("Hello, %s", r.PathVars["name"])
    }).
    Run()
}
```

Too sex no? That code make the same thing what the code above:

>  Code using default library from Go

```go
package main
import (
    "net/http"
    "fmt"
    "strings"
)

func main() {
    http.HandleFunc("/",func Hello (w http.ResponseWriter, r *http.Request) {
        path := strings.Split(r.URL.Path, "/")

        w.WriteHeader(200)
        w.Write([]byte(fmt.Sprintf(
            "Hello, %s", path[len(path)-1],
        )))
    })

    http.ListenAndServe(":8000", nil)
}
```

# [Md Documentation](./docs/readme.md) | [Html Documentation](https://plankiton.github.io/SexPistol)

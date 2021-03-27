

# Sex Pistol

----

A Sex web micro framework for GoLang.

<center><img src="assets/Icon.png" alt="Sex Pistol Icon" align="left" style="max-height: 700px;max-width: 50%"/>

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
    router := new(Sex.Pistol)
    router.Add("/{name}", func (r Sex.Request) (string, int) {
        return Sex.Fmt("Hello, %s", r.PathVars["name"]), 200
    })
    router.Run()
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

# Documentation

-----



## Responses

Lets play with the Sex Function types:

```go
package main
import "github.com/Plankiton/SexPistol"
func main() {
    router := new(Sex.Pistol)
    router.Add("/{name}", Hello)
    router.Run()
}
```

String function way

```go
func Hello(r Sex.Request) (string, int) {
    return Sex.Fmt(`{
    	"Hello": "%s"
    }`, r.PathVars["name"]), 200
}
```

Interface way

> That type take lists, maps and structs into a json output

```go
func (r Sex.Resquest) (interface {}, int) {
    return map[string]string {
        "Hello": r.PathVars["name"],
    }, 200
}
```

And you can too make a custom response:

```go
func (r Sex.Resquest) (Sex.Response, int) {
    res := r.MkResponse()
    res.SetBody(Sex.Jsonify(map[string]string {
        "Hello": r.PathVars["name"],
    }))
    
    return res, 200
}

// When you are using the custom response the status_code are opcional

func (r Sex.Resquest) (Sex.Response) {
    res := r.MkResponse()
    res.SetBody(Sex.Jsonify(map[string]string {
        "Hello": r.PathVars["name"],
    }))
    
    res.SetStatus(200) // Opcional, the default status code are "200"
    res.SetCookie("name", r.PathVars, 1000) // Setting cookies (key, value, expires)
    res.Header().Set("Content-Type", "application/json") // Setting headers (key, value)
    return res
}
```

We have a template for response in case of Rest APIs, it optional but are a way to make a beautiful API:

```go
return Sex.Bullet {
    Type: "Error|Sucess",
    Message: "Response description",
    Data: "Response data"
}, 200
```


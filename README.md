# Sex Pistol

----

A Sex web micro framework for GoLang.

<center><img src="assets/sexpistol-hd.svg" alt="Sex Pistol Icon" align="center" style="zoom: 40%; max-height: 900px;" /><img src="assets/sexpistol-white-theme.svg" alt="Sex Pistol Icon" align="center" style="zoom: 40%; max-height: 900px;-moz-transform: scaleX(-1);-o-transform: scaleX(-1);-webkit-transform: scaleX(-1);transform: scaleX(-1);" /></center>

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

How you did look, that is a minimal Sex web application posssible:

## Responses

```go
package main
import "github.com/Plankiton/SexPistol"
func main() {
    router := new(Sex.Pistol)
    router.Add("/{name}", Hello)
    router.Run()
}

func Hello(r Sex.Request) (string, int) {
    return Sex.Fmt(`{
    	"Hello": "%s"
    }`, r.PathVars["name"]), 200
}
```

But, the Pistol to support various function types, between it have:

```go
func (r Sex.Resquest) (interface {}, int) {
    return map[string]string {
        "Hello": r.PathVars["name"],
    }, 200
}
```

> The output of function above is automatically parsed to json format.

And you can too make a custom response:

```go
func (r Sex.Resquest) (Sex.Response, int) {
    res := r.MkResponse()
    res.SetBody(Sex.Jsonify(map[string]string {
        "Hello": r.PathVars["name"],
    }))
    
    res.Header().Set("Content-Type", "application/json")
    return res, 200
}

// When you are using the custom response the status_code are opcional

func (r Sex.Resquest) (Sex.Response) {
    res := r.MkResponse()
    res.SetBody(Sex.Jsonify(map[string]string {
        "Hello": r.PathVars["name"],
    }))
    
    res.SetStatus(200) // Opcional, the default status code are "200"
    
    res.Header().Set("Content-Type", "application/json")
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




# Sex Pistol

----

A Sex web micro framework for GoLang.

<img src="assets/Icon.png" alt="Sex Pistol Icon" align="left" style="max-height: 700px;max-width: 30%"/>

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
        return Sex.Fmt("Hello, %s", r.PathVars["name"]), Sex.StatusOk
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
    }`, r.PathVars["name"]), Sex.StatusOk
}
```

Interface way

> That type take lists, maps and structs into a json output

```go
func (r Sex.Request) (interface {}, int) {
    return map[string]string {
        "Hello": r.PathVars["name"],
    }, Sex.StatusOk
}
```

And you can too make a custom response:

```go
func (r Sex.Request) (Sex.Response, int) {
    res := r.MkResponse()
    res.SetBody(Sex.Jsonify(map[string]string {
        "Hello": r.PathVars["name"],
    }))
    
    return res, Sex.StatusOk
}

// When you are using the custom response the status_code are opcional

func (r Sex.Request) (Sex.Response) {
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
Sex.Bullet {
    Type: "Error|Sucess",
    Message: "Response description",
    Data: "Response data"
}
```

## Requests

Well, how you did look, the request is handle on function by arguments and it have very util properties:

```go
func (r Sex.Request) Sex.Response {
    res := r.MkResponse()          // Create Sex.Response struct
    r.Header.Get("Authorization")  // Get a header
    r.Cookies.Get("name")          // Get a cookie
    r.PathVars["name"]             // Get path variable 
    
    r.ParseForm()
    r.Form.Get("name")             // Get url encoded form field
    
    var data map[string]interface{}
    r.JsonBody(&data)             // Parse json body to maps, structs or lists
    raw_body := r.RawBody()       // Get byte list with request body
    r.Body                        // io.ReadCloser Body for manual handle 
}
```

**`Sex.Request` have all `*http.Request` properties and functions, because of that you can use any tutorial for to handle `Sex.Request`**

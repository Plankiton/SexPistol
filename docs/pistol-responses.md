# Pistol | Responses

Pistol is the http server library, basicaly we can make Rest APIs, or Site server in simple way:

```go
package main
import "github.com/Plankiton/SexPistol"
func main() {
    new(Sex.Pistol).
    Add("/{name}", func (r Sex.Request) string {
        return Sex.Fmt("Hello, %s", r.PathVars["name"])
    }).
    Run()
}
```

That code make a server for a custom "hello world", basicaly we are taking an path variable and puting than on response, that is a string, but you can return another types on response:

### Bytes/String function way

```go
func (r Sex.Request) (string, int) {
    return []byte(
        Sex.Fmt(`{
    		"Hello": "%s"
    	}`, r.PathVars["name"])
    ), Sex.StatusOk
}

func (r Sex.Request) (string, int) {
    return Sex.Fmt(`{
    	"Hello": "%s"
    }`, r.PathVars["name"]), Sex.StatusOk
}
```

> Status code are opitional, if you want to return "200" (or StatusOk), is just hide the int on declaration of functions and second return value



### Json way

> That type take lists, maps and structs into a json output

```go
func (r Sex.Request) (Sex.Json, int) {
    return map[string]string {
        "Hello": r.PathVars["name"],
    }, Sex.StatusOk
}

// Too works with structs, arrays and wath ever
type Res struct {
    Hello string `json:"Hello"`
} 

func (r Sex.Request) (Sex.Json, int) {
    res := Res {
        Hello: r.PathVars["name"],
    }
    return res, Sex.StatusOk
}
```

### Custom way

And you can too make a custom response:

```go
func (r Sex.Request) (Sex.Response, int) {
    res := r.MkResponse()
    res.SetBody(Sex.Jsonify(map[string]string {
        "Hello": r.PathVars["name"],
    }))
    
    return res, Sex.StatusOk
}
```

### Bullet for pistol

We have a template for response in case of Rest APIs, it optional but are a way to make beautiful API returns:

> That template is for use on Json return way

```go
Sex.Bullet {
    Type: "Error|Success",
    Message: "Response description",
    Data: "Response data"
}
```

# [Pistol Requests](./pistol-requests.html)

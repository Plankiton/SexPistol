# Sex Pistol - Web Micro FrameWork for Golang



![Sex Pistol Logo](Icon.png)

A micro framework for golang, made for who hates the golang way of create web applications, and transforming this:

```go
package main
import ("net/http";"fmt";"strings")
func Hello (w http.ResponseWriter, r *http.Request) {

    path := strings.Split(r.URL.Path, "/")
    
    w.WriteHeader(200)
    w.Write([]byte(fmt.Sprint(
        "Hello, ", path[len(path)-1],
    )))
}

func main() {
    http.HandleFunc("/", Hello)
    http.ListenAndServe(":8000", nil)
}
```

On this:

```go
package main
import ("github.com/Plankiton/SexPistol";"fmt")
func Hello (r sex.Request) ([]byte, int) {
    return []byte(fmt.Sprintf("Hello, %s", r.PathVars["name"])), 200
}

func main() {
    router := new(sex.Pistol)
    router.Add("get", "/{name}", nil, Hello)
    router.Run("/", 8000)
}
```

So sex no?

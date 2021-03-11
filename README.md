# Sex Pistol - Web Micro FrameWork for Golang



<img src="Icon.png" align="center" alt="Sex Pistol Logo" style="zoom:80%;" />



A Sex method of to make web applications with GoLang.

# Get Started

You can to install the Sex Pistol framework with a simple `go get github.com/Plankiton/SexPistol` or adding `github.com/Plankiton/SexPistol` to the `go.mod` of your project.



## First Code



The principal objective of Sex Pistol is change the default method of make Rest APIs into a new and better method:

>  Code using default library from Go

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

>  Sex Pistol code

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


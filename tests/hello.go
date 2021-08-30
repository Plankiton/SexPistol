package main

import (
	"net/http"

	Sex "github.com/Plankiton/SexPistol"
)

func main() {
	router := Sex.NewPistol()
	router.
		Add(`/hello/{name}`, func(r Sex.Request) (string, int) {
			name := r.PathVars["name"]
			Sex.Logf("Hello %s", name)
			return Sex.Fmt("Hello %s", name), 200
		}).
		Add(`/hello/{name}/joao`, func(r Sex.Request) (string, int) {
			name := r.PathVars["name"]
			Sex.Logf("Hello %s", name)
			return Sex.Fmt("Hello %s", name), 200
		}, "get").
		Add("/api", func(r Sex.Request) (Sex.Json, int) {
			return Sex.Bullet{
				Message: "Joao eh gay",
			}, 200
		}).
		Add("/joao/logo", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Joao eh gay"))
			w.WriteHeader(300)
		})

	plugin := func(h http.Handler) http.Handler {
		handler := http.NewServeMux()
		handler.Handle("/", handler)
		handler.HandleFunc("/plugin/", func(http.ResponseWriter, *http.Request) {})
		return handler
	}

	Sex.Err(router.Run(plugin))
}

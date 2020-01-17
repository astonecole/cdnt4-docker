package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/nanoninja/bulma"
)

var addr = flag.String("addr", ":3000", "Adresse du serveur")

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, GoWorld</h1>")
}

type Message struct {
	Text string "json:text"
}

func api(w http.ResponseWriter, r *http.Request) {
	m := Message{}
	json.NewDecoder(r.Body).Decode(&m)
	fmt.Println(m.Text)
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()

	m := http.NewServeMux()
	m.HandleFunc("/", home)
	m.HandleFunc("/api", api)

	handler := bulma.BasicAuth("Connection", logger(m), bulma.User{
		"toto": "0000",
		"joe":  "1234",
	})

	http.ListenAndServe(*addr, handler)
}

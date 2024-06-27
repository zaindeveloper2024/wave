package main

import (
	"fmt"
	"net/http"

	"github.com/zaindeveloper2024/wave/wave"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ping %s\n", r.URL.Query().Get("name"))
}

func main() {
	w := wave.New()
	w.Handle("/ping", handler)
	w.Start(":8080")
}

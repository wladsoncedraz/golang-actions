package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Deploy with Actions GitHub")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

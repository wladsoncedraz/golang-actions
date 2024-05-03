package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Here we Go!")
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

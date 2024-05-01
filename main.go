package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Wellcome to my Website! I'm Wladson Cedraz!")
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

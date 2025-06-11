// main.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go on Azure Free Tier!")
	})
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "✅ Go Web App is running on Azure App Service!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :80...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

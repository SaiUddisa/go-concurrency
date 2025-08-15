package main

import (
	"fmt"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	color := "red"
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	query := r.URL.Query()
	color = query.Get("color")
	fmt.Fprintf(w, "<h1 style='color:%s'>Hello, World!</h1>", color)
}
func main() {
	http.HandleFunc("/", handleHello)

	fmt.Println("Server running at port 8080")
	http.ListenAndServe(":8080", nil)
}

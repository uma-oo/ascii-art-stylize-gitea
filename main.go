package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"asciiWeb/handler"
)

func main() {
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", intercept(fileServer)))
	http.HandleFunc("/", handler.HandleMainPage)
	http.HandleFunc("/ascii-art", handler.HandleAsciiArt)
	fmt.Println("Server starting at http://localhost:6500")
	log.Fatal(http.ListenAndServe(":6500", nil))
}

func intercept(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

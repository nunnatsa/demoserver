package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dist/index.html")
}

func favicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "image/x-icon")
	http.ServeFile(w, r, "dist/favicon.ico")
}

func logMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("got request; method: %s; path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func main() {
	http.Handle("/hello", logMiddleware(hello))
	http.Handle("/favicon.ico", logMiddleware(favicon))
	http.Handle("/", logMiddleware(index))


	log.Println("starting the demo-server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

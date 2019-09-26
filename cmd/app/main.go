package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi"
)

func init() {
	mime.AddExtensionType("js", "text/javascript")
	mime.AddExtensionType("css", "text/css")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	apiServer := "http://localhost:3000"
	url, err := url.Parse(apiServer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("url", url)
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
}

func main() {
	fs := http.FileServer(http.Dir("dist"))

	r := chi.NewRouter()
	r.Get("/api/*", apiHandler)
	r.Post("/api/*", apiHandler)
	r.Put("/api/*", apiHandler)
	r.Delete("/api/*", apiHandler)

	r.Get("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
	log.Fatal(http.ListenAndServe(":8080", r))
}

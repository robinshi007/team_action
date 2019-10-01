package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/rakyll/statik/fs"

	_ "team_action/statik"
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
	fmt.Println("proxy to => ", url)
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
}

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	//fs := http.FileServer(http.Dir("dist"))
	fss := http.FileServer(statikFS)

	r := chi.NewRouter()
	r.Get("/api/*", apiHandler)
	r.Post("/api/*", apiHandler)
	r.Put("/api/*", apiHandler)
	r.Delete("/api/*", apiHandler)

	r.Get("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fss.ServeHTTP(w, r)
	}))
	fmt.Println("listen on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

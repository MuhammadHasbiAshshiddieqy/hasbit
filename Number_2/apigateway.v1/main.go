package main

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

func main() {
    user, err := url.Parse("http://127.0.0.1:8081")
    if err != nil {
        log.Fatal(err)
    }
	external, err := url.Parse("http://127.0.0.1:8082")
    if err != nil {
        log.Fatal(err)
    }
    http.Handle("/user/", http.StripPrefix("/user/", httputil.NewSingleHostReverseProxy(user)))
	http.Handle("/external/", http.StripPrefix("/external/", httputil.NewSingleHostReverseProxy(external)))
	log.Println("Server running on port :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

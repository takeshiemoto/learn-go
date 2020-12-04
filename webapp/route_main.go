package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello WebApp, %s", r.URL.Path[1:])
}

func err(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Error!, %s", r.URL.Path[1:])
}

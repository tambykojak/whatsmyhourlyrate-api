package main

import (
	"fmt"
	"net/http"
)

func (a *app) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is where magic will happen %s", r.URL.Path[0:])
}

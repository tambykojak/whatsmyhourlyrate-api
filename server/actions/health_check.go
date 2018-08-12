package actions

import (
	"fmt"
	"net/http"
)

// HandleHealthCheck TODO
func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is where magic will happen %s", r.URL.Path[0:])
}

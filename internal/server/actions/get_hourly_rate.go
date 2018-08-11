package actions

import (
	"fmt"
	"net/http"
)

// HandleGetHourlyRate TODO
func HandleGetHourlyRate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is where magic will happen %s", r.URL.Path[0:])
}

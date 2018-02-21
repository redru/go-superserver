package apis

import (
	"net/http"
)

// Default
type Default struct {
}

// PersonsGet is
func PersonsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

package web

import (
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Theme home page"))
}
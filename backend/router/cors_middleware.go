package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CorsMiddlware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r, ps)
	}
}

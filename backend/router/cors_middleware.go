package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var AllowedOrigins map[string]bool = map[string]bool{
	"http://localhost:5173":                       true,
	"http://localhost:5174":                       true,
	"http://localhost:3000":                       true,
	"https://cors.compressibleflowcalculator.com": true,
}

func CorsMiddlware(next httprouter.Handle) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		if r.Method != "OPTIONS" {

			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization,Cookie")
			w.Header().Set("Access-Control-Max-Age", "86400")
			origin := r.Header.Get("Origin")

			if AllowedOrigins[origin] {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			//w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

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

package http

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func logger(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		start := time.Now()
		n(w, r, p)
		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			//r,
			time.Since(start),
		)
	}
}

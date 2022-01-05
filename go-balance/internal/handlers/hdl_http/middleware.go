package hdl_http

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/go-hexa/go-balance/pkg"

)

func MiddleWareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("-------------------------------------------- \n")
		log.Println("MiddleWareHandler (INICIO)")

		log.Println(r.Header.Get("jwt"))

		if len(r.Header.Get("jwt")) == 0 {
			pkg.HandlerHttpError(w, pkg.ErrUnauthorized)
			json.NewEncoder(w).Encode(pkg.ErrUnauthorized)
			return
		}
	
		log.Println("MiddleWareHandler (FIM)")
		log.Printf("-------------------------------------------- \n")
		
		next.ServeHTTP(w, r)
	})
}
package routers

import (
	s "day15tblSelling/controllers"

	"github.com/gorilla/mux"
)

func setSellingRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/selling", s.GetSelling).Methods("GET")
	return router
}

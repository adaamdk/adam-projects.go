package routers

import (
	s "day15sellOfficer/controllers"

	"github.com/gorilla/mux"
)

func setOfficerRouter(router *mux.Router) *mux.Router {

	router.HandleFunc("/officer", s.GetOfficer).Methods("GET")
	return router
}

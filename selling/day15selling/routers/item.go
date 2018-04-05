package routers

import (
	s "day15selling/controllers"

	"github.com/gorilla/mux"
)

func setItemRouter(router *mux.Router) *mux.Router {
	// 1b. Buat func pd controller
	router.HandleFunc("/item", s.GetItem).Methods("GET")
	return router
}

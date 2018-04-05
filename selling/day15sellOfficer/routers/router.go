package routers

import (
	"github.com/gorilla/mux"
)

func InitRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	//Set route from table
	router = setOfficerRouter(router)
	return router
}

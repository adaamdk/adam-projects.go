package routers

import (
	"github.com/gorilla/mux"
)

func InitRouters() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	//Atur routing dari tabel
	router = setItemRouter(router)
	return router
}

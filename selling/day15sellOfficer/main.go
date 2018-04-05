package main

import (
	"day15sellOfficer/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouters()
	log.Fatal(http.ListenAndServe(":8871", router))
}

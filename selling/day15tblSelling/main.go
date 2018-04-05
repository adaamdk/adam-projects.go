package main

import (
	"day15tblSelling/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouters()
	log.Fatal(http.ListenAndServe(":8872", router))
}

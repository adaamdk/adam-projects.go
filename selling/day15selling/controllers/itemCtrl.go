package controllers

import (
	"day15selling/data"
	"encoding/json"
	"net/http"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	context := Context{}

	c := DBinit(context.DB)
	defer c.Close()

	rpstori := &data.ItemRepo{c}
	item := data.GetAll(rpstori)

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(item)
	w.Write(mdl)
}

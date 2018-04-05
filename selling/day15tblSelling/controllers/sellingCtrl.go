package controllers

import (
	"day15tblSelling/data"
	"encoding/json"
	"net/http"
)

func GetSelling(w http.ResponseWriter, r *http.Request) {
	context := Context{}

	ct := DBinit(context.DB)
	defer ct.Close()

	repos := &data.SellingRepo{ct}
	selling := data.GetAll(repos)

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(selling)
	w.Write(mdl)
}

package controllers

import (
	"day15sellOfficer/data"
	"encoding/json"
	"net/http"
)

func GetOfficer(w http.ResponseWriter, r *http.Request) {
	context := Context{}

	ct := DBinit(context.DB)
	defer ct.Close()

	repos := &data.OfficerRepo{ct}
	officer := data.GetAll(repos)

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(officer)

	w.Write(mdl)
}

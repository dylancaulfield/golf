package handlers

import (
	"encoding/json"
	"github.com/dyldawg/golf/models"
	"github.com/gorilla/mux"
	"net/http"
)

func ScoresHandler(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	scores, err := models.GetScoresForResult(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	js, err := json.Marshal(scores)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)

}
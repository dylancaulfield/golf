package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dyldawg/golf/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func PlayersHandler(w http.ResponseWriter, r *http.Request){

	players, err := models.GetPlayers()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	js, err := json.Marshal(players)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)

}

func NewPlayerHandler(w http.ResponseWriter, r *http.Request){

	var p models.Player

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		fmt.Println(err)

		return
	}

	err = json.Unmarshal(body, &p)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		fmt.Println(err)

		return
	}

	//Get from JSON body
	err = models.NewPlayer(p)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(http.StatusText(http.StatusOK)))

}

func PlayerResultsHandler(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	results, err := models.GetPlayerResults(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}



	js, err := json.Marshal(results)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)

}
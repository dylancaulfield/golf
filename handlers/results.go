package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dyldawg/golf/models"
	"io/ioutil"
	"net/http"
)


func ResultsHandler(w http.ResponseWriter, r *http.Request){

	results, err := models.GetResults()
	if err != nil {

		fmt.Println(err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	js, err := json.Marshal(results)
	if err != nil {

		fmt.Println(err)

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)

}

func NewResultHandler(w http.ResponseWriter, r *http.Request) {


	var newResult models.JsonResult

	//Read Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	//Parse JSON
	err = json.Unmarshal(body, &newResult)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		fmt.Println(err)

		return
	}

	//Create Result
	err = models.CreateResult(newResult)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		fmt.Println(err)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(http.StatusText(http.StatusOK)))

}

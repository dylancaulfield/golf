package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dyldawg/golf/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func NewCourseHandler(w http.ResponseWriter, r *http.Request) {

	var c models.Course

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		fmt.Println(err)

		return
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		fmt.Println(err)

		return
	}

	if c.Par < 54 || c.Par > 100 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	//Get from JSON body
	err = models.NewCourse(c)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(http.StatusText(http.StatusOK)))

}

func CoursesHandler(w http.ResponseWriter, r *http.Request) {

	courses, err := models.GetCourses()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	js, err := json.Marshal(courses)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)

}

func GetCourseHandler(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	course, err := models.GetCourse(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	js, err := json.Marshal(course)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)

}

func UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {

	var c models.Course

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	if c.Par < 54 || c.Par > 100 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	// get from json body
	err = models.UpdateCourse(c)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		fmt.Println(err)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(http.StatusText(http.StatusOK)))

}

func DeleteCourseHandler(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	err := models.DeleteCourse(vars["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		fmt.Println(err)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(http.StatusText(http.StatusOK)))

}

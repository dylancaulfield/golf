package models

import (
	"github.com/google/uuid"
	"log"
)

type JsonResult struct {
	Course string  `json:"course"`
	Date   string  `json:"date"`
	Scores []Score `json:"scores"`
}

type Result struct {
	Id     string `json:"id"`
	Course string `json:"course"`
	Date   string `json:"date"`
}

func CreateResult(result JsonResult) error {

	tx, err := getDatabase().Begin()
	if err != nil {
		return err
	}

	uid, _ := uuid.NewUUID()

	_, err = tx.Query("INSERT INTO results VALUES (?, ?, ?)", uid.String(), result.Course, result.Date)
	if err != nil {
		handleError(tx.Rollback())
		return err
	}

	prepared, err := tx.Prepare("INSERT INTO scores VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	for _, e := range result.Scores {

		_, err := prepared.Exec(e.Player, e.Score, uid.String())
		if err != nil {
			handleError(tx.Rollback())
			return err
		}

	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil

}

func GetResults() ([]Result, error) {

	dbResults, err := getDatabase().Query("SELECT * FROM results")
	if err != nil {
		return []Result{}, err
	}

	var results []Result

	for dbResults.Next() {
		var r Result

		err = dbResults.Scan(&r.Id, &r.Course, &r.Date)
		if err != nil {
			return []Result{}, err
		}

		results = append(results, r)
	}

	return results, nil

}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

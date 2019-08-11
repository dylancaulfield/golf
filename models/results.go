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

type Score struct {
	Player string `json:"player"`
	Score int `json:"score"`
	Result string `json:"result"`
}

type Result struct {
	Id     string `json:"id"`
	Course string `json:"course"`
	Date   string `json:"date"`
}

type resultWithAllInfo struct {
	Id     string  `json:"id"`
	Date   string  `json:"date"`
	Course string  `json:"course"`
	Par    int     `json:"par"`
	Scores []playerScore `json:"scores"`
}

type playerScore struct {
	PlayerId string `json:"player_id"`
	PlayerName string `json:"player_name"`
	Score int `json:"score"`
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

	dbResults, err := getDatabase().Query("SELECT results.id, courses.name, results.date FROM results, courses WHERE courses.id=results.course;")
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

func GetResult(id string) (resultWithAllInfo, error) {

	var r resultWithAllInfo
	var scores []playerScore


	// Scores and Their Players
	results, err := getDatabase().Query("SELECT DISTINCT players.id, players.name, scores.score FROM players, scores, results WHERE scores.player=players.id AND scores.result=?;", id)
	if err != nil {
		return resultWithAllInfo{}, err
	}

	for results.Next() {
		var s playerScore

		err = results.Scan(&s.PlayerId, &s.PlayerName, &s.Score)
		if err != nil {
			return resultWithAllInfo{}, err
		}

		scores = append(scores, s)
	}

	r.Scores = scores

	// Result and Course
	err = getDatabase().QueryRow("SELECT results.id, results.date, courses.name AS course, courses.par FROM results, courses WHERE results.course=courses.id AND results.id=?;", id).Scan(&r.Id, &r.Date, &r.Course, &r.Par)
	if err != nil {
		return resultWithAllInfo{}, err
	}

	return r, nil

}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

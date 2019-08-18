package models

import (
	"github.com/google/uuid"
)

type Player struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PlayerResult struct {
	Date   string `json:"date"`
	Score  int    `json:"score"`
	Course string `json:"course"`
	Par    int    `json:"par"`
}

type PlayerData struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Results []playerResult `json:"results"`
}

type playerResult struct {
	ResultId string `json:"result_id"`
	Date string `json:"date"`
	CourseId string `json:"course_id"`
	CourseName string `json:"course_name"`
	Par int `json:"par"`
	Score int `json:"score"`
}

func NewPlayer(player Player) error {

	uid, _ := uuid.NewUUID()

	_, err := getDatabase().Query("INSERT INTO players VALUES (?, ?)", uid.String(), player.Name)
	if err != nil {
		return err
	}

	return nil

}

func GetPlayers() ([]Player, error) {

	results, err := getDatabase().Query("SELECT * FROM players")
	if err != nil {
		return []Player{}, err
	}

	var players []Player

	for results.Next() {
		var player Player

		err = results.Scan(&player.Id, &player.Name)
		if err != nil {
			return []Player{}, err
		}

		players = append(players, player)
	}

	return players, nil

}

func GetPlayer(id string) (PlayerData, error) {

	var p PlayerData

	// Get Player
	err := getDatabase().QueryRow("SELECT players.id, players.name FROM players WHERE players.id=?", id).Scan(&p.Id, &p.Name)
	if err != nil {
		return PlayerData{}, err
	}

	// Get PlayerResults
	results, err := getDatabase().Query("SELECT results.id AS result_id, results.date, courses.id AS course_id, courses.name AS course_name, courses.par, scores.score FROM courses, players, results, scores WHERE scores.player=players.id AND scores.result=results.id AND results.course=courses.id AND players.id=?;", id)
	if err != nil {
		return PlayerData{}, err
	}

	for results.Next() {
		var pR playerResult

		err = results.Scan(&pR.ResultId, &pR.Date, &pR.CourseId, &pR.CourseName, &pR.Par, &pR.Score)
		if err != nil {
			return PlayerData{}, err
		}

		p.Results = append(p.Results, pR)
	}

	return p, nil

}











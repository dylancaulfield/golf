package models

import (
	"fmt"
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

func GetPlayer(id string) (Player, error) {

	var p Player

	result := getDatabase().QueryRow("SELECT players.id, players.name FROM players WHERE players.id=?", id)

	err := result.Scan(&p.Id, &p.Name)
	if err != nil {
		return Player{}, err
	}

	return p, nil

}

func GetPlayerResults(id string) ([]PlayerResult, error) {

	results, err := getDatabase().Query("SELECT results.date, scores.score, courses.name, courses.par FROM players, scores, courses, results WHERE scores.player=players.id AND scores.result=results.id AND results.course=courses.id AND players.id=?;", id)
	if err != nil {

		fmt.Println(err)

		return []PlayerResult{}, err
	}

	var playersResults []PlayerResult

	for results.Next() {
		var playerResult PlayerResult

		err = results.Scan(&playerResult.Date, &playerResult.Score, &playerResult.Course, &playerResult.Par)
		if err != nil {
			return []PlayerResult{}, err
		}

		playersResults = append(playersResults, playerResult)

	}

	return playersResults, nil

}

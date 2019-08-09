package models

import "github.com/google/uuid"

type Player struct {
	Id   string `json:"id"`
	Name string `json:"name"`
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

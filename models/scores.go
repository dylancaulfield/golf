package models

type Score struct {
	Player string `json:"player"`
	Score  int    `json:"score"`
	Result string `json:"result"`
}

type ScoreNoResult struct {
	Player string `json:"player"`
	Score  int    `json:"score"`
}

func GetScoresForResult(id string) ([]ScoreNoResult, error) {

	results, err := getDatabase().Query("SELECT player, score FROM scores WHERE result = ?", id)
	if err != nil {
		return []ScoreNoResult{}, err
	}

	var scores []ScoreNoResult

	for results.Next() {
		var score ScoreNoResult

		err = results.Scan(&score.Player, &score.Score)
		if err != nil {
			return []ScoreNoResult{}, err
		}

		scores = append(scores, score)
	}

	return scores, nil

}




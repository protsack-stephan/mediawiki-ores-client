package ores

type response struct {
	Models map[Model]ModelInfo `json:"models"`
	Scores map[int]map[Model]struct {
		Error struct {
			Type    string `json:"type"`
			Message string `json:"message"`
		} `json:"error"`
		Score map[string]interface{} `json:"score"`
	} `json:"scores"`
}

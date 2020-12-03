package ores

type response struct {
	Models map[Model]struct {
		Version string `json:"version"`
	} `json:"models"`
	Scores map[int]map[Model]struct {
		Score map[string]interface{} `json:"score"`
	} `json:"scores"`
}

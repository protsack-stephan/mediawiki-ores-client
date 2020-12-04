package ores

type responseError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type response struct {
	Models map[Model]ModelInfo `json:"models"`
	Scores map[int]map[Model]struct {
		Error responseError          `json:"error"`
		Score map[string]interface{} `json:"score"`
	} `json:"scores"`
}

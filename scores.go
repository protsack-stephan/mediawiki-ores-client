package ores

// ScoreError response error messages (missing revision for example)
type ScoreError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// ScoreArticleQualityProbability ores the scores values for article quality scores
type ScoreArticleQualityProbability struct {
	B     float64 `json:"B"`
	C     float64 `json:"C"`
	FA    float64 `json:"FA"`
	GA    float64 `json:"GA"`
	Start float64 `json:"Start"`
	Stub  float64 `json:"Stub"`
}

// ScoreArticleQuality quality of the article model score
type ScoreArticleQuality struct {
	Prediction  string                         `json:"prediction"`
	Probability ScoreArticleQualityProbability `json:"probability"`
}

// ScoreDamagingProbability stores the scores values for damaging scores
type ScoreDamagingProbability struct {
	False float64 `json:"false"`
	True  float64 `json:"true"`
}

// ScoreDamaging score to determine if revision is damaging
type ScoreDamaging struct {
	Prediction  bool                     `json:"prediction"`
	Probability ScoreDamagingProbability `json:"probability"`
}

// ScoreGoodFaithProbability stores the scores values for goodfaith scores
type ScoreGoodFaithProbability struct {
	False float64 `json:"false"`
	True  float64 `json:"true"`
}

// ScoreGoodFaith revision good faith score
type ScoreGoodFaith struct {
	Prediction  bool                      `json:"prediction"`
	Probability ScoreGoodFaithProbability `json:"probability"`
}

// Scores ORES API response type
type Scores struct {
	Models map[Model]ModelInfo `json:"models"`
	Scores map[int]struct {
		Articlequality *struct {
			Score *ScoreArticleQuality `json:"score"`
			Error *ScoreError          `json:"error"`
		} `json:"articlequality"`
		Damaging *struct {
			Score *ScoreDamaging `json:"score"`
			Error *ScoreError    `json:"error"`
		} `json:"damaging"`
		Goodfaith struct {
			Score *ScoreGoodFaith `json:"score"`
			Error *ScoreError     `json:"error"`
		} `json:"goodfaith"`
	} `json:"scores"`
}

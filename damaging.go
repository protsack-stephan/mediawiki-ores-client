package ores

// Damaging struct for model
type Damaging struct {
	Prediction  bool `json:"prediction"`
	Probability struct {
		False float64 `json:"false"`
		True  float64 `json:"true"`
	} `json:"probability"`
}

func (dmg *Damaging) fromMap(score map[string]interface{}) error {
	switch score["prediction"].(type) {
	case bool:
		dmg.Prediction = score["prediction"].(bool)
	default:
		return ErrInvalidDataInterface
	}

	var probability map[string]interface{}

	switch score["probability"].(type) {
	case map[string]interface{}:
		probability = score["probability"].(map[string]interface{})
	default:
		return ErrInvalidDataInterface
	}

	switch probability["true"].(type) {
	case float64:
		dmg.Probability.True = probability["true"].(float64)
	default:
		return ErrInvalidDataInterface
	}

	switch probability["false"].(type) {
	case float64:
		dmg.Probability.False = probability["false"].(float64)
	default:
		return ErrInvalidDataInterface
	}

	return nil
}

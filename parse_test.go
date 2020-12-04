package ores

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var parseTestData = []struct {
	Rev              int
	Prediction       bool
	ProbabilityFalse float64
	ProbabilityTrue  float64
}{
	{
		1,
		false,
		0.7,
		0.3,
	},
	{
		2,
		true,
		0.3,
		0.6,
	},
}

const parseTestDBName = "testwiki"
const parseTestModel = ModelDamaging
const parseTestTemplate = `{"%s":{"models":{"damaging":{"version":"0.5.1"}},"scores":{"%d":{"damaging":{"score":{"prediction":%t,"probability":{"false":%f,"true":%f}}}},"%d":{"damaging":{"score":{"prediction":%t,"probability":{"false":%f,"true":%f}}}}}}}`
const parseTestErrorRev = 10
const parseTestErrorTemplate = `{"%s":{"models":{"damaging":{"version":"0.5.1"}},"scores":{"%d":{"damaging":{"error":{"message":"RevisionNotFound:Couldnotfindrevision({revision}:%d)","type":"RevisionNotFound"}}}}}}`

func TestParse(t *testing.T) {
	body := fmt.Sprintf(
		parseTestTemplate,
		parseTestDBName,
		parseTestData[0].Rev,
		parseTestData[0].Prediction,
		parseTestData[0].ProbabilityFalse,
		parseTestData[0].ProbabilityTrue,
		parseTestData[1].Rev,
		parseTestData[1].Prediction,
		parseTestData[1].ProbabilityFalse,
		parseTestData[1].ProbabilityTrue)

	revs, err := parse([]byte(body), parseTestDBName, parseTestModel, parseTestData[0].Rev, parseTestData[1].Rev)
	assert.NoError(t, err)

	for _, info := range parseTestData {
		assert.Contains(t, revs, info.Rev)
		res := revs[info.Rev]

		assert.Equal(t, info.Prediction, res["prediction"].(bool))
		assert.Equal(t, info.ProbabilityFalse, res["probability"].(map[string]interface{})["false"].(float64))
		assert.Equal(t, info.ProbabilityTrue, res["probability"].(map[string]interface{})["true"].(float64))
	}
}

func TestParseError(t *testing.T) {
	body := fmt.Sprintf(parseTestErrorTemplate, parseTestDBName, parseTestErrorRev, parseTestErrorRev)
	_, err := parse([]byte(body), parseTestDBName, parseTestModel, parseTestErrorRev)
	assert.Error(t, err)
}

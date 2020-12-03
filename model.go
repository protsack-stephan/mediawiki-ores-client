package ores

var support = map[string]map[Model]bool{}

// Supported ores models
const (
	ModelDamaging Model = "damaging"
)

// Model ORES scoring type
type Model string

// Supports check if model is available for certain database
func (m Model) Supports(dbName string) bool {
	return true
}

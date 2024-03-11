package models

type NHTown struct {
	Money              float64               `json:"money"`
	Name               string                `json:"name"`
	Buildings          map[string]NHBuilding `json:"buildings"`
	HappeningsHappened map[string]struct{}   `json:"happeningsHappened"`
}

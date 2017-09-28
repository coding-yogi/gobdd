package requestmodels

// RecommendationRequest ...
type RecommendationRequest struct {
	StudentID string `json:"studId"`
}

// ConfigRequest ...
type ConfigRequest struct {
	Main ConfigRequestMain `json:"main"`
}

// ConfigRequestMain ...
type ConfigRequestMain struct {
	IsEnabled bool `json:"isEnabled"`
}

// PreferenceRequest ...
type PreferenceRequest struct {
	ID         string `json:"id"`
	Dismissed  bool   `json:"dismissed"`
	Read       bool   `json:"read"`
	Bookmarked bool   `json:"bookmarked"`
	Liked      bool   `json:"liked"`
}

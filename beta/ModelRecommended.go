// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RecommendedAction undocumented
type RecommendedAction struct {
	// Object is the base model of RecommendedAction
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ActionWebURL undocumented
	ActionWebURL *string `json:"actionWebUrl,omitempty"`
	// PotentialScoreImpact undocumented
	PotentialScoreImpact *float64 `json:"potentialScoreImpact,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
}

func NewRecommendedAction() (*RecommendedAction, error) {
	newRecommendedAction := &RecommendedAction{
		ODataType: "#microsoft.graph.RecommendedAction",
	}
	return newRecommendedAction, nil
}

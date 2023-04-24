// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Recommendation undocumented
type Recommendation struct {
	// RecommendationBase is the base model of Recommendation
	RecommendationBase

	ODataType string `json:"@odata.type,omitempty"`
}

func NewRecommendation() (*Recommendation, error) {
	newRecommendation := &Recommendation{
		ODataType: "#microsoft.graph.Recommendation",
	}
	return newRecommendation, nil
}

// RecommendationBase undocumented
type RecommendationBase struct {
	// Entity is the base model of RecommendationBase
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ActionSteps undocumented
	ActionSteps []ActionStep `json:"actionSteps,omitempty"`
	// Benefits undocumented
	Benefits *string `json:"benefits,omitempty"`
	// Category undocumented
	Category *RecommendationCategory `json:"category,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// CurrentScore undocumented
	CurrentScore *float64 `json:"currentScore,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// FeatureAreas undocumented
	FeatureAreas []RecommendationFeatureAreas `json:"featureAreas,omitempty"`
	// ImpactStartDateTime undocumented
	ImpactStartDateTime *time.Time `json:"impactStartDateTime,omitempty"`
	// ImpactType undocumented
	ImpactType *string `json:"impactType,omitempty"`
	// Insights undocumented
	Insights *string `json:"insights,omitempty"`
	// LastCheckedDateTime undocumented
	LastCheckedDateTime *time.Time `json:"lastCheckedDateTime,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// MaxScore undocumented
	MaxScore *float64 `json:"maxScore,omitempty"`
	// PostponeUntilDateTime undocumented
	PostponeUntilDateTime *time.Time `json:"postponeUntilDateTime,omitempty"`
	// Priority undocumented
	Priority *RecommendationPriority `json:"priority,omitempty"`
	// RecommendationType undocumented
	RecommendationType *RecommendationType `json:"recommendationType,omitempty"`
	// RemediationImpact undocumented
	RemediationImpact *string `json:"remediationImpact,omitempty"`
	// Status undocumented
	Status *RecommendationStatus `json:"status,omitempty"`
	// ImpactedResources undocumented
	ImpactedResources []ImpactedResource `json:"impactedResources,omitempty"`
}

func NewRecommendationBase() (*RecommendationBase, error) {
	newRecommendationBase := &RecommendationBase{
		ODataType: "#microsoft.graph.RecommendationBase",
	}
	return newRecommendationBase, nil
}

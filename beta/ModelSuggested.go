// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SuggestedEnrollmentLimit undocumented
type SuggestedEnrollmentLimit struct {
	// Object is the base model of SuggestedEnrollmentLimit
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// SuggestedDailyLimit undocumented
	SuggestedDailyLimit *int `json:"suggestedDailyLimit,omitempty"`
}

func NewSuggestedEnrollmentLimit() (*SuggestedEnrollmentLimit, error) {
	newSuggestedEnrollmentLimit := &SuggestedEnrollmentLimit{
		ODataType: "#microsoft.graph.SuggestedEnrollmentLimit",
	}
	return newSuggestedEnrollmentLimit, nil
}
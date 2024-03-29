// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ControlScore undocumented
type ControlScore struct {
	// Object is the base model of ControlScore
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ControlCategory undocumented
	ControlCategory *string `json:"controlCategory,omitempty"`
	// ControlName undocumented
	ControlName *string `json:"controlName,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// Score undocumented
	Score *float64 `json:"score,omitempty"`
}

func NewControlScore() (*ControlScore, error) {
	newControlScore := &ControlScore{
		ODataType: "#microsoft.graph.ControlScore",
	}
	return newControlScore, nil
}

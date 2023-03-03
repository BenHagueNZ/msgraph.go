// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RubricCriterion undocumented
type RubricCriterion struct {
	// Object is the base model of RubricCriterion
	Object
	// Description undocumented
	Description *EducationItemBody `json:"description,omitempty"`
}

// RubricLevel undocumented
type RubricLevel struct {
	// Object is the base model of RubricLevel
	Object
	// LevelID undocumented
	LevelID *string `json:"levelId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *EducationItemBody `json:"description,omitempty"`
	// Grading undocumented
	Grading *EducationAssignmentGradeType `json:"grading,omitempty"`
}

// RubricQuality undocumented
type RubricQuality struct {
	// Object is the base model of RubricQuality
	Object
	// QualityID undocumented
	QualityID *string `json:"qualityId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *EducationItemBody `json:"description,omitempty"`
	// Weight undocumented
	Weight *float64 `json:"weight,omitempty"`
	// Criteria undocumented
	Criteria []RubricCriterion `json:"criteria,omitempty"`
}

// RubricQualityFeedbackModel undocumented
type RubricQualityFeedbackModel struct {
	// Object is the base model of RubricQualityFeedbackModel
	Object
	// QualityID undocumented
	QualityID *string `json:"qualityId,omitempty"`
	// Feedback undocumented
	Feedback *EducationItemBody `json:"feedback,omitempty"`
}

// RubricQualitySelectedColumnModel undocumented
type RubricQualitySelectedColumnModel struct {
	// Object is the base model of RubricQualitySelectedColumnModel
	Object
	// QualityID undocumented
	QualityID *string `json:"qualityId,omitempty"`
	// ColumnID undocumented
	ColumnID *string `json:"columnId,omitempty"`
}

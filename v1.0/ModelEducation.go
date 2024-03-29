// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// EducationAssignment undocumented
type EducationAssignment struct {
	// Entity is the base model of EducationAssignment
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AddedStudentAction undocumented
	AddedStudentAction *EducationAddedStudentAction `json:"addedStudentAction,omitempty"`
	// AddToCalendarAction undocumented
	AddToCalendarAction *EducationAddToCalendarOptions `json:"addToCalendarAction,omitempty"`
	// AllowLateSubmissions undocumented
	AllowLateSubmissions *bool `json:"allowLateSubmissions,omitempty"`
	// AllowStudentsToAddResourcesToSubmission undocumented
	AllowStudentsToAddResourcesToSubmission *bool `json:"allowStudentsToAddResourcesToSubmission,omitempty"`
	// AssignDateTime undocumented
	AssignDateTime *time.Time `json:"assignDateTime,omitempty"`
	// AssignedDateTime undocumented
	AssignedDateTime *time.Time `json:"assignedDateTime,omitempty"`
	// AssignTo undocumented
	AssignTo *EducationAssignmentRecipient `json:"assignTo,omitempty"`
	// ClassID undocumented
	ClassID *string `json:"classId,omitempty"`
	// CloseDateTime undocumented
	CloseDateTime *time.Time `json:"closeDateTime,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// DueDateTime undocumented
	DueDateTime *time.Time `json:"dueDateTime,omitempty"`
	// FeedbackResourcesFolderURL undocumented
	FeedbackResourcesFolderURL *string `json:"feedbackResourcesFolderUrl,omitempty"`
	// Grading undocumented
	Grading *EducationAssignmentGradeType `json:"grading,omitempty"`
	// Instructions undocumented
	Instructions *EducationItemBody `json:"instructions,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// NotificationChannelURL undocumented
	NotificationChannelURL *string `json:"notificationChannelUrl,omitempty"`
	// ResourcesFolderURL undocumented
	ResourcesFolderURL *string `json:"resourcesFolderUrl,omitempty"`
	// Status undocumented
	Status *EducationAssignmentStatus `json:"status,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
	// Categories undocumented
	Categories []EducationCategory `json:"categories,omitempty"`
	// Resources undocumented
	Resources []EducationAssignmentResource `json:"resources,omitempty"`
	// Rubric undocumented
	Rubric *EducationRubric `json:"rubric,omitempty"`
	// Submissions undocumented
	Submissions []EducationSubmission `json:"submissions,omitempty"`
}

func NewEducationAssignment() (*EducationAssignment, error) {
	newEducationAssignment := &EducationAssignment{
		ODataType: "#microsoft.graph.EducationAssignment",
	}
	return newEducationAssignment, nil
}

// EducationAssignmentClassRecipient undocumented
type EducationAssignmentClassRecipient struct {
	// EducationAssignmentRecipient is the base model of EducationAssignmentClassRecipient
	EducationAssignmentRecipient

	ODataType string `json:"@odata.type,omitempty"`
}

func NewEducationAssignmentClassRecipient() (*EducationAssignmentClassRecipient, error) {
	newEducationAssignmentClassRecipient := &EducationAssignmentClassRecipient{
		ODataType: "#microsoft.graph.EducationAssignmentClassRecipient",
	}
	return newEducationAssignmentClassRecipient, nil
}

// EducationAssignmentDefaults undocumented
type EducationAssignmentDefaults struct {
	// Entity is the base model of EducationAssignmentDefaults
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AddedStudentAction undocumented
	AddedStudentAction *EducationAddedStudentAction `json:"addedStudentAction,omitempty"`
	// AddToCalendarAction undocumented
	AddToCalendarAction *EducationAddToCalendarOptions `json:"addToCalendarAction,omitempty"`
	// DueTime undocumented
	DueTime *TimeOfDay `json:"dueTime,omitempty"`
	// NotificationChannelURL undocumented
	NotificationChannelURL *string `json:"notificationChannelUrl,omitempty"`
}

func NewEducationAssignmentDefaults() (*EducationAssignmentDefaults, error) {
	newEducationAssignmentDefaults := &EducationAssignmentDefaults{
		ODataType: "#microsoft.graph.EducationAssignmentDefaults",
	}
	return newEducationAssignmentDefaults, nil
}

// EducationAssignmentGrade undocumented
type EducationAssignmentGrade struct {
	// Object is the base model of EducationAssignmentGrade
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// GradedBy undocumented
	GradedBy *IdentitySet `json:"gradedBy,omitempty"`
	// GradedDateTime undocumented
	GradedDateTime *time.Time `json:"gradedDateTime,omitempty"`
}

func NewEducationAssignmentGrade() (*EducationAssignmentGrade, error) {
	newEducationAssignmentGrade := &EducationAssignmentGrade{
		ODataType: "#microsoft.graph.EducationAssignmentGrade",
	}
	return newEducationAssignmentGrade, nil
}

// EducationAssignmentGradeType undocumented
type EducationAssignmentGradeType struct {
	// Object is the base model of EducationAssignmentGradeType
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewEducationAssignmentGradeType() (*EducationAssignmentGradeType, error) {
	newEducationAssignmentGradeType := &EducationAssignmentGradeType{
		ODataType: "#microsoft.graph.EducationAssignmentGradeType",
	}
	return newEducationAssignmentGradeType, nil
}

// EducationAssignmentGroupRecipient undocumented
type EducationAssignmentGroupRecipient struct {
	// EducationAssignmentRecipient is the base model of EducationAssignmentGroupRecipient
	EducationAssignmentRecipient

	ODataType string `json:"@odata.type,omitempty"`
}

func NewEducationAssignmentGroupRecipient() (*EducationAssignmentGroupRecipient, error) {
	newEducationAssignmentGroupRecipient := &EducationAssignmentGroupRecipient{
		ODataType: "#microsoft.graph.EducationAssignmentGroupRecipient",
	}
	return newEducationAssignmentGroupRecipient, nil
}

// EducationAssignmentIndividualRecipient undocumented
type EducationAssignmentIndividualRecipient struct {
	// EducationAssignmentRecipient is the base model of EducationAssignmentIndividualRecipient
	EducationAssignmentRecipient

	ODataType string `json:"@odata.type,omitempty"`
	// Recipients undocumented
	Recipients []string `json:"recipients,omitempty"`
}

func NewEducationAssignmentIndividualRecipient() (*EducationAssignmentIndividualRecipient, error) {
	newEducationAssignmentIndividualRecipient := &EducationAssignmentIndividualRecipient{
		ODataType: "#microsoft.graph.EducationAssignmentIndividualRecipient",
	}
	return newEducationAssignmentIndividualRecipient, nil
}

// EducationAssignmentPointsGrade undocumented
type EducationAssignmentPointsGrade struct {
	// EducationAssignmentGrade is the base model of EducationAssignmentPointsGrade
	EducationAssignmentGrade

	ODataType string `json:"@odata.type,omitempty"`
	// Points undocumented
	Points *float64 `json:"points,omitempty"`
}

func NewEducationAssignmentPointsGrade() (*EducationAssignmentPointsGrade, error) {
	newEducationAssignmentPointsGrade := &EducationAssignmentPointsGrade{
		ODataType: "#microsoft.graph.EducationAssignmentPointsGrade",
	}
	return newEducationAssignmentPointsGrade, nil
}

// EducationAssignmentPointsGradeType undocumented
type EducationAssignmentPointsGradeType struct {
	// EducationAssignmentGradeType is the base model of EducationAssignmentPointsGradeType
	EducationAssignmentGradeType

	ODataType string `json:"@odata.type,omitempty"`
	// MaxPoints undocumented
	MaxPoints *float64 `json:"maxPoints,omitempty"`
}

func NewEducationAssignmentPointsGradeType() (*EducationAssignmentPointsGradeType, error) {
	newEducationAssignmentPointsGradeType := &EducationAssignmentPointsGradeType{
		ODataType: "#microsoft.graph.EducationAssignmentPointsGradeType",
	}
	return newEducationAssignmentPointsGradeType, nil
}

// EducationAssignmentRecipient undocumented
type EducationAssignmentRecipient struct {
	// Object is the base model of EducationAssignmentRecipient
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewEducationAssignmentRecipient() (*EducationAssignmentRecipient, error) {
	newEducationAssignmentRecipient := &EducationAssignmentRecipient{
		ODataType: "#microsoft.graph.EducationAssignmentRecipient",
	}
	return newEducationAssignmentRecipient, nil
}

// EducationAssignmentResource undocumented
type EducationAssignmentResource struct {
	// Entity is the base model of EducationAssignmentResource
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DistributeForStudentWork undocumented
	DistributeForStudentWork *bool `json:"distributeForStudentWork,omitempty"`
	// Resource undocumented
	Resource *EducationResource `json:"resource,omitempty"`
}

func NewEducationAssignmentResource() (*EducationAssignmentResource, error) {
	newEducationAssignmentResource := &EducationAssignmentResource{
		ODataType: "#microsoft.graph.EducationAssignmentResource",
	}
	return newEducationAssignmentResource, nil
}

// EducationAssignmentSettings undocumented
type EducationAssignmentSettings struct {
	// Entity is the base model of EducationAssignmentSettings
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// SubmissionAnimationDisabled undocumented
	SubmissionAnimationDisabled *bool `json:"submissionAnimationDisabled,omitempty"`
}

func NewEducationAssignmentSettings() (*EducationAssignmentSettings, error) {
	newEducationAssignmentSettings := &EducationAssignmentSettings{
		ODataType: "#microsoft.graph.EducationAssignmentSettings",
	}
	return newEducationAssignmentSettings, nil
}

// EducationCategory undocumented
type EducationCategory struct {
	// Entity is the base model of EducationCategory
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

func NewEducationCategory() (*EducationCategory, error) {
	newEducationCategory := &EducationCategory{
		ODataType: "#microsoft.graph.EducationCategory",
	}
	return newEducationCategory, nil
}

// EducationClass undocumented
type EducationClass struct {
	// Entity is the base model of EducationClass
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ClassCode undocumented
	ClassCode *string `json:"classCode,omitempty"`
	// Course undocumented
	Course *EducationCourse `json:"course,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// ExternalName undocumented
	ExternalName *string `json:"externalName,omitempty"`
	// ExternalSource undocumented
	ExternalSource *EducationExternalSource `json:"externalSource,omitempty"`
	// ExternalSourceDetail undocumented
	ExternalSourceDetail *string `json:"externalSourceDetail,omitempty"`
	// Grade undocumented
	Grade *string `json:"grade,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// Term undocumented
	Term *EducationTerm `json:"term,omitempty"`
	// AssignmentCategories undocumented
	AssignmentCategories []EducationCategory `json:"assignmentCategories,omitempty"`
	// AssignmentDefaults undocumented
	AssignmentDefaults *EducationAssignmentDefaults `json:"assignmentDefaults,omitempty"`
	// Assignments undocumented
	Assignments []EducationAssignment `json:"assignments,omitempty"`
	// AssignmentSettings undocumented
	AssignmentSettings *EducationAssignmentSettings `json:"assignmentSettings,omitempty"`
	// Group undocumented
	Group *Group `json:"group,omitempty"`
	// Members undocumented
	Members []EducationUser `json:"members,omitempty"`
	// Schools undocumented
	Schools []EducationSchool `json:"schools,omitempty"`
	// Teachers undocumented
	Teachers []EducationUser `json:"teachers,omitempty"`
}

func NewEducationClass() (*EducationClass, error) {
	newEducationClass := &EducationClass{
		ODataType: "#microsoft.graph.EducationClass",
	}
	return newEducationClass, nil
}

// EducationCourse undocumented
type EducationCourse struct {
	// Object is the base model of EducationCourse
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CourseNumber undocumented
	CourseNumber *string `json:"courseNumber,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// Subject undocumented
	Subject *string `json:"subject,omitempty"`
}

func NewEducationCourse() (*EducationCourse, error) {
	newEducationCourse := &EducationCourse{
		ODataType: "#microsoft.graph.EducationCourse",
	}
	return newEducationCourse, nil
}

// EducationExcelResource undocumented
type EducationExcelResource struct {
	// EducationResource is the base model of EducationExcelResource
	EducationResource

	ODataType string `json:"@odata.type,omitempty"`
	// FileURL undocumented
	FileURL *string `json:"fileUrl,omitempty"`
}

func NewEducationExcelResource() (*EducationExcelResource, error) {
	newEducationExcelResource := &EducationExcelResource{
		ODataType: "#microsoft.graph.EducationExcelResource",
	}
	return newEducationExcelResource, nil
}

// EducationExternalResource undocumented
type EducationExternalResource struct {
	// EducationResource is the base model of EducationExternalResource
	EducationResource

	ODataType string `json:"@odata.type,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}

func NewEducationExternalResource() (*EducationExternalResource, error) {
	newEducationExternalResource := &EducationExternalResource{
		ODataType: "#microsoft.graph.EducationExternalResource",
	}
	return newEducationExternalResource, nil
}

// EducationFeedback undocumented
type EducationFeedback struct {
	// Object is the base model of EducationFeedback
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// FeedbackBy undocumented
	FeedbackBy *IdentitySet `json:"feedbackBy,omitempty"`
	// FeedbackDateTime undocumented
	FeedbackDateTime *time.Time `json:"feedbackDateTime,omitempty"`
	// Text undocumented
	Text *EducationItemBody `json:"text,omitempty"`
}

func NewEducationFeedback() (*EducationFeedback, error) {
	newEducationFeedback := &EducationFeedback{
		ODataType: "#microsoft.graph.EducationFeedback",
	}
	return newEducationFeedback, nil
}

// EducationFeedbackOutcome undocumented
type EducationFeedbackOutcome struct {
	// EducationOutcome is the base model of EducationFeedbackOutcome
	EducationOutcome

	ODataType string `json:"@odata.type,omitempty"`
	// Feedback undocumented
	Feedback *EducationFeedback `json:"feedback,omitempty"`
	// PublishedFeedback undocumented
	PublishedFeedback *EducationFeedback `json:"publishedFeedback,omitempty"`
}

func NewEducationFeedbackOutcome() (*EducationFeedbackOutcome, error) {
	newEducationFeedbackOutcome := &EducationFeedbackOutcome{
		ODataType: "#microsoft.graph.EducationFeedbackOutcome",
	}
	return newEducationFeedbackOutcome, nil
}

// EducationFeedbackResourceOutcome undocumented
type EducationFeedbackResourceOutcome struct {
	// EducationOutcome is the base model of EducationFeedbackResourceOutcome
	EducationOutcome

	ODataType string `json:"@odata.type,omitempty"`
	// FeedbackResource undocumented
	FeedbackResource *EducationResource `json:"feedbackResource,omitempty"`
	// ResourceStatus undocumented
	ResourceStatus *EducationFeedbackResourceOutcomeStatus `json:"resourceStatus,omitempty"`
}

func NewEducationFeedbackResourceOutcome() (*EducationFeedbackResourceOutcome, error) {
	newEducationFeedbackResourceOutcome := &EducationFeedbackResourceOutcome{
		ODataType: "#microsoft.graph.EducationFeedbackResourceOutcome",
	}
	return newEducationFeedbackResourceOutcome, nil
}

// EducationFileResource undocumented
type EducationFileResource struct {
	// EducationResource is the base model of EducationFileResource
	EducationResource

	ODataType string `json:"@odata.type,omitempty"`
	// FileURL undocumented
	FileURL *string `json:"fileUrl,omitempty"`
}

func NewEducationFileResource() (*EducationFileResource, error) {
	newEducationFileResource := &EducationFileResource{
		ODataType: "#microsoft.graph.EducationFileResource",
	}
	return newEducationFileResource, nil
}

// EducationItemBody undocumented
type EducationItemBody struct {
	// Object is the base model of EducationItemBody
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Content undocumented
	Content *string `json:"content,omitempty"`
	// ContentType undocumented
	ContentType *BodyType `json:"contentType,omitempty"`
}

func NewEducationItemBody() (*EducationItemBody, error) {
	newEducationItemBody := &EducationItemBody{
		ODataType: "#microsoft.graph.EducationItemBody",
	}
	return newEducationItemBody, nil
}

// EducationLinkResource undocumented
type EducationLinkResource struct {
	// EducationResource is the base model of EducationLinkResource
	EducationResource

	ODataType string `json:"@odata.type,omitempty"`
	// Link undocumented
	Link *string `json:"link,omitempty"`
}

func NewEducationLinkResource() (*EducationLinkResource, error) {
	newEducationLinkResource := &EducationLinkResource{
		ODataType: "#microsoft.graph.EducationLinkResource",
	}
	return newEducationLinkResource, nil
}

// EducationMediaResource undocumented
type EducationMediaResource struct {
	// EducationResource is the base model of EducationMediaResource
	EducationResource

	ODataType string `json:"@odata.type,omitempty"`
	// FileURL undocumented
	FileURL *string `json:"fileUrl,omitempty"`
}

func NewEducationMediaResource() (*EducationMediaResource, error) {
	newEducationMediaResource := &EducationMediaResource{
		ODataType: "#microsoft.graph.EducationMediaResource",
	}
	return newEducationMediaResource, nil
}

// EducationOnPremisesInfo undocumented
type EducationOnPremisesInfo struct {
	// Object is the base model of EducationOnPremisesInfo
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ImmutableID undocumented
	ImmutableID *string `json:"immutableId,omitempty"`
}

func NewEducationOnPremisesInfo() (*EducationOnPremisesInfo, error) {
	newEducationOnPremisesInfo := &EducationOnPremisesInfo{
		ODataType: "#microsoft.graph.EducationOnPremisesInfo",
	}
	return newEducationOnPremisesInfo, nil
}

// EducationOrganization undocumented
type EducationOrganization struct {
	// Entity is the base model of EducationOrganization
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalSource undocumented
	ExternalSource *EducationExternalSource `json:"externalSource,omitempty"`
	// ExternalSourceDetail undocumented
	ExternalSourceDetail *string `json:"externalSourceDetail,omitempty"`
}

func NewEducationOrganization() (*EducationOrganization, error) {
	newEducationOrganization := &EducationOrganization{
		ODataType: "#microsoft.graph.EducationOrganization",
	}
	return newEducationOrganization, nil
}

// EducationOutcome undocumented
type EducationOutcome struct {
	// Entity is the base model of EducationOutcome
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

func NewEducationOutcome() (*EducationOutcome, error) {
	newEducationOutcome := &EducationOutcome{
		ODataType: "#microsoft.graph.EducationOutcome",
	}
	return newEducationOutcome, nil
}

// EducationPointsOutcome undocumented
type EducationPointsOutcome struct {
	// EducationOutcome is the base model of EducationPointsOutcome
	EducationOutcome

	ODataType string `json:"@odata.type,omitempty"`
	// Points undocumented
	Points *EducationAssignmentPointsGrade `json:"points,omitempty"`
	// PublishedPoints undocumented
	PublishedPoints *EducationAssignmentPointsGrade `json:"publishedPoints,omitempty"`
}

func NewEducationPointsOutcome() (*EducationPointsOutcome, error) {
	newEducationPointsOutcome := &EducationPointsOutcome{
		ODataType: "#microsoft.graph.EducationPointsOutcome",
	}
	return newEducationPointsOutcome, nil
}

// EducationPowerPointResource undocumented
type EducationPowerPointResource struct {
	// EducationResource is the base model of EducationPowerPointResource
	EducationResource

	ODataType string `json:"@odata.type,omitempty"`
	// FileURL undocumented
	FileURL *string `json:"fileUrl,omitempty"`
}

func NewEducationPowerPointResource() (*EducationPowerPointResource, error) {
	newEducationPowerPointResource := &EducationPowerPointResource{
		ODataType: "#microsoft.graph.EducationPowerPointResource",
	}
	return newEducationPowerPointResource, nil
}

// EducationResource undocumented
type EducationResource struct {
	// Object is the base model of EducationResource
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

func NewEducationResource() (*EducationResource, error) {
	newEducationResource := &EducationResource{
		ODataType: "#microsoft.graph.EducationResource",
	}
	return newEducationResource, nil
}

// EducationRoot undocumented
type EducationRoot struct {
	// Object is the base model of EducationRoot
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Classes undocumented
	Classes []EducationClass `json:"classes,omitempty"`
	// Me undocumented
	Me *EducationUser `json:"me,omitempty"`
	// Schools undocumented
	Schools []EducationSchool `json:"schools,omitempty"`
	// Users undocumented
	Users []EducationUser `json:"users,omitempty"`
}

func NewEducationRoot() (*EducationRoot, error) {
	newEducationRoot := &EducationRoot{
		ODataType: "#microsoft.graph.EducationRoot",
	}
	return newEducationRoot, nil
}

// EducationRubric undocumented
type EducationRubric struct {
	// Entity is the base model of EducationRubric
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *EducationItemBody `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Grading undocumented
	Grading *EducationAssignmentGradeType `json:"grading,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Levels undocumented
	Levels []RubricLevel `json:"levels,omitempty"`
	// Qualities undocumented
	Qualities []RubricQuality `json:"qualities,omitempty"`
}

func NewEducationRubric() (*EducationRubric, error) {
	newEducationRubric := &EducationRubric{
		ODataType: "#microsoft.graph.EducationRubric",
	}
	return newEducationRubric, nil
}

// EducationRubricOutcome undocumented
type EducationRubricOutcome struct {
	// EducationOutcome is the base model of EducationRubricOutcome
	EducationOutcome

	ODataType string `json:"@odata.type,omitempty"`
	// PublishedRubricQualityFeedback undocumented
	PublishedRubricQualityFeedback []RubricQualityFeedbackModel `json:"publishedRubricQualityFeedback,omitempty"`
	// PublishedRubricQualitySelectedLevels undocumented
	PublishedRubricQualitySelectedLevels []RubricQualitySelectedColumnModel `json:"publishedRubricQualitySelectedLevels,omitempty"`
	// RubricQualityFeedback undocumented
	RubricQualityFeedback []RubricQualityFeedbackModel `json:"rubricQualityFeedback,omitempty"`
	// RubricQualitySelectedLevels undocumented
	RubricQualitySelectedLevels []RubricQualitySelectedColumnModel `json:"rubricQualitySelectedLevels,omitempty"`
}

func NewEducationRubricOutcome() (*EducationRubricOutcome, error) {
	newEducationRubricOutcome := &EducationRubricOutcome{
		ODataType: "#microsoft.graph.EducationRubricOutcome",
	}
	return newEducationRubricOutcome, nil
}

// EducationSchool undocumented
type EducationSchool struct {
	// EducationOrganization is the base model of EducationSchool
	EducationOrganization

	ODataType string `json:"@odata.type,omitempty"`
	// Address undocumented
	Address *PhysicalAddress `json:"address,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// ExternalPrincipalID undocumented
	ExternalPrincipalID *string `json:"externalPrincipalId,omitempty"`
	// Fax undocumented
	Fax *string `json:"fax,omitempty"`
	// HighestGrade undocumented
	HighestGrade *string `json:"highestGrade,omitempty"`
	// LowestGrade undocumented
	LowestGrade *string `json:"lowestGrade,omitempty"`
	// Phone undocumented
	Phone *string `json:"phone,omitempty"`
	// PrincipalEmail undocumented
	PrincipalEmail *string `json:"principalEmail,omitempty"`
	// PrincipalName undocumented
	PrincipalName *string `json:"principalName,omitempty"`
	// SchoolNumber undocumented
	SchoolNumber *string `json:"schoolNumber,omitempty"`
	// AdministrativeUnit undocumented
	AdministrativeUnit *AdministrativeUnit `json:"administrativeUnit,omitempty"`
	// Classes undocumented
	Classes []EducationClass `json:"classes,omitempty"`
	// Users undocumented
	Users []EducationUser `json:"users,omitempty"`
}

func NewEducationSchool() (*EducationSchool, error) {
	newEducationSchool := &EducationSchool{
		ODataType: "#microsoft.graph.EducationSchool",
	}
	return newEducationSchool, nil
}

// EducationStudent undocumented
type EducationStudent struct {
	// Object is the base model of EducationStudent
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// BirthDate undocumented
	BirthDate *Date `json:"birthDate,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// Gender undocumented
	Gender *EducationGender `json:"gender,omitempty"`
	// Grade undocumented
	Grade *string `json:"grade,omitempty"`
	// GraduationYear undocumented
	GraduationYear *string `json:"graduationYear,omitempty"`
	// StudentNumber undocumented
	StudentNumber *string `json:"studentNumber,omitempty"`
}

func NewEducationStudent() (*EducationStudent, error) {
	newEducationStudent := &EducationStudent{
		ODataType: "#microsoft.graph.EducationStudent",
	}
	return newEducationStudent, nil
}

// EducationSubmission undocumented
type EducationSubmission struct {
	// Entity is the base model of EducationSubmission
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ReassignedBy undocumented
	ReassignedBy *IdentitySet `json:"reassignedBy,omitempty"`
	// ReassignedDateTime undocumented
	ReassignedDateTime *time.Time `json:"reassignedDateTime,omitempty"`
	// Recipient undocumented
	Recipient *EducationSubmissionRecipient `json:"recipient,omitempty"`
	// ResourcesFolderURL undocumented
	ResourcesFolderURL *string `json:"resourcesFolderUrl,omitempty"`
	// ReturnedBy undocumented
	ReturnedBy *IdentitySet `json:"returnedBy,omitempty"`
	// ReturnedDateTime undocumented
	ReturnedDateTime *time.Time `json:"returnedDateTime,omitempty"`
	// Status undocumented
	Status *EducationSubmissionStatus `json:"status,omitempty"`
	// SubmittedBy undocumented
	SubmittedBy *IdentitySet `json:"submittedBy,omitempty"`
	// SubmittedDateTime undocumented
	SubmittedDateTime *time.Time `json:"submittedDateTime,omitempty"`
	// UnsubmittedBy undocumented
	UnsubmittedBy *IdentitySet `json:"unsubmittedBy,omitempty"`
	// UnsubmittedDateTime undocumented
	UnsubmittedDateTime *time.Time `json:"unsubmittedDateTime,omitempty"`
	// Outcomes undocumented
	Outcomes []EducationOutcome `json:"outcomes,omitempty"`
	// Resources undocumented
	Resources []EducationSubmissionResource `json:"resources,omitempty"`
	// SubmittedResources undocumented
	SubmittedResources []EducationSubmissionResource `json:"submittedResources,omitempty"`
}

func NewEducationSubmission() (*EducationSubmission, error) {
	newEducationSubmission := &EducationSubmission{
		ODataType: "#microsoft.graph.EducationSubmission",
	}
	return newEducationSubmission, nil
}

// EducationSubmissionIndividualRecipient undocumented
type EducationSubmissionIndividualRecipient struct {
	// EducationSubmissionRecipient is the base model of EducationSubmissionIndividualRecipient
	EducationSubmissionRecipient

	ODataType string `json:"@odata.type,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
}

func NewEducationSubmissionIndividualRecipient() (*EducationSubmissionIndividualRecipient, error) {
	newEducationSubmissionIndividualRecipient := &EducationSubmissionIndividualRecipient{
		ODataType: "#microsoft.graph.EducationSubmissionIndividualRecipient",
	}
	return newEducationSubmissionIndividualRecipient, nil
}

// EducationSubmissionRecipient undocumented
type EducationSubmissionRecipient struct {
	// Object is the base model of EducationSubmissionRecipient
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewEducationSubmissionRecipient() (*EducationSubmissionRecipient, error) {
	newEducationSubmissionRecipient := &EducationSubmissionRecipient{
		ODataType: "#microsoft.graph.EducationSubmissionRecipient",
	}
	return newEducationSubmissionRecipient, nil
}

// EducationSubmissionResource undocumented
type EducationSubmissionResource struct {
	// Entity is the base model of EducationSubmissionResource
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AssignmentResourceURL undocumented
	AssignmentResourceURL *string `json:"assignmentResourceUrl,omitempty"`
	// Resource undocumented
	Resource *EducationResource `json:"resource,omitempty"`
}

func NewEducationSubmissionResource() (*EducationSubmissionResource, error) {
	newEducationSubmissionResource := &EducationSubmissionResource{
		ODataType: "#microsoft.graph.EducationSubmissionResource",
	}
	return newEducationSubmissionResource, nil
}

// EducationTeacher undocumented
type EducationTeacher struct {
	// Object is the base model of EducationTeacher
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// TeacherNumber undocumented
	TeacherNumber *string `json:"teacherNumber,omitempty"`
}

func NewEducationTeacher() (*EducationTeacher, error) {
	newEducationTeacher := &EducationTeacher{
		ODataType: "#microsoft.graph.EducationTeacher",
	}
	return newEducationTeacher, nil
}

// EducationTeamsAppResource undocumented
type EducationTeamsAppResource struct {
	// EducationResource is the base model of EducationTeamsAppResource
	EducationResource

	ODataType string `json:"@odata.type,omitempty"`
	// AppIconWebURL undocumented
	AppIconWebURL *string `json:"appIconWebUrl,omitempty"`
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// TeamsEmbeddedContentURL undocumented
	TeamsEmbeddedContentURL *string `json:"teamsEmbeddedContentUrl,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}

func NewEducationTeamsAppResource() (*EducationTeamsAppResource, error) {
	newEducationTeamsAppResource := &EducationTeamsAppResource{
		ODataType: "#microsoft.graph.EducationTeamsAppResource",
	}
	return newEducationTeamsAppResource, nil
}

// EducationTerm undocumented
type EducationTerm struct {
	// Object is the base model of EducationTerm
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EndDate undocumented
	EndDate *Date `json:"endDate,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// StartDate undocumented
	StartDate *Date `json:"startDate,omitempty"`
}

func NewEducationTerm() (*EducationTerm, error) {
	newEducationTerm := &EducationTerm{
		ODataType: "#microsoft.graph.EducationTerm",
	}
	return newEducationTerm, nil
}

// EducationUser undocumented
type EducationUser struct {
	// Entity is the base model of EducationUser
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// RelatedContacts undocumented
	RelatedContacts []RelatedContact `json:"relatedContacts,omitempty"`
	// AccountEnabled undocumented
	AccountEnabled *bool `json:"accountEnabled,omitempty"`
	// AssignedLicenses undocumented
	AssignedLicenses []AssignedLicense `json:"assignedLicenses,omitempty"`
	// AssignedPlans undocumented
	AssignedPlans []AssignedPlan `json:"assignedPlans,omitempty"`
	// BusinessPhones undocumented
	BusinessPhones []string `json:"businessPhones,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// Department undocumented
	Department *string `json:"department,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalSource undocumented
	ExternalSource *EducationExternalSource `json:"externalSource,omitempty"`
	// ExternalSourceDetail undocumented
	ExternalSourceDetail *string `json:"externalSourceDetail,omitempty"`
	// GivenName undocumented
	GivenName *string `json:"givenName,omitempty"`
	// Mail undocumented
	Mail *string `json:"mail,omitempty"`
	// MailingAddress undocumented
	MailingAddress *PhysicalAddress `json:"mailingAddress,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// MiddleName undocumented
	MiddleName *string `json:"middleName,omitempty"`
	// MobilePhone undocumented
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// OfficeLocation undocumented
	OfficeLocation *string `json:"officeLocation,omitempty"`
	// OnPremisesInfo undocumented
	OnPremisesInfo *EducationOnPremisesInfo `json:"onPremisesInfo,omitempty"`
	// PasswordPolicies undocumented
	PasswordPolicies *string `json:"passwordPolicies,omitempty"`
	// PasswordProfile undocumented
	PasswordProfile *PasswordProfile `json:"passwordProfile,omitempty"`
	// PreferredLanguage undocumented
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	// PrimaryRole undocumented
	PrimaryRole *EducationUserRole `json:"primaryRole,omitempty"`
	// ProvisionedPlans undocumented
	ProvisionedPlans []ProvisionedPlan `json:"provisionedPlans,omitempty"`
	// RefreshTokensValidFromDateTime undocumented
	RefreshTokensValidFromDateTime *time.Time `json:"refreshTokensValidFromDateTime,omitempty"`
	// ResidenceAddress undocumented
	ResidenceAddress *PhysicalAddress `json:"residenceAddress,omitempty"`
	// ShowInAddressList undocumented
	ShowInAddressList *bool `json:"showInAddressList,omitempty"`
	// Student undocumented
	Student *EducationStudent `json:"student,omitempty"`
	// Surname undocumented
	Surname *string `json:"surname,omitempty"`
	// Teacher undocumented
	Teacher *EducationTeacher `json:"teacher,omitempty"`
	// UsageLocation undocumented
	UsageLocation *string `json:"usageLocation,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// UserType undocumented
	UserType *string `json:"userType,omitempty"`
	// Assignments undocumented
	Assignments []EducationAssignment `json:"assignments,omitempty"`
	// Rubrics undocumented
	Rubrics []EducationRubric `json:"rubrics,omitempty"`
	// Classes undocumented
	Classes []EducationClass `json:"classes,omitempty"`
	// Schools undocumented
	Schools []EducationSchool `json:"schools,omitempty"`
	// TaughtClasses undocumented
	TaughtClasses []EducationClass `json:"taughtClasses,omitempty"`
	// User undocumented
	User *User `json:"user,omitempty"`
}

func NewEducationUser() (*EducationUser, error) {
	newEducationUser := &EducationUser{
		ODataType: "#microsoft.graph.EducationUser",
	}
	return newEducationUser, nil
}

// EducationWordResource undocumented
type EducationWordResource struct {
	// EducationResource is the base model of EducationWordResource
	EducationResource

	ODataType string `json:"@odata.type,omitempty"`
	// FileURL undocumented
	FileURL *string `json:"fileUrl,omitempty"`
}

func NewEducationWordResource() (*EducationWordResource, error) {
	newEducationWordResource := &EducationWordResource{
		ODataType: "#microsoft.graph.EducationWordResource",
	}
	return newEducationWordResource, nil
}

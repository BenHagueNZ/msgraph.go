// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// SubjectRightsRequestObject undocumented
type SubjectRightsRequestObject struct {
	// Entity is the base model of SubjectRightsRequestObject
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AssignedTo undocumented
	AssignedTo *Identity `json:"assignedTo,omitempty"`
	// ClosedDateTime undocumented
	ClosedDateTime *time.Time `json:"closedDateTime,omitempty"`
	// ContentQuery undocumented
	ContentQuery *string `json:"contentQuery,omitempty"`
	// CreatedBy undocumented
	CreatedBy *IdentitySet `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DataSubject undocumented
	DataSubject *DataSubject `json:"dataSubject,omitempty"`
	// DataSubjectType undocumented
	DataSubjectType *DataSubjectType `json:"dataSubjectType,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// History undocumented
	History []SubjectRightsRequestHistory `json:"history,omitempty"`
	// IncludeAllVersions undocumented
	IncludeAllVersions *bool `json:"includeAllVersions,omitempty"`
	// IncludeAuthoredContent undocumented
	IncludeAuthoredContent *bool `json:"includeAuthoredContent,omitempty"`
	// Insight undocumented
	Insight *SubjectRightsRequestDetail `json:"insight,omitempty"`
	// InternalDueDateTime undocumented
	InternalDueDateTime *time.Time `json:"internalDueDateTime,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Mailboxlocations undocumented
	Mailboxlocations *SubjectRightsRequestMailboxLocation `json:"mailboxlocations,omitempty"`
	// PauseAfterEstimate undocumented
	PauseAfterEstimate *bool `json:"pauseAfterEstimate,omitempty"`
	// Regulations undocumented
	Regulations []string `json:"regulations,omitempty"`
	// Sitelocations undocumented
	Sitelocations *SubjectRightsRequestSiteLocation `json:"sitelocations,omitempty"`
	// Stages undocumented
	Stages []SubjectRightsRequestStageDetail `json:"stages,omitempty"`
	// Status undocumented
	Status *SubjectRightsRequestStatus `json:"status,omitempty"`
	// Type undocumented
	Type *SubjectRightsRequestType `json:"type,omitempty"`
	// Approvers undocumented
	Approvers []User `json:"approvers,omitempty"`
	// Collaborators undocumented
	Collaborators []User `json:"collaborators,omitempty"`
	// Notes undocumented
	Notes []AuthoredNote `json:"notes,omitempty"`
	// Team undocumented
	Team *Team `json:"team,omitempty"`
}

func NewSubjectRightsRequestObject() (*SubjectRightsRequestObject, error) {
	newSubjectRightsRequestObject := &SubjectRightsRequestObject{
		ODataType: "#microsoft.graph.SubjectRightsRequestObject",
	}
	return newSubjectRightsRequestObject, nil
}

// SubjectRightsRequestAllMailboxLocation undocumented
type SubjectRightsRequestAllMailboxLocation struct {
	// SubjectRightsRequestMailboxLocation is the base model of SubjectRightsRequestAllMailboxLocation
	SubjectRightsRequestMailboxLocation

	ODataType string `json:"@odata.type,omitempty"`
}

func NewSubjectRightsRequestAllMailboxLocation() (*SubjectRightsRequestAllMailboxLocation, error) {
	newSubjectRightsRequestAllMailboxLocation := &SubjectRightsRequestAllMailboxLocation{
		ODataType: "#microsoft.graph.SubjectRightsRequestAllMailboxLocation",
	}
	return newSubjectRightsRequestAllMailboxLocation, nil
}

// SubjectRightsRequestAllSiteLocation undocumented
type SubjectRightsRequestAllSiteLocation struct {
	// SubjectRightsRequestSiteLocation is the base model of SubjectRightsRequestAllSiteLocation
	SubjectRightsRequestSiteLocation

	ODataType string `json:"@odata.type,omitempty"`
}

func NewSubjectRightsRequestAllSiteLocation() (*SubjectRightsRequestAllSiteLocation, error) {
	newSubjectRightsRequestAllSiteLocation := &SubjectRightsRequestAllSiteLocation{
		ODataType: "#microsoft.graph.SubjectRightsRequestAllSiteLocation",
	}
	return newSubjectRightsRequestAllSiteLocation, nil
}

// SubjectRightsRequestDetail undocumented
type SubjectRightsRequestDetail struct {
	// Object is the base model of SubjectRightsRequestDetail
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ExcludedItemCount undocumented
	ExcludedItemCount *int `json:"excludedItemCount,omitempty"`
	// InsightCounts undocumented
	InsightCounts []KeyValuePair `json:"insightCounts,omitempty"`
	// ItemCount undocumented
	ItemCount *int `json:"itemCount,omitempty"`
	// ItemNeedReview undocumented
	ItemNeedReview *int `json:"itemNeedReview,omitempty"`
	// ProductItemCounts undocumented
	ProductItemCounts []KeyValuePair `json:"productItemCounts,omitempty"`
	// SignedOffItemCount undocumented
	SignedOffItemCount *int `json:"signedOffItemCount,omitempty"`
	// TotalItemSize undocumented
	TotalItemSize *int `json:"totalItemSize,omitempty"`
}

func NewSubjectRightsRequestDetail() (*SubjectRightsRequestDetail, error) {
	newSubjectRightsRequestDetail := &SubjectRightsRequestDetail{
		ODataType: "#microsoft.graph.SubjectRightsRequestDetail",
	}
	return newSubjectRightsRequestDetail, nil
}

// SubjectRightsRequestEnumeratedMailboxLocation undocumented
type SubjectRightsRequestEnumeratedMailboxLocation struct {
	// SubjectRightsRequestMailboxLocation is the base model of SubjectRightsRequestEnumeratedMailboxLocation
	SubjectRightsRequestMailboxLocation

	ODataType string `json:"@odata.type,omitempty"`
	// Upns undocumented
	Upns []string `json:"upns,omitempty"`
}

func NewSubjectRightsRequestEnumeratedMailboxLocation() (*SubjectRightsRequestEnumeratedMailboxLocation, error) {
	newSubjectRightsRequestEnumeratedMailboxLocation := &SubjectRightsRequestEnumeratedMailboxLocation{
		ODataType: "#microsoft.graph.SubjectRightsRequestEnumeratedMailboxLocation",
	}
	return newSubjectRightsRequestEnumeratedMailboxLocation, nil
}

// SubjectRightsRequestEnumeratedSiteLocation undocumented
type SubjectRightsRequestEnumeratedSiteLocation struct {
	// SubjectRightsRequestSiteLocation is the base model of SubjectRightsRequestEnumeratedSiteLocation
	SubjectRightsRequestSiteLocation

	ODataType string `json:"@odata.type,omitempty"`
	// Urls undocumented
	Urls []string `json:"urls,omitempty"`
}

func NewSubjectRightsRequestEnumeratedSiteLocation() (*SubjectRightsRequestEnumeratedSiteLocation, error) {
	newSubjectRightsRequestEnumeratedSiteLocation := &SubjectRightsRequestEnumeratedSiteLocation{
		ODataType: "#microsoft.graph.SubjectRightsRequestEnumeratedSiteLocation",
	}
	return newSubjectRightsRequestEnumeratedSiteLocation, nil
}

// SubjectRightsRequestHistory undocumented
type SubjectRightsRequestHistory struct {
	// Object is the base model of SubjectRightsRequestHistory
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ChangedBy undocumented
	ChangedBy *IdentitySet `json:"changedBy,omitempty"`
	// EventDateTime undocumented
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`
	// Stage undocumented
	Stage *SubjectRightsRequestStage `json:"stage,omitempty"`
	// StageStatus undocumented
	StageStatus *SubjectRightsRequestStageStatus `json:"stageStatus,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

func NewSubjectRightsRequestHistory() (*SubjectRightsRequestHistory, error) {
	newSubjectRightsRequestHistory := &SubjectRightsRequestHistory{
		ODataType: "#microsoft.graph.SubjectRightsRequestHistory",
	}
	return newSubjectRightsRequestHistory, nil
}

// SubjectRightsRequestMailboxLocation undocumented
type SubjectRightsRequestMailboxLocation struct {
	// Object is the base model of SubjectRightsRequestMailboxLocation
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewSubjectRightsRequestMailboxLocation() (*SubjectRightsRequestMailboxLocation, error) {
	newSubjectRightsRequestMailboxLocation := &SubjectRightsRequestMailboxLocation{
		ODataType: "#microsoft.graph.SubjectRightsRequestMailboxLocation",
	}
	return newSubjectRightsRequestMailboxLocation, nil
}

// SubjectRightsRequestSiteLocation undocumented
type SubjectRightsRequestSiteLocation struct {
	// Object is the base model of SubjectRightsRequestSiteLocation
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewSubjectRightsRequestSiteLocation() (*SubjectRightsRequestSiteLocation, error) {
	newSubjectRightsRequestSiteLocation := &SubjectRightsRequestSiteLocation{
		ODataType: "#microsoft.graph.SubjectRightsRequestSiteLocation",
	}
	return newSubjectRightsRequestSiteLocation, nil
}

// SubjectRightsRequestStageDetail undocumented
type SubjectRightsRequestStageDetail struct {
	// Object is the base model of SubjectRightsRequestStageDetail
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Error undocumented
	Error *PublicError `json:"error,omitempty"`
	// Stage undocumented
	Stage *SubjectRightsRequestStage `json:"stage,omitempty"`
	// Status undocumented
	Status *SubjectRightsRequestStageStatus `json:"status,omitempty"`
}

func NewSubjectRightsRequestStageDetail() (*SubjectRightsRequestStageDetail, error) {
	newSubjectRightsRequestStageDetail := &SubjectRightsRequestStageDetail{
		ODataType: "#microsoft.graph.SubjectRightsRequestStageDetail",
	}
	return newSubjectRightsRequestStageDetail, nil
}

// SubjectSet undocumented
type SubjectSet struct {
	// Object is the base model of SubjectSet
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewSubjectSet() (*SubjectSet, error) {
	newSubjectSet := &SubjectSet{
		ODataType: "#microsoft.graph.SubjectSet",
	}
	return newSubjectSet, nil
}

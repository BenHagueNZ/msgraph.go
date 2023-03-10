// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// SubjectRightsRequestStage undocumented
type SubjectRightsRequestStage string

const (
	// SubjectRightsRequestStageVContentRetrieval undocumented
	SubjectRightsRequestStageVContentRetrieval SubjectRightsRequestStage = "contentRetrieval"
	// SubjectRightsRequestStageVContentReview undocumented
	SubjectRightsRequestStageVContentReview SubjectRightsRequestStage = "contentReview"
	// SubjectRightsRequestStageVGenerateReport undocumented
	SubjectRightsRequestStageVGenerateReport SubjectRightsRequestStage = "generateReport"
	// SubjectRightsRequestStageVContentDeletion undocumented
	SubjectRightsRequestStageVContentDeletion SubjectRightsRequestStage = "contentDeletion"
	// SubjectRightsRequestStageVCaseResolved undocumented
	SubjectRightsRequestStageVCaseResolved SubjectRightsRequestStage = "caseResolved"
	// SubjectRightsRequestStageVContentEstimate undocumented
	SubjectRightsRequestStageVContentEstimate SubjectRightsRequestStage = "contentEstimate"
	// SubjectRightsRequestStageVUnknownFutureValue undocumented
	SubjectRightsRequestStageVUnknownFutureValue SubjectRightsRequestStage = "unknownFutureValue"
)

var (
	// SubjectRightsRequestStagePContentRetrieval is a pointer to SubjectRightsRequestStageVContentRetrieval
	SubjectRightsRequestStagePContentRetrieval = &_SubjectRightsRequestStagePContentRetrieval
	// SubjectRightsRequestStagePContentReview is a pointer to SubjectRightsRequestStageVContentReview
	SubjectRightsRequestStagePContentReview = &_SubjectRightsRequestStagePContentReview
	// SubjectRightsRequestStagePGenerateReport is a pointer to SubjectRightsRequestStageVGenerateReport
	SubjectRightsRequestStagePGenerateReport = &_SubjectRightsRequestStagePGenerateReport
	// SubjectRightsRequestStagePContentDeletion is a pointer to SubjectRightsRequestStageVContentDeletion
	SubjectRightsRequestStagePContentDeletion = &_SubjectRightsRequestStagePContentDeletion
	// SubjectRightsRequestStagePCaseResolved is a pointer to SubjectRightsRequestStageVCaseResolved
	SubjectRightsRequestStagePCaseResolved = &_SubjectRightsRequestStagePCaseResolved
	// SubjectRightsRequestStagePContentEstimate is a pointer to SubjectRightsRequestStageVContentEstimate
	SubjectRightsRequestStagePContentEstimate = &_SubjectRightsRequestStagePContentEstimate
	// SubjectRightsRequestStagePUnknownFutureValue is a pointer to SubjectRightsRequestStageVUnknownFutureValue
	SubjectRightsRequestStagePUnknownFutureValue = &_SubjectRightsRequestStagePUnknownFutureValue
)

var (
	_SubjectRightsRequestStagePContentRetrieval   = SubjectRightsRequestStageVContentRetrieval
	_SubjectRightsRequestStagePContentReview      = SubjectRightsRequestStageVContentReview
	_SubjectRightsRequestStagePGenerateReport     = SubjectRightsRequestStageVGenerateReport
	_SubjectRightsRequestStagePContentDeletion    = SubjectRightsRequestStageVContentDeletion
	_SubjectRightsRequestStagePCaseResolved       = SubjectRightsRequestStageVCaseResolved
	_SubjectRightsRequestStagePContentEstimate    = SubjectRightsRequestStageVContentEstimate
	_SubjectRightsRequestStagePUnknownFutureValue = SubjectRightsRequestStageVUnknownFutureValue
)

// SubjectRightsRequestStageStatus undocumented
type SubjectRightsRequestStageStatus string

const (
	// SubjectRightsRequestStageStatusVNotStarted undocumented
	SubjectRightsRequestStageStatusVNotStarted SubjectRightsRequestStageStatus = "notStarted"
	// SubjectRightsRequestStageStatusVCurrent undocumented
	SubjectRightsRequestStageStatusVCurrent SubjectRightsRequestStageStatus = "current"
	// SubjectRightsRequestStageStatusVCompleted undocumented
	SubjectRightsRequestStageStatusVCompleted SubjectRightsRequestStageStatus = "completed"
	// SubjectRightsRequestStageStatusVFailed undocumented
	SubjectRightsRequestStageStatusVFailed SubjectRightsRequestStageStatus = "failed"
	// SubjectRightsRequestStageStatusVUnknownFutureValue undocumented
	SubjectRightsRequestStageStatusVUnknownFutureValue SubjectRightsRequestStageStatus = "unknownFutureValue"
)

var (
	// SubjectRightsRequestStageStatusPNotStarted is a pointer to SubjectRightsRequestStageStatusVNotStarted
	SubjectRightsRequestStageStatusPNotStarted = &_SubjectRightsRequestStageStatusPNotStarted
	// SubjectRightsRequestStageStatusPCurrent is a pointer to SubjectRightsRequestStageStatusVCurrent
	SubjectRightsRequestStageStatusPCurrent = &_SubjectRightsRequestStageStatusPCurrent
	// SubjectRightsRequestStageStatusPCompleted is a pointer to SubjectRightsRequestStageStatusVCompleted
	SubjectRightsRequestStageStatusPCompleted = &_SubjectRightsRequestStageStatusPCompleted
	// SubjectRightsRequestStageStatusPFailed is a pointer to SubjectRightsRequestStageStatusVFailed
	SubjectRightsRequestStageStatusPFailed = &_SubjectRightsRequestStageStatusPFailed
	// SubjectRightsRequestStageStatusPUnknownFutureValue is a pointer to SubjectRightsRequestStageStatusVUnknownFutureValue
	SubjectRightsRequestStageStatusPUnknownFutureValue = &_SubjectRightsRequestStageStatusPUnknownFutureValue
)

var (
	_SubjectRightsRequestStageStatusPNotStarted         = SubjectRightsRequestStageStatusVNotStarted
	_SubjectRightsRequestStageStatusPCurrent            = SubjectRightsRequestStageStatusVCurrent
	_SubjectRightsRequestStageStatusPCompleted          = SubjectRightsRequestStageStatusVCompleted
	_SubjectRightsRequestStageStatusPFailed             = SubjectRightsRequestStageStatusVFailed
	_SubjectRightsRequestStageStatusPUnknownFutureValue = SubjectRightsRequestStageStatusVUnknownFutureValue
)

// SubjectRightsRequestStatus undocumented
type SubjectRightsRequestStatus string

const (
	// SubjectRightsRequestStatusVActive undocumented
	SubjectRightsRequestStatusVActive SubjectRightsRequestStatus = "active"
	// SubjectRightsRequestStatusVClosed undocumented
	SubjectRightsRequestStatusVClosed SubjectRightsRequestStatus = "closed"
	// SubjectRightsRequestStatusVUnknownFutureValue undocumented
	SubjectRightsRequestStatusVUnknownFutureValue SubjectRightsRequestStatus = "unknownFutureValue"
)

var (
	// SubjectRightsRequestStatusPActive is a pointer to SubjectRightsRequestStatusVActive
	SubjectRightsRequestStatusPActive = &_SubjectRightsRequestStatusPActive
	// SubjectRightsRequestStatusPClosed is a pointer to SubjectRightsRequestStatusVClosed
	SubjectRightsRequestStatusPClosed = &_SubjectRightsRequestStatusPClosed
	// SubjectRightsRequestStatusPUnknownFutureValue is a pointer to SubjectRightsRequestStatusVUnknownFutureValue
	SubjectRightsRequestStatusPUnknownFutureValue = &_SubjectRightsRequestStatusPUnknownFutureValue
)

var (
	_SubjectRightsRequestStatusPActive             = SubjectRightsRequestStatusVActive
	_SubjectRightsRequestStatusPClosed             = SubjectRightsRequestStatusVClosed
	_SubjectRightsRequestStatusPUnknownFutureValue = SubjectRightsRequestStatusVUnknownFutureValue
)

// SubjectRightsRequestType undocumented
type SubjectRightsRequestType string

const (
	// SubjectRightsRequestTypeVExport undocumented
	SubjectRightsRequestTypeVExport SubjectRightsRequestType = "export"
	// SubjectRightsRequestTypeVDelete undocumented
	SubjectRightsRequestTypeVDelete SubjectRightsRequestType = "delete"
	// SubjectRightsRequestTypeVAccess undocumented
	SubjectRightsRequestTypeVAccess SubjectRightsRequestType = "access"
	// SubjectRightsRequestTypeVTagForAction undocumented
	SubjectRightsRequestTypeVTagForAction SubjectRightsRequestType = "tagForAction"
	// SubjectRightsRequestTypeVUnknownFutureValue undocumented
	SubjectRightsRequestTypeVUnknownFutureValue SubjectRightsRequestType = "unknownFutureValue"
)

var (
	// SubjectRightsRequestTypePExport is a pointer to SubjectRightsRequestTypeVExport
	SubjectRightsRequestTypePExport = &_SubjectRightsRequestTypePExport
	// SubjectRightsRequestTypePDelete is a pointer to SubjectRightsRequestTypeVDelete
	SubjectRightsRequestTypePDelete = &_SubjectRightsRequestTypePDelete
	// SubjectRightsRequestTypePAccess is a pointer to SubjectRightsRequestTypeVAccess
	SubjectRightsRequestTypePAccess = &_SubjectRightsRequestTypePAccess
	// SubjectRightsRequestTypePTagForAction is a pointer to SubjectRightsRequestTypeVTagForAction
	SubjectRightsRequestTypePTagForAction = &_SubjectRightsRequestTypePTagForAction
	// SubjectRightsRequestTypePUnknownFutureValue is a pointer to SubjectRightsRequestTypeVUnknownFutureValue
	SubjectRightsRequestTypePUnknownFutureValue = &_SubjectRightsRequestTypePUnknownFutureValue
)

var (
	_SubjectRightsRequestTypePExport             = SubjectRightsRequestTypeVExport
	_SubjectRightsRequestTypePDelete             = SubjectRightsRequestTypeVDelete
	_SubjectRightsRequestTypePAccess             = SubjectRightsRequestTypeVAccess
	_SubjectRightsRequestTypePTagForAction       = SubjectRightsRequestTypeVTagForAction
	_SubjectRightsRequestTypePUnknownFutureValue = SubjectRightsRequestTypeVUnknownFutureValue
)

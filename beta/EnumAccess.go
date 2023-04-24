// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AccessLevel undocumented
type AccessLevel string

const (
	// AccessLevelVEveryone undocumented
	AccessLevelVEveryone AccessLevel = "everyone"
	// AccessLevelVInvited undocumented
	AccessLevelVInvited AccessLevel = "invited"
	// AccessLevelVLocked undocumented
	AccessLevelVLocked AccessLevel = "locked"
	// AccessLevelVSameEnterprise undocumented
	AccessLevelVSameEnterprise AccessLevel = "sameEnterprise"
	// AccessLevelVSameEnterpriseAndFederated undocumented
	AccessLevelVSameEnterpriseAndFederated AccessLevel = "sameEnterpriseAndFederated"
)

var (
	// AccessLevelPEveryone is a pointer to AccessLevelVEveryone
	AccessLevelPEveryone = &_AccessLevelPEveryone
	// AccessLevelPInvited is a pointer to AccessLevelVInvited
	AccessLevelPInvited = &_AccessLevelPInvited
	// AccessLevelPLocked is a pointer to AccessLevelVLocked
	AccessLevelPLocked = &_AccessLevelPLocked
	// AccessLevelPSameEnterprise is a pointer to AccessLevelVSameEnterprise
	AccessLevelPSameEnterprise = &_AccessLevelPSameEnterprise
	// AccessLevelPSameEnterpriseAndFederated is a pointer to AccessLevelVSameEnterpriseAndFederated
	AccessLevelPSameEnterpriseAndFederated = &_AccessLevelPSameEnterpriseAndFederated
)

var (
	_AccessLevelPEveryone                   = AccessLevelVEveryone
	_AccessLevelPInvited                    = AccessLevelVInvited
	_AccessLevelPLocked                     = AccessLevelVLocked
	_AccessLevelPSameEnterprise             = AccessLevelVSameEnterprise
	_AccessLevelPSameEnterpriseAndFederated = AccessLevelVSameEnterpriseAndFederated
)

// AccessPackageAssignmentFilterByCurrentUserOptions undocumented
type AccessPackageAssignmentFilterByCurrentUserOptions string

const (
	// AccessPackageAssignmentFilterByCurrentUserOptionsVTarget undocumented
	AccessPackageAssignmentFilterByCurrentUserOptionsVTarget AccessPackageAssignmentFilterByCurrentUserOptions = "target"
	// AccessPackageAssignmentFilterByCurrentUserOptionsVCreatedBy undocumented
	AccessPackageAssignmentFilterByCurrentUserOptionsVCreatedBy AccessPackageAssignmentFilterByCurrentUserOptions = "createdBy"
	// AccessPackageAssignmentFilterByCurrentUserOptionsVUnknownFutureValue undocumented
	AccessPackageAssignmentFilterByCurrentUserOptionsVUnknownFutureValue AccessPackageAssignmentFilterByCurrentUserOptions = "unknownFutureValue"
)

var (
	// AccessPackageAssignmentFilterByCurrentUserOptionsPTarget is a pointer to AccessPackageAssignmentFilterByCurrentUserOptionsVTarget
	AccessPackageAssignmentFilterByCurrentUserOptionsPTarget = &_AccessPackageAssignmentFilterByCurrentUserOptionsPTarget
	// AccessPackageAssignmentFilterByCurrentUserOptionsPCreatedBy is a pointer to AccessPackageAssignmentFilterByCurrentUserOptionsVCreatedBy
	AccessPackageAssignmentFilterByCurrentUserOptionsPCreatedBy = &_AccessPackageAssignmentFilterByCurrentUserOptionsPCreatedBy
	// AccessPackageAssignmentFilterByCurrentUserOptionsPUnknownFutureValue is a pointer to AccessPackageAssignmentFilterByCurrentUserOptionsVUnknownFutureValue
	AccessPackageAssignmentFilterByCurrentUserOptionsPUnknownFutureValue = &_AccessPackageAssignmentFilterByCurrentUserOptionsPUnknownFutureValue
)

var (
	_AccessPackageAssignmentFilterByCurrentUserOptionsPTarget             = AccessPackageAssignmentFilterByCurrentUserOptionsVTarget
	_AccessPackageAssignmentFilterByCurrentUserOptionsPCreatedBy          = AccessPackageAssignmentFilterByCurrentUserOptionsVCreatedBy
	_AccessPackageAssignmentFilterByCurrentUserOptionsPUnknownFutureValue = AccessPackageAssignmentFilterByCurrentUserOptionsVUnknownFutureValue
)

// AccessPackageAssignmentRequestFilterByCurrentUserOptions undocumented
type AccessPackageAssignmentRequestFilterByCurrentUserOptions string

const (
	// AccessPackageAssignmentRequestFilterByCurrentUserOptionsVTarget undocumented
	AccessPackageAssignmentRequestFilterByCurrentUserOptionsVTarget AccessPackageAssignmentRequestFilterByCurrentUserOptions = "target"
	// AccessPackageAssignmentRequestFilterByCurrentUserOptionsVCreatedBy undocumented
	AccessPackageAssignmentRequestFilterByCurrentUserOptionsVCreatedBy AccessPackageAssignmentRequestFilterByCurrentUserOptions = "createdBy"
	// AccessPackageAssignmentRequestFilterByCurrentUserOptionsVApprover undocumented
	AccessPackageAssignmentRequestFilterByCurrentUserOptionsVApprover AccessPackageAssignmentRequestFilterByCurrentUserOptions = "approver"
	// AccessPackageAssignmentRequestFilterByCurrentUserOptionsVUnknownFutureValue undocumented
	AccessPackageAssignmentRequestFilterByCurrentUserOptionsVUnknownFutureValue AccessPackageAssignmentRequestFilterByCurrentUserOptions = "unknownFutureValue"
)

var (
	// AccessPackageAssignmentRequestFilterByCurrentUserOptionsPTarget is a pointer to AccessPackageAssignmentRequestFilterByCurrentUserOptionsVTarget
	AccessPackageAssignmentRequestFilterByCurrentUserOptionsPTarget = &_AccessPackageAssignmentRequestFilterByCurrentUserOptionsPTarget
	// AccessPackageAssignmentRequestFilterByCurrentUserOptionsPCreatedBy is a pointer to AccessPackageAssignmentRequestFilterByCurrentUserOptionsVCreatedBy
	AccessPackageAssignmentRequestFilterByCurrentUserOptionsPCreatedBy = &_AccessPackageAssignmentRequestFilterByCurrentUserOptionsPCreatedBy
	// AccessPackageAssignmentRequestFilterByCurrentUserOptionsPApprover is a pointer to AccessPackageAssignmentRequestFilterByCurrentUserOptionsVApprover
	AccessPackageAssignmentRequestFilterByCurrentUserOptionsPApprover = &_AccessPackageAssignmentRequestFilterByCurrentUserOptionsPApprover
	// AccessPackageAssignmentRequestFilterByCurrentUserOptionsPUnknownFutureValue is a pointer to AccessPackageAssignmentRequestFilterByCurrentUserOptionsVUnknownFutureValue
	AccessPackageAssignmentRequestFilterByCurrentUserOptionsPUnknownFutureValue = &_AccessPackageAssignmentRequestFilterByCurrentUserOptionsPUnknownFutureValue
)

var (
	_AccessPackageAssignmentRequestFilterByCurrentUserOptionsPTarget             = AccessPackageAssignmentRequestFilterByCurrentUserOptionsVTarget
	_AccessPackageAssignmentRequestFilterByCurrentUserOptionsPCreatedBy          = AccessPackageAssignmentRequestFilterByCurrentUserOptionsVCreatedBy
	_AccessPackageAssignmentRequestFilterByCurrentUserOptionsPApprover           = AccessPackageAssignmentRequestFilterByCurrentUserOptionsVApprover
	_AccessPackageAssignmentRequestFilterByCurrentUserOptionsPUnknownFutureValue = AccessPackageAssignmentRequestFilterByCurrentUserOptionsVUnknownFutureValue
)

// AccessPackageCustomExtensionHandlerStatus undocumented
type AccessPackageCustomExtensionHandlerStatus string

const (
	// AccessPackageCustomExtensionHandlerStatusVRequestSent undocumented
	AccessPackageCustomExtensionHandlerStatusVRequestSent AccessPackageCustomExtensionHandlerStatus = "requestSent"
	// AccessPackageCustomExtensionHandlerStatusVRequestReceived undocumented
	AccessPackageCustomExtensionHandlerStatusVRequestReceived AccessPackageCustomExtensionHandlerStatus = "requestReceived"
	// AccessPackageCustomExtensionHandlerStatusVUnknownFutureValue undocumented
	AccessPackageCustomExtensionHandlerStatusVUnknownFutureValue AccessPackageCustomExtensionHandlerStatus = "unknownFutureValue"
)

var (
	// AccessPackageCustomExtensionHandlerStatusPRequestSent is a pointer to AccessPackageCustomExtensionHandlerStatusVRequestSent
	AccessPackageCustomExtensionHandlerStatusPRequestSent = &_AccessPackageCustomExtensionHandlerStatusPRequestSent
	// AccessPackageCustomExtensionHandlerStatusPRequestReceived is a pointer to AccessPackageCustomExtensionHandlerStatusVRequestReceived
	AccessPackageCustomExtensionHandlerStatusPRequestReceived = &_AccessPackageCustomExtensionHandlerStatusPRequestReceived
	// AccessPackageCustomExtensionHandlerStatusPUnknownFutureValue is a pointer to AccessPackageCustomExtensionHandlerStatusVUnknownFutureValue
	AccessPackageCustomExtensionHandlerStatusPUnknownFutureValue = &_AccessPackageCustomExtensionHandlerStatusPUnknownFutureValue
)

var (
	_AccessPackageCustomExtensionHandlerStatusPRequestSent        = AccessPackageCustomExtensionHandlerStatusVRequestSent
	_AccessPackageCustomExtensionHandlerStatusPRequestReceived    = AccessPackageCustomExtensionHandlerStatusVRequestReceived
	_AccessPackageCustomExtensionHandlerStatusPUnknownFutureValue = AccessPackageCustomExtensionHandlerStatusVUnknownFutureValue
)

// AccessPackageCustomExtensionStage undocumented
type AccessPackageCustomExtensionStage string

const (
	// AccessPackageCustomExtensionStageVAssignmentRequestCreated undocumented
	AccessPackageCustomExtensionStageVAssignmentRequestCreated AccessPackageCustomExtensionStage = "assignmentRequestCreated"
	// AccessPackageCustomExtensionStageVAssignmentRequestApproved undocumented
	AccessPackageCustomExtensionStageVAssignmentRequestApproved AccessPackageCustomExtensionStage = "assignmentRequestApproved"
	// AccessPackageCustomExtensionStageVAssignmentRequestGranted undocumented
	AccessPackageCustomExtensionStageVAssignmentRequestGranted AccessPackageCustomExtensionStage = "assignmentRequestGranted"
	// AccessPackageCustomExtensionStageVAssignmentRequestRemoved undocumented
	AccessPackageCustomExtensionStageVAssignmentRequestRemoved AccessPackageCustomExtensionStage = "assignmentRequestRemoved"
	// AccessPackageCustomExtensionStageVAssignmentFourteenDaysBeforeExpiration undocumented
	AccessPackageCustomExtensionStageVAssignmentFourteenDaysBeforeExpiration AccessPackageCustomExtensionStage = "assignmentFourteenDaysBeforeExpiration"
	// AccessPackageCustomExtensionStageVAssignmentOneDayBeforeExpiration undocumented
	AccessPackageCustomExtensionStageVAssignmentOneDayBeforeExpiration AccessPackageCustomExtensionStage = "assignmentOneDayBeforeExpiration"
	// AccessPackageCustomExtensionStageVUnknownFutureValue undocumented
	AccessPackageCustomExtensionStageVUnknownFutureValue AccessPackageCustomExtensionStage = "unknownFutureValue"
)

var (
	// AccessPackageCustomExtensionStagePAssignmentRequestCreated is a pointer to AccessPackageCustomExtensionStageVAssignmentRequestCreated
	AccessPackageCustomExtensionStagePAssignmentRequestCreated = &_AccessPackageCustomExtensionStagePAssignmentRequestCreated
	// AccessPackageCustomExtensionStagePAssignmentRequestApproved is a pointer to AccessPackageCustomExtensionStageVAssignmentRequestApproved
	AccessPackageCustomExtensionStagePAssignmentRequestApproved = &_AccessPackageCustomExtensionStagePAssignmentRequestApproved
	// AccessPackageCustomExtensionStagePAssignmentRequestGranted is a pointer to AccessPackageCustomExtensionStageVAssignmentRequestGranted
	AccessPackageCustomExtensionStagePAssignmentRequestGranted = &_AccessPackageCustomExtensionStagePAssignmentRequestGranted
	// AccessPackageCustomExtensionStagePAssignmentRequestRemoved is a pointer to AccessPackageCustomExtensionStageVAssignmentRequestRemoved
	AccessPackageCustomExtensionStagePAssignmentRequestRemoved = &_AccessPackageCustomExtensionStagePAssignmentRequestRemoved
	// AccessPackageCustomExtensionStagePAssignmentFourteenDaysBeforeExpiration is a pointer to AccessPackageCustomExtensionStageVAssignmentFourteenDaysBeforeExpiration
	AccessPackageCustomExtensionStagePAssignmentFourteenDaysBeforeExpiration = &_AccessPackageCustomExtensionStagePAssignmentFourteenDaysBeforeExpiration
	// AccessPackageCustomExtensionStagePAssignmentOneDayBeforeExpiration is a pointer to AccessPackageCustomExtensionStageVAssignmentOneDayBeforeExpiration
	AccessPackageCustomExtensionStagePAssignmentOneDayBeforeExpiration = &_AccessPackageCustomExtensionStagePAssignmentOneDayBeforeExpiration
	// AccessPackageCustomExtensionStagePUnknownFutureValue is a pointer to AccessPackageCustomExtensionStageVUnknownFutureValue
	AccessPackageCustomExtensionStagePUnknownFutureValue = &_AccessPackageCustomExtensionStagePUnknownFutureValue
)

var (
	_AccessPackageCustomExtensionStagePAssignmentRequestCreated               = AccessPackageCustomExtensionStageVAssignmentRequestCreated
	_AccessPackageCustomExtensionStagePAssignmentRequestApproved              = AccessPackageCustomExtensionStageVAssignmentRequestApproved
	_AccessPackageCustomExtensionStagePAssignmentRequestGranted               = AccessPackageCustomExtensionStageVAssignmentRequestGranted
	_AccessPackageCustomExtensionStagePAssignmentRequestRemoved               = AccessPackageCustomExtensionStageVAssignmentRequestRemoved
	_AccessPackageCustomExtensionStagePAssignmentFourteenDaysBeforeExpiration = AccessPackageCustomExtensionStageVAssignmentFourteenDaysBeforeExpiration
	_AccessPackageCustomExtensionStagePAssignmentOneDayBeforeExpiration       = AccessPackageCustomExtensionStageVAssignmentOneDayBeforeExpiration
	_AccessPackageCustomExtensionStagePUnknownFutureValue                     = AccessPackageCustomExtensionStageVUnknownFutureValue
)

// AccessPackageFilterByCurrentUserOptions undocumented
type AccessPackageFilterByCurrentUserOptions string

const (
	// AccessPackageFilterByCurrentUserOptionsVAllowedRequestor undocumented
	AccessPackageFilterByCurrentUserOptionsVAllowedRequestor AccessPackageFilterByCurrentUserOptions = "allowedRequestor"
	// AccessPackageFilterByCurrentUserOptionsVUnknownFutureValue undocumented
	AccessPackageFilterByCurrentUserOptionsVUnknownFutureValue AccessPackageFilterByCurrentUserOptions = "unknownFutureValue"
)

var (
	// AccessPackageFilterByCurrentUserOptionsPAllowedRequestor is a pointer to AccessPackageFilterByCurrentUserOptionsVAllowedRequestor
	AccessPackageFilterByCurrentUserOptionsPAllowedRequestor = &_AccessPackageFilterByCurrentUserOptionsPAllowedRequestor
	// AccessPackageFilterByCurrentUserOptionsPUnknownFutureValue is a pointer to AccessPackageFilterByCurrentUserOptionsVUnknownFutureValue
	AccessPackageFilterByCurrentUserOptionsPUnknownFutureValue = &_AccessPackageFilterByCurrentUserOptionsPUnknownFutureValue
)

var (
	_AccessPackageFilterByCurrentUserOptionsPAllowedRequestor   = AccessPackageFilterByCurrentUserOptionsVAllowedRequestor
	_AccessPackageFilterByCurrentUserOptionsPUnknownFutureValue = AccessPackageFilterByCurrentUserOptionsVUnknownFutureValue
)

// AccessPackageSubjectLifecycle undocumented
type AccessPackageSubjectLifecycle string

const (
	// AccessPackageSubjectLifecycleVNotDefined undocumented
	AccessPackageSubjectLifecycleVNotDefined AccessPackageSubjectLifecycle = "notDefined"
	// AccessPackageSubjectLifecycleVNotGoverned undocumented
	AccessPackageSubjectLifecycleVNotGoverned AccessPackageSubjectLifecycle = "notGoverned"
	// AccessPackageSubjectLifecycleVGoverned undocumented
	AccessPackageSubjectLifecycleVGoverned AccessPackageSubjectLifecycle = "governed"
	// AccessPackageSubjectLifecycleVUnknownFutureValue undocumented
	AccessPackageSubjectLifecycleVUnknownFutureValue AccessPackageSubjectLifecycle = "unknownFutureValue"
)

var (
	// AccessPackageSubjectLifecyclePNotDefined is a pointer to AccessPackageSubjectLifecycleVNotDefined
	AccessPackageSubjectLifecyclePNotDefined = &_AccessPackageSubjectLifecyclePNotDefined
	// AccessPackageSubjectLifecyclePNotGoverned is a pointer to AccessPackageSubjectLifecycleVNotGoverned
	AccessPackageSubjectLifecyclePNotGoverned = &_AccessPackageSubjectLifecyclePNotGoverned
	// AccessPackageSubjectLifecyclePGoverned is a pointer to AccessPackageSubjectLifecycleVGoverned
	AccessPackageSubjectLifecyclePGoverned = &_AccessPackageSubjectLifecyclePGoverned
	// AccessPackageSubjectLifecyclePUnknownFutureValue is a pointer to AccessPackageSubjectLifecycleVUnknownFutureValue
	AccessPackageSubjectLifecyclePUnknownFutureValue = &_AccessPackageSubjectLifecyclePUnknownFutureValue
)

var (
	_AccessPackageSubjectLifecyclePNotDefined         = AccessPackageSubjectLifecycleVNotDefined
	_AccessPackageSubjectLifecyclePNotGoverned        = AccessPackageSubjectLifecycleVNotGoverned
	_AccessPackageSubjectLifecyclePGoverned           = AccessPackageSubjectLifecycleVGoverned
	_AccessPackageSubjectLifecyclePUnknownFutureValue = AccessPackageSubjectLifecycleVUnknownFutureValue
)

// AccessReviewHistoryDecisionFilter undocumented
type AccessReviewHistoryDecisionFilter string

const (
	// AccessReviewHistoryDecisionFilterVApprove undocumented
	AccessReviewHistoryDecisionFilterVApprove AccessReviewHistoryDecisionFilter = "approve"
	// AccessReviewHistoryDecisionFilterVDeny undocumented
	AccessReviewHistoryDecisionFilterVDeny AccessReviewHistoryDecisionFilter = "deny"
	// AccessReviewHistoryDecisionFilterVNotReviewed undocumented
	AccessReviewHistoryDecisionFilterVNotReviewed AccessReviewHistoryDecisionFilter = "notReviewed"
	// AccessReviewHistoryDecisionFilterVDontKnow undocumented
	AccessReviewHistoryDecisionFilterVDontKnow AccessReviewHistoryDecisionFilter = "dontKnow"
	// AccessReviewHistoryDecisionFilterVNotNotified undocumented
	AccessReviewHistoryDecisionFilterVNotNotified AccessReviewHistoryDecisionFilter = "notNotified"
	// AccessReviewHistoryDecisionFilterVUnknownFutureValue undocumented
	AccessReviewHistoryDecisionFilterVUnknownFutureValue AccessReviewHistoryDecisionFilter = "unknownFutureValue"
)

var (
	// AccessReviewHistoryDecisionFilterPApprove is a pointer to AccessReviewHistoryDecisionFilterVApprove
	AccessReviewHistoryDecisionFilterPApprove = &_AccessReviewHistoryDecisionFilterPApprove
	// AccessReviewHistoryDecisionFilterPDeny is a pointer to AccessReviewHistoryDecisionFilterVDeny
	AccessReviewHistoryDecisionFilterPDeny = &_AccessReviewHistoryDecisionFilterPDeny
	// AccessReviewHistoryDecisionFilterPNotReviewed is a pointer to AccessReviewHistoryDecisionFilterVNotReviewed
	AccessReviewHistoryDecisionFilterPNotReviewed = &_AccessReviewHistoryDecisionFilterPNotReviewed
	// AccessReviewHistoryDecisionFilterPDontKnow is a pointer to AccessReviewHistoryDecisionFilterVDontKnow
	AccessReviewHistoryDecisionFilterPDontKnow = &_AccessReviewHistoryDecisionFilterPDontKnow
	// AccessReviewHistoryDecisionFilterPNotNotified is a pointer to AccessReviewHistoryDecisionFilterVNotNotified
	AccessReviewHistoryDecisionFilterPNotNotified = &_AccessReviewHistoryDecisionFilterPNotNotified
	// AccessReviewHistoryDecisionFilterPUnknownFutureValue is a pointer to AccessReviewHistoryDecisionFilterVUnknownFutureValue
	AccessReviewHistoryDecisionFilterPUnknownFutureValue = &_AccessReviewHistoryDecisionFilterPUnknownFutureValue
)

var (
	_AccessReviewHistoryDecisionFilterPApprove            = AccessReviewHistoryDecisionFilterVApprove
	_AccessReviewHistoryDecisionFilterPDeny               = AccessReviewHistoryDecisionFilterVDeny
	_AccessReviewHistoryDecisionFilterPNotReviewed        = AccessReviewHistoryDecisionFilterVNotReviewed
	_AccessReviewHistoryDecisionFilterPDontKnow           = AccessReviewHistoryDecisionFilterVDontKnow
	_AccessReviewHistoryDecisionFilterPNotNotified        = AccessReviewHistoryDecisionFilterVNotNotified
	_AccessReviewHistoryDecisionFilterPUnknownFutureValue = AccessReviewHistoryDecisionFilterVUnknownFutureValue
)

// AccessReviewHistoryStatus undocumented
type AccessReviewHistoryStatus string

const (
	// AccessReviewHistoryStatusVDone undocumented
	AccessReviewHistoryStatusVDone AccessReviewHistoryStatus = "done"
	// AccessReviewHistoryStatusVInprogress undocumented
	AccessReviewHistoryStatusVInprogress AccessReviewHistoryStatus = "inprogress"
	// AccessReviewHistoryStatusVError undocumented
	AccessReviewHistoryStatusVError AccessReviewHistoryStatus = "error"
	// AccessReviewHistoryStatusVRequested undocumented
	AccessReviewHistoryStatusVRequested AccessReviewHistoryStatus = "requested"
	// AccessReviewHistoryStatusVUnknownFutureValue undocumented
	AccessReviewHistoryStatusVUnknownFutureValue AccessReviewHistoryStatus = "unknownFutureValue"
)

var (
	// AccessReviewHistoryStatusPDone is a pointer to AccessReviewHistoryStatusVDone
	AccessReviewHistoryStatusPDone = &_AccessReviewHistoryStatusPDone
	// AccessReviewHistoryStatusPInprogress is a pointer to AccessReviewHistoryStatusVInprogress
	AccessReviewHistoryStatusPInprogress = &_AccessReviewHistoryStatusPInprogress
	// AccessReviewHistoryStatusPError is a pointer to AccessReviewHistoryStatusVError
	AccessReviewHistoryStatusPError = &_AccessReviewHistoryStatusPError
	// AccessReviewHistoryStatusPRequested is a pointer to AccessReviewHistoryStatusVRequested
	AccessReviewHistoryStatusPRequested = &_AccessReviewHistoryStatusPRequested
	// AccessReviewHistoryStatusPUnknownFutureValue is a pointer to AccessReviewHistoryStatusVUnknownFutureValue
	AccessReviewHistoryStatusPUnknownFutureValue = &_AccessReviewHistoryStatusPUnknownFutureValue
)

var (
	_AccessReviewHistoryStatusPDone               = AccessReviewHistoryStatusVDone
	_AccessReviewHistoryStatusPInprogress         = AccessReviewHistoryStatusVInprogress
	_AccessReviewHistoryStatusPError              = AccessReviewHistoryStatusVError
	_AccessReviewHistoryStatusPRequested          = AccessReviewHistoryStatusVRequested
	_AccessReviewHistoryStatusPUnknownFutureValue = AccessReviewHistoryStatusVUnknownFutureValue
)

// AccessReviewInstanceDecisionItemFilterByCurrentUserOptions undocumented
type AccessReviewInstanceDecisionItemFilterByCurrentUserOptions string

const (
	// AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsVReviewer undocumented
	AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsVReviewer AccessReviewInstanceDecisionItemFilterByCurrentUserOptions = "reviewer"
	// AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsVUnknownFutureValue undocumented
	AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsVUnknownFutureValue AccessReviewInstanceDecisionItemFilterByCurrentUserOptions = "unknownFutureValue"
)

var (
	// AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsPReviewer is a pointer to AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsVReviewer
	AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsPReviewer = &_AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsPReviewer
	// AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsPUnknownFutureValue is a pointer to AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsVUnknownFutureValue
	AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsPUnknownFutureValue = &_AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsPUnknownFutureValue
)

var (
	_AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsPReviewer           = AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsVReviewer
	_AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsPUnknownFutureValue = AccessReviewInstanceDecisionItemFilterByCurrentUserOptionsVUnknownFutureValue
)

// AccessReviewInstanceFilterByCurrentUserOptions undocumented
type AccessReviewInstanceFilterByCurrentUserOptions string

const (
	// AccessReviewInstanceFilterByCurrentUserOptionsVReviewer undocumented
	AccessReviewInstanceFilterByCurrentUserOptionsVReviewer AccessReviewInstanceFilterByCurrentUserOptions = "reviewer"
	// AccessReviewInstanceFilterByCurrentUserOptionsVUnknownFutureValue undocumented
	AccessReviewInstanceFilterByCurrentUserOptionsVUnknownFutureValue AccessReviewInstanceFilterByCurrentUserOptions = "unknownFutureValue"
)

var (
	// AccessReviewInstanceFilterByCurrentUserOptionsPReviewer is a pointer to AccessReviewInstanceFilterByCurrentUserOptionsVReviewer
	AccessReviewInstanceFilterByCurrentUserOptionsPReviewer = &_AccessReviewInstanceFilterByCurrentUserOptionsPReviewer
	// AccessReviewInstanceFilterByCurrentUserOptionsPUnknownFutureValue is a pointer to AccessReviewInstanceFilterByCurrentUserOptionsVUnknownFutureValue
	AccessReviewInstanceFilterByCurrentUserOptionsPUnknownFutureValue = &_AccessReviewInstanceFilterByCurrentUserOptionsPUnknownFutureValue
)

var (
	_AccessReviewInstanceFilterByCurrentUserOptionsPReviewer           = AccessReviewInstanceFilterByCurrentUserOptionsVReviewer
	_AccessReviewInstanceFilterByCurrentUserOptionsPUnknownFutureValue = AccessReviewInstanceFilterByCurrentUserOptionsVUnknownFutureValue
)

// AccessReviewScheduleDefinitionFilterByCurrentUserOptions undocumented
type AccessReviewScheduleDefinitionFilterByCurrentUserOptions string

const (
	// AccessReviewScheduleDefinitionFilterByCurrentUserOptionsVReviewer undocumented
	AccessReviewScheduleDefinitionFilterByCurrentUserOptionsVReviewer AccessReviewScheduleDefinitionFilterByCurrentUserOptions = "reviewer"
	// AccessReviewScheduleDefinitionFilterByCurrentUserOptionsVUnknownFutureValue undocumented
	AccessReviewScheduleDefinitionFilterByCurrentUserOptionsVUnknownFutureValue AccessReviewScheduleDefinitionFilterByCurrentUserOptions = "unknownFutureValue"
)

var (
	// AccessReviewScheduleDefinitionFilterByCurrentUserOptionsPReviewer is a pointer to AccessReviewScheduleDefinitionFilterByCurrentUserOptionsVReviewer
	AccessReviewScheduleDefinitionFilterByCurrentUserOptionsPReviewer = &_AccessReviewScheduleDefinitionFilterByCurrentUserOptionsPReviewer
	// AccessReviewScheduleDefinitionFilterByCurrentUserOptionsPUnknownFutureValue is a pointer to AccessReviewScheduleDefinitionFilterByCurrentUserOptionsVUnknownFutureValue
	AccessReviewScheduleDefinitionFilterByCurrentUserOptionsPUnknownFutureValue = &_AccessReviewScheduleDefinitionFilterByCurrentUserOptionsPUnknownFutureValue
)

var (
	_AccessReviewScheduleDefinitionFilterByCurrentUserOptionsPReviewer           = AccessReviewScheduleDefinitionFilterByCurrentUserOptionsVReviewer
	_AccessReviewScheduleDefinitionFilterByCurrentUserOptionsPUnknownFutureValue = AccessReviewScheduleDefinitionFilterByCurrentUserOptionsVUnknownFutureValue
)

// AccessReviewStageFilterByCurrentUserOptions undocumented
type AccessReviewStageFilterByCurrentUserOptions string

const (
	// AccessReviewStageFilterByCurrentUserOptionsVReviewer undocumented
	AccessReviewStageFilterByCurrentUserOptionsVReviewer AccessReviewStageFilterByCurrentUserOptions = "reviewer"
	// AccessReviewStageFilterByCurrentUserOptionsVUnknownFutureValue undocumented
	AccessReviewStageFilterByCurrentUserOptionsVUnknownFutureValue AccessReviewStageFilterByCurrentUserOptions = "unknownFutureValue"
)

var (
	// AccessReviewStageFilterByCurrentUserOptionsPReviewer is a pointer to AccessReviewStageFilterByCurrentUserOptionsVReviewer
	AccessReviewStageFilterByCurrentUserOptionsPReviewer = &_AccessReviewStageFilterByCurrentUserOptionsPReviewer
	// AccessReviewStageFilterByCurrentUserOptionsPUnknownFutureValue is a pointer to AccessReviewStageFilterByCurrentUserOptionsVUnknownFutureValue
	AccessReviewStageFilterByCurrentUserOptionsPUnknownFutureValue = &_AccessReviewStageFilterByCurrentUserOptionsPUnknownFutureValue
)

var (
	_AccessReviewStageFilterByCurrentUserOptionsPReviewer           = AccessReviewStageFilterByCurrentUserOptionsVReviewer
	_AccessReviewStageFilterByCurrentUserOptionsPUnknownFutureValue = AccessReviewStageFilterByCurrentUserOptionsVUnknownFutureValue
)

// AccessReviewTimeoutBehavior undocumented
type AccessReviewTimeoutBehavior string

const (
	// AccessReviewTimeoutBehaviorVKeepAccess undocumented
	AccessReviewTimeoutBehaviorVKeepAccess AccessReviewTimeoutBehavior = "keepAccess"
	// AccessReviewTimeoutBehaviorVRemoveAccess undocumented
	AccessReviewTimeoutBehaviorVRemoveAccess AccessReviewTimeoutBehavior = "removeAccess"
	// AccessReviewTimeoutBehaviorVAcceptAccessRecommendation undocumented
	AccessReviewTimeoutBehaviorVAcceptAccessRecommendation AccessReviewTimeoutBehavior = "acceptAccessRecommendation"
	// AccessReviewTimeoutBehaviorVUnknownFutureValue undocumented
	AccessReviewTimeoutBehaviorVUnknownFutureValue AccessReviewTimeoutBehavior = "unknownFutureValue"
)

var (
	// AccessReviewTimeoutBehaviorPKeepAccess is a pointer to AccessReviewTimeoutBehaviorVKeepAccess
	AccessReviewTimeoutBehaviorPKeepAccess = &_AccessReviewTimeoutBehaviorPKeepAccess
	// AccessReviewTimeoutBehaviorPRemoveAccess is a pointer to AccessReviewTimeoutBehaviorVRemoveAccess
	AccessReviewTimeoutBehaviorPRemoveAccess = &_AccessReviewTimeoutBehaviorPRemoveAccess
	// AccessReviewTimeoutBehaviorPAcceptAccessRecommendation is a pointer to AccessReviewTimeoutBehaviorVAcceptAccessRecommendation
	AccessReviewTimeoutBehaviorPAcceptAccessRecommendation = &_AccessReviewTimeoutBehaviorPAcceptAccessRecommendation
	// AccessReviewTimeoutBehaviorPUnknownFutureValue is a pointer to AccessReviewTimeoutBehaviorVUnknownFutureValue
	AccessReviewTimeoutBehaviorPUnknownFutureValue = &_AccessReviewTimeoutBehaviorPUnknownFutureValue
)

var (
	_AccessReviewTimeoutBehaviorPKeepAccess                 = AccessReviewTimeoutBehaviorVKeepAccess
	_AccessReviewTimeoutBehaviorPRemoveAccess               = AccessReviewTimeoutBehaviorVRemoveAccess
	_AccessReviewTimeoutBehaviorPAcceptAccessRecommendation = AccessReviewTimeoutBehaviorVAcceptAccessRecommendation
	_AccessReviewTimeoutBehaviorPUnknownFutureValue         = AccessReviewTimeoutBehaviorVUnknownFutureValue
)

// AccessType undocumented
type AccessType string

const (
	// AccessTypeVGrant undocumented
	AccessTypeVGrant AccessType = "grant"
	// AccessTypeVDeny undocumented
	AccessTypeVDeny AccessType = "deny"
)

var (
	// AccessTypePGrant is a pointer to AccessTypeVGrant
	AccessTypePGrant = &_AccessTypePGrant
	// AccessTypePDeny is a pointer to AccessTypeVDeny
	AccessTypePDeny = &_AccessTypePDeny
)

var (
	_AccessTypePGrant = AccessTypeVGrant
	_AccessTypePDeny  = AccessTypeVDeny
)
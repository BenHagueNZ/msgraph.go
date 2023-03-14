
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// ApprovalFilterByCurrentUserOptions undocumented
type ApprovalFilterByCurrentUserOptions string

const (
    // ApprovalFilterByCurrentUserOptionsVTarget undocumented
    ApprovalFilterByCurrentUserOptionsVTarget ApprovalFilterByCurrentUserOptions = "target"
    // ApprovalFilterByCurrentUserOptionsVCreatedBy undocumented
    ApprovalFilterByCurrentUserOptionsVCreatedBy ApprovalFilterByCurrentUserOptions = "createdBy"
    // ApprovalFilterByCurrentUserOptionsVApprover undocumented
    ApprovalFilterByCurrentUserOptionsVApprover ApprovalFilterByCurrentUserOptions = "approver"
    // ApprovalFilterByCurrentUserOptionsVUnknownFutureValue undocumented
    ApprovalFilterByCurrentUserOptionsVUnknownFutureValue ApprovalFilterByCurrentUserOptions = "unknownFutureValue"
)

var (
    // ApprovalFilterByCurrentUserOptionsPTarget is a pointer to ApprovalFilterByCurrentUserOptionsVTarget
    ApprovalFilterByCurrentUserOptionsPTarget = &_ApprovalFilterByCurrentUserOptionsPTarget
    // ApprovalFilterByCurrentUserOptionsPCreatedBy is a pointer to ApprovalFilterByCurrentUserOptionsVCreatedBy
    ApprovalFilterByCurrentUserOptionsPCreatedBy = &_ApprovalFilterByCurrentUserOptionsPCreatedBy
    // ApprovalFilterByCurrentUserOptionsPApprover is a pointer to ApprovalFilterByCurrentUserOptionsVApprover
    ApprovalFilterByCurrentUserOptionsPApprover = &_ApprovalFilterByCurrentUserOptionsPApprover
    // ApprovalFilterByCurrentUserOptionsPUnknownFutureValue is a pointer to ApprovalFilterByCurrentUserOptionsVUnknownFutureValue
    ApprovalFilterByCurrentUserOptionsPUnknownFutureValue = &_ApprovalFilterByCurrentUserOptionsPUnknownFutureValue
)

var (
    _ApprovalFilterByCurrentUserOptionsPTarget = ApprovalFilterByCurrentUserOptionsVTarget
    _ApprovalFilterByCurrentUserOptionsPCreatedBy = ApprovalFilterByCurrentUserOptionsVCreatedBy
    _ApprovalFilterByCurrentUserOptionsPApprover = ApprovalFilterByCurrentUserOptionsVApprover
    _ApprovalFilterByCurrentUserOptionsPUnknownFutureValue = ApprovalFilterByCurrentUserOptionsVUnknownFutureValue
)

// ApprovalState undocumented
type ApprovalState string

const (
    // ApprovalStateVPending undocumented
    ApprovalStateVPending ApprovalState = "pending"
    // ApprovalStateVApproved undocumented
    ApprovalStateVApproved ApprovalState = "approved"
    // ApprovalStateVDenied undocumented
    ApprovalStateVDenied ApprovalState = "denied"
    // ApprovalStateVAborted undocumented
    ApprovalStateVAborted ApprovalState = "aborted"
    // ApprovalStateVCanceled undocumented
    ApprovalStateVCanceled ApprovalState = "canceled"
)

var (
    // ApprovalStatePPending is a pointer to ApprovalStateVPending
    ApprovalStatePPending = &_ApprovalStatePPending
    // ApprovalStatePApproved is a pointer to ApprovalStateVApproved
    ApprovalStatePApproved = &_ApprovalStatePApproved
    // ApprovalStatePDenied is a pointer to ApprovalStateVDenied
    ApprovalStatePDenied = &_ApprovalStatePDenied
    // ApprovalStatePAborted is a pointer to ApprovalStateVAborted
    ApprovalStatePAborted = &_ApprovalStatePAborted
    // ApprovalStatePCanceled is a pointer to ApprovalStateVCanceled
    ApprovalStatePCanceled = &_ApprovalStatePCanceled
)

var (
    _ApprovalStatePPending = ApprovalStateVPending
    _ApprovalStatePApproved = ApprovalStateVApproved
    _ApprovalStatePDenied = ApprovalStateVDenied
    _ApprovalStatePAborted = ApprovalStateVAborted
    _ApprovalStatePCanceled = ApprovalStateVCanceled
)


// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// DriverApprovalAction undocumented
type DriverApprovalAction string

const (
    // DriverApprovalActionVApprove undocumented
    DriverApprovalActionVApprove DriverApprovalAction = "approve"
    // DriverApprovalActionVDecline undocumented
    DriverApprovalActionVDecline DriverApprovalAction = "decline"
    // DriverApprovalActionVSuspend undocumented
    DriverApprovalActionVSuspend DriverApprovalAction = "suspend"
)

var (
    // DriverApprovalActionPApprove is a pointer to DriverApprovalActionVApprove
    DriverApprovalActionPApprove = &_DriverApprovalActionPApprove
    // DriverApprovalActionPDecline is a pointer to DriverApprovalActionVDecline
    DriverApprovalActionPDecline = &_DriverApprovalActionPDecline
    // DriverApprovalActionPSuspend is a pointer to DriverApprovalActionVSuspend
    DriverApprovalActionPSuspend = &_DriverApprovalActionPSuspend
)

var (
    _DriverApprovalActionPApprove = DriverApprovalActionVApprove
    _DriverApprovalActionPDecline = DriverApprovalActionVDecline
    _DriverApprovalActionPSuspend = DriverApprovalActionVSuspend
)

// DriverApprovalStatus undocumented
type DriverApprovalStatus string

const (
    // DriverApprovalStatusVNeedsReview undocumented
    DriverApprovalStatusVNeedsReview DriverApprovalStatus = "needsReview"
    // DriverApprovalStatusVDeclined undocumented
    DriverApprovalStatusVDeclined DriverApprovalStatus = "declined"
    // DriverApprovalStatusVApproved undocumented
    DriverApprovalStatusVApproved DriverApprovalStatus = "approved"
    // DriverApprovalStatusVSuspended undocumented
    DriverApprovalStatusVSuspended DriverApprovalStatus = "suspended"
)

var (
    // DriverApprovalStatusPNeedsReview is a pointer to DriverApprovalStatusVNeedsReview
    DriverApprovalStatusPNeedsReview = &_DriverApprovalStatusPNeedsReview
    // DriverApprovalStatusPDeclined is a pointer to DriverApprovalStatusVDeclined
    DriverApprovalStatusPDeclined = &_DriverApprovalStatusPDeclined
    // DriverApprovalStatusPApproved is a pointer to DriverApprovalStatusVApproved
    DriverApprovalStatusPApproved = &_DriverApprovalStatusPApproved
    // DriverApprovalStatusPSuspended is a pointer to DriverApprovalStatusVSuspended
    DriverApprovalStatusPSuspended = &_DriverApprovalStatusPSuspended
)

var (
    _DriverApprovalStatusPNeedsReview = DriverApprovalStatusVNeedsReview
    _DriverApprovalStatusPDeclined = DriverApprovalStatusVDeclined
    _DriverApprovalStatusPApproved = DriverApprovalStatusVApproved
    _DriverApprovalStatusPSuspended = DriverApprovalStatusVSuspended
)

// DriverCategory undocumented
type DriverCategory string

const (
    // DriverCategoryVRecommended undocumented
    DriverCategoryVRecommended DriverCategory = "recommended"
    // DriverCategoryVPreviouslyApproved undocumented
    DriverCategoryVPreviouslyApproved DriverCategory = "previouslyApproved"
    // DriverCategoryVOther undocumented
    DriverCategoryVOther DriverCategory = "other"
)

var (
    // DriverCategoryPRecommended is a pointer to DriverCategoryVRecommended
    DriverCategoryPRecommended = &_DriverCategoryPRecommended
    // DriverCategoryPPreviouslyApproved is a pointer to DriverCategoryVPreviouslyApproved
    DriverCategoryPPreviouslyApproved = &_DriverCategoryPPreviouslyApproved
    // DriverCategoryPOther is a pointer to DriverCategoryVOther
    DriverCategoryPOther = &_DriverCategoryPOther
)

var (
    _DriverCategoryPRecommended = DriverCategoryVRecommended
    _DriverCategoryPPreviouslyApproved = DriverCategoryVPreviouslyApproved
    _DriverCategoryPOther = DriverCategoryVOther
)

// DriverUpdateProfileApprovalType undocumented
type DriverUpdateProfileApprovalType string

const (
    // DriverUpdateProfileApprovalTypeVManual undocumented
    DriverUpdateProfileApprovalTypeVManual DriverUpdateProfileApprovalType = "manual"
    // DriverUpdateProfileApprovalTypeVAutomatic undocumented
    DriverUpdateProfileApprovalTypeVAutomatic DriverUpdateProfileApprovalType = "automatic"
)

var (
    // DriverUpdateProfileApprovalTypePManual is a pointer to DriverUpdateProfileApprovalTypeVManual
    DriverUpdateProfileApprovalTypePManual = &_DriverUpdateProfileApprovalTypePManual
    // DriverUpdateProfileApprovalTypePAutomatic is a pointer to DriverUpdateProfileApprovalTypeVAutomatic
    DriverUpdateProfileApprovalTypePAutomatic = &_DriverUpdateProfileApprovalTypePAutomatic
)

var (
    _DriverUpdateProfileApprovalTypePManual = DriverUpdateProfileApprovalTypeVManual
    _DriverUpdateProfileApprovalTypePAutomatic = DriverUpdateProfileApprovalTypeVAutomatic
)

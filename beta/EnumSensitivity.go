// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Sensitivity undocumented
type Sensitivity string

const (
	// SensitivityVNormal undocumented
	SensitivityVNormal Sensitivity = "normal"
	// SensitivityVPersonal undocumented
	SensitivityVPersonal Sensitivity = "personal"
	// SensitivityVPrivate undocumented
	SensitivityVPrivate Sensitivity = "private"
	// SensitivityVConfidential undocumented
	SensitivityVConfidential Sensitivity = "confidential"
)

var (
	// SensitivityPNormal is a pointer to SensitivityVNormal
	SensitivityPNormal = &_SensitivityPNormal
	// SensitivityPPersonal is a pointer to SensitivityVPersonal
	SensitivityPPersonal = &_SensitivityPPersonal
	// SensitivityPPrivate is a pointer to SensitivityVPrivate
	SensitivityPPrivate = &_SensitivityPPrivate
	// SensitivityPConfidential is a pointer to SensitivityVConfidential
	SensitivityPConfidential = &_SensitivityPConfidential
)

var (
	_SensitivityPNormal       = SensitivityVNormal
	_SensitivityPPersonal     = SensitivityVPersonal
	_SensitivityPPrivate      = SensitivityVPrivate
	_SensitivityPConfidential = SensitivityVConfidential
)

// SensitivityLabelAssignmentMethod undocumented
type SensitivityLabelAssignmentMethod string

const (
	// SensitivityLabelAssignmentMethodVStandard undocumented
	SensitivityLabelAssignmentMethodVStandard SensitivityLabelAssignmentMethod = "standard"
	// SensitivityLabelAssignmentMethodVPrivileged undocumented
	SensitivityLabelAssignmentMethodVPrivileged SensitivityLabelAssignmentMethod = "privileged"
	// SensitivityLabelAssignmentMethodVAuto undocumented
	SensitivityLabelAssignmentMethodVAuto SensitivityLabelAssignmentMethod = "auto"
	// SensitivityLabelAssignmentMethodVUnknownFutureValue undocumented
	SensitivityLabelAssignmentMethodVUnknownFutureValue SensitivityLabelAssignmentMethod = "unknownFutureValue"
)

var (
	// SensitivityLabelAssignmentMethodPStandard is a pointer to SensitivityLabelAssignmentMethodVStandard
	SensitivityLabelAssignmentMethodPStandard = &_SensitivityLabelAssignmentMethodPStandard
	// SensitivityLabelAssignmentMethodPPrivileged is a pointer to SensitivityLabelAssignmentMethodVPrivileged
	SensitivityLabelAssignmentMethodPPrivileged = &_SensitivityLabelAssignmentMethodPPrivileged
	// SensitivityLabelAssignmentMethodPAuto is a pointer to SensitivityLabelAssignmentMethodVAuto
	SensitivityLabelAssignmentMethodPAuto = &_SensitivityLabelAssignmentMethodPAuto
	// SensitivityLabelAssignmentMethodPUnknownFutureValue is a pointer to SensitivityLabelAssignmentMethodVUnknownFutureValue
	SensitivityLabelAssignmentMethodPUnknownFutureValue = &_SensitivityLabelAssignmentMethodPUnknownFutureValue
)

var (
	_SensitivityLabelAssignmentMethodPStandard           = SensitivityLabelAssignmentMethodVStandard
	_SensitivityLabelAssignmentMethodPPrivileged         = SensitivityLabelAssignmentMethodVPrivileged
	_SensitivityLabelAssignmentMethodPAuto               = SensitivityLabelAssignmentMethodVAuto
	_SensitivityLabelAssignmentMethodPUnknownFutureValue = SensitivityLabelAssignmentMethodVUnknownFutureValue
)

// SensitivityLabelTarget undocumented
type SensitivityLabelTarget string

const (
	// SensitivityLabelTargetVEmail undocumented
	SensitivityLabelTargetVEmail SensitivityLabelTarget = "email"
	// SensitivityLabelTargetVSite undocumented
	SensitivityLabelTargetVSite SensitivityLabelTarget = "site"
	// SensitivityLabelTargetVUnifiedGroup undocumented
	SensitivityLabelTargetVUnifiedGroup SensitivityLabelTarget = "unifiedGroup"
	// SensitivityLabelTargetVUnknownFutureValue undocumented
	SensitivityLabelTargetVUnknownFutureValue SensitivityLabelTarget = "unknownFutureValue"
	// SensitivityLabelTargetVTeamwork undocumented
	SensitivityLabelTargetVTeamwork SensitivityLabelTarget = "teamwork"
)

var (
	// SensitivityLabelTargetPEmail is a pointer to SensitivityLabelTargetVEmail
	SensitivityLabelTargetPEmail = &_SensitivityLabelTargetPEmail
	// SensitivityLabelTargetPSite is a pointer to SensitivityLabelTargetVSite
	SensitivityLabelTargetPSite = &_SensitivityLabelTargetPSite
	// SensitivityLabelTargetPUnifiedGroup is a pointer to SensitivityLabelTargetVUnifiedGroup
	SensitivityLabelTargetPUnifiedGroup = &_SensitivityLabelTargetPUnifiedGroup
	// SensitivityLabelTargetPUnknownFutureValue is a pointer to SensitivityLabelTargetVUnknownFutureValue
	SensitivityLabelTargetPUnknownFutureValue = &_SensitivityLabelTargetPUnknownFutureValue
	// SensitivityLabelTargetPTeamwork is a pointer to SensitivityLabelTargetVTeamwork
	SensitivityLabelTargetPTeamwork = &_SensitivityLabelTargetPTeamwork
)

var (
	_SensitivityLabelTargetPEmail              = SensitivityLabelTargetVEmail
	_SensitivityLabelTargetPSite               = SensitivityLabelTargetVSite
	_SensitivityLabelTargetPUnifiedGroup       = SensitivityLabelTargetVUnifiedGroup
	_SensitivityLabelTargetPUnknownFutureValue = SensitivityLabelTargetVUnknownFutureValue
	_SensitivityLabelTargetPTeamwork           = SensitivityLabelTargetVTeamwork
)

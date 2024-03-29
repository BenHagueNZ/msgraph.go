// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ActionCapability undocumented
type ActionCapability string

const (
	// ActionCapabilityVEnabled undocumented
	ActionCapabilityVEnabled ActionCapability = "enabled"
	// ActionCapabilityVDisabled undocumented
	ActionCapabilityVDisabled ActionCapability = "disabled"
	// ActionCapabilityVUnknownFutureValue undocumented
	ActionCapabilityVUnknownFutureValue ActionCapability = "unknownFutureValue"
)

var (
	// ActionCapabilityPEnabled is a pointer to ActionCapabilityVEnabled
	ActionCapabilityPEnabled = &_ActionCapabilityPEnabled
	// ActionCapabilityPDisabled is a pointer to ActionCapabilityVDisabled
	ActionCapabilityPDisabled = &_ActionCapabilityPDisabled
	// ActionCapabilityPUnknownFutureValue is a pointer to ActionCapabilityVUnknownFutureValue
	ActionCapabilityPUnknownFutureValue = &_ActionCapabilityPUnknownFutureValue
)

var (
	_ActionCapabilityPEnabled            = ActionCapabilityVEnabled
	_ActionCapabilityPDisabled           = ActionCapabilityVDisabled
	_ActionCapabilityPUnknownFutureValue = ActionCapabilityVUnknownFutureValue
)

// ActionSource undocumented
type ActionSource string

const (
	// ActionSourceVManual undocumented
	ActionSourceVManual ActionSource = "manual"
	// ActionSourceVAutomatic undocumented
	ActionSourceVAutomatic ActionSource = "automatic"
	// ActionSourceVRecommended undocumented
	ActionSourceVRecommended ActionSource = "recommended"
	// ActionSourceVDefault undocumented
	ActionSourceVDefault ActionSource = "default"
)

var (
	// ActionSourcePManual is a pointer to ActionSourceVManual
	ActionSourcePManual = &_ActionSourcePManual
	// ActionSourcePAutomatic is a pointer to ActionSourceVAutomatic
	ActionSourcePAutomatic = &_ActionSourcePAutomatic
	// ActionSourcePRecommended is a pointer to ActionSourceVRecommended
	ActionSourcePRecommended = &_ActionSourcePRecommended
	// ActionSourcePDefault is a pointer to ActionSourceVDefault
	ActionSourcePDefault = &_ActionSourcePDefault
)

var (
	_ActionSourcePManual      = ActionSourceVManual
	_ActionSourcePAutomatic   = ActionSourceVAutomatic
	_ActionSourcePRecommended = ActionSourceVRecommended
	_ActionSourcePDefault     = ActionSourceVDefault
)

// ActionState undocumented
type ActionState string

const (
	// ActionStateVNone undocumented
	ActionStateVNone ActionState = "none"
	// ActionStateVPending undocumented
	ActionStateVPending ActionState = "pending"
	// ActionStateVCanceled undocumented
	ActionStateVCanceled ActionState = "canceled"
	// ActionStateVActive undocumented
	ActionStateVActive ActionState = "active"
	// ActionStateVDone undocumented
	ActionStateVDone ActionState = "done"
	// ActionStateVFailed undocumented
	ActionStateVFailed ActionState = "failed"
	// ActionStateVNotSupported undocumented
	ActionStateVNotSupported ActionState = "notSupported"
)

var (
	// ActionStatePNone is a pointer to ActionStateVNone
	ActionStatePNone = &_ActionStatePNone
	// ActionStatePPending is a pointer to ActionStateVPending
	ActionStatePPending = &_ActionStatePPending
	// ActionStatePCanceled is a pointer to ActionStateVCanceled
	ActionStatePCanceled = &_ActionStatePCanceled
	// ActionStatePActive is a pointer to ActionStateVActive
	ActionStatePActive = &_ActionStatePActive
	// ActionStatePDone is a pointer to ActionStateVDone
	ActionStatePDone = &_ActionStatePDone
	// ActionStatePFailed is a pointer to ActionStateVFailed
	ActionStatePFailed = &_ActionStatePFailed
	// ActionStatePNotSupported is a pointer to ActionStateVNotSupported
	ActionStatePNotSupported = &_ActionStatePNotSupported
)

var (
	_ActionStatePNone         = ActionStateVNone
	_ActionStatePPending      = ActionStateVPending
	_ActionStatePCanceled     = ActionStateVCanceled
	_ActionStatePActive       = ActionStateVActive
	_ActionStatePDone         = ActionStateVDone
	_ActionStatePFailed       = ActionStateVFailed
	_ActionStatePNotSupported = ActionStateVNotSupported
)

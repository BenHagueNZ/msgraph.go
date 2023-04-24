// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ExternalAudienceScope undocumented
type ExternalAudienceScope string

const (
	// ExternalAudienceScopeVNone undocumented
	ExternalAudienceScopeVNone ExternalAudienceScope = "none"
	// ExternalAudienceScopeVContactsOnly undocumented
	ExternalAudienceScopeVContactsOnly ExternalAudienceScope = "contactsOnly"
	// ExternalAudienceScopeVAll undocumented
	ExternalAudienceScopeVAll ExternalAudienceScope = "all"
)

var (
	// ExternalAudienceScopePNone is a pointer to ExternalAudienceScopeVNone
	ExternalAudienceScopePNone = &_ExternalAudienceScopePNone
	// ExternalAudienceScopePContactsOnly is a pointer to ExternalAudienceScopeVContactsOnly
	ExternalAudienceScopePContactsOnly = &_ExternalAudienceScopePContactsOnly
	// ExternalAudienceScopePAll is a pointer to ExternalAudienceScopeVAll
	ExternalAudienceScopePAll = &_ExternalAudienceScopePAll
)

var (
	_ExternalAudienceScopePNone         = ExternalAudienceScopeVNone
	_ExternalAudienceScopePContactsOnly = ExternalAudienceScopeVContactsOnly
	_ExternalAudienceScopePAll          = ExternalAudienceScopeVAll
)

// ExternalAuthenticationType undocumented
type ExternalAuthenticationType string

const (
	// ExternalAuthenticationTypeVPassthru undocumented
	ExternalAuthenticationTypeVPassthru ExternalAuthenticationType = "passthru"
	// ExternalAuthenticationTypeVAadPreAuthentication undocumented
	ExternalAuthenticationTypeVAadPreAuthentication ExternalAuthenticationType = "aadPreAuthentication"
)

var (
	// ExternalAuthenticationTypePPassthru is a pointer to ExternalAuthenticationTypeVPassthru
	ExternalAuthenticationTypePPassthru = &_ExternalAuthenticationTypePPassthru
	// ExternalAuthenticationTypePAadPreAuthentication is a pointer to ExternalAuthenticationTypeVAadPreAuthentication
	ExternalAuthenticationTypePAadPreAuthentication = &_ExternalAuthenticationTypePAadPreAuthentication
)

var (
	_ExternalAuthenticationTypePPassthru             = ExternalAuthenticationTypeVPassthru
	_ExternalAuthenticationTypePAadPreAuthentication = ExternalAuthenticationTypeVAadPreAuthentication
)

// ExternalEmailOtpState undocumented
type ExternalEmailOtpState string

const (
	// ExternalEmailOtpStateVDefault undocumented
	ExternalEmailOtpStateVDefault ExternalEmailOtpState = "default"
	// ExternalEmailOtpStateVEnabled undocumented
	ExternalEmailOtpStateVEnabled ExternalEmailOtpState = "enabled"
	// ExternalEmailOtpStateVDisabled undocumented
	ExternalEmailOtpStateVDisabled ExternalEmailOtpState = "disabled"
	// ExternalEmailOtpStateVUnknownFutureValue undocumented
	ExternalEmailOtpStateVUnknownFutureValue ExternalEmailOtpState = "unknownFutureValue"
)

var (
	// ExternalEmailOtpStatePDefault is a pointer to ExternalEmailOtpStateVDefault
	ExternalEmailOtpStatePDefault = &_ExternalEmailOtpStatePDefault
	// ExternalEmailOtpStatePEnabled is a pointer to ExternalEmailOtpStateVEnabled
	ExternalEmailOtpStatePEnabled = &_ExternalEmailOtpStatePEnabled
	// ExternalEmailOtpStatePDisabled is a pointer to ExternalEmailOtpStateVDisabled
	ExternalEmailOtpStatePDisabled = &_ExternalEmailOtpStatePDisabled
	// ExternalEmailOtpStatePUnknownFutureValue is a pointer to ExternalEmailOtpStateVUnknownFutureValue
	ExternalEmailOtpStatePUnknownFutureValue = &_ExternalEmailOtpStatePUnknownFutureValue
)

var (
	_ExternalEmailOtpStatePDefault            = ExternalEmailOtpStateVDefault
	_ExternalEmailOtpStatePEnabled            = ExternalEmailOtpStateVEnabled
	_ExternalEmailOtpStatePDisabled           = ExternalEmailOtpStateVDisabled
	_ExternalEmailOtpStatePUnknownFutureValue = ExternalEmailOtpStateVUnknownFutureValue
)

// ExternalItemContentType undocumented
type ExternalItemContentType string

const (
	// ExternalItemContentTypeVText undocumented
	ExternalItemContentTypeVText ExternalItemContentType = "text"
	// ExternalItemContentTypeVHTML undocumented
	ExternalItemContentTypeVHTML ExternalItemContentType = "html"
	// ExternalItemContentTypeVUnknownFutureValue undocumented
	ExternalItemContentTypeVUnknownFutureValue ExternalItemContentType = "unknownFutureValue"
)

var (
	// ExternalItemContentTypePText is a pointer to ExternalItemContentTypeVText
	ExternalItemContentTypePText = &_ExternalItemContentTypePText
	// ExternalItemContentTypePHTML is a pointer to ExternalItemContentTypeVHTML
	ExternalItemContentTypePHTML = &_ExternalItemContentTypePHTML
	// ExternalItemContentTypePUnknownFutureValue is a pointer to ExternalItemContentTypeVUnknownFutureValue
	ExternalItemContentTypePUnknownFutureValue = &_ExternalItemContentTypePUnknownFutureValue
)

var (
	_ExternalItemContentTypePText               = ExternalItemContentTypeVText
	_ExternalItemContentTypePHTML               = ExternalItemContentTypeVHTML
	_ExternalItemContentTypePUnknownFutureValue = ExternalItemContentTypeVUnknownFutureValue
)

// ExternalConnectorsAccessType undocumented
type ExternalConnectorsAccessType string

const (
	// ExternalConnectorsAccessTypeVGrant undocumented
	ExternalConnectorsAccessTypeVGrant ExternalConnectorsAccessType = "grant"
	// ExternalConnectorsAccessTypeVDeny undocumented
	ExternalConnectorsAccessTypeVDeny ExternalConnectorsAccessType = "deny"
	// ExternalConnectorsAccessTypeVUnknownFutureValue undocumented
	ExternalConnectorsAccessTypeVUnknownFutureValue ExternalConnectorsAccessType = "unknownFutureValue"
)

var (
	// ExternalConnectorsAccessTypePGrant is a pointer to ExternalConnectorsAccessTypeVGrant
	ExternalConnectorsAccessTypePGrant = &_ExternalConnectorsAccessTypePGrant
	// ExternalConnectorsAccessTypePDeny is a pointer to ExternalConnectorsAccessTypeVDeny
	ExternalConnectorsAccessTypePDeny = &_ExternalConnectorsAccessTypePDeny
	// ExternalConnectorsAccessTypePUnknownFutureValue is a pointer to ExternalConnectorsAccessTypeVUnknownFutureValue
	ExternalConnectorsAccessTypePUnknownFutureValue = &_ExternalConnectorsAccessTypePUnknownFutureValue
)

var (
	_ExternalConnectorsAccessTypePGrant              = ExternalConnectorsAccessTypeVGrant
	_ExternalConnectorsAccessTypePDeny               = ExternalConnectorsAccessTypeVDeny
	_ExternalConnectorsAccessTypePUnknownFutureValue = ExternalConnectorsAccessTypeVUnknownFutureValue
)

// ExternalConnectorsACLType undocumented
type ExternalConnectorsACLType string

const (
	// ExternalConnectorsACLTypeVUser undocumented
	ExternalConnectorsACLTypeVUser ExternalConnectorsACLType = "user"
	// ExternalConnectorsACLTypeVGroup undocumented
	ExternalConnectorsACLTypeVGroup ExternalConnectorsACLType = "group"
	// ExternalConnectorsACLTypeVEveryone undocumented
	ExternalConnectorsACLTypeVEveryone ExternalConnectorsACLType = "everyone"
	// ExternalConnectorsACLTypeVEveryoneExceptGuests undocumented
	ExternalConnectorsACLTypeVEveryoneExceptGuests ExternalConnectorsACLType = "everyoneExceptGuests"
	// ExternalConnectorsACLTypeVExternalGroup undocumented
	ExternalConnectorsACLTypeVExternalGroup ExternalConnectorsACLType = "externalGroup"
	// ExternalConnectorsACLTypeVUnknownFutureValue undocumented
	ExternalConnectorsACLTypeVUnknownFutureValue ExternalConnectorsACLType = "unknownFutureValue"
)

var (
	// ExternalConnectorsACLTypePUser is a pointer to ExternalConnectorsACLTypeVUser
	ExternalConnectorsACLTypePUser = &_ExternalConnectorsACLTypePUser
	// ExternalConnectorsACLTypePGroup is a pointer to ExternalConnectorsACLTypeVGroup
	ExternalConnectorsACLTypePGroup = &_ExternalConnectorsACLTypePGroup
	// ExternalConnectorsACLTypePEveryone is a pointer to ExternalConnectorsACLTypeVEveryone
	ExternalConnectorsACLTypePEveryone = &_ExternalConnectorsACLTypePEveryone
	// ExternalConnectorsACLTypePEveryoneExceptGuests is a pointer to ExternalConnectorsACLTypeVEveryoneExceptGuests
	ExternalConnectorsACLTypePEveryoneExceptGuests = &_ExternalConnectorsACLTypePEveryoneExceptGuests
	// ExternalConnectorsACLTypePExternalGroup is a pointer to ExternalConnectorsACLTypeVExternalGroup
	ExternalConnectorsACLTypePExternalGroup = &_ExternalConnectorsACLTypePExternalGroup
	// ExternalConnectorsACLTypePUnknownFutureValue is a pointer to ExternalConnectorsACLTypeVUnknownFutureValue
	ExternalConnectorsACLTypePUnknownFutureValue = &_ExternalConnectorsACLTypePUnknownFutureValue
)

var (
	_ExternalConnectorsACLTypePUser                 = ExternalConnectorsACLTypeVUser
	_ExternalConnectorsACLTypePGroup                = ExternalConnectorsACLTypeVGroup
	_ExternalConnectorsACLTypePEveryone             = ExternalConnectorsACLTypeVEveryone
	_ExternalConnectorsACLTypePEveryoneExceptGuests = ExternalConnectorsACLTypeVEveryoneExceptGuests
	_ExternalConnectorsACLTypePExternalGroup        = ExternalConnectorsACLTypeVExternalGroup
	_ExternalConnectorsACLTypePUnknownFutureValue   = ExternalConnectorsACLTypeVUnknownFutureValue
)

// ExternalConnectorsConnectionOperationStatus undocumented
type ExternalConnectorsConnectionOperationStatus string

const (
	// ExternalConnectorsConnectionOperationStatusVUnspecified undocumented
	ExternalConnectorsConnectionOperationStatusVUnspecified ExternalConnectorsConnectionOperationStatus = "unspecified"
	// ExternalConnectorsConnectionOperationStatusVInprogress undocumented
	ExternalConnectorsConnectionOperationStatusVInprogress ExternalConnectorsConnectionOperationStatus = "inprogress"
	// ExternalConnectorsConnectionOperationStatusVCompleted undocumented
	ExternalConnectorsConnectionOperationStatusVCompleted ExternalConnectorsConnectionOperationStatus = "completed"
	// ExternalConnectorsConnectionOperationStatusVFailed undocumented
	ExternalConnectorsConnectionOperationStatusVFailed ExternalConnectorsConnectionOperationStatus = "failed"
	// ExternalConnectorsConnectionOperationStatusVUnknownFutureValue undocumented
	ExternalConnectorsConnectionOperationStatusVUnknownFutureValue ExternalConnectorsConnectionOperationStatus = "unknownFutureValue"
)

var (
	// ExternalConnectorsConnectionOperationStatusPUnspecified is a pointer to ExternalConnectorsConnectionOperationStatusVUnspecified
	ExternalConnectorsConnectionOperationStatusPUnspecified = &_ExternalConnectorsConnectionOperationStatusPUnspecified
	// ExternalConnectorsConnectionOperationStatusPInprogress is a pointer to ExternalConnectorsConnectionOperationStatusVInprogress
	ExternalConnectorsConnectionOperationStatusPInprogress = &_ExternalConnectorsConnectionOperationStatusPInprogress
	// ExternalConnectorsConnectionOperationStatusPCompleted is a pointer to ExternalConnectorsConnectionOperationStatusVCompleted
	ExternalConnectorsConnectionOperationStatusPCompleted = &_ExternalConnectorsConnectionOperationStatusPCompleted
	// ExternalConnectorsConnectionOperationStatusPFailed is a pointer to ExternalConnectorsConnectionOperationStatusVFailed
	ExternalConnectorsConnectionOperationStatusPFailed = &_ExternalConnectorsConnectionOperationStatusPFailed
	// ExternalConnectorsConnectionOperationStatusPUnknownFutureValue is a pointer to ExternalConnectorsConnectionOperationStatusVUnknownFutureValue
	ExternalConnectorsConnectionOperationStatusPUnknownFutureValue = &_ExternalConnectorsConnectionOperationStatusPUnknownFutureValue
)

var (
	_ExternalConnectorsConnectionOperationStatusPUnspecified        = ExternalConnectorsConnectionOperationStatusVUnspecified
	_ExternalConnectorsConnectionOperationStatusPInprogress         = ExternalConnectorsConnectionOperationStatusVInprogress
	_ExternalConnectorsConnectionOperationStatusPCompleted          = ExternalConnectorsConnectionOperationStatusVCompleted
	_ExternalConnectorsConnectionOperationStatusPFailed             = ExternalConnectorsConnectionOperationStatusVFailed
	_ExternalConnectorsConnectionOperationStatusPUnknownFutureValue = ExternalConnectorsConnectionOperationStatusVUnknownFutureValue
)

// ExternalConnectorsConnectionState undocumented
type ExternalConnectorsConnectionState string

const (
	// ExternalConnectorsConnectionStateVDraft undocumented
	ExternalConnectorsConnectionStateVDraft ExternalConnectorsConnectionState = "draft"
	// ExternalConnectorsConnectionStateVReady undocumented
	ExternalConnectorsConnectionStateVReady ExternalConnectorsConnectionState = "ready"
	// ExternalConnectorsConnectionStateVObsolete undocumented
	ExternalConnectorsConnectionStateVObsolete ExternalConnectorsConnectionState = "obsolete"
	// ExternalConnectorsConnectionStateVLimitExceeded undocumented
	ExternalConnectorsConnectionStateVLimitExceeded ExternalConnectorsConnectionState = "limitExceeded"
	// ExternalConnectorsConnectionStateVUnknownFutureValue undocumented
	ExternalConnectorsConnectionStateVUnknownFutureValue ExternalConnectorsConnectionState = "unknownFutureValue"
)

var (
	// ExternalConnectorsConnectionStatePDraft is a pointer to ExternalConnectorsConnectionStateVDraft
	ExternalConnectorsConnectionStatePDraft = &_ExternalConnectorsConnectionStatePDraft
	// ExternalConnectorsConnectionStatePReady is a pointer to ExternalConnectorsConnectionStateVReady
	ExternalConnectorsConnectionStatePReady = &_ExternalConnectorsConnectionStatePReady
	// ExternalConnectorsConnectionStatePObsolete is a pointer to ExternalConnectorsConnectionStateVObsolete
	ExternalConnectorsConnectionStatePObsolete = &_ExternalConnectorsConnectionStatePObsolete
	// ExternalConnectorsConnectionStatePLimitExceeded is a pointer to ExternalConnectorsConnectionStateVLimitExceeded
	ExternalConnectorsConnectionStatePLimitExceeded = &_ExternalConnectorsConnectionStatePLimitExceeded
	// ExternalConnectorsConnectionStatePUnknownFutureValue is a pointer to ExternalConnectorsConnectionStateVUnknownFutureValue
	ExternalConnectorsConnectionStatePUnknownFutureValue = &_ExternalConnectorsConnectionStatePUnknownFutureValue
)

var (
	_ExternalConnectorsConnectionStatePDraft              = ExternalConnectorsConnectionStateVDraft
	_ExternalConnectorsConnectionStatePReady              = ExternalConnectorsConnectionStateVReady
	_ExternalConnectorsConnectionStatePObsolete           = ExternalConnectorsConnectionStateVObsolete
	_ExternalConnectorsConnectionStatePLimitExceeded      = ExternalConnectorsConnectionStateVLimitExceeded
	_ExternalConnectorsConnectionStatePUnknownFutureValue = ExternalConnectorsConnectionStateVUnknownFutureValue
)

// ExternalConnectorsContentExperienceType undocumented
type ExternalConnectorsContentExperienceType string

const (
	// ExternalConnectorsContentExperienceTypeVSearch undocumented
	ExternalConnectorsContentExperienceTypeVSearch ExternalConnectorsContentExperienceType = "search"
	// ExternalConnectorsContentExperienceTypeVCompliance undocumented
	ExternalConnectorsContentExperienceTypeVCompliance ExternalConnectorsContentExperienceType = "compliance"
	// ExternalConnectorsContentExperienceTypeVUnknownFutureValue undocumented
	ExternalConnectorsContentExperienceTypeVUnknownFutureValue ExternalConnectorsContentExperienceType = "unknownFutureValue"
)

var (
	// ExternalConnectorsContentExperienceTypePSearch is a pointer to ExternalConnectorsContentExperienceTypeVSearch
	ExternalConnectorsContentExperienceTypePSearch = &_ExternalConnectorsContentExperienceTypePSearch
	// ExternalConnectorsContentExperienceTypePCompliance is a pointer to ExternalConnectorsContentExperienceTypeVCompliance
	ExternalConnectorsContentExperienceTypePCompliance = &_ExternalConnectorsContentExperienceTypePCompliance
	// ExternalConnectorsContentExperienceTypePUnknownFutureValue is a pointer to ExternalConnectorsContentExperienceTypeVUnknownFutureValue
	ExternalConnectorsContentExperienceTypePUnknownFutureValue = &_ExternalConnectorsContentExperienceTypePUnknownFutureValue
)

var (
	_ExternalConnectorsContentExperienceTypePSearch             = ExternalConnectorsContentExperienceTypeVSearch
	_ExternalConnectorsContentExperienceTypePCompliance         = ExternalConnectorsContentExperienceTypeVCompliance
	_ExternalConnectorsContentExperienceTypePUnknownFutureValue = ExternalConnectorsContentExperienceTypeVUnknownFutureValue
)

// ExternalConnectorsExternalActivityType undocumented
type ExternalConnectorsExternalActivityType string

const (
	// ExternalConnectorsExternalActivityTypeVViewed undocumented
	ExternalConnectorsExternalActivityTypeVViewed ExternalConnectorsExternalActivityType = "viewed"
	// ExternalConnectorsExternalActivityTypeVModified undocumented
	ExternalConnectorsExternalActivityTypeVModified ExternalConnectorsExternalActivityType = "modified"
	// ExternalConnectorsExternalActivityTypeVCreated undocumented
	ExternalConnectorsExternalActivityTypeVCreated ExternalConnectorsExternalActivityType = "created"
	// ExternalConnectorsExternalActivityTypeVCommented undocumented
	ExternalConnectorsExternalActivityTypeVCommented ExternalConnectorsExternalActivityType = "commented"
	// ExternalConnectorsExternalActivityTypeVUnknownFutureValue undocumented
	ExternalConnectorsExternalActivityTypeVUnknownFutureValue ExternalConnectorsExternalActivityType = "unknownFutureValue"
)

var (
	// ExternalConnectorsExternalActivityTypePViewed is a pointer to ExternalConnectorsExternalActivityTypeVViewed
	ExternalConnectorsExternalActivityTypePViewed = &_ExternalConnectorsExternalActivityTypePViewed
	// ExternalConnectorsExternalActivityTypePModified is a pointer to ExternalConnectorsExternalActivityTypeVModified
	ExternalConnectorsExternalActivityTypePModified = &_ExternalConnectorsExternalActivityTypePModified
	// ExternalConnectorsExternalActivityTypePCreated is a pointer to ExternalConnectorsExternalActivityTypeVCreated
	ExternalConnectorsExternalActivityTypePCreated = &_ExternalConnectorsExternalActivityTypePCreated
	// ExternalConnectorsExternalActivityTypePCommented is a pointer to ExternalConnectorsExternalActivityTypeVCommented
	ExternalConnectorsExternalActivityTypePCommented = &_ExternalConnectorsExternalActivityTypePCommented
	// ExternalConnectorsExternalActivityTypePUnknownFutureValue is a pointer to ExternalConnectorsExternalActivityTypeVUnknownFutureValue
	ExternalConnectorsExternalActivityTypePUnknownFutureValue = &_ExternalConnectorsExternalActivityTypePUnknownFutureValue
)

var (
	_ExternalConnectorsExternalActivityTypePViewed             = ExternalConnectorsExternalActivityTypeVViewed
	_ExternalConnectorsExternalActivityTypePModified           = ExternalConnectorsExternalActivityTypeVModified
	_ExternalConnectorsExternalActivityTypePCreated            = ExternalConnectorsExternalActivityTypeVCreated
	_ExternalConnectorsExternalActivityTypePCommented          = ExternalConnectorsExternalActivityTypeVCommented
	_ExternalConnectorsExternalActivityTypePUnknownFutureValue = ExternalConnectorsExternalActivityTypeVUnknownFutureValue
)

// ExternalConnectorsExternalItemContentType undocumented
type ExternalConnectorsExternalItemContentType string

const (
	// ExternalConnectorsExternalItemContentTypeVText undocumented
	ExternalConnectorsExternalItemContentTypeVText ExternalConnectorsExternalItemContentType = "text"
	// ExternalConnectorsExternalItemContentTypeVHTML undocumented
	ExternalConnectorsExternalItemContentTypeVHTML ExternalConnectorsExternalItemContentType = "html"
	// ExternalConnectorsExternalItemContentTypeVUnknownFutureValue undocumented
	ExternalConnectorsExternalItemContentTypeVUnknownFutureValue ExternalConnectorsExternalItemContentType = "unknownFutureValue"
)

var (
	// ExternalConnectorsExternalItemContentTypePText is a pointer to ExternalConnectorsExternalItemContentTypeVText
	ExternalConnectorsExternalItemContentTypePText = &_ExternalConnectorsExternalItemContentTypePText
	// ExternalConnectorsExternalItemContentTypePHTML is a pointer to ExternalConnectorsExternalItemContentTypeVHTML
	ExternalConnectorsExternalItemContentTypePHTML = &_ExternalConnectorsExternalItemContentTypePHTML
	// ExternalConnectorsExternalItemContentTypePUnknownFutureValue is a pointer to ExternalConnectorsExternalItemContentTypeVUnknownFutureValue
	ExternalConnectorsExternalItemContentTypePUnknownFutureValue = &_ExternalConnectorsExternalItemContentTypePUnknownFutureValue
)

var (
	_ExternalConnectorsExternalItemContentTypePText               = ExternalConnectorsExternalItemContentTypeVText
	_ExternalConnectorsExternalItemContentTypePHTML               = ExternalConnectorsExternalItemContentTypeVHTML
	_ExternalConnectorsExternalItemContentTypePUnknownFutureValue = ExternalConnectorsExternalItemContentTypeVUnknownFutureValue
)

// ExternalConnectorsIdentitySourceType undocumented
type ExternalConnectorsIdentitySourceType string

const (
	// ExternalConnectorsIdentitySourceTypeVAzureActiveDirectory undocumented
	ExternalConnectorsIdentitySourceTypeVAzureActiveDirectory ExternalConnectorsIdentitySourceType = "azureActiveDirectory"
	// ExternalConnectorsIdentitySourceTypeVExternal undocumented
	ExternalConnectorsIdentitySourceTypeVExternal ExternalConnectorsIdentitySourceType = "external"
	// ExternalConnectorsIdentitySourceTypeVUnknownFutureValue undocumented
	ExternalConnectorsIdentitySourceTypeVUnknownFutureValue ExternalConnectorsIdentitySourceType = "unknownFutureValue"
)

var (
	// ExternalConnectorsIdentitySourceTypePAzureActiveDirectory is a pointer to ExternalConnectorsIdentitySourceTypeVAzureActiveDirectory
	ExternalConnectorsIdentitySourceTypePAzureActiveDirectory = &_ExternalConnectorsIdentitySourceTypePAzureActiveDirectory
	// ExternalConnectorsIdentitySourceTypePExternal is a pointer to ExternalConnectorsIdentitySourceTypeVExternal
	ExternalConnectorsIdentitySourceTypePExternal = &_ExternalConnectorsIdentitySourceTypePExternal
	// ExternalConnectorsIdentitySourceTypePUnknownFutureValue is a pointer to ExternalConnectorsIdentitySourceTypeVUnknownFutureValue
	ExternalConnectorsIdentitySourceTypePUnknownFutureValue = &_ExternalConnectorsIdentitySourceTypePUnknownFutureValue
)

var (
	_ExternalConnectorsIdentitySourceTypePAzureActiveDirectory = ExternalConnectorsIdentitySourceTypeVAzureActiveDirectory
	_ExternalConnectorsIdentitySourceTypePExternal             = ExternalConnectorsIdentitySourceTypeVExternal
	_ExternalConnectorsIdentitySourceTypePUnknownFutureValue   = ExternalConnectorsIdentitySourceTypeVUnknownFutureValue
)

// ExternalConnectorsIdentityType undocumented
type ExternalConnectorsIdentityType string

const (
	// ExternalConnectorsIdentityTypeVUser undocumented
	ExternalConnectorsIdentityTypeVUser ExternalConnectorsIdentityType = "user"
	// ExternalConnectorsIdentityTypeVGroup undocumented
	ExternalConnectorsIdentityTypeVGroup ExternalConnectorsIdentityType = "group"
	// ExternalConnectorsIdentityTypeVExternalGroup undocumented
	ExternalConnectorsIdentityTypeVExternalGroup ExternalConnectorsIdentityType = "externalGroup"
	// ExternalConnectorsIdentityTypeVUnknownFutureValue undocumented
	ExternalConnectorsIdentityTypeVUnknownFutureValue ExternalConnectorsIdentityType = "unknownFutureValue"
)

var (
	// ExternalConnectorsIdentityTypePUser is a pointer to ExternalConnectorsIdentityTypeVUser
	ExternalConnectorsIdentityTypePUser = &_ExternalConnectorsIdentityTypePUser
	// ExternalConnectorsIdentityTypePGroup is a pointer to ExternalConnectorsIdentityTypeVGroup
	ExternalConnectorsIdentityTypePGroup = &_ExternalConnectorsIdentityTypePGroup
	// ExternalConnectorsIdentityTypePExternalGroup is a pointer to ExternalConnectorsIdentityTypeVExternalGroup
	ExternalConnectorsIdentityTypePExternalGroup = &_ExternalConnectorsIdentityTypePExternalGroup
	// ExternalConnectorsIdentityTypePUnknownFutureValue is a pointer to ExternalConnectorsIdentityTypeVUnknownFutureValue
	ExternalConnectorsIdentityTypePUnknownFutureValue = &_ExternalConnectorsIdentityTypePUnknownFutureValue
)

var (
	_ExternalConnectorsIdentityTypePUser               = ExternalConnectorsIdentityTypeVUser
	_ExternalConnectorsIdentityTypePGroup              = ExternalConnectorsIdentityTypeVGroup
	_ExternalConnectorsIdentityTypePExternalGroup      = ExternalConnectorsIdentityTypeVExternalGroup
	_ExternalConnectorsIdentityTypePUnknownFutureValue = ExternalConnectorsIdentityTypeVUnknownFutureValue
)

// ExternalConnectorsImportanceScore undocumented
type ExternalConnectorsImportanceScore string

const (
	// ExternalConnectorsImportanceScoreVLow undocumented
	ExternalConnectorsImportanceScoreVLow ExternalConnectorsImportanceScore = "low"
	// ExternalConnectorsImportanceScoreVMedium undocumented
	ExternalConnectorsImportanceScoreVMedium ExternalConnectorsImportanceScore = "medium"
	// ExternalConnectorsImportanceScoreVHigh undocumented
	ExternalConnectorsImportanceScoreVHigh ExternalConnectorsImportanceScore = "high"
	// ExternalConnectorsImportanceScoreVVeryHigh undocumented
	ExternalConnectorsImportanceScoreVVeryHigh ExternalConnectorsImportanceScore = "veryHigh"
	// ExternalConnectorsImportanceScoreVUnknownFutureValue undocumented
	ExternalConnectorsImportanceScoreVUnknownFutureValue ExternalConnectorsImportanceScore = "unknownFutureValue"
)

var (
	// ExternalConnectorsImportanceScorePLow is a pointer to ExternalConnectorsImportanceScoreVLow
	ExternalConnectorsImportanceScorePLow = &_ExternalConnectorsImportanceScorePLow
	// ExternalConnectorsImportanceScorePMedium is a pointer to ExternalConnectorsImportanceScoreVMedium
	ExternalConnectorsImportanceScorePMedium = &_ExternalConnectorsImportanceScorePMedium
	// ExternalConnectorsImportanceScorePHigh is a pointer to ExternalConnectorsImportanceScoreVHigh
	ExternalConnectorsImportanceScorePHigh = &_ExternalConnectorsImportanceScorePHigh
	// ExternalConnectorsImportanceScorePVeryHigh is a pointer to ExternalConnectorsImportanceScoreVVeryHigh
	ExternalConnectorsImportanceScorePVeryHigh = &_ExternalConnectorsImportanceScorePVeryHigh
	// ExternalConnectorsImportanceScorePUnknownFutureValue is a pointer to ExternalConnectorsImportanceScoreVUnknownFutureValue
	ExternalConnectorsImportanceScorePUnknownFutureValue = &_ExternalConnectorsImportanceScorePUnknownFutureValue
)

var (
	_ExternalConnectorsImportanceScorePLow                = ExternalConnectorsImportanceScoreVLow
	_ExternalConnectorsImportanceScorePMedium             = ExternalConnectorsImportanceScoreVMedium
	_ExternalConnectorsImportanceScorePHigh               = ExternalConnectorsImportanceScoreVHigh
	_ExternalConnectorsImportanceScorePVeryHigh           = ExternalConnectorsImportanceScoreVVeryHigh
	_ExternalConnectorsImportanceScorePUnknownFutureValue = ExternalConnectorsImportanceScoreVUnknownFutureValue
)

// ExternalConnectorsLabel undocumented
type ExternalConnectorsLabel string

const (
	// ExternalConnectorsLabelVTitle undocumented
	ExternalConnectorsLabelVTitle ExternalConnectorsLabel = "title"
	// ExternalConnectorsLabelVURL undocumented
	ExternalConnectorsLabelVURL ExternalConnectorsLabel = "url"
	// ExternalConnectorsLabelVCreatedBy undocumented
	ExternalConnectorsLabelVCreatedBy ExternalConnectorsLabel = "createdBy"
	// ExternalConnectorsLabelVLastModifiedBy undocumented
	ExternalConnectorsLabelVLastModifiedBy ExternalConnectorsLabel = "lastModifiedBy"
	// ExternalConnectorsLabelVAuthors undocumented
	ExternalConnectorsLabelVAuthors ExternalConnectorsLabel = "authors"
	// ExternalConnectorsLabelVCreatedDateTime undocumented
	ExternalConnectorsLabelVCreatedDateTime ExternalConnectorsLabel = "createdDateTime"
	// ExternalConnectorsLabelVLastModifiedDateTime undocumented
	ExternalConnectorsLabelVLastModifiedDateTime ExternalConnectorsLabel = "lastModifiedDateTime"
	// ExternalConnectorsLabelVFileName undocumented
	ExternalConnectorsLabelVFileName ExternalConnectorsLabel = "fileName"
	// ExternalConnectorsLabelVFileExtension undocumented
	ExternalConnectorsLabelVFileExtension ExternalConnectorsLabel = "fileExtension"
	// ExternalConnectorsLabelVUnknownFutureValue undocumented
	ExternalConnectorsLabelVUnknownFutureValue ExternalConnectorsLabel = "unknownFutureValue"
	// ExternalConnectorsLabelVIconURL undocumented
	ExternalConnectorsLabelVIconURL ExternalConnectorsLabel = "iconUrl"
	// ExternalConnectorsLabelVContainerName undocumented
	ExternalConnectorsLabelVContainerName ExternalConnectorsLabel = "containerName"
	// ExternalConnectorsLabelVContainerURL undocumented
	ExternalConnectorsLabelVContainerURL ExternalConnectorsLabel = "containerUrl"
)

var (
	// ExternalConnectorsLabelPTitle is a pointer to ExternalConnectorsLabelVTitle
	ExternalConnectorsLabelPTitle = &_ExternalConnectorsLabelPTitle
	// ExternalConnectorsLabelPURL is a pointer to ExternalConnectorsLabelVURL
	ExternalConnectorsLabelPURL = &_ExternalConnectorsLabelPURL
	// ExternalConnectorsLabelPCreatedBy is a pointer to ExternalConnectorsLabelVCreatedBy
	ExternalConnectorsLabelPCreatedBy = &_ExternalConnectorsLabelPCreatedBy
	// ExternalConnectorsLabelPLastModifiedBy is a pointer to ExternalConnectorsLabelVLastModifiedBy
	ExternalConnectorsLabelPLastModifiedBy = &_ExternalConnectorsLabelPLastModifiedBy
	// ExternalConnectorsLabelPAuthors is a pointer to ExternalConnectorsLabelVAuthors
	ExternalConnectorsLabelPAuthors = &_ExternalConnectorsLabelPAuthors
	// ExternalConnectorsLabelPCreatedDateTime is a pointer to ExternalConnectorsLabelVCreatedDateTime
	ExternalConnectorsLabelPCreatedDateTime = &_ExternalConnectorsLabelPCreatedDateTime
	// ExternalConnectorsLabelPLastModifiedDateTime is a pointer to ExternalConnectorsLabelVLastModifiedDateTime
	ExternalConnectorsLabelPLastModifiedDateTime = &_ExternalConnectorsLabelPLastModifiedDateTime
	// ExternalConnectorsLabelPFileName is a pointer to ExternalConnectorsLabelVFileName
	ExternalConnectorsLabelPFileName = &_ExternalConnectorsLabelPFileName
	// ExternalConnectorsLabelPFileExtension is a pointer to ExternalConnectorsLabelVFileExtension
	ExternalConnectorsLabelPFileExtension = &_ExternalConnectorsLabelPFileExtension
	// ExternalConnectorsLabelPUnknownFutureValue is a pointer to ExternalConnectorsLabelVUnknownFutureValue
	ExternalConnectorsLabelPUnknownFutureValue = &_ExternalConnectorsLabelPUnknownFutureValue
	// ExternalConnectorsLabelPIconURL is a pointer to ExternalConnectorsLabelVIconURL
	ExternalConnectorsLabelPIconURL = &_ExternalConnectorsLabelPIconURL
	// ExternalConnectorsLabelPContainerName is a pointer to ExternalConnectorsLabelVContainerName
	ExternalConnectorsLabelPContainerName = &_ExternalConnectorsLabelPContainerName
	// ExternalConnectorsLabelPContainerURL is a pointer to ExternalConnectorsLabelVContainerURL
	ExternalConnectorsLabelPContainerURL = &_ExternalConnectorsLabelPContainerURL
)

var (
	_ExternalConnectorsLabelPTitle                = ExternalConnectorsLabelVTitle
	_ExternalConnectorsLabelPURL                  = ExternalConnectorsLabelVURL
	_ExternalConnectorsLabelPCreatedBy            = ExternalConnectorsLabelVCreatedBy
	_ExternalConnectorsLabelPLastModifiedBy       = ExternalConnectorsLabelVLastModifiedBy
	_ExternalConnectorsLabelPAuthors              = ExternalConnectorsLabelVAuthors
	_ExternalConnectorsLabelPCreatedDateTime      = ExternalConnectorsLabelVCreatedDateTime
	_ExternalConnectorsLabelPLastModifiedDateTime = ExternalConnectorsLabelVLastModifiedDateTime
	_ExternalConnectorsLabelPFileName             = ExternalConnectorsLabelVFileName
	_ExternalConnectorsLabelPFileExtension        = ExternalConnectorsLabelVFileExtension
	_ExternalConnectorsLabelPUnknownFutureValue   = ExternalConnectorsLabelVUnknownFutureValue
	_ExternalConnectorsLabelPIconURL              = ExternalConnectorsLabelVIconURL
	_ExternalConnectorsLabelPContainerName        = ExternalConnectorsLabelVContainerName
	_ExternalConnectorsLabelPContainerURL         = ExternalConnectorsLabelVContainerURL
)

// ExternalConnectorsPropertyType undocumented
type ExternalConnectorsPropertyType string

const (
	// ExternalConnectorsPropertyTypeVString undocumented
	ExternalConnectorsPropertyTypeVString ExternalConnectorsPropertyType = "string"
	// ExternalConnectorsPropertyTypeVInt64 undocumented
	ExternalConnectorsPropertyTypeVInt64 ExternalConnectorsPropertyType = "int64"
	// ExternalConnectorsPropertyTypeVDouble undocumented
	ExternalConnectorsPropertyTypeVDouble ExternalConnectorsPropertyType = "double"
	// ExternalConnectorsPropertyTypeVDateTime undocumented
	ExternalConnectorsPropertyTypeVDateTime ExternalConnectorsPropertyType = "dateTime"
	// ExternalConnectorsPropertyTypeVBoolean undocumented
	ExternalConnectorsPropertyTypeVBoolean ExternalConnectorsPropertyType = "boolean"
	// ExternalConnectorsPropertyTypeVStringCollection undocumented
	ExternalConnectorsPropertyTypeVStringCollection ExternalConnectorsPropertyType = "stringCollection"
	// ExternalConnectorsPropertyTypeVInt64Collection undocumented
	ExternalConnectorsPropertyTypeVInt64Collection ExternalConnectorsPropertyType = "int64Collection"
	// ExternalConnectorsPropertyTypeVDoubleCollection undocumented
	ExternalConnectorsPropertyTypeVDoubleCollection ExternalConnectorsPropertyType = "doubleCollection"
	// ExternalConnectorsPropertyTypeVDateTimeCollection undocumented
	ExternalConnectorsPropertyTypeVDateTimeCollection ExternalConnectorsPropertyType = "dateTimeCollection"
	// ExternalConnectorsPropertyTypeVUnknownFutureValue undocumented
	ExternalConnectorsPropertyTypeVUnknownFutureValue ExternalConnectorsPropertyType = "unknownFutureValue"
)

var (
	// ExternalConnectorsPropertyTypePString is a pointer to ExternalConnectorsPropertyTypeVString
	ExternalConnectorsPropertyTypePString = &_ExternalConnectorsPropertyTypePString
	// ExternalConnectorsPropertyTypePInt64 is a pointer to ExternalConnectorsPropertyTypeVInt64
	ExternalConnectorsPropertyTypePInt64 = &_ExternalConnectorsPropertyTypePInt64
	// ExternalConnectorsPropertyTypePDouble is a pointer to ExternalConnectorsPropertyTypeVDouble
	ExternalConnectorsPropertyTypePDouble = &_ExternalConnectorsPropertyTypePDouble
	// ExternalConnectorsPropertyTypePDateTime is a pointer to ExternalConnectorsPropertyTypeVDateTime
	ExternalConnectorsPropertyTypePDateTime = &_ExternalConnectorsPropertyTypePDateTime
	// ExternalConnectorsPropertyTypePBoolean is a pointer to ExternalConnectorsPropertyTypeVBoolean
	ExternalConnectorsPropertyTypePBoolean = &_ExternalConnectorsPropertyTypePBoolean
	// ExternalConnectorsPropertyTypePStringCollection is a pointer to ExternalConnectorsPropertyTypeVStringCollection
	ExternalConnectorsPropertyTypePStringCollection = &_ExternalConnectorsPropertyTypePStringCollection
	// ExternalConnectorsPropertyTypePInt64Collection is a pointer to ExternalConnectorsPropertyTypeVInt64Collection
	ExternalConnectorsPropertyTypePInt64Collection = &_ExternalConnectorsPropertyTypePInt64Collection
	// ExternalConnectorsPropertyTypePDoubleCollection is a pointer to ExternalConnectorsPropertyTypeVDoubleCollection
	ExternalConnectorsPropertyTypePDoubleCollection = &_ExternalConnectorsPropertyTypePDoubleCollection
	// ExternalConnectorsPropertyTypePDateTimeCollection is a pointer to ExternalConnectorsPropertyTypeVDateTimeCollection
	ExternalConnectorsPropertyTypePDateTimeCollection = &_ExternalConnectorsPropertyTypePDateTimeCollection
	// ExternalConnectorsPropertyTypePUnknownFutureValue is a pointer to ExternalConnectorsPropertyTypeVUnknownFutureValue
	ExternalConnectorsPropertyTypePUnknownFutureValue = &_ExternalConnectorsPropertyTypePUnknownFutureValue
)

var (
	_ExternalConnectorsPropertyTypePString             = ExternalConnectorsPropertyTypeVString
	_ExternalConnectorsPropertyTypePInt64              = ExternalConnectorsPropertyTypeVInt64
	_ExternalConnectorsPropertyTypePDouble             = ExternalConnectorsPropertyTypeVDouble
	_ExternalConnectorsPropertyTypePDateTime           = ExternalConnectorsPropertyTypeVDateTime
	_ExternalConnectorsPropertyTypePBoolean            = ExternalConnectorsPropertyTypeVBoolean
	_ExternalConnectorsPropertyTypePStringCollection   = ExternalConnectorsPropertyTypeVStringCollection
	_ExternalConnectorsPropertyTypePInt64Collection    = ExternalConnectorsPropertyTypeVInt64Collection
	_ExternalConnectorsPropertyTypePDoubleCollection   = ExternalConnectorsPropertyTypeVDoubleCollection
	_ExternalConnectorsPropertyTypePDateTimeCollection = ExternalConnectorsPropertyTypeVDateTimeCollection
	_ExternalConnectorsPropertyTypePUnknownFutureValue = ExternalConnectorsPropertyTypeVUnknownFutureValue
)

// ExternalConnectorsRuleOperation undocumented
type ExternalConnectorsRuleOperation string

const (
	// ExternalConnectorsRuleOperationVNull undocumented
	ExternalConnectorsRuleOperationVNull ExternalConnectorsRuleOperation = "null"
	// ExternalConnectorsRuleOperationVEquals undocumented
	ExternalConnectorsRuleOperationVEquals ExternalConnectorsRuleOperation = "equals"
	// ExternalConnectorsRuleOperationVNotEquals undocumented
	ExternalConnectorsRuleOperationVNotEquals ExternalConnectorsRuleOperation = "notEquals"
	// ExternalConnectorsRuleOperationVContains undocumented
	ExternalConnectorsRuleOperationVContains ExternalConnectorsRuleOperation = "contains"
	// ExternalConnectorsRuleOperationVNotContains undocumented
	ExternalConnectorsRuleOperationVNotContains ExternalConnectorsRuleOperation = "notContains"
	// ExternalConnectorsRuleOperationVLessThan undocumented
	ExternalConnectorsRuleOperationVLessThan ExternalConnectorsRuleOperation = "lessThan"
	// ExternalConnectorsRuleOperationVGreaterThan undocumented
	ExternalConnectorsRuleOperationVGreaterThan ExternalConnectorsRuleOperation = "greaterThan"
	// ExternalConnectorsRuleOperationVStartsWith undocumented
	ExternalConnectorsRuleOperationVStartsWith ExternalConnectorsRuleOperation = "startsWith"
	// ExternalConnectorsRuleOperationVUnknownFutureValue undocumented
	ExternalConnectorsRuleOperationVUnknownFutureValue ExternalConnectorsRuleOperation = "unknownFutureValue"
)

var (
	// ExternalConnectorsRuleOperationPNull is a pointer to ExternalConnectorsRuleOperationVNull
	ExternalConnectorsRuleOperationPNull = &_ExternalConnectorsRuleOperationPNull
	// ExternalConnectorsRuleOperationPEquals is a pointer to ExternalConnectorsRuleOperationVEquals
	ExternalConnectorsRuleOperationPEquals = &_ExternalConnectorsRuleOperationPEquals
	// ExternalConnectorsRuleOperationPNotEquals is a pointer to ExternalConnectorsRuleOperationVNotEquals
	ExternalConnectorsRuleOperationPNotEquals = &_ExternalConnectorsRuleOperationPNotEquals
	// ExternalConnectorsRuleOperationPContains is a pointer to ExternalConnectorsRuleOperationVContains
	ExternalConnectorsRuleOperationPContains = &_ExternalConnectorsRuleOperationPContains
	// ExternalConnectorsRuleOperationPNotContains is a pointer to ExternalConnectorsRuleOperationVNotContains
	ExternalConnectorsRuleOperationPNotContains = &_ExternalConnectorsRuleOperationPNotContains
	// ExternalConnectorsRuleOperationPLessThan is a pointer to ExternalConnectorsRuleOperationVLessThan
	ExternalConnectorsRuleOperationPLessThan = &_ExternalConnectorsRuleOperationPLessThan
	// ExternalConnectorsRuleOperationPGreaterThan is a pointer to ExternalConnectorsRuleOperationVGreaterThan
	ExternalConnectorsRuleOperationPGreaterThan = &_ExternalConnectorsRuleOperationPGreaterThan
	// ExternalConnectorsRuleOperationPStartsWith is a pointer to ExternalConnectorsRuleOperationVStartsWith
	ExternalConnectorsRuleOperationPStartsWith = &_ExternalConnectorsRuleOperationPStartsWith
	// ExternalConnectorsRuleOperationPUnknownFutureValue is a pointer to ExternalConnectorsRuleOperationVUnknownFutureValue
	ExternalConnectorsRuleOperationPUnknownFutureValue = &_ExternalConnectorsRuleOperationPUnknownFutureValue
)

var (
	_ExternalConnectorsRuleOperationPNull               = ExternalConnectorsRuleOperationVNull
	_ExternalConnectorsRuleOperationPEquals             = ExternalConnectorsRuleOperationVEquals
	_ExternalConnectorsRuleOperationPNotEquals          = ExternalConnectorsRuleOperationVNotEquals
	_ExternalConnectorsRuleOperationPContains           = ExternalConnectorsRuleOperationVContains
	_ExternalConnectorsRuleOperationPNotContains        = ExternalConnectorsRuleOperationVNotContains
	_ExternalConnectorsRuleOperationPLessThan           = ExternalConnectorsRuleOperationVLessThan
	_ExternalConnectorsRuleOperationPGreaterThan        = ExternalConnectorsRuleOperationVGreaterThan
	_ExternalConnectorsRuleOperationPStartsWith         = ExternalConnectorsRuleOperationVStartsWith
	_ExternalConnectorsRuleOperationPUnknownFutureValue = ExternalConnectorsRuleOperationVUnknownFutureValue
)
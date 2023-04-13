// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Identity undocumented
type Identity struct {
	// Object is the base model of Identity
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
}

func NewIdentity() (*Identity, error) {
	newIdentity := &Identity{
		ODataType: "#microsoft.graph.Identity",
	}
	return newIdentity, nil
}

// IdentityAPIConnector undocumented
type IdentityAPIConnector struct {
	// Entity is the base model of IdentityAPIConnector
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AuthenticationConfiguration undocumented
	AuthenticationConfiguration *APIAuthenticationConfigurationBase `json:"authenticationConfiguration,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// TargetURL undocumented
	TargetURL *string `json:"targetUrl,omitempty"`
}

func NewIdentityAPIConnector() (*IdentityAPIConnector, error) {
	newIdentityAPIConnector := &IdentityAPIConnector{
		ODataType: "#microsoft.graph.IdentityApiConnector",
	}
	return newIdentityAPIConnector, nil
}

// IdentityBuiltInUserFlowAttribute undocumented
type IdentityBuiltInUserFlowAttribute struct {
	// IdentityUserFlowAttribute is the base model of IdentityBuiltInUserFlowAttribute
	IdentityUserFlowAttribute

	ODataType string `json:"@odata.type,omitempty"`
}

func NewIdentityBuiltInUserFlowAttribute() (*IdentityBuiltInUserFlowAttribute, error) {
	newIdentityBuiltInUserFlowAttribute := &IdentityBuiltInUserFlowAttribute{
		ODataType: "#microsoft.graph.IdentityBuiltInUserFlowAttribute",
	}
	return newIdentityBuiltInUserFlowAttribute, nil
}

// IdentityContainer undocumented
type IdentityContainer struct {
	// Entity is the base model of IdentityContainer
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// APIConnectors undocumented
	APIConnectors []IdentityAPIConnector `json:"apiConnectors,omitempty"`
	// B2xUserFlows undocumented
	B2xUserFlows []B2xIdentityUserFlow `json:"b2xUserFlows,omitempty"`
	// IdentityProviders undocumented
	IdentityProviders []IdentityProviderBase `json:"identityProviders,omitempty"`
	// UserFlowAttributes undocumented
	UserFlowAttributes []IdentityUserFlowAttribute `json:"userFlowAttributes,omitempty"`
	// ConditionalAccess undocumented
	ConditionalAccess *ConditionalAccessRoot `json:"conditionalAccess,omitempty"`
}

func NewIdentityContainer() (*IdentityContainer, error) {
	newIdentityContainer := &IdentityContainer{
		ODataType: "#microsoft.graph.IdentityContainer",
	}
	return newIdentityContainer, nil
}

// IdentityCustomUserFlowAttribute undocumented
type IdentityCustomUserFlowAttribute struct {
	// IdentityUserFlowAttribute is the base model of IdentityCustomUserFlowAttribute
	IdentityUserFlowAttribute

	ODataType string `json:"@odata.type,omitempty"`
}

func NewIdentityCustomUserFlowAttribute() (*IdentityCustomUserFlowAttribute, error) {
	newIdentityCustomUserFlowAttribute := &IdentityCustomUserFlowAttribute{
		ODataType: "#microsoft.graph.IdentityCustomUserFlowAttribute",
	}
	return newIdentityCustomUserFlowAttribute, nil
}

// IdentityGovernance undocumented
type IdentityGovernance struct {
	// Object is the base model of IdentityGovernance
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AccessReviews undocumented
	AccessReviews *AccessReviewSet `json:"accessReviews,omitempty"`
	// AppConsent undocumented
	AppConsent *AppConsentApprovalRoute `json:"appConsent,omitempty"`
	// TermsOfUse undocumented
	TermsOfUse *TermsOfUseContainer `json:"termsOfUse,omitempty"`
	// EntitlementManagement undocumented
	EntitlementManagement *EntitlementManagement `json:"entitlementManagement,omitempty"`
}

func NewIdentityGovernance() (*IdentityGovernance, error) {
	newIdentityGovernance := &IdentityGovernance{
		ODataType: "#microsoft.graph.IdentityGovernance",
	}
	return newIdentityGovernance, nil
}

// IdentityProtectionRoot undocumented
type IdentityProtectionRoot struct {
	// Object is the base model of IdentityProtectionRoot
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// RiskDetections undocumented
	RiskDetections []RiskDetection `json:"riskDetections,omitempty"`
	// RiskyServicePrincipals undocumented
	RiskyServicePrincipals []RiskyServicePrincipal `json:"riskyServicePrincipals,omitempty"`
	// RiskyUsers undocumented
	RiskyUsers []RiskyUser `json:"riskyUsers,omitempty"`
	// ServicePrincipalRiskDetections undocumented
	ServicePrincipalRiskDetections []ServicePrincipalRiskDetection `json:"servicePrincipalRiskDetections,omitempty"`
}

func NewIdentityProtectionRoot() (*IdentityProtectionRoot, error) {
	newIdentityProtectionRoot := &IdentityProtectionRoot{
		ODataType: "#microsoft.graph.IdentityProtectionRoot",
	}
	return newIdentityProtectionRoot, nil
}

// IdentityProvider undocumented
type IdentityProvider struct {
	// Entity is the base model of IdentityProvider
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ClientID undocumented
	ClientID *string `json:"clientId,omitempty"`
	// ClientSecret undocumented
	ClientSecret *string `json:"clientSecret,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

func NewIdentityProvider() (*IdentityProvider, error) {
	newIdentityProvider := &IdentityProvider{
		ODataType: "#microsoft.graph.IdentityProvider",
	}
	return newIdentityProvider, nil
}

// IdentityProviderBase undocumented
type IdentityProviderBase struct {
	// Entity is the base model of IdentityProviderBase
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

func NewIdentityProviderBase() (*IdentityProviderBase, error) {
	newIdentityProviderBase := &IdentityProviderBase{
		ODataType: "#microsoft.graph.IdentityProviderBase",
	}
	return newIdentityProviderBase, nil
}

// IdentitySecurityDefaultsEnforcementPolicy undocumented
type IdentitySecurityDefaultsEnforcementPolicy struct {
	// PolicyBase is the base model of IdentitySecurityDefaultsEnforcementPolicy
	PolicyBase

	ODataType string `json:"@odata.type,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

func NewIdentitySecurityDefaultsEnforcementPolicy() (*IdentitySecurityDefaultsEnforcementPolicy, error) {
	newIdentitySecurityDefaultsEnforcementPolicy := &IdentitySecurityDefaultsEnforcementPolicy{
		ODataType: "#microsoft.graph.IdentitySecurityDefaultsEnforcementPolicy",
	}
	return newIdentitySecurityDefaultsEnforcementPolicy, nil
}

// IdentitySet undocumented
type IdentitySet struct {
	// Object is the base model of IdentitySet
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Application undocumented
	Application *Identity `json:"application,omitempty"`
	// Device undocumented
	Device *Identity `json:"device,omitempty"`
	// User undocumented
	User *Identity `json:"user,omitempty"`
}

func NewIdentitySet() (*IdentitySet, error) {
	newIdentitySet := &IdentitySet{
		ODataType: "#microsoft.graph.IdentitySet",
	}
	return newIdentitySet, nil
}

// IdentitySource undocumented
type IdentitySource struct {
	// Object is the base model of IdentitySource
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewIdentitySource() (*IdentitySource, error) {
	newIdentitySource := &IdentitySource{
		ODataType: "#microsoft.graph.IdentitySource",
	}
	return newIdentitySource, nil
}

// IdentityUserFlow undocumented
type IdentityUserFlow struct {
	// Entity is the base model of IdentityUserFlow
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// UserFlowType undocumented
	UserFlowType *UserFlowType `json:"userFlowType,omitempty"`
	// UserFlowTypeVersion undocumented
	UserFlowTypeVersion *float64 `json:"userFlowTypeVersion,omitempty"`
}

func NewIdentityUserFlow() (*IdentityUserFlow, error) {
	newIdentityUserFlow := &IdentityUserFlow{
		ODataType: "#microsoft.graph.IdentityUserFlow",
	}
	return newIdentityUserFlow, nil
}

// IdentityUserFlowAttribute undocumented
type IdentityUserFlowAttribute struct {
	// Entity is the base model of IdentityUserFlowAttribute
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DataType undocumented
	DataType *IdentityUserFlowAttributeDataType `json:"dataType,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// UserFlowAttributeType undocumented
	UserFlowAttributeType *IdentityUserFlowAttributeType `json:"userFlowAttributeType,omitempty"`
}

func NewIdentityUserFlowAttribute() (*IdentityUserFlowAttribute, error) {
	newIdentityUserFlowAttribute := &IdentityUserFlowAttribute{
		ODataType: "#microsoft.graph.IdentityUserFlowAttribute",
	}
	return newIdentityUserFlowAttribute, nil
}

// IdentityUserFlowAttributeAssignment undocumented
type IdentityUserFlowAttributeAssignment struct {
	// Entity is the base model of IdentityUserFlowAttributeAssignment
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsOptional undocumented
	IsOptional *bool `json:"isOptional,omitempty"`
	// RequiresVerification undocumented
	RequiresVerification *bool `json:"requiresVerification,omitempty"`
	// UserAttributeValues undocumented
	UserAttributeValues []UserAttributeValuesItem `json:"userAttributeValues,omitempty"`
	// UserInputType undocumented
	UserInputType *IdentityUserFlowAttributeInputType `json:"userInputType,omitempty"`
	// UserAttribute undocumented
	UserAttribute *IdentityUserFlowAttribute `json:"userAttribute,omitempty"`
}

func NewIdentityUserFlowAttributeAssignment() (*IdentityUserFlowAttributeAssignment, error) {
	newIdentityUserFlowAttributeAssignment := &IdentityUserFlowAttributeAssignment{
		ODataType: "#microsoft.graph.IdentityUserFlowAttributeAssignment",
	}
	return newIdentityUserFlowAttributeAssignment, nil
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// DecisionItemPrincipalResourceMembership undocumented
type DecisionItemPrincipalResourceMembership struct {
	// Object is the base model of DecisionItemPrincipalResourceMembership
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// MembershipType undocumented
	MembershipType *DecisionItemPrincipalResourceMembershipType `json:"membershipType,omitempty"`
}

func NewDecisionItemPrincipalResourceMembership() (*DecisionItemPrincipalResourceMembership, error) {
	newDecisionItemPrincipalResourceMembership := &DecisionItemPrincipalResourceMembership{
		ODataType: "#microsoft.graph.DecisionItemPrincipalResourceMembership",
	}
	return newDecisionItemPrincipalResourceMembership, nil
}

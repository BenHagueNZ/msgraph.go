// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Phone undocumented
type Phone struct {
	// Object is the base model of Phone
	Object

	ODataType string `json:"@odata.type"`
	// Language undocumented
	Language *string `json:"language,omitempty"`
	// Number undocumented
	Number *string `json:"number,omitempty"`
	// Region undocumented
	Region *string `json:"region,omitempty"`
	// Type undocumented
	Type *PhoneType `json:"type,omitempty"`
}

func NewPhone() (*Phone, error) {
	newPhone := &Phone{
		ODataType: "#microsoft.graph.Phone",
	}
	return newPhone, nil
}

// PhoneAuthenticationMethod undocumented
type PhoneAuthenticationMethod struct {
	// AuthenticationMethod is the base model of PhoneAuthenticationMethod
	AuthenticationMethod

	ODataType string `json:"@odata.type"`
	// PhoneNumber undocumented
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// PhoneType undocumented
	PhoneType *AuthenticationPhoneType `json:"phoneType,omitempty"`
	// SmsSignInState undocumented
	SmsSignInState *AuthenticationMethodSignInState `json:"smsSignInState,omitempty"`
}

func NewPhoneAuthenticationMethod() (*PhoneAuthenticationMethod, error) {
	newPhoneAuthenticationMethod := &PhoneAuthenticationMethod{
		ODataType: "#microsoft.graph.PhoneAuthenticationMethod",
	}
	return newPhoneAuthenticationMethod, nil
}

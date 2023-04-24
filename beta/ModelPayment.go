// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// PaymentMethod undocumented
type PaymentMethod struct {
	// Entity is the base model of PaymentMethod
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

func NewPaymentMethod() (*PaymentMethod, error) {
	newPaymentMethod := &PaymentMethod{
		ODataType: "#microsoft.graph.PaymentMethod",
	}
	return newPaymentMethod, nil
}

// PaymentTerm undocumented
type PaymentTerm struct {
	// Entity is the base model of PaymentTerm
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// CalculateDiscountOnCreditMemos undocumented
	CalculateDiscountOnCreditMemos *bool `json:"calculateDiscountOnCreditMemos,omitempty"`
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// DiscountDateCalculation undocumented
	DiscountDateCalculation *string `json:"discountDateCalculation,omitempty"`
	// DiscountPercent undocumented
	DiscountPercent *int `json:"discountPercent,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// DueDateCalculation undocumented
	DueDateCalculation *string `json:"dueDateCalculation,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

func NewPaymentTerm() (*PaymentTerm, error) {
	newPaymentTerm := &PaymentTerm{
		ODataType: "#microsoft.graph.PaymentTerm",
	}
	return newPaymentTerm, nil
}

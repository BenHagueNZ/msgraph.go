// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Currency undocumented
type Currency struct {
	// Entity is the base model of Currency
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AmountDecimalPlaces undocumented
	AmountDecimalPlaces *string `json:"amountDecimalPlaces,omitempty"`
	// AmountRoundingPrecision undocumented
	AmountRoundingPrecision *int `json:"amountRoundingPrecision,omitempty"`
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Symbol undocumented
	Symbol *string `json:"symbol,omitempty"`
}

func NewCurrency() (*Currency, error) {
	newCurrency := &Currency{
		ODataType: "#microsoft.graph.Currency",
	}
	return newCurrency, nil
}

// CurrencyColumn undocumented
type CurrencyColumn struct {
	// Object is the base model of CurrencyColumn
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Locale undocumented
	Locale *string `json:"locale,omitempty"`
}

func NewCurrencyColumn() (*CurrencyColumn, error) {
	newCurrencyColumn := &CurrencyColumn{
		ODataType: "#microsoft.graph.CurrencyColumn",
	}
	return newCurrencyColumn, nil
}

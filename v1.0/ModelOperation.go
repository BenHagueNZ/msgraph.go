// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Operation undocumented
type Operation struct {
	// Entity is the base model of Operation
	Entity

	ODataType string `json:"@odata.type"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastActionDateTime undocumented
	LastActionDateTime *time.Time `json:"lastActionDateTime,omitempty"`
	// Status undocumented
	Status *OperationStatus `json:"status,omitempty"`
}

func NewOperation() (*Operation, error) {
	newOperation := &Operation{
		ODataType: "#microsoft.graph.Operation",
	}
	return newOperation, nil
}

// OperationError undocumented
type OperationError struct {
	// Object is the base model of OperationError
	Object

	ODataType string `json:"@odata.type"`
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
}

func NewOperationError() (*OperationError, error) {
	newOperationError := &OperationError{
		ODataType: "#microsoft.graph.OperationError",
	}
	return newOperationError, nil
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// CommsApplication undocumented
type CommsApplication struct {
	// Object is the base model of CommsApplication
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Calls undocumented
	Calls []Call `json:"calls,omitempty"`
	// OnlineMeetings undocumented
	OnlineMeetings []OnlineMeeting `json:"onlineMeetings,omitempty"`
}

func NewCommsApplication() (*CommsApplication, error) {
	newCommsApplication := &CommsApplication{
		ODataType: "#microsoft.graph.CommsApplication",
	}
	return newCommsApplication, nil
}

// CommsNotification undocumented
type CommsNotification struct {
	// Object is the base model of CommsNotification
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// ChangeType undocumented
	ChangeType *ChangeType `json:"changeType,omitempty"`
	// ResourceURL undocumented
	ResourceURL *string `json:"resourceUrl,omitempty"`
}

func NewCommsNotification() (*CommsNotification, error) {
	newCommsNotification := &CommsNotification{
		ODataType: "#microsoft.graph.CommsNotification",
	}
	return newCommsNotification, nil
}

// CommsNotifications undocumented
type CommsNotifications struct {
	// Object is the base model of CommsNotifications
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Value undocumented
	Value []CommsNotification `json:"value,omitempty"`
}

func NewCommsNotifications() (*CommsNotifications, error) {
	newCommsNotifications := &CommsNotifications{
		ODataType: "#microsoft.graph.CommsNotifications",
	}
	return newCommsNotifications, nil
}

// CommsOperation undocumented
type CommsOperation struct {
	// Entity is the base model of CommsOperation
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
	// ResultInfo undocumented
	ResultInfo *ResultInfo `json:"resultInfo,omitempty"`
	// Status undocumented
	Status *OperationStatus `json:"status,omitempty"`
}

func NewCommsOperation() (*CommsOperation, error) {
	newCommsOperation := &CommsOperation{
		ODataType: "#microsoft.graph.CommsOperation",
	}
	return newCommsOperation, nil
}

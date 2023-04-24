// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// DateTimeColumn undocumented
type DateTimeColumn struct {
	// Object is the base model of DateTimeColumn
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayAs undocumented
	DisplayAs *string `json:"displayAs,omitempty"`
	// Format undocumented
	Format *string `json:"format,omitempty"`
}

func NewDateTimeColumn() (*DateTimeColumn, error) {
	newDateTimeColumn := &DateTimeColumn{
		ODataType: "#microsoft.graph.DateTimeColumn",
	}
	return newDateTimeColumn, nil
}

// DateTimeTimeZone undocumented
type DateTimeTimeZone struct {
	// Object is the base model of DateTimeTimeZone
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DateTime undocumented
	DateTime *string `json:"dateTime,omitempty"`
	// TimeZone undocumented
	TimeZone *string `json:"timeZone,omitempty"`
}

func NewDateTimeTimeZone() (*DateTimeTimeZone, error) {
	newDateTimeTimeZone := &DateTimeTimeZone{
		ODataType: "#microsoft.graph.DateTimeTimeZone",
	}
	return newDateTimeTimeZone, nil
}

// DateTimeTimeZoneType undocumented
type DateTimeTimeZoneType struct {
	// Object is the base model of DateTimeTimeZoneType
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// DateTime undocumented
	DateTime *string `json:"dateTime,omitempty"`
}

func NewDateTimeTimeZoneType() (*DateTimeTimeZoneType, error) {
	newDateTimeTimeZoneType := &DateTimeTimeZoneType{
		ODataType: "#microsoft.graph.DateTimeTimeZoneType",
	}
	return newDateTimeTimeZoneType, nil
}

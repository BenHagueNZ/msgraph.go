// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// IosiPadOSWebClip undocumented
type IosiPadOSWebClip struct {
	// MobileApp is the base model of IosiPadOSWebClip
	MobileApp

	ODataType string `json:"@odata.type,omitempty"`
	// AppURL undocumented
	AppURL *string `json:"appUrl,omitempty"`
	// UseManagedBrowser undocumented
	UseManagedBrowser *bool `json:"useManagedBrowser,omitempty"`
}

func NewIosiPadOSWebClip() (*IosiPadOSWebClip, error) {
	newIosiPadOSWebClip := &IosiPadOSWebClip{
		ODataType: "#microsoft.graph.IosiPadOSWebClip",
	}
	return newIosiPadOSWebClip, nil
}
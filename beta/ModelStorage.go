// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Storage undocumented
type Storage struct {
	// Object is the base model of Storage
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewStorage() (*Storage, error) {
	newStorage := &Storage{
		ODataType: "#microsoft.graph.Storage",
	}
	return newStorage, nil
}

// StoragePlanInformation undocumented
type StoragePlanInformation struct {
	// Object is the base model of StoragePlanInformation
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// UpgradeAvailable undocumented
	UpgradeAvailable *bool `json:"upgradeAvailable,omitempty"`
}

func NewStoragePlanInformation() (*StoragePlanInformation, error) {
	newStoragePlanInformation := &StoragePlanInformation{
		ODataType: "#microsoft.graph.StoragePlanInformation",
	}
	return newStoragePlanInformation, nil
}

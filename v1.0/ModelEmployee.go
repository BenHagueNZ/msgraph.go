// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// EmployeeOrgData undocumented
type EmployeeOrgData struct {
	// Object is the base model of EmployeeOrgData
	Object

	ODataType string `json:"@odata.type"`
	// CostCenter undocumented
	CostCenter *string `json:"costCenter,omitempty"`
	// Division undocumented
	Division *string `json:"division,omitempty"`
}

func NewEmployeeOrgData() (*EmployeeOrgData, error) {
	newEmployeeOrgData := &EmployeeOrgData{
		ODataType: "#microsoft.graph.EmployeeOrgData",
	}
	return newEmployeeOrgData, nil
}

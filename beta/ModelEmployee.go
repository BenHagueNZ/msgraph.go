// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Employee undocumented
type Employee struct {
	// Entity is the base model of Employee
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Address undocumented
	Address *PostalAddressType `json:"address,omitempty"`
	// BirthDate undocumented
	BirthDate *Date `json:"birthDate,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Email undocumented
	Email *string `json:"email,omitempty"`
	// EmploymentDate undocumented
	EmploymentDate *Date `json:"employmentDate,omitempty"`
	// GivenName undocumented
	GivenName *string `json:"givenName,omitempty"`
	// JobTitle undocumented
	JobTitle *string `json:"jobTitle,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// MiddleName undocumented
	MiddleName *string `json:"middleName,omitempty"`
	// MobilePhone undocumented
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// Number undocumented
	Number *string `json:"number,omitempty"`
	// PersonalEmail undocumented
	PersonalEmail *string `json:"personalEmail,omitempty"`
	// PhoneNumber undocumented
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// StatisticsGroupCode undocumented
	StatisticsGroupCode *string `json:"statisticsGroupCode,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// Surname undocumented
	Surname *string `json:"surname,omitempty"`
	// TerminationDate undocumented
	TerminationDate *Date `json:"terminationDate,omitempty"`
	// Picture undocumented
	Picture []Picture `json:"picture,omitempty"`
}

func NewEmployee() (*Employee, error) {
	newEmployee := &Employee{
		ODataType: "#microsoft.graph.Employee",
	}
	return newEmployee, nil
}

// EmployeeExperience undocumented
type EmployeeExperience struct {
	// Object is the base model of EmployeeExperience
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// LearningProviders undocumented
	LearningProviders []LearningProvider `json:"learningProviders,omitempty"`
}

func NewEmployeeExperience() (*EmployeeExperience, error) {
	newEmployeeExperience := &EmployeeExperience{
		ODataType: "#microsoft.graph.EmployeeExperience",
	}
	return newEmployeeExperience, nil
}

// EmployeeOrgData undocumented
type EmployeeOrgData struct {
	// Object is the base model of EmployeeOrgData
	Object

	ODataType string `json:"@odata.type,omitempty"`
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
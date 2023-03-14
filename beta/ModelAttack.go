
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// AttackSimulationOperation undocumented
type AttackSimulationOperation struct {
    // LongRunningOperation is the base model of AttackSimulationOperation
    LongRunningOperation
    // PercentageCompleted undocumented
    PercentageCompleted *int `json:"percentageCompleted,omitempty"`
    // TenantID undocumented
    TenantID *string `json:"tenantId,omitempty"`
    // Type undocumented
    Type *AttackSimulationOperationType `json:"type,omitempty"`
}

// AttackSimulationRepeatOffender undocumented
type AttackSimulationRepeatOffender struct {
    // Object is the base model of AttackSimulationRepeatOffender
    Object
    // AttackSimulationUser undocumented
    AttackSimulationUser *AttackSimulationUser `json:"attackSimulationUser,omitempty"`
    // RepeatOffenceCount undocumented
    RepeatOffenceCount *int `json:"repeatOffenceCount,omitempty"`
}

// AttackSimulationRoot undocumented
type AttackSimulationRoot struct {
    // Entity is the base model of AttackSimulationRoot
    Entity
    // Operations undocumented
    Operations []AttackSimulationOperation `json:"operations,omitempty"`
    // Payloads undocumented
    Payloads []Payload `json:"payloads,omitempty"`
    // SimulationAutomations undocumented
    SimulationAutomations []SimulationAutomation `json:"simulationAutomations,omitempty"`
    // Simulations undocumented
    Simulations []Simulation `json:"simulations,omitempty"`
}

// AttackSimulationSimulationUserCoverage undocumented
type AttackSimulationSimulationUserCoverage struct {
    // Object is the base model of AttackSimulationSimulationUserCoverage
    Object
    // AttackSimulationUser undocumented
    AttackSimulationUser *AttackSimulationUser `json:"attackSimulationUser,omitempty"`
    // ClickCount undocumented
    ClickCount *int `json:"clickCount,omitempty"`
    // CompromisedCount undocumented
    CompromisedCount *int `json:"compromisedCount,omitempty"`
    // LatestSimulationDateTime undocumented
    LatestSimulationDateTime *time.Time `json:"latestSimulationDateTime,omitempty"`
    // SimulationCount undocumented
    SimulationCount *int `json:"simulationCount,omitempty"`
}

// AttackSimulationTrainingUserCoverage undocumented
type AttackSimulationTrainingUserCoverage struct {
    // Object is the base model of AttackSimulationTrainingUserCoverage
    Object
    // AttackSimulationUser undocumented
    AttackSimulationUser *AttackSimulationUser `json:"attackSimulationUser,omitempty"`
    // UserTrainings undocumented
    UserTrainings []UserTrainingStatusInfo `json:"userTrainings,omitempty"`
}

// AttackSimulationUser undocumented
type AttackSimulationUser struct {
    // Object is the base model of AttackSimulationUser
    Object
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // Email undocumented
    Email *string `json:"email,omitempty"`
    // UserID undocumented
    UserID *string `json:"userId,omitempty"`
}
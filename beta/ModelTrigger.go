// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TriggerAndScopeBasedConditions undocumented
type TriggerAndScopeBasedConditions struct {
	// WorkflowExecutionConditions is the base model of TriggerAndScopeBasedConditions
	WorkflowExecutionConditions
	// Scope undocumented
	Scope *SubjectSet `json:"scope,omitempty"`
	// Trigger undocumented
	Trigger *WorkflowExecutionTrigger `json:"trigger,omitempty"`
}

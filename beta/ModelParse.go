
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// ParseExpressionResponse undocumented
type ParseExpressionResponse struct {
    // Object is the base model of ParseExpressionResponse
    Object
    // Error undocumented
    Error *PublicError `json:"error,omitempty"`
    // EvaluationResult undocumented
    EvaluationResult []string `json:"evaluationResult,omitempty"`
    // EvaluationSucceeded undocumented
    EvaluationSucceeded *bool `json:"evaluationSucceeded,omitempty"`
    // ParsedExpression undocumented
    ParsedExpression *AttributeMappingSource `json:"parsedExpression,omitempty"`
    // ParsingSucceeded undocumented
    ParsingSucceeded *bool `json:"parsingSucceeded,omitempty"`
}

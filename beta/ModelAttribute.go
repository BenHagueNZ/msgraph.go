
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// AttributeDefinition undocumented
type AttributeDefinition struct {
    // Object is the base model of AttributeDefinition
    Object
    // Anchor undocumented
    Anchor *bool `json:"anchor,omitempty"`
    // APIExpressions undocumented
    APIExpressions []StringKeyStringValuePair `json:"apiExpressions,omitempty"`
    // CaseExact undocumented
    CaseExact *bool `json:"caseExact,omitempty"`
    // DefaultValue undocumented
    DefaultValue *string `json:"defaultValue,omitempty"`
    // FlowNullValues undocumented
    FlowNullValues *bool `json:"flowNullValues,omitempty"`
    // Metadata undocumented
    Metadata []MetadataEntry `json:"metadata,omitempty"`
    // Multivalued undocumented
    Multivalued *bool `json:"multivalued,omitempty"`
    // Mutability undocumented
    Mutability *Mutability `json:"mutability,omitempty"`
    // Name undocumented
    Name *string `json:"name,omitempty"`
    // ReferencedObjects undocumented
    ReferencedObjects []ReferencedObject `json:"referencedObjects,omitempty"`
    // Required undocumented
    Required *bool `json:"required,omitempty"`
    // Type undocumented
    Type *AttributeType `json:"type,omitempty"`
}

// AttributeMapping undocumented
type AttributeMapping struct {
    // Object is the base model of AttributeMapping
    Object
    // DefaultValue undocumented
    DefaultValue *string `json:"defaultValue,omitempty"`
    // ExportMissingReferences undocumented
    ExportMissingReferences *bool `json:"exportMissingReferences,omitempty"`
    // FlowBehavior undocumented
    FlowBehavior *AttributeFlowBehavior `json:"flowBehavior,omitempty"`
    // FlowType undocumented
    FlowType *AttributeFlowType `json:"flowType,omitempty"`
    // MatchingPriority undocumented
    MatchingPriority *int `json:"matchingPriority,omitempty"`
    // Source undocumented
    Source *AttributeMappingSource `json:"source,omitempty"`
    // TargetAttributeName undocumented
    TargetAttributeName *string `json:"targetAttributeName,omitempty"`
}

// AttributeMappingFunctionSchema undocumented
type AttributeMappingFunctionSchema struct {
    // Entity is the base model of AttributeMappingFunctionSchema
    Entity
    // Parameters undocumented
    Parameters []AttributeMappingParameterSchema `json:"parameters,omitempty"`
}

// AttributeMappingParameterSchema undocumented
type AttributeMappingParameterSchema struct {
    // Object is the base model of AttributeMappingParameterSchema
    Object
    // AllowMultipleOccurrences undocumented
    AllowMultipleOccurrences *bool `json:"allowMultipleOccurrences,omitempty"`
    // Name undocumented
    Name *string `json:"name,omitempty"`
    // Required undocumented
    Required *bool `json:"required,omitempty"`
    // Type undocumented
    Type *AttributeType `json:"type,omitempty"`
}

// AttributeMappingSource undocumented
type AttributeMappingSource struct {
    // Object is the base model of AttributeMappingSource
    Object
    // Expression undocumented
    Expression *string `json:"expression,omitempty"`
    // Name undocumented
    Name *string `json:"name,omitempty"`
    // Parameters undocumented
    Parameters []StringKeyAttributeMappingSourceValuePair `json:"parameters,omitempty"`
    // Type undocumented
    Type *AttributeMappingSourceType `json:"type,omitempty"`
}

// AttributeSet undocumented
type AttributeSet struct {
    // Entity is the base model of AttributeSet
    Entity
    // Description undocumented
    Description *string `json:"description,omitempty"`
    // MaxAttributesPerSet undocumented
    MaxAttributesPerSet *int `json:"maxAttributesPerSet,omitempty"`
}

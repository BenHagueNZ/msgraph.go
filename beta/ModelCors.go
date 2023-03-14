
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// CorsConfiguration undocumented
type CorsConfiguration struct {
    // Object is the base model of CorsConfiguration
    Object
    // AllowedHeaders undocumented
    AllowedHeaders []string `json:"allowedHeaders,omitempty"`
    // AllowedMethods undocumented
    AllowedMethods []string `json:"allowedMethods,omitempty"`
    // AllowedOrigins undocumented
    AllowedOrigins []string `json:"allowedOrigins,omitempty"`
    // MaxAgeInSeconds undocumented
    MaxAgeInSeconds *int `json:"maxAgeInSeconds,omitempty"`
    // Resource undocumented
    Resource *string `json:"resource,omitempty"`
}

// CorsConfiguration_v2 undocumented
type CorsConfiguration_v2 struct {
    // Entity is the base model of CorsConfiguration_v2
    Entity
    // AllowedHeaders undocumented
    AllowedHeaders []string `json:"allowedHeaders,omitempty"`
    // AllowedMethods undocumented
    AllowedMethods []string `json:"allowedMethods,omitempty"`
    // AllowedOrigins undocumented
    AllowedOrigins []string `json:"allowedOrigins,omitempty"`
    // MaxAgeInSeconds undocumented
    MaxAgeInSeconds *int `json:"maxAgeInSeconds,omitempty"`
    // Resource undocumented
    Resource *string `json:"resource,omitempty"`
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// LoggedOnUser undocumented
type LoggedOnUser struct {
    // Object is the base model of LoggedOnUser
    Object
    // LastLogOnDateTime undocumented
    LastLogOnDateTime *time.Time `json:"lastLogOnDateTime,omitempty"`
    // UserID undocumented
    UserID *string `json:"userId,omitempty"`
}

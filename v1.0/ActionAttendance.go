// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AttendanceRecord is navigation property rn
func (b *AttendanceRecordRequestBuilder) AttendanceRecord() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

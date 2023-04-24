// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AttendeeRequestBuilder is request builder for Attendee
type AttendeeRequestBuilder struct{ BaseRequestBuilder }

// Request returns AttendeeRequest
func (b *AttendeeRequestBuilder) Request() *AttendeeRequest {
	return &AttendeeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AttendeeRequest is request for Attendee
type AttendeeRequest struct{ BaseRequest }

// Get performs GET request for Attendee
func (r *AttendeeRequest) Get(ctx context.Context) (resObj *Attendee, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Attendee
func (r *AttendeeRequest) Update(ctx context.Context, reqObj *Attendee) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Attendee
func (r *AttendeeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AttendeeAvailabilityRequestBuilder is request builder for AttendeeAvailability
type AttendeeAvailabilityRequestBuilder struct{ BaseRequestBuilder }

// Request returns AttendeeAvailabilityRequest
func (b *AttendeeAvailabilityRequestBuilder) Request() *AttendeeAvailabilityRequest {
	return &AttendeeAvailabilityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AttendeeAvailabilityRequest is request for AttendeeAvailability
type AttendeeAvailabilityRequest struct{ BaseRequest }

// Get performs GET request for AttendeeAvailability
func (r *AttendeeAvailabilityRequest) Get(ctx context.Context) (resObj *AttendeeAvailability, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AttendeeAvailability
func (r *AttendeeAvailabilityRequest) Update(ctx context.Context, reqObj *AttendeeAvailability) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AttendeeAvailability
func (r *AttendeeAvailabilityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AttendeeBaseRequestBuilder is request builder for AttendeeBase
type AttendeeBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns AttendeeBaseRequest
func (b *AttendeeBaseRequestBuilder) Request() *AttendeeBaseRequest {
	return &AttendeeBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AttendeeBaseRequest is request for AttendeeBase
type AttendeeBaseRequest struct{ BaseRequest }

// Get performs GET request for AttendeeBase
func (r *AttendeeBaseRequest) Get(ctx context.Context) (resObj *AttendeeBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AttendeeBase
func (r *AttendeeBaseRequest) Update(ctx context.Context, reqObj *AttendeeBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AttendeeBase
func (r *AttendeeBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
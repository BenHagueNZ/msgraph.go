// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ScheduleRequestBuilder is request builder for Schedule
type ScheduleRequestBuilder struct{ BaseRequestBuilder }

// Request returns ScheduleRequest
func (b *ScheduleRequestBuilder) Request() *ScheduleRequest {
	return &ScheduleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ScheduleRequest is request for Schedule
type ScheduleRequest struct{ BaseRequest }

// Get performs GET request for Schedule
func (r *ScheduleRequest) Get(ctx context.Context) (resObj *Schedule, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Schedule
func (r *ScheduleRequest) Update(ctx context.Context, reqObj *Schedule) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Schedule
func (r *ScheduleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ScheduleChangeRequestObjectRequestBuilder is request builder for ScheduleChangeRequestObject
type ScheduleChangeRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns ScheduleChangeRequestObjectRequest
func (b *ScheduleChangeRequestObjectRequestBuilder) Request() *ScheduleChangeRequestObjectRequest {
	return &ScheduleChangeRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ScheduleChangeRequestObjectRequest is request for ScheduleChangeRequestObject
type ScheduleChangeRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for ScheduleChangeRequestObject
func (r *ScheduleChangeRequestObjectRequest) Get(ctx context.Context) (resObj *ScheduleChangeRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ScheduleChangeRequestObject
func (r *ScheduleChangeRequestObjectRequest) Update(ctx context.Context, reqObj *ScheduleChangeRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ScheduleChangeRequestObject
func (r *ScheduleChangeRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ScheduleEntityRequestBuilder is request builder for ScheduleEntity
type ScheduleEntityRequestBuilder struct{ BaseRequestBuilder }

// Request returns ScheduleEntityRequest
func (b *ScheduleEntityRequestBuilder) Request() *ScheduleEntityRequest {
	return &ScheduleEntityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ScheduleEntityRequest is request for ScheduleEntity
type ScheduleEntityRequest struct{ BaseRequest }

// Get performs GET request for ScheduleEntity
func (r *ScheduleEntityRequest) Get(ctx context.Context) (resObj *ScheduleEntity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ScheduleEntity
func (r *ScheduleEntityRequest) Update(ctx context.Context, reqObj *ScheduleEntity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ScheduleEntity
func (r *ScheduleEntityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ScheduleInformationRequestBuilder is request builder for ScheduleInformation
type ScheduleInformationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ScheduleInformationRequest
func (b *ScheduleInformationRequestBuilder) Request() *ScheduleInformationRequest {
	return &ScheduleInformationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ScheduleInformationRequest is request for ScheduleInformation
type ScheduleInformationRequest struct{ BaseRequest }

// Get performs GET request for ScheduleInformation
func (r *ScheduleInformationRequest) Get(ctx context.Context) (resObj *ScheduleInformation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ScheduleInformation
func (r *ScheduleInformationRequest) Update(ctx context.Context, reqObj *ScheduleInformation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ScheduleInformation
func (r *ScheduleInformationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ScheduleItemRequestBuilder is request builder for ScheduleItem
type ScheduleItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns ScheduleItemRequest
func (b *ScheduleItemRequestBuilder) Request() *ScheduleItemRequest {
	return &ScheduleItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ScheduleItemRequest is request for ScheduleItem
type ScheduleItemRequest struct{ BaseRequest }

// Get performs GET request for ScheduleItem
func (r *ScheduleItemRequest) Get(ctx context.Context) (resObj *ScheduleItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ScheduleItem
func (r *ScheduleItemRequest) Update(ctx context.Context, reqObj *ScheduleItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ScheduleItem
func (r *ScheduleItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

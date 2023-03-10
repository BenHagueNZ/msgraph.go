// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MembersAddedEventMessageDetailRequestBuilder is request builder for MembersAddedEventMessageDetail
type MembersAddedEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns MembersAddedEventMessageDetailRequest
func (b *MembersAddedEventMessageDetailRequestBuilder) Request() *MembersAddedEventMessageDetailRequest {
	return &MembersAddedEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MembersAddedEventMessageDetailRequest is request for MembersAddedEventMessageDetail
type MembersAddedEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for MembersAddedEventMessageDetail
func (r *MembersAddedEventMessageDetailRequest) Get(ctx context.Context) (resObj *MembersAddedEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MembersAddedEventMessageDetail
func (r *MembersAddedEventMessageDetailRequest) Update(ctx context.Context, reqObj *MembersAddedEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MembersAddedEventMessageDetail
func (r *MembersAddedEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MembersDeletedEventMessageDetailRequestBuilder is request builder for MembersDeletedEventMessageDetail
type MembersDeletedEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns MembersDeletedEventMessageDetailRequest
func (b *MembersDeletedEventMessageDetailRequestBuilder) Request() *MembersDeletedEventMessageDetailRequest {
	return &MembersDeletedEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MembersDeletedEventMessageDetailRequest is request for MembersDeletedEventMessageDetail
type MembersDeletedEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for MembersDeletedEventMessageDetail
func (r *MembersDeletedEventMessageDetailRequest) Get(ctx context.Context) (resObj *MembersDeletedEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MembersDeletedEventMessageDetail
func (r *MembersDeletedEventMessageDetailRequest) Update(ctx context.Context, reqObj *MembersDeletedEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MembersDeletedEventMessageDetail
func (r *MembersDeletedEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MembersJoinedEventMessageDetailRequestBuilder is request builder for MembersJoinedEventMessageDetail
type MembersJoinedEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns MembersJoinedEventMessageDetailRequest
func (b *MembersJoinedEventMessageDetailRequestBuilder) Request() *MembersJoinedEventMessageDetailRequest {
	return &MembersJoinedEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MembersJoinedEventMessageDetailRequest is request for MembersJoinedEventMessageDetail
type MembersJoinedEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for MembersJoinedEventMessageDetail
func (r *MembersJoinedEventMessageDetailRequest) Get(ctx context.Context) (resObj *MembersJoinedEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MembersJoinedEventMessageDetail
func (r *MembersJoinedEventMessageDetailRequest) Update(ctx context.Context, reqObj *MembersJoinedEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MembersJoinedEventMessageDetail
func (r *MembersJoinedEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MembersLeftEventMessageDetailRequestBuilder is request builder for MembersLeftEventMessageDetail
type MembersLeftEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns MembersLeftEventMessageDetailRequest
func (b *MembersLeftEventMessageDetailRequestBuilder) Request() *MembersLeftEventMessageDetailRequest {
	return &MembersLeftEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MembersLeftEventMessageDetailRequest is request for MembersLeftEventMessageDetail
type MembersLeftEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for MembersLeftEventMessageDetail
func (r *MembersLeftEventMessageDetailRequest) Get(ctx context.Context) (resObj *MembersLeftEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MembersLeftEventMessageDetail
func (r *MembersLeftEventMessageDetailRequest) Update(ctx context.Context, reqObj *MembersLeftEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MembersLeftEventMessageDetail
func (r *MembersLeftEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

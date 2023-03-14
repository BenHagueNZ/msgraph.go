// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SharingDetailRequestBuilder is request builder for SharingDetail
type SharingDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns SharingDetailRequest
func (b *SharingDetailRequestBuilder) Request() *SharingDetailRequest {
	return &SharingDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SharingDetailRequest is request for SharingDetail
type SharingDetailRequest struct{ BaseRequest }

// Get performs GET request for SharingDetail
func (r *SharingDetailRequest) Get(ctx context.Context) (resObj *SharingDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SharingDetail
func (r *SharingDetailRequest) Update(ctx context.Context, reqObj *SharingDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SharingDetail
func (r *SharingDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SharingInvitationRequestBuilder is request builder for SharingInvitation
type SharingInvitationRequestBuilder struct{ BaseRequestBuilder }

// Request returns SharingInvitationRequest
func (b *SharingInvitationRequestBuilder) Request() *SharingInvitationRequest {
	return &SharingInvitationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SharingInvitationRequest is request for SharingInvitation
type SharingInvitationRequest struct{ BaseRequest }

// Get performs GET request for SharingInvitation
func (r *SharingInvitationRequest) Get(ctx context.Context) (resObj *SharingInvitation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SharingInvitation
func (r *SharingInvitationRequest) Update(ctx context.Context, reqObj *SharingInvitation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SharingInvitation
func (r *SharingInvitationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SharingLinkRequestBuilder is request builder for SharingLink
type SharingLinkRequestBuilder struct{ BaseRequestBuilder }

// Request returns SharingLinkRequest
func (b *SharingLinkRequestBuilder) Request() *SharingLinkRequest {
	return &SharingLinkRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SharingLinkRequest is request for SharingLink
type SharingLinkRequest struct{ BaseRequest }

// Get performs GET request for SharingLink
func (r *SharingLinkRequest) Get(ctx context.Context) (resObj *SharingLink, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SharingLink
func (r *SharingLinkRequest) Update(ctx context.Context, reqObj *SharingLink) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SharingLink
func (r *SharingLinkRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
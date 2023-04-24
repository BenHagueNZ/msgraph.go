// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ConnectedOrganizationRequestBuilder is request builder for ConnectedOrganization
type ConnectedOrganizationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConnectedOrganizationRequest
func (b *ConnectedOrganizationRequestBuilder) Request() *ConnectedOrganizationRequest {
	return &ConnectedOrganizationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConnectedOrganizationRequest is request for ConnectedOrganization
type ConnectedOrganizationRequest struct{ BaseRequest }

// Get performs GET request for ConnectedOrganization
func (r *ConnectedOrganizationRequest) Get(ctx context.Context) (resObj *ConnectedOrganization, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConnectedOrganization
func (r *ConnectedOrganizationRequest) Update(ctx context.Context, reqObj *ConnectedOrganization) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConnectedOrganization
func (r *ConnectedOrganizationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConnectedOrganizationMembersRequestBuilder is request builder for ConnectedOrganizationMembers
type ConnectedOrganizationMembersRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConnectedOrganizationMembersRequest
func (b *ConnectedOrganizationMembersRequestBuilder) Request() *ConnectedOrganizationMembersRequest {
	return &ConnectedOrganizationMembersRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConnectedOrganizationMembersRequest is request for ConnectedOrganizationMembers
type ConnectedOrganizationMembersRequest struct{ BaseRequest }

// Get performs GET request for ConnectedOrganizationMembers
func (r *ConnectedOrganizationMembersRequest) Get(ctx context.Context) (resObj *ConnectedOrganizationMembers, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConnectedOrganizationMembers
func (r *ConnectedOrganizationMembersRequest) Update(ctx context.Context, reqObj *ConnectedOrganizationMembers) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConnectedOrganizationMembers
func (r *ConnectedOrganizationMembersRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// DeletedItemContainerRequestBuilder is request builder for DeletedItemContainer
type DeletedItemContainerRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeletedItemContainerRequest
func (b *DeletedItemContainerRequestBuilder) Request() *DeletedItemContainerRequest {
	return &DeletedItemContainerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeletedItemContainerRequest is request for DeletedItemContainer
type DeletedItemContainerRequest struct{ BaseRequest }

// Get performs GET request for DeletedItemContainer
func (r *DeletedItemContainerRequest) Get(ctx context.Context) (resObj *DeletedItemContainer, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeletedItemContainer
func (r *DeletedItemContainerRequest) Update(ctx context.Context, reqObj *DeletedItemContainer) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeletedItemContainer
func (r *DeletedItemContainerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DeletedRequestBuilder is request builder for Deleted
type DeletedRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeletedRequest
func (b *DeletedRequestBuilder) Request() *DeletedRequest {
	return &DeletedRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeletedRequest is request for Deleted
type DeletedRequest struct{ BaseRequest }

// Get performs GET request for Deleted
func (r *DeletedRequest) Get(ctx context.Context) (resObj *Deleted, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Deleted
func (r *DeletedRequest) Update(ctx context.Context, reqObj *Deleted) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Deleted
func (r *DeletedRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DeletedTeamRequestBuilder is request builder for DeletedTeam
type DeletedTeamRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeletedTeamRequest
func (b *DeletedTeamRequestBuilder) Request() *DeletedTeamRequest {
	return &DeletedTeamRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeletedTeamRequest is request for DeletedTeam
type DeletedTeamRequest struct{ BaseRequest }

// Get performs GET request for DeletedTeam
func (r *DeletedTeamRequest) Get(ctx context.Context) (resObj *DeletedTeam, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeletedTeam
func (r *DeletedTeamRequest) Update(ctx context.Context, reqObj *DeletedTeam) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeletedTeam
func (r *DeletedTeamRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

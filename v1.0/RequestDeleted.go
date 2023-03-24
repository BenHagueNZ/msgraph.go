// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

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

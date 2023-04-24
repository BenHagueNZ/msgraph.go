// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// LoggedOnUserRequestBuilder is request builder for LoggedOnUser
type LoggedOnUserRequestBuilder struct{ BaseRequestBuilder }

// Request returns LoggedOnUserRequest
func (b *LoggedOnUserRequestBuilder) Request() *LoggedOnUserRequest {
	return &LoggedOnUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// LoggedOnUserRequest is request for LoggedOnUser
type LoggedOnUserRequest struct{ BaseRequest }

// Get performs GET request for LoggedOnUser
func (r *LoggedOnUserRequest) Get(ctx context.Context) (resObj *LoggedOnUser, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for LoggedOnUser
func (r *LoggedOnUserRequest) Update(ctx context.Context, reqObj *LoggedOnUser) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for LoggedOnUser
func (r *LoggedOnUserRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

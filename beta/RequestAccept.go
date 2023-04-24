// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AcceptJoinResponseRequestBuilder is request builder for AcceptJoinResponse
type AcceptJoinResponseRequestBuilder struct{ BaseRequestBuilder }

// Request returns AcceptJoinResponseRequest
func (b *AcceptJoinResponseRequestBuilder) Request() *AcceptJoinResponseRequest {
	return &AcceptJoinResponseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AcceptJoinResponseRequest is request for AcceptJoinResponse
type AcceptJoinResponseRequest struct{ BaseRequest }

// Get performs GET request for AcceptJoinResponse
func (r *AcceptJoinResponseRequest) Get(ctx context.Context) (resObj *AcceptJoinResponse, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AcceptJoinResponse
func (r *AcceptJoinResponseRequest) Update(ctx context.Context, reqObj *AcceptJoinResponse) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AcceptJoinResponse
func (r *AcceptJoinResponseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AlterationResponseRequestBuilder is request builder for AlterationResponse
type AlterationResponseRequestBuilder struct{ BaseRequestBuilder }

// Request returns AlterationResponseRequest
func (b *AlterationResponseRequestBuilder) Request() *AlterationResponseRequest {
	return &AlterationResponseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AlterationResponseRequest is request for AlterationResponse
type AlterationResponseRequest struct{ BaseRequest }

// Get performs GET request for AlterationResponse
func (r *AlterationResponseRequest) Get(ctx context.Context) (resObj *AlterationResponse, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AlterationResponse
func (r *AlterationResponseRequest) Update(ctx context.Context, reqObj *AlterationResponse) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AlterationResponse
func (r *AlterationResponseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

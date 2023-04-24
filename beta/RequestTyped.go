// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TypedEmailAddressRequestBuilder is request builder for TypedEmailAddress
type TypedEmailAddressRequestBuilder struct{ BaseRequestBuilder }

// Request returns TypedEmailAddressRequest
func (b *TypedEmailAddressRequestBuilder) Request() *TypedEmailAddressRequest {
	return &TypedEmailAddressRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TypedEmailAddressRequest is request for TypedEmailAddress
type TypedEmailAddressRequest struct{ BaseRequest }

// Get performs GET request for TypedEmailAddress
func (r *TypedEmailAddressRequest) Get(ctx context.Context) (resObj *TypedEmailAddress, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TypedEmailAddress
func (r *TypedEmailAddressRequest) Update(ctx context.Context, reqObj *TypedEmailAddress) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TypedEmailAddress
func (r *TypedEmailAddressRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

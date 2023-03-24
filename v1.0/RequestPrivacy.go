// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PrivacyRequestBuilder is request builder for Privacy
type PrivacyRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivacyRequest
func (b *PrivacyRequestBuilder) Request() *PrivacyRequest {
	return &PrivacyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivacyRequest is request for Privacy
type PrivacyRequest struct{ BaseRequest }

// Get performs GET request for Privacy
func (r *PrivacyRequest) Get(ctx context.Context) (resObj *Privacy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Privacy
func (r *PrivacyRequest) Update(ctx context.Context, reqObj *Privacy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Privacy
func (r *PrivacyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

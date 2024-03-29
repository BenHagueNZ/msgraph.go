// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ParentLabelDetailsRequestBuilder is request builder for ParentLabelDetails
type ParentLabelDetailsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ParentLabelDetailsRequest
func (b *ParentLabelDetailsRequestBuilder) Request() *ParentLabelDetailsRequest {
	return &ParentLabelDetailsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ParentLabelDetailsRequest is request for ParentLabelDetails
type ParentLabelDetailsRequest struct{ BaseRequest }

// Get performs GET request for ParentLabelDetails
func (r *ParentLabelDetailsRequest) Get(ctx context.Context) (resObj *ParentLabelDetails, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ParentLabelDetails
func (r *ParentLabelDetailsRequest) Update(ctx context.Context, reqObj *ParentLabelDetails) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ParentLabelDetails
func (r *ParentLabelDetailsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

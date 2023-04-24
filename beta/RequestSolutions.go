// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SolutionsRootRequestBuilder is request builder for SolutionsRoot
type SolutionsRootRequestBuilder struct{ BaseRequestBuilder }

// Request returns SolutionsRootRequest
func (b *SolutionsRootRequestBuilder) Request() *SolutionsRootRequest {
	return &SolutionsRootRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SolutionsRootRequest is request for SolutionsRoot
type SolutionsRootRequest struct{ BaseRequest }

// Get performs GET request for SolutionsRoot
func (r *SolutionsRootRequest) Get(ctx context.Context) (resObj *SolutionsRoot, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SolutionsRoot
func (r *SolutionsRootRequest) Update(ctx context.Context, reqObj *SolutionsRoot) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SolutionsRoot
func (r *SolutionsRootRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
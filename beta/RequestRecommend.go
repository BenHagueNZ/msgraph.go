// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// RecommendLabelActionRequestBuilder is request builder for RecommendLabelAction
type RecommendLabelActionRequestBuilder struct{ BaseRequestBuilder }

// Request returns RecommendLabelActionRequest
func (b *RecommendLabelActionRequestBuilder) Request() *RecommendLabelActionRequest {
	return &RecommendLabelActionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RecommendLabelActionRequest is request for RecommendLabelAction
type RecommendLabelActionRequest struct{ BaseRequest }

// Get performs GET request for RecommendLabelAction
func (r *RecommendLabelActionRequest) Get(ctx context.Context) (resObj *RecommendLabelAction, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RecommendLabelAction
func (r *RecommendLabelActionRequest) Update(ctx context.Context, reqObj *RecommendLabelAction) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RecommendLabelAction
func (r *RecommendLabelActionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

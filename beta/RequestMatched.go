// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MatchedConditionRequestBuilder is request builder for MatchedCondition
type MatchedConditionRequestBuilder struct{ BaseRequestBuilder }

// Request returns MatchedConditionRequest
func (b *MatchedConditionRequestBuilder) Request() *MatchedConditionRequest {
	return &MatchedConditionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MatchedConditionRequest is request for MatchedCondition
type MatchedConditionRequest struct{ BaseRequest }

// Get performs GET request for MatchedCondition
func (r *MatchedConditionRequest) Get(ctx context.Context) (resObj *MatchedCondition, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MatchedCondition
func (r *MatchedConditionRequest) Update(ctx context.Context, reqObj *MatchedCondition) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MatchedCondition
func (r *MatchedConditionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

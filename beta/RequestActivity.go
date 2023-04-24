// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ActivityBasedTimeoutPolicyRequestBuilder is request builder for ActivityBasedTimeoutPolicy
type ActivityBasedTimeoutPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns ActivityBasedTimeoutPolicyRequest
func (b *ActivityBasedTimeoutPolicyRequestBuilder) Request() *ActivityBasedTimeoutPolicyRequest {
	return &ActivityBasedTimeoutPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ActivityBasedTimeoutPolicyRequest is request for ActivityBasedTimeoutPolicy
type ActivityBasedTimeoutPolicyRequest struct{ BaseRequest }

// Get performs GET request for ActivityBasedTimeoutPolicy
func (r *ActivityBasedTimeoutPolicyRequest) Get(ctx context.Context) (resObj *ActivityBasedTimeoutPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ActivityBasedTimeoutPolicy
func (r *ActivityBasedTimeoutPolicyRequest) Update(ctx context.Context, reqObj *ActivityBasedTimeoutPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ActivityBasedTimeoutPolicy
func (r *ActivityBasedTimeoutPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ActivityHistoryItemRequestBuilder is request builder for ActivityHistoryItem
type ActivityHistoryItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns ActivityHistoryItemRequest
func (b *ActivityHistoryItemRequestBuilder) Request() *ActivityHistoryItemRequest {
	return &ActivityHistoryItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ActivityHistoryItemRequest is request for ActivityHistoryItem
type ActivityHistoryItemRequest struct{ BaseRequest }

// Get performs GET request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Get(ctx context.Context) (resObj *ActivityHistoryItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Update(ctx context.Context, reqObj *ActivityHistoryItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ActivityStatisticsRequestBuilder is request builder for ActivityStatistics
type ActivityStatisticsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ActivityStatisticsRequest
func (b *ActivityStatisticsRequestBuilder) Request() *ActivityStatisticsRequest {
	return &ActivityStatisticsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ActivityStatisticsRequest is request for ActivityStatistics
type ActivityStatisticsRequest struct{ BaseRequest }

// Get performs GET request for ActivityStatistics
func (r *ActivityStatisticsRequest) Get(ctx context.Context) (resObj *ActivityStatistics, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ActivityStatistics
func (r *ActivityStatisticsRequest) Update(ctx context.Context, reqObj *ActivityStatistics) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ActivityStatistics
func (r *ActivityStatisticsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
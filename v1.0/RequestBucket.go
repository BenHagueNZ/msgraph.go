// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// BucketAggregationDefinitionRequestBuilder is request builder for BucketAggregationDefinition
type BucketAggregationDefinitionRequestBuilder struct{ BaseRequestBuilder }

// Request returns BucketAggregationDefinitionRequest
func (b *BucketAggregationDefinitionRequestBuilder) Request() *BucketAggregationDefinitionRequest {
	return &BucketAggregationDefinitionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BucketAggregationDefinitionRequest is request for BucketAggregationDefinition
type BucketAggregationDefinitionRequest struct{ BaseRequest }

// Get performs GET request for BucketAggregationDefinition
func (r *BucketAggregationDefinitionRequest) Get(ctx context.Context) (resObj *BucketAggregationDefinition, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BucketAggregationDefinition
func (r *BucketAggregationDefinitionRequest) Update(ctx context.Context, reqObj *BucketAggregationDefinition) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BucketAggregationDefinition
func (r *BucketAggregationDefinitionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BucketAggregationRangeRequestBuilder is request builder for BucketAggregationRange
type BucketAggregationRangeRequestBuilder struct{ BaseRequestBuilder }

// Request returns BucketAggregationRangeRequest
func (b *BucketAggregationRangeRequestBuilder) Request() *BucketAggregationRangeRequest {
	return &BucketAggregationRangeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BucketAggregationRangeRequest is request for BucketAggregationRange
type BucketAggregationRangeRequest struct{ BaseRequest }

// Get performs GET request for BucketAggregationRange
func (r *BucketAggregationRangeRequest) Get(ctx context.Context) (resObj *BucketAggregationRange, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BucketAggregationRange
func (r *BucketAggregationRangeRequest) Update(ctx context.Context, reqObj *BucketAggregationRange) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BucketAggregationRange
func (r *BucketAggregationRangeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
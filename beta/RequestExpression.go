// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ExpressionEvaluationDetailsRequestBuilder is request builder for ExpressionEvaluationDetails
type ExpressionEvaluationDetailsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExpressionEvaluationDetailsRequest
func (b *ExpressionEvaluationDetailsRequestBuilder) Request() *ExpressionEvaluationDetailsRequest {
	return &ExpressionEvaluationDetailsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExpressionEvaluationDetailsRequest is request for ExpressionEvaluationDetails
type ExpressionEvaluationDetailsRequest struct{ BaseRequest }

// Get performs GET request for ExpressionEvaluationDetails
func (r *ExpressionEvaluationDetailsRequest) Get(ctx context.Context) (resObj *ExpressionEvaluationDetails, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExpressionEvaluationDetails
func (r *ExpressionEvaluationDetailsRequest) Update(ctx context.Context, reqObj *ExpressionEvaluationDetails) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExpressionEvaluationDetails
func (r *ExpressionEvaluationDetailsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ExpressionInputObjectRequestBuilder is request builder for ExpressionInputObject
type ExpressionInputObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExpressionInputObjectRequest
func (b *ExpressionInputObjectRequestBuilder) Request() *ExpressionInputObjectRequest {
	return &ExpressionInputObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExpressionInputObjectRequest is request for ExpressionInputObject
type ExpressionInputObjectRequest struct{ BaseRequest }

// Get performs GET request for ExpressionInputObject
func (r *ExpressionInputObjectRequest) Get(ctx context.Context) (resObj *ExpressionInputObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExpressionInputObject
func (r *ExpressionInputObjectRequest) Update(ctx context.Context, reqObj *ExpressionInputObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExpressionInputObject
func (r *ExpressionInputObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// OperationRequestBuilder is request builder for Operation
type OperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns OperationRequest
func (b *OperationRequestBuilder) Request() *OperationRequest {
	return &OperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OperationRequest is request for Operation
type OperationRequest struct{ BaseRequest }

// Get performs GET request for Operation
func (r *OperationRequest) Get(ctx context.Context) (resObj *Operation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Operation
func (r *OperationRequest) Update(ctx context.Context, reqObj *Operation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Operation
func (r *OperationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OperationApprovalPolicySetRequestBuilder is request builder for OperationApprovalPolicySet
type OperationApprovalPolicySetRequestBuilder struct{ BaseRequestBuilder }

// Request returns OperationApprovalPolicySetRequest
func (b *OperationApprovalPolicySetRequestBuilder) Request() *OperationApprovalPolicySetRequest {
	return &OperationApprovalPolicySetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OperationApprovalPolicySetRequest is request for OperationApprovalPolicySet
type OperationApprovalPolicySetRequest struct{ BaseRequest }

// Get performs GET request for OperationApprovalPolicySet
func (r *OperationApprovalPolicySetRequest) Get(ctx context.Context) (resObj *OperationApprovalPolicySet, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OperationApprovalPolicySet
func (r *OperationApprovalPolicySetRequest) Update(ctx context.Context, reqObj *OperationApprovalPolicySet) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OperationApprovalPolicySet
func (r *OperationApprovalPolicySetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OperationErrorRequestBuilder is request builder for OperationError
type OperationErrorRequestBuilder struct{ BaseRequestBuilder }

// Request returns OperationErrorRequest
func (b *OperationErrorRequestBuilder) Request() *OperationErrorRequest {
	return &OperationErrorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OperationErrorRequest is request for OperationError
type OperationErrorRequest struct{ BaseRequest }

// Get performs GET request for OperationError
func (r *OperationErrorRequest) Get(ctx context.Context) (resObj *OperationError, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OperationError
func (r *OperationErrorRequest) Update(ctx context.Context, reqObj *OperationError) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OperationError
func (r *OperationErrorRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
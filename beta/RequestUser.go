// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// UserRequestBuilder is request builder for User
type UserRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserRequest
func (b *UserRequestBuilder) Request() *UserRequest {
	return &UserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserRequest is request for User
type UserRequest struct{ BaseRequest }

// Get performs GET request for User
func (r *UserRequest) Get(ctx context.Context) (resObj *User, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for User
func (r *UserRequest) Update(ctx context.Context, reqObj *User) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for User
func (r *UserRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserProcessingResultRequestBuilder is request builder for UserProcessingResult
type UserProcessingResultRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserProcessingResultRequest
func (b *UserProcessingResultRequestBuilder) Request() *UserProcessingResultRequest {
	return &UserProcessingResultRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserProcessingResultRequest is request for UserProcessingResult
type UserProcessingResultRequest struct{ BaseRequest }

// Get performs GET request for UserProcessingResult
func (r *UserProcessingResultRequest) Get(ctx context.Context) (resObj *UserProcessingResult, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserProcessingResult
func (r *UserProcessingResultRequest) Update(ctx context.Context, reqObj *UserProcessingResult) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserProcessingResult
func (r *UserProcessingResultRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserSummaryRequestBuilder is request builder for UserSummary
type UserSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserSummaryRequest
func (b *UserSummaryRequestBuilder) Request() *UserSummaryRequest {
	return &UserSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserSummaryRequest is request for UserSummary
type UserSummaryRequest struct{ BaseRequest }

// Get performs GET request for UserSummary
func (r *UserSummaryRequest) Get(ctx context.Context) (resObj *UserSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserSummary
func (r *UserSummaryRequest) Update(ctx context.Context, reqObj *UserSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserSummary
func (r *UserSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AssignmentFilterEvaluateRequestObjectRequestBuilder is request builder for AssignmentFilterEvaluateRequestObject
type AssignmentFilterEvaluateRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns AssignmentFilterEvaluateRequestObjectRequest
func (b *AssignmentFilterEvaluateRequestObjectRequestBuilder) Request() *AssignmentFilterEvaluateRequestObjectRequest {
	return &AssignmentFilterEvaluateRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AssignmentFilterEvaluateRequestObjectRequest is request for AssignmentFilterEvaluateRequestObject
type AssignmentFilterEvaluateRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for AssignmentFilterEvaluateRequestObject
func (r *AssignmentFilterEvaluateRequestObjectRequest) Get(ctx context.Context) (resObj *AssignmentFilterEvaluateRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AssignmentFilterEvaluateRequestObject
func (r *AssignmentFilterEvaluateRequestObjectRequest) Update(ctx context.Context, reqObj *AssignmentFilterEvaluateRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AssignmentFilterEvaluateRequestObject
func (r *AssignmentFilterEvaluateRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AssignmentFilterEvaluationStatusDetailsRequestBuilder is request builder for AssignmentFilterEvaluationStatusDetails
type AssignmentFilterEvaluationStatusDetailsRequestBuilder struct{ BaseRequestBuilder }

// Request returns AssignmentFilterEvaluationStatusDetailsRequest
func (b *AssignmentFilterEvaluationStatusDetailsRequestBuilder) Request() *AssignmentFilterEvaluationStatusDetailsRequest {
	return &AssignmentFilterEvaluationStatusDetailsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AssignmentFilterEvaluationStatusDetailsRequest is request for AssignmentFilterEvaluationStatusDetails
type AssignmentFilterEvaluationStatusDetailsRequest struct{ BaseRequest }

// Get performs GET request for AssignmentFilterEvaluationStatusDetails
func (r *AssignmentFilterEvaluationStatusDetailsRequest) Get(ctx context.Context) (resObj *AssignmentFilterEvaluationStatusDetails, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AssignmentFilterEvaluationStatusDetails
func (r *AssignmentFilterEvaluationStatusDetailsRequest) Update(ctx context.Context, reqObj *AssignmentFilterEvaluationStatusDetails) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AssignmentFilterEvaluationStatusDetails
func (r *AssignmentFilterEvaluationStatusDetailsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AssignmentFilterEvaluationSummaryRequestBuilder is request builder for AssignmentFilterEvaluationSummary
type AssignmentFilterEvaluationSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns AssignmentFilterEvaluationSummaryRequest
func (b *AssignmentFilterEvaluationSummaryRequestBuilder) Request() *AssignmentFilterEvaluationSummaryRequest {
	return &AssignmentFilterEvaluationSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AssignmentFilterEvaluationSummaryRequest is request for AssignmentFilterEvaluationSummary
type AssignmentFilterEvaluationSummaryRequest struct{ BaseRequest }

// Get performs GET request for AssignmentFilterEvaluationSummary
func (r *AssignmentFilterEvaluationSummaryRequest) Get(ctx context.Context) (resObj *AssignmentFilterEvaluationSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AssignmentFilterEvaluationSummary
func (r *AssignmentFilterEvaluationSummaryRequest) Update(ctx context.Context, reqObj *AssignmentFilterEvaluationSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AssignmentFilterEvaluationSummary
func (r *AssignmentFilterEvaluationSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AssignmentFilterStateRequestBuilder is request builder for AssignmentFilterState
type AssignmentFilterStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns AssignmentFilterStateRequest
func (b *AssignmentFilterStateRequestBuilder) Request() *AssignmentFilterStateRequest {
	return &AssignmentFilterStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AssignmentFilterStateRequest is request for AssignmentFilterState
type AssignmentFilterStateRequest struct{ BaseRequest }

// Get performs GET request for AssignmentFilterState
func (r *AssignmentFilterStateRequest) Get(ctx context.Context) (resObj *AssignmentFilterState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AssignmentFilterState
func (r *AssignmentFilterStateRequest) Update(ctx context.Context, reqObj *AssignmentFilterState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AssignmentFilterState
func (r *AssignmentFilterStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AssignmentFilterStatusDetailsRequestBuilder is request builder for AssignmentFilterStatusDetails
type AssignmentFilterStatusDetailsRequestBuilder struct{ BaseRequestBuilder }

// Request returns AssignmentFilterStatusDetailsRequest
func (b *AssignmentFilterStatusDetailsRequestBuilder) Request() *AssignmentFilterStatusDetailsRequest {
	return &AssignmentFilterStatusDetailsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AssignmentFilterStatusDetailsRequest is request for AssignmentFilterStatusDetails
type AssignmentFilterStatusDetailsRequest struct{ BaseRequest }

// Get performs GET request for AssignmentFilterStatusDetails
func (r *AssignmentFilterStatusDetailsRequest) Get(ctx context.Context) (resObj *AssignmentFilterStatusDetails, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AssignmentFilterStatusDetails
func (r *AssignmentFilterStatusDetailsRequest) Update(ctx context.Context, reqObj *AssignmentFilterStatusDetails) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AssignmentFilterStatusDetails
func (r *AssignmentFilterStatusDetailsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AssignmentFilterSupportedPropertyRequestBuilder is request builder for AssignmentFilterSupportedProperty
type AssignmentFilterSupportedPropertyRequestBuilder struct{ BaseRequestBuilder }

// Request returns AssignmentFilterSupportedPropertyRequest
func (b *AssignmentFilterSupportedPropertyRequestBuilder) Request() *AssignmentFilterSupportedPropertyRequest {
	return &AssignmentFilterSupportedPropertyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AssignmentFilterSupportedPropertyRequest is request for AssignmentFilterSupportedProperty
type AssignmentFilterSupportedPropertyRequest struct{ BaseRequest }

// Get performs GET request for AssignmentFilterSupportedProperty
func (r *AssignmentFilterSupportedPropertyRequest) Get(ctx context.Context) (resObj *AssignmentFilterSupportedProperty, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AssignmentFilterSupportedProperty
func (r *AssignmentFilterSupportedPropertyRequest) Update(ctx context.Context, reqObj *AssignmentFilterSupportedProperty) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AssignmentFilterSupportedProperty
func (r *AssignmentFilterSupportedPropertyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AssignmentFilterTypeAndEvaluationResultRequestBuilder is request builder for AssignmentFilterTypeAndEvaluationResult
type AssignmentFilterTypeAndEvaluationResultRequestBuilder struct{ BaseRequestBuilder }

// Request returns AssignmentFilterTypeAndEvaluationResultRequest
func (b *AssignmentFilterTypeAndEvaluationResultRequestBuilder) Request() *AssignmentFilterTypeAndEvaluationResultRequest {
	return &AssignmentFilterTypeAndEvaluationResultRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AssignmentFilterTypeAndEvaluationResultRequest is request for AssignmentFilterTypeAndEvaluationResult
type AssignmentFilterTypeAndEvaluationResultRequest struct{ BaseRequest }

// Get performs GET request for AssignmentFilterTypeAndEvaluationResult
func (r *AssignmentFilterTypeAndEvaluationResultRequest) Get(ctx context.Context) (resObj *AssignmentFilterTypeAndEvaluationResult, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AssignmentFilterTypeAndEvaluationResult
func (r *AssignmentFilterTypeAndEvaluationResultRequest) Update(ctx context.Context, reqObj *AssignmentFilterTypeAndEvaluationResult) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AssignmentFilterTypeAndEvaluationResult
func (r *AssignmentFilterTypeAndEvaluationResultRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AssignmentFilterValidationResultRequestBuilder is request builder for AssignmentFilterValidationResult
type AssignmentFilterValidationResultRequestBuilder struct{ BaseRequestBuilder }

// Request returns AssignmentFilterValidationResultRequest
func (b *AssignmentFilterValidationResultRequestBuilder) Request() *AssignmentFilterValidationResultRequest {
	return &AssignmentFilterValidationResultRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AssignmentFilterValidationResultRequest is request for AssignmentFilterValidationResult
type AssignmentFilterValidationResultRequest struct{ BaseRequest }

// Get performs GET request for AssignmentFilterValidationResult
func (r *AssignmentFilterValidationResultRequest) Get(ctx context.Context) (resObj *AssignmentFilterValidationResult, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AssignmentFilterValidationResult
func (r *AssignmentFilterValidationResultRequest) Update(ctx context.Context, reqObj *AssignmentFilterValidationResult) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AssignmentFilterValidationResult
func (r *AssignmentFilterValidationResultRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AssignmentOrderRequestBuilder is request builder for AssignmentOrder
type AssignmentOrderRequestBuilder struct{ BaseRequestBuilder }

// Request returns AssignmentOrderRequest
func (b *AssignmentOrderRequestBuilder) Request() *AssignmentOrderRequest {
	return &AssignmentOrderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AssignmentOrderRequest is request for AssignmentOrder
type AssignmentOrderRequest struct{ BaseRequest }

// Get performs GET request for AssignmentOrder
func (r *AssignmentOrderRequest) Get(ctx context.Context) (resObj *AssignmentOrder, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AssignmentOrder
func (r *AssignmentOrderRequest) Update(ctx context.Context, reqObj *AssignmentOrder) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AssignmentOrder
func (r *AssignmentOrderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AssignmentReviewSettingsRequestBuilder is request builder for AssignmentReviewSettings
type AssignmentReviewSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns AssignmentReviewSettingsRequest
func (b *AssignmentReviewSettingsRequestBuilder) Request() *AssignmentReviewSettingsRequest {
	return &AssignmentReviewSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AssignmentReviewSettingsRequest is request for AssignmentReviewSettings
type AssignmentReviewSettingsRequest struct{ BaseRequest }

// Get performs GET request for AssignmentReviewSettings
func (r *AssignmentReviewSettingsRequest) Get(ctx context.Context) (resObj *AssignmentReviewSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AssignmentReviewSettings
func (r *AssignmentReviewSettingsRequest) Update(ctx context.Context, reqObj *AssignmentReviewSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AssignmentReviewSettings
func (r *AssignmentReviewSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

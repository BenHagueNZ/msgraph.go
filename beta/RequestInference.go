// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// InferenceClassificationRequestBuilder is request builder for InferenceClassification
type InferenceClassificationRequestBuilder struct{ BaseRequestBuilder }

// Request returns InferenceClassificationRequest
func (b *InferenceClassificationRequestBuilder) Request() *InferenceClassificationRequest {
	return &InferenceClassificationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// InferenceClassificationRequest is request for InferenceClassification
type InferenceClassificationRequest struct{ BaseRequest }

// Get performs GET request for InferenceClassification
func (r *InferenceClassificationRequest) Get(ctx context.Context) (resObj *InferenceClassification, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for InferenceClassification
func (r *InferenceClassificationRequest) Update(ctx context.Context, reqObj *InferenceClassification) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for InferenceClassification
func (r *InferenceClassificationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// InferenceClassificationOverrideRequestBuilder is request builder for InferenceClassificationOverride
type InferenceClassificationOverrideRequestBuilder struct{ BaseRequestBuilder }

// Request returns InferenceClassificationOverrideRequest
func (b *InferenceClassificationOverrideRequestBuilder) Request() *InferenceClassificationOverrideRequest {
	return &InferenceClassificationOverrideRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// InferenceClassificationOverrideRequest is request for InferenceClassificationOverride
type InferenceClassificationOverrideRequest struct{ BaseRequest }

// Get performs GET request for InferenceClassificationOverride
func (r *InferenceClassificationOverrideRequest) Get(ctx context.Context) (resObj *InferenceClassificationOverride, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for InferenceClassificationOverride
func (r *InferenceClassificationOverrideRequest) Update(ctx context.Context, reqObj *InferenceClassificationOverride) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for InferenceClassificationOverride
func (r *InferenceClassificationOverrideRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// InferenceDataRequestBuilder is request builder for InferenceData
type InferenceDataRequestBuilder struct{ BaseRequestBuilder }

// Request returns InferenceDataRequest
func (b *InferenceDataRequestBuilder) Request() *InferenceDataRequest {
	return &InferenceDataRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// InferenceDataRequest is request for InferenceData
type InferenceDataRequest struct{ BaseRequest }

// Get performs GET request for InferenceData
func (r *InferenceDataRequest) Get(ctx context.Context) (resObj *InferenceData, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for InferenceData
func (r *InferenceDataRequest) Update(ctx context.Context, reqObj *InferenceData) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for InferenceData
func (r *InferenceDataRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
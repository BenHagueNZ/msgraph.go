// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SettingSourceRequestBuilder is request builder for SettingSource
type SettingSourceRequestBuilder struct{ BaseRequestBuilder }

// Request returns SettingSourceRequest
func (b *SettingSourceRequestBuilder) Request() *SettingSourceRequest {
	return &SettingSourceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SettingSourceRequest is request for SettingSource
type SettingSourceRequest struct{ BaseRequest }

// Get performs GET request for SettingSource
func (r *SettingSourceRequest) Get(ctx context.Context) (resObj *SettingSource, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SettingSource
func (r *SettingSourceRequest) Update(ctx context.Context, reqObj *SettingSource) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SettingSource
func (r *SettingSourceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SettingStateDeviceSummaryRequestBuilder is request builder for SettingStateDeviceSummary
type SettingStateDeviceSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns SettingStateDeviceSummaryRequest
func (b *SettingStateDeviceSummaryRequestBuilder) Request() *SettingStateDeviceSummaryRequest {
	return &SettingStateDeviceSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SettingStateDeviceSummaryRequest is request for SettingStateDeviceSummary
type SettingStateDeviceSummaryRequest struct{ BaseRequest }

// Get performs GET request for SettingStateDeviceSummary
func (r *SettingStateDeviceSummaryRequest) Get(ctx context.Context) (resObj *SettingStateDeviceSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SettingStateDeviceSummary
func (r *SettingStateDeviceSummaryRequest) Update(ctx context.Context, reqObj *SettingStateDeviceSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SettingStateDeviceSummary
func (r *SettingStateDeviceSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SettingTemplateValueRequestBuilder is request builder for SettingTemplateValue
type SettingTemplateValueRequestBuilder struct{ BaseRequestBuilder }

// Request returns SettingTemplateValueRequest
func (b *SettingTemplateValueRequestBuilder) Request() *SettingTemplateValueRequest {
	return &SettingTemplateValueRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SettingTemplateValueRequest is request for SettingTemplateValue
type SettingTemplateValueRequest struct{ BaseRequest }

// Get performs GET request for SettingTemplateValue
func (r *SettingTemplateValueRequest) Get(ctx context.Context) (resObj *SettingTemplateValue, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SettingTemplateValue
func (r *SettingTemplateValueRequest) Update(ctx context.Context, reqObj *SettingTemplateValue) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SettingTemplateValue
func (r *SettingTemplateValueRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SettingValueRequestBuilder is request builder for SettingValue
type SettingValueRequestBuilder struct{ BaseRequestBuilder }

// Request returns SettingValueRequest
func (b *SettingValueRequestBuilder) Request() *SettingValueRequest {
	return &SettingValueRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SettingValueRequest is request for SettingValue
type SettingValueRequest struct{ BaseRequest }

// Get performs GET request for SettingValue
func (r *SettingValueRequest) Get(ctx context.Context) (resObj *SettingValue, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SettingValue
func (r *SettingValueRequest) Update(ctx context.Context, reqObj *SettingValue) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SettingValue
func (r *SettingValueRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

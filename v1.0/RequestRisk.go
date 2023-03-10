// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// RiskDetectionRequestBuilder is request builder for RiskDetection
type RiskDetectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns RiskDetectionRequest
func (b *RiskDetectionRequestBuilder) Request() *RiskDetectionRequest {
	return &RiskDetectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RiskDetectionRequest is request for RiskDetection
type RiskDetectionRequest struct{ BaseRequest }

// Get performs GET request for RiskDetection
func (r *RiskDetectionRequest) Get(ctx context.Context) (resObj *RiskDetection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RiskDetection
func (r *RiskDetectionRequest) Update(ctx context.Context, reqObj *RiskDetection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RiskDetection
func (r *RiskDetectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RiskServicePrincipalActivityRequestBuilder is request builder for RiskServicePrincipalActivity
type RiskServicePrincipalActivityRequestBuilder struct{ BaseRequestBuilder }

// Request returns RiskServicePrincipalActivityRequest
func (b *RiskServicePrincipalActivityRequestBuilder) Request() *RiskServicePrincipalActivityRequest {
	return &RiskServicePrincipalActivityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RiskServicePrincipalActivityRequest is request for RiskServicePrincipalActivity
type RiskServicePrincipalActivityRequest struct{ BaseRequest }

// Get performs GET request for RiskServicePrincipalActivity
func (r *RiskServicePrincipalActivityRequest) Get(ctx context.Context) (resObj *RiskServicePrincipalActivity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RiskServicePrincipalActivity
func (r *RiskServicePrincipalActivityRequest) Update(ctx context.Context, reqObj *RiskServicePrincipalActivity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RiskServicePrincipalActivity
func (r *RiskServicePrincipalActivityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RiskUserActivityRequestBuilder is request builder for RiskUserActivity
type RiskUserActivityRequestBuilder struct{ BaseRequestBuilder }

// Request returns RiskUserActivityRequest
func (b *RiskUserActivityRequestBuilder) Request() *RiskUserActivityRequest {
	return &RiskUserActivityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RiskUserActivityRequest is request for RiskUserActivity
type RiskUserActivityRequest struct{ BaseRequest }

// Get performs GET request for RiskUserActivity
func (r *RiskUserActivityRequest) Get(ctx context.Context) (resObj *RiskUserActivity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RiskUserActivity
func (r *RiskUserActivityRequest) Update(ctx context.Context, reqObj *RiskUserActivity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RiskUserActivity
func (r *RiskUserActivityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

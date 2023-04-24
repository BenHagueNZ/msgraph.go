// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// DeliveryOptimizationBandwidthRequestBuilder is request builder for DeliveryOptimizationBandwidth
type DeliveryOptimizationBandwidthRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeliveryOptimizationBandwidthRequest
func (b *DeliveryOptimizationBandwidthRequestBuilder) Request() *DeliveryOptimizationBandwidthRequest {
	return &DeliveryOptimizationBandwidthRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeliveryOptimizationBandwidthRequest is request for DeliveryOptimizationBandwidth
type DeliveryOptimizationBandwidthRequest struct{ BaseRequest }

// Get performs GET request for DeliveryOptimizationBandwidth
func (r *DeliveryOptimizationBandwidthRequest) Get(ctx context.Context) (resObj *DeliveryOptimizationBandwidth, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeliveryOptimizationBandwidth
func (r *DeliveryOptimizationBandwidthRequest) Update(ctx context.Context, reqObj *DeliveryOptimizationBandwidth) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeliveryOptimizationBandwidth
func (r *DeliveryOptimizationBandwidthRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DeliveryOptimizationBandwidthAbsoluteRequestBuilder is request builder for DeliveryOptimizationBandwidthAbsolute
type DeliveryOptimizationBandwidthAbsoluteRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeliveryOptimizationBandwidthAbsoluteRequest
func (b *DeliveryOptimizationBandwidthAbsoluteRequestBuilder) Request() *DeliveryOptimizationBandwidthAbsoluteRequest {
	return &DeliveryOptimizationBandwidthAbsoluteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeliveryOptimizationBandwidthAbsoluteRequest is request for DeliveryOptimizationBandwidthAbsolute
type DeliveryOptimizationBandwidthAbsoluteRequest struct{ BaseRequest }

// Get performs GET request for DeliveryOptimizationBandwidthAbsolute
func (r *DeliveryOptimizationBandwidthAbsoluteRequest) Get(ctx context.Context) (resObj *DeliveryOptimizationBandwidthAbsolute, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeliveryOptimizationBandwidthAbsolute
func (r *DeliveryOptimizationBandwidthAbsoluteRequest) Update(ctx context.Context, reqObj *DeliveryOptimizationBandwidthAbsolute) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeliveryOptimizationBandwidthAbsolute
func (r *DeliveryOptimizationBandwidthAbsoluteRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DeliveryOptimizationBandwidthBusinessHoursLimitRequestBuilder is request builder for DeliveryOptimizationBandwidthBusinessHoursLimit
type DeliveryOptimizationBandwidthBusinessHoursLimitRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeliveryOptimizationBandwidthBusinessHoursLimitRequest
func (b *DeliveryOptimizationBandwidthBusinessHoursLimitRequestBuilder) Request() *DeliveryOptimizationBandwidthBusinessHoursLimitRequest {
	return &DeliveryOptimizationBandwidthBusinessHoursLimitRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeliveryOptimizationBandwidthBusinessHoursLimitRequest is request for DeliveryOptimizationBandwidthBusinessHoursLimit
type DeliveryOptimizationBandwidthBusinessHoursLimitRequest struct{ BaseRequest }

// Get performs GET request for DeliveryOptimizationBandwidthBusinessHoursLimit
func (r *DeliveryOptimizationBandwidthBusinessHoursLimitRequest) Get(ctx context.Context) (resObj *DeliveryOptimizationBandwidthBusinessHoursLimit, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeliveryOptimizationBandwidthBusinessHoursLimit
func (r *DeliveryOptimizationBandwidthBusinessHoursLimitRequest) Update(ctx context.Context, reqObj *DeliveryOptimizationBandwidthBusinessHoursLimit) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeliveryOptimizationBandwidthBusinessHoursLimit
func (r *DeliveryOptimizationBandwidthBusinessHoursLimitRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DeliveryOptimizationBandwidthHoursWithPercentageRequestBuilder is request builder for DeliveryOptimizationBandwidthHoursWithPercentage
type DeliveryOptimizationBandwidthHoursWithPercentageRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeliveryOptimizationBandwidthHoursWithPercentageRequest
func (b *DeliveryOptimizationBandwidthHoursWithPercentageRequestBuilder) Request() *DeliveryOptimizationBandwidthHoursWithPercentageRequest {
	return &DeliveryOptimizationBandwidthHoursWithPercentageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeliveryOptimizationBandwidthHoursWithPercentageRequest is request for DeliveryOptimizationBandwidthHoursWithPercentage
type DeliveryOptimizationBandwidthHoursWithPercentageRequest struct{ BaseRequest }

// Get performs GET request for DeliveryOptimizationBandwidthHoursWithPercentage
func (r *DeliveryOptimizationBandwidthHoursWithPercentageRequest) Get(ctx context.Context) (resObj *DeliveryOptimizationBandwidthHoursWithPercentage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeliveryOptimizationBandwidthHoursWithPercentage
func (r *DeliveryOptimizationBandwidthHoursWithPercentageRequest) Update(ctx context.Context, reqObj *DeliveryOptimizationBandwidthHoursWithPercentage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeliveryOptimizationBandwidthHoursWithPercentage
func (r *DeliveryOptimizationBandwidthHoursWithPercentageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DeliveryOptimizationBandwidthPercentageRequestBuilder is request builder for DeliveryOptimizationBandwidthPercentage
type DeliveryOptimizationBandwidthPercentageRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeliveryOptimizationBandwidthPercentageRequest
func (b *DeliveryOptimizationBandwidthPercentageRequestBuilder) Request() *DeliveryOptimizationBandwidthPercentageRequest {
	return &DeliveryOptimizationBandwidthPercentageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeliveryOptimizationBandwidthPercentageRequest is request for DeliveryOptimizationBandwidthPercentage
type DeliveryOptimizationBandwidthPercentageRequest struct{ BaseRequest }

// Get performs GET request for DeliveryOptimizationBandwidthPercentage
func (r *DeliveryOptimizationBandwidthPercentageRequest) Get(ctx context.Context) (resObj *DeliveryOptimizationBandwidthPercentage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeliveryOptimizationBandwidthPercentage
func (r *DeliveryOptimizationBandwidthPercentageRequest) Update(ctx context.Context, reqObj *DeliveryOptimizationBandwidthPercentage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeliveryOptimizationBandwidthPercentage
func (r *DeliveryOptimizationBandwidthPercentageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DeliveryOptimizationGroupIDCustomRequestBuilder is request builder for DeliveryOptimizationGroupIDCustom
type DeliveryOptimizationGroupIDCustomRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeliveryOptimizationGroupIDCustomRequest
func (b *DeliveryOptimizationGroupIDCustomRequestBuilder) Request() *DeliveryOptimizationGroupIDCustomRequest {
	return &DeliveryOptimizationGroupIDCustomRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeliveryOptimizationGroupIDCustomRequest is request for DeliveryOptimizationGroupIDCustom
type DeliveryOptimizationGroupIDCustomRequest struct{ BaseRequest }

// Get performs GET request for DeliveryOptimizationGroupIDCustom
func (r *DeliveryOptimizationGroupIDCustomRequest) Get(ctx context.Context) (resObj *DeliveryOptimizationGroupIDCustom, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeliveryOptimizationGroupIDCustom
func (r *DeliveryOptimizationGroupIDCustomRequest) Update(ctx context.Context, reqObj *DeliveryOptimizationGroupIDCustom) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeliveryOptimizationGroupIDCustom
func (r *DeliveryOptimizationGroupIDCustomRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DeliveryOptimizationGroupIDSourceRequestBuilder is request builder for DeliveryOptimizationGroupIDSource
type DeliveryOptimizationGroupIDSourceRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeliveryOptimizationGroupIDSourceRequest
func (b *DeliveryOptimizationGroupIDSourceRequestBuilder) Request() *DeliveryOptimizationGroupIDSourceRequest {
	return &DeliveryOptimizationGroupIDSourceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeliveryOptimizationGroupIDSourceRequest is request for DeliveryOptimizationGroupIDSource
type DeliveryOptimizationGroupIDSourceRequest struct{ BaseRequest }

// Get performs GET request for DeliveryOptimizationGroupIDSource
func (r *DeliveryOptimizationGroupIDSourceRequest) Get(ctx context.Context) (resObj *DeliveryOptimizationGroupIDSource, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeliveryOptimizationGroupIDSource
func (r *DeliveryOptimizationGroupIDSourceRequest) Update(ctx context.Context, reqObj *DeliveryOptimizationGroupIDSource) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeliveryOptimizationGroupIDSource
func (r *DeliveryOptimizationGroupIDSourceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DeliveryOptimizationGroupIDSourceOptionsRequestBuilder is request builder for DeliveryOptimizationGroupIDSourceOptions
type DeliveryOptimizationGroupIDSourceOptionsRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeliveryOptimizationGroupIDSourceOptionsRequest
func (b *DeliveryOptimizationGroupIDSourceOptionsRequestBuilder) Request() *DeliveryOptimizationGroupIDSourceOptionsRequest {
	return &DeliveryOptimizationGroupIDSourceOptionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeliveryOptimizationGroupIDSourceOptionsRequest is request for DeliveryOptimizationGroupIDSourceOptions
type DeliveryOptimizationGroupIDSourceOptionsRequest struct{ BaseRequest }

// Get performs GET request for DeliveryOptimizationGroupIDSourceOptions
func (r *DeliveryOptimizationGroupIDSourceOptionsRequest) Get(ctx context.Context) (resObj *DeliveryOptimizationGroupIDSourceOptions, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeliveryOptimizationGroupIDSourceOptions
func (r *DeliveryOptimizationGroupIDSourceOptionsRequest) Update(ctx context.Context, reqObj *DeliveryOptimizationGroupIDSourceOptions) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeliveryOptimizationGroupIDSourceOptions
func (r *DeliveryOptimizationGroupIDSourceOptionsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DeliveryOptimizationMaxCacheSizeRequestBuilder is request builder for DeliveryOptimizationMaxCacheSize
type DeliveryOptimizationMaxCacheSizeRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeliveryOptimizationMaxCacheSizeRequest
func (b *DeliveryOptimizationMaxCacheSizeRequestBuilder) Request() *DeliveryOptimizationMaxCacheSizeRequest {
	return &DeliveryOptimizationMaxCacheSizeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeliveryOptimizationMaxCacheSizeRequest is request for DeliveryOptimizationMaxCacheSize
type DeliveryOptimizationMaxCacheSizeRequest struct{ BaseRequest }

// Get performs GET request for DeliveryOptimizationMaxCacheSize
func (r *DeliveryOptimizationMaxCacheSizeRequest) Get(ctx context.Context) (resObj *DeliveryOptimizationMaxCacheSize, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeliveryOptimizationMaxCacheSize
func (r *DeliveryOptimizationMaxCacheSizeRequest) Update(ctx context.Context, reqObj *DeliveryOptimizationMaxCacheSize) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeliveryOptimizationMaxCacheSize
func (r *DeliveryOptimizationMaxCacheSizeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DeliveryOptimizationMaxCacheSizeAbsoluteRequestBuilder is request builder for DeliveryOptimizationMaxCacheSizeAbsolute
type DeliveryOptimizationMaxCacheSizeAbsoluteRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeliveryOptimizationMaxCacheSizeAbsoluteRequest
func (b *DeliveryOptimizationMaxCacheSizeAbsoluteRequestBuilder) Request() *DeliveryOptimizationMaxCacheSizeAbsoluteRequest {
	return &DeliveryOptimizationMaxCacheSizeAbsoluteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeliveryOptimizationMaxCacheSizeAbsoluteRequest is request for DeliveryOptimizationMaxCacheSizeAbsolute
type DeliveryOptimizationMaxCacheSizeAbsoluteRequest struct{ BaseRequest }

// Get performs GET request for DeliveryOptimizationMaxCacheSizeAbsolute
func (r *DeliveryOptimizationMaxCacheSizeAbsoluteRequest) Get(ctx context.Context) (resObj *DeliveryOptimizationMaxCacheSizeAbsolute, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeliveryOptimizationMaxCacheSizeAbsolute
func (r *DeliveryOptimizationMaxCacheSizeAbsoluteRequest) Update(ctx context.Context, reqObj *DeliveryOptimizationMaxCacheSizeAbsolute) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeliveryOptimizationMaxCacheSizeAbsolute
func (r *DeliveryOptimizationMaxCacheSizeAbsoluteRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DeliveryOptimizationMaxCacheSizePercentageRequestBuilder is request builder for DeliveryOptimizationMaxCacheSizePercentage
type DeliveryOptimizationMaxCacheSizePercentageRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeliveryOptimizationMaxCacheSizePercentageRequest
func (b *DeliveryOptimizationMaxCacheSizePercentageRequestBuilder) Request() *DeliveryOptimizationMaxCacheSizePercentageRequest {
	return &DeliveryOptimizationMaxCacheSizePercentageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeliveryOptimizationMaxCacheSizePercentageRequest is request for DeliveryOptimizationMaxCacheSizePercentage
type DeliveryOptimizationMaxCacheSizePercentageRequest struct{ BaseRequest }

// Get performs GET request for DeliveryOptimizationMaxCacheSizePercentage
func (r *DeliveryOptimizationMaxCacheSizePercentageRequest) Get(ctx context.Context) (resObj *DeliveryOptimizationMaxCacheSizePercentage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeliveryOptimizationMaxCacheSizePercentage
func (r *DeliveryOptimizationMaxCacheSizePercentageRequest) Update(ctx context.Context, reqObj *DeliveryOptimizationMaxCacheSizePercentage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeliveryOptimizationMaxCacheSizePercentage
func (r *DeliveryOptimizationMaxCacheSizePercentageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

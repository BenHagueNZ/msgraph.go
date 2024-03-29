// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// CertificateAuthorityRequestBuilder is request builder for CertificateAuthority
type CertificateAuthorityRequestBuilder struct{ BaseRequestBuilder }

// Request returns CertificateAuthorityRequest
func (b *CertificateAuthorityRequestBuilder) Request() *CertificateAuthorityRequest {
	return &CertificateAuthorityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CertificateAuthorityRequest is request for CertificateAuthority
type CertificateAuthorityRequest struct{ BaseRequest }

// Get performs GET request for CertificateAuthority
func (r *CertificateAuthorityRequest) Get(ctx context.Context) (resObj *CertificateAuthority, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CertificateAuthority
func (r *CertificateAuthorityRequest) Update(ctx context.Context, reqObj *CertificateAuthority) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CertificateAuthority
func (r *CertificateAuthorityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CertificateBasedAuthConfigurationRequestBuilder is request builder for CertificateBasedAuthConfiguration
type CertificateBasedAuthConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns CertificateBasedAuthConfigurationRequest
func (b *CertificateBasedAuthConfigurationRequestBuilder) Request() *CertificateBasedAuthConfigurationRequest {
	return &CertificateBasedAuthConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CertificateBasedAuthConfigurationRequest is request for CertificateBasedAuthConfiguration
type CertificateBasedAuthConfigurationRequest struct{ BaseRequest }

// Get performs GET request for CertificateBasedAuthConfiguration
func (r *CertificateBasedAuthConfigurationRequest) Get(ctx context.Context) (resObj *CertificateBasedAuthConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CertificateBasedAuthConfiguration
func (r *CertificateBasedAuthConfigurationRequest) Update(ctx context.Context, reqObj *CertificateBasedAuthConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CertificateBasedAuthConfiguration
func (r *CertificateBasedAuthConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CertificateConnectorDetailsRequestBuilder is request builder for CertificateConnectorDetails
type CertificateConnectorDetailsRequestBuilder struct{ BaseRequestBuilder }

// Request returns CertificateConnectorDetailsRequest
func (b *CertificateConnectorDetailsRequestBuilder) Request() *CertificateConnectorDetailsRequest {
	return &CertificateConnectorDetailsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CertificateConnectorDetailsRequest is request for CertificateConnectorDetails
type CertificateConnectorDetailsRequest struct{ BaseRequest }

// Get performs GET request for CertificateConnectorDetails
func (r *CertificateConnectorDetailsRequest) Get(ctx context.Context) (resObj *CertificateConnectorDetails, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CertificateConnectorDetails
func (r *CertificateConnectorDetailsRequest) Update(ctx context.Context, reqObj *CertificateConnectorDetails) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CertificateConnectorDetails
func (r *CertificateConnectorDetailsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CertificateConnectorHealthMetricValueRequestBuilder is request builder for CertificateConnectorHealthMetricValue
type CertificateConnectorHealthMetricValueRequestBuilder struct{ BaseRequestBuilder }

// Request returns CertificateConnectorHealthMetricValueRequest
func (b *CertificateConnectorHealthMetricValueRequestBuilder) Request() *CertificateConnectorHealthMetricValueRequest {
	return &CertificateConnectorHealthMetricValueRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CertificateConnectorHealthMetricValueRequest is request for CertificateConnectorHealthMetricValue
type CertificateConnectorHealthMetricValueRequest struct{ BaseRequest }

// Get performs GET request for CertificateConnectorHealthMetricValue
func (r *CertificateConnectorHealthMetricValueRequest) Get(ctx context.Context) (resObj *CertificateConnectorHealthMetricValue, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CertificateConnectorHealthMetricValue
func (r *CertificateConnectorHealthMetricValueRequest) Update(ctx context.Context, reqObj *CertificateConnectorHealthMetricValue) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CertificateConnectorHealthMetricValue
func (r *CertificateConnectorHealthMetricValueRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CertificateConnectorSettingRequestBuilder is request builder for CertificateConnectorSetting
type CertificateConnectorSettingRequestBuilder struct{ BaseRequestBuilder }

// Request returns CertificateConnectorSettingRequest
func (b *CertificateConnectorSettingRequestBuilder) Request() *CertificateConnectorSettingRequest {
	return &CertificateConnectorSettingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CertificateConnectorSettingRequest is request for CertificateConnectorSetting
type CertificateConnectorSettingRequest struct{ BaseRequest }

// Get performs GET request for CertificateConnectorSetting
func (r *CertificateConnectorSettingRequest) Get(ctx context.Context) (resObj *CertificateConnectorSetting, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CertificateConnectorSetting
func (r *CertificateConnectorSettingRequest) Update(ctx context.Context, reqObj *CertificateConnectorSetting) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CertificateConnectorSetting
func (r *CertificateConnectorSettingRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type CertificateConnectorDetailsGetHealthMetricsRequestBuilder struct{ BaseRequestBuilder }

// GetHealthMetrics action undocumentedrac
func (b *CertificateConnectorDetailsRequestBuilder) GetHealthMetrics(reqObj *CertificateConnectorDetailsGetHealthMetricsRequestParameter) *CertificateConnectorDetailsGetHealthMetricsRequestBuilder {
	bb := &CertificateConnectorDetailsGetHealthMetricsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/GetHealthMetrics"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type CertificateConnectorDetailsGetHealthMetricsRequest struct{ BaseRequest }

func (b *CertificateConnectorDetailsGetHealthMetricsRequestBuilder) Request() *CertificateConnectorDetailsGetHealthMetricsRequest {
	return &CertificateConnectorDetailsGetHealthMetricsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *CertificateConnectorDetailsGetHealthMetricsRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]KeyLongValuePair, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []KeyLongValuePair
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []KeyLongValuePair
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, _ = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

func (r *CertificateConnectorDetailsGetHealthMetricsRequest) PostN(ctx context.Context, n int) ([]KeyLongValuePair, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

func (r *CertificateConnectorDetailsGetHealthMetricsRequest) Post(ctx context.Context) ([]KeyLongValuePair, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

type CertificateConnectorDetailsGetHealthMetricTimeSeriesRequestBuilder struct{ BaseRequestBuilder }

// GetHealthMetricTimeSeries action undocumentedrac
func (b *CertificateConnectorDetailsRequestBuilder) GetHealthMetricTimeSeries(reqObj *CertificateConnectorDetailsGetHealthMetricTimeSeriesRequestParameter) *CertificateConnectorDetailsGetHealthMetricTimeSeriesRequestBuilder {
	bb := &CertificateConnectorDetailsGetHealthMetricTimeSeriesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/GetHealthMetricTimeSeries"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type CertificateConnectorDetailsGetHealthMetricTimeSeriesRequest struct{ BaseRequest }

func (b *CertificateConnectorDetailsGetHealthMetricTimeSeriesRequestBuilder) Request() *CertificateConnectorDetailsGetHealthMetricTimeSeriesRequest {
	return &CertificateConnectorDetailsGetHealthMetricTimeSeriesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *CertificateConnectorDetailsGetHealthMetricTimeSeriesRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]CertificateConnectorHealthMetricValue, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []CertificateConnectorHealthMetricValue
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []CertificateConnectorHealthMetricValue
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, _ = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

func (r *CertificateConnectorDetailsGetHealthMetricTimeSeriesRequest) PostN(ctx context.Context, n int) ([]CertificateConnectorHealthMetricValue, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

func (r *CertificateConnectorDetailsGetHealthMetricTimeSeriesRequest) Post(ctx context.Context) ([]CertificateConnectorHealthMetricValue, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// WebPartGetPositionOfWebPartRequestParameter undocumented
type WebPartGetPositionOfWebPartRequestParameter struct {
}

// CorsConfigurations returns request builder for CorsConfiguration_v2 collection
func (b *WebApplicationSegmentRequestBuilder) CorsConfigurations() *WebApplicationSegmentCorsConfigurationsCollectionRequestBuilder {
	bb := &WebApplicationSegmentCorsConfigurationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/corsConfigurations"
	return bb
}

// WebApplicationSegmentCorsConfigurationsCollectionRequestBuilder is request builder for CorsConfiguration_v2 collection rcn
type WebApplicationSegmentCorsConfigurationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CorsConfiguration_v2 collection
func (b *WebApplicationSegmentCorsConfigurationsCollectionRequestBuilder) Request() *WebApplicationSegmentCorsConfigurationsCollectionRequest {
	return &WebApplicationSegmentCorsConfigurationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CorsConfiguration_v2 item
func (b *WebApplicationSegmentCorsConfigurationsCollectionRequestBuilder) ID(id string) *CorsConfiguration_v2RequestBuilder {
	bb := &CorsConfiguration_v2RequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WebApplicationSegmentCorsConfigurationsCollectionRequest is request for CorsConfiguration_v2 collection
type WebApplicationSegmentCorsConfigurationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CorsConfiguration_v2 collection
func (r *WebApplicationSegmentCorsConfigurationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]CorsConfiguration_v2, error) {
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
	var values []CorsConfiguration_v2
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
			value  []CorsConfiguration_v2
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
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for CorsConfiguration_v2 collection, max N pages
func (r *WebApplicationSegmentCorsConfigurationsCollectionRequest) GetN(ctx context.Context, n int) ([]CorsConfiguration_v2, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for CorsConfiguration_v2 collection
func (r *WebApplicationSegmentCorsConfigurationsCollectionRequest) Get(ctx context.Context) ([]CorsConfiguration_v2, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for CorsConfiguration_v2 collection
func (r *WebApplicationSegmentCorsConfigurationsCollectionRequest) Add(ctx context.Context, reqObj *CorsConfiguration_v2) (resObj *CorsConfiguration_v2, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ApplicationSegments returns request builder for WebApplicationSegment collection
func (b *WebSegmentConfigurationRequestBuilder) ApplicationSegments() *WebSegmentConfigurationApplicationSegmentsCollectionRequestBuilder {
	bb := &WebSegmentConfigurationApplicationSegmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/applicationSegments"
	return bb
}

// WebSegmentConfigurationApplicationSegmentsCollectionRequestBuilder is request builder for WebApplicationSegment collection rcn
type WebSegmentConfigurationApplicationSegmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WebApplicationSegment collection
func (b *WebSegmentConfigurationApplicationSegmentsCollectionRequestBuilder) Request() *WebSegmentConfigurationApplicationSegmentsCollectionRequest {
	return &WebSegmentConfigurationApplicationSegmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WebApplicationSegment item
func (b *WebSegmentConfigurationApplicationSegmentsCollectionRequestBuilder) ID(id string) *WebApplicationSegmentRequestBuilder {
	bb := &WebApplicationSegmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WebSegmentConfigurationApplicationSegmentsCollectionRequest is request for WebApplicationSegment collection
type WebSegmentConfigurationApplicationSegmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WebApplicationSegment collection
func (r *WebSegmentConfigurationApplicationSegmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]WebApplicationSegment, error) {
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
	var values []WebApplicationSegment
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
			value  []WebApplicationSegment
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
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for WebApplicationSegment collection, max N pages
func (r *WebSegmentConfigurationApplicationSegmentsCollectionRequest) GetN(ctx context.Context, n int) ([]WebApplicationSegment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for WebApplicationSegment collection
func (r *WebSegmentConfigurationApplicationSegmentsCollectionRequest) Get(ctx context.Context) ([]WebApplicationSegment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for WebApplicationSegment collection
func (r *WebSegmentConfigurationApplicationSegmentsCollectionRequest) Add(ctx context.Context, reqObj *WebApplicationSegment) (resObj *WebApplicationSegment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *WebAccountRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// WebApp returns request builder for WebApp collection
func (b *DeviceAppManagementMobileAppsCollectionRequestBuilder) WebApp() *DeviceAppManagementMobileAppsCollectionWebAppCollectionRequestBuilder {
	bb := &DeviceAppManagementMobileAppsCollectionWebAppCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// DeviceAppManagementMobileAppsCollectionWebAppCollectionRequestBuilder is request builder for WebApp collection rcn
type DeviceAppManagementMobileAppsCollectionWebAppCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WebApp collection
func (b *DeviceAppManagementMobileAppsCollectionWebAppCollectionRequestBuilder) Request() *DeviceAppManagementMobileAppsCollectionWebAppCollectionRequest {
	return &DeviceAppManagementMobileAppsCollectionWebAppCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WebApp item
func (b *DeviceAppManagementMobileAppsCollectionWebAppCollectionRequestBuilder) ID(id string) *WebAppRequestBuilder {
	bb := &WebAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceAppManagementMobileAppsCollectionWebAppCollectionRequest is request for WebApp collection
type DeviceAppManagementMobileAppsCollectionWebAppCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WebApp collection
func (r *DeviceAppManagementMobileAppsCollectionWebAppCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]WebApp, error) {
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
	var values []WebApp
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
			value  []WebApp
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
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for WebApp collection, max N pages
func (r *DeviceAppManagementMobileAppsCollectionWebAppCollectionRequest) GetN(ctx context.Context, n int) ([]WebApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for WebApp collection
func (r *DeviceAppManagementMobileAppsCollectionWebAppCollectionRequest) Get(ctx context.Context) ([]WebApp, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for WebApp collection
func (r *DeviceAppManagementMobileAppsCollectionWebAppCollectionRequest) Add(ctx context.Context, reqObj *WebApp) (resObj *WebApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *WebApplicationSegmentRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

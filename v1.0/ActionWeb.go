// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// WebApp returns request builder for WebApp collection
func (b *DeviceAppManagementMobileAppsCollectionRequestBuilder) WebApp() *DeviceAppManagementMobileAppsCollectionWebAppCollectionRequestBuilder {
	bb := &DeviceAppManagementMobileAppsCollectionWebAppCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// DeviceAppManagementMobileAppsCollectionWebAppCollectionRequestBuilder is request builder for WebApp collection
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

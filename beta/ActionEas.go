// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// EasEmailProfileConfigurationBase returns request builder for EasEmailProfileConfigurationBase collection
func (b *DeviceManagementDeviceConfigurationsCollectionRequestBuilder) EasEmailProfileConfigurationBase() *DeviceManagementDeviceConfigurationsCollectionEasEmailProfileConfigurationBaseCollectionRequestBuilder {
	bb := &DeviceManagementDeviceConfigurationsCollectionEasEmailProfileConfigurationBaseCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// DeviceManagementDeviceConfigurationsCollectionEasEmailProfileConfigurationBaseCollectionRequestBuilder is request builder for EasEmailProfileConfigurationBase collection rcn
type DeviceManagementDeviceConfigurationsCollectionEasEmailProfileConfigurationBaseCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EasEmailProfileConfigurationBase collection
func (b *DeviceManagementDeviceConfigurationsCollectionEasEmailProfileConfigurationBaseCollectionRequestBuilder) Request() *DeviceManagementDeviceConfigurationsCollectionEasEmailProfileConfigurationBaseCollectionRequest {
	return &DeviceManagementDeviceConfigurationsCollectionEasEmailProfileConfigurationBaseCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EasEmailProfileConfigurationBase item
func (b *DeviceManagementDeviceConfigurationsCollectionEasEmailProfileConfigurationBaseCollectionRequestBuilder) ID(id string) *EasEmailProfileConfigurationBaseRequestBuilder {
	bb := &EasEmailProfileConfigurationBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementDeviceConfigurationsCollectionEasEmailProfileConfigurationBaseCollectionRequest is request for EasEmailProfileConfigurationBase collection
type DeviceManagementDeviceConfigurationsCollectionEasEmailProfileConfigurationBaseCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EasEmailProfileConfigurationBase collection
func (r *DeviceManagementDeviceConfigurationsCollectionEasEmailProfileConfigurationBaseCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EasEmailProfileConfigurationBase, error) {
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
	var values []EasEmailProfileConfigurationBase
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
			value  []EasEmailProfileConfigurationBase
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

// GetN performs GET request for EasEmailProfileConfigurationBase collection, max N pages
func (r *DeviceManagementDeviceConfigurationsCollectionEasEmailProfileConfigurationBaseCollectionRequest) GetN(ctx context.Context, n int) ([]EasEmailProfileConfigurationBase, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EasEmailProfileConfigurationBase collection
func (r *DeviceManagementDeviceConfigurationsCollectionEasEmailProfileConfigurationBaseCollectionRequest) Get(ctx context.Context) ([]EasEmailProfileConfigurationBase, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EasEmailProfileConfigurationBase collection
func (r *DeviceManagementDeviceConfigurationsCollectionEasEmailProfileConfigurationBaseCollectionRequest) Add(ctx context.Context, reqObj *EasEmailProfileConfigurationBase) (resObj *EasEmailProfileConfigurationBase, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

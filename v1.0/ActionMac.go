// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// MacOSCompliancePolicy returns request builder for MacOSCompliancePolicy collection
func (b *DeviceManagementDeviceCompliancePoliciesCollectionRequestBuilder) MacOSCompliancePolicy() *DeviceManagementDeviceCompliancePoliciesCollectionMacOSCompliancePolicyCollectionRequestBuilder {
	bb := &DeviceManagementDeviceCompliancePoliciesCollectionMacOSCompliancePolicyCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// DeviceManagementDeviceCompliancePoliciesCollectionMacOSCompliancePolicyCollectionRequestBuilder is request builder for MacOSCompliancePolicy collection rcn
type DeviceManagementDeviceCompliancePoliciesCollectionMacOSCompliancePolicyCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MacOSCompliancePolicy collection
func (b *DeviceManagementDeviceCompliancePoliciesCollectionMacOSCompliancePolicyCollectionRequestBuilder) Request() *DeviceManagementDeviceCompliancePoliciesCollectionMacOSCompliancePolicyCollectionRequest {
	return &DeviceManagementDeviceCompliancePoliciesCollectionMacOSCompliancePolicyCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MacOSCompliancePolicy item
func (b *DeviceManagementDeviceCompliancePoliciesCollectionMacOSCompliancePolicyCollectionRequestBuilder) ID(id string) *MacOSCompliancePolicyRequestBuilder {
	bb := &MacOSCompliancePolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementDeviceCompliancePoliciesCollectionMacOSCompliancePolicyCollectionRequest is request for MacOSCompliancePolicy collection
type DeviceManagementDeviceCompliancePoliciesCollectionMacOSCompliancePolicyCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MacOSCompliancePolicy collection
func (r *DeviceManagementDeviceCompliancePoliciesCollectionMacOSCompliancePolicyCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MacOSCompliancePolicy, error) {
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
	var values []MacOSCompliancePolicy
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
			value  []MacOSCompliancePolicy
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

// GetN performs GET request for MacOSCompliancePolicy collection, max N pages
func (r *DeviceManagementDeviceCompliancePoliciesCollectionMacOSCompliancePolicyCollectionRequest) GetN(ctx context.Context, n int) ([]MacOSCompliancePolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MacOSCompliancePolicy collection
func (r *DeviceManagementDeviceCompliancePoliciesCollectionMacOSCompliancePolicyCollectionRequest) Get(ctx context.Context) ([]MacOSCompliancePolicy, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MacOSCompliancePolicy collection
func (r *DeviceManagementDeviceCompliancePoliciesCollectionMacOSCompliancePolicyCollectionRequest) Add(ctx context.Context, reqObj *MacOSCompliancePolicy) (resObj *MacOSCompliancePolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MacOSCustomConfiguration returns request builder for MacOSCustomConfiguration collection
func (b *DeviceManagementDeviceConfigurationsCollectionRequestBuilder) MacOSCustomConfiguration() *DeviceManagementDeviceConfigurationsCollectionMacOSCustomConfigurationCollectionRequestBuilder {
	bb := &DeviceManagementDeviceConfigurationsCollectionMacOSCustomConfigurationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// DeviceManagementDeviceConfigurationsCollectionMacOSCustomConfigurationCollectionRequestBuilder is request builder for MacOSCustomConfiguration collection rcn
type DeviceManagementDeviceConfigurationsCollectionMacOSCustomConfigurationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MacOSCustomConfiguration collection
func (b *DeviceManagementDeviceConfigurationsCollectionMacOSCustomConfigurationCollectionRequestBuilder) Request() *DeviceManagementDeviceConfigurationsCollectionMacOSCustomConfigurationCollectionRequest {
	return &DeviceManagementDeviceConfigurationsCollectionMacOSCustomConfigurationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MacOSCustomConfiguration item
func (b *DeviceManagementDeviceConfigurationsCollectionMacOSCustomConfigurationCollectionRequestBuilder) ID(id string) *MacOSCustomConfigurationRequestBuilder {
	bb := &MacOSCustomConfigurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementDeviceConfigurationsCollectionMacOSCustomConfigurationCollectionRequest is request for MacOSCustomConfiguration collection
type DeviceManagementDeviceConfigurationsCollectionMacOSCustomConfigurationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MacOSCustomConfiguration collection
func (r *DeviceManagementDeviceConfigurationsCollectionMacOSCustomConfigurationCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MacOSCustomConfiguration, error) {
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
	var values []MacOSCustomConfiguration
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
			value  []MacOSCustomConfiguration
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

// GetN performs GET request for MacOSCustomConfiguration collection, max N pages
func (r *DeviceManagementDeviceConfigurationsCollectionMacOSCustomConfigurationCollectionRequest) GetN(ctx context.Context, n int) ([]MacOSCustomConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MacOSCustomConfiguration collection
func (r *DeviceManagementDeviceConfigurationsCollectionMacOSCustomConfigurationCollectionRequest) Get(ctx context.Context) ([]MacOSCustomConfiguration, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MacOSCustomConfiguration collection
func (r *DeviceManagementDeviceConfigurationsCollectionMacOSCustomConfigurationCollectionRequest) Add(ctx context.Context, reqObj *MacOSCustomConfiguration) (resObj *MacOSCustomConfiguration, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MacOSDeviceFeaturesConfiguration returns request builder for MacOSDeviceFeaturesConfiguration collection
func (b *DeviceManagementDeviceConfigurationsCollectionRequestBuilder) MacOSDeviceFeaturesConfiguration() *DeviceManagementDeviceConfigurationsCollectionMacOSDeviceFeaturesConfigurationCollectionRequestBuilder {
	bb := &DeviceManagementDeviceConfigurationsCollectionMacOSDeviceFeaturesConfigurationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// DeviceManagementDeviceConfigurationsCollectionMacOSDeviceFeaturesConfigurationCollectionRequestBuilder is request builder for MacOSDeviceFeaturesConfiguration collection rcn
type DeviceManagementDeviceConfigurationsCollectionMacOSDeviceFeaturesConfigurationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MacOSDeviceFeaturesConfiguration collection
func (b *DeviceManagementDeviceConfigurationsCollectionMacOSDeviceFeaturesConfigurationCollectionRequestBuilder) Request() *DeviceManagementDeviceConfigurationsCollectionMacOSDeviceFeaturesConfigurationCollectionRequest {
	return &DeviceManagementDeviceConfigurationsCollectionMacOSDeviceFeaturesConfigurationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MacOSDeviceFeaturesConfiguration item
func (b *DeviceManagementDeviceConfigurationsCollectionMacOSDeviceFeaturesConfigurationCollectionRequestBuilder) ID(id string) *MacOSDeviceFeaturesConfigurationRequestBuilder {
	bb := &MacOSDeviceFeaturesConfigurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementDeviceConfigurationsCollectionMacOSDeviceFeaturesConfigurationCollectionRequest is request for MacOSDeviceFeaturesConfiguration collection
type DeviceManagementDeviceConfigurationsCollectionMacOSDeviceFeaturesConfigurationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MacOSDeviceFeaturesConfiguration collection
func (r *DeviceManagementDeviceConfigurationsCollectionMacOSDeviceFeaturesConfigurationCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MacOSDeviceFeaturesConfiguration, error) {
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
	var values []MacOSDeviceFeaturesConfiguration
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
			value  []MacOSDeviceFeaturesConfiguration
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

// GetN performs GET request for MacOSDeviceFeaturesConfiguration collection, max N pages
func (r *DeviceManagementDeviceConfigurationsCollectionMacOSDeviceFeaturesConfigurationCollectionRequest) GetN(ctx context.Context, n int) ([]MacOSDeviceFeaturesConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MacOSDeviceFeaturesConfiguration collection
func (r *DeviceManagementDeviceConfigurationsCollectionMacOSDeviceFeaturesConfigurationCollectionRequest) Get(ctx context.Context) ([]MacOSDeviceFeaturesConfiguration, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MacOSDeviceFeaturesConfiguration collection
func (r *DeviceManagementDeviceConfigurationsCollectionMacOSDeviceFeaturesConfigurationCollectionRequest) Add(ctx context.Context, reqObj *MacOSDeviceFeaturesConfiguration) (resObj *MacOSDeviceFeaturesConfiguration, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MacOSGeneralDeviceConfiguration returns request builder for MacOSGeneralDeviceConfiguration collection
func (b *DeviceManagementDeviceConfigurationsCollectionRequestBuilder) MacOSGeneralDeviceConfiguration() *DeviceManagementDeviceConfigurationsCollectionMacOSGeneralDeviceConfigurationCollectionRequestBuilder {
	bb := &DeviceManagementDeviceConfigurationsCollectionMacOSGeneralDeviceConfigurationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// DeviceManagementDeviceConfigurationsCollectionMacOSGeneralDeviceConfigurationCollectionRequestBuilder is request builder for MacOSGeneralDeviceConfiguration collection rcn
type DeviceManagementDeviceConfigurationsCollectionMacOSGeneralDeviceConfigurationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MacOSGeneralDeviceConfiguration collection
func (b *DeviceManagementDeviceConfigurationsCollectionMacOSGeneralDeviceConfigurationCollectionRequestBuilder) Request() *DeviceManagementDeviceConfigurationsCollectionMacOSGeneralDeviceConfigurationCollectionRequest {
	return &DeviceManagementDeviceConfigurationsCollectionMacOSGeneralDeviceConfigurationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MacOSGeneralDeviceConfiguration item
func (b *DeviceManagementDeviceConfigurationsCollectionMacOSGeneralDeviceConfigurationCollectionRequestBuilder) ID(id string) *MacOSGeneralDeviceConfigurationRequestBuilder {
	bb := &MacOSGeneralDeviceConfigurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementDeviceConfigurationsCollectionMacOSGeneralDeviceConfigurationCollectionRequest is request for MacOSGeneralDeviceConfiguration collection
type DeviceManagementDeviceConfigurationsCollectionMacOSGeneralDeviceConfigurationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MacOSGeneralDeviceConfiguration collection
func (r *DeviceManagementDeviceConfigurationsCollectionMacOSGeneralDeviceConfigurationCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MacOSGeneralDeviceConfiguration, error) {
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
	var values []MacOSGeneralDeviceConfiguration
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
			value  []MacOSGeneralDeviceConfiguration
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

// GetN performs GET request for MacOSGeneralDeviceConfiguration collection, max N pages
func (r *DeviceManagementDeviceConfigurationsCollectionMacOSGeneralDeviceConfigurationCollectionRequest) GetN(ctx context.Context, n int) ([]MacOSGeneralDeviceConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MacOSGeneralDeviceConfiguration collection
func (r *DeviceManagementDeviceConfigurationsCollectionMacOSGeneralDeviceConfigurationCollectionRequest) Get(ctx context.Context) ([]MacOSGeneralDeviceConfiguration, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MacOSGeneralDeviceConfiguration collection
func (r *DeviceManagementDeviceConfigurationsCollectionMacOSGeneralDeviceConfigurationCollectionRequest) Add(ctx context.Context, reqObj *MacOSGeneralDeviceConfiguration) (resObj *MacOSGeneralDeviceConfiguration, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MacOSLobApp returns request builder for MacOSLobApp collection
func (b *DeviceAppManagementMobileAppsCollectionRequestBuilder) MacOSLobApp() *DeviceAppManagementMobileAppsCollectionMacOSLobAppCollectionRequestBuilder {
	bb := &DeviceAppManagementMobileAppsCollectionMacOSLobAppCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// DeviceAppManagementMobileAppsCollectionMacOSLobAppCollectionRequestBuilder is request builder for MacOSLobApp collection rcn
type DeviceAppManagementMobileAppsCollectionMacOSLobAppCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MacOSLobApp collection
func (b *DeviceAppManagementMobileAppsCollectionMacOSLobAppCollectionRequestBuilder) Request() *DeviceAppManagementMobileAppsCollectionMacOSLobAppCollectionRequest {
	return &DeviceAppManagementMobileAppsCollectionMacOSLobAppCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MacOSLobApp item
func (b *DeviceAppManagementMobileAppsCollectionMacOSLobAppCollectionRequestBuilder) ID(id string) *MacOSLobAppRequestBuilder {
	bb := &MacOSLobAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceAppManagementMobileAppsCollectionMacOSLobAppCollectionRequest is request for MacOSLobApp collection
type DeviceAppManagementMobileAppsCollectionMacOSLobAppCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MacOSLobApp collection
func (r *DeviceAppManagementMobileAppsCollectionMacOSLobAppCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MacOSLobApp, error) {
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
	var values []MacOSLobApp
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
			value  []MacOSLobApp
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

// GetN performs GET request for MacOSLobApp collection, max N pages
func (r *DeviceAppManagementMobileAppsCollectionMacOSLobAppCollectionRequest) GetN(ctx context.Context, n int) ([]MacOSLobApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MacOSLobApp collection
func (r *DeviceAppManagementMobileAppsCollectionMacOSLobAppCollectionRequest) Get(ctx context.Context) ([]MacOSLobApp, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MacOSLobApp collection
func (r *DeviceAppManagementMobileAppsCollectionMacOSLobAppCollectionRequest) Add(ctx context.Context, reqObj *MacOSLobApp) (resObj *MacOSLobApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MacOSMicrosoftEdgeApp returns request builder for MacOSMicrosoftEdgeApp collection
func (b *DeviceAppManagementMobileAppsCollectionRequestBuilder) MacOSMicrosoftEdgeApp() *DeviceAppManagementMobileAppsCollectionMacOSMicrosoftEdgeAppCollectionRequestBuilder {
	bb := &DeviceAppManagementMobileAppsCollectionMacOSMicrosoftEdgeAppCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// DeviceAppManagementMobileAppsCollectionMacOSMicrosoftEdgeAppCollectionRequestBuilder is request builder for MacOSMicrosoftEdgeApp collection rcn
type DeviceAppManagementMobileAppsCollectionMacOSMicrosoftEdgeAppCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MacOSMicrosoftEdgeApp collection
func (b *DeviceAppManagementMobileAppsCollectionMacOSMicrosoftEdgeAppCollectionRequestBuilder) Request() *DeviceAppManagementMobileAppsCollectionMacOSMicrosoftEdgeAppCollectionRequest {
	return &DeviceAppManagementMobileAppsCollectionMacOSMicrosoftEdgeAppCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MacOSMicrosoftEdgeApp item
func (b *DeviceAppManagementMobileAppsCollectionMacOSMicrosoftEdgeAppCollectionRequestBuilder) ID(id string) *MacOSMicrosoftEdgeAppRequestBuilder {
	bb := &MacOSMicrosoftEdgeAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceAppManagementMobileAppsCollectionMacOSMicrosoftEdgeAppCollectionRequest is request for MacOSMicrosoftEdgeApp collection
type DeviceAppManagementMobileAppsCollectionMacOSMicrosoftEdgeAppCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MacOSMicrosoftEdgeApp collection
func (r *DeviceAppManagementMobileAppsCollectionMacOSMicrosoftEdgeAppCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MacOSMicrosoftEdgeApp, error) {
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
	var values []MacOSMicrosoftEdgeApp
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
			value  []MacOSMicrosoftEdgeApp
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

// GetN performs GET request for MacOSMicrosoftEdgeApp collection, max N pages
func (r *DeviceAppManagementMobileAppsCollectionMacOSMicrosoftEdgeAppCollectionRequest) GetN(ctx context.Context, n int) ([]MacOSMicrosoftEdgeApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MacOSMicrosoftEdgeApp collection
func (r *DeviceAppManagementMobileAppsCollectionMacOSMicrosoftEdgeAppCollectionRequest) Get(ctx context.Context) ([]MacOSMicrosoftEdgeApp, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MacOSMicrosoftEdgeApp collection
func (r *DeviceAppManagementMobileAppsCollectionMacOSMicrosoftEdgeAppCollectionRequest) Add(ctx context.Context, reqObj *MacOSMicrosoftEdgeApp) (resObj *MacOSMicrosoftEdgeApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MacOSOfficeSuiteApp returns request builder for MacOSOfficeSuiteApp collection
func (b *DeviceAppManagementMobileAppsCollectionRequestBuilder) MacOSOfficeSuiteApp() *DeviceAppManagementMobileAppsCollectionMacOSOfficeSuiteAppCollectionRequestBuilder {
	bb := &DeviceAppManagementMobileAppsCollectionMacOSOfficeSuiteAppCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// DeviceAppManagementMobileAppsCollectionMacOSOfficeSuiteAppCollectionRequestBuilder is request builder for MacOSOfficeSuiteApp collection rcn
type DeviceAppManagementMobileAppsCollectionMacOSOfficeSuiteAppCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MacOSOfficeSuiteApp collection
func (b *DeviceAppManagementMobileAppsCollectionMacOSOfficeSuiteAppCollectionRequestBuilder) Request() *DeviceAppManagementMobileAppsCollectionMacOSOfficeSuiteAppCollectionRequest {
	return &DeviceAppManagementMobileAppsCollectionMacOSOfficeSuiteAppCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MacOSOfficeSuiteApp item
func (b *DeviceAppManagementMobileAppsCollectionMacOSOfficeSuiteAppCollectionRequestBuilder) ID(id string) *MacOSOfficeSuiteAppRequestBuilder {
	bb := &MacOSOfficeSuiteAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceAppManagementMobileAppsCollectionMacOSOfficeSuiteAppCollectionRequest is request for MacOSOfficeSuiteApp collection
type DeviceAppManagementMobileAppsCollectionMacOSOfficeSuiteAppCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MacOSOfficeSuiteApp collection
func (r *DeviceAppManagementMobileAppsCollectionMacOSOfficeSuiteAppCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MacOSOfficeSuiteApp, error) {
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
	var values []MacOSOfficeSuiteApp
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
			value  []MacOSOfficeSuiteApp
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

// GetN performs GET request for MacOSOfficeSuiteApp collection, max N pages
func (r *DeviceAppManagementMobileAppsCollectionMacOSOfficeSuiteAppCollectionRequest) GetN(ctx context.Context, n int) ([]MacOSOfficeSuiteApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MacOSOfficeSuiteApp collection
func (r *DeviceAppManagementMobileAppsCollectionMacOSOfficeSuiteAppCollectionRequest) Get(ctx context.Context) ([]MacOSOfficeSuiteApp, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MacOSOfficeSuiteApp collection
func (r *DeviceAppManagementMobileAppsCollectionMacOSOfficeSuiteAppCollectionRequest) Add(ctx context.Context, reqObj *MacOSOfficeSuiteApp) (resObj *MacOSOfficeSuiteApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

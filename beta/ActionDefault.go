// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// Apps returns request builder for ManagedMobileApp collection
func (b *DefaultManagedAppProtectionRequestBuilder) Apps() *DefaultManagedAppProtectionAppsCollectionRequestBuilder {
	bb := &DefaultManagedAppProtectionAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/apps"
	return bb
}

// DefaultManagedAppProtectionAppsCollectionRequestBuilder is request builder for ManagedMobileApp collection rcn
type DefaultManagedAppProtectionAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedMobileApp collection
func (b *DefaultManagedAppProtectionAppsCollectionRequestBuilder) Request() *DefaultManagedAppProtectionAppsCollectionRequest {
	return &DefaultManagedAppProtectionAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedMobileApp item
func (b *DefaultManagedAppProtectionAppsCollectionRequestBuilder) ID(id string) *ManagedMobileAppRequestBuilder {
	bb := &ManagedMobileAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DefaultManagedAppProtectionAppsCollectionRequest is request for ManagedMobileApp collection
type DefaultManagedAppProtectionAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedMobileApp collection
func (r *DefaultManagedAppProtectionAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ManagedMobileApp, error) {
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
	var values []ManagedMobileApp
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
			value  []ManagedMobileApp
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

// GetN performs GET request for ManagedMobileApp collection, max N pages
func (r *DefaultManagedAppProtectionAppsCollectionRequest) GetN(ctx context.Context, n int) ([]ManagedMobileApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ManagedMobileApp collection
func (r *DefaultManagedAppProtectionAppsCollectionRequest) Get(ctx context.Context) ([]ManagedMobileApp, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ManagedMobileApp collection
func (r *DefaultManagedAppProtectionAppsCollectionRequest) Add(ctx context.Context, reqObj *ManagedMobileApp) (resObj *ManagedMobileApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DeploymentSummary is navigation property rn
func (b *DefaultManagedAppProtectionRequestBuilder) DeploymentSummary() *ManagedAppPolicyDeploymentSummaryRequestBuilder {
	bb := &ManagedAppPolicyDeploymentSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deploymentSummary"
	return bb
}

// DefaultDeviceCompliancePolicy returns request builder for DefaultDeviceCompliancePolicy collection
func (b *DeviceManagementDeviceCompliancePoliciesCollectionRequestBuilder) DefaultDeviceCompliancePolicy() *DeviceManagementDeviceCompliancePoliciesCollectionDefaultDeviceCompliancePolicyCollectionRequestBuilder {
	bb := &DeviceManagementDeviceCompliancePoliciesCollectionDefaultDeviceCompliancePolicyCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// DeviceManagementDeviceCompliancePoliciesCollectionDefaultDeviceCompliancePolicyCollectionRequestBuilder is request builder for DefaultDeviceCompliancePolicy collection rcn
type DeviceManagementDeviceCompliancePoliciesCollectionDefaultDeviceCompliancePolicyCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DefaultDeviceCompliancePolicy collection
func (b *DeviceManagementDeviceCompliancePoliciesCollectionDefaultDeviceCompliancePolicyCollectionRequestBuilder) Request() *DeviceManagementDeviceCompliancePoliciesCollectionDefaultDeviceCompliancePolicyCollectionRequest {
	return &DeviceManagementDeviceCompliancePoliciesCollectionDefaultDeviceCompliancePolicyCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DefaultDeviceCompliancePolicy item
func (b *DeviceManagementDeviceCompliancePoliciesCollectionDefaultDeviceCompliancePolicyCollectionRequestBuilder) ID(id string) *DefaultDeviceCompliancePolicyRequestBuilder {
	bb := &DefaultDeviceCompliancePolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementDeviceCompliancePoliciesCollectionDefaultDeviceCompliancePolicyCollectionRequest is request for DefaultDeviceCompliancePolicy collection
type DeviceManagementDeviceCompliancePoliciesCollectionDefaultDeviceCompliancePolicyCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DefaultDeviceCompliancePolicy collection
func (r *DeviceManagementDeviceCompliancePoliciesCollectionDefaultDeviceCompliancePolicyCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DefaultDeviceCompliancePolicy, error) {
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
	var values []DefaultDeviceCompliancePolicy
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
			value  []DefaultDeviceCompliancePolicy
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

// GetN performs GET request for DefaultDeviceCompliancePolicy collection, max N pages
func (r *DeviceManagementDeviceCompliancePoliciesCollectionDefaultDeviceCompliancePolicyCollectionRequest) GetN(ctx context.Context, n int) ([]DefaultDeviceCompliancePolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DefaultDeviceCompliancePolicy collection
func (r *DeviceManagementDeviceCompliancePoliciesCollectionDefaultDeviceCompliancePolicyCollectionRequest) Get(ctx context.Context) ([]DefaultDeviceCompliancePolicy, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DefaultDeviceCompliancePolicy collection
func (r *DeviceManagementDeviceCompliancePoliciesCollectionDefaultDeviceCompliancePolicyCollectionRequest) Add(ctx context.Context, reqObj *DefaultDeviceCompliancePolicy) (resObj *DefaultDeviceCompliancePolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *DefaultUserRoleOverrideRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
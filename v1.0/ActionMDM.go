// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// MDMWindowsInformationProtectionPolicy returns request builder for MDMWindowsInformationProtectionPolicy collection
func (b *DeviceAppManagementManagedAppPoliciesCollectionRequestBuilder) MDMWindowsInformationProtectionPolicy() *DeviceAppManagementManagedAppPoliciesCollectionMDMWindowsInformationProtectionPolicyCollectionRequestBuilder {
	bb := &DeviceAppManagementManagedAppPoliciesCollectionMDMWindowsInformationProtectionPolicyCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// DeviceAppManagementManagedAppPoliciesCollectionMDMWindowsInformationProtectionPolicyCollectionRequestBuilder is request builder for MDMWindowsInformationProtectionPolicy collection rcn
type DeviceAppManagementManagedAppPoliciesCollectionMDMWindowsInformationProtectionPolicyCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MDMWindowsInformationProtectionPolicy collection
func (b *DeviceAppManagementManagedAppPoliciesCollectionMDMWindowsInformationProtectionPolicyCollectionRequestBuilder) Request() *DeviceAppManagementManagedAppPoliciesCollectionMDMWindowsInformationProtectionPolicyCollectionRequest {
	return &DeviceAppManagementManagedAppPoliciesCollectionMDMWindowsInformationProtectionPolicyCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MDMWindowsInformationProtectionPolicy item
func (b *DeviceAppManagementManagedAppPoliciesCollectionMDMWindowsInformationProtectionPolicyCollectionRequestBuilder) ID(id string) *MDMWindowsInformationProtectionPolicyRequestBuilder {
	bb := &MDMWindowsInformationProtectionPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceAppManagementManagedAppPoliciesCollectionMDMWindowsInformationProtectionPolicyCollectionRequest is request for MDMWindowsInformationProtectionPolicy collection
type DeviceAppManagementManagedAppPoliciesCollectionMDMWindowsInformationProtectionPolicyCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MDMWindowsInformationProtectionPolicy collection
func (r *DeviceAppManagementManagedAppPoliciesCollectionMDMWindowsInformationProtectionPolicyCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MDMWindowsInformationProtectionPolicy, error) {
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
	var values []MDMWindowsInformationProtectionPolicy
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
			value  []MDMWindowsInformationProtectionPolicy
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

// GetN performs GET request for MDMWindowsInformationProtectionPolicy collection, max N pages
func (r *DeviceAppManagementManagedAppPoliciesCollectionMDMWindowsInformationProtectionPolicyCollectionRequest) GetN(ctx context.Context, n int) ([]MDMWindowsInformationProtectionPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MDMWindowsInformationProtectionPolicy collection
func (r *DeviceAppManagementManagedAppPoliciesCollectionMDMWindowsInformationProtectionPolicyCollectionRequest) Get(ctx context.Context) ([]MDMWindowsInformationProtectionPolicy, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MDMWindowsInformationProtectionPolicy collection
func (r *DeviceAppManagementManagedAppPoliciesCollectionMDMWindowsInformationProtectionPolicyCollectionRequest) Add(ctx context.Context, reqObj *MDMWindowsInformationProtectionPolicy) (resObj *MDMWindowsInformationProtectionPolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

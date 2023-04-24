// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// MDMWindowsInformationProtectionPolicyCollectionHasPayloadLinksRequestParameter undocumented
type MDMWindowsInformationProtectionPolicyCollectionHasPayloadLinksRequestParameter struct {
	// PayloadIDs undocumented
	PayloadIDs []string `json:"payloadIds,omitempty"`
}

// MDMWindowsInformationProtectionPolicyPolicySetItem returns request builder for MDMWindowsInformationProtectionPolicyPolicySetItem collection
func (b *PolicySetItemsCollectionRequestBuilder) MDMWindowsInformationProtectionPolicyPolicySetItem() *PolicySetItemsCollectionMDMWindowsInformationProtectionPolicyPolicySetItemCollectionRequestBuilder {
	bb := &PolicySetItemsCollectionMDMWindowsInformationProtectionPolicyPolicySetItemCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// PolicySetItemsCollectionMDMWindowsInformationProtectionPolicyPolicySetItemCollectionRequestBuilder is request builder for MDMWindowsInformationProtectionPolicyPolicySetItem collection rcn
type PolicySetItemsCollectionMDMWindowsInformationProtectionPolicyPolicySetItemCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MDMWindowsInformationProtectionPolicyPolicySetItem collection
func (b *PolicySetItemsCollectionMDMWindowsInformationProtectionPolicyPolicySetItemCollectionRequestBuilder) Request() *PolicySetItemsCollectionMDMWindowsInformationProtectionPolicyPolicySetItemCollectionRequest {
	return &PolicySetItemsCollectionMDMWindowsInformationProtectionPolicyPolicySetItemCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MDMWindowsInformationProtectionPolicyPolicySetItem item
func (b *PolicySetItemsCollectionMDMWindowsInformationProtectionPolicyPolicySetItemCollectionRequestBuilder) ID(id string) *MDMWindowsInformationProtectionPolicyPolicySetItemRequestBuilder {
	bb := &MDMWindowsInformationProtectionPolicyPolicySetItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PolicySetItemsCollectionMDMWindowsInformationProtectionPolicyPolicySetItemCollectionRequest is request for MDMWindowsInformationProtectionPolicyPolicySetItem collection
type PolicySetItemsCollectionMDMWindowsInformationProtectionPolicyPolicySetItemCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MDMWindowsInformationProtectionPolicyPolicySetItem collection
func (r *PolicySetItemsCollectionMDMWindowsInformationProtectionPolicyPolicySetItemCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MDMWindowsInformationProtectionPolicyPolicySetItem, error) {
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
	var values []MDMWindowsInformationProtectionPolicyPolicySetItem
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
			value  []MDMWindowsInformationProtectionPolicyPolicySetItem
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

// GetN performs GET request for MDMWindowsInformationProtectionPolicyPolicySetItem collection, max N pages
func (r *PolicySetItemsCollectionMDMWindowsInformationProtectionPolicyPolicySetItemCollectionRequest) GetN(ctx context.Context, n int) ([]MDMWindowsInformationProtectionPolicyPolicySetItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MDMWindowsInformationProtectionPolicyPolicySetItem collection
func (r *PolicySetItemsCollectionMDMWindowsInformationProtectionPolicyPolicySetItemCollectionRequest) Get(ctx context.Context) ([]MDMWindowsInformationProtectionPolicyPolicySetItem, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MDMWindowsInformationProtectionPolicyPolicySetItem collection
func (r *PolicySetItemsCollectionMDMWindowsInformationProtectionPolicyPolicySetItemCollectionRequest) Add(ctx context.Context, reqObj *MDMWindowsInformationProtectionPolicyPolicySetItem) (resObj *MDMWindowsInformationProtectionPolicyPolicySetItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
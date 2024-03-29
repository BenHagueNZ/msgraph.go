// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// Tenants returns request builder for TenantReference collection
func (b *OutboundSharedUserProfileRequestBuilder) Tenants() *OutboundSharedUserProfileTenantsCollectionRequestBuilder {
	bb := &OutboundSharedUserProfileTenantsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tenants"
	return bb
}

// OutboundSharedUserProfileTenantsCollectionRequestBuilder is request builder for TenantReference collection rcn
type OutboundSharedUserProfileTenantsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TenantReference collection
func (b *OutboundSharedUserProfileTenantsCollectionRequestBuilder) Request() *OutboundSharedUserProfileTenantsCollectionRequest {
	return &OutboundSharedUserProfileTenantsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TenantReference item
func (b *OutboundSharedUserProfileTenantsCollectionRequestBuilder) ID(id string) *TenantReferenceRequestBuilder {
	bb := &TenantReferenceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OutboundSharedUserProfileTenantsCollectionRequest is request for TenantReference collection
type OutboundSharedUserProfileTenantsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TenantReference collection
func (r *OutboundSharedUserProfileTenantsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TenantReference, error) {
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
	var values []TenantReference
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
			value  []TenantReference
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

// GetN performs GET request for TenantReference collection, max N pages
func (r *OutboundSharedUserProfileTenantsCollectionRequest) GetN(ctx context.Context, n int) ([]TenantReference, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TenantReference collection
func (r *OutboundSharedUserProfileTenantsCollectionRequest) Get(ctx context.Context) ([]TenantReference, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TenantReference collection
func (r *OutboundSharedUserProfileTenantsCollectionRequest) Add(ctx context.Context, reqObj *TenantReference) (resObj *TenantReference, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// DelegatedAdminCustomers returns request builder for DelegatedAdminCustomer collection
func (b *TenantRelationshipRequestBuilder) DelegatedAdminCustomers() *TenantRelationshipDelegatedAdminCustomersCollectionRequestBuilder {
	bb := &TenantRelationshipDelegatedAdminCustomersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/delegatedAdminCustomers"
	return bb
}

// TenantRelationshipDelegatedAdminCustomersCollectionRequestBuilder is request builder for DelegatedAdminCustomer collection rcn
type TenantRelationshipDelegatedAdminCustomersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DelegatedAdminCustomer collection
func (b *TenantRelationshipDelegatedAdminCustomersCollectionRequestBuilder) Request() *TenantRelationshipDelegatedAdminCustomersCollectionRequest {
	return &TenantRelationshipDelegatedAdminCustomersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DelegatedAdminCustomer item
func (b *TenantRelationshipDelegatedAdminCustomersCollectionRequestBuilder) ID(id string) *DelegatedAdminCustomerRequestBuilder {
	bb := &DelegatedAdminCustomerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TenantRelationshipDelegatedAdminCustomersCollectionRequest is request for DelegatedAdminCustomer collection
type TenantRelationshipDelegatedAdminCustomersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DelegatedAdminCustomer collection
func (r *TenantRelationshipDelegatedAdminCustomersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DelegatedAdminCustomer, error) {
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
	var values []DelegatedAdminCustomer
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
			value  []DelegatedAdminCustomer
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

// GetN performs GET request for DelegatedAdminCustomer collection, max N pages
func (r *TenantRelationshipDelegatedAdminCustomersCollectionRequest) GetN(ctx context.Context, n int) ([]DelegatedAdminCustomer, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DelegatedAdminCustomer collection
func (r *TenantRelationshipDelegatedAdminCustomersCollectionRequest) Get(ctx context.Context) ([]DelegatedAdminCustomer, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DelegatedAdminCustomer collection
func (r *TenantRelationshipDelegatedAdminCustomersCollectionRequest) Add(ctx context.Context, reqObj *DelegatedAdminCustomer) (resObj *DelegatedAdminCustomer, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DelegatedAdminRelationships returns request builder for DelegatedAdminRelationship collection
func (b *TenantRelationshipRequestBuilder) DelegatedAdminRelationships() *TenantRelationshipDelegatedAdminRelationshipsCollectionRequestBuilder {
	bb := &TenantRelationshipDelegatedAdminRelationshipsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/delegatedAdminRelationships"
	return bb
}

// TenantRelationshipDelegatedAdminRelationshipsCollectionRequestBuilder is request builder for DelegatedAdminRelationship collection rcn
type TenantRelationshipDelegatedAdminRelationshipsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DelegatedAdminRelationship collection
func (b *TenantRelationshipDelegatedAdminRelationshipsCollectionRequestBuilder) Request() *TenantRelationshipDelegatedAdminRelationshipsCollectionRequest {
	return &TenantRelationshipDelegatedAdminRelationshipsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DelegatedAdminRelationship item
func (b *TenantRelationshipDelegatedAdminRelationshipsCollectionRequestBuilder) ID(id string) *DelegatedAdminRelationshipRequestBuilder {
	bb := &DelegatedAdminRelationshipRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TenantRelationshipDelegatedAdminRelationshipsCollectionRequest is request for DelegatedAdminRelationship collection
type TenantRelationshipDelegatedAdminRelationshipsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DelegatedAdminRelationship collection
func (r *TenantRelationshipDelegatedAdminRelationshipsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DelegatedAdminRelationship, error) {
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
	var values []DelegatedAdminRelationship
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
			value  []DelegatedAdminRelationship
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

// GetN performs GET request for DelegatedAdminRelationship collection, max N pages
func (r *TenantRelationshipDelegatedAdminRelationshipsCollectionRequest) GetN(ctx context.Context, n int) ([]DelegatedAdminRelationship, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DelegatedAdminRelationship collection
func (r *TenantRelationshipDelegatedAdminRelationshipsCollectionRequest) Get(ctx context.Context) ([]DelegatedAdminRelationship, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DelegatedAdminRelationship collection
func (r *TenantRelationshipDelegatedAdminRelationshipsCollectionRequest) Add(ctx context.Context, reqObj *DelegatedAdminRelationship) (resObj *DelegatedAdminRelationship, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

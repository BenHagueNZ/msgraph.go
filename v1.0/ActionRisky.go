// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// RiskyServicePrincipalCollectionConfirmCompromisedRequestParameter undocumented
type RiskyServicePrincipalCollectionConfirmCompromisedRequestParameter struct {
	// ServicePrincipalIDs undocumented
	ServicePrincipalIDs []string `json:"servicePrincipalIds,omitempty"`
}

// RiskyServicePrincipalCollectionDismissRequestParameter undocumented
type RiskyServicePrincipalCollectionDismissRequestParameter struct {
	// ServicePrincipalIDs undocumented
	ServicePrincipalIDs []string `json:"servicePrincipalIds,omitempty"`
}

// RiskyUserCollectionConfirmCompromisedRequestParameter undocumented
type RiskyUserCollectionConfirmCompromisedRequestParameter struct {
	// UserIDs undocumented
	UserIDs []string `json:"userIds,omitempty"`
}

// RiskyUserCollectionDismissRequestParameter undocumented
type RiskyUserCollectionDismissRequestParameter struct {
	// UserIDs undocumented
	UserIDs []string `json:"userIds,omitempty"`
}

// History returns request builder for RiskyServicePrincipalHistoryItem collection rcn
func (b *RiskyServicePrincipalRequestBuilder) History() *RiskyServicePrincipalHistoryCollectionRequestBuilder {
	bb := &RiskyServicePrincipalHistoryCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/history"
	return bb
}

// RiskyServicePrincipalHistoryCollectionRequestBuilder is request builder for RiskyServicePrincipalHistoryItem collection
type RiskyServicePrincipalHistoryCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for RiskyServicePrincipalHistoryItem collection
func (b *RiskyServicePrincipalHistoryCollectionRequestBuilder) Request() *RiskyServicePrincipalHistoryCollectionRequest {
	return &RiskyServicePrincipalHistoryCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for RiskyServicePrincipalHistoryItem item
func (b *RiskyServicePrincipalHistoryCollectionRequestBuilder) ID(id string) *RiskyServicePrincipalHistoryItemRequestBuilder {
	bb := &RiskyServicePrincipalHistoryItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RiskyServicePrincipalHistoryCollectionRequest is request for RiskyServicePrincipalHistoryItem collection
type RiskyServicePrincipalHistoryCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for RiskyServicePrincipalHistoryItem collection
func (r *RiskyServicePrincipalHistoryCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]RiskyServicePrincipalHistoryItem, error) {
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
	var values []RiskyServicePrincipalHistoryItem
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
			value  []RiskyServicePrincipalHistoryItem
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

// GetN performs GET request for RiskyServicePrincipalHistoryItem collection, max N pages
func (r *RiskyServicePrincipalHistoryCollectionRequest) GetN(ctx context.Context, n int) ([]RiskyServicePrincipalHistoryItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for RiskyServicePrincipalHistoryItem collection
func (r *RiskyServicePrincipalHistoryCollectionRequest) Get(ctx context.Context) ([]RiskyServicePrincipalHistoryItem, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for RiskyServicePrincipalHistoryItem collection
func (r *RiskyServicePrincipalHistoryCollectionRequest) Add(ctx context.Context, reqObj *RiskyServicePrincipalHistoryItem) (resObj *RiskyServicePrincipalHistoryItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// History returns request builder for RiskyUserHistoryItem collection rcn
func (b *RiskyUserRequestBuilder) History() *RiskyUserHistoryCollectionRequestBuilder {
	bb := &RiskyUserHistoryCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/history"
	return bb
}

// RiskyUserHistoryCollectionRequestBuilder is request builder for RiskyUserHistoryItem collection
type RiskyUserHistoryCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for RiskyUserHistoryItem collection
func (b *RiskyUserHistoryCollectionRequestBuilder) Request() *RiskyUserHistoryCollectionRequest {
	return &RiskyUserHistoryCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for RiskyUserHistoryItem item
func (b *RiskyUserHistoryCollectionRequestBuilder) ID(id string) *RiskyUserHistoryItemRequestBuilder {
	bb := &RiskyUserHistoryItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RiskyUserHistoryCollectionRequest is request for RiskyUserHistoryItem collection
type RiskyUserHistoryCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for RiskyUserHistoryItem collection
func (r *RiskyUserHistoryCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]RiskyUserHistoryItem, error) {
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
	var values []RiskyUserHistoryItem
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
			value  []RiskyUserHistoryItem
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

// GetN performs GET request for RiskyUserHistoryItem collection, max N pages
func (r *RiskyUserHistoryCollectionRequest) GetN(ctx context.Context, n int) ([]RiskyUserHistoryItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for RiskyUserHistoryItem collection
func (r *RiskyUserHistoryCollectionRequest) Get(ctx context.Context) ([]RiskyUserHistoryItem, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for RiskyUserHistoryItem collection
func (r *RiskyUserHistoryCollectionRequest) Add(ctx context.Context, reqObj *RiskyUserHistoryItem) (resObj *RiskyUserHistoryItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RiskyServicePrincipal is navigation property rn
func (b *RiskyServicePrincipalRequestBuilder) RiskyServicePrincipal() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// RiskyServicePrincipalHistoryItem returns request builder for RiskyServicePrincipal collection rcn
func (b *RiskyServicePrincipalHistoryItemRequestBuilder) RiskyServicePrincipalHistoryItem() *RiskyServicePrincipalHistoryItemRiskyServicePrincipalHistoryItemCollectionRequestBuilder {
	bb := &RiskyServicePrincipalHistoryItemRiskyServicePrincipalHistoryItemCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/RiskyServicePrincipal"
	return bb
}

// RiskyServicePrincipalHistoryItemRiskyServicePrincipalHistoryItemCollectionRequestBuilder is request builder for RiskyServicePrincipal collection
type RiskyServicePrincipalHistoryItemRiskyServicePrincipalHistoryItemCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for RiskyServicePrincipal collection
func (b *RiskyServicePrincipalHistoryItemRiskyServicePrincipalHistoryItemCollectionRequestBuilder) Request() *RiskyServicePrincipalHistoryItemRiskyServicePrincipalHistoryItemCollectionRequest {
	return &RiskyServicePrincipalHistoryItemRiskyServicePrincipalHistoryItemCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for RiskyServicePrincipal item
func (b *RiskyServicePrincipalHistoryItemRiskyServicePrincipalHistoryItemCollectionRequestBuilder) ID(id string) *RiskyServicePrincipalRequestBuilder {
	bb := &RiskyServicePrincipalRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RiskyServicePrincipalHistoryItemRiskyServicePrincipalHistoryItemCollectionRequest is request for RiskyServicePrincipal collection
type RiskyServicePrincipalHistoryItemRiskyServicePrincipalHistoryItemCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for RiskyServicePrincipal collection
func (r *RiskyServicePrincipalHistoryItemRiskyServicePrincipalHistoryItemCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]RiskyServicePrincipal, error) {
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
	var values []RiskyServicePrincipal
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
			value  []RiskyServicePrincipal
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

// GetN performs GET request for RiskyServicePrincipal collection, max N pages
func (r *RiskyServicePrincipalHistoryItemRiskyServicePrincipalHistoryItemCollectionRequest) GetN(ctx context.Context, n int) ([]RiskyServicePrincipal, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for RiskyServicePrincipal collection
func (r *RiskyServicePrincipalHistoryItemRiskyServicePrincipalHistoryItemCollectionRequest) Get(ctx context.Context) ([]RiskyServicePrincipal, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for RiskyServicePrincipal collection
func (r *RiskyServicePrincipalHistoryItemRiskyServicePrincipalHistoryItemCollectionRequest) Add(ctx context.Context, reqObj *RiskyServicePrincipal) (resObj *RiskyServicePrincipal, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RiskyUser is navigation property rn
func (b *RiskyUserRequestBuilder) RiskyUser() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// RiskyUserHistoryItem returns request builder for RiskyUser collection rcn
func (b *RiskyUserHistoryItemRequestBuilder) RiskyUserHistoryItem() *RiskyUserHistoryItemRiskyUserHistoryItemCollectionRequestBuilder {
	bb := &RiskyUserHistoryItemRiskyUserHistoryItemCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/RiskyUser"
	return bb
}

// RiskyUserHistoryItemRiskyUserHistoryItemCollectionRequestBuilder is request builder for RiskyUser collection
type RiskyUserHistoryItemRiskyUserHistoryItemCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for RiskyUser collection
func (b *RiskyUserHistoryItemRiskyUserHistoryItemCollectionRequestBuilder) Request() *RiskyUserHistoryItemRiskyUserHistoryItemCollectionRequest {
	return &RiskyUserHistoryItemRiskyUserHistoryItemCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for RiskyUser item
func (b *RiskyUserHistoryItemRiskyUserHistoryItemCollectionRequestBuilder) ID(id string) *RiskyUserRequestBuilder {
	bb := &RiskyUserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RiskyUserHistoryItemRiskyUserHistoryItemCollectionRequest is request for RiskyUser collection
type RiskyUserHistoryItemRiskyUserHistoryItemCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for RiskyUser collection
func (r *RiskyUserHistoryItemRiskyUserHistoryItemCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]RiskyUser, error) {
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
	var values []RiskyUser
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
			value  []RiskyUser
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

// GetN performs GET request for RiskyUser collection, max N pages
func (r *RiskyUserHistoryItemRiskyUserHistoryItemCollectionRequest) GetN(ctx context.Context, n int) ([]RiskyUser, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for RiskyUser collection
func (r *RiskyUserHistoryItemRiskyUserHistoryItemCollectionRequest) Get(ctx context.Context) ([]RiskyUser, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for RiskyUser collection
func (r *RiskyUserHistoryItemRiskyUserHistoryItemCollectionRequest) Add(ctx context.Context, reqObj *RiskyUser) (resObj *RiskyUser, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

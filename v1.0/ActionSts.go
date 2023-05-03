// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// AppliesTo returns request builder for DirectoryObject collection
func (b *StsPolicyRequestBuilder) AppliesTo() *StsPolicyAppliesToCollectionRequestBuilder {
	bb := &StsPolicyAppliesToCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appliesTo"
	return bb
}

// StsPolicyAppliesToCollectionRequestBuilder is request builder for DirectoryObject collection rcn
type StsPolicyAppliesToCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *StsPolicyAppliesToCollectionRequestBuilder) Request() *StsPolicyAppliesToCollectionRequest {
	return &StsPolicyAppliesToCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *StsPolicyAppliesToCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// StsPolicyAppliesToCollectionRequest is request for DirectoryObject collection
type StsPolicyAppliesToCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *StsPolicyAppliesToCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *StsPolicyAppliesToCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *StsPolicyAppliesToCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *StsPolicyAppliesToCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// StsPolicy returns request builder for StsPolicy collection
func (b *AdministrativeUnitMembersCollectionRequestBuilder) StsPolicy() *AdministrativeUnitMembersCollectionStsPolicyCollectionRequestBuilder {
	bb := &AdministrativeUnitMembersCollectionStsPolicyCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// AdministrativeUnitMembersCollectionStsPolicyCollectionRequestBuilder is request builder for StsPolicy collection rcn
type AdministrativeUnitMembersCollectionStsPolicyCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for StsPolicy collection
func (b *AdministrativeUnitMembersCollectionStsPolicyCollectionRequestBuilder) Request() *AdministrativeUnitMembersCollectionStsPolicyCollectionRequest {
	return &AdministrativeUnitMembersCollectionStsPolicyCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for StsPolicy item
func (b *AdministrativeUnitMembersCollectionStsPolicyCollectionRequestBuilder) ID(id string) *StsPolicyRequestBuilder {
	bb := &StsPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AdministrativeUnitMembersCollectionStsPolicyCollectionRequest is request for StsPolicy collection
type AdministrativeUnitMembersCollectionStsPolicyCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for StsPolicy collection
func (r *AdministrativeUnitMembersCollectionStsPolicyCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]StsPolicy, error) {
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
	var values []StsPolicy
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
			value  []StsPolicy
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

// GetN performs GET request for StsPolicy collection, max N pages
func (r *AdministrativeUnitMembersCollectionStsPolicyCollectionRequest) GetN(ctx context.Context, n int) ([]StsPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for StsPolicy collection
func (r *AdministrativeUnitMembersCollectionStsPolicyCollectionRequest) Get(ctx context.Context) ([]StsPolicy, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for StsPolicy collection
func (r *AdministrativeUnitMembersCollectionStsPolicyCollectionRequest) Add(ctx context.Context, reqObj *StsPolicy) (resObj *StsPolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

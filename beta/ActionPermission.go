// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// PermissionGrantRequestParameter undocumented
type PermissionGrantRequestParameter struct {
	// Roles undocumented
	Roles []string `json:"roles,omitempty"`
	// Recipients undocumented
	Recipients []DriveRecipient `json:"recipients,omitempty"`
}

// PermissionRevokeGrantsRequestParameter undocumented
type PermissionRevokeGrantsRequestParameter struct {
	// Grantees undocumented
	Grantees []DriveRecipient `json:"grantees,omitempty"`
}

// Excludes returns request builder for PermissionGrantConditionSet collection
func (b *PermissionGrantPolicyRequestBuilder) Excludes() *PermissionGrantPolicyExcludesCollectionRequestBuilder {
	bb := &PermissionGrantPolicyExcludesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/excludes"
	return bb
}

// PermissionGrantPolicyExcludesCollectionRequestBuilder is request builder for PermissionGrantConditionSet collection rcn
type PermissionGrantPolicyExcludesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PermissionGrantConditionSet collection
func (b *PermissionGrantPolicyExcludesCollectionRequestBuilder) Request() *PermissionGrantPolicyExcludesCollectionRequest {
	return &PermissionGrantPolicyExcludesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PermissionGrantConditionSet item
func (b *PermissionGrantPolicyExcludesCollectionRequestBuilder) ID(id string) *PermissionGrantConditionSetRequestBuilder {
	bb := &PermissionGrantConditionSetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PermissionGrantPolicyExcludesCollectionRequest is request for PermissionGrantConditionSet collection
type PermissionGrantPolicyExcludesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PermissionGrantConditionSet collection
func (r *PermissionGrantPolicyExcludesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PermissionGrantConditionSet, error) {
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
	var values []PermissionGrantConditionSet
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
			value  []PermissionGrantConditionSet
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

// GetN performs GET request for PermissionGrantConditionSet collection, max N pages
func (r *PermissionGrantPolicyExcludesCollectionRequest) GetN(ctx context.Context, n int) ([]PermissionGrantConditionSet, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PermissionGrantConditionSet collection
func (r *PermissionGrantPolicyExcludesCollectionRequest) Get(ctx context.Context) ([]PermissionGrantConditionSet, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PermissionGrantConditionSet collection
func (r *PermissionGrantPolicyExcludesCollectionRequest) Add(ctx context.Context, reqObj *PermissionGrantConditionSet) (resObj *PermissionGrantConditionSet, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Includes returns request builder for PermissionGrantConditionSet collection
func (b *PermissionGrantPolicyRequestBuilder) Includes() *PermissionGrantPolicyIncludesCollectionRequestBuilder {
	bb := &PermissionGrantPolicyIncludesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/includes"
	return bb
}

// PermissionGrantPolicyIncludesCollectionRequestBuilder is request builder for PermissionGrantConditionSet collection rcn
type PermissionGrantPolicyIncludesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PermissionGrantConditionSet collection
func (b *PermissionGrantPolicyIncludesCollectionRequestBuilder) Request() *PermissionGrantPolicyIncludesCollectionRequest {
	return &PermissionGrantPolicyIncludesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PermissionGrantConditionSet item
func (b *PermissionGrantPolicyIncludesCollectionRequestBuilder) ID(id string) *PermissionGrantConditionSetRequestBuilder {
	bb := &PermissionGrantConditionSetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PermissionGrantPolicyIncludesCollectionRequest is request for PermissionGrantConditionSet collection
type PermissionGrantPolicyIncludesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PermissionGrantConditionSet collection
func (r *PermissionGrantPolicyIncludesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PermissionGrantConditionSet, error) {
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
	var values []PermissionGrantConditionSet
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
			value  []PermissionGrantConditionSet
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

// GetN performs GET request for PermissionGrantConditionSet collection, max N pages
func (r *PermissionGrantPolicyIncludesCollectionRequest) GetN(ctx context.Context, n int) ([]PermissionGrantConditionSet, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PermissionGrantConditionSet collection
func (r *PermissionGrantPolicyIncludesCollectionRequest) Get(ctx context.Context) ([]PermissionGrantConditionSet, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PermissionGrantConditionSet collection
func (r *PermissionGrantPolicyIncludesCollectionRequest) Add(ctx context.Context, reqObj *PermissionGrantConditionSet) (resObj *PermissionGrantConditionSet, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *PermissionRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *PermissionGrantConditionSetRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

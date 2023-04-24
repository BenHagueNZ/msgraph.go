// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// PermissionRequestBuilder is request builder for Permission
type PermissionRequestBuilder struct{ BaseRequestBuilder }

// Request returns PermissionRequest
func (b *PermissionRequestBuilder) Request() *PermissionRequest {
	return &PermissionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PermissionRequest is request for Permission
type PermissionRequest struct{ BaseRequest }

// Get performs GET request for Permission
func (r *PermissionRequest) Get(ctx context.Context) (resObj *Permission, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Permission
func (r *PermissionRequest) Update(ctx context.Context, reqObj *Permission) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Permission
func (r *PermissionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PermissionGrantConditionSetRequestBuilder is request builder for PermissionGrantConditionSet
type PermissionGrantConditionSetRequestBuilder struct{ BaseRequestBuilder }

// Request returns PermissionGrantConditionSetRequest
func (b *PermissionGrantConditionSetRequestBuilder) Request() *PermissionGrantConditionSetRequest {
	return &PermissionGrantConditionSetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PermissionGrantConditionSetRequest is request for PermissionGrantConditionSet
type PermissionGrantConditionSetRequest struct{ BaseRequest }

// Get performs GET request for PermissionGrantConditionSet
func (r *PermissionGrantConditionSetRequest) Get(ctx context.Context) (resObj *PermissionGrantConditionSet, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PermissionGrantConditionSet
func (r *PermissionGrantConditionSetRequest) Update(ctx context.Context, reqObj *PermissionGrantConditionSet) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PermissionGrantConditionSet
func (r *PermissionGrantConditionSetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PermissionGrantPolicyRequestBuilder is request builder for PermissionGrantPolicy
type PermissionGrantPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns PermissionGrantPolicyRequest
func (b *PermissionGrantPolicyRequestBuilder) Request() *PermissionGrantPolicyRequest {
	return &PermissionGrantPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PermissionGrantPolicyRequest is request for PermissionGrantPolicy
type PermissionGrantPolicyRequest struct{ BaseRequest }

// Get performs GET request for PermissionGrantPolicy
func (r *PermissionGrantPolicyRequest) Get(ctx context.Context) (resObj *PermissionGrantPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PermissionGrantPolicy
func (r *PermissionGrantPolicyRequest) Update(ctx context.Context, reqObj *PermissionGrantPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PermissionGrantPolicy
func (r *PermissionGrantPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PermissionScopeRequestBuilder is request builder for PermissionScope
type PermissionScopeRequestBuilder struct{ BaseRequestBuilder }

// Request returns PermissionScopeRequest
func (b *PermissionScopeRequestBuilder) Request() *PermissionScopeRequest {
	return &PermissionScopeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PermissionScopeRequest is request for PermissionScope
type PermissionScopeRequest struct{ BaseRequest }

// Get performs GET request for PermissionScope
func (r *PermissionScopeRequest) Get(ctx context.Context) (resObj *PermissionScope, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PermissionScope
func (r *PermissionScopeRequest) Update(ctx context.Context, reqObj *PermissionScope) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PermissionScope
func (r *PermissionScopeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type PermissionGrantRequestBuilder struct{ BaseRequestBuilder }

// Grant action undocumentedrac
func (b *PermissionRequestBuilder) Grant(reqObj *PermissionGrantRequestParameter) *PermissionGrantRequestBuilder {
	bb := &PermissionGrantRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Grant"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type PermissionGrantRequest struct{ BaseRequest }

func (b *PermissionGrantRequestBuilder) Request() *PermissionGrantRequest {
	return &PermissionGrantRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *PermissionGrantRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Permission, error) {
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
	var values []Permission
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
			value  []Permission
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
		req, _ = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

func (r *PermissionGrantRequest) PostN(ctx context.Context, n int) ([]Permission, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

func (r *PermissionGrantRequest) Post(ctx context.Context) ([]Permission, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

type PermissionRevokeGrantsRequestBuilder struct{ BaseRequestBuilder }

// RevokeGrants action undocumentedras
func (b *PermissionRequestBuilder) RevokeGrants(reqObj *PermissionRevokeGrantsRequestParameter) *PermissionRevokeGrantsRequestBuilder {
	bb := &PermissionRevokeGrantsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RevokeGrants"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type PermissionRevokeGrantsRequest struct{ BaseRequest }

func (b *PermissionRevokeGrantsRequestBuilder) Request() *PermissionRevokeGrantsRequest {
	return &PermissionRevokeGrantsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *PermissionRevokeGrantsRequest) Post(ctx context.Context) (resObj *Permission, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

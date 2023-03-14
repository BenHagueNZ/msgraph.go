// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// TeamsApps returns request builder for TeamsApp collection
func (b *AppCatalogsRequestBuilder) TeamsApps() *AppCatalogsTeamsAppsCollectionRequestBuilder {
	bb := &AppCatalogsTeamsAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/teamsApps"
	return bb
}

// AppCatalogsTeamsAppsCollectionRequestBuilder is request builder for TeamsApp collection
type AppCatalogsTeamsAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsApp collection
func (b *AppCatalogsTeamsAppsCollectionRequestBuilder) Request() *AppCatalogsTeamsAppsCollectionRequest {
	return &AppCatalogsTeamsAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsApp item
func (b *AppCatalogsTeamsAppsCollectionRequestBuilder) ID(id string) *TeamsAppRequestBuilder {
	bb := &TeamsAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AppCatalogsTeamsAppsCollectionRequest is request for TeamsApp collection
type AppCatalogsTeamsAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamsApp collection
func (r *AppCatalogsTeamsAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamsApp, error) {
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
	var values []TeamsApp
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
			value  []TeamsApp
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

// GetN performs GET request for TeamsApp collection, max N pages
func (r *AppCatalogsTeamsAppsCollectionRequest) GetN(ctx context.Context, n int) ([]TeamsApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamsApp collection
func (r *AppCatalogsTeamsAppsCollectionRequest) Get(ctx context.Context) ([]TeamsApp, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamsApp collection
func (r *AppCatalogsTeamsAppsCollectionRequest) Add(ctx context.Context, reqObj *TeamsApp) (resObj *TeamsApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AppConsentRequests returns request builder for AppConsentRequest collection
func (b *AppConsentApprovalRouteRequestBuilder) AppConsentRequests() *AppConsentApprovalRouteAppConsentRequestsCollectionRequestBuilder {
	bb := &AppConsentApprovalRouteAppConsentRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appConsentRequests"
	return bb
}

// AppConsentApprovalRouteAppConsentRequestsCollectionRequestBuilder is request builder for AppConsentRequest collection
type AppConsentApprovalRouteAppConsentRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AppConsentRequest collection
func (b *AppConsentApprovalRouteAppConsentRequestsCollectionRequestBuilder) Request() *AppConsentApprovalRouteAppConsentRequestsCollectionRequest {
	return &AppConsentApprovalRouteAppConsentRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AppConsentRequest item
func (b *AppConsentApprovalRouteAppConsentRequestsCollectionRequestBuilder) ID(id string) *AppConsentRequestRequestBuilder {
	bb := &AppConsentRequestRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AppConsentApprovalRouteAppConsentRequestsCollectionRequest is request for AppConsentRequest collection
type AppConsentApprovalRouteAppConsentRequestsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AppConsentRequest collection
func (r *AppConsentApprovalRouteAppConsentRequestsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AppConsentRequest, error) {
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
	var values []AppConsentRequest
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
			value  []AppConsentRequest
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

// GetN performs GET request for AppConsentRequest collection, max N pages
func (r *AppConsentApprovalRouteAppConsentRequestsCollectionRequest) GetN(ctx context.Context, n int) ([]AppConsentRequest, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AppConsentRequest collection
func (r *AppConsentApprovalRouteAppConsentRequestsCollectionRequest) Get(ctx context.Context) ([]AppConsentRequest, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AppConsentRequest collection
func (r *AppConsentApprovalRouteAppConsentRequestsCollectionRequest) Add(ctx context.Context, reqObj *AppConsentRequest) (resObj *AppConsentRequest, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// UserConsentRequests returns request builder for UserConsentRequest collection
func (b *AppConsentRequestObjectRequestBuilder) UserConsentRequests() *AppConsentRequestObjectUserConsentRequestsCollectionRequestBuilder {
	bb := &AppConsentRequestObjectUserConsentRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userConsentRequests"
	return bb
}

// AppConsentRequestObjectUserConsentRequestsCollectionRequestBuilder is request builder for UserConsentRequest collection
type AppConsentRequestObjectUserConsentRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UserConsentRequest collection
func (b *AppConsentRequestObjectUserConsentRequestsCollectionRequestBuilder) Request() *AppConsentRequestObjectUserConsentRequestsCollectionRequest {
	return &AppConsentRequestObjectUserConsentRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UserConsentRequest item
func (b *AppConsentRequestObjectUserConsentRequestsCollectionRequestBuilder) ID(id string) *UserConsentRequestRequestBuilder {
	bb := &UserConsentRequestRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AppConsentRequestObjectUserConsentRequestsCollectionRequest is request for UserConsentRequest collection
type AppConsentRequestObjectUserConsentRequestsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UserConsentRequest collection
func (r *AppConsentRequestObjectUserConsentRequestsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]UserConsentRequest, error) {
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
	var values []UserConsentRequest
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
			value  []UserConsentRequest
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

// GetN performs GET request for UserConsentRequest collection, max N pages
func (r *AppConsentRequestObjectUserConsentRequestsCollectionRequest) GetN(ctx context.Context, n int) ([]UserConsentRequest, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for UserConsentRequest collection
func (r *AppConsentRequestObjectUserConsentRequestsCollectionRequest) Get(ctx context.Context) ([]UserConsentRequest, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for UserConsentRequest collection
func (r *AppConsentRequestObjectUserConsentRequestsCollectionRequest) Add(ctx context.Context, reqObj *UserConsentRequest) (resObj *UserConsentRequest, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AppliesTo returns request builder for DirectoryObject collection
func (b *AppManagementPolicyRequestBuilder) AppliesTo() *AppManagementPolicyAppliesToCollectionRequestBuilder {
	bb := &AppManagementPolicyAppliesToCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appliesTo"
	return bb
}

// AppManagementPolicyAppliesToCollectionRequestBuilder is request builder for DirectoryObject collection
type AppManagementPolicyAppliesToCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *AppManagementPolicyAppliesToCollectionRequestBuilder) Request() *AppManagementPolicyAppliesToCollectionRequest {
	return &AppManagementPolicyAppliesToCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *AppManagementPolicyAppliesToCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AppManagementPolicyAppliesToCollectionRequest is request for DirectoryObject collection
type AppManagementPolicyAppliesToCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *AppManagementPolicyAppliesToCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
func (r *AppManagementPolicyAppliesToCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *AppManagementPolicyAppliesToCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *AppManagementPolicyAppliesToCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

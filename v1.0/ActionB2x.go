// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// IdentityProviders returns request builder for IdentityProvider collection
func (b *B2xIdentityUserFlowRequestBuilder) IdentityProviders() *B2xIdentityUserFlowIdentityProvidersCollectionRequestBuilder {
	bb := &B2xIdentityUserFlowIdentityProvidersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityProviders"
	return bb
}

// B2xIdentityUserFlowIdentityProvidersCollectionRequestBuilder is request builder for IdentityProvider collection rcn
type B2xIdentityUserFlowIdentityProvidersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for IdentityProvider collection
func (b *B2xIdentityUserFlowIdentityProvidersCollectionRequestBuilder) Request() *B2xIdentityUserFlowIdentityProvidersCollectionRequest {
	return &B2xIdentityUserFlowIdentityProvidersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for IdentityProvider item
func (b *B2xIdentityUserFlowIdentityProvidersCollectionRequestBuilder) ID(id string) *IdentityProviderRequestBuilder {
	bb := &IdentityProviderRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// B2xIdentityUserFlowIdentityProvidersCollectionRequest is request for IdentityProvider collection
type B2xIdentityUserFlowIdentityProvidersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for IdentityProvider collection
func (r *B2xIdentityUserFlowIdentityProvidersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]IdentityProvider, error) {
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
	var values []IdentityProvider
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
			value  []IdentityProvider
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

// GetN performs GET request for IdentityProvider collection, max N pages
func (r *B2xIdentityUserFlowIdentityProvidersCollectionRequest) GetN(ctx context.Context, n int) ([]IdentityProvider, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for IdentityProvider collection
func (r *B2xIdentityUserFlowIdentityProvidersCollectionRequest) Get(ctx context.Context) ([]IdentityProvider, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for IdentityProvider collection
func (r *B2xIdentityUserFlowIdentityProvidersCollectionRequest) Add(ctx context.Context, reqObj *IdentityProvider) (resObj *IdentityProvider, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Languages returns request builder for UserFlowLanguageConfiguration collection
func (b *B2xIdentityUserFlowRequestBuilder) Languages() *B2xIdentityUserFlowLanguagesCollectionRequestBuilder {
	bb := &B2xIdentityUserFlowLanguagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/languages"
	return bb
}

// B2xIdentityUserFlowLanguagesCollectionRequestBuilder is request builder for UserFlowLanguageConfiguration collection rcn
type B2xIdentityUserFlowLanguagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UserFlowLanguageConfiguration collection
func (b *B2xIdentityUserFlowLanguagesCollectionRequestBuilder) Request() *B2xIdentityUserFlowLanguagesCollectionRequest {
	return &B2xIdentityUserFlowLanguagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UserFlowLanguageConfiguration item
func (b *B2xIdentityUserFlowLanguagesCollectionRequestBuilder) ID(id string) *UserFlowLanguageConfigurationRequestBuilder {
	bb := &UserFlowLanguageConfigurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// B2xIdentityUserFlowLanguagesCollectionRequest is request for UserFlowLanguageConfiguration collection
type B2xIdentityUserFlowLanguagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UserFlowLanguageConfiguration collection
func (r *B2xIdentityUserFlowLanguagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]UserFlowLanguageConfiguration, error) {
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
	var values []UserFlowLanguageConfiguration
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
			value  []UserFlowLanguageConfiguration
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

// GetN performs GET request for UserFlowLanguageConfiguration collection, max N pages
func (r *B2xIdentityUserFlowLanguagesCollectionRequest) GetN(ctx context.Context, n int) ([]UserFlowLanguageConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for UserFlowLanguageConfiguration collection
func (r *B2xIdentityUserFlowLanguagesCollectionRequest) Get(ctx context.Context) ([]UserFlowLanguageConfiguration, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for UserFlowLanguageConfiguration collection
func (r *B2xIdentityUserFlowLanguagesCollectionRequest) Add(ctx context.Context, reqObj *UserFlowLanguageConfiguration) (resObj *UserFlowLanguageConfiguration, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// UserAttributeAssignments returns request builder for IdentityUserFlowAttributeAssignment collection
func (b *B2xIdentityUserFlowRequestBuilder) UserAttributeAssignments() *B2xIdentityUserFlowUserAttributeAssignmentsCollectionRequestBuilder {
	bb := &B2xIdentityUserFlowUserAttributeAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userAttributeAssignments"
	return bb
}

// B2xIdentityUserFlowUserAttributeAssignmentsCollectionRequestBuilder is request builder for IdentityUserFlowAttributeAssignment collection rcn
type B2xIdentityUserFlowUserAttributeAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for IdentityUserFlowAttributeAssignment collection
func (b *B2xIdentityUserFlowUserAttributeAssignmentsCollectionRequestBuilder) Request() *B2xIdentityUserFlowUserAttributeAssignmentsCollectionRequest {
	return &B2xIdentityUserFlowUserAttributeAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for IdentityUserFlowAttributeAssignment item
func (b *B2xIdentityUserFlowUserAttributeAssignmentsCollectionRequestBuilder) ID(id string) *IdentityUserFlowAttributeAssignmentRequestBuilder {
	bb := &IdentityUserFlowAttributeAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// B2xIdentityUserFlowUserAttributeAssignmentsCollectionRequest is request for IdentityUserFlowAttributeAssignment collection
type B2xIdentityUserFlowUserAttributeAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for IdentityUserFlowAttributeAssignment collection
func (r *B2xIdentityUserFlowUserAttributeAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]IdentityUserFlowAttributeAssignment, error) {
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
	var values []IdentityUserFlowAttributeAssignment
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
			value  []IdentityUserFlowAttributeAssignment
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

// GetN performs GET request for IdentityUserFlowAttributeAssignment collection, max N pages
func (r *B2xIdentityUserFlowUserAttributeAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]IdentityUserFlowAttributeAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for IdentityUserFlowAttributeAssignment collection
func (r *B2xIdentityUserFlowUserAttributeAssignmentsCollectionRequest) Get(ctx context.Context) ([]IdentityUserFlowAttributeAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for IdentityUserFlowAttributeAssignment collection
func (r *B2xIdentityUserFlowUserAttributeAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *IdentityUserFlowAttributeAssignment) (resObj *IdentityUserFlowAttributeAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// UserFlowIdentityProviders returns request builder for IdentityProviderBase collection
func (b *B2xIdentityUserFlowRequestBuilder) UserFlowIdentityProviders() *B2xIdentityUserFlowUserFlowIdentityProvidersCollectionRequestBuilder {
	bb := &B2xIdentityUserFlowUserFlowIdentityProvidersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userFlowIdentityProviders"
	return bb
}

// B2xIdentityUserFlowUserFlowIdentityProvidersCollectionRequestBuilder is request builder for IdentityProviderBase collection rcn
type B2xIdentityUserFlowUserFlowIdentityProvidersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for IdentityProviderBase collection
func (b *B2xIdentityUserFlowUserFlowIdentityProvidersCollectionRequestBuilder) Request() *B2xIdentityUserFlowUserFlowIdentityProvidersCollectionRequest {
	return &B2xIdentityUserFlowUserFlowIdentityProvidersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for IdentityProviderBase item
func (b *B2xIdentityUserFlowUserFlowIdentityProvidersCollectionRequestBuilder) ID(id string) *IdentityProviderBaseRequestBuilder {
	bb := &IdentityProviderBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// B2xIdentityUserFlowUserFlowIdentityProvidersCollectionRequest is request for IdentityProviderBase collection
type B2xIdentityUserFlowUserFlowIdentityProvidersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for IdentityProviderBase collection
func (r *B2xIdentityUserFlowUserFlowIdentityProvidersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]IdentityProviderBase, error) {
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
	var values []IdentityProviderBase
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
			value  []IdentityProviderBase
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

// GetN performs GET request for IdentityProviderBase collection, max N pages
func (r *B2xIdentityUserFlowUserFlowIdentityProvidersCollectionRequest) GetN(ctx context.Context, n int) ([]IdentityProviderBase, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for IdentityProviderBase collection
func (r *B2xIdentityUserFlowUserFlowIdentityProvidersCollectionRequest) Get(ctx context.Context) ([]IdentityProviderBase, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for IdentityProviderBase collection
func (r *B2xIdentityUserFlowUserFlowIdentityProvidersCollectionRequest) Add(ctx context.Context, reqObj *IdentityProviderBase) (resObj *IdentityProviderBase, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

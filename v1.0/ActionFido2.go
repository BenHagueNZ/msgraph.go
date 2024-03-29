// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// IncludeTargets returns request builder for AuthenticationMethodTarget collection
func (b *Fido2AuthenticationMethodConfigurationRequestBuilder) IncludeTargets() *Fido2AuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder {
	bb := &Fido2AuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/includeTargets"
	return bb
}

// Fido2AuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder is request builder for AuthenticationMethodTarget collection rcn
type Fido2AuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AuthenticationMethodTarget collection
func (b *Fido2AuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder) Request() *Fido2AuthenticationMethodConfigurationIncludeTargetsCollectionRequest {
	return &Fido2AuthenticationMethodConfigurationIncludeTargetsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AuthenticationMethodTarget item
func (b *Fido2AuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder) ID(id string) *AuthenticationMethodTargetRequestBuilder {
	bb := &AuthenticationMethodTargetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// Fido2AuthenticationMethodConfigurationIncludeTargetsCollectionRequest is request for AuthenticationMethodTarget collection
type Fido2AuthenticationMethodConfigurationIncludeTargetsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AuthenticationMethodTarget collection
func (r *Fido2AuthenticationMethodConfigurationIncludeTargetsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AuthenticationMethodTarget, error) {
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
	var values []AuthenticationMethodTarget
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
			value  []AuthenticationMethodTarget
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

// GetN performs GET request for AuthenticationMethodTarget collection, max N pages
func (r *Fido2AuthenticationMethodConfigurationIncludeTargetsCollectionRequest) GetN(ctx context.Context, n int) ([]AuthenticationMethodTarget, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AuthenticationMethodTarget collection
func (r *Fido2AuthenticationMethodConfigurationIncludeTargetsCollectionRequest) Get(ctx context.Context) ([]AuthenticationMethodTarget, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AuthenticationMethodTarget collection
func (r *Fido2AuthenticationMethodConfigurationIncludeTargetsCollectionRequest) Add(ctx context.Context, reqObj *AuthenticationMethodTarget) (resObj *AuthenticationMethodTarget, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Fido2AuthenticationMethod returns request builder for Fido2AuthenticationMethod collection
func (b *AuthenticationMethodsCollectionRequestBuilder) Fido2AuthenticationMethod() *AuthenticationMethodsCollectionFido2AuthenticationMethodCollectionRequestBuilder {
	bb := &AuthenticationMethodsCollectionFido2AuthenticationMethodCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// AuthenticationMethodsCollectionFido2AuthenticationMethodCollectionRequestBuilder is request builder for Fido2AuthenticationMethod collection rcn
type AuthenticationMethodsCollectionFido2AuthenticationMethodCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Fido2AuthenticationMethod collection
func (b *AuthenticationMethodsCollectionFido2AuthenticationMethodCollectionRequestBuilder) Request() *AuthenticationMethodsCollectionFido2AuthenticationMethodCollectionRequest {
	return &AuthenticationMethodsCollectionFido2AuthenticationMethodCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Fido2AuthenticationMethod item
func (b *AuthenticationMethodsCollectionFido2AuthenticationMethodCollectionRequestBuilder) ID(id string) *Fido2AuthenticationMethodRequestBuilder {
	bb := &Fido2AuthenticationMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuthenticationMethodsCollectionFido2AuthenticationMethodCollectionRequest is request for Fido2AuthenticationMethod collection
type AuthenticationMethodsCollectionFido2AuthenticationMethodCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Fido2AuthenticationMethod collection
func (r *AuthenticationMethodsCollectionFido2AuthenticationMethodCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Fido2AuthenticationMethod, error) {
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
	var values []Fido2AuthenticationMethod
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
			value  []Fido2AuthenticationMethod
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

// GetN performs GET request for Fido2AuthenticationMethod collection, max N pages
func (r *AuthenticationMethodsCollectionFido2AuthenticationMethodCollectionRequest) GetN(ctx context.Context, n int) ([]Fido2AuthenticationMethod, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Fido2AuthenticationMethod collection
func (r *AuthenticationMethodsCollectionFido2AuthenticationMethodCollectionRequest) Get(ctx context.Context) ([]Fido2AuthenticationMethod, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Fido2AuthenticationMethod collection
func (r *AuthenticationMethodsCollectionFido2AuthenticationMethodCollectionRequest) Add(ctx context.Context, reqObj *Fido2AuthenticationMethod) (resObj *Fido2AuthenticationMethod, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Fido2AuthenticationMethodConfiguration returns request builder for Fido2AuthenticationMethodConfiguration collection
func (b *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionRequestBuilder) Fido2AuthenticationMethodConfiguration() *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionFido2AuthenticationMethodConfigurationCollectionRequestBuilder {
	bb := &AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionFido2AuthenticationMethodConfigurationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionFido2AuthenticationMethodConfigurationCollectionRequestBuilder is request builder for Fido2AuthenticationMethodConfiguration collection rcn
type AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionFido2AuthenticationMethodConfigurationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Fido2AuthenticationMethodConfiguration collection
func (b *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionFido2AuthenticationMethodConfigurationCollectionRequestBuilder) Request() *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionFido2AuthenticationMethodConfigurationCollectionRequest {
	return &AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionFido2AuthenticationMethodConfigurationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Fido2AuthenticationMethodConfiguration item
func (b *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionFido2AuthenticationMethodConfigurationCollectionRequestBuilder) ID(id string) *Fido2AuthenticationMethodConfigurationRequestBuilder {
	bb := &Fido2AuthenticationMethodConfigurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionFido2AuthenticationMethodConfigurationCollectionRequest is request for Fido2AuthenticationMethodConfiguration collection
type AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionFido2AuthenticationMethodConfigurationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Fido2AuthenticationMethodConfiguration collection
func (r *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionFido2AuthenticationMethodConfigurationCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Fido2AuthenticationMethodConfiguration, error) {
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
	var values []Fido2AuthenticationMethodConfiguration
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
			value  []Fido2AuthenticationMethodConfiguration
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

// GetN performs GET request for Fido2AuthenticationMethodConfiguration collection, max N pages
func (r *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionFido2AuthenticationMethodConfigurationCollectionRequest) GetN(ctx context.Context, n int) ([]Fido2AuthenticationMethodConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Fido2AuthenticationMethodConfiguration collection
func (r *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionFido2AuthenticationMethodConfigurationCollectionRequest) Get(ctx context.Context) ([]Fido2AuthenticationMethodConfiguration, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Fido2AuthenticationMethodConfiguration collection
func (r *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionFido2AuthenticationMethodConfigurationCollectionRequest) Add(ctx context.Context, reqObj *Fido2AuthenticationMethodConfiguration) (resObj *Fido2AuthenticationMethodConfiguration, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// SocialIdentityProvider returns request builder for SocialIdentityProvider collection
func (b *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionRequestBuilder) SocialIdentityProvider() *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionSocialIdentityProviderCollectionRequestBuilder {
	bb := &B2cIdentityUserFlowUserFlowIdentityProvidersCollectionSocialIdentityProviderCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// B2cIdentityUserFlowUserFlowIdentityProvidersCollectionSocialIdentityProviderCollectionRequestBuilder is request builder for SocialIdentityProvider collection rcn
type B2cIdentityUserFlowUserFlowIdentityProvidersCollectionSocialIdentityProviderCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SocialIdentityProvider collection
func (b *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionSocialIdentityProviderCollectionRequestBuilder) Request() *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionSocialIdentityProviderCollectionRequest {
	return &B2cIdentityUserFlowUserFlowIdentityProvidersCollectionSocialIdentityProviderCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SocialIdentityProvider item
func (b *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionSocialIdentityProviderCollectionRequestBuilder) ID(id string) *SocialIdentityProviderRequestBuilder {
	bb := &SocialIdentityProviderRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// B2cIdentityUserFlowUserFlowIdentityProvidersCollectionSocialIdentityProviderCollectionRequest is request for SocialIdentityProvider collection
type B2cIdentityUserFlowUserFlowIdentityProvidersCollectionSocialIdentityProviderCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SocialIdentityProvider collection
func (r *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionSocialIdentityProviderCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SocialIdentityProvider, error) {
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
	var values []SocialIdentityProvider
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
			value  []SocialIdentityProvider
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

// GetN performs GET request for SocialIdentityProvider collection, max N pages
func (r *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionSocialIdentityProviderCollectionRequest) GetN(ctx context.Context, n int) ([]SocialIdentityProvider, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SocialIdentityProvider collection
func (r *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionSocialIdentityProviderCollectionRequest) Get(ctx context.Context) ([]SocialIdentityProvider, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SocialIdentityProvider collection
func (r *B2cIdentityUserFlowUserFlowIdentityProvidersCollectionSocialIdentityProviderCollectionRequest) Add(ctx context.Context, reqObj *SocialIdentityProvider) (resObj *SocialIdentityProvider, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

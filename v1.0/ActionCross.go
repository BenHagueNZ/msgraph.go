// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// CrossTenantAccessPolicyConfigurationDefaultResetToSystemDefaultRequestParameter undocumented
type CrossTenantAccessPolicyConfigurationDefaultResetToSystemDefaultRequestParameter struct {
}

// Default is navigation property rn
func (b *CrossTenantAccessPolicyRequestBuilder) Default() *CrossTenantAccessPolicyConfigurationDefaultRequestBuilder {
	bb := &CrossTenantAccessPolicyConfigurationDefaultRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/default"
	return bb
}

// Partners returns request builder for CrossTenantAccessPolicyConfigurationPartner collection
func (b *CrossTenantAccessPolicyRequestBuilder) Partners() *CrossTenantAccessPolicyPartnersCollectionRequestBuilder {
	bb := &CrossTenantAccessPolicyPartnersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/partners"
	return bb
}

// CrossTenantAccessPolicyPartnersCollectionRequestBuilder is request builder for CrossTenantAccessPolicyConfigurationPartner collection
type CrossTenantAccessPolicyPartnersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CrossTenantAccessPolicyConfigurationPartner collection
func (b *CrossTenantAccessPolicyPartnersCollectionRequestBuilder) Request() *CrossTenantAccessPolicyPartnersCollectionRequest {
	return &CrossTenantAccessPolicyPartnersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CrossTenantAccessPolicyConfigurationPartner item
func (b *CrossTenantAccessPolicyPartnersCollectionRequestBuilder) ID(id string) *CrossTenantAccessPolicyConfigurationPartnerRequestBuilder {
	bb := &CrossTenantAccessPolicyConfigurationPartnerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CrossTenantAccessPolicyPartnersCollectionRequest is request for CrossTenantAccessPolicyConfigurationPartner collection
type CrossTenantAccessPolicyPartnersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CrossTenantAccessPolicyConfigurationPartner collection
func (r *CrossTenantAccessPolicyPartnersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]CrossTenantAccessPolicyConfigurationPartner, error) {
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
	var values []CrossTenantAccessPolicyConfigurationPartner
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
			value  []CrossTenantAccessPolicyConfigurationPartner
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

// GetN performs GET request for CrossTenantAccessPolicyConfigurationPartner collection, max N pages
func (r *CrossTenantAccessPolicyPartnersCollectionRequest) GetN(ctx context.Context, n int) ([]CrossTenantAccessPolicyConfigurationPartner, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for CrossTenantAccessPolicyConfigurationPartner collection
func (r *CrossTenantAccessPolicyPartnersCollectionRequest) Get(ctx context.Context) ([]CrossTenantAccessPolicyConfigurationPartner, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for CrossTenantAccessPolicyConfigurationPartner collection
func (r *CrossTenantAccessPolicyPartnersCollectionRequest) Add(ctx context.Context, reqObj *CrossTenantAccessPolicyConfigurationPartner) (resObj *CrossTenantAccessPolicyConfigurationPartner, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *CrossTenantAccessPolicyConfigurationDefaultRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// PasswordAuthenticationMethod returns request builder for AuthenticationMethod collection rcn
func (b *PasswordAuthenticationMethodRequestBuilder) PasswordAuthenticationMethod() *PasswordAuthenticationMethodPasswordAuthenticationMethodCollectionRequestBuilder {
	bb := &PasswordAuthenticationMethodPasswordAuthenticationMethodCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/AuthenticationMethod"
	return bb
}

// PasswordAuthenticationMethodPasswordAuthenticationMethodCollectionRequestBuilder is request builder for AuthenticationMethod collection
type PasswordAuthenticationMethodPasswordAuthenticationMethodCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AuthenticationMethod collection
func (b *PasswordAuthenticationMethodPasswordAuthenticationMethodCollectionRequestBuilder) Request() *PasswordAuthenticationMethodPasswordAuthenticationMethodCollectionRequest {
	return &PasswordAuthenticationMethodPasswordAuthenticationMethodCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AuthenticationMethod item
func (b *PasswordAuthenticationMethodPasswordAuthenticationMethodCollectionRequestBuilder) ID(id string) *AuthenticationMethodRequestBuilder {
	bb := &AuthenticationMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PasswordAuthenticationMethodPasswordAuthenticationMethodCollectionRequest is request for AuthenticationMethod collection
type PasswordAuthenticationMethodPasswordAuthenticationMethodCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AuthenticationMethod collection
func (r *PasswordAuthenticationMethodPasswordAuthenticationMethodCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AuthenticationMethod, error) {
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
	var values []AuthenticationMethod
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
			value  []AuthenticationMethod
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

// GetN performs GET request for AuthenticationMethod collection, max N pages
func (r *PasswordAuthenticationMethodPasswordAuthenticationMethodCollectionRequest) GetN(ctx context.Context, n int) ([]AuthenticationMethod, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AuthenticationMethod collection
func (r *PasswordAuthenticationMethodPasswordAuthenticationMethodCollectionRequest) Get(ctx context.Context) ([]AuthenticationMethod, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AuthenticationMethod collection
func (r *PasswordAuthenticationMethodPasswordAuthenticationMethodCollectionRequest) Add(ctx context.Context, reqObj *AuthenticationMethod) (resObj *AuthenticationMethod, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

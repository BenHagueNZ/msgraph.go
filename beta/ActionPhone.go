// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// PhoneAuthenticationMethod returns request builder for PhoneAuthenticationMethod collection
func (b *AuthenticationMethodsCollectionRequestBuilder) PhoneAuthenticationMethod() *AuthenticationMethodsCollectionPhoneAuthenticationMethodCollectionRequestBuilder {
	bb := &AuthenticationMethodsCollectionPhoneAuthenticationMethodCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// AuthenticationMethodsCollectionPhoneAuthenticationMethodCollectionRequestBuilder is request builder for PhoneAuthenticationMethod collection rcn
type AuthenticationMethodsCollectionPhoneAuthenticationMethodCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PhoneAuthenticationMethod collection
func (b *AuthenticationMethodsCollectionPhoneAuthenticationMethodCollectionRequestBuilder) Request() *AuthenticationMethodsCollectionPhoneAuthenticationMethodCollectionRequest {
	return &AuthenticationMethodsCollectionPhoneAuthenticationMethodCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PhoneAuthenticationMethod item
func (b *AuthenticationMethodsCollectionPhoneAuthenticationMethodCollectionRequestBuilder) ID(id string) *PhoneAuthenticationMethodRequestBuilder {
	bb := &PhoneAuthenticationMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuthenticationMethodsCollectionPhoneAuthenticationMethodCollectionRequest is request for PhoneAuthenticationMethod collection
type AuthenticationMethodsCollectionPhoneAuthenticationMethodCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PhoneAuthenticationMethod collection
func (r *AuthenticationMethodsCollectionPhoneAuthenticationMethodCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PhoneAuthenticationMethod, error) {
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
	var values []PhoneAuthenticationMethod
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
			value  []PhoneAuthenticationMethod
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

// GetN performs GET request for PhoneAuthenticationMethod collection, max N pages
func (r *AuthenticationMethodsCollectionPhoneAuthenticationMethodCollectionRequest) GetN(ctx context.Context, n int) ([]PhoneAuthenticationMethod, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PhoneAuthenticationMethod collection
func (r *AuthenticationMethodsCollectionPhoneAuthenticationMethodCollectionRequest) Get(ctx context.Context) ([]PhoneAuthenticationMethod, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PhoneAuthenticationMethod collection
func (r *AuthenticationMethodsCollectionPhoneAuthenticationMethodCollectionRequest) Add(ctx context.Context, reqObj *PhoneAuthenticationMethod) (resObj *PhoneAuthenticationMethod, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

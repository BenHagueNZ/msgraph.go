// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// BrowserSiteListPublishRequestParameter undocumented
type BrowserSiteListPublishRequestParameter struct {
	// Revision undocumented
	Revision *string `json:"revision,omitempty"`
	// Sites undocumented
	Sites []BrowserSite `json:"sites,omitempty"`
	// SharedCookies undocumented
	SharedCookies []BrowserSharedCookie `json:"sharedCookies,omitempty"`
}

// SharedCookies returns request builder for BrowserSharedCookie collection
func (b *BrowserSiteListRequestBuilder) SharedCookies() *BrowserSiteListSharedCookiesCollectionRequestBuilder {
	bb := &BrowserSiteListSharedCookiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sharedCookies"
	return bb
}

// BrowserSiteListSharedCookiesCollectionRequestBuilder is request builder for BrowserSharedCookie collection rcn
type BrowserSiteListSharedCookiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BrowserSharedCookie collection
func (b *BrowserSiteListSharedCookiesCollectionRequestBuilder) Request() *BrowserSiteListSharedCookiesCollectionRequest {
	return &BrowserSiteListSharedCookiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BrowserSharedCookie item
func (b *BrowserSiteListSharedCookiesCollectionRequestBuilder) ID(id string) *BrowserSharedCookieRequestBuilder {
	bb := &BrowserSharedCookieRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// BrowserSiteListSharedCookiesCollectionRequest is request for BrowserSharedCookie collection
type BrowserSiteListSharedCookiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BrowserSharedCookie collection
func (r *BrowserSiteListSharedCookiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]BrowserSharedCookie, error) {
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
	var values []BrowserSharedCookie
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
			value  []BrowserSharedCookie
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

// GetN performs GET request for BrowserSharedCookie collection, max N pages
func (r *BrowserSiteListSharedCookiesCollectionRequest) GetN(ctx context.Context, n int) ([]BrowserSharedCookie, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for BrowserSharedCookie collection
func (r *BrowserSiteListSharedCookiesCollectionRequest) Get(ctx context.Context) ([]BrowserSharedCookie, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for BrowserSharedCookie collection
func (r *BrowserSiteListSharedCookiesCollectionRequest) Add(ctx context.Context, reqObj *BrowserSharedCookie) (resObj *BrowserSharedCookie, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Sites returns request builder for BrowserSite collection
func (b *BrowserSiteListRequestBuilder) Sites() *BrowserSiteListSitesCollectionRequestBuilder {
	bb := &BrowserSiteListSitesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sites"
	return bb
}

// BrowserSiteListSitesCollectionRequestBuilder is request builder for BrowserSite collection rcn
type BrowserSiteListSitesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BrowserSite collection
func (b *BrowserSiteListSitesCollectionRequestBuilder) Request() *BrowserSiteListSitesCollectionRequest {
	return &BrowserSiteListSitesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BrowserSite item
func (b *BrowserSiteListSitesCollectionRequestBuilder) ID(id string) *BrowserSiteRequestBuilder {
	bb := &BrowserSiteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// BrowserSiteListSitesCollectionRequest is request for BrowserSite collection
type BrowserSiteListSitesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BrowserSite collection
func (r *BrowserSiteListSitesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]BrowserSite, error) {
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
	var values []BrowserSite
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
			value  []BrowserSite
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

// GetN performs GET request for BrowserSite collection, max N pages
func (r *BrowserSiteListSitesCollectionRequest) GetN(ctx context.Context, n int) ([]BrowserSite, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for BrowserSite collection
func (r *BrowserSiteListSitesCollectionRequest) Get(ctx context.Context) ([]BrowserSite, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for BrowserSite collection
func (r *BrowserSiteListSitesCollectionRequest) Add(ctx context.Context, reqObj *BrowserSite) (resObj *BrowserSite, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *BrowserSharedCookieRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *BrowserSiteRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *BrowserSiteListRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
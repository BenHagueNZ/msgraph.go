// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// SiteLists returns request builder for BrowserSiteList collection
func (b *InternetExplorerModeRequestBuilder) SiteLists() *InternetExplorerModeSiteListsCollectionRequestBuilder {
	bb := &InternetExplorerModeSiteListsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/siteLists"
	return bb
}

// InternetExplorerModeSiteListsCollectionRequestBuilder is request builder for BrowserSiteList collection rcn
type InternetExplorerModeSiteListsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BrowserSiteList collection
func (b *InternetExplorerModeSiteListsCollectionRequestBuilder) Request() *InternetExplorerModeSiteListsCollectionRequest {
	return &InternetExplorerModeSiteListsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BrowserSiteList item
func (b *InternetExplorerModeSiteListsCollectionRequestBuilder) ID(id string) *BrowserSiteListRequestBuilder {
	bb := &BrowserSiteListRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// InternetExplorerModeSiteListsCollectionRequest is request for BrowserSiteList collection
type InternetExplorerModeSiteListsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BrowserSiteList collection
func (r *InternetExplorerModeSiteListsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]BrowserSiteList, error) {
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
	var values []BrowserSiteList
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
			value  []BrowserSiteList
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

// GetN performs GET request for BrowserSiteList collection, max N pages
func (r *InternetExplorerModeSiteListsCollectionRequest) GetN(ctx context.Context, n int) ([]BrowserSiteList, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for BrowserSiteList collection
func (r *InternetExplorerModeSiteListsCollectionRequest) Get(ctx context.Context) ([]BrowserSiteList, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for BrowserSiteList collection
func (r *InternetExplorerModeSiteListsCollectionRequest) Add(ctx context.Context, reqObj *BrowserSiteList) (resObj *BrowserSiteList, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *InternetExplorerModeRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
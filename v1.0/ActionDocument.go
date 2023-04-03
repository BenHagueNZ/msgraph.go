// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// DocumentSetVersionRestoreRequestParameter undocumented
type DocumentSetVersionRestoreRequestParameter struct {
}

// SharedColumns returns request builder for ColumnDefinition collection rcn
func (b *DocumentSetRequestBuilder) SharedColumns() *DocumentSetSharedColumnsCollectionRequestBuilder {
	bb := &DocumentSetSharedColumnsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sharedColumns"
	return bb
}

// DocumentSetSharedColumnsCollectionRequestBuilder is request builder for ColumnDefinition collection
type DocumentSetSharedColumnsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ColumnDefinition collection
func (b *DocumentSetSharedColumnsCollectionRequestBuilder) Request() *DocumentSetSharedColumnsCollectionRequest {
	return &DocumentSetSharedColumnsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ColumnDefinition item
func (b *DocumentSetSharedColumnsCollectionRequestBuilder) ID(id string) *ColumnDefinitionRequestBuilder {
	bb := &ColumnDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DocumentSetSharedColumnsCollectionRequest is request for ColumnDefinition collection
type DocumentSetSharedColumnsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ColumnDefinition collection
func (r *DocumentSetSharedColumnsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ColumnDefinition, error) {
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
	var values []ColumnDefinition
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
			value  []ColumnDefinition
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

// GetN performs GET request for ColumnDefinition collection, max N pages
func (r *DocumentSetSharedColumnsCollectionRequest) GetN(ctx context.Context, n int) ([]ColumnDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ColumnDefinition collection
func (r *DocumentSetSharedColumnsCollectionRequest) Get(ctx context.Context) ([]ColumnDefinition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ColumnDefinition collection
func (r *DocumentSetSharedColumnsCollectionRequest) Add(ctx context.Context, reqObj *ColumnDefinition) (resObj *ColumnDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// WelcomePageColumns returns request builder for ColumnDefinition collection rcn
func (b *DocumentSetRequestBuilder) WelcomePageColumns() *DocumentSetWelcomePageColumnsCollectionRequestBuilder {
	bb := &DocumentSetWelcomePageColumnsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/welcomePageColumns"
	return bb
}

// DocumentSetWelcomePageColumnsCollectionRequestBuilder is request builder for ColumnDefinition collection
type DocumentSetWelcomePageColumnsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ColumnDefinition collection
func (b *DocumentSetWelcomePageColumnsCollectionRequestBuilder) Request() *DocumentSetWelcomePageColumnsCollectionRequest {
	return &DocumentSetWelcomePageColumnsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ColumnDefinition item
func (b *DocumentSetWelcomePageColumnsCollectionRequestBuilder) ID(id string) *ColumnDefinitionRequestBuilder {
	bb := &ColumnDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DocumentSetWelcomePageColumnsCollectionRequest is request for ColumnDefinition collection
type DocumentSetWelcomePageColumnsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ColumnDefinition collection
func (r *DocumentSetWelcomePageColumnsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ColumnDefinition, error) {
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
	var values []ColumnDefinition
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
			value  []ColumnDefinition
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

// GetN performs GET request for ColumnDefinition collection, max N pages
func (r *DocumentSetWelcomePageColumnsCollectionRequest) GetN(ctx context.Context, n int) ([]ColumnDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ColumnDefinition collection
func (r *DocumentSetWelcomePageColumnsCollectionRequest) Get(ctx context.Context) ([]ColumnDefinition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ColumnDefinition collection
func (r *DocumentSetWelcomePageColumnsCollectionRequest) Add(ctx context.Context, reqObj *ColumnDefinition) (resObj *ColumnDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DocumentSetVersion returns request builder for ListItemVersion collection rcn
func (b *DocumentSetVersionRequestBuilder) DocumentSetVersion() *DocumentSetVersionDocumentSetVersionCollectionRequestBuilder {
	bb := &DocumentSetVersionDocumentSetVersionCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/ListItemVersion"
	return bb
}

// DocumentSetVersionDocumentSetVersionCollectionRequestBuilder is request builder for ListItemVersion collection
type DocumentSetVersionDocumentSetVersionCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ListItemVersion collection
func (b *DocumentSetVersionDocumentSetVersionCollectionRequestBuilder) Request() *DocumentSetVersionDocumentSetVersionCollectionRequest {
	return &DocumentSetVersionDocumentSetVersionCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ListItemVersion item
func (b *DocumentSetVersionDocumentSetVersionCollectionRequestBuilder) ID(id string) *ListItemVersionRequestBuilder {
	bb := &ListItemVersionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DocumentSetVersionDocumentSetVersionCollectionRequest is request for ListItemVersion collection
type DocumentSetVersionDocumentSetVersionCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ListItemVersion collection
func (r *DocumentSetVersionDocumentSetVersionCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ListItemVersion, error) {
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
	var values []ListItemVersion
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
			value  []ListItemVersion
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

// GetN performs GET request for ListItemVersion collection, max N pages
func (r *DocumentSetVersionDocumentSetVersionCollectionRequest) GetN(ctx context.Context, n int) ([]ListItemVersion, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ListItemVersion collection
func (r *DocumentSetVersionDocumentSetVersionCollectionRequest) Get(ctx context.Context) ([]ListItemVersion, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ListItemVersion collection
func (r *DocumentSetVersionDocumentSetVersionCollectionRequest) Add(ctx context.Context, reqObj *ListItemVersion) (resObj *ListItemVersion, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

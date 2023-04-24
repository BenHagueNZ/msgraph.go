// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// ListItemVersionRestoreVersionRequestParameter undocumented
type ListItemVersionRestoreVersionRequestParameter struct {
}

// Columns returns request builder for ColumnDefinition collection
func (b *ListRequestBuilder) Columns() *ListColumnsCollectionRequestBuilder {
	bb := &ListColumnsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/columns"
	return bb
}

// ListColumnsCollectionRequestBuilder is request builder for ColumnDefinition collection rcn
type ListColumnsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ColumnDefinition collection
func (b *ListColumnsCollectionRequestBuilder) Request() *ListColumnsCollectionRequest {
	return &ListColumnsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ColumnDefinition item
func (b *ListColumnsCollectionRequestBuilder) ID(id string) *ColumnDefinitionRequestBuilder {
	bb := &ColumnDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListColumnsCollectionRequest is request for ColumnDefinition collection
type ListColumnsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ColumnDefinition collection
func (r *ListColumnsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ColumnDefinition, error) {
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
func (r *ListColumnsCollectionRequest) GetN(ctx context.Context, n int) ([]ColumnDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ColumnDefinition collection
func (r *ListColumnsCollectionRequest) Get(ctx context.Context) ([]ColumnDefinition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ColumnDefinition collection
func (r *ListColumnsCollectionRequest) Add(ctx context.Context, reqObj *ColumnDefinition) (resObj *ColumnDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ContentTypes returns request builder for ContentType collection
func (b *ListRequestBuilder) ContentTypes() *ListContentTypesCollectionRequestBuilder {
	bb := &ListContentTypesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/contentTypes"
	return bb
}

// ListContentTypesCollectionRequestBuilder is request builder for ContentType collection rcn
type ListContentTypesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ContentType collection
func (b *ListContentTypesCollectionRequestBuilder) Request() *ListContentTypesCollectionRequest {
	return &ListContentTypesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ContentType item
func (b *ListContentTypesCollectionRequestBuilder) ID(id string) *ContentTypeRequestBuilder {
	bb := &ContentTypeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListContentTypesCollectionRequest is request for ContentType collection
type ListContentTypesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ContentType collection
func (r *ListContentTypesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ContentType, error) {
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
	var values []ContentType
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
			value  []ContentType
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

// GetN performs GET request for ContentType collection, max N pages
func (r *ListContentTypesCollectionRequest) GetN(ctx context.Context, n int) ([]ContentType, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ContentType collection
func (r *ListContentTypesCollectionRequest) Get(ctx context.Context) ([]ContentType, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ContentType collection
func (r *ListContentTypesCollectionRequest) Add(ctx context.Context, reqObj *ContentType) (resObj *ContentType, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Drive is navigation property rn
func (b *ListRequestBuilder) Drive() *DriveRequestBuilder {
	bb := &DriveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/drive"
	return bb
}

// Items returns request builder for ListItem collection
func (b *ListRequestBuilder) Items() *ListItemsCollectionRequestBuilder {
	bb := &ListItemsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/items"
	return bb
}

// ListItemsCollectionRequestBuilder is request builder for ListItem collection rcn
type ListItemsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ListItem collection
func (b *ListItemsCollectionRequestBuilder) Request() *ListItemsCollectionRequest {
	return &ListItemsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ListItem item
func (b *ListItemsCollectionRequestBuilder) ID(id string) *ListItemRequestBuilder {
	bb := &ListItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListItemsCollectionRequest is request for ListItem collection
type ListItemsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ListItem collection
func (r *ListItemsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ListItem, error) {
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
	var values []ListItem
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
			value  []ListItem
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

// GetN performs GET request for ListItem collection, max N pages
func (r *ListItemsCollectionRequest) GetN(ctx context.Context, n int) ([]ListItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ListItem collection
func (r *ListItemsCollectionRequest) Get(ctx context.Context) ([]ListItem, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ListItem collection
func (r *ListItemsCollectionRequest) Add(ctx context.Context, reqObj *ListItem) (resObj *ListItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Operations returns request builder for RichLongRunningOperation collection
func (b *ListRequestBuilder) Operations() *ListOperationsCollectionRequestBuilder {
	bb := &ListOperationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/operations"
	return bb
}

// ListOperationsCollectionRequestBuilder is request builder for RichLongRunningOperation collection rcn
type ListOperationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for RichLongRunningOperation collection
func (b *ListOperationsCollectionRequestBuilder) Request() *ListOperationsCollectionRequest {
	return &ListOperationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for RichLongRunningOperation item
func (b *ListOperationsCollectionRequestBuilder) ID(id string) *RichLongRunningOperationRequestBuilder {
	bb := &RichLongRunningOperationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListOperationsCollectionRequest is request for RichLongRunningOperation collection
type ListOperationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for RichLongRunningOperation collection
func (r *ListOperationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]RichLongRunningOperation, error) {
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
	var values []RichLongRunningOperation
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
			value  []RichLongRunningOperation
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

// GetN performs GET request for RichLongRunningOperation collection, max N pages
func (r *ListOperationsCollectionRequest) GetN(ctx context.Context, n int) ([]RichLongRunningOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for RichLongRunningOperation collection
func (r *ListOperationsCollectionRequest) Get(ctx context.Context) ([]RichLongRunningOperation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for RichLongRunningOperation collection
func (r *ListOperationsCollectionRequest) Add(ctx context.Context, reqObj *RichLongRunningOperation) (resObj *RichLongRunningOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Subscriptions returns request builder for Subscription collection
func (b *ListRequestBuilder) Subscriptions() *ListSubscriptionsCollectionRequestBuilder {
	bb := &ListSubscriptionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/subscriptions"
	return bb
}

// ListSubscriptionsCollectionRequestBuilder is request builder for Subscription collection rcn
type ListSubscriptionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Subscription collection
func (b *ListSubscriptionsCollectionRequestBuilder) Request() *ListSubscriptionsCollectionRequest {
	return &ListSubscriptionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Subscription item
func (b *ListSubscriptionsCollectionRequestBuilder) ID(id string) *SubscriptionRequestBuilder {
	bb := &SubscriptionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListSubscriptionsCollectionRequest is request for Subscription collection
type ListSubscriptionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Subscription collection
func (r *ListSubscriptionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Subscription, error) {
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
	var values []Subscription
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
			value  []Subscription
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

// GetN performs GET request for Subscription collection, max N pages
func (r *ListSubscriptionsCollectionRequest) GetN(ctx context.Context, n int) ([]Subscription, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Subscription collection
func (r *ListSubscriptionsCollectionRequest) Get(ctx context.Context) ([]Subscription, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Subscription collection
func (r *ListSubscriptionsCollectionRequest) Add(ctx context.Context, reqObj *Subscription) (resObj *Subscription, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Analytics is navigation property rn
func (b *ListItemRequestBuilder) Analytics() *ItemAnalyticsRequestBuilder {
	bb := &ItemAnalyticsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/analytics"
	return bb
}

// DocumentSetVersions returns request builder for DocumentSetVersion collection
func (b *ListItemRequestBuilder) DocumentSetVersions() *ListItemDocumentSetVersionsCollectionRequestBuilder {
	bb := &ListItemDocumentSetVersionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/documentSetVersions"
	return bb
}

// ListItemDocumentSetVersionsCollectionRequestBuilder is request builder for DocumentSetVersion collection rcn
type ListItemDocumentSetVersionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DocumentSetVersion collection
func (b *ListItemDocumentSetVersionsCollectionRequestBuilder) Request() *ListItemDocumentSetVersionsCollectionRequest {
	return &ListItemDocumentSetVersionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DocumentSetVersion item
func (b *ListItemDocumentSetVersionsCollectionRequestBuilder) ID(id string) *DocumentSetVersionRequestBuilder {
	bb := &DocumentSetVersionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListItemDocumentSetVersionsCollectionRequest is request for DocumentSetVersion collection
type ListItemDocumentSetVersionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DocumentSetVersion collection
func (r *ListItemDocumentSetVersionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DocumentSetVersion, error) {
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
	var values []DocumentSetVersion
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
			value  []DocumentSetVersion
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

// GetN performs GET request for DocumentSetVersion collection, max N pages
func (r *ListItemDocumentSetVersionsCollectionRequest) GetN(ctx context.Context, n int) ([]DocumentSetVersion, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DocumentSetVersion collection
func (r *ListItemDocumentSetVersionsCollectionRequest) Get(ctx context.Context) ([]DocumentSetVersion, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DocumentSetVersion collection
func (r *ListItemDocumentSetVersionsCollectionRequest) Add(ctx context.Context, reqObj *DocumentSetVersion) (resObj *DocumentSetVersion, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DriveItem is navigation property rn
func (b *ListItemRequestBuilder) DriveItem() *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/driveItem"
	return bb
}

// Fields is navigation property rn
func (b *ListItemRequestBuilder) Fields() *FieldValueSetRequestBuilder {
	bb := &FieldValueSetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/fields"
	return bb
}

// Versions returns request builder for ListItemVersion collection
func (b *ListItemRequestBuilder) Versions() *ListItemVersionsCollectionRequestBuilder {
	bb := &ListItemVersionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/versions"
	return bb
}

// ListItemVersionsCollectionRequestBuilder is request builder for ListItemVersion collection rcn
type ListItemVersionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ListItemVersion collection
func (b *ListItemVersionsCollectionRequestBuilder) Request() *ListItemVersionsCollectionRequest {
	return &ListItemVersionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ListItemVersion item
func (b *ListItemVersionsCollectionRequestBuilder) ID(id string) *ListItemVersionRequestBuilder {
	bb := &ListItemVersionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ListItemVersionsCollectionRequest is request for ListItemVersion collection
type ListItemVersionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ListItemVersion collection
func (r *ListItemVersionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ListItemVersion, error) {
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
func (r *ListItemVersionsCollectionRequest) GetN(ctx context.Context, n int) ([]ListItemVersion, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ListItemVersion collection
func (r *ListItemVersionsCollectionRequest) Get(ctx context.Context) ([]ListItemVersion, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ListItemVersion collection
func (r *ListItemVersionsCollectionRequest) Add(ctx context.Context, reqObj *ListItemVersion) (resObj *ListItemVersion, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Fields is navigation property rn
func (b *ListItemVersionRequestBuilder) Fields() *FieldValueSetRequestBuilder {
	bb := &FieldValueSetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/fields"
	return bb
}

// List returns request builder for List collection
func (b *SiteItemsCollectionRequestBuilder) List() *SiteItemsCollectionListCollectionRequestBuilder {
	bb := &SiteItemsCollectionListCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// SiteItemsCollectionListCollectionRequestBuilder is request builder for List collection rcn
type SiteItemsCollectionListCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for List collection
func (b *SiteItemsCollectionListCollectionRequestBuilder) Request() *SiteItemsCollectionListCollectionRequest {
	return &SiteItemsCollectionListCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for List item
func (b *SiteItemsCollectionListCollectionRequestBuilder) ID(id string) *ListRequestBuilder {
	bb := &ListRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteItemsCollectionListCollectionRequest is request for List collection
type SiteItemsCollectionListCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for List collection
func (r *SiteItemsCollectionListCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]List, error) {
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
	var values []List
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
			value  []List
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

// GetN performs GET request for List collection, max N pages
func (r *SiteItemsCollectionListCollectionRequest) GetN(ctx context.Context, n int) ([]List, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for List collection
func (r *SiteItemsCollectionListCollectionRequest) Get(ctx context.Context) ([]List, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for List collection
func (r *SiteItemsCollectionListCollectionRequest) Add(ctx context.Context, reqObj *List) (resObj *List, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ListItem returns request builder for ListItem collection
func (b *SiteItemsCollectionRequestBuilder) ListItem() *SiteItemsCollectionListItemCollectionRequestBuilder {
	bb := &SiteItemsCollectionListItemCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// SiteItemsCollectionListItemCollectionRequestBuilder is request builder for ListItem collection rcn
type SiteItemsCollectionListItemCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ListItem collection
func (b *SiteItemsCollectionListItemCollectionRequestBuilder) Request() *SiteItemsCollectionListItemCollectionRequest {
	return &SiteItemsCollectionListItemCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ListItem item
func (b *SiteItemsCollectionListItemCollectionRequestBuilder) ID(id string) *ListItemRequestBuilder {
	bb := &ListItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteItemsCollectionListItemCollectionRequest is request for ListItem collection
type SiteItemsCollectionListItemCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ListItem collection
func (r *SiteItemsCollectionListItemCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ListItem, error) {
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
	var values []ListItem
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
			value  []ListItem
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

// GetN performs GET request for ListItem collection, max N pages
func (r *SiteItemsCollectionListItemCollectionRequest) GetN(ctx context.Context, n int) ([]ListItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ListItem collection
func (r *SiteItemsCollectionListItemCollectionRequest) Get(ctx context.Context) ([]ListItem, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ListItem collection
func (r *SiteItemsCollectionListItemCollectionRequest) Add(ctx context.Context, reqObj *ListItem) (resObj *ListItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

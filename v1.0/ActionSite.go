// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// SiteCollectionAddRequestParameter undocumented
type SiteCollectionAddRequestParameter struct {
	// Value undocumented
	Value []Site `json:"value,omitempty"`
}

// SiteCollectionRemoveRequestParameter undocumented
type SiteCollectionRemoveRequestParameter struct {
	// Value undocumented
	Value []Site `json:"value,omitempty"`
}

// Analytics is navigation property rn
func (b *SiteRequestBuilder) Analytics() *ItemAnalyticsRequestBuilder {
	bb := &ItemAnalyticsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/analytics"
	return bb
}

// Columns returns request builder for ColumnDefinition collection
func (b *SiteRequestBuilder) Columns() *SiteColumnsCollectionRequestBuilder {
	bb := &SiteColumnsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/columns"
	return bb
}

// SiteColumnsCollectionRequestBuilder is request builder for ColumnDefinition collection
type SiteColumnsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ColumnDefinition collection
func (b *SiteColumnsCollectionRequestBuilder) Request() *SiteColumnsCollectionRequest {
	return &SiteColumnsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ColumnDefinition item
func (b *SiteColumnsCollectionRequestBuilder) ID(id string) *ColumnDefinitionRequestBuilder {
	bb := &ColumnDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteColumnsCollectionRequest is request for ColumnDefinition collection
type SiteColumnsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ColumnDefinition collection
func (r *SiteColumnsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ColumnDefinition, error) {
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
func (r *SiteColumnsCollectionRequest) GetN(ctx context.Context, n int) ([]ColumnDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ColumnDefinition collection
func (r *SiteColumnsCollectionRequest) Get(ctx context.Context) ([]ColumnDefinition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ColumnDefinition collection
func (r *SiteColumnsCollectionRequest) Add(ctx context.Context, reqObj *ColumnDefinition) (resObj *ColumnDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ContentTypes returns request builder for ContentType collection
func (b *SiteRequestBuilder) ContentTypes() *SiteContentTypesCollectionRequestBuilder {
	bb := &SiteContentTypesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/contentTypes"
	return bb
}

// SiteContentTypesCollectionRequestBuilder is request builder for ContentType collection
type SiteContentTypesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ContentType collection
func (b *SiteContentTypesCollectionRequestBuilder) Request() *SiteContentTypesCollectionRequest {
	return &SiteContentTypesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ContentType item
func (b *SiteContentTypesCollectionRequestBuilder) ID(id string) *ContentTypeRequestBuilder {
	bb := &ContentTypeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteContentTypesCollectionRequest is request for ContentType collection
type SiteContentTypesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ContentType collection
func (r *SiteContentTypesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ContentType, error) {
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
func (r *SiteContentTypesCollectionRequest) GetN(ctx context.Context, n int) ([]ContentType, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ContentType collection
func (r *SiteContentTypesCollectionRequest) Get(ctx context.Context) ([]ContentType, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ContentType collection
func (r *SiteContentTypesCollectionRequest) Add(ctx context.Context, reqObj *ContentType) (resObj *ContentType, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Drive is navigation property rn
func (b *SiteRequestBuilder) Drive() *DriveRequestBuilder {
	bb := &DriveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/drive"
	return bb
}

// Drives returns request builder for Drive collection
func (b *SiteRequestBuilder) Drives() *SiteDrivesCollectionRequestBuilder {
	bb := &SiteDrivesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/drives"
	return bb
}

// SiteDrivesCollectionRequestBuilder is request builder for Drive collection
type SiteDrivesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Drive collection
func (b *SiteDrivesCollectionRequestBuilder) Request() *SiteDrivesCollectionRequest {
	return &SiteDrivesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Drive item
func (b *SiteDrivesCollectionRequestBuilder) ID(id string) *DriveRequestBuilder {
	bb := &DriveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteDrivesCollectionRequest is request for Drive collection
type SiteDrivesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Drive collection
func (r *SiteDrivesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Drive, error) {
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
	var values []Drive
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
			value  []Drive
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

// GetN performs GET request for Drive collection, max N pages
func (r *SiteDrivesCollectionRequest) GetN(ctx context.Context, n int) ([]Drive, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Drive collection
func (r *SiteDrivesCollectionRequest) Get(ctx context.Context) ([]Drive, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Drive collection
func (r *SiteDrivesCollectionRequest) Add(ctx context.Context, reqObj *Drive) (resObj *Drive, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ExternalColumns returns request builder for ColumnDefinition collection
func (b *SiteRequestBuilder) ExternalColumns() *SiteExternalColumnsCollectionRequestBuilder {
	bb := &SiteExternalColumnsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/externalColumns"
	return bb
}

// SiteExternalColumnsCollectionRequestBuilder is request builder for ColumnDefinition collection
type SiteExternalColumnsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ColumnDefinition collection
func (b *SiteExternalColumnsCollectionRequestBuilder) Request() *SiteExternalColumnsCollectionRequest {
	return &SiteExternalColumnsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ColumnDefinition item
func (b *SiteExternalColumnsCollectionRequestBuilder) ID(id string) *ColumnDefinitionRequestBuilder {
	bb := &ColumnDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteExternalColumnsCollectionRequest is request for ColumnDefinition collection
type SiteExternalColumnsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ColumnDefinition collection
func (r *SiteExternalColumnsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ColumnDefinition, error) {
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
func (r *SiteExternalColumnsCollectionRequest) GetN(ctx context.Context, n int) ([]ColumnDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ColumnDefinition collection
func (r *SiteExternalColumnsCollectionRequest) Get(ctx context.Context) ([]ColumnDefinition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ColumnDefinition collection
func (r *SiteExternalColumnsCollectionRequest) Add(ctx context.Context, reqObj *ColumnDefinition) (resObj *ColumnDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Items returns request builder for BaseItem collection
func (b *SiteRequestBuilder) Items() *SiteItemsCollectionRequestBuilder {
	bb := &SiteItemsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/items"
	return bb
}

// SiteItemsCollectionRequestBuilder is request builder for BaseItem collection
type SiteItemsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for BaseItem collection
func (b *SiteItemsCollectionRequestBuilder) Request() *SiteItemsCollectionRequest {
	return &SiteItemsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for BaseItem item
func (b *SiteItemsCollectionRequestBuilder) ID(id string) *BaseItemRequestBuilder {
	bb := &BaseItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteItemsCollectionRequest is request for BaseItem collection
type SiteItemsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for BaseItem collection
func (r *SiteItemsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]BaseItem, error) {
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
	var values []BaseItem
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
			value  []BaseItem
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

// GetN performs GET request for BaseItem collection, max N pages
func (r *SiteItemsCollectionRequest) GetN(ctx context.Context, n int) ([]BaseItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for BaseItem collection
func (r *SiteItemsCollectionRequest) Get(ctx context.Context) ([]BaseItem, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for BaseItem collection
func (r *SiteItemsCollectionRequest) Add(ctx context.Context, reqObj *BaseItem) (resObj *BaseItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Lists returns request builder for List collection
func (b *SiteRequestBuilder) Lists() *SiteListsCollectionRequestBuilder {
	bb := &SiteListsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/lists"
	return bb
}

// SiteListsCollectionRequestBuilder is request builder for List collection
type SiteListsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for List collection
func (b *SiteListsCollectionRequestBuilder) Request() *SiteListsCollectionRequest {
	return &SiteListsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for List item
func (b *SiteListsCollectionRequestBuilder) ID(id string) *ListRequestBuilder {
	bb := &ListRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteListsCollectionRequest is request for List collection
type SiteListsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for List collection
func (r *SiteListsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]List, error) {
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
func (r *SiteListsCollectionRequest) GetN(ctx context.Context, n int) ([]List, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for List collection
func (r *SiteListsCollectionRequest) Get(ctx context.Context) ([]List, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for List collection
func (r *SiteListsCollectionRequest) Add(ctx context.Context, reqObj *List) (resObj *List, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Onenote is navigation property rn
func (b *SiteRequestBuilder) Onenote() *OnenoteRequestBuilder {
	bb := &OnenoteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/onenote"
	return bb
}

// Operations returns request builder for RichLongRunningOperation collection
func (b *SiteRequestBuilder) Operations() *SiteOperationsCollectionRequestBuilder {
	bb := &SiteOperationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/operations"
	return bb
}

// SiteOperationsCollectionRequestBuilder is request builder for RichLongRunningOperation collection
type SiteOperationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for RichLongRunningOperation collection
func (b *SiteOperationsCollectionRequestBuilder) Request() *SiteOperationsCollectionRequest {
	return &SiteOperationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for RichLongRunningOperation item
func (b *SiteOperationsCollectionRequestBuilder) ID(id string) *RichLongRunningOperationRequestBuilder {
	bb := &RichLongRunningOperationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteOperationsCollectionRequest is request for RichLongRunningOperation collection
type SiteOperationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for RichLongRunningOperation collection
func (r *SiteOperationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]RichLongRunningOperation, error) {
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
func (r *SiteOperationsCollectionRequest) GetN(ctx context.Context, n int) ([]RichLongRunningOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for RichLongRunningOperation collection
func (r *SiteOperationsCollectionRequest) Get(ctx context.Context) ([]RichLongRunningOperation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for RichLongRunningOperation collection
func (r *SiteOperationsCollectionRequest) Add(ctx context.Context, reqObj *RichLongRunningOperation) (resObj *RichLongRunningOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Permissions returns request builder for Permission collection
func (b *SiteRequestBuilder) Permissions() *SitePermissionsCollectionRequestBuilder {
	bb := &SitePermissionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/permissions"
	return bb
}

// SitePermissionsCollectionRequestBuilder is request builder for Permission collection
type SitePermissionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Permission collection
func (b *SitePermissionsCollectionRequestBuilder) Request() *SitePermissionsCollectionRequest {
	return &SitePermissionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Permission item
func (b *SitePermissionsCollectionRequestBuilder) ID(id string) *PermissionRequestBuilder {
	bb := &PermissionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SitePermissionsCollectionRequest is request for Permission collection
type SitePermissionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Permission collection
func (r *SitePermissionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Permission, error) {
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
	var values []Permission
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
			value  []Permission
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

// GetN performs GET request for Permission collection, max N pages
func (r *SitePermissionsCollectionRequest) GetN(ctx context.Context, n int) ([]Permission, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Permission collection
func (r *SitePermissionsCollectionRequest) Get(ctx context.Context) ([]Permission, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Permission collection
func (r *SitePermissionsCollectionRequest) Add(ctx context.Context, reqObj *Permission) (resObj *Permission, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Sites returns request builder for Site collection
func (b *SiteRequestBuilder) Sites() *SiteSitesCollectionRequestBuilder {
	bb := &SiteSitesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sites"
	return bb
}

// SiteSitesCollectionRequestBuilder is request builder for Site collection
type SiteSitesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Site collection
func (b *SiteSitesCollectionRequestBuilder) Request() *SiteSitesCollectionRequest {
	return &SiteSitesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Site item
func (b *SiteSitesCollectionRequestBuilder) ID(id string) *SiteRequestBuilder {
	bb := &SiteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteSitesCollectionRequest is request for Site collection
type SiteSitesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Site collection
func (r *SiteSitesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Site, error) {
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
	var values []Site
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
			value  []Site
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

// GetN performs GET request for Site collection, max N pages
func (r *SiteSitesCollectionRequest) GetN(ctx context.Context, n int) ([]Site, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Site collection
func (r *SiteSitesCollectionRequest) Get(ctx context.Context) ([]Site, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Site collection
func (r *SiteSitesCollectionRequest) Add(ctx context.Context, reqObj *Site) (resObj *Site, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// TermStore is navigation property rn
func (b *SiteRequestBuilder) TermStore() *TermStoreStoreRequestBuilder {
	bb := &TermStoreStoreRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/termStore"
	return bb
}

// TermStores returns request builder for TermStoreStore collection
func (b *SiteRequestBuilder) TermStores() *SiteTermStoresCollectionRequestBuilder {
	bb := &SiteTermStoresCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/termStores"
	return bb
}

// SiteTermStoresCollectionRequestBuilder is request builder for TermStoreStore collection
type SiteTermStoresCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TermStoreStore collection
func (b *SiteTermStoresCollectionRequestBuilder) Request() *SiteTermStoresCollectionRequest {
	return &SiteTermStoresCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TermStoreStore item
func (b *SiteTermStoresCollectionRequestBuilder) ID(id string) *TermStoreStoreRequestBuilder {
	bb := &TermStoreStoreRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteTermStoresCollectionRequest is request for TermStoreStore collection
type SiteTermStoresCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TermStoreStore collection
func (r *SiteTermStoresCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TermStoreStore, error) {
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
	var values []TermStoreStore
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
			value  []TermStoreStore
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

// GetN performs GET request for TermStoreStore collection, max N pages
func (r *SiteTermStoresCollectionRequest) GetN(ctx context.Context, n int) ([]TermStoreStore, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TermStoreStore collection
func (r *SiteTermStoresCollectionRequest) Get(ctx context.Context) ([]TermStoreStore, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TermStoreStore collection
func (r *SiteTermStoresCollectionRequest) Add(ctx context.Context, reqObj *TermStoreStore) (resObj *TermStoreStore, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Site returns request builder for Site collection
func (b *SiteItemsCollectionRequestBuilder) Site() *SiteItemsCollectionSiteCollectionRequestBuilder {
	bb := &SiteItemsCollectionSiteCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// SiteItemsCollectionSiteCollectionRequestBuilder is request builder for Site collection
type SiteItemsCollectionSiteCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Site collection
func (b *SiteItemsCollectionSiteCollectionRequestBuilder) Request() *SiteItemsCollectionSiteCollectionRequest {
	return &SiteItemsCollectionSiteCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Site item
func (b *SiteItemsCollectionSiteCollectionRequestBuilder) ID(id string) *SiteRequestBuilder {
	bb := &SiteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SiteItemsCollectionSiteCollectionRequest is request for Site collection
type SiteItemsCollectionSiteCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Site collection
func (r *SiteItemsCollectionSiteCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Site, error) {
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
	var values []Site
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
			value  []Site
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

// GetN performs GET request for Site collection, max N pages
func (r *SiteItemsCollectionSiteCollectionRequest) GetN(ctx context.Context, n int) ([]Site, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Site collection
func (r *SiteItemsCollectionSiteCollectionRequest) Get(ctx context.Context) ([]Site, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Site collection
func (r *SiteItemsCollectionSiteCollectionRequest) Add(ctx context.Context, reqObj *Site) (resObj *Site, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

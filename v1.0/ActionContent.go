// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// ContentTypeCollectionAddCopyRequestParameter undocumented
type ContentTypeCollectionAddCopyRequestParameter struct {
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
}

// ContentTypeCollectionAddCopyFromContentTypeHubRequestParameter undocumented
type ContentTypeCollectionAddCopyFromContentTypeHubRequestParameter struct {
	// ContentTypeID undocumented
	ContentTypeID *string `json:"contentTypeId,omitempty"`
}

// ContentTypePublishRequestParameter undocumented
type ContentTypePublishRequestParameter struct {
}

// ContentTypeUnpublishRequestParameter undocumented
type ContentTypeUnpublishRequestParameter struct {
}

// ContentTypeAssociateWithHubSitesRequestParameter undocumented
type ContentTypeAssociateWithHubSitesRequestParameter struct {
	// HubSiteUrls undocumented
	HubSiteUrls []string `json:"hubSiteUrls,omitempty"`
	// PropagateToExistingLists undocumented
	PropagateToExistingLists *bool `json:"propagateToExistingLists,omitempty"`
}

// ContentTypeCopyToDefaultContentLocationRequestParameter undocumented
type ContentTypeCopyToDefaultContentLocationRequestParameter struct {
	// SourceFile undocumented
	SourceFile *ItemReference `json:"sourceFile,omitempty"`
	// DestinationFileName undocumented
	DestinationFileName *string `json:"destinationFileName,omitempty"`
}

// Base is navigation property rn
func (b *ContentTypeRequestBuilder) Base() *ContentTypeRequestBuilder {
	bb := &ContentTypeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/base"
	return bb
}

// BaseTypes returns request builder for ContentType collection
func (b *ContentTypeRequestBuilder) BaseTypes() *ContentTypeBaseTypesCollectionRequestBuilder {
	bb := &ContentTypeBaseTypesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/baseTypes"
	return bb
}

// ContentTypeBaseTypesCollectionRequestBuilder is request builder for ContentType collection rcn
type ContentTypeBaseTypesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ContentType collection
func (b *ContentTypeBaseTypesCollectionRequestBuilder) Request() *ContentTypeBaseTypesCollectionRequest {
	return &ContentTypeBaseTypesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ContentType item
func (b *ContentTypeBaseTypesCollectionRequestBuilder) ID(id string) *ContentTypeRequestBuilder {
	bb := &ContentTypeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ContentTypeBaseTypesCollectionRequest is request for ContentType collection
type ContentTypeBaseTypesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ContentType collection
func (r *ContentTypeBaseTypesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ContentType, error) {
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
func (r *ContentTypeBaseTypesCollectionRequest) GetN(ctx context.Context, n int) ([]ContentType, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ContentType collection
func (r *ContentTypeBaseTypesCollectionRequest) Get(ctx context.Context) ([]ContentType, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ContentType collection
func (r *ContentTypeBaseTypesCollectionRequest) Add(ctx context.Context, reqObj *ContentType) (resObj *ContentType, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ColumnLinks returns request builder for ColumnLink collection
func (b *ContentTypeRequestBuilder) ColumnLinks() *ContentTypeColumnLinksCollectionRequestBuilder {
	bb := &ContentTypeColumnLinksCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/columnLinks"
	return bb
}

// ContentTypeColumnLinksCollectionRequestBuilder is request builder for ColumnLink collection rcn
type ContentTypeColumnLinksCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ColumnLink collection
func (b *ContentTypeColumnLinksCollectionRequestBuilder) Request() *ContentTypeColumnLinksCollectionRequest {
	return &ContentTypeColumnLinksCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ColumnLink item
func (b *ContentTypeColumnLinksCollectionRequestBuilder) ID(id string) *ColumnLinkRequestBuilder {
	bb := &ColumnLinkRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ContentTypeColumnLinksCollectionRequest is request for ColumnLink collection
type ContentTypeColumnLinksCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ColumnLink collection
func (r *ContentTypeColumnLinksCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ColumnLink, error) {
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
	var values []ColumnLink
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
			value  []ColumnLink
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

// GetN performs GET request for ColumnLink collection, max N pages
func (r *ContentTypeColumnLinksCollectionRequest) GetN(ctx context.Context, n int) ([]ColumnLink, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ColumnLink collection
func (r *ContentTypeColumnLinksCollectionRequest) Get(ctx context.Context) ([]ColumnLink, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ColumnLink collection
func (r *ContentTypeColumnLinksCollectionRequest) Add(ctx context.Context, reqObj *ColumnLink) (resObj *ColumnLink, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ColumnPositions returns request builder for ColumnDefinition collection
func (b *ContentTypeRequestBuilder) ColumnPositions() *ContentTypeColumnPositionsCollectionRequestBuilder {
	bb := &ContentTypeColumnPositionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/columnPositions"
	return bb
}

// ContentTypeColumnPositionsCollectionRequestBuilder is request builder for ColumnDefinition collection rcn
type ContentTypeColumnPositionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ColumnDefinition collection
func (b *ContentTypeColumnPositionsCollectionRequestBuilder) Request() *ContentTypeColumnPositionsCollectionRequest {
	return &ContentTypeColumnPositionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ColumnDefinition item
func (b *ContentTypeColumnPositionsCollectionRequestBuilder) ID(id string) *ColumnDefinitionRequestBuilder {
	bb := &ColumnDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ContentTypeColumnPositionsCollectionRequest is request for ColumnDefinition collection
type ContentTypeColumnPositionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ColumnDefinition collection
func (r *ContentTypeColumnPositionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ColumnDefinition, error) {
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
func (r *ContentTypeColumnPositionsCollectionRequest) GetN(ctx context.Context, n int) ([]ColumnDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ColumnDefinition collection
func (r *ContentTypeColumnPositionsCollectionRequest) Get(ctx context.Context) ([]ColumnDefinition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ColumnDefinition collection
func (r *ContentTypeColumnPositionsCollectionRequest) Add(ctx context.Context, reqObj *ColumnDefinition) (resObj *ColumnDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Columns returns request builder for ColumnDefinition collection
func (b *ContentTypeRequestBuilder) Columns() *ContentTypeColumnsCollectionRequestBuilder {
	bb := &ContentTypeColumnsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/columns"
	return bb
}

// ContentTypeColumnsCollectionRequestBuilder is request builder for ColumnDefinition collection rcn
type ContentTypeColumnsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ColumnDefinition collection
func (b *ContentTypeColumnsCollectionRequestBuilder) Request() *ContentTypeColumnsCollectionRequest {
	return &ContentTypeColumnsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ColumnDefinition item
func (b *ContentTypeColumnsCollectionRequestBuilder) ID(id string) *ColumnDefinitionRequestBuilder {
	bb := &ColumnDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ContentTypeColumnsCollectionRequest is request for ColumnDefinition collection
type ContentTypeColumnsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ColumnDefinition collection
func (r *ContentTypeColumnsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ColumnDefinition, error) {
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
func (r *ContentTypeColumnsCollectionRequest) GetN(ctx context.Context, n int) ([]ColumnDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ColumnDefinition collection
func (r *ContentTypeColumnsCollectionRequest) Get(ctx context.Context) ([]ColumnDefinition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ColumnDefinition collection
func (r *ContentTypeColumnsCollectionRequest) Add(ctx context.Context, reqObj *ColumnDefinition) (resObj *ColumnDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *ContentSharingSessionRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *ContentTypeRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

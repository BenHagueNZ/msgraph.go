// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// MobileAppAssignRequestParameter undocumented
type MobileAppAssignRequestParameter struct {
	// MobileAppAssignments undocumented
	MobileAppAssignments []MobileAppAssignment `json:"mobileAppAssignments,omitempty"`
}

// MobileAppContentFileCommitRequestParameter undocumented
type MobileAppContentFileCommitRequestParameter struct {
	// FileEncryptionInfo undocumented
	FileEncryptionInfo *FileEncryptionInfo `json:"fileEncryptionInfo,omitempty"`
}

// MobileAppContentFileRenewUploadRequestParameter undocumented
type MobileAppContentFileRenewUploadRequestParameter struct {
}

// Assignments returns request builder for MobileAppAssignment collection
func (b *MobileAppRequestBuilder) Assignments() *MobileAppAssignmentsCollectionRequestBuilder {
	bb := &MobileAppAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// MobileAppAssignmentsCollectionRequestBuilder is request builder for MobileAppAssignment collection
type MobileAppAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppAssignment collection
func (b *MobileAppAssignmentsCollectionRequestBuilder) Request() *MobileAppAssignmentsCollectionRequest {
	return &MobileAppAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppAssignment item
func (b *MobileAppAssignmentsCollectionRequestBuilder) ID(id string) *MobileAppAssignmentRequestBuilder {
	bb := &MobileAppAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppAssignmentsCollectionRequest is request for MobileAppAssignment collection
type MobileAppAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileAppAssignment collection
func (r *MobileAppAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MobileAppAssignment, error) {
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
	var values []MobileAppAssignment
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
			value  []MobileAppAssignment
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

// GetN performs GET request for MobileAppAssignment collection, max N pages
func (r *MobileAppAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]MobileAppAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MobileAppAssignment collection
func (r *MobileAppAssignmentsCollectionRequest) Get(ctx context.Context) ([]MobileAppAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MobileAppAssignment collection
func (r *MobileAppAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *MobileAppAssignment) (resObj *MobileAppAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Categories returns request builder for MobileAppCategory collection
func (b *MobileAppRequestBuilder) Categories() *MobileAppCategoriesCollectionRequestBuilder {
	bb := &MobileAppCategoriesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/categories"
	return bb
}

// MobileAppCategoriesCollectionRequestBuilder is request builder for MobileAppCategory collection
type MobileAppCategoriesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppCategory collection
func (b *MobileAppCategoriesCollectionRequestBuilder) Request() *MobileAppCategoriesCollectionRequest {
	return &MobileAppCategoriesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppCategory item
func (b *MobileAppCategoriesCollectionRequestBuilder) ID(id string) *MobileAppCategoryRequestBuilder {
	bb := &MobileAppCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppCategoriesCollectionRequest is request for MobileAppCategory collection
type MobileAppCategoriesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileAppCategory collection
func (r *MobileAppCategoriesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MobileAppCategory, error) {
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
	var values []MobileAppCategory
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
			value  []MobileAppCategory
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

// GetN performs GET request for MobileAppCategory collection, max N pages
func (r *MobileAppCategoriesCollectionRequest) GetN(ctx context.Context, n int) ([]MobileAppCategory, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MobileAppCategory collection
func (r *MobileAppCategoriesCollectionRequest) Get(ctx context.Context) ([]MobileAppCategory, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MobileAppCategory collection
func (r *MobileAppCategoriesCollectionRequest) Add(ctx context.Context, reqObj *MobileAppCategory) (resObj *MobileAppCategory, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Files returns request builder for MobileAppContentFile collection
func (b *MobileAppContentRequestBuilder) Files() *MobileAppContentFilesCollectionRequestBuilder {
	bb := &MobileAppContentFilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/files"
	return bb
}

// MobileAppContentFilesCollectionRequestBuilder is request builder for MobileAppContentFile collection
type MobileAppContentFilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppContentFile collection
func (b *MobileAppContentFilesCollectionRequestBuilder) Request() *MobileAppContentFilesCollectionRequest {
	return &MobileAppContentFilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppContentFile item
func (b *MobileAppContentFilesCollectionRequestBuilder) ID(id string) *MobileAppContentFileRequestBuilder {
	bb := &MobileAppContentFileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppContentFilesCollectionRequest is request for MobileAppContentFile collection
type MobileAppContentFilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileAppContentFile collection
func (r *MobileAppContentFilesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MobileAppContentFile, error) {
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
	var values []MobileAppContentFile
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
			value  []MobileAppContentFile
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

// GetN performs GET request for MobileAppContentFile collection, max N pages
func (r *MobileAppContentFilesCollectionRequest) GetN(ctx context.Context, n int) ([]MobileAppContentFile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MobileAppContentFile collection
func (r *MobileAppContentFilesCollectionRequest) Get(ctx context.Context) ([]MobileAppContentFile, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MobileAppContentFile collection
func (r *MobileAppContentFilesCollectionRequest) Add(ctx context.Context, reqObj *MobileAppContentFile) (resObj *MobileAppContentFile, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ContentVersions returns request builder for MobileAppContent collection
func (b *MobileLobAppRequestBuilder) ContentVersions() *MobileLobAppContentVersionsCollectionRequestBuilder {
	bb := &MobileLobAppContentVersionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/contentVersions"
	return bb
}

// MobileLobAppContentVersionsCollectionRequestBuilder is request builder for MobileAppContent collection
type MobileLobAppContentVersionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppContent collection
func (b *MobileLobAppContentVersionsCollectionRequestBuilder) Request() *MobileLobAppContentVersionsCollectionRequest {
	return &MobileLobAppContentVersionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppContent item
func (b *MobileLobAppContentVersionsCollectionRequestBuilder) ID(id string) *MobileAppContentRequestBuilder {
	bb := &MobileAppContentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileLobAppContentVersionsCollectionRequest is request for MobileAppContent collection
type MobileLobAppContentVersionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileAppContent collection
func (r *MobileLobAppContentVersionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MobileAppContent, error) {
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
	var values []MobileAppContent
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
			value  []MobileAppContent
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

// GetN performs GET request for MobileAppContent collection, max N pages
func (r *MobileLobAppContentVersionsCollectionRequest) GetN(ctx context.Context, n int) ([]MobileAppContent, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for MobileAppContent collection
func (r *MobileLobAppContentVersionsCollectionRequest) Get(ctx context.Context) ([]MobileAppContent, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for MobileAppContent collection
func (r *MobileLobAppContentVersionsCollectionRequest) Add(ctx context.Context, reqObj *MobileAppContent) (resObj *MobileAppContent, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

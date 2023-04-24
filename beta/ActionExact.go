// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// ExactMatchDataStoreLookupRequestParameter undocumented
type ExactMatchDataStoreLookupRequestParameter struct {
	// Key undocumented
	Key *string `json:"key,omitempty"`
	// Values undocumented
	Values []string `json:"values,omitempty"`
	// ResultColumnNames undocumented
	ResultColumnNames []string `json:"resultColumnNames,omitempty"`
}

// ExactMatchSessionCancelRequestParameter undocumented
type ExactMatchSessionCancelRequestParameter struct {
}

// ExactMatchSessionRenewRequestParameter undocumented
type ExactMatchSessionRenewRequestParameter struct {
}

// ExactMatchSessionCommitRequestParameter undocumented
type ExactMatchSessionCommitRequestParameter struct {
}

// Sessions returns request builder for ExactMatchSession collection
func (b *ExactMatchDataStoreRequestBuilder) Sessions() *ExactMatchDataStoreSessionsCollectionRequestBuilder {
	bb := &ExactMatchDataStoreSessionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sessions"
	return bb
}

// ExactMatchDataStoreSessionsCollectionRequestBuilder is request builder for ExactMatchSession collection rcn
type ExactMatchDataStoreSessionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ExactMatchSession collection
func (b *ExactMatchDataStoreSessionsCollectionRequestBuilder) Request() *ExactMatchDataStoreSessionsCollectionRequest {
	return &ExactMatchDataStoreSessionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ExactMatchSession item
func (b *ExactMatchDataStoreSessionsCollectionRequestBuilder) ID(id string) *ExactMatchSessionRequestBuilder {
	bb := &ExactMatchSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ExactMatchDataStoreSessionsCollectionRequest is request for ExactMatchSession collection
type ExactMatchDataStoreSessionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ExactMatchSession collection
func (r *ExactMatchDataStoreSessionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ExactMatchSession, error) {
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
	var values []ExactMatchSession
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
			value  []ExactMatchSession
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

// GetN performs GET request for ExactMatchSession collection, max N pages
func (r *ExactMatchDataStoreSessionsCollectionRequest) GetN(ctx context.Context, n int) ([]ExactMatchSession, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ExactMatchSession collection
func (r *ExactMatchDataStoreSessionsCollectionRequest) Get(ctx context.Context) ([]ExactMatchSession, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ExactMatchSession collection
func (r *ExactMatchDataStoreSessionsCollectionRequest) Add(ctx context.Context, reqObj *ExactMatchSession) (resObj *ExactMatchSession, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// MatchingRows returns request builder for LookupResultRow collection
func (b *ExactMatchLookupJobRequestBuilder) MatchingRows() *ExactMatchLookupJobMatchingRowsCollectionRequestBuilder {
	bb := &ExactMatchLookupJobMatchingRowsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/matchingRows"
	return bb
}

// ExactMatchLookupJobMatchingRowsCollectionRequestBuilder is request builder for LookupResultRow collection rcn
type ExactMatchLookupJobMatchingRowsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for LookupResultRow collection
func (b *ExactMatchLookupJobMatchingRowsCollectionRequestBuilder) Request() *ExactMatchLookupJobMatchingRowsCollectionRequest {
	return &ExactMatchLookupJobMatchingRowsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for LookupResultRow item
func (b *ExactMatchLookupJobMatchingRowsCollectionRequestBuilder) ID(id string) *LookupResultRowRequestBuilder {
	bb := &LookupResultRowRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ExactMatchLookupJobMatchingRowsCollectionRequest is request for LookupResultRow collection
type ExactMatchLookupJobMatchingRowsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for LookupResultRow collection
func (r *ExactMatchLookupJobMatchingRowsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]LookupResultRow, error) {
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
	var values []LookupResultRow
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
			value  []LookupResultRow
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

// GetN performs GET request for LookupResultRow collection, max N pages
func (r *ExactMatchLookupJobMatchingRowsCollectionRequest) GetN(ctx context.Context, n int) ([]LookupResultRow, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for LookupResultRow collection
func (r *ExactMatchLookupJobMatchingRowsCollectionRequest) Get(ctx context.Context) ([]LookupResultRow, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for LookupResultRow collection
func (r *ExactMatchLookupJobMatchingRowsCollectionRequest) Add(ctx context.Context, reqObj *LookupResultRow) (resObj *LookupResultRow, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// UploadAgent is navigation property rn
func (b *ExactMatchSessionRequestBuilder) UploadAgent() *ExactMatchUploadAgentRequestBuilder {
	bb := &ExactMatchUploadAgentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/uploadAgent"
	return bb
}

// Entity is navigation property rn
func (b *ExactMatchDataStoreBaseRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *ExactMatchJobBaseRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *ExactMatchUploadAgentRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// TaskProcessingResults returns request builder for TaskProcessingResult collection
func (b *RunRequestBuilder) TaskProcessingResults() *RunTaskProcessingResultsCollectionRequestBuilder {
	bb := &RunTaskProcessingResultsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/taskProcessingResults"
	return bb
}

// RunTaskProcessingResultsCollectionRequestBuilder is request builder for TaskProcessingResult collection
type RunTaskProcessingResultsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TaskProcessingResult collection
func (b *RunTaskProcessingResultsCollectionRequestBuilder) Request() *RunTaskProcessingResultsCollectionRequest {
	return &RunTaskProcessingResultsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TaskProcessingResult item
func (b *RunTaskProcessingResultsCollectionRequestBuilder) ID(id string) *TaskProcessingResultRequestBuilder {
	bb := &TaskProcessingResultRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RunTaskProcessingResultsCollectionRequest is request for TaskProcessingResult collection
type RunTaskProcessingResultsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TaskProcessingResult collection
func (r *RunTaskProcessingResultsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TaskProcessingResult, error) {
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
	var values []TaskProcessingResult
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
			value  []TaskProcessingResult
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

// GetN performs GET request for TaskProcessingResult collection, max N pages
func (r *RunTaskProcessingResultsCollectionRequest) GetN(ctx context.Context, n int) ([]TaskProcessingResult, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TaskProcessingResult collection
func (r *RunTaskProcessingResultsCollectionRequest) Get(ctx context.Context) ([]TaskProcessingResult, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TaskProcessingResult collection
func (r *RunTaskProcessingResultsCollectionRequest) Add(ctx context.Context, reqObj *TaskProcessingResult) (resObj *TaskProcessingResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// UserProcessingResults returns request builder for UserProcessingResult collection
func (b *RunRequestBuilder) UserProcessingResults() *RunUserProcessingResultsCollectionRequestBuilder {
	bb := &RunUserProcessingResultsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/userProcessingResults"
	return bb
}

// RunUserProcessingResultsCollectionRequestBuilder is request builder for UserProcessingResult collection
type RunUserProcessingResultsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UserProcessingResult collection
func (b *RunUserProcessingResultsCollectionRequestBuilder) Request() *RunUserProcessingResultsCollectionRequest {
	return &RunUserProcessingResultsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UserProcessingResult item
func (b *RunUserProcessingResultsCollectionRequestBuilder) ID(id string) *UserProcessingResultRequestBuilder {
	bb := &UserProcessingResultRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RunUserProcessingResultsCollectionRequest is request for UserProcessingResult collection
type RunUserProcessingResultsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UserProcessingResult collection
func (r *RunUserProcessingResultsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]UserProcessingResult, error) {
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
	var values []UserProcessingResult
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
			value  []UserProcessingResult
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

// GetN performs GET request for UserProcessingResult collection, max N pages
func (r *RunUserProcessingResultsCollectionRequest) GetN(ctx context.Context, n int) ([]UserProcessingResult, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for UserProcessingResult collection
func (r *RunUserProcessingResultsCollectionRequest) Get(ctx context.Context) ([]UserProcessingResult, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for UserProcessingResult collection
func (r *RunUserProcessingResultsCollectionRequest) Add(ctx context.Context, reqObj *UserProcessingResult) (resObj *UserProcessingResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

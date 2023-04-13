// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// StopHoldMusicOperation returns request builder for StopHoldMusicOperation collection
func (b *CallOperationsCollectionRequestBuilder) StopHoldMusicOperation() *CallOperationsCollectionStopHoldMusicOperationCollectionRequestBuilder {
	bb := &CallOperationsCollectionStopHoldMusicOperationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// CallOperationsCollectionStopHoldMusicOperationCollectionRequestBuilder is request builder for StopHoldMusicOperation collection
type CallOperationsCollectionStopHoldMusicOperationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for StopHoldMusicOperation collection
func (b *CallOperationsCollectionStopHoldMusicOperationCollectionRequestBuilder) Request() *CallOperationsCollectionStopHoldMusicOperationCollectionRequest {
	return &CallOperationsCollectionStopHoldMusicOperationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for StopHoldMusicOperation item
func (b *CallOperationsCollectionStopHoldMusicOperationCollectionRequestBuilder) ID(id string) *StopHoldMusicOperationRequestBuilder {
	bb := &StopHoldMusicOperationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallOperationsCollectionStopHoldMusicOperationCollectionRequest is request for StopHoldMusicOperation collection
type CallOperationsCollectionStopHoldMusicOperationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for StopHoldMusicOperation collection
func (r *CallOperationsCollectionStopHoldMusicOperationCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]StopHoldMusicOperation, error) {
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
	var values []StopHoldMusicOperation
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
			value  []StopHoldMusicOperation
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

// GetN performs GET request for StopHoldMusicOperation collection, max N pages
func (r *CallOperationsCollectionStopHoldMusicOperationCollectionRequest) GetN(ctx context.Context, n int) ([]StopHoldMusicOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for StopHoldMusicOperation collection
func (r *CallOperationsCollectionStopHoldMusicOperationCollectionRequest) Get(ctx context.Context) ([]StopHoldMusicOperation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for StopHoldMusicOperation collection
func (r *CallOperationsCollectionStopHoldMusicOperationCollectionRequest) Add(ctx context.Context, reqObj *StopHoldMusicOperation) (resObj *StopHoldMusicOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

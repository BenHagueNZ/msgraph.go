// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// RecordOperation returns request builder for RecordOperation collection
func (b *CallOperationsCollectionRequestBuilder) RecordOperation() *CallOperationsCollectionRecordOperationCollectionRequestBuilder {
	bb := &CallOperationsCollectionRecordOperationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// CallOperationsCollectionRecordOperationCollectionRequestBuilder is request builder for RecordOperation collection rcn
type CallOperationsCollectionRecordOperationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for RecordOperation collection
func (b *CallOperationsCollectionRecordOperationCollectionRequestBuilder) Request() *CallOperationsCollectionRecordOperationCollectionRequest {
	return &CallOperationsCollectionRecordOperationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for RecordOperation item
func (b *CallOperationsCollectionRecordOperationCollectionRequestBuilder) ID(id string) *RecordOperationRequestBuilder {
	bb := &RecordOperationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallOperationsCollectionRecordOperationCollectionRequest is request for RecordOperation collection
type CallOperationsCollectionRecordOperationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for RecordOperation collection
func (r *CallOperationsCollectionRecordOperationCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]RecordOperation, error) {
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
	var values []RecordOperation
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
			value  []RecordOperation
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

// GetN performs GET request for RecordOperation collection, max N pages
func (r *CallOperationsCollectionRecordOperationCollectionRequest) GetN(ctx context.Context, n int) ([]RecordOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for RecordOperation collection
func (r *CallOperationsCollectionRecordOperationCollectionRequest) Get(ctx context.Context) ([]RecordOperation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for RecordOperation collection
func (r *CallOperationsCollectionRecordOperationCollectionRequest) Add(ctx context.Context, reqObj *RecordOperation) (resObj *RecordOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

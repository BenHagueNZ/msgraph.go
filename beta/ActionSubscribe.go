// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// SubscribeToToneOperation returns request builder for SubscribeToToneOperation collection
func (b *CallOperationsCollectionRequestBuilder) SubscribeToToneOperation() *CallOperationsCollectionSubscribeToToneOperationCollectionRequestBuilder {
	bb := &CallOperationsCollectionSubscribeToToneOperationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// CallOperationsCollectionSubscribeToToneOperationCollectionRequestBuilder is request builder for SubscribeToToneOperation collection rcn
type CallOperationsCollectionSubscribeToToneOperationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SubscribeToToneOperation collection
func (b *CallOperationsCollectionSubscribeToToneOperationCollectionRequestBuilder) Request() *CallOperationsCollectionSubscribeToToneOperationCollectionRequest {
	return &CallOperationsCollectionSubscribeToToneOperationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SubscribeToToneOperation item
func (b *CallOperationsCollectionSubscribeToToneOperationCollectionRequestBuilder) ID(id string) *SubscribeToToneOperationRequestBuilder {
	bb := &SubscribeToToneOperationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallOperationsCollectionSubscribeToToneOperationCollectionRequest is request for SubscribeToToneOperation collection
type CallOperationsCollectionSubscribeToToneOperationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SubscribeToToneOperation collection
func (r *CallOperationsCollectionSubscribeToToneOperationCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SubscribeToToneOperation, error) {
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
	var values []SubscribeToToneOperation
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
			value  []SubscribeToToneOperation
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

// GetN performs GET request for SubscribeToToneOperation collection, max N pages
func (r *CallOperationsCollectionSubscribeToToneOperationCollectionRequest) GetN(ctx context.Context, n int) ([]SubscribeToToneOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SubscribeToToneOperation collection
func (r *CallOperationsCollectionSubscribeToToneOperationCollectionRequest) Get(ctx context.Context) ([]SubscribeToToneOperation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SubscribeToToneOperation collection
func (r *CallOperationsCollectionSubscribeToToneOperationCollectionRequest) Add(ctx context.Context, reqObj *SubscribeToToneOperation) (resObj *SubscribeToToneOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

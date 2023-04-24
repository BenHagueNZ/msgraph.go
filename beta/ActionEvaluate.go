// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// EvaluateLabelJobResponse returns request builder for EvaluateLabelJobResponse collection
func (b *DataClassificationServiceClassifyFileJobsCollectionRequestBuilder) EvaluateLabelJobResponse() *DataClassificationServiceClassifyFileJobsCollectionEvaluateLabelJobResponseCollectionRequestBuilder {
	bb := &DataClassificationServiceClassifyFileJobsCollectionEvaluateLabelJobResponseCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// DataClassificationServiceClassifyFileJobsCollectionEvaluateLabelJobResponseCollectionRequestBuilder is request builder for EvaluateLabelJobResponse collection rcn
type DataClassificationServiceClassifyFileJobsCollectionEvaluateLabelJobResponseCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EvaluateLabelJobResponse collection
func (b *DataClassificationServiceClassifyFileJobsCollectionEvaluateLabelJobResponseCollectionRequestBuilder) Request() *DataClassificationServiceClassifyFileJobsCollectionEvaluateLabelJobResponseCollectionRequest {
	return &DataClassificationServiceClassifyFileJobsCollectionEvaluateLabelJobResponseCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EvaluateLabelJobResponse item
func (b *DataClassificationServiceClassifyFileJobsCollectionEvaluateLabelJobResponseCollectionRequestBuilder) ID(id string) *EvaluateLabelJobResponseRequestBuilder {
	bb := &EvaluateLabelJobResponseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DataClassificationServiceClassifyFileJobsCollectionEvaluateLabelJobResponseCollectionRequest is request for EvaluateLabelJobResponse collection
type DataClassificationServiceClassifyFileJobsCollectionEvaluateLabelJobResponseCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EvaluateLabelJobResponse collection
func (r *DataClassificationServiceClassifyFileJobsCollectionEvaluateLabelJobResponseCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EvaluateLabelJobResponse, error) {
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
	var values []EvaluateLabelJobResponse
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
			value  []EvaluateLabelJobResponse
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

// GetN performs GET request for EvaluateLabelJobResponse collection, max N pages
func (r *DataClassificationServiceClassifyFileJobsCollectionEvaluateLabelJobResponseCollectionRequest) GetN(ctx context.Context, n int) ([]EvaluateLabelJobResponse, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EvaluateLabelJobResponse collection
func (r *DataClassificationServiceClassifyFileJobsCollectionEvaluateLabelJobResponseCollectionRequest) Get(ctx context.Context) ([]EvaluateLabelJobResponse, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EvaluateLabelJobResponse collection
func (r *DataClassificationServiceClassifyFileJobsCollectionEvaluateLabelJobResponseCollectionRequest) Add(ctx context.Context, reqObj *EvaluateLabelJobResponse) (resObj *EvaluateLabelJobResponse, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
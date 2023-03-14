// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// Stages returns request builder for ApprovalStage collection
func (b *ApprovalRequestBuilder) Stages() *ApprovalStagesCollectionRequestBuilder {
	bb := &ApprovalStagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/stages"
	return bb
}

// ApprovalStagesCollectionRequestBuilder is request builder for ApprovalStage collection
type ApprovalStagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ApprovalStage collection
func (b *ApprovalStagesCollectionRequestBuilder) Request() *ApprovalStagesCollectionRequest {
	return &ApprovalStagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ApprovalStage item
func (b *ApprovalStagesCollectionRequestBuilder) ID(id string) *ApprovalStageRequestBuilder {
	bb := &ApprovalStageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApprovalStagesCollectionRequest is request for ApprovalStage collection
type ApprovalStagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ApprovalStage collection
func (r *ApprovalStagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ApprovalStage, error) {
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
	var values []ApprovalStage
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
			value  []ApprovalStage
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

// GetN performs GET request for ApprovalStage collection, max N pages
func (r *ApprovalStagesCollectionRequest) GetN(ctx context.Context, n int) ([]ApprovalStage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ApprovalStage collection
func (r *ApprovalStagesCollectionRequest) Get(ctx context.Context) ([]ApprovalStage, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ApprovalStage collection
func (r *ApprovalStagesCollectionRequest) Add(ctx context.Context, reqObj *ApprovalStage) (resObj *ApprovalStage, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
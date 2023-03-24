// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// SearchAggregationRequestBuilder is request builder for SearchAggregation
type SearchAggregationRequestBuilder struct{ BaseRequestBuilder }

// Request returns SearchAggregationRequest
func (b *SearchAggregationRequestBuilder) Request() *SearchAggregationRequest {
	return &SearchAggregationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SearchAggregationRequest is request for SearchAggregation
type SearchAggregationRequest struct{ BaseRequest }

// Get performs GET request for SearchAggregation
func (r *SearchAggregationRequest) Get(ctx context.Context) (resObj *SearchAggregation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SearchAggregation
func (r *SearchAggregationRequest) Update(ctx context.Context, reqObj *SearchAggregation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SearchAggregation
func (r *SearchAggregationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SearchAlterationRequestBuilder is request builder for SearchAlteration
type SearchAlterationRequestBuilder struct{ BaseRequestBuilder }

// Request returns SearchAlterationRequest
func (b *SearchAlterationRequestBuilder) Request() *SearchAlterationRequest {
	return &SearchAlterationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SearchAlterationRequest is request for SearchAlteration
type SearchAlterationRequest struct{ BaseRequest }

// Get performs GET request for SearchAlteration
func (r *SearchAlterationRequest) Get(ctx context.Context) (resObj *SearchAlteration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SearchAlteration
func (r *SearchAlterationRequest) Update(ctx context.Context, reqObj *SearchAlteration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SearchAlteration
func (r *SearchAlterationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SearchAlterationOptionsRequestBuilder is request builder for SearchAlterationOptions
type SearchAlterationOptionsRequestBuilder struct{ BaseRequestBuilder }

// Request returns SearchAlterationOptionsRequest
func (b *SearchAlterationOptionsRequestBuilder) Request() *SearchAlterationOptionsRequest {
	return &SearchAlterationOptionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SearchAlterationOptionsRequest is request for SearchAlterationOptions
type SearchAlterationOptionsRequest struct{ BaseRequest }

// Get performs GET request for SearchAlterationOptions
func (r *SearchAlterationOptionsRequest) Get(ctx context.Context) (resObj *SearchAlterationOptions, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SearchAlterationOptions
func (r *SearchAlterationOptionsRequest) Update(ctx context.Context, reqObj *SearchAlterationOptions) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SearchAlterationOptions
func (r *SearchAlterationOptionsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SearchBucketRequestBuilder is request builder for SearchBucket
type SearchBucketRequestBuilder struct{ BaseRequestBuilder }

// Request returns SearchBucketRequest
func (b *SearchBucketRequestBuilder) Request() *SearchBucketRequest {
	return &SearchBucketRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SearchBucketRequest is request for SearchBucket
type SearchBucketRequest struct{ BaseRequest }

// Get performs GET request for SearchBucket
func (r *SearchBucketRequest) Get(ctx context.Context) (resObj *SearchBucket, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SearchBucket
func (r *SearchBucketRequest) Update(ctx context.Context, reqObj *SearchBucket) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SearchBucket
func (r *SearchBucketRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SearchEntityRequestBuilder is request builder for SearchEntity
type SearchEntityRequestBuilder struct{ BaseRequestBuilder }

// Request returns SearchEntityRequest
func (b *SearchEntityRequestBuilder) Request() *SearchEntityRequest {
	return &SearchEntityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SearchEntityRequest is request for SearchEntity
type SearchEntityRequest struct{ BaseRequest }

// Get performs GET request for SearchEntity
func (r *SearchEntityRequest) Get(ctx context.Context) (resObj *SearchEntity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SearchEntity
func (r *SearchEntityRequest) Update(ctx context.Context, reqObj *SearchEntity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SearchEntity
func (r *SearchEntityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SearchHitRequestBuilder is request builder for SearchHit
type SearchHitRequestBuilder struct{ BaseRequestBuilder }

// Request returns SearchHitRequest
func (b *SearchHitRequestBuilder) Request() *SearchHitRequest {
	return &SearchHitRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SearchHitRequest is request for SearchHit
type SearchHitRequest struct{ BaseRequest }

// Get performs GET request for SearchHit
func (r *SearchHitRequest) Get(ctx context.Context) (resObj *SearchHit, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SearchHit
func (r *SearchHitRequest) Update(ctx context.Context, reqObj *SearchHit) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SearchHit
func (r *SearchHitRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SearchHitsContainerRequestBuilder is request builder for SearchHitsContainer
type SearchHitsContainerRequestBuilder struct{ BaseRequestBuilder }

// Request returns SearchHitsContainerRequest
func (b *SearchHitsContainerRequestBuilder) Request() *SearchHitsContainerRequest {
	return &SearchHitsContainerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SearchHitsContainerRequest is request for SearchHitsContainer
type SearchHitsContainerRequest struct{ BaseRequest }

// Get performs GET request for SearchHitsContainer
func (r *SearchHitsContainerRequest) Get(ctx context.Context) (resObj *SearchHitsContainer, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SearchHitsContainer
func (r *SearchHitsContainerRequest) Update(ctx context.Context, reqObj *SearchHitsContainer) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SearchHitsContainer
func (r *SearchHitsContainerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SearchQueryRequestBuilder is request builder for SearchQuery
type SearchQueryRequestBuilder struct{ BaseRequestBuilder }

// Request returns SearchQueryRequest
func (b *SearchQueryRequestBuilder) Request() *SearchQueryRequest {
	return &SearchQueryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SearchQueryRequest is request for SearchQuery
type SearchQueryRequest struct{ BaseRequest }

// Get performs GET request for SearchQuery
func (r *SearchQueryRequest) Get(ctx context.Context) (resObj *SearchQuery, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SearchQuery
func (r *SearchQueryRequest) Update(ctx context.Context, reqObj *SearchQuery) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SearchQuery
func (r *SearchQueryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SearchRequestObjectRequestBuilder is request builder for SearchRequestObject
type SearchRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns SearchRequestObjectRequest
func (b *SearchRequestObjectRequestBuilder) Request() *SearchRequestObjectRequest {
	return &SearchRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SearchRequestObjectRequest is request for SearchRequestObject
type SearchRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for SearchRequestObject
func (r *SearchRequestObjectRequest) Get(ctx context.Context) (resObj *SearchRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SearchRequestObject
func (r *SearchRequestObjectRequest) Update(ctx context.Context, reqObj *SearchRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SearchRequestObject
func (r *SearchRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SearchResponseRequestBuilder is request builder for SearchResponse
type SearchResponseRequestBuilder struct{ BaseRequestBuilder }

// Request returns SearchResponseRequest
func (b *SearchResponseRequestBuilder) Request() *SearchResponseRequest {
	return &SearchResponseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SearchResponseRequest is request for SearchResponse
type SearchResponseRequest struct{ BaseRequest }

// Get performs GET request for SearchResponse
func (r *SearchResponseRequest) Get(ctx context.Context) (resObj *SearchResponse, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SearchResponse
func (r *SearchResponseRequest) Update(ctx context.Context, reqObj *SearchResponse) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SearchResponse
func (r *SearchResponseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SearchResultRequestBuilder is request builder for SearchResult
type SearchResultRequestBuilder struct{ BaseRequestBuilder }

// Request returns SearchResultRequest
func (b *SearchResultRequestBuilder) Request() *SearchResultRequest {
	return &SearchResultRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SearchResultRequest is request for SearchResult
type SearchResultRequest struct{ BaseRequest }

// Get performs GET request for SearchResult
func (r *SearchResultRequest) Get(ctx context.Context) (resObj *SearchResult, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SearchResult
func (r *SearchResultRequest) Update(ctx context.Context, reqObj *SearchResult) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SearchResult
func (r *SearchResultRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type SearchEntityQueryRequestBuilder struct{ BaseRequestBuilder }

// Query action undocumented
func (b *SearchEntityRequestBuilder) Query(reqObj *SearchEntityQueryRequestParameter) *SearchEntityQueryRequestBuilder {
	bb := &SearchEntityQueryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Query"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SearchEntityQueryRequest struct{ BaseRequest }

func (b *SearchEntityQueryRequestBuilder) Request() *SearchEntityQueryRequest {
	return &SearchEntityQueryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SearchEntityQueryRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SearchResponse, error) {
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
	var values []SearchResponse
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
			value  []SearchResponse
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

func (r *SearchEntityQueryRequest) PostN(ctx context.Context, n int) ([]SearchResponse, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

func (r *SearchEntityQueryRequest) Post(ctx context.Context) ([]SearchResponse, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

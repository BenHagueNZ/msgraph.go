// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// Buckets returns request builder for PlannerBucket collection
func (b *PlannerRequestBuilder) Buckets() *PlannerBucketsCollectionRequestBuilder {
	bb := &PlannerBucketsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/buckets"
	return bb
}

// PlannerBucketsCollectionRequestBuilder is request builder for PlannerBucket collection
type PlannerBucketsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerBucket collection
func (b *PlannerBucketsCollectionRequestBuilder) Request() *PlannerBucketsCollectionRequest {
	return &PlannerBucketsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerBucket item
func (b *PlannerBucketsCollectionRequestBuilder) ID(id string) *PlannerBucketRequestBuilder {
	bb := &PlannerBucketRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerBucketsCollectionRequest is request for PlannerBucket collection
type PlannerBucketsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PlannerBucket collection
func (r *PlannerBucketsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PlannerBucket, error) {
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
	var values []PlannerBucket
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
			value  []PlannerBucket
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

// GetN performs GET request for PlannerBucket collection, max N pages
func (r *PlannerBucketsCollectionRequest) GetN(ctx context.Context, n int) ([]PlannerBucket, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PlannerBucket collection
func (r *PlannerBucketsCollectionRequest) Get(ctx context.Context) ([]PlannerBucket, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PlannerBucket collection
func (r *PlannerBucketsCollectionRequest) Add(ctx context.Context, reqObj *PlannerBucket) (resObj *PlannerBucket, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Plans returns request builder for PlannerPlan collection
func (b *PlannerRequestBuilder) Plans() *PlannerPlansCollectionRequestBuilder {
	bb := &PlannerPlansCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/plans"
	return bb
}

// PlannerPlansCollectionRequestBuilder is request builder for PlannerPlan collection
type PlannerPlansCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerPlan collection
func (b *PlannerPlansCollectionRequestBuilder) Request() *PlannerPlansCollectionRequest {
	return &PlannerPlansCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerPlan item
func (b *PlannerPlansCollectionRequestBuilder) ID(id string) *PlannerPlanRequestBuilder {
	bb := &PlannerPlanRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerPlansCollectionRequest is request for PlannerPlan collection
type PlannerPlansCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PlannerPlan collection
func (r *PlannerPlansCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PlannerPlan, error) {
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
	var values []PlannerPlan
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
			value  []PlannerPlan
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

// GetN performs GET request for PlannerPlan collection, max N pages
func (r *PlannerPlansCollectionRequest) GetN(ctx context.Context, n int) ([]PlannerPlan, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PlannerPlan collection
func (r *PlannerPlansCollectionRequest) Get(ctx context.Context) ([]PlannerPlan, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PlannerPlan collection
func (r *PlannerPlansCollectionRequest) Add(ctx context.Context, reqObj *PlannerPlan) (resObj *PlannerPlan, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Tasks returns request builder for PlannerTask collection
func (b *PlannerRequestBuilder) Tasks() *PlannerTasksCollectionRequestBuilder {
	bb := &PlannerTasksCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tasks"
	return bb
}

// PlannerTasksCollectionRequestBuilder is request builder for PlannerTask collection
type PlannerTasksCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerTask collection
func (b *PlannerTasksCollectionRequestBuilder) Request() *PlannerTasksCollectionRequest {
	return &PlannerTasksCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerTask item
func (b *PlannerTasksCollectionRequestBuilder) ID(id string) *PlannerTaskRequestBuilder {
	bb := &PlannerTaskRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerTasksCollectionRequest is request for PlannerTask collection
type PlannerTasksCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PlannerTask collection
func (r *PlannerTasksCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PlannerTask, error) {
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
	var values []PlannerTask
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
			value  []PlannerTask
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

// GetN performs GET request for PlannerTask collection, max N pages
func (r *PlannerTasksCollectionRequest) GetN(ctx context.Context, n int) ([]PlannerTask, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PlannerTask collection
func (r *PlannerTasksCollectionRequest) Get(ctx context.Context) ([]PlannerTask, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PlannerTask collection
func (r *PlannerTasksCollectionRequest) Add(ctx context.Context, reqObj *PlannerTask) (resObj *PlannerTask, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Tasks returns request builder for PlannerTask collection
func (b *PlannerBucketRequestBuilder) Tasks() *PlannerBucketTasksCollectionRequestBuilder {
	bb := &PlannerBucketTasksCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tasks"
	return bb
}

// PlannerBucketTasksCollectionRequestBuilder is request builder for PlannerTask collection
type PlannerBucketTasksCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerTask collection
func (b *PlannerBucketTasksCollectionRequestBuilder) Request() *PlannerBucketTasksCollectionRequest {
	return &PlannerBucketTasksCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerTask item
func (b *PlannerBucketTasksCollectionRequestBuilder) ID(id string) *PlannerTaskRequestBuilder {
	bb := &PlannerTaskRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerBucketTasksCollectionRequest is request for PlannerTask collection
type PlannerBucketTasksCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PlannerTask collection
func (r *PlannerBucketTasksCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PlannerTask, error) {
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
	var values []PlannerTask
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
			value  []PlannerTask
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

// GetN performs GET request for PlannerTask collection, max N pages
func (r *PlannerBucketTasksCollectionRequest) GetN(ctx context.Context, n int) ([]PlannerTask, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PlannerTask collection
func (r *PlannerBucketTasksCollectionRequest) Get(ctx context.Context) ([]PlannerTask, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PlannerTask collection
func (r *PlannerBucketTasksCollectionRequest) Add(ctx context.Context, reqObj *PlannerTask) (resObj *PlannerTask, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Plans returns request builder for PlannerPlan collection
func (b *PlannerGroupRequestBuilder) Plans() *PlannerGroupPlansCollectionRequestBuilder {
	bb := &PlannerGroupPlansCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/plans"
	return bb
}

// PlannerGroupPlansCollectionRequestBuilder is request builder for PlannerPlan collection
type PlannerGroupPlansCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerPlan collection
func (b *PlannerGroupPlansCollectionRequestBuilder) Request() *PlannerGroupPlansCollectionRequest {
	return &PlannerGroupPlansCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerPlan item
func (b *PlannerGroupPlansCollectionRequestBuilder) ID(id string) *PlannerPlanRequestBuilder {
	bb := &PlannerPlanRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerGroupPlansCollectionRequest is request for PlannerPlan collection
type PlannerGroupPlansCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PlannerPlan collection
func (r *PlannerGroupPlansCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PlannerPlan, error) {
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
	var values []PlannerPlan
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
			value  []PlannerPlan
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

// GetN performs GET request for PlannerPlan collection, max N pages
func (r *PlannerGroupPlansCollectionRequest) GetN(ctx context.Context, n int) ([]PlannerPlan, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PlannerPlan collection
func (r *PlannerGroupPlansCollectionRequest) Get(ctx context.Context) ([]PlannerPlan, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PlannerPlan collection
func (r *PlannerGroupPlansCollectionRequest) Add(ctx context.Context, reqObj *PlannerPlan) (resObj *PlannerPlan, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Buckets returns request builder for PlannerBucket collection
func (b *PlannerPlanRequestBuilder) Buckets() *PlannerPlanBucketsCollectionRequestBuilder {
	bb := &PlannerPlanBucketsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/buckets"
	return bb
}

// PlannerPlanBucketsCollectionRequestBuilder is request builder for PlannerBucket collection
type PlannerPlanBucketsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerBucket collection
func (b *PlannerPlanBucketsCollectionRequestBuilder) Request() *PlannerPlanBucketsCollectionRequest {
	return &PlannerPlanBucketsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerBucket item
func (b *PlannerPlanBucketsCollectionRequestBuilder) ID(id string) *PlannerBucketRequestBuilder {
	bb := &PlannerBucketRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerPlanBucketsCollectionRequest is request for PlannerBucket collection
type PlannerPlanBucketsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PlannerBucket collection
func (r *PlannerPlanBucketsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PlannerBucket, error) {
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
	var values []PlannerBucket
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
			value  []PlannerBucket
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

// GetN performs GET request for PlannerBucket collection, max N pages
func (r *PlannerPlanBucketsCollectionRequest) GetN(ctx context.Context, n int) ([]PlannerBucket, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PlannerBucket collection
func (r *PlannerPlanBucketsCollectionRequest) Get(ctx context.Context) ([]PlannerBucket, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PlannerBucket collection
func (r *PlannerPlanBucketsCollectionRequest) Add(ctx context.Context, reqObj *PlannerBucket) (resObj *PlannerBucket, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Details is navigation property
func (b *PlannerPlanRequestBuilder) Details() *PlannerPlanDetailsRequestBuilder {
	bb := &PlannerPlanDetailsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/details"
	return bb
}

// Tasks returns request builder for PlannerTask collection
func (b *PlannerPlanRequestBuilder) Tasks() *PlannerPlanTasksCollectionRequestBuilder {
	bb := &PlannerPlanTasksCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tasks"
	return bb
}

// PlannerPlanTasksCollectionRequestBuilder is request builder for PlannerTask collection
type PlannerPlanTasksCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerTask collection
func (b *PlannerPlanTasksCollectionRequestBuilder) Request() *PlannerPlanTasksCollectionRequest {
	return &PlannerPlanTasksCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerTask item
func (b *PlannerPlanTasksCollectionRequestBuilder) ID(id string) *PlannerTaskRequestBuilder {
	bb := &PlannerTaskRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerPlanTasksCollectionRequest is request for PlannerTask collection
type PlannerPlanTasksCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PlannerTask collection
func (r *PlannerPlanTasksCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PlannerTask, error) {
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
	var values []PlannerTask
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
			value  []PlannerTask
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

// GetN performs GET request for PlannerTask collection, max N pages
func (r *PlannerPlanTasksCollectionRequest) GetN(ctx context.Context, n int) ([]PlannerTask, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PlannerTask collection
func (r *PlannerPlanTasksCollectionRequest) Get(ctx context.Context) ([]PlannerTask, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PlannerTask collection
func (r *PlannerPlanTasksCollectionRequest) Add(ctx context.Context, reqObj *PlannerTask) (resObj *PlannerTask, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AssignedToTaskBoardFormat is navigation property
func (b *PlannerTaskRequestBuilder) AssignedToTaskBoardFormat() *PlannerAssignedToTaskBoardTaskFormatRequestBuilder {
	bb := &PlannerAssignedToTaskBoardTaskFormatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignedToTaskBoardFormat"
	return bb
}

// BucketTaskBoardFormat is navigation property
func (b *PlannerTaskRequestBuilder) BucketTaskBoardFormat() *PlannerBucketTaskBoardTaskFormatRequestBuilder {
	bb := &PlannerBucketTaskBoardTaskFormatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/bucketTaskBoardFormat"
	return bb
}

// Details is navigation property
func (b *PlannerTaskRequestBuilder) Details() *PlannerTaskDetailsRequestBuilder {
	bb := &PlannerTaskDetailsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/details"
	return bb
}

// ProgressTaskBoardFormat is navigation property
func (b *PlannerTaskRequestBuilder) ProgressTaskBoardFormat() *PlannerProgressTaskBoardTaskFormatRequestBuilder {
	bb := &PlannerProgressTaskBoardTaskFormatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/progressTaskBoardFormat"
	return bb
}

// All returns request builder for PlannerDelta collection
func (b *PlannerUserRequestBuilder) All() *PlannerUserAllCollectionRequestBuilder {
	bb := &PlannerUserAllCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/all"
	return bb
}

// PlannerUserAllCollectionRequestBuilder is request builder for PlannerDelta collection
type PlannerUserAllCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerDelta collection
func (b *PlannerUserAllCollectionRequestBuilder) Request() *PlannerUserAllCollectionRequest {
	return &PlannerUserAllCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerDelta item
func (b *PlannerUserAllCollectionRequestBuilder) ID(id string) *PlannerDeltaRequestBuilder {
	bb := &PlannerDeltaRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerUserAllCollectionRequest is request for PlannerDelta collection
type PlannerUserAllCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PlannerDelta collection
func (r *PlannerUserAllCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PlannerDelta, error) {
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
	var values []PlannerDelta
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
			value  []PlannerDelta
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

// GetN performs GET request for PlannerDelta collection, max N pages
func (r *PlannerUserAllCollectionRequest) GetN(ctx context.Context, n int) ([]PlannerDelta, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PlannerDelta collection
func (r *PlannerUserAllCollectionRequest) Get(ctx context.Context) ([]PlannerDelta, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PlannerDelta collection
func (r *PlannerUserAllCollectionRequest) Add(ctx context.Context, reqObj *PlannerDelta) (resObj *PlannerDelta, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// FavoritePlans returns request builder for PlannerPlan collection
func (b *PlannerUserRequestBuilder) FavoritePlans() *PlannerUserFavoritePlansCollectionRequestBuilder {
	bb := &PlannerUserFavoritePlansCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/favoritePlans"
	return bb
}

// PlannerUserFavoritePlansCollectionRequestBuilder is request builder for PlannerPlan collection
type PlannerUserFavoritePlansCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerPlan collection
func (b *PlannerUserFavoritePlansCollectionRequestBuilder) Request() *PlannerUserFavoritePlansCollectionRequest {
	return &PlannerUserFavoritePlansCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerPlan item
func (b *PlannerUserFavoritePlansCollectionRequestBuilder) ID(id string) *PlannerPlanRequestBuilder {
	bb := &PlannerPlanRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerUserFavoritePlansCollectionRequest is request for PlannerPlan collection
type PlannerUserFavoritePlansCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PlannerPlan collection
func (r *PlannerUserFavoritePlansCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PlannerPlan, error) {
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
	var values []PlannerPlan
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
			value  []PlannerPlan
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

// GetN performs GET request for PlannerPlan collection, max N pages
func (r *PlannerUserFavoritePlansCollectionRequest) GetN(ctx context.Context, n int) ([]PlannerPlan, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PlannerPlan collection
func (r *PlannerUserFavoritePlansCollectionRequest) Get(ctx context.Context) ([]PlannerPlan, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PlannerPlan collection
func (r *PlannerUserFavoritePlansCollectionRequest) Add(ctx context.Context, reqObj *PlannerPlan) (resObj *PlannerPlan, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Plans returns request builder for PlannerPlan collection
func (b *PlannerUserRequestBuilder) Plans() *PlannerUserPlansCollectionRequestBuilder {
	bb := &PlannerUserPlansCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/plans"
	return bb
}

// PlannerUserPlansCollectionRequestBuilder is request builder for PlannerPlan collection
type PlannerUserPlansCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerPlan collection
func (b *PlannerUserPlansCollectionRequestBuilder) Request() *PlannerUserPlansCollectionRequest {
	return &PlannerUserPlansCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerPlan item
func (b *PlannerUserPlansCollectionRequestBuilder) ID(id string) *PlannerPlanRequestBuilder {
	bb := &PlannerPlanRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerUserPlansCollectionRequest is request for PlannerPlan collection
type PlannerUserPlansCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PlannerPlan collection
func (r *PlannerUserPlansCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PlannerPlan, error) {
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
	var values []PlannerPlan
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
			value  []PlannerPlan
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

// GetN performs GET request for PlannerPlan collection, max N pages
func (r *PlannerUserPlansCollectionRequest) GetN(ctx context.Context, n int) ([]PlannerPlan, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PlannerPlan collection
func (r *PlannerUserPlansCollectionRequest) Get(ctx context.Context) ([]PlannerPlan, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PlannerPlan collection
func (r *PlannerUserPlansCollectionRequest) Add(ctx context.Context, reqObj *PlannerPlan) (resObj *PlannerPlan, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RecentPlans returns request builder for PlannerPlan collection
func (b *PlannerUserRequestBuilder) RecentPlans() *PlannerUserRecentPlansCollectionRequestBuilder {
	bb := &PlannerUserRecentPlansCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/recentPlans"
	return bb
}

// PlannerUserRecentPlansCollectionRequestBuilder is request builder for PlannerPlan collection
type PlannerUserRecentPlansCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerPlan collection
func (b *PlannerUserRecentPlansCollectionRequestBuilder) Request() *PlannerUserRecentPlansCollectionRequest {
	return &PlannerUserRecentPlansCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerPlan item
func (b *PlannerUserRecentPlansCollectionRequestBuilder) ID(id string) *PlannerPlanRequestBuilder {
	bb := &PlannerPlanRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerUserRecentPlansCollectionRequest is request for PlannerPlan collection
type PlannerUserRecentPlansCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PlannerPlan collection
func (r *PlannerUserRecentPlansCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PlannerPlan, error) {
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
	var values []PlannerPlan
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
			value  []PlannerPlan
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

// GetN performs GET request for PlannerPlan collection, max N pages
func (r *PlannerUserRecentPlansCollectionRequest) GetN(ctx context.Context, n int) ([]PlannerPlan, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PlannerPlan collection
func (r *PlannerUserRecentPlansCollectionRequest) Get(ctx context.Context) ([]PlannerPlan, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PlannerPlan collection
func (r *PlannerUserRecentPlansCollectionRequest) Add(ctx context.Context, reqObj *PlannerPlan) (resObj *PlannerPlan, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Tasks returns request builder for PlannerTask collection
func (b *PlannerUserRequestBuilder) Tasks() *PlannerUserTasksCollectionRequestBuilder {
	bb := &PlannerUserTasksCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tasks"
	return bb
}

// PlannerUserTasksCollectionRequestBuilder is request builder for PlannerTask collection
type PlannerUserTasksCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerTask collection
func (b *PlannerUserTasksCollectionRequestBuilder) Request() *PlannerUserTasksCollectionRequest {
	return &PlannerUserTasksCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerTask item
func (b *PlannerUserTasksCollectionRequestBuilder) ID(id string) *PlannerTaskRequestBuilder {
	bb := &PlannerTaskRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerUserTasksCollectionRequest is request for PlannerTask collection
type PlannerUserTasksCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PlannerTask collection
func (r *PlannerUserTasksCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PlannerTask, error) {
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
	var values []PlannerTask
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
			value  []PlannerTask
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

// GetN performs GET request for PlannerTask collection, max N pages
func (r *PlannerUserTasksCollectionRequest) GetN(ctx context.Context, n int) ([]PlannerTask, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PlannerTask collection
func (r *PlannerUserTasksCollectionRequest) Get(ctx context.Context) ([]PlannerTask, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PlannerTask collection
func (r *PlannerUserTasksCollectionRequest) Add(ctx context.Context, reqObj *PlannerTask) (resObj *PlannerTask, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

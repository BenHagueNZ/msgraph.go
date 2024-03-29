// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// Office365ActiveUserCountsRequestBuilder is request builder for Office365ActiveUserCounts
type Office365ActiveUserCountsRequestBuilder struct{ BaseRequestBuilder }

// Request returns Office365ActiveUserCountsRequest
func (b *Office365ActiveUserCountsRequestBuilder) Request() *Office365ActiveUserCountsRequest {
	return &Office365ActiveUserCountsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Office365ActiveUserCountsRequest is request for Office365ActiveUserCounts
type Office365ActiveUserCountsRequest struct{ BaseRequest }

// Get performs GET request for Office365ActiveUserCounts
func (r *Office365ActiveUserCountsRequest) Get(ctx context.Context) (resObj *Office365ActiveUserCounts, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Office365ActiveUserCounts
func (r *Office365ActiveUserCountsRequest) Update(ctx context.Context, reqObj *Office365ActiveUserCounts) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Office365ActiveUserCounts
func (r *Office365ActiveUserCountsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Office365ActiveUserDetailRequestBuilder is request builder for Office365ActiveUserDetail
type Office365ActiveUserDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns Office365ActiveUserDetailRequest
func (b *Office365ActiveUserDetailRequestBuilder) Request() *Office365ActiveUserDetailRequest {
	return &Office365ActiveUserDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Office365ActiveUserDetailRequest is request for Office365ActiveUserDetail
type Office365ActiveUserDetailRequest struct{ BaseRequest }

// Get performs GET request for Office365ActiveUserDetail
func (r *Office365ActiveUserDetailRequest) Get(ctx context.Context) (resObj *Office365ActiveUserDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Office365ActiveUserDetail
func (r *Office365ActiveUserDetailRequest) Update(ctx context.Context, reqObj *Office365ActiveUserDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Office365ActiveUserDetail
func (r *Office365ActiveUserDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Office365GroupsActivityCountsRequestBuilder is request builder for Office365GroupsActivityCounts
type Office365GroupsActivityCountsRequestBuilder struct{ BaseRequestBuilder }

// Request returns Office365GroupsActivityCountsRequest
func (b *Office365GroupsActivityCountsRequestBuilder) Request() *Office365GroupsActivityCountsRequest {
	return &Office365GroupsActivityCountsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Office365GroupsActivityCountsRequest is request for Office365GroupsActivityCounts
type Office365GroupsActivityCountsRequest struct{ BaseRequest }

// Get performs GET request for Office365GroupsActivityCounts
func (r *Office365GroupsActivityCountsRequest) Get(ctx context.Context) (resObj *Office365GroupsActivityCounts, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Office365GroupsActivityCounts
func (r *Office365GroupsActivityCountsRequest) Update(ctx context.Context, reqObj *Office365GroupsActivityCounts) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Office365GroupsActivityCounts
func (r *Office365GroupsActivityCountsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Office365GroupsActivityDetailRequestBuilder is request builder for Office365GroupsActivityDetail
type Office365GroupsActivityDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns Office365GroupsActivityDetailRequest
func (b *Office365GroupsActivityDetailRequestBuilder) Request() *Office365GroupsActivityDetailRequest {
	return &Office365GroupsActivityDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Office365GroupsActivityDetailRequest is request for Office365GroupsActivityDetail
type Office365GroupsActivityDetailRequest struct{ BaseRequest }

// Get performs GET request for Office365GroupsActivityDetail
func (r *Office365GroupsActivityDetailRequest) Get(ctx context.Context) (resObj *Office365GroupsActivityDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Office365GroupsActivityDetail
func (r *Office365GroupsActivityDetailRequest) Update(ctx context.Context, reqObj *Office365GroupsActivityDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Office365GroupsActivityDetail
func (r *Office365GroupsActivityDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Office365GroupsActivityFileCountsRequestBuilder is request builder for Office365GroupsActivityFileCounts
type Office365GroupsActivityFileCountsRequestBuilder struct{ BaseRequestBuilder }

// Request returns Office365GroupsActivityFileCountsRequest
func (b *Office365GroupsActivityFileCountsRequestBuilder) Request() *Office365GroupsActivityFileCountsRequest {
	return &Office365GroupsActivityFileCountsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Office365GroupsActivityFileCountsRequest is request for Office365GroupsActivityFileCounts
type Office365GroupsActivityFileCountsRequest struct{ BaseRequest }

// Get performs GET request for Office365GroupsActivityFileCounts
func (r *Office365GroupsActivityFileCountsRequest) Get(ctx context.Context) (resObj *Office365GroupsActivityFileCounts, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Office365GroupsActivityFileCounts
func (r *Office365GroupsActivityFileCountsRequest) Update(ctx context.Context, reqObj *Office365GroupsActivityFileCounts) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Office365GroupsActivityFileCounts
func (r *Office365GroupsActivityFileCountsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Office365GroupsActivityGroupCountsRequestBuilder is request builder for Office365GroupsActivityGroupCounts
type Office365GroupsActivityGroupCountsRequestBuilder struct{ BaseRequestBuilder }

// Request returns Office365GroupsActivityGroupCountsRequest
func (b *Office365GroupsActivityGroupCountsRequestBuilder) Request() *Office365GroupsActivityGroupCountsRequest {
	return &Office365GroupsActivityGroupCountsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Office365GroupsActivityGroupCountsRequest is request for Office365GroupsActivityGroupCounts
type Office365GroupsActivityGroupCountsRequest struct{ BaseRequest }

// Get performs GET request for Office365GroupsActivityGroupCounts
func (r *Office365GroupsActivityGroupCountsRequest) Get(ctx context.Context) (resObj *Office365GroupsActivityGroupCounts, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Office365GroupsActivityGroupCounts
func (r *Office365GroupsActivityGroupCountsRequest) Update(ctx context.Context, reqObj *Office365GroupsActivityGroupCounts) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Office365GroupsActivityGroupCounts
func (r *Office365GroupsActivityGroupCountsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Office365GroupsActivityStorageRequestBuilder is request builder for Office365GroupsActivityStorage
type Office365GroupsActivityStorageRequestBuilder struct{ BaseRequestBuilder }

// Request returns Office365GroupsActivityStorageRequest
func (b *Office365GroupsActivityStorageRequestBuilder) Request() *Office365GroupsActivityStorageRequest {
	return &Office365GroupsActivityStorageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Office365GroupsActivityStorageRequest is request for Office365GroupsActivityStorage
type Office365GroupsActivityStorageRequest struct{ BaseRequest }

// Get performs GET request for Office365GroupsActivityStorage
func (r *Office365GroupsActivityStorageRequest) Get(ctx context.Context) (resObj *Office365GroupsActivityStorage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Office365GroupsActivityStorage
func (r *Office365GroupsActivityStorageRequest) Update(ctx context.Context, reqObj *Office365GroupsActivityStorage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Office365GroupsActivityStorage
func (r *Office365GroupsActivityStorageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Office365ServicesUserCountsRequestBuilder is request builder for Office365ServicesUserCounts
type Office365ServicesUserCountsRequestBuilder struct{ BaseRequestBuilder }

// Request returns Office365ServicesUserCountsRequest
func (b *Office365ServicesUserCountsRequestBuilder) Request() *Office365ServicesUserCountsRequest {
	return &Office365ServicesUserCountsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Office365ServicesUserCountsRequest is request for Office365ServicesUserCounts
type Office365ServicesUserCountsRequest struct{ BaseRequest }

// Get performs GET request for Office365ServicesUserCounts
func (r *Office365ServicesUserCountsRequest) Get(ctx context.Context) (resObj *Office365ServicesUserCounts, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Office365ServicesUserCounts
func (r *Office365ServicesUserCountsRequest) Update(ctx context.Context, reqObj *Office365ServicesUserCounts) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Office365ServicesUserCounts
func (r *Office365ServicesUserCountsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

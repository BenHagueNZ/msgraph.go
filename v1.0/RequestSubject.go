// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SubjectRightsRequestDetailRequestBuilder is request builder for SubjectRightsRequestDetail
type SubjectRightsRequestDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns SubjectRightsRequestDetailRequest
func (b *SubjectRightsRequestDetailRequestBuilder) Request() *SubjectRightsRequestDetailRequest {
	return &SubjectRightsRequestDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SubjectRightsRequestDetailRequest is request for SubjectRightsRequestDetail
type SubjectRightsRequestDetailRequest struct{ BaseRequest }

// Get performs GET request for SubjectRightsRequestDetail
func (r *SubjectRightsRequestDetailRequest) Get(ctx context.Context) (resObj *SubjectRightsRequestDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SubjectRightsRequestDetail
func (r *SubjectRightsRequestDetailRequest) Update(ctx context.Context, reqObj *SubjectRightsRequestDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SubjectRightsRequestDetail
func (r *SubjectRightsRequestDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SubjectRightsRequestHistoryRequestBuilder is request builder for SubjectRightsRequestHistory
type SubjectRightsRequestHistoryRequestBuilder struct{ BaseRequestBuilder }

// Request returns SubjectRightsRequestHistoryRequest
func (b *SubjectRightsRequestHistoryRequestBuilder) Request() *SubjectRightsRequestHistoryRequest {
	return &SubjectRightsRequestHistoryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SubjectRightsRequestHistoryRequest is request for SubjectRightsRequestHistory
type SubjectRightsRequestHistoryRequest struct{ BaseRequest }

// Get performs GET request for SubjectRightsRequestHistory
func (r *SubjectRightsRequestHistoryRequest) Get(ctx context.Context) (resObj *SubjectRightsRequestHistory, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SubjectRightsRequestHistory
func (r *SubjectRightsRequestHistoryRequest) Update(ctx context.Context, reqObj *SubjectRightsRequestHistory) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SubjectRightsRequestHistory
func (r *SubjectRightsRequestHistoryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SubjectRightsRequestObjectRequestBuilder is request builder for SubjectRightsRequestObject
type SubjectRightsRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns SubjectRightsRequestObjectRequest
func (b *SubjectRightsRequestObjectRequestBuilder) Request() *SubjectRightsRequestObjectRequest {
	return &SubjectRightsRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SubjectRightsRequestObjectRequest is request for SubjectRightsRequestObject
type SubjectRightsRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for SubjectRightsRequestObject
func (r *SubjectRightsRequestObjectRequest) Get(ctx context.Context) (resObj *SubjectRightsRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SubjectRightsRequestObject
func (r *SubjectRightsRequestObjectRequest) Update(ctx context.Context, reqObj *SubjectRightsRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SubjectRightsRequestObject
func (r *SubjectRightsRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SubjectRightsRequestStageDetailRequestBuilder is request builder for SubjectRightsRequestStageDetail
type SubjectRightsRequestStageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns SubjectRightsRequestStageDetailRequest
func (b *SubjectRightsRequestStageDetailRequestBuilder) Request() *SubjectRightsRequestStageDetailRequest {
	return &SubjectRightsRequestStageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SubjectRightsRequestStageDetailRequest is request for SubjectRightsRequestStageDetail
type SubjectRightsRequestStageDetailRequest struct{ BaseRequest }

// Get performs GET request for SubjectRightsRequestStageDetail
func (r *SubjectRightsRequestStageDetailRequest) Get(ctx context.Context) (resObj *SubjectRightsRequestStageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SubjectRightsRequestStageDetail
func (r *SubjectRightsRequestStageDetailRequest) Update(ctx context.Context, reqObj *SubjectRightsRequestStageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SubjectRightsRequestStageDetail
func (r *SubjectRightsRequestStageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SubjectSetRequestBuilder is request builder for SubjectSet
type SubjectSetRequestBuilder struct{ BaseRequestBuilder }

// Request returns SubjectSetRequest
func (b *SubjectSetRequestBuilder) Request() *SubjectSetRequest {
	return &SubjectSetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SubjectSetRequest is request for SubjectSet
type SubjectSetRequest struct{ BaseRequest }

// Get performs GET request for SubjectSet
func (r *SubjectSetRequest) Get(ctx context.Context) (resObj *SubjectSet, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SubjectSet
func (r *SubjectSetRequest) Update(ctx context.Context, reqObj *SubjectSet) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SubjectSet
func (r *SubjectSetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

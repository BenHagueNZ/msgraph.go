// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// OnlineMeetingRequestBuilder is request builder for OnlineMeeting
type OnlineMeetingRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnlineMeetingRequest
func (b *OnlineMeetingRequestBuilder) Request() *OnlineMeetingRequest {
	return &OnlineMeetingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnlineMeetingRequest is request for OnlineMeeting
type OnlineMeetingRequest struct{ BaseRequest }

// Get performs GET request for OnlineMeeting
func (r *OnlineMeetingRequest) Get(ctx context.Context) (resObj *OnlineMeeting, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnlineMeeting
func (r *OnlineMeetingRequest) Update(ctx context.Context, reqObj *OnlineMeeting) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnlineMeeting
func (r *OnlineMeetingRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OnlineMeetingInfoRequestBuilder is request builder for OnlineMeetingInfo
type OnlineMeetingInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnlineMeetingInfoRequest
func (b *OnlineMeetingInfoRequestBuilder) Request() *OnlineMeetingInfoRequest {
	return &OnlineMeetingInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnlineMeetingInfoRequest is request for OnlineMeetingInfo
type OnlineMeetingInfoRequest struct{ BaseRequest }

// Get performs GET request for OnlineMeetingInfo
func (r *OnlineMeetingInfoRequest) Get(ctx context.Context) (resObj *OnlineMeetingInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnlineMeetingInfo
func (r *OnlineMeetingInfoRequest) Update(ctx context.Context, reqObj *OnlineMeetingInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnlineMeetingInfo
func (r *OnlineMeetingInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OnlineMeetingRestrictedRequestBuilder is request builder for OnlineMeetingRestricted
type OnlineMeetingRestrictedRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnlineMeetingRestrictedRequest
func (b *OnlineMeetingRestrictedRequestBuilder) Request() *OnlineMeetingRestrictedRequest {
	return &OnlineMeetingRestrictedRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnlineMeetingRestrictedRequest is request for OnlineMeetingRestricted
type OnlineMeetingRestrictedRequest struct{ BaseRequest }

// Get performs GET request for OnlineMeetingRestricted
func (r *OnlineMeetingRestrictedRequest) Get(ctx context.Context) (resObj *OnlineMeetingRestricted, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnlineMeetingRestricted
func (r *OnlineMeetingRestrictedRequest) Update(ctx context.Context, reqObj *OnlineMeetingRestricted) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnlineMeetingRestricted
func (r *OnlineMeetingRestrictedRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type OnlineMeetingCollectionCreateOrGetRequestBuilder struct{ BaseRequestBuilder }

// CreateOrGet action undocumentedras
func (b *CloudCommunicationsOnlineMeetingsCollectionRequestBuilder) CreateOrGet(reqObj *OnlineMeetingCollectionCreateOrGetRequestParameter) *OnlineMeetingCollectionCreateOrGetRequestBuilder {
	bb := &OnlineMeetingCollectionCreateOrGetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CreateOrGet"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// CreateOrGet action undocumentedras
func (b *CommsApplicationOnlineMeetingsCollectionRequestBuilder) CreateOrGet(reqObj *OnlineMeetingCollectionCreateOrGetRequestParameter) *OnlineMeetingCollectionCreateOrGetRequestBuilder {
	bb := &OnlineMeetingCollectionCreateOrGetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CreateOrGet"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// CreateOrGet action undocumentedras
func (b *UserOnlineMeetingsCollectionRequestBuilder) CreateOrGet(reqObj *OnlineMeetingCollectionCreateOrGetRequestParameter) *OnlineMeetingCollectionCreateOrGetRequestBuilder {
	bb := &OnlineMeetingCollectionCreateOrGetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CreateOrGet"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type OnlineMeetingCollectionCreateOrGetRequest struct{ BaseRequest }

func (b *OnlineMeetingCollectionCreateOrGetRequestBuilder) Request() *OnlineMeetingCollectionCreateOrGetRequest {
	return &OnlineMeetingCollectionCreateOrGetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *OnlineMeetingCollectionCreateOrGetRequest) Post(ctx context.Context) (resObj *OnlineMeeting, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ImpactedResourceRequestBuilder is request builder for ImpactedResource
type ImpactedResourceRequestBuilder struct{ BaseRequestBuilder }

// Request returns ImpactedResourceRequest
func (b *ImpactedResourceRequestBuilder) Request() *ImpactedResourceRequest {
	return &ImpactedResourceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ImpactedResourceRequest is request for ImpactedResource
type ImpactedResourceRequest struct{ BaseRequest }

// Get performs GET request for ImpactedResource
func (r *ImpactedResourceRequest) Get(ctx context.Context) (resObj *ImpactedResource, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ImpactedResource
func (r *ImpactedResourceRequest) Update(ctx context.Context, reqObj *ImpactedResource) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ImpactedResource
func (r *ImpactedResourceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type ImpactedResourceCompleteRequestBuilder struct{ BaseRequestBuilder }

// Complete action undocumentedras
func (b *ImpactedResourceRequestBuilder) Complete(reqObj *ImpactedResourceCompleteRequestParameter) *ImpactedResourceCompleteRequestBuilder {
	bb := &ImpactedResourceCompleteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Complete"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ImpactedResourceCompleteRequest struct{ BaseRequest }

func (b *ImpactedResourceCompleteRequestBuilder) Request() *ImpactedResourceCompleteRequest {
	return &ImpactedResourceCompleteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ImpactedResourceCompleteRequest) Post(ctx context.Context) (resObj *ImpactedResource, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type ImpactedResourceDismissRequestBuilder struct{ BaseRequestBuilder }

// Dismiss action undocumentedras
func (b *ImpactedResourceRequestBuilder) Dismiss(reqObj *ImpactedResourceDismissRequestParameter) *ImpactedResourceDismissRequestBuilder {
	bb := &ImpactedResourceDismissRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Dismiss"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ImpactedResourceDismissRequest struct{ BaseRequest }

func (b *ImpactedResourceDismissRequestBuilder) Request() *ImpactedResourceDismissRequest {
	return &ImpactedResourceDismissRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ImpactedResourceDismissRequest) Post(ctx context.Context) (resObj *ImpactedResource, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type ImpactedResourcePostponeRequestBuilder struct{ BaseRequestBuilder }

// Postpone action undocumentedras
func (b *ImpactedResourceRequestBuilder) Postpone(reqObj *ImpactedResourcePostponeRequestParameter) *ImpactedResourcePostponeRequestBuilder {
	bb := &ImpactedResourcePostponeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Postpone"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ImpactedResourcePostponeRequest struct{ BaseRequest }

func (b *ImpactedResourcePostponeRequestBuilder) Request() *ImpactedResourcePostponeRequest {
	return &ImpactedResourcePostponeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ImpactedResourcePostponeRequest) Post(ctx context.Context) (resObj *ImpactedResource, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type ImpactedResourceReactivateRequestBuilder struct{ BaseRequestBuilder }

// Reactivate action undocumentedras
func (b *ImpactedResourceRequestBuilder) Reactivate(reqObj *ImpactedResourceReactivateRequestParameter) *ImpactedResourceReactivateRequestBuilder {
	bb := &ImpactedResourceReactivateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Reactivate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ImpactedResourceReactivateRequest struct{ BaseRequest }

func (b *ImpactedResourceReactivateRequestBuilder) Request() *ImpactedResourceReactivateRequest {
	return &ImpactedResourceReactivateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ImpactedResourceReactivateRequest) Post(ctx context.Context) (resObj *ImpactedResource, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

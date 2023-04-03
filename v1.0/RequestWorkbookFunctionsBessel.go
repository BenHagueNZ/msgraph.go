// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsBesselIRequestBuilder struct{ BaseRequestBuilder }

// BesselI action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) BesselI(reqObj *WorkbookFunctionsBesselIRequestParameter) *WorkbookFunctionsBesselIRequestBuilder {
	bb := &WorkbookFunctionsBesselIRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/BesselI"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsBesselIRequest struct{ BaseRequest }

func (b *WorkbookFunctionsBesselIRequestBuilder) Request() *WorkbookFunctionsBesselIRequest {
	return &WorkbookFunctionsBesselIRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsBesselIRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsBesselJRequestBuilder struct{ BaseRequestBuilder }

// BesselJ action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) BesselJ(reqObj *WorkbookFunctionsBesselJRequestParameter) *WorkbookFunctionsBesselJRequestBuilder {
	bb := &WorkbookFunctionsBesselJRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/BesselJ"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsBesselJRequest struct{ BaseRequest }

func (b *WorkbookFunctionsBesselJRequestBuilder) Request() *WorkbookFunctionsBesselJRequest {
	return &WorkbookFunctionsBesselJRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsBesselJRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsBesselKRequestBuilder struct{ BaseRequestBuilder }

// BesselK action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) BesselK(reqObj *WorkbookFunctionsBesselKRequestParameter) *WorkbookFunctionsBesselKRequestBuilder {
	bb := &WorkbookFunctionsBesselKRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/BesselK"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsBesselKRequest struct{ BaseRequest }

func (b *WorkbookFunctionsBesselKRequestBuilder) Request() *WorkbookFunctionsBesselKRequest {
	return &WorkbookFunctionsBesselKRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsBesselKRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsBesselYRequestBuilder struct{ BaseRequestBuilder }

// BesselY action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) BesselY(reqObj *WorkbookFunctionsBesselYRequestParameter) *WorkbookFunctionsBesselYRequestBuilder {
	bb := &WorkbookFunctionsBesselYRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/BesselY"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsBesselYRequest struct{ BaseRequest }

func (b *WorkbookFunctionsBesselYRequestBuilder) Request() *WorkbookFunctionsBesselYRequest {
	return &WorkbookFunctionsBesselYRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsBesselYRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsT_DistRequestBuilder struct{ BaseRequestBuilder }

// T_Dist action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) T_Dist(reqObj *WorkbookFunctionsT_DistRequestParameter) *WorkbookFunctionsT_DistRequestBuilder {
	bb := &WorkbookFunctionsT_DistRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/T_Dist"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsT_DistRequest struct{ BaseRequest }

func (b *WorkbookFunctionsT_DistRequestBuilder) Request() *WorkbookFunctionsT_DistRequest {
	return &WorkbookFunctionsT_DistRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsT_DistRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsT_Dist_2TRequestBuilder struct{ BaseRequestBuilder }

// T_Dist_2T action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) T_Dist_2T(reqObj *WorkbookFunctionsT_Dist_2TRequestParameter) *WorkbookFunctionsT_Dist_2TRequestBuilder {
	bb := &WorkbookFunctionsT_Dist_2TRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/T_Dist_2T"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsT_Dist_2TRequest struct{ BaseRequest }

func (b *WorkbookFunctionsT_Dist_2TRequestBuilder) Request() *WorkbookFunctionsT_Dist_2TRequest {
	return &WorkbookFunctionsT_Dist_2TRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsT_Dist_2TRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsT_Dist_RTRequestBuilder struct{ BaseRequestBuilder }

// T_Dist_RT action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) T_Dist_RT(reqObj *WorkbookFunctionsT_Dist_RTRequestParameter) *WorkbookFunctionsT_Dist_RTRequestBuilder {
	bb := &WorkbookFunctionsT_Dist_RTRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/T_Dist_RT"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsT_Dist_RTRequest struct{ BaseRequest }

func (b *WorkbookFunctionsT_Dist_RTRequestBuilder) Request() *WorkbookFunctionsT_Dist_RTRequest {
	return &WorkbookFunctionsT_Dist_RTRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsT_Dist_RTRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsT_InvRequestBuilder struct{ BaseRequestBuilder }

// T_Inv action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) T_Inv(reqObj *WorkbookFunctionsT_InvRequestParameter) *WorkbookFunctionsT_InvRequestBuilder {
	bb := &WorkbookFunctionsT_InvRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/T_Inv"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsT_InvRequest struct{ BaseRequest }

func (b *WorkbookFunctionsT_InvRequestBuilder) Request() *WorkbookFunctionsT_InvRequest {
	return &WorkbookFunctionsT_InvRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsT_InvRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsT_Inv_2TRequestBuilder struct{ BaseRequestBuilder }

// T_Inv_2T action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) T_Inv_2T(reqObj *WorkbookFunctionsT_Inv_2TRequestParameter) *WorkbookFunctionsT_Inv_2TRequestBuilder {
	bb := &WorkbookFunctionsT_Inv_2TRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/T_Inv_2T"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsT_Inv_2TRequest struct{ BaseRequest }

func (b *WorkbookFunctionsT_Inv_2TRequestBuilder) Request() *WorkbookFunctionsT_Inv_2TRequest {
	return &WorkbookFunctionsT_Inv_2TRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsT_Inv_2TRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

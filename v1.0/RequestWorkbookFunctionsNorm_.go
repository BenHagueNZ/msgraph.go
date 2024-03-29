// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsNorm_DistRequestBuilder struct{ BaseRequestBuilder }

// Norm_Dist action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Norm_Dist(reqObj *WorkbookFunctionsNorm_DistRequestParameter) *WorkbookFunctionsNorm_DistRequestBuilder {
	bb := &WorkbookFunctionsNorm_DistRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Norm_Dist"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsNorm_DistRequest struct{ BaseRequest }

func (b *WorkbookFunctionsNorm_DistRequestBuilder) Request() *WorkbookFunctionsNorm_DistRequest {
	return &WorkbookFunctionsNorm_DistRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsNorm_DistRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsNorm_InvRequestBuilder struct{ BaseRequestBuilder }

// Norm_Inv action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Norm_Inv(reqObj *WorkbookFunctionsNorm_InvRequestParameter) *WorkbookFunctionsNorm_InvRequestBuilder {
	bb := &WorkbookFunctionsNorm_InvRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Norm_Inv"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsNorm_InvRequest struct{ BaseRequest }

func (b *WorkbookFunctionsNorm_InvRequestBuilder) Request() *WorkbookFunctionsNorm_InvRequest {
	return &WorkbookFunctionsNorm_InvRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsNorm_InvRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsNorm_S_DistRequestBuilder struct{ BaseRequestBuilder }

// Norm_S_Dist action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Norm_S_Dist(reqObj *WorkbookFunctionsNorm_S_DistRequestParameter) *WorkbookFunctionsNorm_S_DistRequestBuilder {
	bb := &WorkbookFunctionsNorm_S_DistRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Norm_S_Dist"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsNorm_S_DistRequest struct{ BaseRequest }

func (b *WorkbookFunctionsNorm_S_DistRequestBuilder) Request() *WorkbookFunctionsNorm_S_DistRequest {
	return &WorkbookFunctionsNorm_S_DistRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsNorm_S_DistRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsNorm_S_InvRequestBuilder struct{ BaseRequestBuilder }

// Norm_S_Inv action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Norm_S_Inv(reqObj *WorkbookFunctionsNorm_S_InvRequestParameter) *WorkbookFunctionsNorm_S_InvRequestBuilder {
	bb := &WorkbookFunctionsNorm_S_InvRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Norm_S_Inv"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsNorm_S_InvRequest struct{ BaseRequest }

func (b *WorkbookFunctionsNorm_S_InvRequestBuilder) Request() *WorkbookFunctionsNorm_S_InvRequest {
	return &WorkbookFunctionsNorm_S_InvRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsNorm_S_InvRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

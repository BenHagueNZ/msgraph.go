// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsChiSq_DistRequestBuilder struct{ BaseRequestBuilder }

// ChiSq_Dist action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) ChiSq_Dist(reqObj *WorkbookFunctionsChiSq_DistRequestParameter) *WorkbookFunctionsChiSq_DistRequestBuilder {
	bb := &WorkbookFunctionsChiSq_DistRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ChiSq_Dist"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsChiSq_DistRequest struct{ BaseRequest }

func (b *WorkbookFunctionsChiSq_DistRequestBuilder) Request() *WorkbookFunctionsChiSq_DistRequest {
	return &WorkbookFunctionsChiSq_DistRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsChiSq_DistRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsChiSq_Dist_RTRequestBuilder struct{ BaseRequestBuilder }

// ChiSq_Dist_RT action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) ChiSq_Dist_RT(reqObj *WorkbookFunctionsChiSq_Dist_RTRequestParameter) *WorkbookFunctionsChiSq_Dist_RTRequestBuilder {
	bb := &WorkbookFunctionsChiSq_Dist_RTRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ChiSq_Dist_RT"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsChiSq_Dist_RTRequest struct{ BaseRequest }

func (b *WorkbookFunctionsChiSq_Dist_RTRequestBuilder) Request() *WorkbookFunctionsChiSq_Dist_RTRequest {
	return &WorkbookFunctionsChiSq_Dist_RTRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsChiSq_Dist_RTRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsChiSq_InvRequestBuilder struct{ BaseRequestBuilder }

// ChiSq_Inv action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) ChiSq_Inv(reqObj *WorkbookFunctionsChiSq_InvRequestParameter) *WorkbookFunctionsChiSq_InvRequestBuilder {
	bb := &WorkbookFunctionsChiSq_InvRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ChiSq_Inv"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsChiSq_InvRequest struct{ BaseRequest }

func (b *WorkbookFunctionsChiSq_InvRequestBuilder) Request() *WorkbookFunctionsChiSq_InvRequest {
	return &WorkbookFunctionsChiSq_InvRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsChiSq_InvRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsChiSq_Inv_RTRequestBuilder struct{ BaseRequestBuilder }

// ChiSq_Inv_RT action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) ChiSq_Inv_RT(reqObj *WorkbookFunctionsChiSq_Inv_RTRequestParameter) *WorkbookFunctionsChiSq_Inv_RTRequestBuilder {
	bb := &WorkbookFunctionsChiSq_Inv_RTRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ChiSq_Inv_RT"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsChiSq_Inv_RTRequest struct{ BaseRequest }

func (b *WorkbookFunctionsChiSq_Inv_RTRequestBuilder) Request() *WorkbookFunctionsChiSq_Inv_RTRequest {
	return &WorkbookFunctionsChiSq_Inv_RTRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsChiSq_Inv_RTRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

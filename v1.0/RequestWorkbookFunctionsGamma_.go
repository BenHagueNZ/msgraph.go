// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsGamma_DistRequestBuilder struct{ BaseRequestBuilder }

// Gamma_Dist action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Gamma_Dist(reqObj *WorkbookFunctionsGamma_DistRequestParameter) *WorkbookFunctionsGamma_DistRequestBuilder {
	bb := &WorkbookFunctionsGamma_DistRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Gamma_Dist"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsGamma_DistRequest struct{ BaseRequest }

func (b *WorkbookFunctionsGamma_DistRequestBuilder) Request() *WorkbookFunctionsGamma_DistRequest {
	return &WorkbookFunctionsGamma_DistRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsGamma_DistRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsGamma_InvRequestBuilder struct{ BaseRequestBuilder }

// Gamma_Inv action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Gamma_Inv(reqObj *WorkbookFunctionsGamma_InvRequestParameter) *WorkbookFunctionsGamma_InvRequestBuilder {
	bb := &WorkbookFunctionsGamma_InvRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Gamma_Inv"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsGamma_InvRequest struct{ BaseRequest }

func (b *WorkbookFunctionsGamma_InvRequestBuilder) Request() *WorkbookFunctionsGamma_InvRequest {
	return &WorkbookFunctionsGamma_InvRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsGamma_InvRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

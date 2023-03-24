// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsRoundRequestBuilder struct{ BaseRequestBuilder }

// Round action undocumented
func (b *WorkbookFunctionsRequestBuilder) Round(reqObj *WorkbookFunctionsRoundRequestParameter) *WorkbookFunctionsRoundRequestBuilder {
	bb := &WorkbookFunctionsRoundRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Round"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsRoundRequest struct{ BaseRequest }

func (b *WorkbookFunctionsRoundRequestBuilder) Request() *WorkbookFunctionsRoundRequest {
	return &WorkbookFunctionsRoundRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsRoundRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsRoundDownRequestBuilder struct{ BaseRequestBuilder }

// RoundDown action undocumented
func (b *WorkbookFunctionsRequestBuilder) RoundDown(reqObj *WorkbookFunctionsRoundDownRequestParameter) *WorkbookFunctionsRoundDownRequestBuilder {
	bb := &WorkbookFunctionsRoundDownRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RoundDown"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsRoundDownRequest struct{ BaseRequest }

func (b *WorkbookFunctionsRoundDownRequestBuilder) Request() *WorkbookFunctionsRoundDownRequest {
	return &WorkbookFunctionsRoundDownRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsRoundDownRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsRoundUpRequestBuilder struct{ BaseRequestBuilder }

// RoundUp action undocumented
func (b *WorkbookFunctionsRequestBuilder) RoundUp(reqObj *WorkbookFunctionsRoundUpRequestParameter) *WorkbookFunctionsRoundUpRequestBuilder {
	bb := &WorkbookFunctionsRoundUpRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RoundUp"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsRoundUpRequest struct{ BaseRequest }

func (b *WorkbookFunctionsRoundUpRequestBuilder) Request() *WorkbookFunctionsRoundUpRequest {
	return &WorkbookFunctionsRoundUpRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsRoundUpRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

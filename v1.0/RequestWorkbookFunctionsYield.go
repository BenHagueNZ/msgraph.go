// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsYieldRequestBuilder struct{ BaseRequestBuilder }

// Yield action undocumented
func (b *WorkbookFunctionsRequestBuilder) Yield(reqObj *WorkbookFunctionsYieldRequestParameter) *WorkbookFunctionsYieldRequestBuilder {
	bb := &WorkbookFunctionsYieldRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/yield"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsYieldRequest struct{ BaseRequest }

func (b *WorkbookFunctionsYieldRequestBuilder) Request() *WorkbookFunctionsYieldRequest {
	return &WorkbookFunctionsYieldRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsYieldRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsYieldDiscRequestBuilder struct{ BaseRequestBuilder }

// YieldDisc action undocumented
func (b *WorkbookFunctionsRequestBuilder) YieldDisc(reqObj *WorkbookFunctionsYieldDiscRequestParameter) *WorkbookFunctionsYieldDiscRequestBuilder {
	bb := &WorkbookFunctionsYieldDiscRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/yieldDisc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsYieldDiscRequest struct{ BaseRequest }

func (b *WorkbookFunctionsYieldDiscRequestBuilder) Request() *WorkbookFunctionsYieldDiscRequest {
	return &WorkbookFunctionsYieldDiscRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsYieldDiscRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsYieldMatRequestBuilder struct{ BaseRequestBuilder }

// YieldMat action undocumented
func (b *WorkbookFunctionsRequestBuilder) YieldMat(reqObj *WorkbookFunctionsYieldMatRequestParameter) *WorkbookFunctionsYieldMatRequestBuilder {
	bb := &WorkbookFunctionsYieldMatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/yieldMat"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsYieldMatRequest struct{ BaseRequest }

func (b *WorkbookFunctionsYieldMatRequestBuilder) Request() *WorkbookFunctionsYieldMatRequest {
	return &WorkbookFunctionsYieldMatRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsYieldMatRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

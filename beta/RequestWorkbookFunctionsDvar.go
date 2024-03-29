// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsDvarRequestBuilder struct{ BaseRequestBuilder }

// Dvar action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Dvar(reqObj *WorkbookFunctionsDvarRequestParameter) *WorkbookFunctionsDvarRequestBuilder {
	bb := &WorkbookFunctionsDvarRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Dvar"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsDvarRequest struct{ BaseRequest }

func (b *WorkbookFunctionsDvarRequestBuilder) Request() *WorkbookFunctionsDvarRequest {
	return &WorkbookFunctionsDvarRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsDvarRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsDvarPRequestBuilder struct{ BaseRequestBuilder }

// DvarP action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) DvarP(reqObj *WorkbookFunctionsDvarPRequestParameter) *WorkbookFunctionsDvarPRequestBuilder {
	bb := &WorkbookFunctionsDvarPRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/DvarP"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsDvarPRequest struct{ BaseRequest }

func (b *WorkbookFunctionsDvarPRequestBuilder) Request() *WorkbookFunctionsDvarPRequest {
	return &WorkbookFunctionsDvarPRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsDvarPRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

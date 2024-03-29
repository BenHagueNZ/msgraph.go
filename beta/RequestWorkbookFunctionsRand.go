// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsRandRequestBuilder struct{ BaseRequestBuilder }

// Rand action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Rand(reqObj *WorkbookFunctionsRandRequestParameter) *WorkbookFunctionsRandRequestBuilder {
	bb := &WorkbookFunctionsRandRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Rand"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsRandRequest struct{ BaseRequest }

func (b *WorkbookFunctionsRandRequestBuilder) Request() *WorkbookFunctionsRandRequest {
	return &WorkbookFunctionsRandRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsRandRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsRandBetweenRequestBuilder struct{ BaseRequestBuilder }

// RandBetween action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) RandBetween(reqObj *WorkbookFunctionsRandBetweenRequestParameter) *WorkbookFunctionsRandBetweenRequestBuilder {
	bb := &WorkbookFunctionsRandBetweenRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RandBetween"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsRandBetweenRequest struct{ BaseRequest }

func (b *WorkbookFunctionsRandBetweenRequestBuilder) Request() *WorkbookFunctionsRandBetweenRequest {
	return &WorkbookFunctionsRandBetweenRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsRandBetweenRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

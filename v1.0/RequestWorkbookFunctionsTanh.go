// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsTanhRequestBuilder struct{ BaseRequestBuilder }

// Tanh action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Tanh(reqObj *WorkbookFunctionsTanhRequestParameter) *WorkbookFunctionsTanhRequestBuilder {
	bb := &WorkbookFunctionsTanhRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Tanh"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsTanhRequest struct{ BaseRequest }

func (b *WorkbookFunctionsTanhRequestBuilder) Request() *WorkbookFunctionsTanhRequest {
	return &WorkbookFunctionsTanhRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsTanhRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

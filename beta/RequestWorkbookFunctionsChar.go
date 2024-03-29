// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsCharRequestBuilder struct{ BaseRequestBuilder }

// Char action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Char(reqObj *WorkbookFunctionsCharRequestParameter) *WorkbookFunctionsCharRequestBuilder {
	bb := &WorkbookFunctionsCharRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Char"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsCharRequest struct{ BaseRequest }

func (b *WorkbookFunctionsCharRequestBuilder) Request() *WorkbookFunctionsCharRequest {
	return &WorkbookFunctionsCharRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsCharRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

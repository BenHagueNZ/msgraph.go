// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsRightbRequestBuilder struct{ BaseRequestBuilder }

// Rightb action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Rightb(reqObj *WorkbookFunctionsRightbRequestParameter) *WorkbookFunctionsRightbRequestBuilder {
	bb := &WorkbookFunctionsRightbRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Rightb"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsRightbRequest struct{ BaseRequest }

func (b *WorkbookFunctionsRightbRequestBuilder) Request() *WorkbookFunctionsRightbRequest {
	return &WorkbookFunctionsRightbRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsRightbRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
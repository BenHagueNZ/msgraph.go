// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsReptRequestBuilder struct{ BaseRequestBuilder }

// Rept action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Rept(reqObj *WorkbookFunctionsReptRequestParameter) *WorkbookFunctionsReptRequestBuilder {
	bb := &WorkbookFunctionsReptRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Rept"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsReptRequest struct{ BaseRequest }

func (b *WorkbookFunctionsReptRequestBuilder) Request() *WorkbookFunctionsReptRequest {
	return &WorkbookFunctionsReptRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsReptRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

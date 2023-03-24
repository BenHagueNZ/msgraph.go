// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsSecRequestBuilder struct{ BaseRequestBuilder }

// Sec action undocumented
func (b *WorkbookFunctionsRequestBuilder) Sec(reqObj *WorkbookFunctionsSecRequestParameter) *WorkbookFunctionsSecRequestBuilder {
	bb := &WorkbookFunctionsSecRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Sec"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsSecRequest struct{ BaseRequest }

func (b *WorkbookFunctionsSecRequestBuilder) Request() *WorkbookFunctionsSecRequest {
	return &WorkbookFunctionsSecRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsSecRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

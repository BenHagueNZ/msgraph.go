// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsSydRequestBuilder struct{ BaseRequestBuilder }

// Syd action undocumented
func (b *WorkbookFunctionsRequestBuilder) Syd(reqObj *WorkbookFunctionsSydRequestParameter) *WorkbookFunctionsSydRequestBuilder {
	bb := &WorkbookFunctionsSydRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Syd"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsSydRequest struct{ BaseRequest }

func (b *WorkbookFunctionsSydRequestBuilder) Request() *WorkbookFunctionsSydRequest {
	return &WorkbookFunctionsSydRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsSydRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsRriRequestBuilder struct{ BaseRequestBuilder }

// Rri action undocumented
func (b *WorkbookFunctionsRequestBuilder) Rri(reqObj *WorkbookFunctionsRriRequestParameter) *WorkbookFunctionsRriRequestBuilder {
	bb := &WorkbookFunctionsRriRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/rri"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsRriRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsRriRequestBuilder) Request() *WorkbookFunctionsRriRequest {
	return &WorkbookFunctionsRriRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsRriRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

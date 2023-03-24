// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsIsrefRequestBuilder struct{ BaseRequestBuilder }

// Isref action undocumented
func (b *WorkbookFunctionsRequestBuilder) Isref(reqObj *WorkbookFunctionsIsrefRequestParameter) *WorkbookFunctionsIsrefRequestBuilder {
	bb := &WorkbookFunctionsIsrefRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Isref"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsIsrefRequest struct{ BaseRequest }

func (b *WorkbookFunctionsIsrefRequestBuilder) Request() *WorkbookFunctionsIsrefRequest {
	return &WorkbookFunctionsIsrefRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsIsrefRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

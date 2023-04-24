// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsSmallRequestBuilder struct{ BaseRequestBuilder }

// Small action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Small(reqObj *WorkbookFunctionsSmallRequestParameter) *WorkbookFunctionsSmallRequestBuilder {
	bb := &WorkbookFunctionsSmallRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Small"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsSmallRequest struct{ BaseRequest }

func (b *WorkbookFunctionsSmallRequestBuilder) Request() *WorkbookFunctionsSmallRequest {
	return &WorkbookFunctionsSmallRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsSmallRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsAbsRequestBuilder struct{ BaseRequestBuilder }

// Abs action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Abs(reqObj *WorkbookFunctionsAbsRequestParameter) *WorkbookFunctionsAbsRequestBuilder {
	bb := &WorkbookFunctionsAbsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Abs"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsAbsRequest struct{ BaseRequest }

func (b *WorkbookFunctionsAbsRequestBuilder) Request() *WorkbookFunctionsAbsRequest {
	return &WorkbookFunctionsAbsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsAbsRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

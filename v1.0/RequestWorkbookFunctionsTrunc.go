// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsTruncRequestBuilder struct{ BaseRequestBuilder }

// Trunc action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Trunc(reqObj *WorkbookFunctionsTruncRequestParameter) *WorkbookFunctionsTruncRequestBuilder {
	bb := &WorkbookFunctionsTruncRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Trunc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsTruncRequest struct{ BaseRequest }

func (b *WorkbookFunctionsTruncRequestBuilder) Request() *WorkbookFunctionsTruncRequest {
	return &WorkbookFunctionsTruncRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsTruncRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

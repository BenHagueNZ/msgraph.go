// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsTRequestBuilder struct{ BaseRequestBuilder }

// T action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) T(reqObj *WorkbookFunctionsTRequestParameter) *WorkbookFunctionsTRequestBuilder {
	bb := &WorkbookFunctionsTRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/T"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsTRequest struct{ BaseRequest }

func (b *WorkbookFunctionsTRequestBuilder) Request() *WorkbookFunctionsTRequest {
	return &WorkbookFunctionsTRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsTRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
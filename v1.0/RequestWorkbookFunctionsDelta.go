// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsDeltaRequestBuilder struct{ BaseRequestBuilder }

// Delta action undocumented
func (b *WorkbookFunctionsRequestBuilder) Delta(reqObj *WorkbookFunctionsDeltaRequestParameter) *WorkbookFunctionsDeltaRequestBuilder {
	bb := &WorkbookFunctionsDeltaRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Delta"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsDeltaRequest struct{ BaseRequest }

func (b *WorkbookFunctionsDeltaRequestBuilder) Request() *WorkbookFunctionsDeltaRequest {
	return &WorkbookFunctionsDeltaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsDeltaRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

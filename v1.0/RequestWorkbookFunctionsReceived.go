// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsReceivedRequestBuilder struct{ BaseRequestBuilder }

// Received action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Received(reqObj *WorkbookFunctionsReceivedRequestParameter) *WorkbookFunctionsReceivedRequestBuilder {
	bb := &WorkbookFunctionsReceivedRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Received"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsReceivedRequest struct{ BaseRequest }

func (b *WorkbookFunctionsReceivedRequestBuilder) Request() *WorkbookFunctionsReceivedRequest {
	return &WorkbookFunctionsReceivedRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsReceivedRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

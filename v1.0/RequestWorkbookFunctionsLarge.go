// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsLargeRequestBuilder struct{ BaseRequestBuilder }

// Large action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Large(reqObj *WorkbookFunctionsLargeRequestParameter) *WorkbookFunctionsLargeRequestBuilder {
	bb := &WorkbookFunctionsLargeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Large"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsLargeRequest struct{ BaseRequest }

func (b *WorkbookFunctionsLargeRequestBuilder) Request() *WorkbookFunctionsLargeRequest {
	return &WorkbookFunctionsLargeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsLargeRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

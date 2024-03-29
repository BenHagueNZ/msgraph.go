// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsPiRequestBuilder struct{ BaseRequestBuilder }

// Pi action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Pi(reqObj *WorkbookFunctionsPiRequestParameter) *WorkbookFunctionsPiRequestBuilder {
	bb := &WorkbookFunctionsPiRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Pi"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsPiRequest struct{ BaseRequest }

func (b *WorkbookFunctionsPiRequestBuilder) Request() *WorkbookFunctionsPiRequest {
	return &WorkbookFunctionsPiRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsPiRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

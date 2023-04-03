// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsBaseRequestBuilder struct{ BaseRequestBuilder }

// Base action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Base(reqObj *WorkbookFunctionsBaseRequestParameter) *WorkbookFunctionsBaseRequestBuilder {
	bb := &WorkbookFunctionsBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Base"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsBaseRequest struct{ BaseRequest }

func (b *WorkbookFunctionsBaseRequestBuilder) Request() *WorkbookFunctionsBaseRequest {
	return &WorkbookFunctionsBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsBaseRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

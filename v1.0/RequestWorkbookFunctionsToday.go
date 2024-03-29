// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsTodayRequestBuilder struct{ BaseRequestBuilder }

// Today action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Today(reqObj *WorkbookFunctionsTodayRequestParameter) *WorkbookFunctionsTodayRequestBuilder {
	bb := &WorkbookFunctionsTodayRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Today"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsTodayRequest struct{ BaseRequest }

func (b *WorkbookFunctionsTodayRequestBuilder) Request() *WorkbookFunctionsTodayRequest {
	return &WorkbookFunctionsTodayRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsTodayRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

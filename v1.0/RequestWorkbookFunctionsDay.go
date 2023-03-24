// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsDayRequestBuilder struct{ BaseRequestBuilder }

// Day action undocumented
func (b *WorkbookFunctionsRequestBuilder) Day(reqObj *WorkbookFunctionsDayRequestParameter) *WorkbookFunctionsDayRequestBuilder {
	bb := &WorkbookFunctionsDayRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Day"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsDayRequest struct{ BaseRequest }

func (b *WorkbookFunctionsDayRequestBuilder) Request() *WorkbookFunctionsDayRequest {
	return &WorkbookFunctionsDayRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsDayRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

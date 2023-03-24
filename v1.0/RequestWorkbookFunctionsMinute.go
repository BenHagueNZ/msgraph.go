// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsMinuteRequestBuilder struct{ BaseRequestBuilder }

// Minute action undocumented
func (b *WorkbookFunctionsRequestBuilder) Minute(reqObj *WorkbookFunctionsMinuteRequestParameter) *WorkbookFunctionsMinuteRequestBuilder {
	bb := &WorkbookFunctionsMinuteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Minute"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsMinuteRequest struct{ BaseRequest }

func (b *WorkbookFunctionsMinuteRequestBuilder) Request() *WorkbookFunctionsMinuteRequest {
	return &WorkbookFunctionsMinuteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsMinuteRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

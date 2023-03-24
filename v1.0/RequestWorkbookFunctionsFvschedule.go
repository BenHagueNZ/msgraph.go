// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsFvscheduleRequestBuilder struct{ BaseRequestBuilder }

// Fvschedule action undocumented
func (b *WorkbookFunctionsRequestBuilder) Fvschedule(reqObj *WorkbookFunctionsFvscheduleRequestParameter) *WorkbookFunctionsFvscheduleRequestBuilder {
	bb := &WorkbookFunctionsFvscheduleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Fvschedule"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsFvscheduleRequest struct{ BaseRequest }

func (b *WorkbookFunctionsFvscheduleRequestBuilder) Request() *WorkbookFunctionsFvscheduleRequest {
	return &WorkbookFunctionsFvscheduleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsFvscheduleRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

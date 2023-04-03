// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsWeekdayRequestBuilder struct{ BaseRequestBuilder }

// Weekday action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Weekday(reqObj *WorkbookFunctionsWeekdayRequestParameter) *WorkbookFunctionsWeekdayRequestBuilder {
	bb := &WorkbookFunctionsWeekdayRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Weekday"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsWeekdayRequest struct{ BaseRequest }

func (b *WorkbookFunctionsWeekdayRequestBuilder) Request() *WorkbookFunctionsWeekdayRequest {
	return &WorkbookFunctionsWeekdayRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsWeekdayRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

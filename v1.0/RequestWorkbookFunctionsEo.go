// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsEoMonthRequestBuilder struct{ BaseRequestBuilder }

// EoMonth action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) EoMonth(reqObj *WorkbookFunctionsEoMonthRequestParameter) *WorkbookFunctionsEoMonthRequestBuilder {
	bb := &WorkbookFunctionsEoMonthRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/EoMonth"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsEoMonthRequest struct{ BaseRequest }

func (b *WorkbookFunctionsEoMonthRequestBuilder) Request() *WorkbookFunctionsEoMonthRequest {
	return &WorkbookFunctionsEoMonthRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsEoMonthRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

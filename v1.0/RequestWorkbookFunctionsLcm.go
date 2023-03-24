// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsLcmRequestBuilder struct{ BaseRequestBuilder }

// Lcm action undocumented
func (b *WorkbookFunctionsRequestBuilder) Lcm(reqObj *WorkbookFunctionsLcmRequestParameter) *WorkbookFunctionsLcmRequestBuilder {
	bb := &WorkbookFunctionsLcmRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Lcm"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsLcmRequest struct{ BaseRequest }

func (b *WorkbookFunctionsLcmRequestBuilder) Request() *WorkbookFunctionsLcmRequest {
	return &WorkbookFunctionsLcmRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsLcmRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsRightRequestBuilder struct{ BaseRequestBuilder }

// Right action undocumented
func (b *WorkbookFunctionsRequestBuilder) Right(reqObj *WorkbookFunctionsRightRequestParameter) *WorkbookFunctionsRightRequestBuilder {
	bb := &WorkbookFunctionsRightRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Right"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsRightRequest struct{ BaseRequest }

func (b *WorkbookFunctionsRightRequestBuilder) Request() *WorkbookFunctionsRightRequest {
	return &WorkbookFunctionsRightRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsRightRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

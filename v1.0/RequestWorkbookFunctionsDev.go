// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsDevSqRequestBuilder struct{ BaseRequestBuilder }

// DevSq action undocumented
func (b *WorkbookFunctionsRequestBuilder) DevSq(reqObj *WorkbookFunctionsDevSqRequestParameter) *WorkbookFunctionsDevSqRequestBuilder {
	bb := &WorkbookFunctionsDevSqRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/DevSq"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsDevSqRequest struct{ BaseRequest }

func (b *WorkbookFunctionsDevSqRequestBuilder) Request() *WorkbookFunctionsDevSqRequest {
	return &WorkbookFunctionsDevSqRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsDevSqRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

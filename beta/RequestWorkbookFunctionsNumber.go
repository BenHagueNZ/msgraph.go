// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsNumberValueRequestBuilder struct{ BaseRequestBuilder }

// NumberValue action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) NumberValue(reqObj *WorkbookFunctionsNumberValueRequestParameter) *WorkbookFunctionsNumberValueRequestBuilder {
	bb := &WorkbookFunctionsNumberValueRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/NumberValue"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsNumberValueRequest struct{ BaseRequest }

func (b *WorkbookFunctionsNumberValueRequestBuilder) Request() *WorkbookFunctionsNumberValueRequest {
	return &WorkbookFunctionsNumberValueRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsNumberValueRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsFalseRequestBuilder struct{ BaseRequestBuilder }

// False action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) False(reqObj *WorkbookFunctionsFalseRequestParameter) *WorkbookFunctionsFalseRequestBuilder {
	bb := &WorkbookFunctionsFalseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/False"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsFalseRequest struct{ BaseRequest }

func (b *WorkbookFunctionsFalseRequestBuilder) Request() *WorkbookFunctionsFalseRequest {
	return &WorkbookFunctionsFalseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsFalseRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsXorRequestBuilder struct{ BaseRequestBuilder }

// Xor action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Xor(reqObj *WorkbookFunctionsXorRequestParameter) *WorkbookFunctionsXorRequestBuilder {
	bb := &WorkbookFunctionsXorRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Xor"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsXorRequest struct{ BaseRequest }

func (b *WorkbookFunctionsXorRequestBuilder) Request() *WorkbookFunctionsXorRequest {
	return &WorkbookFunctionsXorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsXorRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

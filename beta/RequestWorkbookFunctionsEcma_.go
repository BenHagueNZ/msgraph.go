// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsEcma_CeilingRequestBuilder struct{ BaseRequestBuilder }

// Ecma_Ceiling action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Ecma_Ceiling(reqObj *WorkbookFunctionsEcma_CeilingRequestParameter) *WorkbookFunctionsEcma_CeilingRequestBuilder {
	bb := &WorkbookFunctionsEcma_CeilingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Ecma_Ceiling"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsEcma_CeilingRequest struct{ BaseRequest }

func (b *WorkbookFunctionsEcma_CeilingRequestBuilder) Request() *WorkbookFunctionsEcma_CeilingRequest {
	return &WorkbookFunctionsEcma_CeilingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsEcma_CeilingRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

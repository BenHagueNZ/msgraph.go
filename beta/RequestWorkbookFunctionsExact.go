// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsExactRequestBuilder struct{ BaseRequestBuilder }

// Exact action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Exact(reqObj *WorkbookFunctionsExactRequestParameter) *WorkbookFunctionsExactRequestBuilder {
	bb := &WorkbookFunctionsExactRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Exact"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsExactRequest struct{ BaseRequest }

func (b *WorkbookFunctionsExactRequestBuilder) Request() *WorkbookFunctionsExactRequest {
	return &WorkbookFunctionsExactRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsExactRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

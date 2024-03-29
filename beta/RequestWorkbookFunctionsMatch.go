// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsMatchRequestBuilder struct{ BaseRequestBuilder }

// Match action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Match(reqObj *WorkbookFunctionsMatchRequestParameter) *WorkbookFunctionsMatchRequestBuilder {
	bb := &WorkbookFunctionsMatchRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Match"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsMatchRequest struct{ BaseRequest }

func (b *WorkbookFunctionsMatchRequestBuilder) Request() *WorkbookFunctionsMatchRequest {
	return &WorkbookFunctionsMatchRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsMatchRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

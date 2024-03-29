// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsDiscRequestBuilder struct{ BaseRequestBuilder }

// Disc action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Disc(reqObj *WorkbookFunctionsDiscRequestParameter) *WorkbookFunctionsDiscRequestBuilder {
	bb := &WorkbookFunctionsDiscRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Disc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsDiscRequest struct{ BaseRequest }

func (b *WorkbookFunctionsDiscRequestBuilder) Request() *WorkbookFunctionsDiscRequest {
	return &WorkbookFunctionsDiscRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsDiscRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

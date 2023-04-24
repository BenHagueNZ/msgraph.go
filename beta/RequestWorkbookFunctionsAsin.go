// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsAsinRequestBuilder struct{ BaseRequestBuilder }

// Asin action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Asin(reqObj *WorkbookFunctionsAsinRequestParameter) *WorkbookFunctionsAsinRequestBuilder {
	bb := &WorkbookFunctionsAsinRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Asin"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsAsinRequest struct{ BaseRequest }

func (b *WorkbookFunctionsAsinRequestBuilder) Request() *WorkbookFunctionsAsinRequest {
	return &WorkbookFunctionsAsinRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsAsinRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

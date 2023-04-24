// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsTimevalueRequestBuilder struct{ BaseRequestBuilder }

// Timevalue action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Timevalue(reqObj *WorkbookFunctionsTimevalueRequestParameter) *WorkbookFunctionsTimevalueRequestBuilder {
	bb := &WorkbookFunctionsTimevalueRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Timevalue"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsTimevalueRequest struct{ BaseRequest }

func (b *WorkbookFunctionsTimevalueRequestBuilder) Request() *WorkbookFunctionsTimevalueRequest {
	return &WorkbookFunctionsTimevalueRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsTimevalueRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

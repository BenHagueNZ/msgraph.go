// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsDmaxRequestBuilder struct{ BaseRequestBuilder }

// Dmax action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Dmax(reqObj *WorkbookFunctionsDmaxRequestParameter) *WorkbookFunctionsDmaxRequestBuilder {
	bb := &WorkbookFunctionsDmaxRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Dmax"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsDmaxRequest struct{ BaseRequest }

func (b *WorkbookFunctionsDmaxRequestBuilder) Request() *WorkbookFunctionsDmaxRequest {
	return &WorkbookFunctionsDmaxRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsDmaxRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

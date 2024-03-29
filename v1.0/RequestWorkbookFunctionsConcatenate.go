// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsConcatenateRequestBuilder struct{ BaseRequestBuilder }

// Concatenate action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Concatenate(reqObj *WorkbookFunctionsConcatenateRequestParameter) *WorkbookFunctionsConcatenateRequestBuilder {
	bb := &WorkbookFunctionsConcatenateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Concatenate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsConcatenateRequest struct{ BaseRequest }

func (b *WorkbookFunctionsConcatenateRequestBuilder) Request() *WorkbookFunctionsConcatenateRequest {
	return &WorkbookFunctionsConcatenateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsConcatenateRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

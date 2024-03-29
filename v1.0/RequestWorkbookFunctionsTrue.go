// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsTrueRequestBuilder struct{ BaseRequestBuilder }

// True action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) True(reqObj *WorkbookFunctionsTrueRequestParameter) *WorkbookFunctionsTrueRequestBuilder {
	bb := &WorkbookFunctionsTrueRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/True"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsTrueRequest struct{ BaseRequest }

func (b *WorkbookFunctionsTrueRequestBuilder) Request() *WorkbookFunctionsTrueRequest {
	return &WorkbookFunctionsTrueRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsTrueRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

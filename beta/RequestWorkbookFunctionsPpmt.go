// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsPpmtRequestBuilder struct{ BaseRequestBuilder }

// Ppmt action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Ppmt(reqObj *WorkbookFunctionsPpmtRequestParameter) *WorkbookFunctionsPpmtRequestBuilder {
	bb := &WorkbookFunctionsPpmtRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Ppmt"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsPpmtRequest struct{ BaseRequest }

func (b *WorkbookFunctionsPpmtRequestBuilder) Request() *WorkbookFunctionsPpmtRequest {
	return &WorkbookFunctionsPpmtRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsPpmtRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

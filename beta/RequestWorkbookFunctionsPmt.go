// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsPmtRequestBuilder struct{ BaseRequestBuilder }

// Pmt action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Pmt(reqObj *WorkbookFunctionsPmtRequestParameter) *WorkbookFunctionsPmtRequestBuilder {
	bb := &WorkbookFunctionsPmtRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Pmt"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsPmtRequest struct{ BaseRequest }

func (b *WorkbookFunctionsPmtRequestBuilder) Request() *WorkbookFunctionsPmtRequest {
	return &WorkbookFunctionsPmtRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsPmtRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsIpmtRequestBuilder struct{ BaseRequestBuilder }

// Ipmt action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Ipmt(reqObj *WorkbookFunctionsIpmtRequestParameter) *WorkbookFunctionsIpmtRequestBuilder {
	bb := &WorkbookFunctionsIpmtRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Ipmt"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsIpmtRequest struct{ BaseRequest }

func (b *WorkbookFunctionsIpmtRequestBuilder) Request() *WorkbookFunctionsIpmtRequest {
	return &WorkbookFunctionsIpmtRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsIpmtRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

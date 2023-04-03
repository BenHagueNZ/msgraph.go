// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsDbcsRequestBuilder struct{ BaseRequestBuilder }

// Dbcs action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Dbcs(reqObj *WorkbookFunctionsDbcsRequestParameter) *WorkbookFunctionsDbcsRequestBuilder {
	bb := &WorkbookFunctionsDbcsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Dbcs"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsDbcsRequest struct{ BaseRequest }

func (b *WorkbookFunctionsDbcsRequestBuilder) Request() *WorkbookFunctionsDbcsRequest {
	return &WorkbookFunctionsDbcsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsDbcsRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

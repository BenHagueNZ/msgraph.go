// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsAveDevRequestBuilder struct{ BaseRequestBuilder }

// AveDev action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) AveDev(reqObj *WorkbookFunctionsAveDevRequestParameter) *WorkbookFunctionsAveDevRequestBuilder {
	bb := &WorkbookFunctionsAveDevRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AveDev"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsAveDevRequest struct{ BaseRequest }

func (b *WorkbookFunctionsAveDevRequestBuilder) Request() *WorkbookFunctionsAveDevRequest {
	return &WorkbookFunctionsAveDevRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsAveDevRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
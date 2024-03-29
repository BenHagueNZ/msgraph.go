// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsMdurationRequestBuilder struct{ BaseRequestBuilder }

// Mduration action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Mduration(reqObj *WorkbookFunctionsMdurationRequestParameter) *WorkbookFunctionsMdurationRequestBuilder {
	bb := &WorkbookFunctionsMdurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Mduration"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsMdurationRequest struct{ BaseRequest }

func (b *WorkbookFunctionsMdurationRequestBuilder) Request() *WorkbookFunctionsMdurationRequest {
	return &WorkbookFunctionsMdurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsMdurationRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

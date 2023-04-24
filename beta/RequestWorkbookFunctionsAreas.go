// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsAreasRequestBuilder struct{ BaseRequestBuilder }

// Areas action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Areas(reqObj *WorkbookFunctionsAreasRequestParameter) *WorkbookFunctionsAreasRequestBuilder {
	bb := &WorkbookFunctionsAreasRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Areas"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsAreasRequest struct{ BaseRequest }

func (b *WorkbookFunctionsAreasRequestBuilder) Request() *WorkbookFunctionsAreasRequest {
	return &WorkbookFunctionsAreasRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsAreasRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
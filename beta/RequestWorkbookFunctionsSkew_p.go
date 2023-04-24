// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsSkew_pRequestBuilder struct{ BaseRequestBuilder }

// Skew_p action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Skew_p(reqObj *WorkbookFunctionsSkew_pRequestParameter) *WorkbookFunctionsSkew_pRequestBuilder {
	bb := &WorkbookFunctionsSkew_pRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Skew_p"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsSkew_pRequest struct{ BaseRequest }

func (b *WorkbookFunctionsSkew_pRequestBuilder) Request() *WorkbookFunctionsSkew_pRequest {
	return &WorkbookFunctionsSkew_pRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsSkew_pRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsRateRequestBuilder struct{ BaseRequestBuilder }

// Rate action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Rate(reqObj *WorkbookFunctionsRateRequestParameter) *WorkbookFunctionsRateRequestBuilder {
	bb := &WorkbookFunctionsRateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Rate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsRateRequest struct{ BaseRequest }

func (b *WorkbookFunctionsRateRequestBuilder) Request() *WorkbookFunctionsRateRequest {
	return &WorkbookFunctionsRateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsRateRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

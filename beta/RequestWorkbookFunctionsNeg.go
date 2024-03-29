// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsNegBinom_DistRequestBuilder struct{ BaseRequestBuilder }

// NegBinom_Dist action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) NegBinom_Dist(reqObj *WorkbookFunctionsNegBinom_DistRequestParameter) *WorkbookFunctionsNegBinom_DistRequestBuilder {
	bb := &WorkbookFunctionsNegBinom_DistRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/NegBinom_Dist"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsNegBinom_DistRequest struct{ BaseRequest }

func (b *WorkbookFunctionsNegBinom_DistRequestBuilder) Request() *WorkbookFunctionsNegBinom_DistRequest {
	return &WorkbookFunctionsNegBinom_DistRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsNegBinom_DistRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

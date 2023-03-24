// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsPoisson_DistRequestBuilder struct{ BaseRequestBuilder }

// Poisson_Dist action undocumented
func (b *WorkbookFunctionsRequestBuilder) Poisson_Dist(reqObj *WorkbookFunctionsPoisson_DistRequestParameter) *WorkbookFunctionsPoisson_DistRequestBuilder {
	bb := &WorkbookFunctionsPoisson_DistRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Poisson_Dist"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsPoisson_DistRequest struct{ BaseRequest }

func (b *WorkbookFunctionsPoisson_DistRequestBuilder) Request() *WorkbookFunctionsPoisson_DistRequest {
	return &WorkbookFunctionsPoisson_DistRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsPoisson_DistRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

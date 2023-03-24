// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsWeibull_DistRequestBuilder struct{ BaseRequestBuilder }

// Weibull_Dist action undocumented
func (b *WorkbookFunctionsRequestBuilder) Weibull_Dist(reqObj *WorkbookFunctionsWeibull_DistRequestParameter) *WorkbookFunctionsWeibull_DistRequestBuilder {
	bb := &WorkbookFunctionsWeibull_DistRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Weibull_Dist"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsWeibull_DistRequest struct{ BaseRequest }

func (b *WorkbookFunctionsWeibull_DistRequestBuilder) Request() *WorkbookFunctionsWeibull_DistRequest {
	return &WorkbookFunctionsWeibull_DistRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsWeibull_DistRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

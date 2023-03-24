// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsSubstituteRequestBuilder struct{ BaseRequestBuilder }

// Substitute action undocumented
func (b *WorkbookFunctionsRequestBuilder) Substitute(reqObj *WorkbookFunctionsSubstituteRequestParameter) *WorkbookFunctionsSubstituteRequestBuilder {
	bb := &WorkbookFunctionsSubstituteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Substitute"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsSubstituteRequest struct{ BaseRequest }

func (b *WorkbookFunctionsSubstituteRequestBuilder) Request() *WorkbookFunctionsSubstituteRequest {
	return &WorkbookFunctionsSubstituteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsSubstituteRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

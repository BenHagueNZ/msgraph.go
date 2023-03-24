// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsMultiNomialRequestBuilder struct{ BaseRequestBuilder }

// MultiNomial action undocumented
func (b *WorkbookFunctionsRequestBuilder) MultiNomial(reqObj *WorkbookFunctionsMultiNomialRequestParameter) *WorkbookFunctionsMultiNomialRequestBuilder {
	bb := &WorkbookFunctionsMultiNomialRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/MultiNomial"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsMultiNomialRequest struct{ BaseRequest }

func (b *WorkbookFunctionsMultiNomialRequestBuilder) Request() *WorkbookFunctionsMultiNomialRequest {
	return &WorkbookFunctionsMultiNomialRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsMultiNomialRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsChooseRequestBuilder struct{ BaseRequestBuilder }

// Choose action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Choose(reqObj *WorkbookFunctionsChooseRequestParameter) *WorkbookFunctionsChooseRequestBuilder {
	bb := &WorkbookFunctionsChooseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Choose"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsChooseRequest struct{ BaseRequest }

func (b *WorkbookFunctionsChooseRequestBuilder) Request() *WorkbookFunctionsChooseRequest {
	return &WorkbookFunctionsChooseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsChooseRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

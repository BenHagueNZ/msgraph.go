// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsSheetsRequestBuilder struct{ BaseRequestBuilder }

// Sheets action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Sheets(reqObj *WorkbookFunctionsSheetsRequestParameter) *WorkbookFunctionsSheetsRequestBuilder {
	bb := &WorkbookFunctionsSheetsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Sheets"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsSheetsRequest struct{ BaseRequest }

func (b *WorkbookFunctionsSheetsRequestBuilder) Request() *WorkbookFunctionsSheetsRequest {
	return &WorkbookFunctionsSheetsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsSheetsRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
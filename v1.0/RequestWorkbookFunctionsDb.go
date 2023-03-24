// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsDbRequestBuilder struct{ BaseRequestBuilder }

// Db action undocumented
func (b *WorkbookFunctionsRequestBuilder) Db(reqObj *WorkbookFunctionsDbRequestParameter) *WorkbookFunctionsDbRequestBuilder {
	bb := &WorkbookFunctionsDbRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Db"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsDbRequest struct{ BaseRequest }

func (b *WorkbookFunctionsDbRequestBuilder) Request() *WorkbookFunctionsDbRequest {
	return &WorkbookFunctionsDbRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsDbRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

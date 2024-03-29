// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsAcothRequestBuilder struct{ BaseRequestBuilder }

// Acoth action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Acoth(reqObj *WorkbookFunctionsAcothRequestParameter) *WorkbookFunctionsAcothRequestBuilder {
	bb := &WorkbookFunctionsAcothRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Acoth"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsAcothRequest struct{ BaseRequest }

func (b *WorkbookFunctionsAcothRequestBuilder) Request() *WorkbookFunctionsAcothRequest {
	return &WorkbookFunctionsAcothRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsAcothRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

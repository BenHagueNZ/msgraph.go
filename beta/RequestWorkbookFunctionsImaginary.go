// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsImaginaryRequestBuilder struct{ BaseRequestBuilder }

// Imaginary action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Imaginary(reqObj *WorkbookFunctionsImaginaryRequestParameter) *WorkbookFunctionsImaginaryRequestBuilder {
	bb := &WorkbookFunctionsImaginaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Imaginary"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImaginaryRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImaginaryRequestBuilder) Request() *WorkbookFunctionsImaginaryRequest {
	return &WorkbookFunctionsImaginaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImaginaryRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
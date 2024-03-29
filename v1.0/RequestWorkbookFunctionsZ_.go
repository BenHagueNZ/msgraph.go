// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsZ_TestRequestBuilder struct{ BaseRequestBuilder }

// Z_Test action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Z_Test(reqObj *WorkbookFunctionsZ_TestRequestParameter) *WorkbookFunctionsZ_TestRequestBuilder {
	bb := &WorkbookFunctionsZ_TestRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Z_Test"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsZ_TestRequest struct{ BaseRequest }

func (b *WorkbookFunctionsZ_TestRequestBuilder) Request() *WorkbookFunctionsZ_TestRequest {
	return &WorkbookFunctionsZ_TestRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsZ_TestRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

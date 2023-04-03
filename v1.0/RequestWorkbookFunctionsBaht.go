// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsBahtTextRequestBuilder struct{ BaseRequestBuilder }

// BahtText action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) BahtText(reqObj *WorkbookFunctionsBahtTextRequestParameter) *WorkbookFunctionsBahtTextRequestBuilder {
	bb := &WorkbookFunctionsBahtTextRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/BahtText"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsBahtTextRequest struct{ BaseRequest }

func (b *WorkbookFunctionsBahtTextRequestBuilder) Request() *WorkbookFunctionsBahtTextRequest {
	return &WorkbookFunctionsBahtTextRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsBahtTextRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

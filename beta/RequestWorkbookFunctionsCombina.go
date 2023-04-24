// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsCombinaRequestBuilder struct{ BaseRequestBuilder }

// Combina action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Combina(reqObj *WorkbookFunctionsCombinaRequestParameter) *WorkbookFunctionsCombinaRequestBuilder {
	bb := &WorkbookFunctionsCombinaRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Combina"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsCombinaRequest struct{ BaseRequest }

func (b *WorkbookFunctionsCombinaRequestBuilder) Request() *WorkbookFunctionsCombinaRequest {
	return &WorkbookFunctionsCombinaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsCombinaRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
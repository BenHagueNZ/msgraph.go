// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsVlookupRequestBuilder struct{ BaseRequestBuilder }

// Vlookup action undocumented
func (b *WorkbookFunctionsRequestBuilder) Vlookup(reqObj *WorkbookFunctionsVlookupRequestParameter) *WorkbookFunctionsVlookupRequestBuilder {
	bb := &WorkbookFunctionsVlookupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Vlookup"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsVlookupRequest struct{ BaseRequest }

func (b *WorkbookFunctionsVlookupRequestBuilder) Request() *WorkbookFunctionsVlookupRequest {
	return &WorkbookFunctionsVlookupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsVlookupRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

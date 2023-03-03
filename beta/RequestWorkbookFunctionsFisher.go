// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsFisherRequestBuilder struct{ BaseRequestBuilder }

// Fisher action undocumented
func (b *WorkbookFunctionsRequestBuilder) Fisher(reqObj *WorkbookFunctionsFisherRequestParameter) *WorkbookFunctionsFisherRequestBuilder {
	bb := &WorkbookFunctionsFisherRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/fisher"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsFisherRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsFisherRequestBuilder) Request() *WorkbookFunctionsFisherRequest {
	return &WorkbookFunctionsFisherRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsFisherRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type WorkbookFunctionsFisherInvRequestBuilder struct{ BaseRequestBuilder }

// FisherInv action undocumented
func (b *WorkbookFunctionsRequestBuilder) FisherInv(reqObj *WorkbookFunctionsFisherInvRequestParameter) *WorkbookFunctionsFisherInvRequestBuilder {
	bb := &WorkbookFunctionsFisherInvRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/fisherInv"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsFisherInvRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsFisherInvRequestBuilder) Request() *WorkbookFunctionsFisherInvRequest {
	return &WorkbookFunctionsFisherInvRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsFisherInvRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

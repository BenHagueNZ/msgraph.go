// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsDollarRequestBuilder struct{ BaseRequestBuilder }

// Dollar action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Dollar(reqObj *WorkbookFunctionsDollarRequestParameter) *WorkbookFunctionsDollarRequestBuilder {
	bb := &WorkbookFunctionsDollarRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Dollar"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsDollarRequest struct{ BaseRequest }

func (b *WorkbookFunctionsDollarRequestBuilder) Request() *WorkbookFunctionsDollarRequest {
	return &WorkbookFunctionsDollarRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsDollarRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsDollarDeRequestBuilder struct{ BaseRequestBuilder }

// DollarDe action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) DollarDe(reqObj *WorkbookFunctionsDollarDeRequestParameter) *WorkbookFunctionsDollarDeRequestBuilder {
	bb := &WorkbookFunctionsDollarDeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/DollarDe"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsDollarDeRequest struct{ BaseRequest }

func (b *WorkbookFunctionsDollarDeRequestBuilder) Request() *WorkbookFunctionsDollarDeRequest {
	return &WorkbookFunctionsDollarDeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsDollarDeRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsDollarFrRequestBuilder struct{ BaseRequestBuilder }

// DollarFr action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) DollarFr(reqObj *WorkbookFunctionsDollarFrRequestParameter) *WorkbookFunctionsDollarFrRequestBuilder {
	bb := &WorkbookFunctionsDollarFrRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/DollarFr"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsDollarFrRequest struct{ BaseRequest }

func (b *WorkbookFunctionsDollarFrRequestBuilder) Request() *WorkbookFunctionsDollarFrRequest {
	return &WorkbookFunctionsDollarFrRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsDollarFrRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

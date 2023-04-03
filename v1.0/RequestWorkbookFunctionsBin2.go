// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsBin2DecRequestBuilder struct{ BaseRequestBuilder }

// Bin2Dec action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Bin2Dec(reqObj *WorkbookFunctionsBin2DecRequestParameter) *WorkbookFunctionsBin2DecRequestBuilder {
	bb := &WorkbookFunctionsBin2DecRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Bin2Dec"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsBin2DecRequest struct{ BaseRequest }

func (b *WorkbookFunctionsBin2DecRequestBuilder) Request() *WorkbookFunctionsBin2DecRequest {
	return &WorkbookFunctionsBin2DecRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsBin2DecRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsBin2HexRequestBuilder struct{ BaseRequestBuilder }

// Bin2Hex action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Bin2Hex(reqObj *WorkbookFunctionsBin2HexRequestParameter) *WorkbookFunctionsBin2HexRequestBuilder {
	bb := &WorkbookFunctionsBin2HexRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Bin2Hex"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsBin2HexRequest struct{ BaseRequest }

func (b *WorkbookFunctionsBin2HexRequestBuilder) Request() *WorkbookFunctionsBin2HexRequest {
	return &WorkbookFunctionsBin2HexRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsBin2HexRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsBin2OctRequestBuilder struct{ BaseRequestBuilder }

// Bin2Oct action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Bin2Oct(reqObj *WorkbookFunctionsBin2OctRequestParameter) *WorkbookFunctionsBin2OctRequestBuilder {
	bb := &WorkbookFunctionsBin2OctRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Bin2Oct"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsBin2OctRequest struct{ BaseRequest }

func (b *WorkbookFunctionsBin2OctRequestBuilder) Request() *WorkbookFunctionsBin2OctRequest {
	return &WorkbookFunctionsBin2OctRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsBin2OctRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

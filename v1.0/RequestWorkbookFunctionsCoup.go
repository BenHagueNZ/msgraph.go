// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsCoupDayBsRequestBuilder struct{ BaseRequestBuilder }

// CoupDayBs action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) CoupDayBs(reqObj *WorkbookFunctionsCoupDayBsRequestParameter) *WorkbookFunctionsCoupDayBsRequestBuilder {
	bb := &WorkbookFunctionsCoupDayBsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CoupDayBs"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsCoupDayBsRequest struct{ BaseRequest }

func (b *WorkbookFunctionsCoupDayBsRequestBuilder) Request() *WorkbookFunctionsCoupDayBsRequest {
	return &WorkbookFunctionsCoupDayBsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsCoupDayBsRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsCoupDaysRequestBuilder struct{ BaseRequestBuilder }

// CoupDays action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) CoupDays(reqObj *WorkbookFunctionsCoupDaysRequestParameter) *WorkbookFunctionsCoupDaysRequestBuilder {
	bb := &WorkbookFunctionsCoupDaysRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CoupDays"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsCoupDaysRequest struct{ BaseRequest }

func (b *WorkbookFunctionsCoupDaysRequestBuilder) Request() *WorkbookFunctionsCoupDaysRequest {
	return &WorkbookFunctionsCoupDaysRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsCoupDaysRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsCoupDaysNcRequestBuilder struct{ BaseRequestBuilder }

// CoupDaysNc action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) CoupDaysNc(reqObj *WorkbookFunctionsCoupDaysNcRequestParameter) *WorkbookFunctionsCoupDaysNcRequestBuilder {
	bb := &WorkbookFunctionsCoupDaysNcRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CoupDaysNc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsCoupDaysNcRequest struct{ BaseRequest }

func (b *WorkbookFunctionsCoupDaysNcRequestBuilder) Request() *WorkbookFunctionsCoupDaysNcRequest {
	return &WorkbookFunctionsCoupDaysNcRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsCoupDaysNcRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsCoupNcdRequestBuilder struct{ BaseRequestBuilder }

// CoupNcd action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) CoupNcd(reqObj *WorkbookFunctionsCoupNcdRequestParameter) *WorkbookFunctionsCoupNcdRequestBuilder {
	bb := &WorkbookFunctionsCoupNcdRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CoupNcd"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsCoupNcdRequest struct{ BaseRequest }

func (b *WorkbookFunctionsCoupNcdRequestBuilder) Request() *WorkbookFunctionsCoupNcdRequest {
	return &WorkbookFunctionsCoupNcdRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsCoupNcdRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsCoupNumRequestBuilder struct{ BaseRequestBuilder }

// CoupNum action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) CoupNum(reqObj *WorkbookFunctionsCoupNumRequestParameter) *WorkbookFunctionsCoupNumRequestBuilder {
	bb := &WorkbookFunctionsCoupNumRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CoupNum"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsCoupNumRequest struct{ BaseRequest }

func (b *WorkbookFunctionsCoupNumRequestBuilder) Request() *WorkbookFunctionsCoupNumRequest {
	return &WorkbookFunctionsCoupNumRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsCoupNumRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsCoupPcdRequestBuilder struct{ BaseRequestBuilder }

// CoupPcd action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) CoupPcd(reqObj *WorkbookFunctionsCoupPcdRequestParameter) *WorkbookFunctionsCoupPcdRequestBuilder {
	bb := &WorkbookFunctionsCoupPcdRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CoupPcd"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsCoupPcdRequest struct{ BaseRequest }

func (b *WorkbookFunctionsCoupPcdRequestBuilder) Request() *WorkbookFunctionsCoupPcdRequest {
	return &WorkbookFunctionsCoupPcdRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsCoupPcdRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

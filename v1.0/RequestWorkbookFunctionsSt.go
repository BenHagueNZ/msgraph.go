// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsStDevARequestBuilder struct{ BaseRequestBuilder }

// StDevA action undocumented
func (b *WorkbookFunctionsRequestBuilder) StDevA(reqObj *WorkbookFunctionsStDevARequestParameter) *WorkbookFunctionsStDevARequestBuilder {
	bb := &WorkbookFunctionsStDevARequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/StDevA"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsStDevARequest struct{ BaseRequest }

func (b *WorkbookFunctionsStDevARequestBuilder) Request() *WorkbookFunctionsStDevARequest {
	return &WorkbookFunctionsStDevARequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsStDevARequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsStDevPARequestBuilder struct{ BaseRequestBuilder }

// StDevPA action undocumented
func (b *WorkbookFunctionsRequestBuilder) StDevPA(reqObj *WorkbookFunctionsStDevPARequestParameter) *WorkbookFunctionsStDevPARequestBuilder {
	bb := &WorkbookFunctionsStDevPARequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/StDevPA"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsStDevPARequest struct{ BaseRequest }

func (b *WorkbookFunctionsStDevPARequestBuilder) Request() *WorkbookFunctionsStDevPARequest {
	return &WorkbookFunctionsStDevPARequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsStDevPARequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsStDev_PRequestBuilder struct{ BaseRequestBuilder }

// StDev_P action undocumented
func (b *WorkbookFunctionsRequestBuilder) StDev_P(reqObj *WorkbookFunctionsStDev_PRequestParameter) *WorkbookFunctionsStDev_PRequestBuilder {
	bb := &WorkbookFunctionsStDev_PRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/StDev_P"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsStDev_PRequest struct{ BaseRequest }

func (b *WorkbookFunctionsStDev_PRequestBuilder) Request() *WorkbookFunctionsStDev_PRequest {
	return &WorkbookFunctionsStDev_PRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsStDev_PRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsStDev_SRequestBuilder struct{ BaseRequestBuilder }

// StDev_S action undocumented
func (b *WorkbookFunctionsRequestBuilder) StDev_S(reqObj *WorkbookFunctionsStDev_SRequestParameter) *WorkbookFunctionsStDev_SRequestBuilder {
	bb := &WorkbookFunctionsStDev_SRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/StDev_S"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsStDev_SRequest struct{ BaseRequest }

func (b *WorkbookFunctionsStDev_SRequestBuilder) Request() *WorkbookFunctionsStDev_SRequest {
	return &WorkbookFunctionsStDev_SRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsStDev_SRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

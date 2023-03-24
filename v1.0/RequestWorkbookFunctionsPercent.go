// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsPercentRank_ExcRequestBuilder struct{ BaseRequestBuilder }

// PercentRank_Exc action undocumented
func (b *WorkbookFunctionsRequestBuilder) PercentRank_Exc(reqObj *WorkbookFunctionsPercentRank_ExcRequestParameter) *WorkbookFunctionsPercentRank_ExcRequestBuilder {
	bb := &WorkbookFunctionsPercentRank_ExcRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/PercentRank_Exc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsPercentRank_ExcRequest struct{ BaseRequest }

func (b *WorkbookFunctionsPercentRank_ExcRequestBuilder) Request() *WorkbookFunctionsPercentRank_ExcRequest {
	return &WorkbookFunctionsPercentRank_ExcRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsPercentRank_ExcRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsPercentRank_IncRequestBuilder struct{ BaseRequestBuilder }

// PercentRank_Inc action undocumented
func (b *WorkbookFunctionsRequestBuilder) PercentRank_Inc(reqObj *WorkbookFunctionsPercentRank_IncRequestParameter) *WorkbookFunctionsPercentRank_IncRequestBuilder {
	bb := &WorkbookFunctionsPercentRank_IncRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/PercentRank_Inc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsPercentRank_IncRequest struct{ BaseRequest }

func (b *WorkbookFunctionsPercentRank_IncRequestBuilder) Request() *WorkbookFunctionsPercentRank_IncRequest {
	return &WorkbookFunctionsPercentRank_IncRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsPercentRank_IncRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

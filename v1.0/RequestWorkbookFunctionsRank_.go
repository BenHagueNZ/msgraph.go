// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsRank_AvgRequestBuilder struct{ BaseRequestBuilder }

// Rank_Avg action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Rank_Avg(reqObj *WorkbookFunctionsRank_AvgRequestParameter) *WorkbookFunctionsRank_AvgRequestBuilder {
	bb := &WorkbookFunctionsRank_AvgRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Rank_Avg"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsRank_AvgRequest struct{ BaseRequest }

func (b *WorkbookFunctionsRank_AvgRequestBuilder) Request() *WorkbookFunctionsRank_AvgRequest {
	return &WorkbookFunctionsRank_AvgRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsRank_AvgRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsRank_EqRequestBuilder struct{ BaseRequestBuilder }

// Rank_Eq action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) Rank_Eq(reqObj *WorkbookFunctionsRank_EqRequestParameter) *WorkbookFunctionsRank_EqRequestBuilder {
	bb := &WorkbookFunctionsRank_EqRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Rank_Eq"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsRank_EqRequest struct{ BaseRequest }

func (b *WorkbookFunctionsRank_EqRequestBuilder) Request() *WorkbookFunctionsRank_EqRequest {
	return &WorkbookFunctionsRank_EqRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsRank_EqRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

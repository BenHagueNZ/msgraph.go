// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsHarMeanRequestBuilder struct{ BaseRequestBuilder }

// HarMean action undocumentedras
func (b *WorkbookFunctionsRequestBuilder) HarMean(reqObj *WorkbookFunctionsHarMeanRequestParameter) *WorkbookFunctionsHarMeanRequestBuilder {
	bb := &WorkbookFunctionsHarMeanRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/HarMean"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsHarMeanRequest struct{ BaseRequest }

func (b *WorkbookFunctionsHarMeanRequestBuilder) Request() *WorkbookFunctionsHarMeanRequest {
	return &WorkbookFunctionsHarMeanRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsHarMeanRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

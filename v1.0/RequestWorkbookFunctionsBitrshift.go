// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsBitrshiftRequestBuilder struct{ BaseRequestBuilder }

// Bitrshift action undocumented
func (b *WorkbookFunctionsRequestBuilder) Bitrshift(reqObj *WorkbookFunctionsBitrshiftRequestParameter) *WorkbookFunctionsBitrshiftRequestBuilder {
	bb := &WorkbookFunctionsBitrshiftRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Bitrshift"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsBitrshiftRequest struct{ BaseRequest }

func (b *WorkbookFunctionsBitrshiftRequestBuilder) Request() *WorkbookFunctionsBitrshiftRequest {
	return &WorkbookFunctionsBitrshiftRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsBitrshiftRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

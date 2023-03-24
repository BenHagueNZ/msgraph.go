// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

type WorkbookFunctionsImAbsRequestBuilder struct{ BaseRequestBuilder }

// ImAbs action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImAbs(reqObj *WorkbookFunctionsImAbsRequestParameter) *WorkbookFunctionsImAbsRequestBuilder {
	bb := &WorkbookFunctionsImAbsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImAbs"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImAbsRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImAbsRequestBuilder) Request() *WorkbookFunctionsImAbsRequest {
	return &WorkbookFunctionsImAbsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImAbsRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImArgumentRequestBuilder struct{ BaseRequestBuilder }

// ImArgument action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImArgument(reqObj *WorkbookFunctionsImArgumentRequestParameter) *WorkbookFunctionsImArgumentRequestBuilder {
	bb := &WorkbookFunctionsImArgumentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImArgument"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImArgumentRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImArgumentRequestBuilder) Request() *WorkbookFunctionsImArgumentRequest {
	return &WorkbookFunctionsImArgumentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImArgumentRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImConjugateRequestBuilder struct{ BaseRequestBuilder }

// ImConjugate action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImConjugate(reqObj *WorkbookFunctionsImConjugateRequestParameter) *WorkbookFunctionsImConjugateRequestBuilder {
	bb := &WorkbookFunctionsImConjugateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImConjugate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImConjugateRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImConjugateRequestBuilder) Request() *WorkbookFunctionsImConjugateRequest {
	return &WorkbookFunctionsImConjugateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImConjugateRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImCosRequestBuilder struct{ BaseRequestBuilder }

// ImCos action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImCos(reqObj *WorkbookFunctionsImCosRequestParameter) *WorkbookFunctionsImCosRequestBuilder {
	bb := &WorkbookFunctionsImCosRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImCos"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImCosRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImCosRequestBuilder) Request() *WorkbookFunctionsImCosRequest {
	return &WorkbookFunctionsImCosRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImCosRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImCoshRequestBuilder struct{ BaseRequestBuilder }

// ImCosh action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImCosh(reqObj *WorkbookFunctionsImCoshRequestParameter) *WorkbookFunctionsImCoshRequestBuilder {
	bb := &WorkbookFunctionsImCoshRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImCosh"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImCoshRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImCoshRequestBuilder) Request() *WorkbookFunctionsImCoshRequest {
	return &WorkbookFunctionsImCoshRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImCoshRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImCotRequestBuilder struct{ BaseRequestBuilder }

// ImCot action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImCot(reqObj *WorkbookFunctionsImCotRequestParameter) *WorkbookFunctionsImCotRequestBuilder {
	bb := &WorkbookFunctionsImCotRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImCot"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImCotRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImCotRequestBuilder) Request() *WorkbookFunctionsImCotRequest {
	return &WorkbookFunctionsImCotRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImCotRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImCscRequestBuilder struct{ BaseRequestBuilder }

// ImCsc action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImCsc(reqObj *WorkbookFunctionsImCscRequestParameter) *WorkbookFunctionsImCscRequestBuilder {
	bb := &WorkbookFunctionsImCscRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImCsc"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImCscRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImCscRequestBuilder) Request() *WorkbookFunctionsImCscRequest {
	return &WorkbookFunctionsImCscRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImCscRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImCschRequestBuilder struct{ BaseRequestBuilder }

// ImCsch action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImCsch(reqObj *WorkbookFunctionsImCschRequestParameter) *WorkbookFunctionsImCschRequestBuilder {
	bb := &WorkbookFunctionsImCschRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImCsch"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImCschRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImCschRequestBuilder) Request() *WorkbookFunctionsImCschRequest {
	return &WorkbookFunctionsImCschRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImCschRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImDivRequestBuilder struct{ BaseRequestBuilder }

// ImDiv action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImDiv(reqObj *WorkbookFunctionsImDivRequestParameter) *WorkbookFunctionsImDivRequestBuilder {
	bb := &WorkbookFunctionsImDivRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImDiv"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImDivRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImDivRequestBuilder) Request() *WorkbookFunctionsImDivRequest {
	return &WorkbookFunctionsImDivRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImDivRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImExpRequestBuilder struct{ BaseRequestBuilder }

// ImExp action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImExp(reqObj *WorkbookFunctionsImExpRequestParameter) *WorkbookFunctionsImExpRequestBuilder {
	bb := &WorkbookFunctionsImExpRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImExp"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImExpRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImExpRequestBuilder) Request() *WorkbookFunctionsImExpRequest {
	return &WorkbookFunctionsImExpRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImExpRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImLnRequestBuilder struct{ BaseRequestBuilder }

// ImLn action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImLn(reqObj *WorkbookFunctionsImLnRequestParameter) *WorkbookFunctionsImLnRequestBuilder {
	bb := &WorkbookFunctionsImLnRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImLn"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImLnRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImLnRequestBuilder) Request() *WorkbookFunctionsImLnRequest {
	return &WorkbookFunctionsImLnRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImLnRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImLog10RequestBuilder struct{ BaseRequestBuilder }

// ImLog10 action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImLog10(reqObj *WorkbookFunctionsImLog10RequestParameter) *WorkbookFunctionsImLog10RequestBuilder {
	bb := &WorkbookFunctionsImLog10RequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImLog10"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImLog10Request struct{ BaseRequest }

func (b *WorkbookFunctionsImLog10RequestBuilder) Request() *WorkbookFunctionsImLog10Request {
	return &WorkbookFunctionsImLog10Request{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImLog10Request) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImLog2RequestBuilder struct{ BaseRequestBuilder }

// ImLog2 action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImLog2(reqObj *WorkbookFunctionsImLog2RequestParameter) *WorkbookFunctionsImLog2RequestBuilder {
	bb := &WorkbookFunctionsImLog2RequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImLog2"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImLog2Request struct{ BaseRequest }

func (b *WorkbookFunctionsImLog2RequestBuilder) Request() *WorkbookFunctionsImLog2Request {
	return &WorkbookFunctionsImLog2Request{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImLog2Request) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImPowerRequestBuilder struct{ BaseRequestBuilder }

// ImPower action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImPower(reqObj *WorkbookFunctionsImPowerRequestParameter) *WorkbookFunctionsImPowerRequestBuilder {
	bb := &WorkbookFunctionsImPowerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImPower"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImPowerRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImPowerRequestBuilder) Request() *WorkbookFunctionsImPowerRequest {
	return &WorkbookFunctionsImPowerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImPowerRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImProductRequestBuilder struct{ BaseRequestBuilder }

// ImProduct action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImProduct(reqObj *WorkbookFunctionsImProductRequestParameter) *WorkbookFunctionsImProductRequestBuilder {
	bb := &WorkbookFunctionsImProductRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImProduct"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImProductRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImProductRequestBuilder) Request() *WorkbookFunctionsImProductRequest {
	return &WorkbookFunctionsImProductRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImProductRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImRealRequestBuilder struct{ BaseRequestBuilder }

// ImReal action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImReal(reqObj *WorkbookFunctionsImRealRequestParameter) *WorkbookFunctionsImRealRequestBuilder {
	bb := &WorkbookFunctionsImRealRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImReal"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImRealRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImRealRequestBuilder) Request() *WorkbookFunctionsImRealRequest {
	return &WorkbookFunctionsImRealRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImRealRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImSecRequestBuilder struct{ BaseRequestBuilder }

// ImSec action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImSec(reqObj *WorkbookFunctionsImSecRequestParameter) *WorkbookFunctionsImSecRequestBuilder {
	bb := &WorkbookFunctionsImSecRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImSec"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImSecRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImSecRequestBuilder) Request() *WorkbookFunctionsImSecRequest {
	return &WorkbookFunctionsImSecRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImSecRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImSechRequestBuilder struct{ BaseRequestBuilder }

// ImSech action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImSech(reqObj *WorkbookFunctionsImSechRequestParameter) *WorkbookFunctionsImSechRequestBuilder {
	bb := &WorkbookFunctionsImSechRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImSech"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImSechRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImSechRequestBuilder) Request() *WorkbookFunctionsImSechRequest {
	return &WorkbookFunctionsImSechRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImSechRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImSinRequestBuilder struct{ BaseRequestBuilder }

// ImSin action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImSin(reqObj *WorkbookFunctionsImSinRequestParameter) *WorkbookFunctionsImSinRequestBuilder {
	bb := &WorkbookFunctionsImSinRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImSin"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImSinRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImSinRequestBuilder) Request() *WorkbookFunctionsImSinRequest {
	return &WorkbookFunctionsImSinRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImSinRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImSinhRequestBuilder struct{ BaseRequestBuilder }

// ImSinh action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImSinh(reqObj *WorkbookFunctionsImSinhRequestParameter) *WorkbookFunctionsImSinhRequestBuilder {
	bb := &WorkbookFunctionsImSinhRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImSinh"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImSinhRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImSinhRequestBuilder) Request() *WorkbookFunctionsImSinhRequest {
	return &WorkbookFunctionsImSinhRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImSinhRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImSqrtRequestBuilder struct{ BaseRequestBuilder }

// ImSqrt action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImSqrt(reqObj *WorkbookFunctionsImSqrtRequestParameter) *WorkbookFunctionsImSqrtRequestBuilder {
	bb := &WorkbookFunctionsImSqrtRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImSqrt"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImSqrtRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImSqrtRequestBuilder) Request() *WorkbookFunctionsImSqrtRequest {
	return &WorkbookFunctionsImSqrtRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImSqrtRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImSubRequestBuilder struct{ BaseRequestBuilder }

// ImSub action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImSub(reqObj *WorkbookFunctionsImSubRequestParameter) *WorkbookFunctionsImSubRequestBuilder {
	bb := &WorkbookFunctionsImSubRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImSub"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImSubRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImSubRequestBuilder) Request() *WorkbookFunctionsImSubRequest {
	return &WorkbookFunctionsImSubRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImSubRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImSumRequestBuilder struct{ BaseRequestBuilder }

// ImSum action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImSum(reqObj *WorkbookFunctionsImSumRequestParameter) *WorkbookFunctionsImSumRequestBuilder {
	bb := &WorkbookFunctionsImSumRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImSum"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImSumRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImSumRequestBuilder) Request() *WorkbookFunctionsImSumRequest {
	return &WorkbookFunctionsImSumRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImSumRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type WorkbookFunctionsImTanRequestBuilder struct{ BaseRequestBuilder }

// ImTan action undocumented
func (b *WorkbookFunctionsRequestBuilder) ImTan(reqObj *WorkbookFunctionsImTanRequestParameter) *WorkbookFunctionsImTanRequestBuilder {
	bb := &WorkbookFunctionsImTanRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ImTan"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type WorkbookFunctionsImTanRequest struct{ BaseRequest }

func (b *WorkbookFunctionsImTanRequestBuilder) Request() *WorkbookFunctionsImTanRequest {
	return &WorkbookFunctionsImTanRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *WorkbookFunctionsImTanRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

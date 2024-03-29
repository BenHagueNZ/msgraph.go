// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SalesCreditMemoRequestBuilder is request builder for SalesCreditMemo
type SalesCreditMemoRequestBuilder struct{ BaseRequestBuilder }

// Request returns SalesCreditMemoRequest
func (b *SalesCreditMemoRequestBuilder) Request() *SalesCreditMemoRequest {
	return &SalesCreditMemoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SalesCreditMemoRequest is request for SalesCreditMemo
type SalesCreditMemoRequest struct{ BaseRequest }

// Get performs GET request for SalesCreditMemo
func (r *SalesCreditMemoRequest) Get(ctx context.Context) (resObj *SalesCreditMemo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SalesCreditMemo
func (r *SalesCreditMemoRequest) Update(ctx context.Context, reqObj *SalesCreditMemo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SalesCreditMemo
func (r *SalesCreditMemoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SalesCreditMemoLineRequestBuilder is request builder for SalesCreditMemoLine
type SalesCreditMemoLineRequestBuilder struct{ BaseRequestBuilder }

// Request returns SalesCreditMemoLineRequest
func (b *SalesCreditMemoLineRequestBuilder) Request() *SalesCreditMemoLineRequest {
	return &SalesCreditMemoLineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SalesCreditMemoLineRequest is request for SalesCreditMemoLine
type SalesCreditMemoLineRequest struct{ BaseRequest }

// Get performs GET request for SalesCreditMemoLine
func (r *SalesCreditMemoLineRequest) Get(ctx context.Context) (resObj *SalesCreditMemoLine, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SalesCreditMemoLine
func (r *SalesCreditMemoLineRequest) Update(ctx context.Context, reqObj *SalesCreditMemoLine) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SalesCreditMemoLine
func (r *SalesCreditMemoLineRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SalesInvoiceRequestBuilder is request builder for SalesInvoice
type SalesInvoiceRequestBuilder struct{ BaseRequestBuilder }

// Request returns SalesInvoiceRequest
func (b *SalesInvoiceRequestBuilder) Request() *SalesInvoiceRequest {
	return &SalesInvoiceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SalesInvoiceRequest is request for SalesInvoice
type SalesInvoiceRequest struct{ BaseRequest }

// Get performs GET request for SalesInvoice
func (r *SalesInvoiceRequest) Get(ctx context.Context) (resObj *SalesInvoice, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SalesInvoice
func (r *SalesInvoiceRequest) Update(ctx context.Context, reqObj *SalesInvoice) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SalesInvoice
func (r *SalesInvoiceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SalesInvoiceLineRequestBuilder is request builder for SalesInvoiceLine
type SalesInvoiceLineRequestBuilder struct{ BaseRequestBuilder }

// Request returns SalesInvoiceLineRequest
func (b *SalesInvoiceLineRequestBuilder) Request() *SalesInvoiceLineRequest {
	return &SalesInvoiceLineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SalesInvoiceLineRequest is request for SalesInvoiceLine
type SalesInvoiceLineRequest struct{ BaseRequest }

// Get performs GET request for SalesInvoiceLine
func (r *SalesInvoiceLineRequest) Get(ctx context.Context) (resObj *SalesInvoiceLine, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SalesInvoiceLine
func (r *SalesInvoiceLineRequest) Update(ctx context.Context, reqObj *SalesInvoiceLine) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SalesInvoiceLine
func (r *SalesInvoiceLineRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SalesOrderRequestBuilder is request builder for SalesOrder
type SalesOrderRequestBuilder struct{ BaseRequestBuilder }

// Request returns SalesOrderRequest
func (b *SalesOrderRequestBuilder) Request() *SalesOrderRequest {
	return &SalesOrderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SalesOrderRequest is request for SalesOrder
type SalesOrderRequest struct{ BaseRequest }

// Get performs GET request for SalesOrder
func (r *SalesOrderRequest) Get(ctx context.Context) (resObj *SalesOrder, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SalesOrder
func (r *SalesOrderRequest) Update(ctx context.Context, reqObj *SalesOrder) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SalesOrder
func (r *SalesOrderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SalesOrderLineRequestBuilder is request builder for SalesOrderLine
type SalesOrderLineRequestBuilder struct{ BaseRequestBuilder }

// Request returns SalesOrderLineRequest
func (b *SalesOrderLineRequestBuilder) Request() *SalesOrderLineRequest {
	return &SalesOrderLineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SalesOrderLineRequest is request for SalesOrderLine
type SalesOrderLineRequest struct{ BaseRequest }

// Get performs GET request for SalesOrderLine
func (r *SalesOrderLineRequest) Get(ctx context.Context) (resObj *SalesOrderLine, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SalesOrderLine
func (r *SalesOrderLineRequest) Update(ctx context.Context, reqObj *SalesOrderLine) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SalesOrderLine
func (r *SalesOrderLineRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SalesQuoteRequestBuilder is request builder for SalesQuote
type SalesQuoteRequestBuilder struct{ BaseRequestBuilder }

// Request returns SalesQuoteRequest
func (b *SalesQuoteRequestBuilder) Request() *SalesQuoteRequest {
	return &SalesQuoteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SalesQuoteRequest is request for SalesQuote
type SalesQuoteRequest struct{ BaseRequest }

// Get performs GET request for SalesQuote
func (r *SalesQuoteRequest) Get(ctx context.Context) (resObj *SalesQuote, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SalesQuote
func (r *SalesQuoteRequest) Update(ctx context.Context, reqObj *SalesQuote) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SalesQuote
func (r *SalesQuoteRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SalesQuoteLineRequestBuilder is request builder for SalesQuoteLine
type SalesQuoteLineRequestBuilder struct{ BaseRequestBuilder }

// Request returns SalesQuoteLineRequest
func (b *SalesQuoteLineRequestBuilder) Request() *SalesQuoteLineRequest {
	return &SalesQuoteLineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SalesQuoteLineRequest is request for SalesQuoteLine
type SalesQuoteLineRequest struct{ BaseRequest }

// Get performs GET request for SalesQuoteLine
func (r *SalesQuoteLineRequest) Get(ctx context.Context) (resObj *SalesQuoteLine, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SalesQuoteLine
func (r *SalesQuoteLineRequest) Update(ctx context.Context, reqObj *SalesQuoteLine) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SalesQuoteLine
func (r *SalesQuoteLineRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type SalesInvoiceCancelRequestBuilder struct{ BaseRequestBuilder }

// Cancel action undocumentedrav
func (b *SalesInvoiceRequestBuilder) Cancel(reqObj *SalesInvoiceCancelRequestParameter) *SalesInvoiceCancelRequestBuilder {
	bb := &SalesInvoiceCancelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Cancel"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SalesInvoiceCancelRequest struct{ BaseRequest }

func (b *SalesInvoiceCancelRequestBuilder) Request() *SalesInvoiceCancelRequest {
	return &SalesInvoiceCancelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SalesInvoiceCancelRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type SalesInvoiceSendRequestBuilder struct{ BaseRequestBuilder }

// Send action undocumentedrav
func (b *SalesInvoiceRequestBuilder) Send(reqObj *SalesInvoiceSendRequestParameter) *SalesInvoiceSendRequestBuilder {
	bb := &SalesInvoiceSendRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Send"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SalesInvoiceSendRequest struct{ BaseRequest }

func (b *SalesInvoiceSendRequestBuilder) Request() *SalesInvoiceSendRequest {
	return &SalesInvoiceSendRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SalesInvoiceSendRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type SalesInvoiceCancelAndSendRequestBuilder struct{ BaseRequestBuilder }

// CancelAndSend action undocumentedrav
func (b *SalesInvoiceRequestBuilder) CancelAndSend(reqObj *SalesInvoiceCancelAndSendRequestParameter) *SalesInvoiceCancelAndSendRequestBuilder {
	bb := &SalesInvoiceCancelAndSendRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CancelAndSend"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SalesInvoiceCancelAndSendRequest struct{ BaseRequest }

func (b *SalesInvoiceCancelAndSendRequestBuilder) Request() *SalesInvoiceCancelAndSendRequest {
	return &SalesInvoiceCancelAndSendRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SalesInvoiceCancelAndSendRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type SalesInvoicePostRequestBuilder struct{ BaseRequestBuilder }

// Post action undocumentedrav
func (b *SalesInvoiceRequestBuilder) Post(reqObj *SalesInvoicePostRequestParameter) *SalesInvoicePostRequestBuilder {
	bb := &SalesInvoicePostRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Post"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SalesInvoicePostRequest struct{ BaseRequest }

func (b *SalesInvoicePostRequestBuilder) Request() *SalesInvoicePostRequest {
	return &SalesInvoicePostRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SalesInvoicePostRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type SalesInvoicePostAndSendRequestBuilder struct{ BaseRequestBuilder }

// PostAndSend action undocumentedrav
func (b *SalesInvoiceRequestBuilder) PostAndSend(reqObj *SalesInvoicePostAndSendRequestParameter) *SalesInvoicePostAndSendRequestBuilder {
	bb := &SalesInvoicePostAndSendRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/PostAndSend"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SalesInvoicePostAndSendRequest struct{ BaseRequest }

func (b *SalesInvoicePostAndSendRequestBuilder) Request() *SalesInvoicePostAndSendRequest {
	return &SalesInvoicePostAndSendRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SalesInvoicePostAndSendRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type SalesQuoteSendRequestBuilder struct{ BaseRequestBuilder }

// Send action undocumentedrav
func (b *SalesQuoteRequestBuilder) Send(reqObj *SalesQuoteSendRequestParameter) *SalesQuoteSendRequestBuilder {
	bb := &SalesQuoteSendRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Send"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SalesQuoteSendRequest struct{ BaseRequest }

func (b *SalesQuoteSendRequestBuilder) Request() *SalesQuoteSendRequest {
	return &SalesQuoteSendRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SalesQuoteSendRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type SalesQuoteMakeInvoiceRequestBuilder struct{ BaseRequestBuilder }

// MakeInvoice action undocumentedrav
func (b *SalesQuoteRequestBuilder) MakeInvoice(reqObj *SalesQuoteMakeInvoiceRequestParameter) *SalesQuoteMakeInvoiceRequestBuilder {
	bb := &SalesQuoteMakeInvoiceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/MakeInvoice"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SalesQuoteMakeInvoiceRequest struct{ BaseRequest }

func (b *SalesQuoteMakeInvoiceRequestBuilder) Request() *SalesQuoteMakeInvoiceRequest {
	return &SalesQuoteMakeInvoiceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SalesQuoteMakeInvoiceRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

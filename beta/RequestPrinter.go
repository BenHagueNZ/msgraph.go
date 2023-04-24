// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PrinterRequestBuilder is request builder for Printer
type PrinterRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrinterRequest
func (b *PrinterRequestBuilder) Request() *PrinterRequest {
	return &PrinterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrinterRequest is request for Printer
type PrinterRequest struct{ BaseRequest }

// Get performs GET request for Printer
func (r *PrinterRequest) Get(ctx context.Context) (resObj *Printer, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Printer
func (r *PrinterRequest) Update(ctx context.Context, reqObj *Printer) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Printer
func (r *PrinterRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrinterBaseRequestBuilder is request builder for PrinterBase
type PrinterBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrinterBaseRequest
func (b *PrinterBaseRequestBuilder) Request() *PrinterBaseRequest {
	return &PrinterBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrinterBaseRequest is request for PrinterBase
type PrinterBaseRequest struct{ BaseRequest }

// Get performs GET request for PrinterBase
func (r *PrinterBaseRequest) Get(ctx context.Context) (resObj *PrinterBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrinterBase
func (r *PrinterBaseRequest) Update(ctx context.Context, reqObj *PrinterBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrinterBase
func (r *PrinterBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrinterCapabilitiesRequestBuilder is request builder for PrinterCapabilities
type PrinterCapabilitiesRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrinterCapabilitiesRequest
func (b *PrinterCapabilitiesRequestBuilder) Request() *PrinterCapabilitiesRequest {
	return &PrinterCapabilitiesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrinterCapabilitiesRequest is request for PrinterCapabilities
type PrinterCapabilitiesRequest struct{ BaseRequest }

// Get performs GET request for PrinterCapabilities
func (r *PrinterCapabilitiesRequest) Get(ctx context.Context) (resObj *PrinterCapabilities, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrinterCapabilities
func (r *PrinterCapabilitiesRequest) Update(ctx context.Context, reqObj *PrinterCapabilities) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrinterCapabilities
func (r *PrinterCapabilitiesRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrinterCreateOperationRequestBuilder is request builder for PrinterCreateOperation
type PrinterCreateOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrinterCreateOperationRequest
func (b *PrinterCreateOperationRequestBuilder) Request() *PrinterCreateOperationRequest {
	return &PrinterCreateOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrinterCreateOperationRequest is request for PrinterCreateOperation
type PrinterCreateOperationRequest struct{ BaseRequest }

// Get performs GET request for PrinterCreateOperation
func (r *PrinterCreateOperationRequest) Get(ctx context.Context) (resObj *PrinterCreateOperation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrinterCreateOperation
func (r *PrinterCreateOperationRequest) Update(ctx context.Context, reqObj *PrinterCreateOperation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrinterCreateOperation
func (r *PrinterCreateOperationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrinterDefaultsRequestBuilder is request builder for PrinterDefaults
type PrinterDefaultsRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrinterDefaultsRequest
func (b *PrinterDefaultsRequestBuilder) Request() *PrinterDefaultsRequest {
	return &PrinterDefaultsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrinterDefaultsRequest is request for PrinterDefaults
type PrinterDefaultsRequest struct{ BaseRequest }

// Get performs GET request for PrinterDefaults
func (r *PrinterDefaultsRequest) Get(ctx context.Context) (resObj *PrinterDefaults, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrinterDefaults
func (r *PrinterDefaultsRequest) Update(ctx context.Context, reqObj *PrinterDefaults) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrinterDefaults
func (r *PrinterDefaultsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrinterDocumentConfigurationRequestBuilder is request builder for PrinterDocumentConfiguration
type PrinterDocumentConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrinterDocumentConfigurationRequest
func (b *PrinterDocumentConfigurationRequestBuilder) Request() *PrinterDocumentConfigurationRequest {
	return &PrinterDocumentConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrinterDocumentConfigurationRequest is request for PrinterDocumentConfiguration
type PrinterDocumentConfigurationRequest struct{ BaseRequest }

// Get performs GET request for PrinterDocumentConfiguration
func (r *PrinterDocumentConfigurationRequest) Get(ctx context.Context) (resObj *PrinterDocumentConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrinterDocumentConfiguration
func (r *PrinterDocumentConfigurationRequest) Update(ctx context.Context, reqObj *PrinterDocumentConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrinterDocumentConfiguration
func (r *PrinterDocumentConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrinterLocationRequestBuilder is request builder for PrinterLocation
type PrinterLocationRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrinterLocationRequest
func (b *PrinterLocationRequestBuilder) Request() *PrinterLocationRequest {
	return &PrinterLocationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrinterLocationRequest is request for PrinterLocation
type PrinterLocationRequest struct{ BaseRequest }

// Get performs GET request for PrinterLocation
func (r *PrinterLocationRequest) Get(ctx context.Context) (resObj *PrinterLocation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrinterLocation
func (r *PrinterLocationRequest) Update(ctx context.Context, reqObj *PrinterLocation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrinterLocation
func (r *PrinterLocationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrinterShareRequestBuilder is request builder for PrinterShare
type PrinterShareRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrinterShareRequest
func (b *PrinterShareRequestBuilder) Request() *PrinterShareRequest {
	return &PrinterShareRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrinterShareRequest is request for PrinterShare
type PrinterShareRequest struct{ BaseRequest }

// Get performs GET request for PrinterShare
func (r *PrinterShareRequest) Get(ctx context.Context) (resObj *PrinterShare, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrinterShare
func (r *PrinterShareRequest) Update(ctx context.Context, reqObj *PrinterShare) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrinterShare
func (r *PrinterShareRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrinterShareViewpointRequestBuilder is request builder for PrinterShareViewpoint
type PrinterShareViewpointRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrinterShareViewpointRequest
func (b *PrinterShareViewpointRequestBuilder) Request() *PrinterShareViewpointRequest {
	return &PrinterShareViewpointRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrinterShareViewpointRequest is request for PrinterShareViewpoint
type PrinterShareViewpointRequest struct{ BaseRequest }

// Get performs GET request for PrinterShareViewpoint
func (r *PrinterShareViewpointRequest) Get(ctx context.Context) (resObj *PrinterShareViewpoint, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrinterShareViewpoint
func (r *PrinterShareViewpointRequest) Update(ctx context.Context, reqObj *PrinterShareViewpoint) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrinterShareViewpoint
func (r *PrinterShareViewpointRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrinterStatusRequestBuilder is request builder for PrinterStatus
type PrinterStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrinterStatusRequest
func (b *PrinterStatusRequestBuilder) Request() *PrinterStatusRequest {
	return &PrinterStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrinterStatusRequest is request for PrinterStatus
type PrinterStatusRequest struct{ BaseRequest }

// Get performs GET request for PrinterStatus
func (r *PrinterStatusRequest) Get(ctx context.Context) (resObj *PrinterStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrinterStatus
func (r *PrinterStatusRequest) Update(ctx context.Context, reqObj *PrinterStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrinterStatus
func (r *PrinterStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrinterUsageSummaryRequestBuilder is request builder for PrinterUsageSummary
type PrinterUsageSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrinterUsageSummaryRequest
func (b *PrinterUsageSummaryRequestBuilder) Request() *PrinterUsageSummaryRequest {
	return &PrinterUsageSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrinterUsageSummaryRequest is request for PrinterUsageSummary
type PrinterUsageSummaryRequest struct{ BaseRequest }

// Get performs GET request for PrinterUsageSummary
func (r *PrinterUsageSummaryRequest) Get(ctx context.Context) (resObj *PrinterUsageSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrinterUsageSummary
func (r *PrinterUsageSummaryRequest) Update(ctx context.Context, reqObj *PrinterUsageSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrinterUsageSummary
func (r *PrinterUsageSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type PrinterCollectionCreateRequestBuilder struct{ BaseRequestBuilder }

// Create action undocumentedrav
func (b *PrintPrintersCollectionRequestBuilder) Create(reqObj *PrinterCollectionCreateRequestParameter) *PrinterCollectionCreateRequestBuilder {
	bb := &PrinterCollectionCreateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Create"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type PrinterCollectionCreateRequest struct{ BaseRequest }

func (b *PrinterCollectionCreateRequestBuilder) Request() *PrinterCollectionCreateRequest {
	return &PrinterCollectionCreateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *PrinterCollectionCreateRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type PrinterResetDefaultsRequestBuilder struct{ BaseRequestBuilder }

// ResetDefaults action undocumentedrav
func (b *PrinterRequestBuilder) ResetDefaults(reqObj *PrinterResetDefaultsRequestParameter) *PrinterResetDefaultsRequestBuilder {
	bb := &PrinterResetDefaultsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ResetDefaults"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type PrinterResetDefaultsRequest struct{ BaseRequest }

func (b *PrinterResetDefaultsRequestBuilder) Request() *PrinterResetDefaultsRequest {
	return &PrinterResetDefaultsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *PrinterResetDefaultsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type PrinterRestoreFactoryDefaultsRequestBuilder struct{ BaseRequestBuilder }

// RestoreFactoryDefaults action undocumentedrav
func (b *PrinterRequestBuilder) RestoreFactoryDefaults(reqObj *PrinterRestoreFactoryDefaultsRequestParameter) *PrinterRestoreFactoryDefaultsRequestBuilder {
	bb := &PrinterRestoreFactoryDefaultsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RestoreFactoryDefaults"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type PrinterRestoreFactoryDefaultsRequest struct{ BaseRequest }

func (b *PrinterRestoreFactoryDefaultsRequestBuilder) Request() *PrinterRestoreFactoryDefaultsRequest {
	return &PrinterRestoreFactoryDefaultsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *PrinterRestoreFactoryDefaultsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PrintRequestBuilder is request builder for Print
type PrintRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintRequest
func (b *PrintRequestBuilder) Request() *PrintRequest {
	return &PrintRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintRequest is request for Print
type PrintRequest struct{ BaseRequest }

// Get performs GET request for Print
func (r *PrintRequest) Get(ctx context.Context) (resObj *Print, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Print
func (r *PrintRequest) Update(ctx context.Context, reqObj *Print) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Print
func (r *PrintRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintCertificateSigningRequestObjectRequestBuilder is request builder for PrintCertificateSigningRequestObject
type PrintCertificateSigningRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintCertificateSigningRequestObjectRequest
func (b *PrintCertificateSigningRequestObjectRequestBuilder) Request() *PrintCertificateSigningRequestObjectRequest {
	return &PrintCertificateSigningRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintCertificateSigningRequestObjectRequest is request for PrintCertificateSigningRequestObject
type PrintCertificateSigningRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for PrintCertificateSigningRequestObject
func (r *PrintCertificateSigningRequestObjectRequest) Get(ctx context.Context) (resObj *PrintCertificateSigningRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintCertificateSigningRequestObject
func (r *PrintCertificateSigningRequestObjectRequest) Update(ctx context.Context, reqObj *PrintCertificateSigningRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintCertificateSigningRequestObject
func (r *PrintCertificateSigningRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintConnectorRequestBuilder is request builder for PrintConnector
type PrintConnectorRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintConnectorRequest
func (b *PrintConnectorRequestBuilder) Request() *PrintConnectorRequest {
	return &PrintConnectorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintConnectorRequest is request for PrintConnector
type PrintConnectorRequest struct{ BaseRequest }

// Get performs GET request for PrintConnector
func (r *PrintConnectorRequest) Get(ctx context.Context) (resObj *PrintConnector, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintConnector
func (r *PrintConnectorRequest) Update(ctx context.Context, reqObj *PrintConnector) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintConnector
func (r *PrintConnectorRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintDocumentRequestBuilder is request builder for PrintDocument
type PrintDocumentRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintDocumentRequest
func (b *PrintDocumentRequestBuilder) Request() *PrintDocumentRequest {
	return &PrintDocumentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintDocumentRequest is request for PrintDocument
type PrintDocumentRequest struct{ BaseRequest }

// Get performs GET request for PrintDocument
func (r *PrintDocumentRequest) Get(ctx context.Context) (resObj *PrintDocument, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintDocument
func (r *PrintDocumentRequest) Update(ctx context.Context, reqObj *PrintDocument) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintDocument
func (r *PrintDocumentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintDocumentUploadPropertiesRequestBuilder is request builder for PrintDocumentUploadProperties
type PrintDocumentUploadPropertiesRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintDocumentUploadPropertiesRequest
func (b *PrintDocumentUploadPropertiesRequestBuilder) Request() *PrintDocumentUploadPropertiesRequest {
	return &PrintDocumentUploadPropertiesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintDocumentUploadPropertiesRequest is request for PrintDocumentUploadProperties
type PrintDocumentUploadPropertiesRequest struct{ BaseRequest }

// Get performs GET request for PrintDocumentUploadProperties
func (r *PrintDocumentUploadPropertiesRequest) Get(ctx context.Context) (resObj *PrintDocumentUploadProperties, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintDocumentUploadProperties
func (r *PrintDocumentUploadPropertiesRequest) Update(ctx context.Context, reqObj *PrintDocumentUploadProperties) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintDocumentUploadProperties
func (r *PrintDocumentUploadPropertiesRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintJobRequestBuilder is request builder for PrintJob
type PrintJobRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintJobRequest
func (b *PrintJobRequestBuilder) Request() *PrintJobRequest {
	return &PrintJobRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintJobRequest is request for PrintJob
type PrintJobRequest struct{ BaseRequest }

// Get performs GET request for PrintJob
func (r *PrintJobRequest) Get(ctx context.Context) (resObj *PrintJob, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintJob
func (r *PrintJobRequest) Update(ctx context.Context, reqObj *PrintJob) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintJob
func (r *PrintJobRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintJobConfigurationRequestBuilder is request builder for PrintJobConfiguration
type PrintJobConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintJobConfigurationRequest
func (b *PrintJobConfigurationRequestBuilder) Request() *PrintJobConfigurationRequest {
	return &PrintJobConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintJobConfigurationRequest is request for PrintJobConfiguration
type PrintJobConfigurationRequest struct{ BaseRequest }

// Get performs GET request for PrintJobConfiguration
func (r *PrintJobConfigurationRequest) Get(ctx context.Context) (resObj *PrintJobConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintJobConfiguration
func (r *PrintJobConfigurationRequest) Update(ctx context.Context, reqObj *PrintJobConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintJobConfiguration
func (r *PrintJobConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintJobStatusRequestBuilder is request builder for PrintJobStatus
type PrintJobStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintJobStatusRequest
func (b *PrintJobStatusRequestBuilder) Request() *PrintJobStatusRequest {
	return &PrintJobStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintJobStatusRequest is request for PrintJobStatus
type PrintJobStatusRequest struct{ BaseRequest }

// Get performs GET request for PrintJobStatus
func (r *PrintJobStatusRequest) Get(ctx context.Context) (resObj *PrintJobStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintJobStatus
func (r *PrintJobStatusRequest) Update(ctx context.Context, reqObj *PrintJobStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintJobStatus
func (r *PrintJobStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintMarginRequestBuilder is request builder for PrintMargin
type PrintMarginRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintMarginRequest
func (b *PrintMarginRequestBuilder) Request() *PrintMarginRequest {
	return &PrintMarginRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintMarginRequest is request for PrintMargin
type PrintMarginRequest struct{ BaseRequest }

// Get performs GET request for PrintMargin
func (r *PrintMarginRequest) Get(ctx context.Context) (resObj *PrintMargin, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintMargin
func (r *PrintMarginRequest) Update(ctx context.Context, reqObj *PrintMargin) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintMargin
func (r *PrintMarginRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintOperationRequestBuilder is request builder for PrintOperation
type PrintOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintOperationRequest
func (b *PrintOperationRequestBuilder) Request() *PrintOperationRequest {
	return &PrintOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintOperationRequest is request for PrintOperation
type PrintOperationRequest struct{ BaseRequest }

// Get performs GET request for PrintOperation
func (r *PrintOperationRequest) Get(ctx context.Context) (resObj *PrintOperation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintOperation
func (r *PrintOperationRequest) Update(ctx context.Context, reqObj *PrintOperation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintOperation
func (r *PrintOperationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintOperationStatusRequestBuilder is request builder for PrintOperationStatus
type PrintOperationStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintOperationStatusRequest
func (b *PrintOperationStatusRequestBuilder) Request() *PrintOperationStatusRequest {
	return &PrintOperationStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintOperationStatusRequest is request for PrintOperationStatus
type PrintOperationStatusRequest struct{ BaseRequest }

// Get performs GET request for PrintOperationStatus
func (r *PrintOperationStatusRequest) Get(ctx context.Context) (resObj *PrintOperationStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintOperationStatus
func (r *PrintOperationStatusRequest) Update(ctx context.Context, reqObj *PrintOperationStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintOperationStatus
func (r *PrintOperationStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintServiceRequestBuilder is request builder for PrintService
type PrintServiceRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintServiceRequest
func (b *PrintServiceRequestBuilder) Request() *PrintServiceRequest {
	return &PrintServiceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintServiceRequest is request for PrintService
type PrintServiceRequest struct{ BaseRequest }

// Get performs GET request for PrintService
func (r *PrintServiceRequest) Get(ctx context.Context) (resObj *PrintService, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintService
func (r *PrintServiceRequest) Update(ctx context.Context, reqObj *PrintService) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintService
func (r *PrintServiceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintServiceEndpointRequestBuilder is request builder for PrintServiceEndpoint
type PrintServiceEndpointRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintServiceEndpointRequest
func (b *PrintServiceEndpointRequestBuilder) Request() *PrintServiceEndpointRequest {
	return &PrintServiceEndpointRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintServiceEndpointRequest is request for PrintServiceEndpoint
type PrintServiceEndpointRequest struct{ BaseRequest }

// Get performs GET request for PrintServiceEndpoint
func (r *PrintServiceEndpointRequest) Get(ctx context.Context) (resObj *PrintServiceEndpoint, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintServiceEndpoint
func (r *PrintServiceEndpointRequest) Update(ctx context.Context, reqObj *PrintServiceEndpoint) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintServiceEndpoint
func (r *PrintServiceEndpointRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintSettingsRequestBuilder is request builder for PrintSettings
type PrintSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintSettingsRequest
func (b *PrintSettingsRequestBuilder) Request() *PrintSettingsRequest {
	return &PrintSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintSettingsRequest is request for PrintSettings
type PrintSettingsRequest struct{ BaseRequest }

// Get performs GET request for PrintSettings
func (r *PrintSettingsRequest) Get(ctx context.Context) (resObj *PrintSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintSettings
func (r *PrintSettingsRequest) Update(ctx context.Context, reqObj *PrintSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintSettings
func (r *PrintSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintTaskRequestBuilder is request builder for PrintTask
type PrintTaskRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintTaskRequest
func (b *PrintTaskRequestBuilder) Request() *PrintTaskRequest {
	return &PrintTaskRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintTaskRequest is request for PrintTask
type PrintTaskRequest struct{ BaseRequest }

// Get performs GET request for PrintTask
func (r *PrintTaskRequest) Get(ctx context.Context) (resObj *PrintTask, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintTask
func (r *PrintTaskRequest) Update(ctx context.Context, reqObj *PrintTask) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintTask
func (r *PrintTaskRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintTaskDefinitionRequestBuilder is request builder for PrintTaskDefinition
type PrintTaskDefinitionRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintTaskDefinitionRequest
func (b *PrintTaskDefinitionRequestBuilder) Request() *PrintTaskDefinitionRequest {
	return &PrintTaskDefinitionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintTaskDefinitionRequest is request for PrintTaskDefinition
type PrintTaskDefinitionRequest struct{ BaseRequest }

// Get performs GET request for PrintTaskDefinition
func (r *PrintTaskDefinitionRequest) Get(ctx context.Context) (resObj *PrintTaskDefinition, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintTaskDefinition
func (r *PrintTaskDefinitionRequest) Update(ctx context.Context, reqObj *PrintTaskDefinition) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintTaskDefinition
func (r *PrintTaskDefinitionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintTaskStatusRequestBuilder is request builder for PrintTaskStatus
type PrintTaskStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintTaskStatusRequest
func (b *PrintTaskStatusRequestBuilder) Request() *PrintTaskStatusRequest {
	return &PrintTaskStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintTaskStatusRequest is request for PrintTaskStatus
type PrintTaskStatusRequest struct{ BaseRequest }

// Get performs GET request for PrintTaskStatus
func (r *PrintTaskStatusRequest) Get(ctx context.Context) (resObj *PrintTaskStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintTaskStatus
func (r *PrintTaskStatusRequest) Update(ctx context.Context, reqObj *PrintTaskStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintTaskStatus
func (r *PrintTaskStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintTaskTriggerRequestBuilder is request builder for PrintTaskTrigger
type PrintTaskTriggerRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintTaskTriggerRequest
func (b *PrintTaskTriggerRequestBuilder) Request() *PrintTaskTriggerRequest {
	return &PrintTaskTriggerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintTaskTriggerRequest is request for PrintTaskTrigger
type PrintTaskTriggerRequest struct{ BaseRequest }

// Get performs GET request for PrintTaskTrigger
func (r *PrintTaskTriggerRequest) Get(ctx context.Context) (resObj *PrintTaskTrigger, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintTaskTrigger
func (r *PrintTaskTriggerRequest) Update(ctx context.Context, reqObj *PrintTaskTrigger) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintTaskTrigger
func (r *PrintTaskTriggerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintUsageRequestBuilder is request builder for PrintUsage
type PrintUsageRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintUsageRequest
func (b *PrintUsageRequestBuilder) Request() *PrintUsageRequest {
	return &PrintUsageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintUsageRequest is request for PrintUsage
type PrintUsageRequest struct{ BaseRequest }

// Get performs GET request for PrintUsage
func (r *PrintUsageRequest) Get(ctx context.Context) (resObj *PrintUsage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintUsage
func (r *PrintUsageRequest) Update(ctx context.Context, reqObj *PrintUsage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintUsage
func (r *PrintUsageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintUsageByPrinterRequestBuilder is request builder for PrintUsageByPrinter
type PrintUsageByPrinterRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintUsageByPrinterRequest
func (b *PrintUsageByPrinterRequestBuilder) Request() *PrintUsageByPrinterRequest {
	return &PrintUsageByPrinterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintUsageByPrinterRequest is request for PrintUsageByPrinter
type PrintUsageByPrinterRequest struct{ BaseRequest }

// Get performs GET request for PrintUsageByPrinter
func (r *PrintUsageByPrinterRequest) Get(ctx context.Context) (resObj *PrintUsageByPrinter, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintUsageByPrinter
func (r *PrintUsageByPrinterRequest) Update(ctx context.Context, reqObj *PrintUsageByPrinter) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintUsageByPrinter
func (r *PrintUsageByPrinterRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrintUsageByUserRequestBuilder is request builder for PrintUsageByUser
type PrintUsageByUserRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrintUsageByUserRequest
func (b *PrintUsageByUserRequestBuilder) Request() *PrintUsageByUserRequest {
	return &PrintUsageByUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrintUsageByUserRequest is request for PrintUsageByUser
type PrintUsageByUserRequest struct{ BaseRequest }

// Get performs GET request for PrintUsageByUser
func (r *PrintUsageByUserRequest) Get(ctx context.Context) (resObj *PrintUsageByUser, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrintUsageByUser
func (r *PrintUsageByUserRequest) Update(ctx context.Context, reqObj *PrintUsageByUser) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrintUsageByUser
func (r *PrintUsageByUserRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
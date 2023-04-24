// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SynchronizationRequestBuilder is request builder for Synchronization
type SynchronizationRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationRequest
func (b *SynchronizationRequestBuilder) Request() *SynchronizationRequest {
	return &SynchronizationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationRequest is request for Synchronization
type SynchronizationRequest struct{ BaseRequest }

// Get performs GET request for Synchronization
func (r *SynchronizationRequest) Get(ctx context.Context) (resObj *Synchronization, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Synchronization
func (r *SynchronizationRequest) Update(ctx context.Context, reqObj *Synchronization) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Synchronization
func (r *SynchronizationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SynchronizationErrorRequestBuilder is request builder for SynchronizationError
type SynchronizationErrorRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationErrorRequest
func (b *SynchronizationErrorRequestBuilder) Request() *SynchronizationErrorRequest {
	return &SynchronizationErrorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationErrorRequest is request for SynchronizationError
type SynchronizationErrorRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationError
func (r *SynchronizationErrorRequest) Get(ctx context.Context) (resObj *SynchronizationError, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationError
func (r *SynchronizationErrorRequest) Update(ctx context.Context, reqObj *SynchronizationError) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationError
func (r *SynchronizationErrorRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SynchronizationJobRequestBuilder is request builder for SynchronizationJob
type SynchronizationJobRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationJobRequest
func (b *SynchronizationJobRequestBuilder) Request() *SynchronizationJobRequest {
	return &SynchronizationJobRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationJobRequest is request for SynchronizationJob
type SynchronizationJobRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationJob
func (r *SynchronizationJobRequest) Get(ctx context.Context) (resObj *SynchronizationJob, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationJob
func (r *SynchronizationJobRequest) Update(ctx context.Context, reqObj *SynchronizationJob) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationJob
func (r *SynchronizationJobRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SynchronizationJobApplicationParametersRequestBuilder is request builder for SynchronizationJobApplicationParameters
type SynchronizationJobApplicationParametersRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationJobApplicationParametersRequest
func (b *SynchronizationJobApplicationParametersRequestBuilder) Request() *SynchronizationJobApplicationParametersRequest {
	return &SynchronizationJobApplicationParametersRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationJobApplicationParametersRequest is request for SynchronizationJobApplicationParameters
type SynchronizationJobApplicationParametersRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationJobApplicationParameters
func (r *SynchronizationJobApplicationParametersRequest) Get(ctx context.Context) (resObj *SynchronizationJobApplicationParameters, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationJobApplicationParameters
func (r *SynchronizationJobApplicationParametersRequest) Update(ctx context.Context, reqObj *SynchronizationJobApplicationParameters) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationJobApplicationParameters
func (r *SynchronizationJobApplicationParametersRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SynchronizationJobRestartCriteriaRequestBuilder is request builder for SynchronizationJobRestartCriteria
type SynchronizationJobRestartCriteriaRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationJobRestartCriteriaRequest
func (b *SynchronizationJobRestartCriteriaRequestBuilder) Request() *SynchronizationJobRestartCriteriaRequest {
	return &SynchronizationJobRestartCriteriaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationJobRestartCriteriaRequest is request for SynchronizationJobRestartCriteria
type SynchronizationJobRestartCriteriaRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationJobRestartCriteria
func (r *SynchronizationJobRestartCriteriaRequest) Get(ctx context.Context) (resObj *SynchronizationJobRestartCriteria, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationJobRestartCriteria
func (r *SynchronizationJobRestartCriteriaRequest) Update(ctx context.Context, reqObj *SynchronizationJobRestartCriteria) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationJobRestartCriteria
func (r *SynchronizationJobRestartCriteriaRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SynchronizationJobSubjectRequestBuilder is request builder for SynchronizationJobSubject
type SynchronizationJobSubjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationJobSubjectRequest
func (b *SynchronizationJobSubjectRequestBuilder) Request() *SynchronizationJobSubjectRequest {
	return &SynchronizationJobSubjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationJobSubjectRequest is request for SynchronizationJobSubject
type SynchronizationJobSubjectRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationJobSubject
func (r *SynchronizationJobSubjectRequest) Get(ctx context.Context) (resObj *SynchronizationJobSubject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationJobSubject
func (r *SynchronizationJobSubjectRequest) Update(ctx context.Context, reqObj *SynchronizationJobSubject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationJobSubject
func (r *SynchronizationJobSubjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SynchronizationLinkedObjectsRequestBuilder is request builder for SynchronizationLinkedObjects
type SynchronizationLinkedObjectsRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationLinkedObjectsRequest
func (b *SynchronizationLinkedObjectsRequestBuilder) Request() *SynchronizationLinkedObjectsRequest {
	return &SynchronizationLinkedObjectsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationLinkedObjectsRequest is request for SynchronizationLinkedObjects
type SynchronizationLinkedObjectsRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationLinkedObjects
func (r *SynchronizationLinkedObjectsRequest) Get(ctx context.Context) (resObj *SynchronizationLinkedObjects, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationLinkedObjects
func (r *SynchronizationLinkedObjectsRequest) Update(ctx context.Context, reqObj *SynchronizationLinkedObjects) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationLinkedObjects
func (r *SynchronizationLinkedObjectsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SynchronizationProgressRequestBuilder is request builder for SynchronizationProgress
type SynchronizationProgressRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationProgressRequest
func (b *SynchronizationProgressRequestBuilder) Request() *SynchronizationProgressRequest {
	return &SynchronizationProgressRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationProgressRequest is request for SynchronizationProgress
type SynchronizationProgressRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationProgress
func (r *SynchronizationProgressRequest) Get(ctx context.Context) (resObj *SynchronizationProgress, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationProgress
func (r *SynchronizationProgressRequest) Update(ctx context.Context, reqObj *SynchronizationProgress) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationProgress
func (r *SynchronizationProgressRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SynchronizationQuarantineRequestBuilder is request builder for SynchronizationQuarantine
type SynchronizationQuarantineRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationQuarantineRequest
func (b *SynchronizationQuarantineRequestBuilder) Request() *SynchronizationQuarantineRequest {
	return &SynchronizationQuarantineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationQuarantineRequest is request for SynchronizationQuarantine
type SynchronizationQuarantineRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationQuarantine
func (r *SynchronizationQuarantineRequest) Get(ctx context.Context) (resObj *SynchronizationQuarantine, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationQuarantine
func (r *SynchronizationQuarantineRequest) Update(ctx context.Context, reqObj *SynchronizationQuarantine) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationQuarantine
func (r *SynchronizationQuarantineRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SynchronizationRuleRequestBuilder is request builder for SynchronizationRule
type SynchronizationRuleRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationRuleRequest
func (b *SynchronizationRuleRequestBuilder) Request() *SynchronizationRuleRequest {
	return &SynchronizationRuleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationRuleRequest is request for SynchronizationRule
type SynchronizationRuleRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationRule
func (r *SynchronizationRuleRequest) Get(ctx context.Context) (resObj *SynchronizationRule, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationRule
func (r *SynchronizationRuleRequest) Update(ctx context.Context, reqObj *SynchronizationRule) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationRule
func (r *SynchronizationRuleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SynchronizationScheduleRequestBuilder is request builder for SynchronizationSchedule
type SynchronizationScheduleRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationScheduleRequest
func (b *SynchronizationScheduleRequestBuilder) Request() *SynchronizationScheduleRequest {
	return &SynchronizationScheduleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationScheduleRequest is request for SynchronizationSchedule
type SynchronizationScheduleRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationSchedule
func (r *SynchronizationScheduleRequest) Get(ctx context.Context) (resObj *SynchronizationSchedule, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationSchedule
func (r *SynchronizationScheduleRequest) Update(ctx context.Context, reqObj *SynchronizationSchedule) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationSchedule
func (r *SynchronizationScheduleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SynchronizationSchemaRequestBuilder is request builder for SynchronizationSchema
type SynchronizationSchemaRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationSchemaRequest
func (b *SynchronizationSchemaRequestBuilder) Request() *SynchronizationSchemaRequest {
	return &SynchronizationSchemaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationSchemaRequest is request for SynchronizationSchema
type SynchronizationSchemaRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationSchema
func (r *SynchronizationSchemaRequest) Get(ctx context.Context) (resObj *SynchronizationSchema, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationSchema
func (r *SynchronizationSchemaRequest) Update(ctx context.Context, reqObj *SynchronizationSchema) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationSchema
func (r *SynchronizationSchemaRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SynchronizationSecretKeyStringValuePairRequestBuilder is request builder for SynchronizationSecretKeyStringValuePair
type SynchronizationSecretKeyStringValuePairRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationSecretKeyStringValuePairRequest
func (b *SynchronizationSecretKeyStringValuePairRequestBuilder) Request() *SynchronizationSecretKeyStringValuePairRequest {
	return &SynchronizationSecretKeyStringValuePairRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationSecretKeyStringValuePairRequest is request for SynchronizationSecretKeyStringValuePair
type SynchronizationSecretKeyStringValuePairRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationSecretKeyStringValuePair
func (r *SynchronizationSecretKeyStringValuePairRequest) Get(ctx context.Context) (resObj *SynchronizationSecretKeyStringValuePair, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationSecretKeyStringValuePair
func (r *SynchronizationSecretKeyStringValuePairRequest) Update(ctx context.Context, reqObj *SynchronizationSecretKeyStringValuePair) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationSecretKeyStringValuePair
func (r *SynchronizationSecretKeyStringValuePairRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SynchronizationStatusRequestBuilder is request builder for SynchronizationStatus
type SynchronizationStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationStatusRequest
func (b *SynchronizationStatusRequestBuilder) Request() *SynchronizationStatusRequest {
	return &SynchronizationStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationStatusRequest is request for SynchronizationStatus
type SynchronizationStatusRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationStatus
func (r *SynchronizationStatusRequest) Get(ctx context.Context) (resObj *SynchronizationStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationStatus
func (r *SynchronizationStatusRequest) Update(ctx context.Context, reqObj *SynchronizationStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationStatus
func (r *SynchronizationStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SynchronizationTaskExecutionRequestBuilder is request builder for SynchronizationTaskExecution
type SynchronizationTaskExecutionRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationTaskExecutionRequest
func (b *SynchronizationTaskExecutionRequestBuilder) Request() *SynchronizationTaskExecutionRequest {
	return &SynchronizationTaskExecutionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationTaskExecutionRequest is request for SynchronizationTaskExecution
type SynchronizationTaskExecutionRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationTaskExecution
func (r *SynchronizationTaskExecutionRequest) Get(ctx context.Context) (resObj *SynchronizationTaskExecution, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationTaskExecution
func (r *SynchronizationTaskExecutionRequest) Update(ctx context.Context, reqObj *SynchronizationTaskExecution) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationTaskExecution
func (r *SynchronizationTaskExecutionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SynchronizationTemplateRequestBuilder is request builder for SynchronizationTemplate
type SynchronizationTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationTemplateRequest
func (b *SynchronizationTemplateRequestBuilder) Request() *SynchronizationTemplateRequest {
	return &SynchronizationTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationTemplateRequest is request for SynchronizationTemplate
type SynchronizationTemplateRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationTemplate
func (r *SynchronizationTemplateRequest) Get(ctx context.Context) (resObj *SynchronizationTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationTemplate
func (r *SynchronizationTemplateRequest) Update(ctx context.Context, reqObj *SynchronizationTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationTemplate
func (r *SynchronizationTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type SynchronizationJobCollectionValidateCredentialsRequestBuilder struct{ BaseRequestBuilder }

// ValidateCredentials action undocumentedrav
func (b *SynchronizationJobsCollectionRequestBuilder) ValidateCredentials(reqObj *SynchronizationJobCollectionValidateCredentialsRequestParameter) *SynchronizationJobCollectionValidateCredentialsRequestBuilder {
	bb := &SynchronizationJobCollectionValidateCredentialsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ValidateCredentials"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SynchronizationJobCollectionValidateCredentialsRequest struct{ BaseRequest }

func (b *SynchronizationJobCollectionValidateCredentialsRequestBuilder) Request() *SynchronizationJobCollectionValidateCredentialsRequest {
	return &SynchronizationJobCollectionValidateCredentialsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SynchronizationJobCollectionValidateCredentialsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type SynchronizationAcquireAccessTokenRequestBuilder struct{ BaseRequestBuilder }

// AcquireAccessToken action undocumentedrav
func (b *SynchronizationRequestBuilder) AcquireAccessToken(reqObj *SynchronizationAcquireAccessTokenRequestParameter) *SynchronizationAcquireAccessTokenRequestBuilder {
	bb := &SynchronizationAcquireAccessTokenRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AcquireAccessToken"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SynchronizationAcquireAccessTokenRequest struct{ BaseRequest }

func (b *SynchronizationAcquireAccessTokenRequestBuilder) Request() *SynchronizationAcquireAccessTokenRequest {
	return &SynchronizationAcquireAccessTokenRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SynchronizationAcquireAccessTokenRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type SynchronizationJobStartRequestBuilder struct{ BaseRequestBuilder }

// Start action undocumentedrav
func (b *SynchronizationJobRequestBuilder) Start(reqObj *SynchronizationJobStartRequestParameter) *SynchronizationJobStartRequestBuilder {
	bb := &SynchronizationJobStartRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Start"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SynchronizationJobStartRequest struct{ BaseRequest }

func (b *SynchronizationJobStartRequestBuilder) Request() *SynchronizationJobStartRequest {
	return &SynchronizationJobStartRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SynchronizationJobStartRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type SynchronizationJobStopRequestBuilder struct{ BaseRequestBuilder }

// Stop action undocumentedrav
func (b *SynchronizationJobRequestBuilder) Stop(reqObj *SynchronizationJobStopRequestParameter) *SynchronizationJobStopRequestBuilder {
	bb := &SynchronizationJobStopRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Stop"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SynchronizationJobStopRequest struct{ BaseRequest }

func (b *SynchronizationJobStopRequestBuilder) Request() *SynchronizationJobStopRequest {
	return &SynchronizationJobStopRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SynchronizationJobStopRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type SynchronizationJobPauseRequestBuilder struct{ BaseRequestBuilder }

// Pause action undocumentedrav
func (b *SynchronizationJobRequestBuilder) Pause(reqObj *SynchronizationJobPauseRequestParameter) *SynchronizationJobPauseRequestBuilder {
	bb := &SynchronizationJobPauseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Pause"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SynchronizationJobPauseRequest struct{ BaseRequest }

func (b *SynchronizationJobPauseRequestBuilder) Request() *SynchronizationJobPauseRequest {
	return &SynchronizationJobPauseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SynchronizationJobPauseRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type SynchronizationJobProvisionOnDemandRequestBuilder struct{ BaseRequestBuilder }

// ProvisionOnDemand action undocumentedras
func (b *SynchronizationJobRequestBuilder) ProvisionOnDemand(reqObj *SynchronizationJobProvisionOnDemandRequestParameter) *SynchronizationJobProvisionOnDemandRequestBuilder {
	bb := &SynchronizationJobProvisionOnDemandRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ProvisionOnDemand"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SynchronizationJobProvisionOnDemandRequest struct{ BaseRequest }

func (b *SynchronizationJobProvisionOnDemandRequestBuilder) Request() *SynchronizationJobProvisionOnDemandRequest {
	return &SynchronizationJobProvisionOnDemandRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SynchronizationJobProvisionOnDemandRequest) Post(ctx context.Context) (resObj *StringKeyStringValuePair, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type SynchronizationJobRestartRequestBuilder struct{ BaseRequestBuilder }

// Restart action undocumentedrav
func (b *SynchronizationJobRequestBuilder) Restart(reqObj *SynchronizationJobRestartRequestParameter) *SynchronizationJobRestartRequestBuilder {
	bb := &SynchronizationJobRestartRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Restart"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SynchronizationJobRestartRequest struct{ BaseRequest }

func (b *SynchronizationJobRestartRequestBuilder) Request() *SynchronizationJobRestartRequest {
	return &SynchronizationJobRestartRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SynchronizationJobRestartRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type SynchronizationJobValidateCredentialsRequestBuilder struct{ BaseRequestBuilder }

// ValidateCredentials action undocumentedrav
func (b *SynchronizationJobRequestBuilder) ValidateCredentials(reqObj *SynchronizationJobValidateCredentialsRequestParameter) *SynchronizationJobValidateCredentialsRequestBuilder {
	bb := &SynchronizationJobValidateCredentialsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ValidateCredentials"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SynchronizationJobValidateCredentialsRequest struct{ BaseRequest }

func (b *SynchronizationJobValidateCredentialsRequestBuilder) Request() *SynchronizationJobValidateCredentialsRequest {
	return &SynchronizationJobValidateCredentialsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SynchronizationJobValidateCredentialsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type SynchronizationSchemaParseExpressionRequestBuilder struct{ BaseRequestBuilder }

// ParseExpression action undocumentedras
func (b *SynchronizationSchemaRequestBuilder) ParseExpression(reqObj *SynchronizationSchemaParseExpressionRequestParameter) *SynchronizationSchemaParseExpressionRequestBuilder {
	bb := &SynchronizationSchemaParseExpressionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ParseExpression"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type SynchronizationSchemaParseExpressionRequest struct{ BaseRequest }

func (b *SynchronizationSchemaParseExpressionRequestBuilder) Request() *SynchronizationSchemaParseExpressionRequest {
	return &SynchronizationSchemaParseExpressionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *SynchronizationSchemaParseExpressionRequest) Post(ctx context.Context) (resObj *ParseExpressionResponse, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

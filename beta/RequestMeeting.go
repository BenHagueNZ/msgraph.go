// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MeetingActivityStatisticsRequestBuilder is request builder for MeetingActivityStatistics
type MeetingActivityStatisticsRequestBuilder struct{ BaseRequestBuilder }

// Request returns MeetingActivityStatisticsRequest
func (b *MeetingActivityStatisticsRequestBuilder) Request() *MeetingActivityStatisticsRequest {
	return &MeetingActivityStatisticsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MeetingActivityStatisticsRequest is request for MeetingActivityStatistics
type MeetingActivityStatisticsRequest struct{ BaseRequest }

// Get performs GET request for MeetingActivityStatistics
func (r *MeetingActivityStatisticsRequest) Get(ctx context.Context) (resObj *MeetingActivityStatistics, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MeetingActivityStatistics
func (r *MeetingActivityStatisticsRequest) Update(ctx context.Context, reqObj *MeetingActivityStatistics) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MeetingActivityStatistics
func (r *MeetingActivityStatisticsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MeetingAttendanceReportRequestBuilder is request builder for MeetingAttendanceReport
type MeetingAttendanceReportRequestBuilder struct{ BaseRequestBuilder }

// Request returns MeetingAttendanceReportRequest
func (b *MeetingAttendanceReportRequestBuilder) Request() *MeetingAttendanceReportRequest {
	return &MeetingAttendanceReportRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MeetingAttendanceReportRequest is request for MeetingAttendanceReport
type MeetingAttendanceReportRequest struct{ BaseRequest }

// Get performs GET request for MeetingAttendanceReport
func (r *MeetingAttendanceReportRequest) Get(ctx context.Context) (resObj *MeetingAttendanceReport, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MeetingAttendanceReport
func (r *MeetingAttendanceReportRequest) Update(ctx context.Context, reqObj *MeetingAttendanceReport) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MeetingAttendanceReport
func (r *MeetingAttendanceReportRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MeetingCapabilityRequestBuilder is request builder for MeetingCapability
type MeetingCapabilityRequestBuilder struct{ BaseRequestBuilder }

// Request returns MeetingCapabilityRequest
func (b *MeetingCapabilityRequestBuilder) Request() *MeetingCapabilityRequest {
	return &MeetingCapabilityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MeetingCapabilityRequest is request for MeetingCapability
type MeetingCapabilityRequest struct{ BaseRequest }

// Get performs GET request for MeetingCapability
func (r *MeetingCapabilityRequest) Get(ctx context.Context) (resObj *MeetingCapability, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MeetingCapability
func (r *MeetingCapabilityRequest) Update(ctx context.Context, reqObj *MeetingCapability) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MeetingCapability
func (r *MeetingCapabilityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MeetingInfoRequestBuilder is request builder for MeetingInfo
type MeetingInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns MeetingInfoRequest
func (b *MeetingInfoRequestBuilder) Request() *MeetingInfoRequest {
	return &MeetingInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MeetingInfoRequest is request for MeetingInfo
type MeetingInfoRequest struct{ BaseRequest }

// Get performs GET request for MeetingInfo
func (r *MeetingInfoRequest) Get(ctx context.Context) (resObj *MeetingInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MeetingInfo
func (r *MeetingInfoRequest) Update(ctx context.Context, reqObj *MeetingInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MeetingInfo
func (r *MeetingInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MeetingParticipantInfoRequestBuilder is request builder for MeetingParticipantInfo
type MeetingParticipantInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns MeetingParticipantInfoRequest
func (b *MeetingParticipantInfoRequestBuilder) Request() *MeetingParticipantInfoRequest {
	return &MeetingParticipantInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MeetingParticipantInfoRequest is request for MeetingParticipantInfo
type MeetingParticipantInfoRequest struct{ BaseRequest }

// Get performs GET request for MeetingParticipantInfo
func (r *MeetingParticipantInfoRequest) Get(ctx context.Context) (resObj *MeetingParticipantInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MeetingParticipantInfo
func (r *MeetingParticipantInfoRequest) Update(ctx context.Context, reqObj *MeetingParticipantInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MeetingParticipantInfo
func (r *MeetingParticipantInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MeetingParticipantsRequestBuilder is request builder for MeetingParticipants
type MeetingParticipantsRequestBuilder struct{ BaseRequestBuilder }

// Request returns MeetingParticipantsRequest
func (b *MeetingParticipantsRequestBuilder) Request() *MeetingParticipantsRequest {
	return &MeetingParticipantsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MeetingParticipantsRequest is request for MeetingParticipants
type MeetingParticipantsRequest struct{ BaseRequest }

// Get performs GET request for MeetingParticipants
func (r *MeetingParticipantsRequest) Get(ctx context.Context) (resObj *MeetingParticipants, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MeetingParticipants
func (r *MeetingParticipantsRequest) Update(ctx context.Context, reqObj *MeetingParticipants) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MeetingParticipants
func (r *MeetingParticipantsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MeetingPolicyUpdatedEventMessageDetailRequestBuilder is request builder for MeetingPolicyUpdatedEventMessageDetail
type MeetingPolicyUpdatedEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns MeetingPolicyUpdatedEventMessageDetailRequest
func (b *MeetingPolicyUpdatedEventMessageDetailRequestBuilder) Request() *MeetingPolicyUpdatedEventMessageDetailRequest {
	return &MeetingPolicyUpdatedEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MeetingPolicyUpdatedEventMessageDetailRequest is request for MeetingPolicyUpdatedEventMessageDetail
type MeetingPolicyUpdatedEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for MeetingPolicyUpdatedEventMessageDetail
func (r *MeetingPolicyUpdatedEventMessageDetailRequest) Get(ctx context.Context) (resObj *MeetingPolicyUpdatedEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MeetingPolicyUpdatedEventMessageDetail
func (r *MeetingPolicyUpdatedEventMessageDetailRequest) Update(ctx context.Context, reqObj *MeetingPolicyUpdatedEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MeetingPolicyUpdatedEventMessageDetail
func (r *MeetingPolicyUpdatedEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MeetingRegistrantRequestBuilder is request builder for MeetingRegistrant
type MeetingRegistrantRequestBuilder struct{ BaseRequestBuilder }

// Request returns MeetingRegistrantRequest
func (b *MeetingRegistrantRequestBuilder) Request() *MeetingRegistrantRequest {
	return &MeetingRegistrantRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MeetingRegistrantRequest is request for MeetingRegistrant
type MeetingRegistrantRequest struct{ BaseRequest }

// Get performs GET request for MeetingRegistrant
func (r *MeetingRegistrantRequest) Get(ctx context.Context) (resObj *MeetingRegistrant, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MeetingRegistrant
func (r *MeetingRegistrantRequest) Update(ctx context.Context, reqObj *MeetingRegistrant) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MeetingRegistrant
func (r *MeetingRegistrantRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MeetingRegistrantBaseRequestBuilder is request builder for MeetingRegistrantBase
type MeetingRegistrantBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns MeetingRegistrantBaseRequest
func (b *MeetingRegistrantBaseRequestBuilder) Request() *MeetingRegistrantBaseRequest {
	return &MeetingRegistrantBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MeetingRegistrantBaseRequest is request for MeetingRegistrantBase
type MeetingRegistrantBaseRequest struct{ BaseRequest }

// Get performs GET request for MeetingRegistrantBase
func (r *MeetingRegistrantBaseRequest) Get(ctx context.Context) (resObj *MeetingRegistrantBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MeetingRegistrantBase
func (r *MeetingRegistrantBaseRequest) Update(ctx context.Context, reqObj *MeetingRegistrantBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MeetingRegistrantBase
func (r *MeetingRegistrantBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MeetingRegistrationRequestBuilder is request builder for MeetingRegistration
type MeetingRegistrationRequestBuilder struct{ BaseRequestBuilder }

// Request returns MeetingRegistrationRequest
func (b *MeetingRegistrationRequestBuilder) Request() *MeetingRegistrationRequest {
	return &MeetingRegistrationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MeetingRegistrationRequest is request for MeetingRegistration
type MeetingRegistrationRequest struct{ BaseRequest }

// Get performs GET request for MeetingRegistration
func (r *MeetingRegistrationRequest) Get(ctx context.Context) (resObj *MeetingRegistration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MeetingRegistration
func (r *MeetingRegistrationRequest) Update(ctx context.Context, reqObj *MeetingRegistration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MeetingRegistration
func (r *MeetingRegistrationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MeetingRegistrationBaseRequestBuilder is request builder for MeetingRegistrationBase
type MeetingRegistrationBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns MeetingRegistrationBaseRequest
func (b *MeetingRegistrationBaseRequestBuilder) Request() *MeetingRegistrationBaseRequest {
	return &MeetingRegistrationBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MeetingRegistrationBaseRequest is request for MeetingRegistrationBase
type MeetingRegistrationBaseRequest struct{ BaseRequest }

// Get performs GET request for MeetingRegistrationBase
func (r *MeetingRegistrationBaseRequest) Get(ctx context.Context) (resObj *MeetingRegistrationBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MeetingRegistrationBase
func (r *MeetingRegistrationBaseRequest) Update(ctx context.Context, reqObj *MeetingRegistrationBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MeetingRegistrationBase
func (r *MeetingRegistrationBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MeetingRegistrationQuestionRequestBuilder is request builder for MeetingRegistrationQuestion
type MeetingRegistrationQuestionRequestBuilder struct{ BaseRequestBuilder }

// Request returns MeetingRegistrationQuestionRequest
func (b *MeetingRegistrationQuestionRequestBuilder) Request() *MeetingRegistrationQuestionRequest {
	return &MeetingRegistrationQuestionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MeetingRegistrationQuestionRequest is request for MeetingRegistrationQuestion
type MeetingRegistrationQuestionRequest struct{ BaseRequest }

// Get performs GET request for MeetingRegistrationQuestion
func (r *MeetingRegistrationQuestionRequest) Get(ctx context.Context) (resObj *MeetingRegistrationQuestion, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MeetingRegistrationQuestion
func (r *MeetingRegistrationQuestionRequest) Update(ctx context.Context, reqObj *MeetingRegistrationQuestion) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MeetingRegistrationQuestion
func (r *MeetingRegistrationQuestionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MeetingSpeakerRequestBuilder is request builder for MeetingSpeaker
type MeetingSpeakerRequestBuilder struct{ BaseRequestBuilder }

// Request returns MeetingSpeakerRequest
func (b *MeetingSpeakerRequestBuilder) Request() *MeetingSpeakerRequest {
	return &MeetingSpeakerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MeetingSpeakerRequest is request for MeetingSpeaker
type MeetingSpeakerRequest struct{ BaseRequest }

// Get performs GET request for MeetingSpeaker
func (r *MeetingSpeakerRequest) Get(ctx context.Context) (resObj *MeetingSpeaker, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MeetingSpeaker
func (r *MeetingSpeakerRequest) Update(ctx context.Context, reqObj *MeetingSpeaker) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MeetingSpeaker
func (r *MeetingSpeakerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MeetingTimeSuggestionRequestBuilder is request builder for MeetingTimeSuggestion
type MeetingTimeSuggestionRequestBuilder struct{ BaseRequestBuilder }

// Request returns MeetingTimeSuggestionRequest
func (b *MeetingTimeSuggestionRequestBuilder) Request() *MeetingTimeSuggestionRequest {
	return &MeetingTimeSuggestionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MeetingTimeSuggestionRequest is request for MeetingTimeSuggestion
type MeetingTimeSuggestionRequest struct{ BaseRequest }

// Get performs GET request for MeetingTimeSuggestion
func (r *MeetingTimeSuggestionRequest) Get(ctx context.Context) (resObj *MeetingTimeSuggestion, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MeetingTimeSuggestion
func (r *MeetingTimeSuggestionRequest) Update(ctx context.Context, reqObj *MeetingTimeSuggestion) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MeetingTimeSuggestion
func (r *MeetingTimeSuggestionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MeetingTimeSuggestionsResultRequestBuilder is request builder for MeetingTimeSuggestionsResult
type MeetingTimeSuggestionsResultRequestBuilder struct{ BaseRequestBuilder }

// Request returns MeetingTimeSuggestionsResultRequest
func (b *MeetingTimeSuggestionsResultRequestBuilder) Request() *MeetingTimeSuggestionsResultRequest {
	return &MeetingTimeSuggestionsResultRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MeetingTimeSuggestionsResultRequest is request for MeetingTimeSuggestionsResult
type MeetingTimeSuggestionsResultRequest struct{ BaseRequest }

// Get performs GET request for MeetingTimeSuggestionsResult
func (r *MeetingTimeSuggestionsResultRequest) Get(ctx context.Context) (resObj *MeetingTimeSuggestionsResult, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MeetingTimeSuggestionsResult
func (r *MeetingTimeSuggestionsResultRequest) Update(ctx context.Context, reqObj *MeetingTimeSuggestionsResult) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MeetingTimeSuggestionsResult
func (r *MeetingTimeSuggestionsResultRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

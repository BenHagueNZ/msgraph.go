// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// UserRequestBuilder is request builder for User
type UserRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserRequest
func (b *UserRequestBuilder) Request() *UserRequest {
	return &UserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserRequest is request for User
type UserRequest struct{ BaseRequest }

// Get performs GET request for User
func (r *UserRequest) Get(ctx context.Context) (resObj *User, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for User
func (r *UserRequest) Update(ctx context.Context, reqObj *User) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for User
func (r *UserRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserActivityRequestBuilder is request builder for UserActivity
type UserActivityRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserActivityRequest
func (b *UserActivityRequestBuilder) Request() *UserActivityRequest {
	return &UserActivityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserActivityRequest is request for UserActivity
type UserActivityRequest struct{ BaseRequest }

// Get performs GET request for UserActivity
func (r *UserActivityRequest) Get(ctx context.Context) (resObj *UserActivity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserActivity
func (r *UserActivityRequest) Update(ctx context.Context, reqObj *UserActivity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserActivity
func (r *UserActivityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserAttributeValuesItemRequestBuilder is request builder for UserAttributeValuesItem
type UserAttributeValuesItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserAttributeValuesItemRequest
func (b *UserAttributeValuesItemRequestBuilder) Request() *UserAttributeValuesItemRequest {
	return &UserAttributeValuesItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserAttributeValuesItemRequest is request for UserAttributeValuesItem
type UserAttributeValuesItemRequest struct{ BaseRequest }

// Get performs GET request for UserAttributeValuesItem
func (r *UserAttributeValuesItemRequest) Get(ctx context.Context) (resObj *UserAttributeValuesItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserAttributeValuesItem
func (r *UserAttributeValuesItemRequest) Update(ctx context.Context, reqObj *UserAttributeValuesItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserAttributeValuesItem
func (r *UserAttributeValuesItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserConsentRequestObjectRequestBuilder is request builder for UserConsentRequestObject
type UserConsentRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserConsentRequestObjectRequest
func (b *UserConsentRequestObjectRequestBuilder) Request() *UserConsentRequestObjectRequest {
	return &UserConsentRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserConsentRequestObjectRequest is request for UserConsentRequestObject
type UserConsentRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for UserConsentRequestObject
func (r *UserConsentRequestObjectRequest) Get(ctx context.Context) (resObj *UserConsentRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserConsentRequestObject
func (r *UserConsentRequestObjectRequest) Update(ctx context.Context, reqObj *UserConsentRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserConsentRequestObject
func (r *UserConsentRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserExperienceAnalyticsDevicePerformanceRequestBuilder is request builder for UserExperienceAnalyticsDevicePerformance
type UserExperienceAnalyticsDevicePerformanceRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserExperienceAnalyticsDevicePerformanceRequest
func (b *UserExperienceAnalyticsDevicePerformanceRequestBuilder) Request() *UserExperienceAnalyticsDevicePerformanceRequest {
	return &UserExperienceAnalyticsDevicePerformanceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserExperienceAnalyticsDevicePerformanceRequest is request for UserExperienceAnalyticsDevicePerformance
type UserExperienceAnalyticsDevicePerformanceRequest struct{ BaseRequest }

// Get performs GET request for UserExperienceAnalyticsDevicePerformance
func (r *UserExperienceAnalyticsDevicePerformanceRequest) Get(ctx context.Context) (resObj *UserExperienceAnalyticsDevicePerformance, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserExperienceAnalyticsDevicePerformance
func (r *UserExperienceAnalyticsDevicePerformanceRequest) Update(ctx context.Context, reqObj *UserExperienceAnalyticsDevicePerformance) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserExperienceAnalyticsDevicePerformance
func (r *UserExperienceAnalyticsDevicePerformanceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserFlowAPIConnectorConfigurationRequestBuilder is request builder for UserFlowAPIConnectorConfiguration
type UserFlowAPIConnectorConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserFlowAPIConnectorConfigurationRequest
func (b *UserFlowAPIConnectorConfigurationRequestBuilder) Request() *UserFlowAPIConnectorConfigurationRequest {
	return &UserFlowAPIConnectorConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserFlowAPIConnectorConfigurationRequest is request for UserFlowAPIConnectorConfiguration
type UserFlowAPIConnectorConfigurationRequest struct{ BaseRequest }

// Get performs GET request for UserFlowAPIConnectorConfiguration
func (r *UserFlowAPIConnectorConfigurationRequest) Get(ctx context.Context) (resObj *UserFlowAPIConnectorConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserFlowAPIConnectorConfiguration
func (r *UserFlowAPIConnectorConfigurationRequest) Update(ctx context.Context, reqObj *UserFlowAPIConnectorConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserFlowAPIConnectorConfiguration
func (r *UserFlowAPIConnectorConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserFlowLanguageConfigurationRequestBuilder is request builder for UserFlowLanguageConfiguration
type UserFlowLanguageConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserFlowLanguageConfigurationRequest
func (b *UserFlowLanguageConfigurationRequestBuilder) Request() *UserFlowLanguageConfigurationRequest {
	return &UserFlowLanguageConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserFlowLanguageConfigurationRequest is request for UserFlowLanguageConfiguration
type UserFlowLanguageConfigurationRequest struct{ BaseRequest }

// Get performs GET request for UserFlowLanguageConfiguration
func (r *UserFlowLanguageConfigurationRequest) Get(ctx context.Context) (resObj *UserFlowLanguageConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserFlowLanguageConfiguration
func (r *UserFlowLanguageConfigurationRequest) Update(ctx context.Context, reqObj *UserFlowLanguageConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserFlowLanguageConfiguration
func (r *UserFlowLanguageConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserFlowLanguagePageRequestBuilder is request builder for UserFlowLanguagePage
type UserFlowLanguagePageRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserFlowLanguagePageRequest
func (b *UserFlowLanguagePageRequestBuilder) Request() *UserFlowLanguagePageRequest {
	return &UserFlowLanguagePageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserFlowLanguagePageRequest is request for UserFlowLanguagePage
type UserFlowLanguagePageRequest struct{ BaseRequest }

// Get performs GET request for UserFlowLanguagePage
func (r *UserFlowLanguagePageRequest) Get(ctx context.Context) (resObj *UserFlowLanguagePage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserFlowLanguagePage
func (r *UserFlowLanguagePageRequest) Update(ctx context.Context, reqObj *UserFlowLanguagePage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserFlowLanguagePage
func (r *UserFlowLanguagePageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserIdentityRequestBuilder is request builder for UserIdentity
type UserIdentityRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserIdentityRequest
func (b *UserIdentityRequestBuilder) Request() *UserIdentityRequest {
	return &UserIdentityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserIdentityRequest is request for UserIdentity
type UserIdentityRequest struct{ BaseRequest }

// Get performs GET request for UserIdentity
func (r *UserIdentityRequest) Get(ctx context.Context) (resObj *UserIdentity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserIdentity
func (r *UserIdentityRequest) Update(ctx context.Context, reqObj *UserIdentity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserIdentity
func (r *UserIdentityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserInstallStateSummaryRequestBuilder is request builder for UserInstallStateSummary
type UserInstallStateSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserInstallStateSummaryRequest
func (b *UserInstallStateSummaryRequestBuilder) Request() *UserInstallStateSummaryRequest {
	return &UserInstallStateSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserInstallStateSummaryRequest is request for UserInstallStateSummary
type UserInstallStateSummaryRequest struct{ BaseRequest }

// Get performs GET request for UserInstallStateSummary
func (r *UserInstallStateSummaryRequest) Get(ctx context.Context) (resObj *UserInstallStateSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserInstallStateSummary
func (r *UserInstallStateSummaryRequest) Update(ctx context.Context, reqObj *UserInstallStateSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserInstallStateSummary
func (r *UserInstallStateSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserScopeTeamsAppInstallationRequestBuilder is request builder for UserScopeTeamsAppInstallation
type UserScopeTeamsAppInstallationRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserScopeTeamsAppInstallationRequest
func (b *UserScopeTeamsAppInstallationRequestBuilder) Request() *UserScopeTeamsAppInstallationRequest {
	return &UserScopeTeamsAppInstallationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserScopeTeamsAppInstallationRequest is request for UserScopeTeamsAppInstallation
type UserScopeTeamsAppInstallationRequest struct{ BaseRequest }

// Get performs GET request for UserScopeTeamsAppInstallation
func (r *UserScopeTeamsAppInstallationRequest) Get(ctx context.Context) (resObj *UserScopeTeamsAppInstallation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserScopeTeamsAppInstallation
func (r *UserScopeTeamsAppInstallationRequest) Update(ctx context.Context, reqObj *UserScopeTeamsAppInstallation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserScopeTeamsAppInstallation
func (r *UserScopeTeamsAppInstallationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserSecurityStateRequestBuilder is request builder for UserSecurityState
type UserSecurityStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserSecurityStateRequest
func (b *UserSecurityStateRequestBuilder) Request() *UserSecurityStateRequest {
	return &UserSecurityStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserSecurityStateRequest is request for UserSecurityState
type UserSecurityStateRequest struct{ BaseRequest }

// Get performs GET request for UserSecurityState
func (r *UserSecurityStateRequest) Get(ctx context.Context) (resObj *UserSecurityState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserSecurityState
func (r *UserSecurityStateRequest) Update(ctx context.Context, reqObj *UserSecurityState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserSecurityState
func (r *UserSecurityStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserSettingsRequestBuilder is request builder for UserSettings
type UserSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserSettingsRequest
func (b *UserSettingsRequestBuilder) Request() *UserSettingsRequest {
	return &UserSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserSettingsRequest is request for UserSettings
type UserSettingsRequest struct{ BaseRequest }

// Get performs GET request for UserSettings
func (r *UserSettingsRequest) Get(ctx context.Context) (resObj *UserSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserSettings
func (r *UserSettingsRequest) Update(ctx context.Context, reqObj *UserSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserSettings
func (r *UserSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserSimulationDetailsRequestBuilder is request builder for UserSimulationDetails
type UserSimulationDetailsRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserSimulationDetailsRequest
func (b *UserSimulationDetailsRequestBuilder) Request() *UserSimulationDetailsRequest {
	return &UserSimulationDetailsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserSimulationDetailsRequest is request for UserSimulationDetails
type UserSimulationDetailsRequest struct{ BaseRequest }

// Get performs GET request for UserSimulationDetails
func (r *UserSimulationDetailsRequest) Get(ctx context.Context) (resObj *UserSimulationDetails, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserSimulationDetails
func (r *UserSimulationDetailsRequest) Update(ctx context.Context, reqObj *UserSimulationDetails) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserSimulationDetails
func (r *UserSimulationDetailsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserSimulationEventInfoRequestBuilder is request builder for UserSimulationEventInfo
type UserSimulationEventInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserSimulationEventInfoRequest
func (b *UserSimulationEventInfoRequestBuilder) Request() *UserSimulationEventInfoRequest {
	return &UserSimulationEventInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserSimulationEventInfoRequest is request for UserSimulationEventInfo
type UserSimulationEventInfoRequest struct{ BaseRequest }

// Get performs GET request for UserSimulationEventInfo
func (r *UserSimulationEventInfoRequest) Get(ctx context.Context) (resObj *UserSimulationEventInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserSimulationEventInfo
func (r *UserSimulationEventInfoRequest) Update(ctx context.Context, reqObj *UserSimulationEventInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserSimulationEventInfo
func (r *UserSimulationEventInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserTeamworkRequestBuilder is request builder for UserTeamwork
type UserTeamworkRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserTeamworkRequest
func (b *UserTeamworkRequestBuilder) Request() *UserTeamworkRequest {
	return &UserTeamworkRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserTeamworkRequest is request for UserTeamwork
type UserTeamworkRequest struct{ BaseRequest }

// Get performs GET request for UserTeamwork
func (r *UserTeamworkRequest) Get(ctx context.Context) (resObj *UserTeamwork, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserTeamwork
func (r *UserTeamworkRequest) Update(ctx context.Context, reqObj *UserTeamwork) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserTeamwork
func (r *UserTeamworkRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserTrainingContentEventInfoRequestBuilder is request builder for UserTrainingContentEventInfo
type UserTrainingContentEventInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserTrainingContentEventInfoRequest
func (b *UserTrainingContentEventInfoRequestBuilder) Request() *UserTrainingContentEventInfoRequest {
	return &UserTrainingContentEventInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserTrainingContentEventInfoRequest is request for UserTrainingContentEventInfo
type UserTrainingContentEventInfoRequest struct{ BaseRequest }

// Get performs GET request for UserTrainingContentEventInfo
func (r *UserTrainingContentEventInfoRequest) Get(ctx context.Context) (resObj *UserTrainingContentEventInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserTrainingContentEventInfo
func (r *UserTrainingContentEventInfoRequest) Update(ctx context.Context, reqObj *UserTrainingContentEventInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserTrainingContentEventInfo
func (r *UserTrainingContentEventInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserTrainingEventInfoRequestBuilder is request builder for UserTrainingEventInfo
type UserTrainingEventInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserTrainingEventInfoRequest
func (b *UserTrainingEventInfoRequestBuilder) Request() *UserTrainingEventInfoRequest {
	return &UserTrainingEventInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserTrainingEventInfoRequest is request for UserTrainingEventInfo
type UserTrainingEventInfoRequest struct{ BaseRequest }

// Get performs GET request for UserTrainingEventInfo
func (r *UserTrainingEventInfoRequest) Get(ctx context.Context) (resObj *UserTrainingEventInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserTrainingEventInfo
func (r *UserTrainingEventInfoRequest) Update(ctx context.Context, reqObj *UserTrainingEventInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserTrainingEventInfo
func (r *UserTrainingEventInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UserTrainingStatusInfoRequestBuilder is request builder for UserTrainingStatusInfo
type UserTrainingStatusInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserTrainingStatusInfoRequest
func (b *UserTrainingStatusInfoRequestBuilder) Request() *UserTrainingStatusInfoRequest {
	return &UserTrainingStatusInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserTrainingStatusInfoRequest is request for UserTrainingStatusInfo
type UserTrainingStatusInfoRequest struct{ BaseRequest }

// Get performs GET request for UserTrainingStatusInfo
func (r *UserTrainingStatusInfoRequest) Get(ctx context.Context) (resObj *UserTrainingStatusInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserTrainingStatusInfo
func (r *UserTrainingStatusInfoRequest) Update(ctx context.Context, reqObj *UserTrainingStatusInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserTrainingStatusInfo
func (r *UserTrainingStatusInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type UserAssignLicenseRequestBuilder struct{ BaseRequestBuilder }

// AssignLicense action undocumentedras
func (b *UserRequestBuilder) AssignLicense(reqObj *UserAssignLicenseRequestParameter) *UserAssignLicenseRequestBuilder {
	bb := &UserAssignLicenseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AssignLicense"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type UserAssignLicenseRequest struct{ BaseRequest }

func (b *UserAssignLicenseRequestBuilder) Request() *UserAssignLicenseRequest {
	return &UserAssignLicenseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *UserAssignLicenseRequest) Post(ctx context.Context) (resObj *User, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type UserChangePasswordRequestBuilder struct{ BaseRequestBuilder }

// ChangePassword action undocumentedrav
func (b *UserRequestBuilder) ChangePassword(reqObj *UserChangePasswordRequestParameter) *UserChangePasswordRequestBuilder {
	bb := &UserChangePasswordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ChangePassword"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type UserChangePasswordRequest struct{ BaseRequest }

func (b *UserChangePasswordRequestBuilder) Request() *UserChangePasswordRequest {
	return &UserChangePasswordRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *UserChangePasswordRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type UserReprocessLicenseAssignmentRequestBuilder struct{ BaseRequestBuilder }

// ReprocessLicenseAssignment action undocumentedras
func (b *UserRequestBuilder) ReprocessLicenseAssignment(reqObj *UserReprocessLicenseAssignmentRequestParameter) *UserReprocessLicenseAssignmentRequestBuilder {
	bb := &UserReprocessLicenseAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ReprocessLicenseAssignment"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type UserReprocessLicenseAssignmentRequest struct{ BaseRequest }

func (b *UserReprocessLicenseAssignmentRequestBuilder) Request() *UserReprocessLicenseAssignmentRequest {
	return &UserReprocessLicenseAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *UserReprocessLicenseAssignmentRequest) Post(ctx context.Context) (resObj *User, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type UserRevokeSignInSessionsRequestBuilder struct{ BaseRequestBuilder }

// RevokeSignInSessions action undocumentedras
func (b *UserRequestBuilder) RevokeSignInSessions(reqObj *UserRevokeSignInSessionsRequestParameter) *UserRevokeSignInSessionsRequestBuilder {
	bb := &UserRevokeSignInSessionsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RevokeSignInSessions"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type UserRevokeSignInSessionsRequest struct{ BaseRequest }

func (b *UserRevokeSignInSessionsRequestBuilder) Request() *UserRevokeSignInSessionsRequest {
	return &UserRevokeSignInSessionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *UserRevokeSignInSessionsRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type UserFindMeetingTimesRequestBuilder struct{ BaseRequestBuilder }

// FindMeetingTimes action undocumentedras
func (b *UserRequestBuilder) FindMeetingTimes(reqObj *UserFindMeetingTimesRequestParameter) *UserFindMeetingTimesRequestBuilder {
	bb := &UserFindMeetingTimesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/FindMeetingTimes"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type UserFindMeetingTimesRequest struct{ BaseRequest }

func (b *UserFindMeetingTimesRequestBuilder) Request() *UserFindMeetingTimesRequest {
	return &UserFindMeetingTimesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *UserFindMeetingTimesRequest) Post(ctx context.Context) (resObj *MeetingTimeSuggestionsResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type UserGetMailTipsRequestBuilder struct{ BaseRequestBuilder }

// GetMailTips action undocumentedrac
func (b *UserRequestBuilder) GetMailTips(reqObj *UserGetMailTipsRequestParameter) *UserGetMailTipsRequestBuilder {
	bb := &UserGetMailTipsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/GetMailTips"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type UserGetMailTipsRequest struct{ BaseRequest }

func (b *UserGetMailTipsRequestBuilder) Request() *UserGetMailTipsRequest {
	return &UserGetMailTipsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *UserGetMailTipsRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]MailTips, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MailTips
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []MailTips
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, _ = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

func (r *UserGetMailTipsRequest) PostN(ctx context.Context, n int) ([]MailTips, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

func (r *UserGetMailTipsRequest) Post(ctx context.Context) ([]MailTips, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

type UserSendMailRequestBuilder struct{ BaseRequestBuilder }

// SendMail action undocumentedrav
func (b *UserRequestBuilder) SendMail(reqObj *UserSendMailRequestParameter) *UserSendMailRequestBuilder {
	bb := &UserSendMailRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/SendMail"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type UserSendMailRequest struct{ BaseRequest }

func (b *UserSendMailRequestBuilder) Request() *UserSendMailRequest {
	return &UserSendMailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *UserSendMailRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type UserTranslateExchangeIDsRequestBuilder struct{ BaseRequestBuilder }

// TranslateExchangeIDs action undocumentedrac
func (b *UserRequestBuilder) TranslateExchangeIDs(reqObj *UserTranslateExchangeIDsRequestParameter) *UserTranslateExchangeIDsRequestBuilder {
	bb := &UserTranslateExchangeIDsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/TranslateExchangeIds"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type UserTranslateExchangeIDsRequest struct{ BaseRequest }

func (b *UserTranslateExchangeIDsRequestBuilder) Request() *UserTranslateExchangeIDsRequest {
	return &UserTranslateExchangeIDsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *UserTranslateExchangeIDsRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ConvertIDResult, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ConvertIDResult
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ConvertIDResult
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, _ = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

func (r *UserTranslateExchangeIDsRequest) PostN(ctx context.Context, n int) ([]ConvertIDResult, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

func (r *UserTranslateExchangeIDsRequest) Post(ctx context.Context) ([]ConvertIDResult, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

type UserRemoveAllDevicesFromManagementRequestBuilder struct{ BaseRequestBuilder }

// RemoveAllDevicesFromManagement action undocumentedrav
func (b *UserRequestBuilder) RemoveAllDevicesFromManagement(reqObj *UserRemoveAllDevicesFromManagementRequestParameter) *UserRemoveAllDevicesFromManagementRequestBuilder {
	bb := &UserRemoveAllDevicesFromManagementRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RemoveAllDevicesFromManagement"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type UserRemoveAllDevicesFromManagementRequest struct{ BaseRequest }

func (b *UserRemoveAllDevicesFromManagementRequestBuilder) Request() *UserRemoveAllDevicesFromManagementRequest {
	return &UserRemoveAllDevicesFromManagementRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *UserRemoveAllDevicesFromManagementRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder struct{ BaseRequestBuilder }

// WipeManagedAppRegistrationsByDeviceTag action undocumentedrav
func (b *UserRequestBuilder) WipeManagedAppRegistrationsByDeviceTag(reqObj *UserWipeManagedAppRegistrationsByDeviceTagRequestParameter) *UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder {
	bb := &UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/WipeManagedAppRegistrationsByDeviceTag"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type UserWipeManagedAppRegistrationsByDeviceTagRequest struct{ BaseRequest }

func (b *UserWipeManagedAppRegistrationsByDeviceTagRequestBuilder) Request() *UserWipeManagedAppRegistrationsByDeviceTagRequest {
	return &UserWipeManagedAppRegistrationsByDeviceTagRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *UserWipeManagedAppRegistrationsByDeviceTagRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type UserExportPersonalDataRequestBuilder struct{ BaseRequestBuilder }

// ExportPersonalData action undocumentedrav
func (b *UserRequestBuilder) ExportPersonalData(reqObj *UserExportPersonalDataRequestParameter) *UserExportPersonalDataRequestBuilder {
	bb := &UserExportPersonalDataRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ExportPersonalData"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type UserExportPersonalDataRequest struct{ BaseRequest }

func (b *UserExportPersonalDataRequestBuilder) Request() *UserExportPersonalDataRequest {
	return &UserExportPersonalDataRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *UserExportPersonalDataRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type UserTeamworkSendActivityNotificationRequestBuilder struct{ BaseRequestBuilder }

// SendActivityNotification action undocumentedrav
func (b *UserTeamworkRequestBuilder) SendActivityNotification(reqObj *UserTeamworkSendActivityNotificationRequestParameter) *UserTeamworkSendActivityNotificationRequestBuilder {
	bb := &UserTeamworkSendActivityNotificationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/SendActivityNotification"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type UserTeamworkSendActivityNotificationRequest struct{ BaseRequest }

func (b *UserTeamworkSendActivityNotificationRequestBuilder) Request() *UserTeamworkSendActivityNotificationRequest {
	return &UserTeamworkSendActivityNotificationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *UserTeamworkSendActivityNotificationRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

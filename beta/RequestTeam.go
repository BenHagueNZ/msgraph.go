// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TeamRequestBuilder is request builder for Team
type TeamRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamRequest
func (b *TeamRequestBuilder) Request() *TeamRequest {
	return &TeamRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamRequest is request for Team
type TeamRequest struct{ BaseRequest }

// Get performs GET request for Team
func (r *TeamRequest) Get(ctx context.Context) (resObj *Team, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Team
func (r *TeamRequest) Update(ctx context.Context, reqObj *Team) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Team
func (r *TeamRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamArchivedEventMessageDetailRequestBuilder is request builder for TeamArchivedEventMessageDetail
type TeamArchivedEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamArchivedEventMessageDetailRequest
func (b *TeamArchivedEventMessageDetailRequestBuilder) Request() *TeamArchivedEventMessageDetailRequest {
	return &TeamArchivedEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamArchivedEventMessageDetailRequest is request for TeamArchivedEventMessageDetail
type TeamArchivedEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for TeamArchivedEventMessageDetail
func (r *TeamArchivedEventMessageDetailRequest) Get(ctx context.Context) (resObj *TeamArchivedEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamArchivedEventMessageDetail
func (r *TeamArchivedEventMessageDetailRequest) Update(ctx context.Context, reqObj *TeamArchivedEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamArchivedEventMessageDetail
func (r *TeamArchivedEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamClassSettingsRequestBuilder is request builder for TeamClassSettings
type TeamClassSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamClassSettingsRequest
func (b *TeamClassSettingsRequestBuilder) Request() *TeamClassSettingsRequest {
	return &TeamClassSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamClassSettingsRequest is request for TeamClassSettings
type TeamClassSettingsRequest struct{ BaseRequest }

// Get performs GET request for TeamClassSettings
func (r *TeamClassSettingsRequest) Get(ctx context.Context) (resObj *TeamClassSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamClassSettings
func (r *TeamClassSettingsRequest) Update(ctx context.Context, reqObj *TeamClassSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamClassSettings
func (r *TeamClassSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamCreatedEventMessageDetailRequestBuilder is request builder for TeamCreatedEventMessageDetail
type TeamCreatedEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamCreatedEventMessageDetailRequest
func (b *TeamCreatedEventMessageDetailRequestBuilder) Request() *TeamCreatedEventMessageDetailRequest {
	return &TeamCreatedEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamCreatedEventMessageDetailRequest is request for TeamCreatedEventMessageDetail
type TeamCreatedEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for TeamCreatedEventMessageDetail
func (r *TeamCreatedEventMessageDetailRequest) Get(ctx context.Context) (resObj *TeamCreatedEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamCreatedEventMessageDetail
func (r *TeamCreatedEventMessageDetailRequest) Update(ctx context.Context, reqObj *TeamCreatedEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamCreatedEventMessageDetail
func (r *TeamCreatedEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamDescriptionUpdatedEventMessageDetailRequestBuilder is request builder for TeamDescriptionUpdatedEventMessageDetail
type TeamDescriptionUpdatedEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamDescriptionUpdatedEventMessageDetailRequest
func (b *TeamDescriptionUpdatedEventMessageDetailRequestBuilder) Request() *TeamDescriptionUpdatedEventMessageDetailRequest {
	return &TeamDescriptionUpdatedEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamDescriptionUpdatedEventMessageDetailRequest is request for TeamDescriptionUpdatedEventMessageDetail
type TeamDescriptionUpdatedEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for TeamDescriptionUpdatedEventMessageDetail
func (r *TeamDescriptionUpdatedEventMessageDetailRequest) Get(ctx context.Context) (resObj *TeamDescriptionUpdatedEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamDescriptionUpdatedEventMessageDetail
func (r *TeamDescriptionUpdatedEventMessageDetailRequest) Update(ctx context.Context, reqObj *TeamDescriptionUpdatedEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamDescriptionUpdatedEventMessageDetail
func (r *TeamDescriptionUpdatedEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamDiscoverySettingsRequestBuilder is request builder for TeamDiscoverySettings
type TeamDiscoverySettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamDiscoverySettingsRequest
func (b *TeamDiscoverySettingsRequestBuilder) Request() *TeamDiscoverySettingsRequest {
	return &TeamDiscoverySettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamDiscoverySettingsRequest is request for TeamDiscoverySettings
type TeamDiscoverySettingsRequest struct{ BaseRequest }

// Get performs GET request for TeamDiscoverySettings
func (r *TeamDiscoverySettingsRequest) Get(ctx context.Context) (resObj *TeamDiscoverySettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamDiscoverySettings
func (r *TeamDiscoverySettingsRequest) Update(ctx context.Context, reqObj *TeamDiscoverySettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamDiscoverySettings
func (r *TeamDiscoverySettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamFunSettingsRequestBuilder is request builder for TeamFunSettings
type TeamFunSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamFunSettingsRequest
func (b *TeamFunSettingsRequestBuilder) Request() *TeamFunSettingsRequest {
	return &TeamFunSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamFunSettingsRequest is request for TeamFunSettings
type TeamFunSettingsRequest struct{ BaseRequest }

// Get performs GET request for TeamFunSettings
func (r *TeamFunSettingsRequest) Get(ctx context.Context) (resObj *TeamFunSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamFunSettings
func (r *TeamFunSettingsRequest) Update(ctx context.Context, reqObj *TeamFunSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamFunSettings
func (r *TeamFunSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamGuestSettingsRequestBuilder is request builder for TeamGuestSettings
type TeamGuestSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamGuestSettingsRequest
func (b *TeamGuestSettingsRequestBuilder) Request() *TeamGuestSettingsRequest {
	return &TeamGuestSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamGuestSettingsRequest is request for TeamGuestSettings
type TeamGuestSettingsRequest struct{ BaseRequest }

// Get performs GET request for TeamGuestSettings
func (r *TeamGuestSettingsRequest) Get(ctx context.Context) (resObj *TeamGuestSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamGuestSettings
func (r *TeamGuestSettingsRequest) Update(ctx context.Context, reqObj *TeamGuestSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamGuestSettings
func (r *TeamGuestSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamInfoRequestBuilder is request builder for TeamInfo
type TeamInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamInfoRequest
func (b *TeamInfoRequestBuilder) Request() *TeamInfoRequest {
	return &TeamInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamInfoRequest is request for TeamInfo
type TeamInfoRequest struct{ BaseRequest }

// Get performs GET request for TeamInfo
func (r *TeamInfoRequest) Get(ctx context.Context) (resObj *TeamInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamInfo
func (r *TeamInfoRequest) Update(ctx context.Context, reqObj *TeamInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamInfo
func (r *TeamInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamJoiningDisabledEventMessageDetailRequestBuilder is request builder for TeamJoiningDisabledEventMessageDetail
type TeamJoiningDisabledEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamJoiningDisabledEventMessageDetailRequest
func (b *TeamJoiningDisabledEventMessageDetailRequestBuilder) Request() *TeamJoiningDisabledEventMessageDetailRequest {
	return &TeamJoiningDisabledEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamJoiningDisabledEventMessageDetailRequest is request for TeamJoiningDisabledEventMessageDetail
type TeamJoiningDisabledEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for TeamJoiningDisabledEventMessageDetail
func (r *TeamJoiningDisabledEventMessageDetailRequest) Get(ctx context.Context) (resObj *TeamJoiningDisabledEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamJoiningDisabledEventMessageDetail
func (r *TeamJoiningDisabledEventMessageDetailRequest) Update(ctx context.Context, reqObj *TeamJoiningDisabledEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamJoiningDisabledEventMessageDetail
func (r *TeamJoiningDisabledEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamJoiningEnabledEventMessageDetailRequestBuilder is request builder for TeamJoiningEnabledEventMessageDetail
type TeamJoiningEnabledEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamJoiningEnabledEventMessageDetailRequest
func (b *TeamJoiningEnabledEventMessageDetailRequestBuilder) Request() *TeamJoiningEnabledEventMessageDetailRequest {
	return &TeamJoiningEnabledEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamJoiningEnabledEventMessageDetailRequest is request for TeamJoiningEnabledEventMessageDetail
type TeamJoiningEnabledEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for TeamJoiningEnabledEventMessageDetail
func (r *TeamJoiningEnabledEventMessageDetailRequest) Get(ctx context.Context) (resObj *TeamJoiningEnabledEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamJoiningEnabledEventMessageDetail
func (r *TeamJoiningEnabledEventMessageDetailRequest) Update(ctx context.Context, reqObj *TeamJoiningEnabledEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamJoiningEnabledEventMessageDetail
func (r *TeamJoiningEnabledEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamMemberSettingsRequestBuilder is request builder for TeamMemberSettings
type TeamMemberSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamMemberSettingsRequest
func (b *TeamMemberSettingsRequestBuilder) Request() *TeamMemberSettingsRequest {
	return &TeamMemberSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamMemberSettingsRequest is request for TeamMemberSettings
type TeamMemberSettingsRequest struct{ BaseRequest }

// Get performs GET request for TeamMemberSettings
func (r *TeamMemberSettingsRequest) Get(ctx context.Context) (resObj *TeamMemberSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamMemberSettings
func (r *TeamMemberSettingsRequest) Update(ctx context.Context, reqObj *TeamMemberSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamMemberSettings
func (r *TeamMemberSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamMembersNotificationRecipientRequestBuilder is request builder for TeamMembersNotificationRecipient
type TeamMembersNotificationRecipientRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamMembersNotificationRecipientRequest
func (b *TeamMembersNotificationRecipientRequestBuilder) Request() *TeamMembersNotificationRecipientRequest {
	return &TeamMembersNotificationRecipientRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamMembersNotificationRecipientRequest is request for TeamMembersNotificationRecipient
type TeamMembersNotificationRecipientRequest struct{ BaseRequest }

// Get performs GET request for TeamMembersNotificationRecipient
func (r *TeamMembersNotificationRecipientRequest) Get(ctx context.Context) (resObj *TeamMembersNotificationRecipient, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamMembersNotificationRecipient
func (r *TeamMembersNotificationRecipientRequest) Update(ctx context.Context, reqObj *TeamMembersNotificationRecipient) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamMembersNotificationRecipient
func (r *TeamMembersNotificationRecipientRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamMessagingSettingsRequestBuilder is request builder for TeamMessagingSettings
type TeamMessagingSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamMessagingSettingsRequest
func (b *TeamMessagingSettingsRequestBuilder) Request() *TeamMessagingSettingsRequest {
	return &TeamMessagingSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamMessagingSettingsRequest is request for TeamMessagingSettings
type TeamMessagingSettingsRequest struct{ BaseRequest }

// Get performs GET request for TeamMessagingSettings
func (r *TeamMessagingSettingsRequest) Get(ctx context.Context) (resObj *TeamMessagingSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamMessagingSettings
func (r *TeamMessagingSettingsRequest) Update(ctx context.Context, reqObj *TeamMessagingSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamMessagingSettings
func (r *TeamMessagingSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamRenamedEventMessageDetailRequestBuilder is request builder for TeamRenamedEventMessageDetail
type TeamRenamedEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamRenamedEventMessageDetailRequest
func (b *TeamRenamedEventMessageDetailRequestBuilder) Request() *TeamRenamedEventMessageDetailRequest {
	return &TeamRenamedEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamRenamedEventMessageDetailRequest is request for TeamRenamedEventMessageDetail
type TeamRenamedEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for TeamRenamedEventMessageDetail
func (r *TeamRenamedEventMessageDetailRequest) Get(ctx context.Context) (resObj *TeamRenamedEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamRenamedEventMessageDetail
func (r *TeamRenamedEventMessageDetailRequest) Update(ctx context.Context, reqObj *TeamRenamedEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamRenamedEventMessageDetail
func (r *TeamRenamedEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamSummaryRequestBuilder is request builder for TeamSummary
type TeamSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamSummaryRequest
func (b *TeamSummaryRequestBuilder) Request() *TeamSummaryRequest {
	return &TeamSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamSummaryRequest is request for TeamSummary
type TeamSummaryRequest struct{ BaseRequest }

// Get performs GET request for TeamSummary
func (r *TeamSummaryRequest) Get(ctx context.Context) (resObj *TeamSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamSummary
func (r *TeamSummaryRequest) Update(ctx context.Context, reqObj *TeamSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamSummary
func (r *TeamSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamTemplateRequestBuilder is request builder for TeamTemplate
type TeamTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamTemplateRequest
func (b *TeamTemplateRequestBuilder) Request() *TeamTemplateRequest {
	return &TeamTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamTemplateRequest is request for TeamTemplate
type TeamTemplateRequest struct{ BaseRequest }

// Get performs GET request for TeamTemplate
func (r *TeamTemplateRequest) Get(ctx context.Context) (resObj *TeamTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamTemplate
func (r *TeamTemplateRequest) Update(ctx context.Context, reqObj *TeamTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamTemplate
func (r *TeamTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamTemplateDefinitionRequestBuilder is request builder for TeamTemplateDefinition
type TeamTemplateDefinitionRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamTemplateDefinitionRequest
func (b *TeamTemplateDefinitionRequestBuilder) Request() *TeamTemplateDefinitionRequest {
	return &TeamTemplateDefinitionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamTemplateDefinitionRequest is request for TeamTemplateDefinition
type TeamTemplateDefinitionRequest struct{ BaseRequest }

// Get performs GET request for TeamTemplateDefinition
func (r *TeamTemplateDefinitionRequest) Get(ctx context.Context) (resObj *TeamTemplateDefinition, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamTemplateDefinition
func (r *TeamTemplateDefinitionRequest) Update(ctx context.Context, reqObj *TeamTemplateDefinition) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamTemplateDefinition
func (r *TeamTemplateDefinitionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamUnarchivedEventMessageDetailRequestBuilder is request builder for TeamUnarchivedEventMessageDetail
type TeamUnarchivedEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamUnarchivedEventMessageDetailRequest
func (b *TeamUnarchivedEventMessageDetailRequestBuilder) Request() *TeamUnarchivedEventMessageDetailRequest {
	return &TeamUnarchivedEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamUnarchivedEventMessageDetailRequest is request for TeamUnarchivedEventMessageDetail
type TeamUnarchivedEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for TeamUnarchivedEventMessageDetail
func (r *TeamUnarchivedEventMessageDetailRequest) Get(ctx context.Context) (resObj *TeamUnarchivedEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TeamUnarchivedEventMessageDetail
func (r *TeamUnarchivedEventMessageDetailRequest) Update(ctx context.Context, reqObj *TeamUnarchivedEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TeamUnarchivedEventMessageDetail
func (r *TeamUnarchivedEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type TeamCloneRequestBuilder struct{ BaseRequestBuilder }

// Clone action undocumentedrav
func (b *TeamRequestBuilder) Clone(reqObj *TeamCloneRequestParameter) *TeamCloneRequestBuilder {
	bb := &TeamCloneRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Clone"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type TeamCloneRequest struct{ BaseRequest }

func (b *TeamCloneRequestBuilder) Request() *TeamCloneRequest {
	return &TeamCloneRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *TeamCloneRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type TeamArchiveRequestBuilder struct{ BaseRequestBuilder }

// Archive action undocumentedrav
func (b *TeamRequestBuilder) Archive(reqObj *TeamArchiveRequestParameter) *TeamArchiveRequestBuilder {
	bb := &TeamArchiveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Archive"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type TeamArchiveRequest struct{ BaseRequest }

func (b *TeamArchiveRequestBuilder) Request() *TeamArchiveRequest {
	return &TeamArchiveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *TeamArchiveRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type TeamUnarchiveRequestBuilder struct{ BaseRequestBuilder }

// Unarchive action undocumentedrav
func (b *TeamRequestBuilder) Unarchive(reqObj *TeamUnarchiveRequestParameter) *TeamUnarchiveRequestBuilder {
	bb := &TeamUnarchiveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Unarchive"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type TeamUnarchiveRequest struct{ BaseRequest }

func (b *TeamUnarchiveRequestBuilder) Request() *TeamUnarchiveRequest {
	return &TeamUnarchiveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *TeamUnarchiveRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type TeamCompleteMigrationRequestBuilder struct{ BaseRequestBuilder }

// CompleteMigration action undocumentedrav
func (b *TeamRequestBuilder) CompleteMigration(reqObj *TeamCompleteMigrationRequestParameter) *TeamCompleteMigrationRequestBuilder {
	bb := &TeamCompleteMigrationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CompleteMigration"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type TeamCompleteMigrationRequest struct{ BaseRequest }

func (b *TeamCompleteMigrationRequestBuilder) Request() *TeamCompleteMigrationRequest {
	return &TeamCompleteMigrationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *TeamCompleteMigrationRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type TeamSendActivityNotificationRequestBuilder struct{ BaseRequestBuilder }

// SendActivityNotification action undocumentedrav
func (b *TeamRequestBuilder) SendActivityNotification(reqObj *TeamSendActivityNotificationRequestParameter) *TeamSendActivityNotificationRequestBuilder {
	bb := &TeamSendActivityNotificationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/SendActivityNotification"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type TeamSendActivityNotificationRequest struct{ BaseRequest }

func (b *TeamSendActivityNotificationRequestBuilder) Request() *TeamSendActivityNotificationRequest {
	return &TeamSendActivityNotificationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *TeamSendActivityNotificationRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
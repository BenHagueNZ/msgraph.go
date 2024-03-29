// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// MobileAppRequestBuilder is request builder for MobileApp
type MobileAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppRequest
func (b *MobileAppRequestBuilder) Request() *MobileAppRequest {
	return &MobileAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppRequest is request for MobileApp
type MobileAppRequest struct{ BaseRequest }

// Get performs GET request for MobileApp
func (r *MobileAppRequest) Get(ctx context.Context) (resObj *MobileApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileApp
func (r *MobileAppRequest) Update(ctx context.Context, reqObj *MobileApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileApp
func (r *MobileAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MobileAppAssignmentRequestBuilder is request builder for MobileAppAssignment
type MobileAppAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppAssignmentRequest
func (b *MobileAppAssignmentRequestBuilder) Request() *MobileAppAssignmentRequest {
	return &MobileAppAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppAssignmentRequest is request for MobileAppAssignment
type MobileAppAssignmentRequest struct{ BaseRequest }

// Get performs GET request for MobileAppAssignment
func (r *MobileAppAssignmentRequest) Get(ctx context.Context) (resObj *MobileAppAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileAppAssignment
func (r *MobileAppAssignmentRequest) Update(ctx context.Context, reqObj *MobileAppAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileAppAssignment
func (r *MobileAppAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MobileAppAssignmentSettingsRequestBuilder is request builder for MobileAppAssignmentSettings
type MobileAppAssignmentSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppAssignmentSettingsRequest
func (b *MobileAppAssignmentSettingsRequestBuilder) Request() *MobileAppAssignmentSettingsRequest {
	return &MobileAppAssignmentSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppAssignmentSettingsRequest is request for MobileAppAssignmentSettings
type MobileAppAssignmentSettingsRequest struct{ BaseRequest }

// Get performs GET request for MobileAppAssignmentSettings
func (r *MobileAppAssignmentSettingsRequest) Get(ctx context.Context) (resObj *MobileAppAssignmentSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileAppAssignmentSettings
func (r *MobileAppAssignmentSettingsRequest) Update(ctx context.Context, reqObj *MobileAppAssignmentSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileAppAssignmentSettings
func (r *MobileAppAssignmentSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MobileAppCategoryRequestBuilder is request builder for MobileAppCategory
type MobileAppCategoryRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppCategoryRequest
func (b *MobileAppCategoryRequestBuilder) Request() *MobileAppCategoryRequest {
	return &MobileAppCategoryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppCategoryRequest is request for MobileAppCategory
type MobileAppCategoryRequest struct{ BaseRequest }

// Get performs GET request for MobileAppCategory
func (r *MobileAppCategoryRequest) Get(ctx context.Context) (resObj *MobileAppCategory, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileAppCategory
func (r *MobileAppCategoryRequest) Update(ctx context.Context, reqObj *MobileAppCategory) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileAppCategory
func (r *MobileAppCategoryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MobileAppContentRequestBuilder is request builder for MobileAppContent
type MobileAppContentRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppContentRequest
func (b *MobileAppContentRequestBuilder) Request() *MobileAppContentRequest {
	return &MobileAppContentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppContentRequest is request for MobileAppContent
type MobileAppContentRequest struct{ BaseRequest }

// Get performs GET request for MobileAppContent
func (r *MobileAppContentRequest) Get(ctx context.Context) (resObj *MobileAppContent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileAppContent
func (r *MobileAppContentRequest) Update(ctx context.Context, reqObj *MobileAppContent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileAppContent
func (r *MobileAppContentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MobileAppContentFileRequestBuilder is request builder for MobileAppContentFile
type MobileAppContentFileRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppContentFileRequest
func (b *MobileAppContentFileRequestBuilder) Request() *MobileAppContentFileRequest {
	return &MobileAppContentFileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppContentFileRequest is request for MobileAppContentFile
type MobileAppContentFileRequest struct{ BaseRequest }

// Get performs GET request for MobileAppContentFile
func (r *MobileAppContentFileRequest) Get(ctx context.Context) (resObj *MobileAppContentFile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileAppContentFile
func (r *MobileAppContentFileRequest) Update(ctx context.Context, reqObj *MobileAppContentFile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileAppContentFile
func (r *MobileAppContentFileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MobileAppIdentifierRequestBuilder is request builder for MobileAppIdentifier
type MobileAppIdentifierRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppIdentifierRequest
func (b *MobileAppIdentifierRequestBuilder) Request() *MobileAppIdentifierRequest {
	return &MobileAppIdentifierRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppIdentifierRequest is request for MobileAppIdentifier
type MobileAppIdentifierRequest struct{ BaseRequest }

// Get performs GET request for MobileAppIdentifier
func (r *MobileAppIdentifierRequest) Get(ctx context.Context) (resObj *MobileAppIdentifier, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileAppIdentifier
func (r *MobileAppIdentifierRequest) Update(ctx context.Context, reqObj *MobileAppIdentifier) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileAppIdentifier
func (r *MobileAppIdentifierRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MobileAppInstallTimeSettingsRequestBuilder is request builder for MobileAppInstallTimeSettings
type MobileAppInstallTimeSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppInstallTimeSettingsRequest
func (b *MobileAppInstallTimeSettingsRequestBuilder) Request() *MobileAppInstallTimeSettingsRequest {
	return &MobileAppInstallTimeSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppInstallTimeSettingsRequest is request for MobileAppInstallTimeSettings
type MobileAppInstallTimeSettingsRequest struct{ BaseRequest }

// Get performs GET request for MobileAppInstallTimeSettings
func (r *MobileAppInstallTimeSettingsRequest) Get(ctx context.Context) (resObj *MobileAppInstallTimeSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileAppInstallTimeSettings
func (r *MobileAppInstallTimeSettingsRequest) Update(ctx context.Context, reqObj *MobileAppInstallTimeSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileAppInstallTimeSettings
func (r *MobileAppInstallTimeSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MobileContainedAppRequestBuilder is request builder for MobileContainedApp
type MobileContainedAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileContainedAppRequest
func (b *MobileContainedAppRequestBuilder) Request() *MobileContainedAppRequest {
	return &MobileContainedAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileContainedAppRequest is request for MobileContainedApp
type MobileContainedAppRequest struct{ BaseRequest }

// Get performs GET request for MobileContainedApp
func (r *MobileContainedAppRequest) Get(ctx context.Context) (resObj *MobileContainedApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileContainedApp
func (r *MobileContainedAppRequest) Update(ctx context.Context, reqObj *MobileContainedApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileContainedApp
func (r *MobileContainedAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MobileLobAppRequestBuilder is request builder for MobileLobApp
type MobileLobAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileLobAppRequest
func (b *MobileLobAppRequestBuilder) Request() *MobileLobAppRequest {
	return &MobileLobAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileLobAppRequest is request for MobileLobApp
type MobileLobAppRequest struct{ BaseRequest }

// Get performs GET request for MobileLobApp
func (r *MobileLobAppRequest) Get(ctx context.Context) (resObj *MobileLobApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileLobApp
func (r *MobileLobAppRequest) Update(ctx context.Context, reqObj *MobileLobApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileLobApp
func (r *MobileLobAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MobileThreatDefenseConnectorRequestBuilder is request builder for MobileThreatDefenseConnector
type MobileThreatDefenseConnectorRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileThreatDefenseConnectorRequest
func (b *MobileThreatDefenseConnectorRequestBuilder) Request() *MobileThreatDefenseConnectorRequest {
	return &MobileThreatDefenseConnectorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileThreatDefenseConnectorRequest is request for MobileThreatDefenseConnector
type MobileThreatDefenseConnectorRequest struct{ BaseRequest }

// Get performs GET request for MobileThreatDefenseConnector
func (r *MobileThreatDefenseConnectorRequest) Get(ctx context.Context) (resObj *MobileThreatDefenseConnector, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileThreatDefenseConnector
func (r *MobileThreatDefenseConnectorRequest) Update(ctx context.Context, reqObj *MobileThreatDefenseConnector) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileThreatDefenseConnector
func (r *MobileThreatDefenseConnectorRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type MobileAppAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumentedrav
func (b *MobileAppRequestBuilder) Assign(reqObj *MobileAppAssignRequestParameter) *MobileAppAssignRequestBuilder {
	bb := &MobileAppAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type MobileAppAssignRequest struct{ BaseRequest }

func (b *MobileAppAssignRequestBuilder) Request() *MobileAppAssignRequest {
	return &MobileAppAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *MobileAppAssignRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type MobileAppContentFileCommitRequestBuilder struct{ BaseRequestBuilder }

// Commit action undocumentedrav
func (b *MobileAppContentFileRequestBuilder) Commit(reqObj *MobileAppContentFileCommitRequestParameter) *MobileAppContentFileCommitRequestBuilder {
	bb := &MobileAppContentFileCommitRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Commit"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type MobileAppContentFileCommitRequest struct{ BaseRequest }

func (b *MobileAppContentFileCommitRequestBuilder) Request() *MobileAppContentFileCommitRequest {
	return &MobileAppContentFileCommitRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *MobileAppContentFileCommitRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type MobileAppContentFileRenewUploadRequestBuilder struct{ BaseRequestBuilder }

// RenewUpload action undocumentedrav
func (b *MobileAppContentFileRequestBuilder) RenewUpload(reqObj *MobileAppContentFileRenewUploadRequestParameter) *MobileAppContentFileRenewUploadRequestBuilder {
	bb := &MobileAppContentFileRenewUploadRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RenewUpload"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type MobileAppContentFileRenewUploadRequest struct{ BaseRequest }

func (b *MobileAppContentFileRenewUploadRequestBuilder) Request() *MobileAppContentFileRenewUploadRequest {
	return &MobileAppContentFileRenewUploadRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *MobileAppContentFileRenewUploadRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

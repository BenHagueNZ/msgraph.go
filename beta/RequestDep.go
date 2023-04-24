// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// DepEnrollmentBaseProfileRequestBuilder is request builder for DepEnrollmentBaseProfile
type DepEnrollmentBaseProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns DepEnrollmentBaseProfileRequest
func (b *DepEnrollmentBaseProfileRequestBuilder) Request() *DepEnrollmentBaseProfileRequest {
	return &DepEnrollmentBaseProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DepEnrollmentBaseProfileRequest is request for DepEnrollmentBaseProfile
type DepEnrollmentBaseProfileRequest struct{ BaseRequest }

// Get performs GET request for DepEnrollmentBaseProfile
func (r *DepEnrollmentBaseProfileRequest) Get(ctx context.Context) (resObj *DepEnrollmentBaseProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DepEnrollmentBaseProfile
func (r *DepEnrollmentBaseProfileRequest) Update(ctx context.Context, reqObj *DepEnrollmentBaseProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DepEnrollmentBaseProfile
func (r *DepEnrollmentBaseProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DepEnrollmentProfileRequestBuilder is request builder for DepEnrollmentProfile
type DepEnrollmentProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns DepEnrollmentProfileRequest
func (b *DepEnrollmentProfileRequestBuilder) Request() *DepEnrollmentProfileRequest {
	return &DepEnrollmentProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DepEnrollmentProfileRequest is request for DepEnrollmentProfile
type DepEnrollmentProfileRequest struct{ BaseRequest }

// Get performs GET request for DepEnrollmentProfile
func (r *DepEnrollmentProfileRequest) Get(ctx context.Context) (resObj *DepEnrollmentProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DepEnrollmentProfile
func (r *DepEnrollmentProfileRequest) Update(ctx context.Context, reqObj *DepEnrollmentProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DepEnrollmentProfile
func (r *DepEnrollmentProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DepIOSEnrollmentProfileRequestBuilder is request builder for DepIOSEnrollmentProfile
type DepIOSEnrollmentProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns DepIOSEnrollmentProfileRequest
func (b *DepIOSEnrollmentProfileRequestBuilder) Request() *DepIOSEnrollmentProfileRequest {
	return &DepIOSEnrollmentProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DepIOSEnrollmentProfileRequest is request for DepIOSEnrollmentProfile
type DepIOSEnrollmentProfileRequest struct{ BaseRequest }

// Get performs GET request for DepIOSEnrollmentProfile
func (r *DepIOSEnrollmentProfileRequest) Get(ctx context.Context) (resObj *DepIOSEnrollmentProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DepIOSEnrollmentProfile
func (r *DepIOSEnrollmentProfileRequest) Update(ctx context.Context, reqObj *DepIOSEnrollmentProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DepIOSEnrollmentProfile
func (r *DepIOSEnrollmentProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DepMacOSEnrollmentProfileRequestBuilder is request builder for DepMacOSEnrollmentProfile
type DepMacOSEnrollmentProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns DepMacOSEnrollmentProfileRequest
func (b *DepMacOSEnrollmentProfileRequestBuilder) Request() *DepMacOSEnrollmentProfileRequest {
	return &DepMacOSEnrollmentProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DepMacOSEnrollmentProfileRequest is request for DepMacOSEnrollmentProfile
type DepMacOSEnrollmentProfileRequest struct{ BaseRequest }

// Get performs GET request for DepMacOSEnrollmentProfile
func (r *DepMacOSEnrollmentProfileRequest) Get(ctx context.Context) (resObj *DepMacOSEnrollmentProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DepMacOSEnrollmentProfile
func (r *DepMacOSEnrollmentProfileRequest) Update(ctx context.Context, reqObj *DepMacOSEnrollmentProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DepMacOSEnrollmentProfile
func (r *DepMacOSEnrollmentProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DepOnboardingSettingRequestBuilder is request builder for DepOnboardingSetting
type DepOnboardingSettingRequestBuilder struct{ BaseRequestBuilder }

// Request returns DepOnboardingSettingRequest
func (b *DepOnboardingSettingRequestBuilder) Request() *DepOnboardingSettingRequest {
	return &DepOnboardingSettingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DepOnboardingSettingRequest is request for DepOnboardingSetting
type DepOnboardingSettingRequest struct{ BaseRequest }

// Get performs GET request for DepOnboardingSetting
func (r *DepOnboardingSettingRequest) Get(ctx context.Context) (resObj *DepOnboardingSetting, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DepOnboardingSetting
func (r *DepOnboardingSettingRequest) Update(ctx context.Context, reqObj *DepOnboardingSetting) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DepOnboardingSetting
func (r *DepOnboardingSettingRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type DepOnboardingSettingGenerateEncryptionPublicKeyRequestBuilder struct{ BaseRequestBuilder }

// GenerateEncryptionPublicKey action undocumentedras
func (b *DepOnboardingSettingRequestBuilder) GenerateEncryptionPublicKey(reqObj *DepOnboardingSettingGenerateEncryptionPublicKeyRequestParameter) *DepOnboardingSettingGenerateEncryptionPublicKeyRequestBuilder {
	bb := &DepOnboardingSettingGenerateEncryptionPublicKeyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/GenerateEncryptionPublicKey"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DepOnboardingSettingGenerateEncryptionPublicKeyRequest struct{ BaseRequest }

func (b *DepOnboardingSettingGenerateEncryptionPublicKeyRequestBuilder) Request() *DepOnboardingSettingGenerateEncryptionPublicKeyRequest {
	return &DepOnboardingSettingGenerateEncryptionPublicKeyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DepOnboardingSettingGenerateEncryptionPublicKeyRequest) Post(ctx context.Context) (resObj *string, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type DepOnboardingSettingShareForSchoolDataSyncServiceRequestBuilder struct{ BaseRequestBuilder }

// ShareForSchoolDataSyncService action undocumentedrav
func (b *DepOnboardingSettingRequestBuilder) ShareForSchoolDataSyncService(reqObj *DepOnboardingSettingShareForSchoolDataSyncServiceRequestParameter) *DepOnboardingSettingShareForSchoolDataSyncServiceRequestBuilder {
	bb := &DepOnboardingSettingShareForSchoolDataSyncServiceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ShareForSchoolDataSyncService"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DepOnboardingSettingShareForSchoolDataSyncServiceRequest struct{ BaseRequest }

func (b *DepOnboardingSettingShareForSchoolDataSyncServiceRequestBuilder) Request() *DepOnboardingSettingShareForSchoolDataSyncServiceRequest {
	return &DepOnboardingSettingShareForSchoolDataSyncServiceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DepOnboardingSettingShareForSchoolDataSyncServiceRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type DepOnboardingSettingSyncWithAppleDeviceEnrollmentProgramRequestBuilder struct{ BaseRequestBuilder }

// SyncWithAppleDeviceEnrollmentProgram action undocumentedrav
func (b *DepOnboardingSettingRequestBuilder) SyncWithAppleDeviceEnrollmentProgram(reqObj *DepOnboardingSettingSyncWithAppleDeviceEnrollmentProgramRequestParameter) *DepOnboardingSettingSyncWithAppleDeviceEnrollmentProgramRequestBuilder {
	bb := &DepOnboardingSettingSyncWithAppleDeviceEnrollmentProgramRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/SyncWithAppleDeviceEnrollmentProgram"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DepOnboardingSettingSyncWithAppleDeviceEnrollmentProgramRequest struct{ BaseRequest }

func (b *DepOnboardingSettingSyncWithAppleDeviceEnrollmentProgramRequestBuilder) Request() *DepOnboardingSettingSyncWithAppleDeviceEnrollmentProgramRequest {
	return &DepOnboardingSettingSyncWithAppleDeviceEnrollmentProgramRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DepOnboardingSettingSyncWithAppleDeviceEnrollmentProgramRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type DepOnboardingSettingUnshareForSchoolDataSyncServiceRequestBuilder struct{ BaseRequestBuilder }

// UnshareForSchoolDataSyncService action undocumentedrav
func (b *DepOnboardingSettingRequestBuilder) UnshareForSchoolDataSyncService(reqObj *DepOnboardingSettingUnshareForSchoolDataSyncServiceRequestParameter) *DepOnboardingSettingUnshareForSchoolDataSyncServiceRequestBuilder {
	bb := &DepOnboardingSettingUnshareForSchoolDataSyncServiceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/UnshareForSchoolDataSyncService"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DepOnboardingSettingUnshareForSchoolDataSyncServiceRequest struct{ BaseRequest }

func (b *DepOnboardingSettingUnshareForSchoolDataSyncServiceRequestBuilder) Request() *DepOnboardingSettingUnshareForSchoolDataSyncServiceRequest {
	return &DepOnboardingSettingUnshareForSchoolDataSyncServiceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DepOnboardingSettingUnshareForSchoolDataSyncServiceRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type DepOnboardingSettingUploadDepTokenRequestBuilder struct{ BaseRequestBuilder }

// UploadDepToken action undocumentedrav
func (b *DepOnboardingSettingRequestBuilder) UploadDepToken(reqObj *DepOnboardingSettingUploadDepTokenRequestParameter) *DepOnboardingSettingUploadDepTokenRequestBuilder {
	bb := &DepOnboardingSettingUploadDepTokenRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/UploadDepToken"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DepOnboardingSettingUploadDepTokenRequest struct{ BaseRequest }

func (b *DepOnboardingSettingUploadDepTokenRequestBuilder) Request() *DepOnboardingSettingUploadDepTokenRequest {
	return &DepOnboardingSettingUploadDepTokenRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DepOnboardingSettingUploadDepTokenRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
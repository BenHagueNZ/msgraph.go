// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TenantAppManagementPolicyRequestBuilder is request builder for TenantAppManagementPolicy
type TenantAppManagementPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns TenantAppManagementPolicyRequest
func (b *TenantAppManagementPolicyRequestBuilder) Request() *TenantAppManagementPolicyRequest {
	return &TenantAppManagementPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TenantAppManagementPolicyRequest is request for TenantAppManagementPolicy
type TenantAppManagementPolicyRequest struct{ BaseRequest }

// Get performs GET request for TenantAppManagementPolicy
func (r *TenantAppManagementPolicyRequest) Get(ctx context.Context) (resObj *TenantAppManagementPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TenantAppManagementPolicy
func (r *TenantAppManagementPolicyRequest) Update(ctx context.Context, reqObj *TenantAppManagementPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TenantAppManagementPolicy
func (r *TenantAppManagementPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TenantAttachRBACRequestBuilder is request builder for TenantAttachRBAC
type TenantAttachRBACRequestBuilder struct{ BaseRequestBuilder }

// Request returns TenantAttachRBACRequest
func (b *TenantAttachRBACRequestBuilder) Request() *TenantAttachRBACRequest {
	return &TenantAttachRBACRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TenantAttachRBACRequest is request for TenantAttachRBAC
type TenantAttachRBACRequest struct{ BaseRequest }

// Get performs GET request for TenantAttachRBAC
func (r *TenantAttachRBACRequest) Get(ctx context.Context) (resObj *TenantAttachRBAC, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TenantAttachRBAC
func (r *TenantAttachRBACRequest) Update(ctx context.Context, reqObj *TenantAttachRBAC) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TenantAttachRBAC
func (r *TenantAttachRBACRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TenantAttachRBACStateRequestBuilder is request builder for TenantAttachRBACState
type TenantAttachRBACStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns TenantAttachRBACStateRequest
func (b *TenantAttachRBACStateRequestBuilder) Request() *TenantAttachRBACStateRequest {
	return &TenantAttachRBACStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TenantAttachRBACStateRequest is request for TenantAttachRBACState
type TenantAttachRBACStateRequest struct{ BaseRequest }

// Get performs GET request for TenantAttachRBACState
func (r *TenantAttachRBACStateRequest) Get(ctx context.Context) (resObj *TenantAttachRBACState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TenantAttachRBACState
func (r *TenantAttachRBACStateRequest) Update(ctx context.Context, reqObj *TenantAttachRBACState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TenantAttachRBACState
func (r *TenantAttachRBACStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TenantInformationRequestBuilder is request builder for TenantInformation
type TenantInformationRequestBuilder struct{ BaseRequestBuilder }

// Request returns TenantInformationRequest
func (b *TenantInformationRequestBuilder) Request() *TenantInformationRequest {
	return &TenantInformationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TenantInformationRequest is request for TenantInformation
type TenantInformationRequest struct{ BaseRequest }

// Get performs GET request for TenantInformation
func (r *TenantInformationRequest) Get(ctx context.Context) (resObj *TenantInformation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TenantInformation
func (r *TenantInformationRequest) Update(ctx context.Context, reqObj *TenantInformation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TenantInformation
func (r *TenantInformationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TenantReferenceRequestBuilder is request builder for TenantReference
type TenantReferenceRequestBuilder struct{ BaseRequestBuilder }

// Request returns TenantReferenceRequest
func (b *TenantReferenceRequestBuilder) Request() *TenantReferenceRequest {
	return &TenantReferenceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TenantReferenceRequest is request for TenantReference
type TenantReferenceRequest struct{ BaseRequest }

// Get performs GET request for TenantReference
func (r *TenantReferenceRequest) Get(ctx context.Context) (resObj *TenantReference, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TenantReference
func (r *TenantReferenceRequest) Update(ctx context.Context, reqObj *TenantReference) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TenantReference
func (r *TenantReferenceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TenantRelationshipRequestBuilder is request builder for TenantRelationship
type TenantRelationshipRequestBuilder struct{ BaseRequestBuilder }

// Request returns TenantRelationshipRequest
func (b *TenantRelationshipRequestBuilder) Request() *TenantRelationshipRequest {
	return &TenantRelationshipRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TenantRelationshipRequest is request for TenantRelationship
type TenantRelationshipRequest struct{ BaseRequest }

// Get performs GET request for TenantRelationship
func (r *TenantRelationshipRequest) Get(ctx context.Context) (resObj *TenantRelationship, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TenantRelationship
func (r *TenantRelationshipRequest) Update(ctx context.Context, reqObj *TenantRelationship) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TenantRelationship
func (r *TenantRelationshipRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TenantRelationshipAccessPolicyBaseRequestBuilder is request builder for TenantRelationshipAccessPolicyBase
type TenantRelationshipAccessPolicyBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns TenantRelationshipAccessPolicyBaseRequest
func (b *TenantRelationshipAccessPolicyBaseRequestBuilder) Request() *TenantRelationshipAccessPolicyBaseRequest {
	return &TenantRelationshipAccessPolicyBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TenantRelationshipAccessPolicyBaseRequest is request for TenantRelationshipAccessPolicyBase
type TenantRelationshipAccessPolicyBaseRequest struct{ BaseRequest }

// Get performs GET request for TenantRelationshipAccessPolicyBase
func (r *TenantRelationshipAccessPolicyBaseRequest) Get(ctx context.Context) (resObj *TenantRelationshipAccessPolicyBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TenantRelationshipAccessPolicyBase
func (r *TenantRelationshipAccessPolicyBaseRequest) Update(ctx context.Context, reqObj *TenantRelationshipAccessPolicyBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TenantRelationshipAccessPolicyBase
func (r *TenantRelationshipAccessPolicyBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TenantSetupInfoRequestBuilder is request builder for TenantSetupInfo
type TenantSetupInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns TenantSetupInfoRequest
func (b *TenantSetupInfoRequestBuilder) Request() *TenantSetupInfoRequest {
	return &TenantSetupInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TenantSetupInfoRequest is request for TenantSetupInfo
type TenantSetupInfoRequest struct{ BaseRequest }

// Get performs GET request for TenantSetupInfo
func (r *TenantSetupInfoRequest) Get(ctx context.Context) (resObj *TenantSetupInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TenantSetupInfo
func (r *TenantSetupInfoRequest) Update(ctx context.Context, reqObj *TenantSetupInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TenantSetupInfo
func (r *TenantSetupInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type TenantAttachRBACEnableRequestBuilder struct{ BaseRequestBuilder }

// Enable action undocumentedrav
func (b *TenantAttachRBACRequestBuilder) Enable(reqObj *TenantAttachRBACEnableRequestParameter) *TenantAttachRBACEnableRequestBuilder {
	bb := &TenantAttachRBACEnableRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Enable"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type TenantAttachRBACEnableRequest struct{ BaseRequest }

func (b *TenantAttachRBACEnableRequestBuilder) Request() *TenantAttachRBACEnableRequest {
	return &TenantAttachRBACEnableRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *TenantAttachRBACEnableRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type TenantReferenceRemovePersonalDataRequestBuilder struct{ BaseRequestBuilder }

// RemovePersonalData action undocumentedrav
func (b *TenantReferenceRequestBuilder) RemovePersonalData(reqObj *TenantReferenceRemovePersonalDataRequestParameter) *TenantReferenceRemovePersonalDataRequestBuilder {
	bb := &TenantReferenceRemovePersonalDataRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RemovePersonalData"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type TenantReferenceRemovePersonalDataRequest struct{ BaseRequest }

func (b *TenantReferenceRemovePersonalDataRequestBuilder) Request() *TenantReferenceRemovePersonalDataRequest {
	return &TenantReferenceRemovePersonalDataRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *TenantReferenceRemovePersonalDataRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

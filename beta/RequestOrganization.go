// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// OrganizationRequestBuilder is request builder for Organization
type OrganizationRequestBuilder struct{ BaseRequestBuilder }

// Request returns OrganizationRequest
func (b *OrganizationRequestBuilder) Request() *OrganizationRequest {
	return &OrganizationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OrganizationRequest is request for Organization
type OrganizationRequest struct{ BaseRequest }

// Get performs GET request for Organization
func (r *OrganizationRequest) Get(ctx context.Context) (resObj *Organization, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Organization
func (r *OrganizationRequest) Update(ctx context.Context, reqObj *Organization) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Organization
func (r *OrganizationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OrganizationSettingsRequestBuilder is request builder for OrganizationSettings
type OrganizationSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns OrganizationSettingsRequest
func (b *OrganizationSettingsRequestBuilder) Request() *OrganizationSettingsRequest {
	return &OrganizationSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OrganizationSettingsRequest is request for OrganizationSettings
type OrganizationSettingsRequest struct{ BaseRequest }

// Get performs GET request for OrganizationSettings
func (r *OrganizationSettingsRequest) Get(ctx context.Context) (resObj *OrganizationSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OrganizationSettings
func (r *OrganizationSettingsRequest) Update(ctx context.Context, reqObj *OrganizationSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OrganizationSettings
func (r *OrganizationSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type OrganizationSetMobileDeviceManagementAuthorityRequestBuilder struct{ BaseRequestBuilder }

// SetMobileDeviceManagementAuthority action undocumentedras
func (b *OrganizationRequestBuilder) SetMobileDeviceManagementAuthority(reqObj *OrganizationSetMobileDeviceManagementAuthorityRequestParameter) *OrganizationSetMobileDeviceManagementAuthorityRequestBuilder {
	bb := &OrganizationSetMobileDeviceManagementAuthorityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/SetMobileDeviceManagementAuthority"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type OrganizationSetMobileDeviceManagementAuthorityRequest struct{ BaseRequest }

func (b *OrganizationSetMobileDeviceManagementAuthorityRequestBuilder) Request() *OrganizationSetMobileDeviceManagementAuthorityRequest {
	return &OrganizationSetMobileDeviceManagementAuthorityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *OrganizationSetMobileDeviceManagementAuthorityRequest) Post(ctx context.Context) (resObj *int, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

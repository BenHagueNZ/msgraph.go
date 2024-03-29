// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ConditionalAccessAllExternalTenantsRequestBuilder is request builder for ConditionalAccessAllExternalTenants
type ConditionalAccessAllExternalTenantsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessAllExternalTenantsRequest
func (b *ConditionalAccessAllExternalTenantsRequestBuilder) Request() *ConditionalAccessAllExternalTenantsRequest {
	return &ConditionalAccessAllExternalTenantsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessAllExternalTenantsRequest is request for ConditionalAccessAllExternalTenants
type ConditionalAccessAllExternalTenantsRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessAllExternalTenants
func (r *ConditionalAccessAllExternalTenantsRequest) Get(ctx context.Context) (resObj *ConditionalAccessAllExternalTenants, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessAllExternalTenants
func (r *ConditionalAccessAllExternalTenantsRequest) Update(ctx context.Context, reqObj *ConditionalAccessAllExternalTenants) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessAllExternalTenants
func (r *ConditionalAccessAllExternalTenantsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessApplicationsRequestBuilder is request builder for ConditionalAccessApplications
type ConditionalAccessApplicationsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessApplicationsRequest
func (b *ConditionalAccessApplicationsRequestBuilder) Request() *ConditionalAccessApplicationsRequest {
	return &ConditionalAccessApplicationsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessApplicationsRequest is request for ConditionalAccessApplications
type ConditionalAccessApplicationsRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessApplications
func (r *ConditionalAccessApplicationsRequest) Get(ctx context.Context) (resObj *ConditionalAccessApplications, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessApplications
func (r *ConditionalAccessApplicationsRequest) Update(ctx context.Context, reqObj *ConditionalAccessApplications) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessApplications
func (r *ConditionalAccessApplicationsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessClientApplicationsRequestBuilder is request builder for ConditionalAccessClientApplications
type ConditionalAccessClientApplicationsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessClientApplicationsRequest
func (b *ConditionalAccessClientApplicationsRequestBuilder) Request() *ConditionalAccessClientApplicationsRequest {
	return &ConditionalAccessClientApplicationsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessClientApplicationsRequest is request for ConditionalAccessClientApplications
type ConditionalAccessClientApplicationsRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessClientApplications
func (r *ConditionalAccessClientApplicationsRequest) Get(ctx context.Context) (resObj *ConditionalAccessClientApplications, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessClientApplications
func (r *ConditionalAccessClientApplicationsRequest) Update(ctx context.Context, reqObj *ConditionalAccessClientApplications) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessClientApplications
func (r *ConditionalAccessClientApplicationsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessConditionSetRequestBuilder is request builder for ConditionalAccessConditionSet
type ConditionalAccessConditionSetRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessConditionSetRequest
func (b *ConditionalAccessConditionSetRequestBuilder) Request() *ConditionalAccessConditionSetRequest {
	return &ConditionalAccessConditionSetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessConditionSetRequest is request for ConditionalAccessConditionSet
type ConditionalAccessConditionSetRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessConditionSet
func (r *ConditionalAccessConditionSetRequest) Get(ctx context.Context) (resObj *ConditionalAccessConditionSet, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessConditionSet
func (r *ConditionalAccessConditionSetRequest) Update(ctx context.Context, reqObj *ConditionalAccessConditionSet) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessConditionSet
func (r *ConditionalAccessConditionSetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessDeviceStatesRequestBuilder is request builder for ConditionalAccessDeviceStates
type ConditionalAccessDeviceStatesRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessDeviceStatesRequest
func (b *ConditionalAccessDeviceStatesRequestBuilder) Request() *ConditionalAccessDeviceStatesRequest {
	return &ConditionalAccessDeviceStatesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessDeviceStatesRequest is request for ConditionalAccessDeviceStates
type ConditionalAccessDeviceStatesRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessDeviceStates
func (r *ConditionalAccessDeviceStatesRequest) Get(ctx context.Context) (resObj *ConditionalAccessDeviceStates, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessDeviceStates
func (r *ConditionalAccessDeviceStatesRequest) Update(ctx context.Context, reqObj *ConditionalAccessDeviceStates) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessDeviceStates
func (r *ConditionalAccessDeviceStatesRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessDevicesRequestBuilder is request builder for ConditionalAccessDevices
type ConditionalAccessDevicesRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessDevicesRequest
func (b *ConditionalAccessDevicesRequestBuilder) Request() *ConditionalAccessDevicesRequest {
	return &ConditionalAccessDevicesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessDevicesRequest is request for ConditionalAccessDevices
type ConditionalAccessDevicesRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessDevices
func (r *ConditionalAccessDevicesRequest) Get(ctx context.Context) (resObj *ConditionalAccessDevices, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessDevices
func (r *ConditionalAccessDevicesRequest) Update(ctx context.Context, reqObj *ConditionalAccessDevices) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessDevices
func (r *ConditionalAccessDevicesRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessEnumeratedExternalTenantsRequestBuilder is request builder for ConditionalAccessEnumeratedExternalTenants
type ConditionalAccessEnumeratedExternalTenantsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessEnumeratedExternalTenantsRequest
func (b *ConditionalAccessEnumeratedExternalTenantsRequestBuilder) Request() *ConditionalAccessEnumeratedExternalTenantsRequest {
	return &ConditionalAccessEnumeratedExternalTenantsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessEnumeratedExternalTenantsRequest is request for ConditionalAccessEnumeratedExternalTenants
type ConditionalAccessEnumeratedExternalTenantsRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessEnumeratedExternalTenants
func (r *ConditionalAccessEnumeratedExternalTenantsRequest) Get(ctx context.Context) (resObj *ConditionalAccessEnumeratedExternalTenants, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessEnumeratedExternalTenants
func (r *ConditionalAccessEnumeratedExternalTenantsRequest) Update(ctx context.Context, reqObj *ConditionalAccessEnumeratedExternalTenants) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessEnumeratedExternalTenants
func (r *ConditionalAccessEnumeratedExternalTenantsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessExternalTenantsRequestBuilder is request builder for ConditionalAccessExternalTenants
type ConditionalAccessExternalTenantsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessExternalTenantsRequest
func (b *ConditionalAccessExternalTenantsRequestBuilder) Request() *ConditionalAccessExternalTenantsRequest {
	return &ConditionalAccessExternalTenantsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessExternalTenantsRequest is request for ConditionalAccessExternalTenants
type ConditionalAccessExternalTenantsRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessExternalTenants
func (r *ConditionalAccessExternalTenantsRequest) Get(ctx context.Context) (resObj *ConditionalAccessExternalTenants, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessExternalTenants
func (r *ConditionalAccessExternalTenantsRequest) Update(ctx context.Context, reqObj *ConditionalAccessExternalTenants) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessExternalTenants
func (r *ConditionalAccessExternalTenantsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessFilterRequestBuilder is request builder for ConditionalAccessFilter
type ConditionalAccessFilterRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessFilterRequest
func (b *ConditionalAccessFilterRequestBuilder) Request() *ConditionalAccessFilterRequest {
	return &ConditionalAccessFilterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessFilterRequest is request for ConditionalAccessFilter
type ConditionalAccessFilterRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessFilter
func (r *ConditionalAccessFilterRequest) Get(ctx context.Context) (resObj *ConditionalAccessFilter, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessFilter
func (r *ConditionalAccessFilterRequest) Update(ctx context.Context, reqObj *ConditionalAccessFilter) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessFilter
func (r *ConditionalAccessFilterRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessGrantControlsRequestBuilder is request builder for ConditionalAccessGrantControls
type ConditionalAccessGrantControlsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessGrantControlsRequest
func (b *ConditionalAccessGrantControlsRequestBuilder) Request() *ConditionalAccessGrantControlsRequest {
	return &ConditionalAccessGrantControlsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessGrantControlsRequest is request for ConditionalAccessGrantControls
type ConditionalAccessGrantControlsRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessGrantControls
func (r *ConditionalAccessGrantControlsRequest) Get(ctx context.Context) (resObj *ConditionalAccessGrantControls, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessGrantControls
func (r *ConditionalAccessGrantControlsRequest) Update(ctx context.Context, reqObj *ConditionalAccessGrantControls) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessGrantControls
func (r *ConditionalAccessGrantControlsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessGuestsOrExternalUsersRequestBuilder is request builder for ConditionalAccessGuestsOrExternalUsers
type ConditionalAccessGuestsOrExternalUsersRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessGuestsOrExternalUsersRequest
func (b *ConditionalAccessGuestsOrExternalUsersRequestBuilder) Request() *ConditionalAccessGuestsOrExternalUsersRequest {
	return &ConditionalAccessGuestsOrExternalUsersRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessGuestsOrExternalUsersRequest is request for ConditionalAccessGuestsOrExternalUsers
type ConditionalAccessGuestsOrExternalUsersRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessGuestsOrExternalUsers
func (r *ConditionalAccessGuestsOrExternalUsersRequest) Get(ctx context.Context) (resObj *ConditionalAccessGuestsOrExternalUsers, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessGuestsOrExternalUsers
func (r *ConditionalAccessGuestsOrExternalUsersRequest) Update(ctx context.Context, reqObj *ConditionalAccessGuestsOrExternalUsers) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessGuestsOrExternalUsers
func (r *ConditionalAccessGuestsOrExternalUsersRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessLocationsRequestBuilder is request builder for ConditionalAccessLocations
type ConditionalAccessLocationsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessLocationsRequest
func (b *ConditionalAccessLocationsRequestBuilder) Request() *ConditionalAccessLocationsRequest {
	return &ConditionalAccessLocationsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessLocationsRequest is request for ConditionalAccessLocations
type ConditionalAccessLocationsRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessLocations
func (r *ConditionalAccessLocationsRequest) Get(ctx context.Context) (resObj *ConditionalAccessLocations, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessLocations
func (r *ConditionalAccessLocationsRequest) Update(ctx context.Context, reqObj *ConditionalAccessLocations) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessLocations
func (r *ConditionalAccessLocationsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessPlatformsRequestBuilder is request builder for ConditionalAccessPlatforms
type ConditionalAccessPlatformsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessPlatformsRequest
func (b *ConditionalAccessPlatformsRequestBuilder) Request() *ConditionalAccessPlatformsRequest {
	return &ConditionalAccessPlatformsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessPlatformsRequest is request for ConditionalAccessPlatforms
type ConditionalAccessPlatformsRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessPlatforms
func (r *ConditionalAccessPlatformsRequest) Get(ctx context.Context) (resObj *ConditionalAccessPlatforms, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessPlatforms
func (r *ConditionalAccessPlatformsRequest) Update(ctx context.Context, reqObj *ConditionalAccessPlatforms) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessPlatforms
func (r *ConditionalAccessPlatformsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessPolicyRequestBuilder is request builder for ConditionalAccessPolicy
type ConditionalAccessPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessPolicyRequest
func (b *ConditionalAccessPolicyRequestBuilder) Request() *ConditionalAccessPolicyRequest {
	return &ConditionalAccessPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessPolicyRequest is request for ConditionalAccessPolicy
type ConditionalAccessPolicyRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessPolicy
func (r *ConditionalAccessPolicyRequest) Get(ctx context.Context) (resObj *ConditionalAccessPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessPolicy
func (r *ConditionalAccessPolicyRequest) Update(ctx context.Context, reqObj *ConditionalAccessPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessPolicy
func (r *ConditionalAccessPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessPolicyDetailRequestBuilder is request builder for ConditionalAccessPolicyDetail
type ConditionalAccessPolicyDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessPolicyDetailRequest
func (b *ConditionalAccessPolicyDetailRequestBuilder) Request() *ConditionalAccessPolicyDetailRequest {
	return &ConditionalAccessPolicyDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessPolicyDetailRequest is request for ConditionalAccessPolicyDetail
type ConditionalAccessPolicyDetailRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessPolicyDetail
func (r *ConditionalAccessPolicyDetailRequest) Get(ctx context.Context) (resObj *ConditionalAccessPolicyDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessPolicyDetail
func (r *ConditionalAccessPolicyDetailRequest) Update(ctx context.Context, reqObj *ConditionalAccessPolicyDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessPolicyDetail
func (r *ConditionalAccessPolicyDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessRootRequestBuilder is request builder for ConditionalAccessRoot
type ConditionalAccessRootRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessRootRequest
func (b *ConditionalAccessRootRequestBuilder) Request() *ConditionalAccessRootRequest {
	return &ConditionalAccessRootRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessRootRequest is request for ConditionalAccessRoot
type ConditionalAccessRootRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessRoot
func (r *ConditionalAccessRootRequest) Get(ctx context.Context) (resObj *ConditionalAccessRoot, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessRoot
func (r *ConditionalAccessRootRequest) Update(ctx context.Context, reqObj *ConditionalAccessRoot) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessRoot
func (r *ConditionalAccessRootRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessRuleSatisfiedRequestBuilder is request builder for ConditionalAccessRuleSatisfied
type ConditionalAccessRuleSatisfiedRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessRuleSatisfiedRequest
func (b *ConditionalAccessRuleSatisfiedRequestBuilder) Request() *ConditionalAccessRuleSatisfiedRequest {
	return &ConditionalAccessRuleSatisfiedRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessRuleSatisfiedRequest is request for ConditionalAccessRuleSatisfied
type ConditionalAccessRuleSatisfiedRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessRuleSatisfied
func (r *ConditionalAccessRuleSatisfiedRequest) Get(ctx context.Context) (resObj *ConditionalAccessRuleSatisfied, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessRuleSatisfied
func (r *ConditionalAccessRuleSatisfiedRequest) Update(ctx context.Context, reqObj *ConditionalAccessRuleSatisfied) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessRuleSatisfied
func (r *ConditionalAccessRuleSatisfiedRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessSessionControlRequestBuilder is request builder for ConditionalAccessSessionControl
type ConditionalAccessSessionControlRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessSessionControlRequest
func (b *ConditionalAccessSessionControlRequestBuilder) Request() *ConditionalAccessSessionControlRequest {
	return &ConditionalAccessSessionControlRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessSessionControlRequest is request for ConditionalAccessSessionControl
type ConditionalAccessSessionControlRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessSessionControl
func (r *ConditionalAccessSessionControlRequest) Get(ctx context.Context) (resObj *ConditionalAccessSessionControl, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessSessionControl
func (r *ConditionalAccessSessionControlRequest) Update(ctx context.Context, reqObj *ConditionalAccessSessionControl) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessSessionControl
func (r *ConditionalAccessSessionControlRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessSessionControlsRequestBuilder is request builder for ConditionalAccessSessionControls
type ConditionalAccessSessionControlsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessSessionControlsRequest
func (b *ConditionalAccessSessionControlsRequestBuilder) Request() *ConditionalAccessSessionControlsRequest {
	return &ConditionalAccessSessionControlsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessSessionControlsRequest is request for ConditionalAccessSessionControls
type ConditionalAccessSessionControlsRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessSessionControls
func (r *ConditionalAccessSessionControlsRequest) Get(ctx context.Context) (resObj *ConditionalAccessSessionControls, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessSessionControls
func (r *ConditionalAccessSessionControlsRequest) Update(ctx context.Context, reqObj *ConditionalAccessSessionControls) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessSessionControls
func (r *ConditionalAccessSessionControlsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessTemplateRequestBuilder is request builder for ConditionalAccessTemplate
type ConditionalAccessTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessTemplateRequest
func (b *ConditionalAccessTemplateRequestBuilder) Request() *ConditionalAccessTemplateRequest {
	return &ConditionalAccessTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessTemplateRequest is request for ConditionalAccessTemplate
type ConditionalAccessTemplateRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessTemplate
func (r *ConditionalAccessTemplateRequest) Get(ctx context.Context) (resObj *ConditionalAccessTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessTemplate
func (r *ConditionalAccessTemplateRequest) Update(ctx context.Context, reqObj *ConditionalAccessTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessTemplate
func (r *ConditionalAccessTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConditionalAccessUsersRequestBuilder is request builder for ConditionalAccessUsers
type ConditionalAccessUsersRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConditionalAccessUsersRequest
func (b *ConditionalAccessUsersRequestBuilder) Request() *ConditionalAccessUsersRequest {
	return &ConditionalAccessUsersRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConditionalAccessUsersRequest is request for ConditionalAccessUsers
type ConditionalAccessUsersRequest struct{ BaseRequest }

// Get performs GET request for ConditionalAccessUsers
func (r *ConditionalAccessUsersRequest) Get(ctx context.Context) (resObj *ConditionalAccessUsers, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConditionalAccessUsers
func (r *ConditionalAccessUsersRequest) Update(ctx context.Context, reqObj *ConditionalAccessUsers) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConditionalAccessUsers
func (r *ConditionalAccessUsersRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

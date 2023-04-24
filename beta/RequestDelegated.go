// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// DelegatedAdminAccessAssignmentRequestBuilder is request builder for DelegatedAdminAccessAssignment
type DelegatedAdminAccessAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns DelegatedAdminAccessAssignmentRequest
func (b *DelegatedAdminAccessAssignmentRequestBuilder) Request() *DelegatedAdminAccessAssignmentRequest {
	return &DelegatedAdminAccessAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DelegatedAdminAccessAssignmentRequest is request for DelegatedAdminAccessAssignment
type DelegatedAdminAccessAssignmentRequest struct{ BaseRequest }

// Get performs GET request for DelegatedAdminAccessAssignment
func (r *DelegatedAdminAccessAssignmentRequest) Get(ctx context.Context) (resObj *DelegatedAdminAccessAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DelegatedAdminAccessAssignment
func (r *DelegatedAdminAccessAssignmentRequest) Update(ctx context.Context, reqObj *DelegatedAdminAccessAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DelegatedAdminAccessAssignment
func (r *DelegatedAdminAccessAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DelegatedAdminAccessContainerRequestBuilder is request builder for DelegatedAdminAccessContainer
type DelegatedAdminAccessContainerRequestBuilder struct{ BaseRequestBuilder }

// Request returns DelegatedAdminAccessContainerRequest
func (b *DelegatedAdminAccessContainerRequestBuilder) Request() *DelegatedAdminAccessContainerRequest {
	return &DelegatedAdminAccessContainerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DelegatedAdminAccessContainerRequest is request for DelegatedAdminAccessContainer
type DelegatedAdminAccessContainerRequest struct{ BaseRequest }

// Get performs GET request for DelegatedAdminAccessContainer
func (r *DelegatedAdminAccessContainerRequest) Get(ctx context.Context) (resObj *DelegatedAdminAccessContainer, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DelegatedAdminAccessContainer
func (r *DelegatedAdminAccessContainerRequest) Update(ctx context.Context, reqObj *DelegatedAdminAccessContainer) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DelegatedAdminAccessContainer
func (r *DelegatedAdminAccessContainerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DelegatedAdminAccessDetailsRequestBuilder is request builder for DelegatedAdminAccessDetails
type DelegatedAdminAccessDetailsRequestBuilder struct{ BaseRequestBuilder }

// Request returns DelegatedAdminAccessDetailsRequest
func (b *DelegatedAdminAccessDetailsRequestBuilder) Request() *DelegatedAdminAccessDetailsRequest {
	return &DelegatedAdminAccessDetailsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DelegatedAdminAccessDetailsRequest is request for DelegatedAdminAccessDetails
type DelegatedAdminAccessDetailsRequest struct{ BaseRequest }

// Get performs GET request for DelegatedAdminAccessDetails
func (r *DelegatedAdminAccessDetailsRequest) Get(ctx context.Context) (resObj *DelegatedAdminAccessDetails, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DelegatedAdminAccessDetails
func (r *DelegatedAdminAccessDetailsRequest) Update(ctx context.Context, reqObj *DelegatedAdminAccessDetails) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DelegatedAdminAccessDetails
func (r *DelegatedAdminAccessDetailsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DelegatedAdminCustomerRequestBuilder is request builder for DelegatedAdminCustomer
type DelegatedAdminCustomerRequestBuilder struct{ BaseRequestBuilder }

// Request returns DelegatedAdminCustomerRequest
func (b *DelegatedAdminCustomerRequestBuilder) Request() *DelegatedAdminCustomerRequest {
	return &DelegatedAdminCustomerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DelegatedAdminCustomerRequest is request for DelegatedAdminCustomer
type DelegatedAdminCustomerRequest struct{ BaseRequest }

// Get performs GET request for DelegatedAdminCustomer
func (r *DelegatedAdminCustomerRequest) Get(ctx context.Context) (resObj *DelegatedAdminCustomer, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DelegatedAdminCustomer
func (r *DelegatedAdminCustomerRequest) Update(ctx context.Context, reqObj *DelegatedAdminCustomer) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DelegatedAdminCustomer
func (r *DelegatedAdminCustomerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DelegatedAdminRelationshipRequestBuilder is request builder for DelegatedAdminRelationship
type DelegatedAdminRelationshipRequestBuilder struct{ BaseRequestBuilder }

// Request returns DelegatedAdminRelationshipRequest
func (b *DelegatedAdminRelationshipRequestBuilder) Request() *DelegatedAdminRelationshipRequest {
	return &DelegatedAdminRelationshipRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DelegatedAdminRelationshipRequest is request for DelegatedAdminRelationship
type DelegatedAdminRelationshipRequest struct{ BaseRequest }

// Get performs GET request for DelegatedAdminRelationship
func (r *DelegatedAdminRelationshipRequest) Get(ctx context.Context) (resObj *DelegatedAdminRelationship, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DelegatedAdminRelationship
func (r *DelegatedAdminRelationshipRequest) Update(ctx context.Context, reqObj *DelegatedAdminRelationship) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DelegatedAdminRelationship
func (r *DelegatedAdminRelationshipRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DelegatedAdminRelationshipCustomerParticipantRequestBuilder is request builder for DelegatedAdminRelationshipCustomerParticipant
type DelegatedAdminRelationshipCustomerParticipantRequestBuilder struct{ BaseRequestBuilder }

// Request returns DelegatedAdminRelationshipCustomerParticipantRequest
func (b *DelegatedAdminRelationshipCustomerParticipantRequestBuilder) Request() *DelegatedAdminRelationshipCustomerParticipantRequest {
	return &DelegatedAdminRelationshipCustomerParticipantRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DelegatedAdminRelationshipCustomerParticipantRequest is request for DelegatedAdminRelationshipCustomerParticipant
type DelegatedAdminRelationshipCustomerParticipantRequest struct{ BaseRequest }

// Get performs GET request for DelegatedAdminRelationshipCustomerParticipant
func (r *DelegatedAdminRelationshipCustomerParticipantRequest) Get(ctx context.Context) (resObj *DelegatedAdminRelationshipCustomerParticipant, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DelegatedAdminRelationshipCustomerParticipant
func (r *DelegatedAdminRelationshipCustomerParticipantRequest) Update(ctx context.Context, reqObj *DelegatedAdminRelationshipCustomerParticipant) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DelegatedAdminRelationshipCustomerParticipant
func (r *DelegatedAdminRelationshipCustomerParticipantRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DelegatedAdminRelationshipOperationRequestBuilder is request builder for DelegatedAdminRelationshipOperation
type DelegatedAdminRelationshipOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns DelegatedAdminRelationshipOperationRequest
func (b *DelegatedAdminRelationshipOperationRequestBuilder) Request() *DelegatedAdminRelationshipOperationRequest {
	return &DelegatedAdminRelationshipOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DelegatedAdminRelationshipOperationRequest is request for DelegatedAdminRelationshipOperation
type DelegatedAdminRelationshipOperationRequest struct{ BaseRequest }

// Get performs GET request for DelegatedAdminRelationshipOperation
func (r *DelegatedAdminRelationshipOperationRequest) Get(ctx context.Context) (resObj *DelegatedAdminRelationshipOperation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DelegatedAdminRelationshipOperation
func (r *DelegatedAdminRelationshipOperationRequest) Update(ctx context.Context, reqObj *DelegatedAdminRelationshipOperation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DelegatedAdminRelationshipOperation
func (r *DelegatedAdminRelationshipOperationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DelegatedAdminRelationshipRequestObjectRequestBuilder is request builder for DelegatedAdminRelationshipRequestObject
type DelegatedAdminRelationshipRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns DelegatedAdminRelationshipRequestObjectRequest
func (b *DelegatedAdminRelationshipRequestObjectRequestBuilder) Request() *DelegatedAdminRelationshipRequestObjectRequest {
	return &DelegatedAdminRelationshipRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DelegatedAdminRelationshipRequestObjectRequest is request for DelegatedAdminRelationshipRequestObject
type DelegatedAdminRelationshipRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for DelegatedAdminRelationshipRequestObject
func (r *DelegatedAdminRelationshipRequestObjectRequest) Get(ctx context.Context) (resObj *DelegatedAdminRelationshipRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DelegatedAdminRelationshipRequestObject
func (r *DelegatedAdminRelationshipRequestObjectRequest) Update(ctx context.Context, reqObj *DelegatedAdminRelationshipRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DelegatedAdminRelationshipRequestObject
func (r *DelegatedAdminRelationshipRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DelegatedAdminServiceManagementDetailRequestBuilder is request builder for DelegatedAdminServiceManagementDetail
type DelegatedAdminServiceManagementDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns DelegatedAdminServiceManagementDetailRequest
func (b *DelegatedAdminServiceManagementDetailRequestBuilder) Request() *DelegatedAdminServiceManagementDetailRequest {
	return &DelegatedAdminServiceManagementDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DelegatedAdminServiceManagementDetailRequest is request for DelegatedAdminServiceManagementDetail
type DelegatedAdminServiceManagementDetailRequest struct{ BaseRequest }

// Get performs GET request for DelegatedAdminServiceManagementDetail
func (r *DelegatedAdminServiceManagementDetailRequest) Get(ctx context.Context) (resObj *DelegatedAdminServiceManagementDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DelegatedAdminServiceManagementDetail
func (r *DelegatedAdminServiceManagementDetailRequest) Update(ctx context.Context, reqObj *DelegatedAdminServiceManagementDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DelegatedAdminServiceManagementDetail
func (r *DelegatedAdminServiceManagementDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DelegatedPermissionClassificationRequestBuilder is request builder for DelegatedPermissionClassification
type DelegatedPermissionClassificationRequestBuilder struct{ BaseRequestBuilder }

// Request returns DelegatedPermissionClassificationRequest
func (b *DelegatedPermissionClassificationRequestBuilder) Request() *DelegatedPermissionClassificationRequest {
	return &DelegatedPermissionClassificationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DelegatedPermissionClassificationRequest is request for DelegatedPermissionClassification
type DelegatedPermissionClassificationRequest struct{ BaseRequest }

// Get performs GET request for DelegatedPermissionClassification
func (r *DelegatedPermissionClassificationRequest) Get(ctx context.Context) (resObj *DelegatedPermissionClassification, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DelegatedPermissionClassification
func (r *DelegatedPermissionClassificationRequest) Update(ctx context.Context, reqObj *DelegatedPermissionClassification) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DelegatedPermissionClassification
func (r *DelegatedPermissionClassificationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// RoleAssignmentRequestBuilder is request builder for RoleAssignment
type RoleAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns RoleAssignmentRequest
func (b *RoleAssignmentRequestBuilder) Request() *RoleAssignmentRequest {
	return &RoleAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RoleAssignmentRequest is request for RoleAssignment
type RoleAssignmentRequest struct{ BaseRequest }

// Get performs GET request for RoleAssignment
func (r *RoleAssignmentRequest) Get(ctx context.Context) (resObj *RoleAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RoleAssignment
func (r *RoleAssignmentRequest) Update(ctx context.Context, reqObj *RoleAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RoleAssignment
func (r *RoleAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RoleDefinitionRequestBuilder is request builder for RoleDefinition
type RoleDefinitionRequestBuilder struct{ BaseRequestBuilder }

// Request returns RoleDefinitionRequest
func (b *RoleDefinitionRequestBuilder) Request() *RoleDefinitionRequest {
	return &RoleDefinitionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RoleDefinitionRequest is request for RoleDefinition
type RoleDefinitionRequest struct{ BaseRequest }

// Get performs GET request for RoleDefinition
func (r *RoleDefinitionRequest) Get(ctx context.Context) (resObj *RoleDefinition, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RoleDefinition
func (r *RoleDefinitionRequest) Update(ctx context.Context, reqObj *RoleDefinition) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RoleDefinition
func (r *RoleDefinitionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RoleManagementRequestBuilder is request builder for RoleManagement
type RoleManagementRequestBuilder struct{ BaseRequestBuilder }

// Request returns RoleManagementRequest
func (b *RoleManagementRequestBuilder) Request() *RoleManagementRequest {
	return &RoleManagementRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RoleManagementRequest is request for RoleManagement
type RoleManagementRequest struct{ BaseRequest }

// Get performs GET request for RoleManagement
func (r *RoleManagementRequest) Get(ctx context.Context) (resObj *RoleManagement, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RoleManagement
func (r *RoleManagementRequest) Update(ctx context.Context, reqObj *RoleManagement) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RoleManagement
func (r *RoleManagementRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RolePermissionRequestBuilder is request builder for RolePermission
type RolePermissionRequestBuilder struct{ BaseRequestBuilder }

// Request returns RolePermissionRequest
func (b *RolePermissionRequestBuilder) Request() *RolePermissionRequest {
	return &RolePermissionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RolePermissionRequest is request for RolePermission
type RolePermissionRequest struct{ BaseRequest }

// Get performs GET request for RolePermission
func (r *RolePermissionRequest) Get(ctx context.Context) (resObj *RolePermission, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RolePermission
func (r *RolePermissionRequest) Update(ctx context.Context, reqObj *RolePermission) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RolePermission
func (r *RolePermissionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

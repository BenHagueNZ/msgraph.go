// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// EmployeeRequestBuilder is request builder for Employee
type EmployeeRequestBuilder struct{ BaseRequestBuilder }

// Request returns EmployeeRequest
func (b *EmployeeRequestBuilder) Request() *EmployeeRequest {
	return &EmployeeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EmployeeRequest is request for Employee
type EmployeeRequest struct{ BaseRequest }

// Get performs GET request for Employee
func (r *EmployeeRequest) Get(ctx context.Context) (resObj *Employee, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Employee
func (r *EmployeeRequest) Update(ctx context.Context, reqObj *Employee) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Employee
func (r *EmployeeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// EmployeeExperienceRequestBuilder is request builder for EmployeeExperience
type EmployeeExperienceRequestBuilder struct{ BaseRequestBuilder }

// Request returns EmployeeExperienceRequest
func (b *EmployeeExperienceRequestBuilder) Request() *EmployeeExperienceRequest {
	return &EmployeeExperienceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EmployeeExperienceRequest is request for EmployeeExperience
type EmployeeExperienceRequest struct{ BaseRequest }

// Get performs GET request for EmployeeExperience
func (r *EmployeeExperienceRequest) Get(ctx context.Context) (resObj *EmployeeExperience, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EmployeeExperience
func (r *EmployeeExperienceRequest) Update(ctx context.Context, reqObj *EmployeeExperience) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EmployeeExperience
func (r *EmployeeExperienceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// EmployeeOrgDataRequestBuilder is request builder for EmployeeOrgData
type EmployeeOrgDataRequestBuilder struct{ BaseRequestBuilder }

// Request returns EmployeeOrgDataRequest
func (b *EmployeeOrgDataRequestBuilder) Request() *EmployeeOrgDataRequest {
	return &EmployeeOrgDataRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EmployeeOrgDataRequest is request for EmployeeOrgData
type EmployeeOrgDataRequest struct{ BaseRequest }

// Get performs GET request for EmployeeOrgData
func (r *EmployeeOrgDataRequest) Get(ctx context.Context) (resObj *EmployeeOrgData, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EmployeeOrgData
func (r *EmployeeOrgDataRequest) Update(ctx context.Context, reqObj *EmployeeOrgData) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EmployeeOrgData
func (r *EmployeeOrgDataRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

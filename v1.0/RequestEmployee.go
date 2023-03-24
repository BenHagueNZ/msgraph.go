// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

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
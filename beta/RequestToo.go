// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TooManyGlobalAdminsAssignedToTenantAlertConfigurationRequestBuilder is request builder for TooManyGlobalAdminsAssignedToTenantAlertConfiguration
type TooManyGlobalAdminsAssignedToTenantAlertConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns TooManyGlobalAdminsAssignedToTenantAlertConfigurationRequest
func (b *TooManyGlobalAdminsAssignedToTenantAlertConfigurationRequestBuilder) Request() *TooManyGlobalAdminsAssignedToTenantAlertConfigurationRequest {
	return &TooManyGlobalAdminsAssignedToTenantAlertConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TooManyGlobalAdminsAssignedToTenantAlertConfigurationRequest is request for TooManyGlobalAdminsAssignedToTenantAlertConfiguration
type TooManyGlobalAdminsAssignedToTenantAlertConfigurationRequest struct{ BaseRequest }

// Get performs GET request for TooManyGlobalAdminsAssignedToTenantAlertConfiguration
func (r *TooManyGlobalAdminsAssignedToTenantAlertConfigurationRequest) Get(ctx context.Context) (resObj *TooManyGlobalAdminsAssignedToTenantAlertConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TooManyGlobalAdminsAssignedToTenantAlertConfiguration
func (r *TooManyGlobalAdminsAssignedToTenantAlertConfigurationRequest) Update(ctx context.Context, reqObj *TooManyGlobalAdminsAssignedToTenantAlertConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TooManyGlobalAdminsAssignedToTenantAlertConfiguration
func (r *TooManyGlobalAdminsAssignedToTenantAlertConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TooManyGlobalAdminsAssignedToTenantAlertIncidentRequestBuilder is request builder for TooManyGlobalAdminsAssignedToTenantAlertIncident
type TooManyGlobalAdminsAssignedToTenantAlertIncidentRequestBuilder struct{ BaseRequestBuilder }

// Request returns TooManyGlobalAdminsAssignedToTenantAlertIncidentRequest
func (b *TooManyGlobalAdminsAssignedToTenantAlertIncidentRequestBuilder) Request() *TooManyGlobalAdminsAssignedToTenantAlertIncidentRequest {
	return &TooManyGlobalAdminsAssignedToTenantAlertIncidentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TooManyGlobalAdminsAssignedToTenantAlertIncidentRequest is request for TooManyGlobalAdminsAssignedToTenantAlertIncident
type TooManyGlobalAdminsAssignedToTenantAlertIncidentRequest struct{ BaseRequest }

// Get performs GET request for TooManyGlobalAdminsAssignedToTenantAlertIncident
func (r *TooManyGlobalAdminsAssignedToTenantAlertIncidentRequest) Get(ctx context.Context) (resObj *TooManyGlobalAdminsAssignedToTenantAlertIncident, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TooManyGlobalAdminsAssignedToTenantAlertIncident
func (r *TooManyGlobalAdminsAssignedToTenantAlertIncidentRequest) Update(ctx context.Context, reqObj *TooManyGlobalAdminsAssignedToTenantAlertIncident) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TooManyGlobalAdminsAssignedToTenantAlertIncident
func (r *TooManyGlobalAdminsAssignedToTenantAlertIncidentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// NoMFAOnRoleActivationAlertConfigurationRequestBuilder is request builder for NoMFAOnRoleActivationAlertConfiguration
type NoMFAOnRoleActivationAlertConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns NoMFAOnRoleActivationAlertConfigurationRequest
func (b *NoMFAOnRoleActivationAlertConfigurationRequestBuilder) Request() *NoMFAOnRoleActivationAlertConfigurationRequest {
	return &NoMFAOnRoleActivationAlertConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// NoMFAOnRoleActivationAlertConfigurationRequest is request for NoMFAOnRoleActivationAlertConfiguration
type NoMFAOnRoleActivationAlertConfigurationRequest struct{ BaseRequest }

// Get performs GET request for NoMFAOnRoleActivationAlertConfiguration
func (r *NoMFAOnRoleActivationAlertConfigurationRequest) Get(ctx context.Context) (resObj *NoMFAOnRoleActivationAlertConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for NoMFAOnRoleActivationAlertConfiguration
func (r *NoMFAOnRoleActivationAlertConfigurationRequest) Update(ctx context.Context, reqObj *NoMFAOnRoleActivationAlertConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for NoMFAOnRoleActivationAlertConfiguration
func (r *NoMFAOnRoleActivationAlertConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// NoMFAOnRoleActivationAlertIncidentRequestBuilder is request builder for NoMFAOnRoleActivationAlertIncident
type NoMFAOnRoleActivationAlertIncidentRequestBuilder struct{ BaseRequestBuilder }

// Request returns NoMFAOnRoleActivationAlertIncidentRequest
func (b *NoMFAOnRoleActivationAlertIncidentRequestBuilder) Request() *NoMFAOnRoleActivationAlertIncidentRequest {
	return &NoMFAOnRoleActivationAlertIncidentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// NoMFAOnRoleActivationAlertIncidentRequest is request for NoMFAOnRoleActivationAlertIncident
type NoMFAOnRoleActivationAlertIncidentRequest struct{ BaseRequest }

// Get performs GET request for NoMFAOnRoleActivationAlertIncident
func (r *NoMFAOnRoleActivationAlertIncidentRequest) Get(ctx context.Context) (resObj *NoMFAOnRoleActivationAlertIncident, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for NoMFAOnRoleActivationAlertIncident
func (r *NoMFAOnRoleActivationAlertIncidentRequest) Update(ctx context.Context, reqObj *NoMFAOnRoleActivationAlertIncident) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for NoMFAOnRoleActivationAlertIncident
func (r *NoMFAOnRoleActivationAlertIncidentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
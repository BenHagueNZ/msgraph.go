// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// RedundantAssignmentAlertConfigurationRequestBuilder is request builder for RedundantAssignmentAlertConfiguration
type RedundantAssignmentAlertConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns RedundantAssignmentAlertConfigurationRequest
func (b *RedundantAssignmentAlertConfigurationRequestBuilder) Request() *RedundantAssignmentAlertConfigurationRequest {
	return &RedundantAssignmentAlertConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RedundantAssignmentAlertConfigurationRequest is request for RedundantAssignmentAlertConfiguration
type RedundantAssignmentAlertConfigurationRequest struct{ BaseRequest }

// Get performs GET request for RedundantAssignmentAlertConfiguration
func (r *RedundantAssignmentAlertConfigurationRequest) Get(ctx context.Context) (resObj *RedundantAssignmentAlertConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RedundantAssignmentAlertConfiguration
func (r *RedundantAssignmentAlertConfigurationRequest) Update(ctx context.Context, reqObj *RedundantAssignmentAlertConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RedundantAssignmentAlertConfiguration
func (r *RedundantAssignmentAlertConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RedundantAssignmentAlertIncidentRequestBuilder is request builder for RedundantAssignmentAlertIncident
type RedundantAssignmentAlertIncidentRequestBuilder struct{ BaseRequestBuilder }

// Request returns RedundantAssignmentAlertIncidentRequest
func (b *RedundantAssignmentAlertIncidentRequestBuilder) Request() *RedundantAssignmentAlertIncidentRequest {
	return &RedundantAssignmentAlertIncidentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RedundantAssignmentAlertIncidentRequest is request for RedundantAssignmentAlertIncident
type RedundantAssignmentAlertIncidentRequest struct{ BaseRequest }

// Get performs GET request for RedundantAssignmentAlertIncident
func (r *RedundantAssignmentAlertIncidentRequest) Get(ctx context.Context) (resObj *RedundantAssignmentAlertIncident, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RedundantAssignmentAlertIncident
func (r *RedundantAssignmentAlertIncidentRequest) Update(ctx context.Context, reqObj *RedundantAssignmentAlertIncident) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RedundantAssignmentAlertIncident
func (r *RedundantAssignmentAlertIncidentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

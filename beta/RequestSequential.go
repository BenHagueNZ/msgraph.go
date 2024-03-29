// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SequentialActivationRenewalsAlertConfigurationRequestBuilder is request builder for SequentialActivationRenewalsAlertConfiguration
type SequentialActivationRenewalsAlertConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns SequentialActivationRenewalsAlertConfigurationRequest
func (b *SequentialActivationRenewalsAlertConfigurationRequestBuilder) Request() *SequentialActivationRenewalsAlertConfigurationRequest {
	return &SequentialActivationRenewalsAlertConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SequentialActivationRenewalsAlertConfigurationRequest is request for SequentialActivationRenewalsAlertConfiguration
type SequentialActivationRenewalsAlertConfigurationRequest struct{ BaseRequest }

// Get performs GET request for SequentialActivationRenewalsAlertConfiguration
func (r *SequentialActivationRenewalsAlertConfigurationRequest) Get(ctx context.Context) (resObj *SequentialActivationRenewalsAlertConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SequentialActivationRenewalsAlertConfiguration
func (r *SequentialActivationRenewalsAlertConfigurationRequest) Update(ctx context.Context, reqObj *SequentialActivationRenewalsAlertConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SequentialActivationRenewalsAlertConfiguration
func (r *SequentialActivationRenewalsAlertConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SequentialActivationRenewalsAlertIncidentRequestBuilder is request builder for SequentialActivationRenewalsAlertIncident
type SequentialActivationRenewalsAlertIncidentRequestBuilder struct{ BaseRequestBuilder }

// Request returns SequentialActivationRenewalsAlertIncidentRequest
func (b *SequentialActivationRenewalsAlertIncidentRequestBuilder) Request() *SequentialActivationRenewalsAlertIncidentRequest {
	return &SequentialActivationRenewalsAlertIncidentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SequentialActivationRenewalsAlertIncidentRequest is request for SequentialActivationRenewalsAlertIncident
type SequentialActivationRenewalsAlertIncidentRequest struct{ BaseRequest }

// Get performs GET request for SequentialActivationRenewalsAlertIncident
func (r *SequentialActivationRenewalsAlertIncidentRequest) Get(ctx context.Context) (resObj *SequentialActivationRenewalsAlertIncident, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SequentialActivationRenewalsAlertIncident
func (r *SequentialActivationRenewalsAlertIncidentRequest) Update(ctx context.Context, reqObj *SequentialActivationRenewalsAlertIncident) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SequentialActivationRenewalsAlertIncident
func (r *SequentialActivationRenewalsAlertIncidentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

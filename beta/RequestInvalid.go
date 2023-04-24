// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// InvalidLicenseAlertConfigurationRequestBuilder is request builder for InvalidLicenseAlertConfiguration
type InvalidLicenseAlertConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns InvalidLicenseAlertConfigurationRequest
func (b *InvalidLicenseAlertConfigurationRequestBuilder) Request() *InvalidLicenseAlertConfigurationRequest {
	return &InvalidLicenseAlertConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// InvalidLicenseAlertConfigurationRequest is request for InvalidLicenseAlertConfiguration
type InvalidLicenseAlertConfigurationRequest struct{ BaseRequest }

// Get performs GET request for InvalidLicenseAlertConfiguration
func (r *InvalidLicenseAlertConfigurationRequest) Get(ctx context.Context) (resObj *InvalidLicenseAlertConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for InvalidLicenseAlertConfiguration
func (r *InvalidLicenseAlertConfigurationRequest) Update(ctx context.Context, reqObj *InvalidLicenseAlertConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for InvalidLicenseAlertConfiguration
func (r *InvalidLicenseAlertConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// InvalidLicenseAlertIncidentRequestBuilder is request builder for InvalidLicenseAlertIncident
type InvalidLicenseAlertIncidentRequestBuilder struct{ BaseRequestBuilder }

// Request returns InvalidLicenseAlertIncidentRequest
func (b *InvalidLicenseAlertIncidentRequestBuilder) Request() *InvalidLicenseAlertIncidentRequest {
	return &InvalidLicenseAlertIncidentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// InvalidLicenseAlertIncidentRequest is request for InvalidLicenseAlertIncident
type InvalidLicenseAlertIncidentRequest struct{ BaseRequest }

// Get performs GET request for InvalidLicenseAlertIncident
func (r *InvalidLicenseAlertIncidentRequest) Get(ctx context.Context) (resObj *InvalidLicenseAlertIncident, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for InvalidLicenseAlertIncident
func (r *InvalidLicenseAlertIncidentRequest) Update(ctx context.Context, reqObj *InvalidLicenseAlertIncident) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for InvalidLicenseAlertIncident
func (r *InvalidLicenseAlertIncidentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

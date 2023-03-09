// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// OnPremisesAccidentalDeletionPreventionRequestBuilder is request builder for OnPremisesAccidentalDeletionPrevention
type OnPremisesAccidentalDeletionPreventionRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesAccidentalDeletionPreventionRequest
func (b *OnPremisesAccidentalDeletionPreventionRequestBuilder) Request() *OnPremisesAccidentalDeletionPreventionRequest {
	return &OnPremisesAccidentalDeletionPreventionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesAccidentalDeletionPreventionRequest is request for OnPremisesAccidentalDeletionPrevention
type OnPremisesAccidentalDeletionPreventionRequest struct{ BaseRequest }

// Get performs GET request for OnPremisesAccidentalDeletionPrevention
func (r *OnPremisesAccidentalDeletionPreventionRequest) Get(ctx context.Context) (resObj *OnPremisesAccidentalDeletionPrevention, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnPremisesAccidentalDeletionPrevention
func (r *OnPremisesAccidentalDeletionPreventionRequest) Update(ctx context.Context, reqObj *OnPremisesAccidentalDeletionPrevention) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnPremisesAccidentalDeletionPrevention
func (r *OnPremisesAccidentalDeletionPreventionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OnPremisesConditionalAccessSettingsRequestBuilder is request builder for OnPremisesConditionalAccessSettings
type OnPremisesConditionalAccessSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesConditionalAccessSettingsRequest
func (b *OnPremisesConditionalAccessSettingsRequestBuilder) Request() *OnPremisesConditionalAccessSettingsRequest {
	return &OnPremisesConditionalAccessSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesConditionalAccessSettingsRequest is request for OnPremisesConditionalAccessSettings
type OnPremisesConditionalAccessSettingsRequest struct{ BaseRequest }

// Get performs GET request for OnPremisesConditionalAccessSettings
func (r *OnPremisesConditionalAccessSettingsRequest) Get(ctx context.Context) (resObj *OnPremisesConditionalAccessSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnPremisesConditionalAccessSettings
func (r *OnPremisesConditionalAccessSettingsRequest) Update(ctx context.Context, reqObj *OnPremisesConditionalAccessSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnPremisesConditionalAccessSettings
func (r *OnPremisesConditionalAccessSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OnPremisesDirectorySynchronizationRequestBuilder is request builder for OnPremisesDirectorySynchronization
type OnPremisesDirectorySynchronizationRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesDirectorySynchronizationRequest
func (b *OnPremisesDirectorySynchronizationRequestBuilder) Request() *OnPremisesDirectorySynchronizationRequest {
	return &OnPremisesDirectorySynchronizationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesDirectorySynchronizationRequest is request for OnPremisesDirectorySynchronization
type OnPremisesDirectorySynchronizationRequest struct{ BaseRequest }

// Get performs GET request for OnPremisesDirectorySynchronization
func (r *OnPremisesDirectorySynchronizationRequest) Get(ctx context.Context) (resObj *OnPremisesDirectorySynchronization, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnPremisesDirectorySynchronization
func (r *OnPremisesDirectorySynchronizationRequest) Update(ctx context.Context, reqObj *OnPremisesDirectorySynchronization) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnPremisesDirectorySynchronization
func (r *OnPremisesDirectorySynchronizationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OnPremisesDirectorySynchronizationConfigurationRequestBuilder is request builder for OnPremisesDirectorySynchronizationConfiguration
type OnPremisesDirectorySynchronizationConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesDirectorySynchronizationConfigurationRequest
func (b *OnPremisesDirectorySynchronizationConfigurationRequestBuilder) Request() *OnPremisesDirectorySynchronizationConfigurationRequest {
	return &OnPremisesDirectorySynchronizationConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesDirectorySynchronizationConfigurationRequest is request for OnPremisesDirectorySynchronizationConfiguration
type OnPremisesDirectorySynchronizationConfigurationRequest struct{ BaseRequest }

// Get performs GET request for OnPremisesDirectorySynchronizationConfiguration
func (r *OnPremisesDirectorySynchronizationConfigurationRequest) Get(ctx context.Context) (resObj *OnPremisesDirectorySynchronizationConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnPremisesDirectorySynchronizationConfiguration
func (r *OnPremisesDirectorySynchronizationConfigurationRequest) Update(ctx context.Context, reqObj *OnPremisesDirectorySynchronizationConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnPremisesDirectorySynchronizationConfiguration
func (r *OnPremisesDirectorySynchronizationConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OnPremisesDirectorySynchronizationFeatureRequestBuilder is request builder for OnPremisesDirectorySynchronizationFeature
type OnPremisesDirectorySynchronizationFeatureRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesDirectorySynchronizationFeatureRequest
func (b *OnPremisesDirectorySynchronizationFeatureRequestBuilder) Request() *OnPremisesDirectorySynchronizationFeatureRequest {
	return &OnPremisesDirectorySynchronizationFeatureRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesDirectorySynchronizationFeatureRequest is request for OnPremisesDirectorySynchronizationFeature
type OnPremisesDirectorySynchronizationFeatureRequest struct{ BaseRequest }

// Get performs GET request for OnPremisesDirectorySynchronizationFeature
func (r *OnPremisesDirectorySynchronizationFeatureRequest) Get(ctx context.Context) (resObj *OnPremisesDirectorySynchronizationFeature, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnPremisesDirectorySynchronizationFeature
func (r *OnPremisesDirectorySynchronizationFeatureRequest) Update(ctx context.Context, reqObj *OnPremisesDirectorySynchronizationFeature) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnPremisesDirectorySynchronizationFeature
func (r *OnPremisesDirectorySynchronizationFeatureRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OnPremisesExtensionAttributesRequestBuilder is request builder for OnPremisesExtensionAttributes
type OnPremisesExtensionAttributesRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesExtensionAttributesRequest
func (b *OnPremisesExtensionAttributesRequestBuilder) Request() *OnPremisesExtensionAttributesRequest {
	return &OnPremisesExtensionAttributesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesExtensionAttributesRequest is request for OnPremisesExtensionAttributes
type OnPremisesExtensionAttributesRequest struct{ BaseRequest }

// Get performs GET request for OnPremisesExtensionAttributes
func (r *OnPremisesExtensionAttributesRequest) Get(ctx context.Context) (resObj *OnPremisesExtensionAttributes, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnPremisesExtensionAttributes
func (r *OnPremisesExtensionAttributesRequest) Update(ctx context.Context, reqObj *OnPremisesExtensionAttributes) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnPremisesExtensionAttributes
func (r *OnPremisesExtensionAttributesRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// OnPremisesProvisioningErrorRequestBuilder is request builder for OnPremisesProvisioningError
type OnPremisesProvisioningErrorRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnPremisesProvisioningErrorRequest
func (b *OnPremisesProvisioningErrorRequestBuilder) Request() *OnPremisesProvisioningErrorRequest {
	return &OnPremisesProvisioningErrorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnPremisesProvisioningErrorRequest is request for OnPremisesProvisioningError
type OnPremisesProvisioningErrorRequest struct{ BaseRequest }

// Get performs GET request for OnPremisesProvisioningError
func (r *OnPremisesProvisioningErrorRequest) Get(ctx context.Context) (resObj *OnPremisesProvisioningError, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnPremisesProvisioningError
func (r *OnPremisesProvisioningErrorRequest) Update(ctx context.Context, reqObj *OnPremisesProvisioningError) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnPremisesProvisioningError
func (r *OnPremisesProvisioningErrorRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

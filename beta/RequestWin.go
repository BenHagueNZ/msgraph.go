// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// WinGetAppRequestBuilder is request builder for WinGetApp
type WinGetAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns WinGetAppRequest
func (b *WinGetAppRequestBuilder) Request() *WinGetAppRequest {
	return &WinGetAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WinGetAppRequest is request for WinGetApp
type WinGetAppRequest struct{ BaseRequest }

// Get performs GET request for WinGetApp
func (r *WinGetAppRequest) Get(ctx context.Context) (resObj *WinGetApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WinGetApp
func (r *WinGetAppRequest) Update(ctx context.Context, reqObj *WinGetApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WinGetApp
func (r *WinGetAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// WinGetAppAssignmentSettingsRequestBuilder is request builder for WinGetAppAssignmentSettings
type WinGetAppAssignmentSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns WinGetAppAssignmentSettingsRequest
func (b *WinGetAppAssignmentSettingsRequestBuilder) Request() *WinGetAppAssignmentSettingsRequest {
	return &WinGetAppAssignmentSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WinGetAppAssignmentSettingsRequest is request for WinGetAppAssignmentSettings
type WinGetAppAssignmentSettingsRequest struct{ BaseRequest }

// Get performs GET request for WinGetAppAssignmentSettings
func (r *WinGetAppAssignmentSettingsRequest) Get(ctx context.Context) (resObj *WinGetAppAssignmentSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WinGetAppAssignmentSettings
func (r *WinGetAppAssignmentSettingsRequest) Update(ctx context.Context, reqObj *WinGetAppAssignmentSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WinGetAppAssignmentSettings
func (r *WinGetAppAssignmentSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// WinGetAppInstallExperienceRequestBuilder is request builder for WinGetAppInstallExperience
type WinGetAppInstallExperienceRequestBuilder struct{ BaseRequestBuilder }

// Request returns WinGetAppInstallExperienceRequest
func (b *WinGetAppInstallExperienceRequestBuilder) Request() *WinGetAppInstallExperienceRequest {
	return &WinGetAppInstallExperienceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WinGetAppInstallExperienceRequest is request for WinGetAppInstallExperience
type WinGetAppInstallExperienceRequest struct{ BaseRequest }

// Get performs GET request for WinGetAppInstallExperience
func (r *WinGetAppInstallExperienceRequest) Get(ctx context.Context) (resObj *WinGetAppInstallExperience, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WinGetAppInstallExperience
func (r *WinGetAppInstallExperienceRequest) Update(ctx context.Context, reqObj *WinGetAppInstallExperience) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WinGetAppInstallExperience
func (r *WinGetAppInstallExperienceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// WinGetAppInstallTimeSettingsRequestBuilder is request builder for WinGetAppInstallTimeSettings
type WinGetAppInstallTimeSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns WinGetAppInstallTimeSettingsRequest
func (b *WinGetAppInstallTimeSettingsRequestBuilder) Request() *WinGetAppInstallTimeSettingsRequest {
	return &WinGetAppInstallTimeSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WinGetAppInstallTimeSettingsRequest is request for WinGetAppInstallTimeSettings
type WinGetAppInstallTimeSettingsRequest struct{ BaseRequest }

// Get performs GET request for WinGetAppInstallTimeSettings
func (r *WinGetAppInstallTimeSettingsRequest) Get(ctx context.Context) (resObj *WinGetAppInstallTimeSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WinGetAppInstallTimeSettings
func (r *WinGetAppInstallTimeSettingsRequest) Update(ctx context.Context, reqObj *WinGetAppInstallTimeSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WinGetAppInstallTimeSettings
func (r *WinGetAppInstallTimeSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// WinGetAppRestartSettingsRequestBuilder is request builder for WinGetAppRestartSettings
type WinGetAppRestartSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns WinGetAppRestartSettingsRequest
func (b *WinGetAppRestartSettingsRequestBuilder) Request() *WinGetAppRestartSettingsRequest {
	return &WinGetAppRestartSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WinGetAppRestartSettingsRequest is request for WinGetAppRestartSettings
type WinGetAppRestartSettingsRequest struct{ BaseRequest }

// Get performs GET request for WinGetAppRestartSettings
func (r *WinGetAppRestartSettingsRequest) Get(ctx context.Context) (resObj *WinGetAppRestartSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WinGetAppRestartSettings
func (r *WinGetAppRestartSettingsRequest) Update(ctx context.Context, reqObj *WinGetAppRestartSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WinGetAppRestartSettings
func (r *WinGetAppRestartSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

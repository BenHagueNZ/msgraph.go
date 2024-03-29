// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// BitLockerFixedDrivePolicyRequestBuilder is request builder for BitLockerFixedDrivePolicy
type BitLockerFixedDrivePolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns BitLockerFixedDrivePolicyRequest
func (b *BitLockerFixedDrivePolicyRequestBuilder) Request() *BitLockerFixedDrivePolicyRequest {
	return &BitLockerFixedDrivePolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BitLockerFixedDrivePolicyRequest is request for BitLockerFixedDrivePolicy
type BitLockerFixedDrivePolicyRequest struct{ BaseRequest }

// Get performs GET request for BitLockerFixedDrivePolicy
func (r *BitLockerFixedDrivePolicyRequest) Get(ctx context.Context) (resObj *BitLockerFixedDrivePolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BitLockerFixedDrivePolicy
func (r *BitLockerFixedDrivePolicyRequest) Update(ctx context.Context, reqObj *BitLockerFixedDrivePolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BitLockerFixedDrivePolicy
func (r *BitLockerFixedDrivePolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BitLockerRecoveryOptionsRequestBuilder is request builder for BitLockerRecoveryOptions
type BitLockerRecoveryOptionsRequestBuilder struct{ BaseRequestBuilder }

// Request returns BitLockerRecoveryOptionsRequest
func (b *BitLockerRecoveryOptionsRequestBuilder) Request() *BitLockerRecoveryOptionsRequest {
	return &BitLockerRecoveryOptionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BitLockerRecoveryOptionsRequest is request for BitLockerRecoveryOptions
type BitLockerRecoveryOptionsRequest struct{ BaseRequest }

// Get performs GET request for BitLockerRecoveryOptions
func (r *BitLockerRecoveryOptionsRequest) Get(ctx context.Context) (resObj *BitLockerRecoveryOptions, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BitLockerRecoveryOptions
func (r *BitLockerRecoveryOptionsRequest) Update(ctx context.Context, reqObj *BitLockerRecoveryOptions) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BitLockerRecoveryOptions
func (r *BitLockerRecoveryOptionsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BitLockerRemovableDrivePolicyRequestBuilder is request builder for BitLockerRemovableDrivePolicy
type BitLockerRemovableDrivePolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns BitLockerRemovableDrivePolicyRequest
func (b *BitLockerRemovableDrivePolicyRequestBuilder) Request() *BitLockerRemovableDrivePolicyRequest {
	return &BitLockerRemovableDrivePolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BitLockerRemovableDrivePolicyRequest is request for BitLockerRemovableDrivePolicy
type BitLockerRemovableDrivePolicyRequest struct{ BaseRequest }

// Get performs GET request for BitLockerRemovableDrivePolicy
func (r *BitLockerRemovableDrivePolicyRequest) Get(ctx context.Context) (resObj *BitLockerRemovableDrivePolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BitLockerRemovableDrivePolicy
func (r *BitLockerRemovableDrivePolicyRequest) Update(ctx context.Context, reqObj *BitLockerRemovableDrivePolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BitLockerRemovableDrivePolicy
func (r *BitLockerRemovableDrivePolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BitLockerSystemDrivePolicyRequestBuilder is request builder for BitLockerSystemDrivePolicy
type BitLockerSystemDrivePolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns BitLockerSystemDrivePolicyRequest
func (b *BitLockerSystemDrivePolicyRequestBuilder) Request() *BitLockerSystemDrivePolicyRequest {
	return &BitLockerSystemDrivePolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BitLockerSystemDrivePolicyRequest is request for BitLockerSystemDrivePolicy
type BitLockerSystemDrivePolicyRequest struct{ BaseRequest }

// Get performs GET request for BitLockerSystemDrivePolicy
func (r *BitLockerSystemDrivePolicyRequest) Get(ctx context.Context) (resObj *BitLockerSystemDrivePolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BitLockerSystemDrivePolicy
func (r *BitLockerSystemDrivePolicyRequest) Update(ctx context.Context, reqObj *BitLockerSystemDrivePolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BitLockerSystemDrivePolicy
func (r *BitLockerSystemDrivePolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

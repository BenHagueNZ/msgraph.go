// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// NetworkConnectionRequestBuilder is request builder for NetworkConnection
type NetworkConnectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns NetworkConnectionRequest
func (b *NetworkConnectionRequestBuilder) Request() *NetworkConnectionRequest {
	return &NetworkConnectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// NetworkConnectionRequest is request for NetworkConnection
type NetworkConnectionRequest struct{ BaseRequest }

// Get performs GET request for NetworkConnection
func (r *NetworkConnectionRequest) Get(ctx context.Context) (resObj *NetworkConnection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for NetworkConnection
func (r *NetworkConnectionRequest) Update(ctx context.Context, reqObj *NetworkConnection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for NetworkConnection
func (r *NetworkConnectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// NetworkInterfaceRequestBuilder is request builder for NetworkInterface
type NetworkInterfaceRequestBuilder struct{ BaseRequestBuilder }

// Request returns NetworkInterfaceRequest
func (b *NetworkInterfaceRequestBuilder) Request() *NetworkInterfaceRequest {
	return &NetworkInterfaceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// NetworkInterfaceRequest is request for NetworkInterface
type NetworkInterfaceRequest struct{ BaseRequest }

// Get performs GET request for NetworkInterface
func (r *NetworkInterfaceRequest) Get(ctx context.Context) (resObj *NetworkInterface, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for NetworkInterface
func (r *NetworkInterfaceRequest) Update(ctx context.Context, reqObj *NetworkInterface) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for NetworkInterface
func (r *NetworkInterfaceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// NetworkLocationDetailRequestBuilder is request builder for NetworkLocationDetail
type NetworkLocationDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns NetworkLocationDetailRequest
func (b *NetworkLocationDetailRequestBuilder) Request() *NetworkLocationDetailRequest {
	return &NetworkLocationDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// NetworkLocationDetailRequest is request for NetworkLocationDetail
type NetworkLocationDetailRequest struct{ BaseRequest }

// Get performs GET request for NetworkLocationDetail
func (r *NetworkLocationDetailRequest) Get(ctx context.Context) (resObj *NetworkLocationDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for NetworkLocationDetail
func (r *NetworkLocationDetailRequest) Update(ctx context.Context, reqObj *NetworkLocationDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for NetworkLocationDetail
func (r *NetworkLocationDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

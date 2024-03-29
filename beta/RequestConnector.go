// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ConnectorRequestBuilder is request builder for Connector
type ConnectorRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConnectorRequest
func (b *ConnectorRequestBuilder) Request() *ConnectorRequest {
	return &ConnectorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConnectorRequest is request for Connector
type ConnectorRequest struct{ BaseRequest }

// Get performs GET request for Connector
func (r *ConnectorRequest) Get(ctx context.Context) (resObj *Connector, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Connector
func (r *ConnectorRequest) Update(ctx context.Context, reqObj *Connector) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Connector
func (r *ConnectorRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConnectorGroupRequestBuilder is request builder for ConnectorGroup
type ConnectorGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConnectorGroupRequest
func (b *ConnectorGroupRequestBuilder) Request() *ConnectorGroupRequest {
	return &ConnectorGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConnectorGroupRequest is request for ConnectorGroup
type ConnectorGroupRequest struct{ BaseRequest }

// Get performs GET request for ConnectorGroup
func (r *ConnectorGroupRequest) Get(ctx context.Context) (resObj *ConnectorGroup, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConnectorGroup
func (r *ConnectorGroupRequest) Update(ctx context.Context, reqObj *ConnectorGroup) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConnectorGroup
func (r *ConnectorGroupRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConnectorStatusDetailsRequestBuilder is request builder for ConnectorStatusDetails
type ConnectorStatusDetailsRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConnectorStatusDetailsRequest
func (b *ConnectorStatusDetailsRequestBuilder) Request() *ConnectorStatusDetailsRequest {
	return &ConnectorStatusDetailsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConnectorStatusDetailsRequest is request for ConnectorStatusDetails
type ConnectorStatusDetailsRequest struct{ BaseRequest }

// Get performs GET request for ConnectorStatusDetails
func (r *ConnectorStatusDetailsRequest) Get(ctx context.Context) (resObj *ConnectorStatusDetails, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConnectorStatusDetails
func (r *ConnectorStatusDetailsRequest) Update(ctx context.Context, reqObj *ConnectorStatusDetails) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConnectorStatusDetails
func (r *ConnectorStatusDetailsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// InternetExplorerModeRequestBuilder is request builder for InternetExplorerMode
type InternetExplorerModeRequestBuilder struct{ BaseRequestBuilder }

// Request returns InternetExplorerModeRequest
func (b *InternetExplorerModeRequestBuilder) Request() *InternetExplorerModeRequest {
	return &InternetExplorerModeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// InternetExplorerModeRequest is request for InternetExplorerMode
type InternetExplorerModeRequest struct{ BaseRequest }

// Get performs GET request for InternetExplorerMode
func (r *InternetExplorerModeRequest) Get(ctx context.Context) (resObj *InternetExplorerMode, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for InternetExplorerMode
func (r *InternetExplorerModeRequest) Update(ctx context.Context, reqObj *InternetExplorerMode) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for InternetExplorerMode
func (r *InternetExplorerModeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// InternetMessageHeaderRequestBuilder is request builder for InternetMessageHeader
type InternetMessageHeaderRequestBuilder struct{ BaseRequestBuilder }

// Request returns InternetMessageHeaderRequest
func (b *InternetMessageHeaderRequestBuilder) Request() *InternetMessageHeaderRequest {
	return &InternetMessageHeaderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// InternetMessageHeaderRequest is request for InternetMessageHeader
type InternetMessageHeaderRequest struct{ BaseRequest }

// Get performs GET request for InternetMessageHeader
func (r *InternetMessageHeaderRequest) Get(ctx context.Context) (resObj *InternetMessageHeader, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for InternetMessageHeader
func (r *InternetMessageHeaderRequest) Update(ctx context.Context, reqObj *InternetMessageHeader) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for InternetMessageHeader
func (r *InternetMessageHeaderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
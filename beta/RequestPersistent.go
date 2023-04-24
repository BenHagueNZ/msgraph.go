// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PersistentBrowserSessionControlRequestBuilder is request builder for PersistentBrowserSessionControl
type PersistentBrowserSessionControlRequestBuilder struct{ BaseRequestBuilder }

// Request returns PersistentBrowserSessionControlRequest
func (b *PersistentBrowserSessionControlRequestBuilder) Request() *PersistentBrowserSessionControlRequest {
	return &PersistentBrowserSessionControlRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PersistentBrowserSessionControlRequest is request for PersistentBrowserSessionControl
type PersistentBrowserSessionControlRequest struct{ BaseRequest }

// Get performs GET request for PersistentBrowserSessionControl
func (r *PersistentBrowserSessionControlRequest) Get(ctx context.Context) (resObj *PersistentBrowserSessionControl, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PersistentBrowserSessionControl
func (r *PersistentBrowserSessionControlRequest) Update(ctx context.Context, reqObj *PersistentBrowserSessionControl) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PersistentBrowserSessionControl
func (r *PersistentBrowserSessionControlRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
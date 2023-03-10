// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// DefenderDetectedMalwareActionsRequestBuilder is request builder for DefenderDetectedMalwareActions
type DefenderDetectedMalwareActionsRequestBuilder struct{ BaseRequestBuilder }

// Request returns DefenderDetectedMalwareActionsRequest
func (b *DefenderDetectedMalwareActionsRequestBuilder) Request() *DefenderDetectedMalwareActionsRequest {
	return &DefenderDetectedMalwareActionsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DefenderDetectedMalwareActionsRequest is request for DefenderDetectedMalwareActions
type DefenderDetectedMalwareActionsRequest struct{ BaseRequest }

// Get performs GET request for DefenderDetectedMalwareActions
func (r *DefenderDetectedMalwareActionsRequest) Get(ctx context.Context) (resObj *DefenderDetectedMalwareActions, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DefenderDetectedMalwareActions
func (r *DefenderDetectedMalwareActionsRequest) Update(ctx context.Context, reqObj *DefenderDetectedMalwareActions) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DefenderDetectedMalwareActions
func (r *DefenderDetectedMalwareActionsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

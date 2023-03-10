// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// TicketInfoRequestBuilder is request builder for TicketInfo
type TicketInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns TicketInfoRequest
func (b *TicketInfoRequestBuilder) Request() *TicketInfoRequest {
	return &TicketInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TicketInfoRequest is request for TicketInfo
type TicketInfoRequest struct{ BaseRequest }

// Get performs GET request for TicketInfo
func (r *TicketInfoRequest) Get(ctx context.Context) (resObj *TicketInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TicketInfo
func (r *TicketInfoRequest) Update(ctx context.Context, reqObj *TicketInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TicketInfo
func (r *TicketInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

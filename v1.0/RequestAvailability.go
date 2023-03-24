// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AvailabilityItemRequestBuilder is request builder for AvailabilityItem
type AvailabilityItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns AvailabilityItemRequest
func (b *AvailabilityItemRequestBuilder) Request() *AvailabilityItemRequest {
	return &AvailabilityItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AvailabilityItemRequest is request for AvailabilityItem
type AvailabilityItemRequest struct{ BaseRequest }

// Get performs GET request for AvailabilityItem
func (r *AvailabilityItemRequest) Get(ctx context.Context) (resObj *AvailabilityItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AvailabilityItem
func (r *AvailabilityItemRequest) Update(ctx context.Context, reqObj *AvailabilityItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AvailabilityItem
func (r *AvailabilityItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

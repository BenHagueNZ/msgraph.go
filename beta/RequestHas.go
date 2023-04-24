// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// HasPayloadLinkResultItemRequestBuilder is request builder for HasPayloadLinkResultItem
type HasPayloadLinkResultItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns HasPayloadLinkResultItemRequest
func (b *HasPayloadLinkResultItemRequestBuilder) Request() *HasPayloadLinkResultItemRequest {
	return &HasPayloadLinkResultItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// HasPayloadLinkResultItemRequest is request for HasPayloadLinkResultItem
type HasPayloadLinkResultItemRequest struct{ BaseRequest }

// Get performs GET request for HasPayloadLinkResultItem
func (r *HasPayloadLinkResultItemRequest) Get(ctx context.Context) (resObj *HasPayloadLinkResultItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for HasPayloadLinkResultItem
func (r *HasPayloadLinkResultItemRequest) Update(ctx context.Context, reqObj *HasPayloadLinkResultItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for HasPayloadLinkResultItem
func (r *HasPayloadLinkResultItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PageLinksRequestBuilder is request builder for PageLinks
type PageLinksRequestBuilder struct{ BaseRequestBuilder }

// Request returns PageLinksRequest
func (b *PageLinksRequestBuilder) Request() *PageLinksRequest {
	return &PageLinksRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PageLinksRequest is request for PageLinks
type PageLinksRequest struct{ BaseRequest }

// Get performs GET request for PageLinks
func (r *PageLinksRequest) Get(ctx context.Context) (resObj *PageLinks, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PageLinks
func (r *PageLinksRequest) Update(ctx context.Context, reqObj *PageLinks) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PageLinks
func (r *PageLinksRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

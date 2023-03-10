// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SectionGroupRequestBuilder is request builder for SectionGroup
type SectionGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns SectionGroupRequest
func (b *SectionGroupRequestBuilder) Request() *SectionGroupRequest {
	return &SectionGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SectionGroupRequest is request for SectionGroup
type SectionGroupRequest struct{ BaseRequest }

// Get performs GET request for SectionGroup
func (r *SectionGroupRequest) Get(ctx context.Context) (resObj *SectionGroup, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SectionGroup
func (r *SectionGroupRequest) Update(ctx context.Context, reqObj *SectionGroup) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SectionGroup
func (r *SectionGroupRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SectionLinksRequestBuilder is request builder for SectionLinks
type SectionLinksRequestBuilder struct{ BaseRequestBuilder }

// Request returns SectionLinksRequest
func (b *SectionLinksRequestBuilder) Request() *SectionLinksRequest {
	return &SectionLinksRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SectionLinksRequest is request for SectionLinks
type SectionLinksRequest struct{ BaseRequest }

// Get performs GET request for SectionLinks
func (r *SectionLinksRequest) Get(ctx context.Context) (resObj *SectionLinks, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SectionLinks
func (r *SectionLinksRequest) Update(ctx context.Context, reqObj *SectionLinks) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SectionLinks
func (r *SectionLinksRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

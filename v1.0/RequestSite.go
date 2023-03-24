// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SiteRequestBuilder is request builder for Site
type SiteRequestBuilder struct{ BaseRequestBuilder }

// Request returns SiteRequest
func (b *SiteRequestBuilder) Request() *SiteRequest {
	return &SiteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SiteRequest is request for Site
type SiteRequest struct{ BaseRequest }

// Get performs GET request for Site
func (r *SiteRequest) Get(ctx context.Context) (resObj *Site, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Site
func (r *SiteRequest) Update(ctx context.Context, reqObj *Site) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Site
func (r *SiteRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SiteCollectionRequestBuilder is request builder for SiteCollection
type SiteCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns SiteCollectionRequest
func (b *SiteCollectionRequestBuilder) Request() *SiteCollectionRequest {
	return &SiteCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SiteCollectionRequest is request for SiteCollection
type SiteCollectionRequest struct{ BaseRequest }

// Get performs GET request for SiteCollection
func (r *SiteCollectionRequest) Get(ctx context.Context) (resObj *SiteCollection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SiteCollection
func (r *SiteCollectionRequest) Update(ctx context.Context, reqObj *SiteCollection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SiteCollection
func (r *SiteCollectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

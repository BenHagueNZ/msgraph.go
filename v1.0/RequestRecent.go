// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// RecentNotebookRequestBuilder is request builder for RecentNotebook
type RecentNotebookRequestBuilder struct{ BaseRequestBuilder }

// Request returns RecentNotebookRequest
func (b *RecentNotebookRequestBuilder) Request() *RecentNotebookRequest {
	return &RecentNotebookRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RecentNotebookRequest is request for RecentNotebook
type RecentNotebookRequest struct{ BaseRequest }

// Get performs GET request for RecentNotebook
func (r *RecentNotebookRequest) Get(ctx context.Context) (resObj *RecentNotebook, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RecentNotebook
func (r *RecentNotebookRequest) Update(ctx context.Context, reqObj *RecentNotebook) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RecentNotebook
func (r *RecentNotebookRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RecentNotebookLinksRequestBuilder is request builder for RecentNotebookLinks
type RecentNotebookLinksRequestBuilder struct{ BaseRequestBuilder }

// Request returns RecentNotebookLinksRequest
func (b *RecentNotebookLinksRequestBuilder) Request() *RecentNotebookLinksRequest {
	return &RecentNotebookLinksRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RecentNotebookLinksRequest is request for RecentNotebookLinks
type RecentNotebookLinksRequest struct{ BaseRequest }

// Get performs GET request for RecentNotebookLinks
func (r *RecentNotebookLinksRequest) Get(ctx context.Context) (resObj *RecentNotebookLinks, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RecentNotebookLinks
func (r *RecentNotebookLinksRequest) Update(ctx context.Context, reqObj *RecentNotebookLinks) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RecentNotebookLinks
func (r *RecentNotebookLinksRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

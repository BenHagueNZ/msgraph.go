// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// FolderRequestBuilder is request builder for Folder
type FolderRequestBuilder struct{ BaseRequestBuilder }

// Request returns FolderRequest
func (b *FolderRequestBuilder) Request() *FolderRequest {
	return &FolderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// FolderRequest is request for Folder
type FolderRequest struct{ BaseRequest }

// Get performs GET request for Folder
func (r *FolderRequest) Get(ctx context.Context) (resObj *Folder, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Folder
func (r *FolderRequest) Update(ctx context.Context, reqObj *Folder) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Folder
func (r *FolderRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// FolderViewRequestBuilder is request builder for FolderView
type FolderViewRequestBuilder struct{ BaseRequestBuilder }

// Request returns FolderViewRequest
func (b *FolderViewRequestBuilder) Request() *FolderViewRequest {
	return &FolderViewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// FolderViewRequest is request for FolderView
type FolderViewRequest struct{ BaseRequest }

// Get performs GET request for FolderView
func (r *FolderViewRequest) Get(ctx context.Context) (resObj *FolderView, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for FolderView
func (r *FolderViewRequest) Update(ctx context.Context, reqObj *FolderView) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for FolderView
func (r *FolderViewRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

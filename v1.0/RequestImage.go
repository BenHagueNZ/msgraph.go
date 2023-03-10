// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ImageRequestBuilder is request builder for Image
type ImageRequestBuilder struct{ BaseRequestBuilder }

// Request returns ImageRequest
func (b *ImageRequestBuilder) Request() *ImageRequest {
	return &ImageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ImageRequest is request for Image
type ImageRequest struct{ BaseRequest }

// Get performs GET request for Image
func (r *ImageRequest) Get(ctx context.Context) (resObj *Image, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Image
func (r *ImageRequest) Update(ctx context.Context, reqObj *Image) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Image
func (r *ImageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ImageInfoRequestBuilder is request builder for ImageInfo
type ImageInfoRequestBuilder struct{ BaseRequestBuilder }

// Request returns ImageInfoRequest
func (b *ImageInfoRequestBuilder) Request() *ImageInfoRequest {
	return &ImageInfoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ImageInfoRequest is request for ImageInfo
type ImageInfoRequest struct{ BaseRequest }

// Get performs GET request for ImageInfo
func (r *ImageInfoRequest) Get(ctx context.Context) (resObj *ImageInfo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ImageInfo
func (r *ImageInfoRequest) Update(ctx context.Context, reqObj *ImageInfo) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ImageInfo
func (r *ImageInfoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

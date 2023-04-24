// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ProfileRequestBuilder is request builder for Profile
type ProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProfileRequest
func (b *ProfileRequestBuilder) Request() *ProfileRequest {
	return &ProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProfileRequest is request for Profile
type ProfileRequest struct{ BaseRequest }

// Get performs GET request for Profile
func (r *ProfileRequest) Get(ctx context.Context) (resObj *Profile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Profile
func (r *ProfileRequest) Update(ctx context.Context, reqObj *Profile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Profile
func (r *ProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ProfileCardAnnotationRequestBuilder is request builder for ProfileCardAnnotation
type ProfileCardAnnotationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProfileCardAnnotationRequest
func (b *ProfileCardAnnotationRequestBuilder) Request() *ProfileCardAnnotationRequest {
	return &ProfileCardAnnotationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProfileCardAnnotationRequest is request for ProfileCardAnnotation
type ProfileCardAnnotationRequest struct{ BaseRequest }

// Get performs GET request for ProfileCardAnnotation
func (r *ProfileCardAnnotationRequest) Get(ctx context.Context) (resObj *ProfileCardAnnotation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ProfileCardAnnotation
func (r *ProfileCardAnnotationRequest) Update(ctx context.Context, reqObj *ProfileCardAnnotation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ProfileCardAnnotation
func (r *ProfileCardAnnotationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ProfileCardPropertyRequestBuilder is request builder for ProfileCardProperty
type ProfileCardPropertyRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProfileCardPropertyRequest
func (b *ProfileCardPropertyRequestBuilder) Request() *ProfileCardPropertyRequest {
	return &ProfileCardPropertyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProfileCardPropertyRequest is request for ProfileCardProperty
type ProfileCardPropertyRequest struct{ BaseRequest }

// Get performs GET request for ProfileCardProperty
func (r *ProfileCardPropertyRequest) Get(ctx context.Context) (resObj *ProfileCardProperty, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ProfileCardProperty
func (r *ProfileCardPropertyRequest) Update(ctx context.Context, reqObj *ProfileCardProperty) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ProfileCardProperty
func (r *ProfileCardPropertyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ProfilePhotoRequestBuilder is request builder for ProfilePhoto
type ProfilePhotoRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProfilePhotoRequest
func (b *ProfilePhotoRequestBuilder) Request() *ProfilePhotoRequest {
	return &ProfilePhotoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProfilePhotoRequest is request for ProfilePhoto
type ProfilePhotoRequest struct{ BaseRequest }

// Get performs GET request for ProfilePhoto
func (r *ProfilePhotoRequest) Get(ctx context.Context) (resObj *ProfilePhoto, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ProfilePhoto
func (r *ProfilePhotoRequest) Update(ctx context.Context, reqObj *ProfilePhoto) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ProfilePhoto
func (r *ProfilePhotoRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

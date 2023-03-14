
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// UserRequestBuilder is request builder for User
type UserRequestBuilder struct { BaseRequestBuilder }

// Request returns UserRequest
func (b *UserRequestBuilder) Request() *UserRequest {
    return &UserRequest{
        BaseRequest: BaseRequest{ baseURL: b.baseURL, client:  b.client },
    }
}

// UserRequest is request for User
type UserRequest struct{ BaseRequest }

// Get performs GET request for User
func (r *UserRequest) Get(ctx context.Context) (resObj *User, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for User
func (r *UserRequest) Update(ctx context.Context, reqObj *User) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for User
func (r *UserRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

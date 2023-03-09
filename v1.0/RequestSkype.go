// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SkypeForBusinessUserConversationMemberRequestBuilder is request builder for SkypeForBusinessUserConversationMember
type SkypeForBusinessUserConversationMemberRequestBuilder struct{ BaseRequestBuilder }

// Request returns SkypeForBusinessUserConversationMemberRequest
func (b *SkypeForBusinessUserConversationMemberRequestBuilder) Request() *SkypeForBusinessUserConversationMemberRequest {
	return &SkypeForBusinessUserConversationMemberRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SkypeForBusinessUserConversationMemberRequest is request for SkypeForBusinessUserConversationMember
type SkypeForBusinessUserConversationMemberRequest struct{ BaseRequest }

// Get performs GET request for SkypeForBusinessUserConversationMember
func (r *SkypeForBusinessUserConversationMemberRequest) Get(ctx context.Context) (resObj *SkypeForBusinessUserConversationMember, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SkypeForBusinessUserConversationMember
func (r *SkypeForBusinessUserConversationMemberRequest) Update(ctx context.Context, reqObj *SkypeForBusinessUserConversationMember) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SkypeForBusinessUserConversationMember
func (r *SkypeForBusinessUserConversationMemberRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SkypeUserConversationMemberRequestBuilder is request builder for SkypeUserConversationMember
type SkypeUserConversationMemberRequestBuilder struct{ BaseRequestBuilder }

// Request returns SkypeUserConversationMemberRequest
func (b *SkypeUserConversationMemberRequestBuilder) Request() *SkypeUserConversationMemberRequest {
	return &SkypeUserConversationMemberRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SkypeUserConversationMemberRequest is request for SkypeUserConversationMember
type SkypeUserConversationMemberRequest struct{ BaseRequest }

// Get performs GET request for SkypeUserConversationMember
func (r *SkypeUserConversationMemberRequest) Get(ctx context.Context) (resObj *SkypeUserConversationMember, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SkypeUserConversationMember
func (r *SkypeUserConversationMemberRequest) Update(ctx context.Context, reqObj *SkypeUserConversationMember) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SkypeUserConversationMember
func (r *SkypeUserConversationMemberRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

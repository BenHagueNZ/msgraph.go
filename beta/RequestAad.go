// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AadUserConversationMemberRequestBuilder is request builder for AadUserConversationMember
type AadUserConversationMemberRequestBuilder struct{ BaseRequestBuilder }

// Request returns AadUserConversationMemberRequest
func (b *AadUserConversationMemberRequestBuilder) Request() *AadUserConversationMemberRequest {
	return &AadUserConversationMemberRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AadUserConversationMemberRequest is request for AadUserConversationMember
type AadUserConversationMemberRequest struct{ BaseRequest }

// Get performs GET request for AadUserConversationMember
func (r *AadUserConversationMemberRequest) Get(ctx context.Context) (resObj *AadUserConversationMember, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AadUserConversationMember
func (r *AadUserConversationMemberRequest) Update(ctx context.Context, reqObj *AadUserConversationMember) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AadUserConversationMember
func (r *AadUserConversationMemberRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AadUserConversationMemberResultRequestBuilder is request builder for AadUserConversationMemberResult
type AadUserConversationMemberResultRequestBuilder struct{ BaseRequestBuilder }

// Request returns AadUserConversationMemberResultRequest
func (b *AadUserConversationMemberResultRequestBuilder) Request() *AadUserConversationMemberResultRequest {
	return &AadUserConversationMemberResultRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AadUserConversationMemberResultRequest is request for AadUserConversationMemberResult
type AadUserConversationMemberResultRequest struct{ BaseRequest }

// Get performs GET request for AadUserConversationMemberResult
func (r *AadUserConversationMemberResultRequest) Get(ctx context.Context) (resObj *AadUserConversationMemberResult, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AadUserConversationMemberResult
func (r *AadUserConversationMemberResultRequest) Update(ctx context.Context, reqObj *AadUserConversationMemberResult) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AadUserConversationMemberResult
func (r *AadUserConversationMemberResultRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AadUserNotificationRecipientRequestBuilder is request builder for AadUserNotificationRecipient
type AadUserNotificationRecipientRequestBuilder struct{ BaseRequestBuilder }

// Request returns AadUserNotificationRecipientRequest
func (b *AadUserNotificationRecipientRequestBuilder) Request() *AadUserNotificationRecipientRequest {
	return &AadUserNotificationRecipientRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AadUserNotificationRecipientRequest is request for AadUserNotificationRecipient
type AadUserNotificationRecipientRequest struct{ BaseRequest }

// Get performs GET request for AadUserNotificationRecipient
func (r *AadUserNotificationRecipientRequest) Get(ctx context.Context) (resObj *AadUserNotificationRecipient, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AadUserNotificationRecipient
func (r *AadUserNotificationRecipientRequest) Update(ctx context.Context, reqObj *AadUserNotificationRecipient) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AadUserNotificationRecipient
func (r *AadUserNotificationRecipientRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

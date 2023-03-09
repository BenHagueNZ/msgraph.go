// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// ConversationRequestBuilder is request builder for Conversation
type ConversationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConversationRequest
func (b *ConversationRequestBuilder) Request() *ConversationRequest {
	return &ConversationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConversationRequest is request for Conversation
type ConversationRequest struct{ BaseRequest }

// Get performs GET request for Conversation
func (r *ConversationRequest) Get(ctx context.Context) (resObj *Conversation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Conversation
func (r *ConversationRequest) Update(ctx context.Context, reqObj *Conversation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Conversation
func (r *ConversationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConversationMemberRequestBuilder is request builder for ConversationMember
type ConversationMemberRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConversationMemberRequest
func (b *ConversationMemberRequestBuilder) Request() *ConversationMemberRequest {
	return &ConversationMemberRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConversationMemberRequest is request for ConversationMember
type ConversationMemberRequest struct{ BaseRequest }

// Get performs GET request for ConversationMember
func (r *ConversationMemberRequest) Get(ctx context.Context) (resObj *ConversationMember, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConversationMember
func (r *ConversationMemberRequest) Update(ctx context.Context, reqObj *ConversationMember) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConversationMember
func (r *ConversationMemberRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConversationMemberRoleUpdatedEventMessageDetailRequestBuilder is request builder for ConversationMemberRoleUpdatedEventMessageDetail
type ConversationMemberRoleUpdatedEventMessageDetailRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConversationMemberRoleUpdatedEventMessageDetailRequest
func (b *ConversationMemberRoleUpdatedEventMessageDetailRequestBuilder) Request() *ConversationMemberRoleUpdatedEventMessageDetailRequest {
	return &ConversationMemberRoleUpdatedEventMessageDetailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConversationMemberRoleUpdatedEventMessageDetailRequest is request for ConversationMemberRoleUpdatedEventMessageDetail
type ConversationMemberRoleUpdatedEventMessageDetailRequest struct{ BaseRequest }

// Get performs GET request for ConversationMemberRoleUpdatedEventMessageDetail
func (r *ConversationMemberRoleUpdatedEventMessageDetailRequest) Get(ctx context.Context) (resObj *ConversationMemberRoleUpdatedEventMessageDetail, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConversationMemberRoleUpdatedEventMessageDetail
func (r *ConversationMemberRoleUpdatedEventMessageDetailRequest) Update(ctx context.Context, reqObj *ConversationMemberRoleUpdatedEventMessageDetail) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConversationMemberRoleUpdatedEventMessageDetail
func (r *ConversationMemberRoleUpdatedEventMessageDetailRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ConversationThreadRequestBuilder is request builder for ConversationThread
type ConversationThreadRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConversationThreadRequest
func (b *ConversationThreadRequestBuilder) Request() *ConversationThreadRequest {
	return &ConversationThreadRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConversationThreadRequest is request for ConversationThread
type ConversationThreadRequest struct{ BaseRequest }

// Get performs GET request for ConversationThread
func (r *ConversationThreadRequest) Get(ctx context.Context) (resObj *ConversationThread, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConversationThread
func (r *ConversationThreadRequest) Update(ctx context.Context, reqObj *ConversationThread) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConversationThread
func (r *ConversationThreadRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type ConversationMemberCollectionAddRequestBuilder struct{ BaseRequestBuilder }

// Add action undocumented
func (b *ChannelMembersCollectionRequestBuilder) Add(reqObj *ConversationMemberCollectionAddRequestParameter) *ConversationMemberCollectionAddRequestBuilder {
	bb := &ConversationMemberCollectionAddRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/add"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// Add action undocumented
func (b *ChatMembersCollectionRequestBuilder) Add(reqObj *ConversationMemberCollectionAddRequestParameter) *ConversationMemberCollectionAddRequestBuilder {
	bb := &ConversationMemberCollectionAddRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/add"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// Add action undocumented
func (b *SharedWithChannelTeamInfoAllowedMembersCollectionRequestBuilder) Add(reqObj *ConversationMemberCollectionAddRequestParameter) *ConversationMemberCollectionAddRequestBuilder {
	bb := &ConversationMemberCollectionAddRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/add"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// Add action undocumented
func (b *TeamMembersCollectionRequestBuilder) Add(reqObj *ConversationMemberCollectionAddRequestParameter) *ConversationMemberCollectionAddRequestBuilder {
	bb := &ConversationMemberCollectionAddRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/add"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ConversationMemberCollectionAddRequest struct{ BaseRequest }

func (b *ConversationMemberCollectionAddRequestBuilder) Request() *ConversationMemberCollectionAddRequest {
	return &ConversationMemberCollectionAddRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ConversationMemberCollectionAddRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ActionResultPart, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ActionResultPart
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []ActionResultPart
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

func (r *ConversationMemberCollectionAddRequest) PostN(ctx context.Context, n int) ([]ActionResultPart, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

func (r *ConversationMemberCollectionAddRequest) Post(ctx context.Context) ([]ActionResultPart, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

type ConversationThreadReplyRequestBuilder struct{ BaseRequestBuilder }

// Reply action undocumented
func (b *ConversationThreadRequestBuilder) Reply(reqObj *ConversationThreadReplyRequestParameter) *ConversationThreadReplyRequestBuilder {
	bb := &ConversationThreadReplyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/reply"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type ConversationThreadReplyRequest struct{ BaseRequest }

func (b *ConversationThreadReplyRequestBuilder) Request() *ConversationThreadReplyRequest {
	return &ConversationThreadReplyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *ConversationThreadReplyRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

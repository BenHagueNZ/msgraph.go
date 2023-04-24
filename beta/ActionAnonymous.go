// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// AnonymousGuestConversationMember returns request builder for AnonymousGuestConversationMember collection
func (b *ChannelMembersCollectionRequestBuilder) AnonymousGuestConversationMember() *ChannelMembersCollectionAnonymousGuestConversationMemberCollectionRequestBuilder {
	bb := &ChannelMembersCollectionAnonymousGuestConversationMemberCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// ChannelMembersCollectionAnonymousGuestConversationMemberCollectionRequestBuilder is request builder for AnonymousGuestConversationMember collection rcn
type ChannelMembersCollectionAnonymousGuestConversationMemberCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AnonymousGuestConversationMember collection
func (b *ChannelMembersCollectionAnonymousGuestConversationMemberCollectionRequestBuilder) Request() *ChannelMembersCollectionAnonymousGuestConversationMemberCollectionRequest {
	return &ChannelMembersCollectionAnonymousGuestConversationMemberCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AnonymousGuestConversationMember item
func (b *ChannelMembersCollectionAnonymousGuestConversationMemberCollectionRequestBuilder) ID(id string) *AnonymousGuestConversationMemberRequestBuilder {
	bb := &AnonymousGuestConversationMemberRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChannelMembersCollectionAnonymousGuestConversationMemberCollectionRequest is request for AnonymousGuestConversationMember collection
type ChannelMembersCollectionAnonymousGuestConversationMemberCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AnonymousGuestConversationMember collection
func (r *ChannelMembersCollectionAnonymousGuestConversationMemberCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AnonymousGuestConversationMember, error) {
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
	var values []AnonymousGuestConversationMember
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
			value  []AnonymousGuestConversationMember
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

// GetN performs GET request for AnonymousGuestConversationMember collection, max N pages
func (r *ChannelMembersCollectionAnonymousGuestConversationMemberCollectionRequest) GetN(ctx context.Context, n int) ([]AnonymousGuestConversationMember, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AnonymousGuestConversationMember collection
func (r *ChannelMembersCollectionAnonymousGuestConversationMemberCollectionRequest) Get(ctx context.Context) ([]AnonymousGuestConversationMember, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AnonymousGuestConversationMember collection
func (r *ChannelMembersCollectionAnonymousGuestConversationMemberCollectionRequest) Add(ctx context.Context, reqObj *AnonymousGuestConversationMember) (resObj *AnonymousGuestConversationMember, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
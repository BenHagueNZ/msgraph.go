// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// SkypeForBusinessUserConversationMember returns request builder for SkypeForBusinessUserConversationMember collection
func (b *ChannelMembersCollectionRequestBuilder) SkypeForBusinessUserConversationMember() *ChannelMembersCollectionSkypeForBusinessUserConversationMemberCollectionRequestBuilder {
	bb := &ChannelMembersCollectionSkypeForBusinessUserConversationMemberCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// ChannelMembersCollectionSkypeForBusinessUserConversationMemberCollectionRequestBuilder is request builder for SkypeForBusinessUserConversationMember collection rcn
type ChannelMembersCollectionSkypeForBusinessUserConversationMemberCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SkypeForBusinessUserConversationMember collection
func (b *ChannelMembersCollectionSkypeForBusinessUserConversationMemberCollectionRequestBuilder) Request() *ChannelMembersCollectionSkypeForBusinessUserConversationMemberCollectionRequest {
	return &ChannelMembersCollectionSkypeForBusinessUserConversationMemberCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SkypeForBusinessUserConversationMember item
func (b *ChannelMembersCollectionSkypeForBusinessUserConversationMemberCollectionRequestBuilder) ID(id string) *SkypeForBusinessUserConversationMemberRequestBuilder {
	bb := &SkypeForBusinessUserConversationMemberRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChannelMembersCollectionSkypeForBusinessUserConversationMemberCollectionRequest is request for SkypeForBusinessUserConversationMember collection
type ChannelMembersCollectionSkypeForBusinessUserConversationMemberCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SkypeForBusinessUserConversationMember collection
func (r *ChannelMembersCollectionSkypeForBusinessUserConversationMemberCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SkypeForBusinessUserConversationMember, error) {
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
	var values []SkypeForBusinessUserConversationMember
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
			value  []SkypeForBusinessUserConversationMember
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

// GetN performs GET request for SkypeForBusinessUserConversationMember collection, max N pages
func (r *ChannelMembersCollectionSkypeForBusinessUserConversationMemberCollectionRequest) GetN(ctx context.Context, n int) ([]SkypeForBusinessUserConversationMember, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SkypeForBusinessUserConversationMember collection
func (r *ChannelMembersCollectionSkypeForBusinessUserConversationMemberCollectionRequest) Get(ctx context.Context) ([]SkypeForBusinessUserConversationMember, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SkypeForBusinessUserConversationMember collection
func (r *ChannelMembersCollectionSkypeForBusinessUserConversationMemberCollectionRequest) Add(ctx context.Context, reqObj *SkypeForBusinessUserConversationMember) (resObj *SkypeForBusinessUserConversationMember, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SkypeUserConversationMember returns request builder for SkypeUserConversationMember collection
func (b *ChannelMembersCollectionRequestBuilder) SkypeUserConversationMember() *ChannelMembersCollectionSkypeUserConversationMemberCollectionRequestBuilder {
	bb := &ChannelMembersCollectionSkypeUserConversationMemberCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// ChannelMembersCollectionSkypeUserConversationMemberCollectionRequestBuilder is request builder for SkypeUserConversationMember collection rcn
type ChannelMembersCollectionSkypeUserConversationMemberCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SkypeUserConversationMember collection
func (b *ChannelMembersCollectionSkypeUserConversationMemberCollectionRequestBuilder) Request() *ChannelMembersCollectionSkypeUserConversationMemberCollectionRequest {
	return &ChannelMembersCollectionSkypeUserConversationMemberCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SkypeUserConversationMember item
func (b *ChannelMembersCollectionSkypeUserConversationMemberCollectionRequestBuilder) ID(id string) *SkypeUserConversationMemberRequestBuilder {
	bb := &SkypeUserConversationMemberRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChannelMembersCollectionSkypeUserConversationMemberCollectionRequest is request for SkypeUserConversationMember collection
type ChannelMembersCollectionSkypeUserConversationMemberCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SkypeUserConversationMember collection
func (r *ChannelMembersCollectionSkypeUserConversationMemberCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SkypeUserConversationMember, error) {
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
	var values []SkypeUserConversationMember
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
			value  []SkypeUserConversationMember
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

// GetN performs GET request for SkypeUserConversationMember collection, max N pages
func (r *ChannelMembersCollectionSkypeUserConversationMemberCollectionRequest) GetN(ctx context.Context, n int) ([]SkypeUserConversationMember, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SkypeUserConversationMember collection
func (r *ChannelMembersCollectionSkypeUserConversationMemberCollectionRequest) Get(ctx context.Context) ([]SkypeUserConversationMember, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SkypeUserConversationMember collection
func (r *ChannelMembersCollectionSkypeUserConversationMemberCollectionRequest) Add(ctx context.Context, reqObj *SkypeUserConversationMember) (resObj *SkypeUserConversationMember, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

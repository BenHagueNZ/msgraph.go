// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// ChatSendActivityNotificationRequestParameter undocumented
type ChatSendActivityNotificationRequestParameter struct {
	// Topic undocumented
	Topic *TeamworkActivityTopic `json:"topic,omitempty"`
	// ActivityType undocumented
	ActivityType *string `json:"activityType,omitempty"`
	// ChainID undocumented
	ChainID *int `json:"chainId,omitempty"`
	// PreviewText undocumented
	PreviewText *ItemBody `json:"previewText,omitempty"`
	// TemplateParameters undocumented
	TemplateParameters []KeyValuePair `json:"templateParameters,omitempty"`
	// Recipient undocumented
	Recipient *TeamworkNotificationRecipient `json:"recipient,omitempty"`
}

// ChatHideForUserRequestParameter undocumented
type ChatHideForUserRequestParameter struct {
	// User undocumented
	User *TeamworkUserIdentity `json:"user,omitempty"`
}

// ChatMarkChatReadForUserRequestParameter undocumented
type ChatMarkChatReadForUserRequestParameter struct {
	// User undocumented
	User *TeamworkUserIdentity `json:"user,omitempty"`
}

// ChatMarkChatUnreadForUserRequestParameter undocumented
type ChatMarkChatUnreadForUserRequestParameter struct {
	// User undocumented
	User *TeamworkUserIdentity `json:"user,omitempty"`
	// LastMessageReadDateTime undocumented
	LastMessageReadDateTime *time.Time `json:"lastMessageReadDateTime,omitempty"`
}

// ChatUnhideForUserRequestParameter undocumented
type ChatUnhideForUserRequestParameter struct {
	// User undocumented
	User *TeamworkUserIdentity `json:"user,omitempty"`
}

// ChatMessageSoftDeleteRequestParameter undocumented
type ChatMessageSoftDeleteRequestParameter struct {
}

// ChatMessageUndoSoftDeleteRequestParameter undocumented
type ChatMessageUndoSoftDeleteRequestParameter struct {
}

// InstalledApps returns request builder for TeamsAppInstallation collection
func (b *ChatRequestBuilder) InstalledApps() *ChatInstalledAppsCollectionRequestBuilder {
	bb := &ChatInstalledAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/installedApps"
	return bb
}

// ChatInstalledAppsCollectionRequestBuilder is request builder for TeamsAppInstallation collection
type ChatInstalledAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsAppInstallation collection
func (b *ChatInstalledAppsCollectionRequestBuilder) Request() *ChatInstalledAppsCollectionRequest {
	return &ChatInstalledAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsAppInstallation item
func (b *ChatInstalledAppsCollectionRequestBuilder) ID(id string) *TeamsAppInstallationRequestBuilder {
	bb := &TeamsAppInstallationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChatInstalledAppsCollectionRequest is request for TeamsAppInstallation collection
type ChatInstalledAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamsAppInstallation collection
func (r *ChatInstalledAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamsAppInstallation, error) {
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
	var values []TeamsAppInstallation
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
			value  []TeamsAppInstallation
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

// GetN performs GET request for TeamsAppInstallation collection, max N pages
func (r *ChatInstalledAppsCollectionRequest) GetN(ctx context.Context, n int) ([]TeamsAppInstallation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamsAppInstallation collection
func (r *ChatInstalledAppsCollectionRequest) Get(ctx context.Context) ([]TeamsAppInstallation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamsAppInstallation collection
func (r *ChatInstalledAppsCollectionRequest) Add(ctx context.Context, reqObj *TeamsAppInstallation) (resObj *TeamsAppInstallation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// LastMessagePreview is navigation property rn
func (b *ChatRequestBuilder) LastMessagePreview() *ChatMessageInfoRequestBuilder {
	bb := &ChatMessageInfoRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/lastMessagePreview"
	return bb
}

// Members returns request builder for ConversationMember collection
func (b *ChatRequestBuilder) Members() *ChatMembersCollectionRequestBuilder {
	bb := &ChatMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/members"
	return bb
}

// ChatMembersCollectionRequestBuilder is request builder for ConversationMember collection
type ChatMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ConversationMember collection
func (b *ChatMembersCollectionRequestBuilder) Request() *ChatMembersCollectionRequest {
	return &ChatMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ConversationMember item
func (b *ChatMembersCollectionRequestBuilder) ID(id string) *ConversationMemberRequestBuilder {
	bb := &ConversationMemberRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChatMembersCollectionRequest is request for ConversationMember collection
type ChatMembersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ConversationMember collection
func (r *ChatMembersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ConversationMember, error) {
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
	var values []ConversationMember
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
			value  []ConversationMember
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

// GetN performs GET request for ConversationMember collection, max N pages
func (r *ChatMembersCollectionRequest) GetN(ctx context.Context, n int) ([]ConversationMember, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ConversationMember collection
func (r *ChatMembersCollectionRequest) Get(ctx context.Context) ([]ConversationMember, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ConversationMember collection
func (r *ChatMembersCollectionRequest) Add(ctx context.Context, reqObj *ConversationMember) (resObj *ConversationMember, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Messages returns request builder for ChatMessage collection
func (b *ChatRequestBuilder) Messages() *ChatMessagesCollectionRequestBuilder {
	bb := &ChatMessagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/messages"
	return bb
}

// ChatMessagesCollectionRequestBuilder is request builder for ChatMessage collection
type ChatMessagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ChatMessage collection
func (b *ChatMessagesCollectionRequestBuilder) Request() *ChatMessagesCollectionRequest {
	return &ChatMessagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ChatMessage item
func (b *ChatMessagesCollectionRequestBuilder) ID(id string) *ChatMessageRequestBuilder {
	bb := &ChatMessageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChatMessagesCollectionRequest is request for ChatMessage collection
type ChatMessagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ChatMessage collection
func (r *ChatMessagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ChatMessage, error) {
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
	var values []ChatMessage
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
			value  []ChatMessage
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

// GetN performs GET request for ChatMessage collection, max N pages
func (r *ChatMessagesCollectionRequest) GetN(ctx context.Context, n int) ([]ChatMessage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ChatMessage collection
func (r *ChatMessagesCollectionRequest) Get(ctx context.Context) ([]ChatMessage, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ChatMessage collection
func (r *ChatMessagesCollectionRequest) Add(ctx context.Context, reqObj *ChatMessage) (resObj *ChatMessage, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// PinnedMessages returns request builder for PinnedChatMessageInfo collection
func (b *ChatRequestBuilder) PinnedMessages() *ChatPinnedMessagesCollectionRequestBuilder {
	bb := &ChatPinnedMessagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/pinnedMessages"
	return bb
}

// ChatPinnedMessagesCollectionRequestBuilder is request builder for PinnedChatMessageInfo collection
type ChatPinnedMessagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PinnedChatMessageInfo collection
func (b *ChatPinnedMessagesCollectionRequestBuilder) Request() *ChatPinnedMessagesCollectionRequest {
	return &ChatPinnedMessagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PinnedChatMessageInfo item
func (b *ChatPinnedMessagesCollectionRequestBuilder) ID(id string) *PinnedChatMessageInfoRequestBuilder {
	bb := &PinnedChatMessageInfoRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChatPinnedMessagesCollectionRequest is request for PinnedChatMessageInfo collection
type ChatPinnedMessagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PinnedChatMessageInfo collection
func (r *ChatPinnedMessagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PinnedChatMessageInfo, error) {
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
	var values []PinnedChatMessageInfo
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
			value  []PinnedChatMessageInfo
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

// GetN performs GET request for PinnedChatMessageInfo collection, max N pages
func (r *ChatPinnedMessagesCollectionRequest) GetN(ctx context.Context, n int) ([]PinnedChatMessageInfo, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PinnedChatMessageInfo collection
func (r *ChatPinnedMessagesCollectionRequest) Get(ctx context.Context) ([]PinnedChatMessageInfo, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PinnedChatMessageInfo collection
func (r *ChatPinnedMessagesCollectionRequest) Add(ctx context.Context, reqObj *PinnedChatMessageInfo) (resObj *PinnedChatMessageInfo, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Tabs returns request builder for TeamsTab collection
func (b *ChatRequestBuilder) Tabs() *ChatTabsCollectionRequestBuilder {
	bb := &ChatTabsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tabs"
	return bb
}

// ChatTabsCollectionRequestBuilder is request builder for TeamsTab collection
type ChatTabsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsTab collection
func (b *ChatTabsCollectionRequestBuilder) Request() *ChatTabsCollectionRequest {
	return &ChatTabsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsTab item
func (b *ChatTabsCollectionRequestBuilder) ID(id string) *TeamsTabRequestBuilder {
	bb := &TeamsTabRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChatTabsCollectionRequest is request for TeamsTab collection
type ChatTabsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamsTab collection
func (r *ChatTabsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamsTab, error) {
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
	var values []TeamsTab
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
			value  []TeamsTab
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

// GetN performs GET request for TeamsTab collection, max N pages
func (r *ChatTabsCollectionRequest) GetN(ctx context.Context, n int) ([]TeamsTab, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamsTab collection
func (r *ChatTabsCollectionRequest) Get(ctx context.Context) ([]TeamsTab, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamsTab collection
func (r *ChatTabsCollectionRequest) Add(ctx context.Context, reqObj *TeamsTab) (resObj *TeamsTab, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// HostedContents returns request builder for ChatMessageHostedContent collection
func (b *ChatMessageRequestBuilder) HostedContents() *ChatMessageHostedContentsCollectionRequestBuilder {
	bb := &ChatMessageHostedContentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/hostedContents"
	return bb
}

// ChatMessageHostedContentsCollectionRequestBuilder is request builder for ChatMessageHostedContent collection
type ChatMessageHostedContentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ChatMessageHostedContent collection
func (b *ChatMessageHostedContentsCollectionRequestBuilder) Request() *ChatMessageHostedContentsCollectionRequest {
	return &ChatMessageHostedContentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ChatMessageHostedContent item
func (b *ChatMessageHostedContentsCollectionRequestBuilder) ID(id string) *ChatMessageHostedContentRequestBuilder {
	bb := &ChatMessageHostedContentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChatMessageHostedContentsCollectionRequest is request for ChatMessageHostedContent collection
type ChatMessageHostedContentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ChatMessageHostedContent collection
func (r *ChatMessageHostedContentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ChatMessageHostedContent, error) {
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
	var values []ChatMessageHostedContent
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
			value  []ChatMessageHostedContent
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

// GetN performs GET request for ChatMessageHostedContent collection, max N pages
func (r *ChatMessageHostedContentsCollectionRequest) GetN(ctx context.Context, n int) ([]ChatMessageHostedContent, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ChatMessageHostedContent collection
func (r *ChatMessageHostedContentsCollectionRequest) Get(ctx context.Context) ([]ChatMessageHostedContent, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ChatMessageHostedContent collection
func (r *ChatMessageHostedContentsCollectionRequest) Add(ctx context.Context, reqObj *ChatMessageHostedContent) (resObj *ChatMessageHostedContent, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Replies returns request builder for ChatMessage collection
func (b *ChatMessageRequestBuilder) Replies() *ChatMessageRepliesCollectionRequestBuilder {
	bb := &ChatMessageRepliesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/replies"
	return bb
}

// ChatMessageRepliesCollectionRequestBuilder is request builder for ChatMessage collection
type ChatMessageRepliesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ChatMessage collection
func (b *ChatMessageRepliesCollectionRequestBuilder) Request() *ChatMessageRepliesCollectionRequest {
	return &ChatMessageRepliesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ChatMessage item
func (b *ChatMessageRepliesCollectionRequestBuilder) ID(id string) *ChatMessageRequestBuilder {
	bb := &ChatMessageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ChatMessageRepliesCollectionRequest is request for ChatMessage collection
type ChatMessageRepliesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ChatMessage collection
func (r *ChatMessageRepliesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ChatMessage, error) {
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
	var values []ChatMessage
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
			value  []ChatMessage
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

// GetN performs GET request for ChatMessage collection, max N pages
func (r *ChatMessageRepliesCollectionRequest) GetN(ctx context.Context, n int) ([]ChatMessage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ChatMessage collection
func (r *ChatMessageRepliesCollectionRequest) Get(ctx context.Context) ([]ChatMessage, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ChatMessage collection
func (r *ChatMessageRepliesCollectionRequest) Add(ctx context.Context, reqObj *ChatMessage) (resObj *ChatMessage, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *ChatRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *ChatMessageRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *ChatMessageInfoRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

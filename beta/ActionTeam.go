// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// TeamCloneRequestParameter undocumented
type TeamCloneRequestParameter struct {
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// MailNickname undocumented
	MailNickname *string `json:"mailNickname,omitempty"`
	// Classification undocumented
	Classification *string `json:"classification,omitempty"`
	// Visibility undocumented
	Visibility *TeamVisibilityType `json:"visibility,omitempty"`
	// PartsToClone undocumented
	PartsToClone *ClonableTeamParts `json:"partsToClone,omitempty"`
}

// TeamArchiveRequestParameter undocumented
type TeamArchiveRequestParameter struct {
	// ShouldSetSpoSiteReadOnlyForMembers undocumented
	ShouldSetSpoSiteReadOnlyForMembers *bool `json:"shouldSetSpoSiteReadOnlyForMembers,omitempty"`
}

// TeamUnarchiveRequestParameter undocumented
type TeamUnarchiveRequestParameter struct {
}

// TeamCompleteMigrationRequestParameter undocumented
type TeamCompleteMigrationRequestParameter struct {
}

// TeamSendActivityNotificationRequestParameter undocumented
type TeamSendActivityNotificationRequestParameter struct {
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

// AllChannels returns request builder for Channel collection
func (b *TeamRequestBuilder) AllChannels() *TeamAllChannelsCollectionRequestBuilder {
	bb := &TeamAllChannelsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/allChannels"
	return bb
}

// TeamAllChannelsCollectionRequestBuilder is request builder for Channel collection rcn
type TeamAllChannelsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Channel collection
func (b *TeamAllChannelsCollectionRequestBuilder) Request() *TeamAllChannelsCollectionRequest {
	return &TeamAllChannelsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Channel item
func (b *TeamAllChannelsCollectionRequestBuilder) ID(id string) *ChannelRequestBuilder {
	bb := &ChannelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamAllChannelsCollectionRequest is request for Channel collection
type TeamAllChannelsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Channel collection
func (r *TeamAllChannelsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Channel, error) {
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
	var values []Channel
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
			value  []Channel
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

// GetN performs GET request for Channel collection, max N pages
func (r *TeamAllChannelsCollectionRequest) GetN(ctx context.Context, n int) ([]Channel, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Channel collection
func (r *TeamAllChannelsCollectionRequest) Get(ctx context.Context) ([]Channel, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Channel collection
func (r *TeamAllChannelsCollectionRequest) Add(ctx context.Context, reqObj *Channel) (resObj *Channel, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Channels returns request builder for Channel collection
func (b *TeamRequestBuilder) Channels() *TeamChannelsCollectionRequestBuilder {
	bb := &TeamChannelsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/channels"
	return bb
}

// TeamChannelsCollectionRequestBuilder is request builder for Channel collection rcn
type TeamChannelsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Channel collection
func (b *TeamChannelsCollectionRequestBuilder) Request() *TeamChannelsCollectionRequest {
	return &TeamChannelsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Channel item
func (b *TeamChannelsCollectionRequestBuilder) ID(id string) *ChannelRequestBuilder {
	bb := &ChannelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamChannelsCollectionRequest is request for Channel collection
type TeamChannelsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Channel collection
func (r *TeamChannelsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Channel, error) {
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
	var values []Channel
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
			value  []Channel
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

// GetN performs GET request for Channel collection, max N pages
func (r *TeamChannelsCollectionRequest) GetN(ctx context.Context, n int) ([]Channel, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Channel collection
func (r *TeamChannelsCollectionRequest) Get(ctx context.Context) ([]Channel, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Channel collection
func (r *TeamChannelsCollectionRequest) Add(ctx context.Context, reqObj *Channel) (resObj *Channel, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Group is navigation property rn
func (b *TeamRequestBuilder) Group() *GroupRequestBuilder {
	bb := &GroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/group"
	return bb
}

// IncomingChannels returns request builder for Channel collection
func (b *TeamRequestBuilder) IncomingChannels() *TeamIncomingChannelsCollectionRequestBuilder {
	bb := &TeamIncomingChannelsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/incomingChannels"
	return bb
}

// TeamIncomingChannelsCollectionRequestBuilder is request builder for Channel collection rcn
type TeamIncomingChannelsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Channel collection
func (b *TeamIncomingChannelsCollectionRequestBuilder) Request() *TeamIncomingChannelsCollectionRequest {
	return &TeamIncomingChannelsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Channel item
func (b *TeamIncomingChannelsCollectionRequestBuilder) ID(id string) *ChannelRequestBuilder {
	bb := &ChannelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamIncomingChannelsCollectionRequest is request for Channel collection
type TeamIncomingChannelsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Channel collection
func (r *TeamIncomingChannelsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Channel, error) {
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
	var values []Channel
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
			value  []Channel
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

// GetN performs GET request for Channel collection, max N pages
func (r *TeamIncomingChannelsCollectionRequest) GetN(ctx context.Context, n int) ([]Channel, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Channel collection
func (r *TeamIncomingChannelsCollectionRequest) Get(ctx context.Context) ([]Channel, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Channel collection
func (r *TeamIncomingChannelsCollectionRequest) Add(ctx context.Context, reqObj *Channel) (resObj *Channel, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// InstalledApps returns request builder for TeamsAppInstallation collection
func (b *TeamRequestBuilder) InstalledApps() *TeamInstalledAppsCollectionRequestBuilder {
	bb := &TeamInstalledAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/installedApps"
	return bb
}

// TeamInstalledAppsCollectionRequestBuilder is request builder for TeamsAppInstallation collection rcn
type TeamInstalledAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsAppInstallation collection
func (b *TeamInstalledAppsCollectionRequestBuilder) Request() *TeamInstalledAppsCollectionRequest {
	return &TeamInstalledAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsAppInstallation item
func (b *TeamInstalledAppsCollectionRequestBuilder) ID(id string) *TeamsAppInstallationRequestBuilder {
	bb := &TeamsAppInstallationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamInstalledAppsCollectionRequest is request for TeamsAppInstallation collection
type TeamInstalledAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamsAppInstallation collection
func (r *TeamInstalledAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamsAppInstallation, error) {
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
func (r *TeamInstalledAppsCollectionRequest) GetN(ctx context.Context, n int) ([]TeamsAppInstallation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamsAppInstallation collection
func (r *TeamInstalledAppsCollectionRequest) Get(ctx context.Context) ([]TeamsAppInstallation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamsAppInstallation collection
func (r *TeamInstalledAppsCollectionRequest) Add(ctx context.Context, reqObj *TeamsAppInstallation) (resObj *TeamsAppInstallation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Members returns request builder for ConversationMember collection
func (b *TeamRequestBuilder) Members() *TeamMembersCollectionRequestBuilder {
	bb := &TeamMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/members"
	return bb
}

// TeamMembersCollectionRequestBuilder is request builder for ConversationMember collection rcn
type TeamMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ConversationMember collection
func (b *TeamMembersCollectionRequestBuilder) Request() *TeamMembersCollectionRequest {
	return &TeamMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ConversationMember item
func (b *TeamMembersCollectionRequestBuilder) ID(id string) *ConversationMemberRequestBuilder {
	bb := &ConversationMemberRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamMembersCollectionRequest is request for ConversationMember collection
type TeamMembersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ConversationMember collection
func (r *TeamMembersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ConversationMember, error) {
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
func (r *TeamMembersCollectionRequest) GetN(ctx context.Context, n int) ([]ConversationMember, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ConversationMember collection
func (r *TeamMembersCollectionRequest) Get(ctx context.Context) ([]ConversationMember, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ConversationMember collection
func (r *TeamMembersCollectionRequest) Add(ctx context.Context, reqObj *ConversationMember) (resObj *ConversationMember, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Operations returns request builder for TeamsAsyncOperation collection
func (b *TeamRequestBuilder) Operations() *TeamOperationsCollectionRequestBuilder {
	bb := &TeamOperationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/operations"
	return bb
}

// TeamOperationsCollectionRequestBuilder is request builder for TeamsAsyncOperation collection rcn
type TeamOperationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsAsyncOperation collection
func (b *TeamOperationsCollectionRequestBuilder) Request() *TeamOperationsCollectionRequest {
	return &TeamOperationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsAsyncOperation item
func (b *TeamOperationsCollectionRequestBuilder) ID(id string) *TeamsAsyncOperationRequestBuilder {
	bb := &TeamsAsyncOperationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamOperationsCollectionRequest is request for TeamsAsyncOperation collection
type TeamOperationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamsAsyncOperation collection
func (r *TeamOperationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamsAsyncOperation, error) {
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
	var values []TeamsAsyncOperation
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
			value  []TeamsAsyncOperation
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

// GetN performs GET request for TeamsAsyncOperation collection, max N pages
func (r *TeamOperationsCollectionRequest) GetN(ctx context.Context, n int) ([]TeamsAsyncOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamsAsyncOperation collection
func (r *TeamOperationsCollectionRequest) Get(ctx context.Context) ([]TeamsAsyncOperation, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamsAsyncOperation collection
func (r *TeamOperationsCollectionRequest) Add(ctx context.Context, reqObj *TeamsAsyncOperation) (resObj *TeamsAsyncOperation, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Owners returns request builder for User collection
func (b *TeamRequestBuilder) Owners() *TeamOwnersCollectionRequestBuilder {
	bb := &TeamOwnersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/owners"
	return bb
}

// TeamOwnersCollectionRequestBuilder is request builder for User collection rcn
type TeamOwnersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for User collection
func (b *TeamOwnersCollectionRequestBuilder) Request() *TeamOwnersCollectionRequest {
	return &TeamOwnersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for User item
func (b *TeamOwnersCollectionRequestBuilder) ID(id string) *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamOwnersCollectionRequest is request for User collection
type TeamOwnersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for User collection
func (r *TeamOwnersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]User, error) {
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
	var values []User
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
			value  []User
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

// GetN performs GET request for User collection, max N pages
func (r *TeamOwnersCollectionRequest) GetN(ctx context.Context, n int) ([]User, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for User collection
func (r *TeamOwnersCollectionRequest) Get(ctx context.Context) ([]User, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for User collection
func (r *TeamOwnersCollectionRequest) Add(ctx context.Context, reqObj *User) (resObj *User, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// PermissionGrants returns request builder for ResourceSpecificPermissionGrant collection
func (b *TeamRequestBuilder) PermissionGrants() *TeamPermissionGrantsCollectionRequestBuilder {
	bb := &TeamPermissionGrantsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/permissionGrants"
	return bb
}

// TeamPermissionGrantsCollectionRequestBuilder is request builder for ResourceSpecificPermissionGrant collection rcn
type TeamPermissionGrantsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ResourceSpecificPermissionGrant collection
func (b *TeamPermissionGrantsCollectionRequestBuilder) Request() *TeamPermissionGrantsCollectionRequest {
	return &TeamPermissionGrantsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ResourceSpecificPermissionGrant item
func (b *TeamPermissionGrantsCollectionRequestBuilder) ID(id string) *ResourceSpecificPermissionGrantRequestBuilder {
	bb := &ResourceSpecificPermissionGrantRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamPermissionGrantsCollectionRequest is request for ResourceSpecificPermissionGrant collection
type TeamPermissionGrantsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ResourceSpecificPermissionGrant collection
func (r *TeamPermissionGrantsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ResourceSpecificPermissionGrant, error) {
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
	var values []ResourceSpecificPermissionGrant
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
			value  []ResourceSpecificPermissionGrant
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

// GetN performs GET request for ResourceSpecificPermissionGrant collection, max N pages
func (r *TeamPermissionGrantsCollectionRequest) GetN(ctx context.Context, n int) ([]ResourceSpecificPermissionGrant, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ResourceSpecificPermissionGrant collection
func (r *TeamPermissionGrantsCollectionRequest) Get(ctx context.Context) ([]ResourceSpecificPermissionGrant, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ResourceSpecificPermissionGrant collection
func (r *TeamPermissionGrantsCollectionRequest) Add(ctx context.Context, reqObj *ResourceSpecificPermissionGrant) (resObj *ResourceSpecificPermissionGrant, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Photo is navigation property rn
func (b *TeamRequestBuilder) Photo() *ProfilePhotoRequestBuilder {
	bb := &ProfilePhotoRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/photo"
	return bb
}

// PrimaryChannel is navigation property rn
func (b *TeamRequestBuilder) PrimaryChannel() *ChannelRequestBuilder {
	bb := &ChannelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/primaryChannel"
	return bb
}

// Schedule is navigation property rn
func (b *TeamRequestBuilder) Schedule() *ScheduleRequestBuilder {
	bb := &ScheduleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/schedule"
	return bb
}

// Tags returns request builder for TeamworkTag collection
func (b *TeamRequestBuilder) Tags() *TeamTagsCollectionRequestBuilder {
	bb := &TeamTagsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tags"
	return bb
}

// TeamTagsCollectionRequestBuilder is request builder for TeamworkTag collection rcn
type TeamTagsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamworkTag collection
func (b *TeamTagsCollectionRequestBuilder) Request() *TeamTagsCollectionRequest {
	return &TeamTagsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamworkTag item
func (b *TeamTagsCollectionRequestBuilder) ID(id string) *TeamworkTagRequestBuilder {
	bb := &TeamworkTagRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamTagsCollectionRequest is request for TeamworkTag collection
type TeamTagsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamworkTag collection
func (r *TeamTagsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamworkTag, error) {
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
	var values []TeamworkTag
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
			value  []TeamworkTag
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

// GetN performs GET request for TeamworkTag collection, max N pages
func (r *TeamTagsCollectionRequest) GetN(ctx context.Context, n int) ([]TeamworkTag, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamworkTag collection
func (r *TeamTagsCollectionRequest) Get(ctx context.Context) ([]TeamworkTag, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamworkTag collection
func (r *TeamTagsCollectionRequest) Add(ctx context.Context, reqObj *TeamworkTag) (resObj *TeamworkTag, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Template is navigation property rn
func (b *TeamRequestBuilder) Template() *TeamsTemplateRequestBuilder {
	bb := &TeamsTemplateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/template"
	return bb
}

// TemplateDefinition is navigation property rn
func (b *TeamRequestBuilder) TemplateDefinition() *TeamTemplateDefinitionRequestBuilder {
	bb := &TeamTemplateDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/templateDefinition"
	return bb
}

// Team is navigation property rn
func (b *TeamInfoRequestBuilder) Team() *TeamRequestBuilder {
	bb := &TeamRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/team"
	return bb
}

// Definitions returns request builder for TeamTemplateDefinition collection
func (b *TeamTemplateRequestBuilder) Definitions() *TeamTemplateDefinitionsCollectionRequestBuilder {
	bb := &TeamTemplateDefinitionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/definitions"
	return bb
}

// TeamTemplateDefinitionsCollectionRequestBuilder is request builder for TeamTemplateDefinition collection rcn
type TeamTemplateDefinitionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamTemplateDefinition collection
func (b *TeamTemplateDefinitionsCollectionRequestBuilder) Request() *TeamTemplateDefinitionsCollectionRequest {
	return &TeamTemplateDefinitionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamTemplateDefinition item
func (b *TeamTemplateDefinitionsCollectionRequestBuilder) ID(id string) *TeamTemplateDefinitionRequestBuilder {
	bb := &TeamTemplateDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamTemplateDefinitionsCollectionRequest is request for TeamTemplateDefinition collection
type TeamTemplateDefinitionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamTemplateDefinition collection
func (r *TeamTemplateDefinitionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamTemplateDefinition, error) {
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
	var values []TeamTemplateDefinition
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
			value  []TeamTemplateDefinition
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

// GetN performs GET request for TeamTemplateDefinition collection, max N pages
func (r *TeamTemplateDefinitionsCollectionRequest) GetN(ctx context.Context, n int) ([]TeamTemplateDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamTemplateDefinition collection
func (r *TeamTemplateDefinitionsCollectionRequest) Get(ctx context.Context) ([]TeamTemplateDefinition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamTemplateDefinition collection
func (r *TeamTemplateDefinitionsCollectionRequest) Add(ctx context.Context, reqObj *TeamTemplateDefinition) (resObj *TeamTemplateDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// TeamDefinition is navigation property rn
func (b *TeamTemplateDefinitionRequestBuilder) TeamDefinition() *TeamRequestBuilder {
	bb := &TeamRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/teamDefinition"
	return bb
}

// Entity is navigation property rn
func (b *TeamRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TeamInfoRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TeamTemplateRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TeamTemplateDefinitionRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
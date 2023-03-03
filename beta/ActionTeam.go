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

// Apps returns request builder for TeamsCatalogApp collection
func (b *TeamRequestBuilder) Apps() *TeamAppsCollectionRequestBuilder {
	bb := &TeamAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/apps"
	return bb
}

// TeamAppsCollectionRequestBuilder is request builder for TeamsCatalogApp collection
type TeamAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsCatalogApp collection
func (b *TeamAppsCollectionRequestBuilder) Request() *TeamAppsCollectionRequest {
	return &TeamAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsCatalogApp item
func (b *TeamAppsCollectionRequestBuilder) ID(id string) *TeamsCatalogAppRequestBuilder {
	bb := &TeamsCatalogAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamAppsCollectionRequest is request for TeamsCatalogApp collection
type TeamAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamsCatalogApp collection
func (r *TeamAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamsCatalogApp, error) {
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
	var values []TeamsCatalogApp
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
			value  []TeamsCatalogApp
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

// GetN performs GET request for TeamsCatalogApp collection, max N pages
func (r *TeamAppsCollectionRequest) GetN(ctx context.Context, n int) ([]TeamsCatalogApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamsCatalogApp collection
func (r *TeamAppsCollectionRequest) Get(ctx context.Context) ([]TeamsCatalogApp, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamsCatalogApp collection
func (r *TeamAppsCollectionRequest) Add(ctx context.Context, reqObj *TeamsCatalogApp) (resObj *TeamsCatalogApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Channels returns request builder for Channel collection
func (b *TeamRequestBuilder) Channels() *TeamChannelsCollectionRequestBuilder {
	bb := &TeamChannelsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/channels"
	return bb
}

// TeamChannelsCollectionRequestBuilder is request builder for Channel collection
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

// Group is navigation property
func (b *TeamRequestBuilder) Group() *GroupRequestBuilder {
	bb := &GroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/group"
	return bb
}

// InstalledApps returns request builder for TeamsAppInstallation collection
func (b *TeamRequestBuilder) InstalledApps() *TeamInstalledAppsCollectionRequestBuilder {
	bb := &TeamInstalledAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/installedApps"
	return bb
}

// TeamInstalledAppsCollectionRequestBuilder is request builder for TeamsAppInstallation collection
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

// Operations returns request builder for TeamsAsyncOperation collection
func (b *TeamRequestBuilder) Operations() *TeamOperationsCollectionRequestBuilder {
	bb := &TeamOperationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/operations"
	return bb
}

// TeamOperationsCollectionRequestBuilder is request builder for TeamsAsyncOperation collection
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

// TeamOwnersCollectionRequestBuilder is request builder for User collection
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

// Photo is navigation property
func (b *TeamRequestBuilder) Photo() *ProfilePhotoRequestBuilder {
	bb := &ProfilePhotoRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/photo"
	return bb
}

// PrimaryChannel is navigation property
func (b *TeamRequestBuilder) PrimaryChannel() *ChannelRequestBuilder {
	bb := &ChannelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/primaryChannel"
	return bb
}

// Schedule is navigation property
func (b *TeamRequestBuilder) Schedule() *ScheduleRequestBuilder {
	bb := &ScheduleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/schedule"
	return bb
}

// Template is navigation property
func (b *TeamRequestBuilder) Template() *TeamsTemplateRequestBuilder {
	bb := &TeamsTemplateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/template"
	return bb
}

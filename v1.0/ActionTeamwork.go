// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// TeamworkSendActivityNotificationToRecipientsRequestParameter undocumented
type TeamworkSendActivityNotificationToRecipientsRequestParameter struct {
	// Topic undocumented
	Topic *TeamworkActivityTopic `json:"topic,omitempty"`
	// ActivityType undocumented
	ActivityType *string `json:"activityType,omitempty"`
	// ChainID undocumented
	ChainID *int `json:"chainId,omitempty"`
	// PreviewText undocumented
	PreviewText *ItemBody `json:"previewText,omitempty"`
	// TeamsAppID undocumented
	TeamsAppID *string `json:"teamsAppId,omitempty"`
	// TemplateParameters undocumented
	TemplateParameters []KeyValuePair `json:"templateParameters,omitempty"`
	// Recipients undocumented
	Recipients []TeamworkNotificationRecipient `json:"recipients,omitempty"`
}

// DeletedTeams returns request builder for DeletedTeam collection rcn
func (b *TeamworkRequestBuilder) DeletedTeams() *TeamworkDeletedTeamsCollectionRequestBuilder {
	bb := &TeamworkDeletedTeamsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deletedTeams"
	return bb
}

// TeamworkDeletedTeamsCollectionRequestBuilder is request builder for DeletedTeam collection
type TeamworkDeletedTeamsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeletedTeam collection
func (b *TeamworkDeletedTeamsCollectionRequestBuilder) Request() *TeamworkDeletedTeamsCollectionRequest {
	return &TeamworkDeletedTeamsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeletedTeam item
func (b *TeamworkDeletedTeamsCollectionRequestBuilder) ID(id string) *DeletedTeamRequestBuilder {
	bb := &DeletedTeamRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamworkDeletedTeamsCollectionRequest is request for DeletedTeam collection
type TeamworkDeletedTeamsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeletedTeam collection
func (r *TeamworkDeletedTeamsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DeletedTeam, error) {
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
	var values []DeletedTeam
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
			value  []DeletedTeam
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

// GetN performs GET request for DeletedTeam collection, max N pages
func (r *TeamworkDeletedTeamsCollectionRequest) GetN(ctx context.Context, n int) ([]DeletedTeam, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DeletedTeam collection
func (r *TeamworkDeletedTeamsCollectionRequest) Get(ctx context.Context) ([]DeletedTeam, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DeletedTeam collection
func (r *TeamworkDeletedTeamsCollectionRequest) Add(ctx context.Context, reqObj *DeletedTeam) (resObj *DeletedTeam, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// WorkforceIntegrations returns request builder for WorkforceIntegration collection rcn
func (b *TeamworkRequestBuilder) WorkforceIntegrations() *TeamworkWorkforceIntegrationsCollectionRequestBuilder {
	bb := &TeamworkWorkforceIntegrationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/workforceIntegrations"
	return bb
}

// TeamworkWorkforceIntegrationsCollectionRequestBuilder is request builder for WorkforceIntegration collection
type TeamworkWorkforceIntegrationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WorkforceIntegration collection
func (b *TeamworkWorkforceIntegrationsCollectionRequestBuilder) Request() *TeamworkWorkforceIntegrationsCollectionRequest {
	return &TeamworkWorkforceIntegrationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WorkforceIntegration item
func (b *TeamworkWorkforceIntegrationsCollectionRequestBuilder) ID(id string) *WorkforceIntegrationRequestBuilder {
	bb := &WorkforceIntegrationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamworkWorkforceIntegrationsCollectionRequest is request for WorkforceIntegration collection
type TeamworkWorkforceIntegrationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WorkforceIntegration collection
func (r *TeamworkWorkforceIntegrationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]WorkforceIntegration, error) {
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
	var values []WorkforceIntegration
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
			value  []WorkforceIntegration
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

// GetN performs GET request for WorkforceIntegration collection, max N pages
func (r *TeamworkWorkforceIntegrationsCollectionRequest) GetN(ctx context.Context, n int) ([]WorkforceIntegration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for WorkforceIntegration collection
func (r *TeamworkWorkforceIntegrationsCollectionRequest) Get(ctx context.Context) ([]WorkforceIntegration, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for WorkforceIntegration collection
func (r *TeamworkWorkforceIntegrationsCollectionRequest) Add(ctx context.Context, reqObj *WorkforceIntegration) (resObj *WorkforceIntegration, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Members returns request builder for TeamworkTagMember collection rcn
func (b *TeamworkTagRequestBuilder) Members() *TeamworkTagMembersCollectionRequestBuilder {
	bb := &TeamworkTagMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/members"
	return bb
}

// TeamworkTagMembersCollectionRequestBuilder is request builder for TeamworkTagMember collection
type TeamworkTagMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamworkTagMember collection
func (b *TeamworkTagMembersCollectionRequestBuilder) Request() *TeamworkTagMembersCollectionRequest {
	return &TeamworkTagMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamworkTagMember item
func (b *TeamworkTagMembersCollectionRequestBuilder) ID(id string) *TeamworkTagMemberRequestBuilder {
	bb := &TeamworkTagMemberRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TeamworkTagMembersCollectionRequest is request for TeamworkTagMember collection
type TeamworkTagMembersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamworkTagMember collection
func (r *TeamworkTagMembersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TeamworkTagMember, error) {
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
	var values []TeamworkTagMember
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
			value  []TeamworkTagMember
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

// GetN performs GET request for TeamworkTagMember collection, max N pages
func (r *TeamworkTagMembersCollectionRequest) GetN(ctx context.Context, n int) ([]TeamworkTagMember, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TeamworkTagMember collection
func (r *TeamworkTagMembersCollectionRequest) Get(ctx context.Context) ([]TeamworkTagMember, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TeamworkTagMember collection
func (r *TeamworkTagMembersCollectionRequest) Add(ctx context.Context, reqObj *TeamworkTagMember) (resObj *TeamworkTagMember, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Teamwork is navigation property rn
func (b *TeamworkRequestBuilder) Teamwork() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// TeamworkBot is navigation property rn
func (b *TeamworkBotRequestBuilder) TeamworkBot() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// TeamworkHostedContent is navigation property rn
func (b *TeamworkHostedContentRequestBuilder) TeamworkHostedContent() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// TeamworkTag is navigation property rn
func (b *TeamworkTagRequestBuilder) TeamworkTag() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// TeamworkTagMember is navigation property rn
func (b *TeamworkTagMemberRequestBuilder) TeamworkTagMember() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

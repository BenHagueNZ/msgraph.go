// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// GroupRequestBuilder is request builder for Group
type GroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns GroupRequest
func (b *GroupRequestBuilder) Request() *GroupRequest {
	return &GroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GroupRequest is request for Group
type GroupRequest struct{ BaseRequest }

// Get performs GET request for Group
func (r *GroupRequest) Get(ctx context.Context) (resObj *Group, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Group
func (r *GroupRequest) Update(ctx context.Context, reqObj *Group) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Group
func (r *GroupRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// GroupAssignmentTargetRequestBuilder is request builder for GroupAssignmentTarget
type GroupAssignmentTargetRequestBuilder struct{ BaseRequestBuilder }

// Request returns GroupAssignmentTargetRequest
func (b *GroupAssignmentTargetRequestBuilder) Request() *GroupAssignmentTargetRequest {
	return &GroupAssignmentTargetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GroupAssignmentTargetRequest is request for GroupAssignmentTarget
type GroupAssignmentTargetRequest struct{ BaseRequest }

// Get performs GET request for GroupAssignmentTarget
func (r *GroupAssignmentTargetRequest) Get(ctx context.Context) (resObj *GroupAssignmentTarget, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for GroupAssignmentTarget
func (r *GroupAssignmentTargetRequest) Update(ctx context.Context, reqObj *GroupAssignmentTarget) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for GroupAssignmentTarget
func (r *GroupAssignmentTargetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// GroupLifecyclePolicyRequestBuilder is request builder for GroupLifecyclePolicy
type GroupLifecyclePolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns GroupLifecyclePolicyRequest
func (b *GroupLifecyclePolicyRequestBuilder) Request() *GroupLifecyclePolicyRequest {
	return &GroupLifecyclePolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GroupLifecyclePolicyRequest is request for GroupLifecyclePolicy
type GroupLifecyclePolicyRequest struct{ BaseRequest }

// Get performs GET request for GroupLifecyclePolicy
func (r *GroupLifecyclePolicyRequest) Get(ctx context.Context) (resObj *GroupLifecyclePolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for GroupLifecyclePolicy
func (r *GroupLifecyclePolicyRequest) Update(ctx context.Context, reqObj *GroupLifecyclePolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for GroupLifecyclePolicy
func (r *GroupLifecyclePolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// GroupMembersRequestBuilder is request builder for GroupMembers
type GroupMembersRequestBuilder struct{ BaseRequestBuilder }

// Request returns GroupMembersRequest
func (b *GroupMembersRequestBuilder) Request() *GroupMembersRequest {
	return &GroupMembersRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GroupMembersRequest is request for GroupMembers
type GroupMembersRequest struct{ BaseRequest }

// Get performs GET request for GroupMembers
func (r *GroupMembersRequest) Get(ctx context.Context) (resObj *GroupMembers, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for GroupMembers
func (r *GroupMembersRequest) Update(ctx context.Context, reqObj *GroupMembers) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for GroupMembers
func (r *GroupMembersRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// GroupSettingRequestBuilder is request builder for GroupSetting
type GroupSettingRequestBuilder struct{ BaseRequestBuilder }

// Request returns GroupSettingRequest
func (b *GroupSettingRequestBuilder) Request() *GroupSettingRequest {
	return &GroupSettingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GroupSettingRequest is request for GroupSetting
type GroupSettingRequest struct{ BaseRequest }

// Get performs GET request for GroupSetting
func (r *GroupSettingRequest) Get(ctx context.Context) (resObj *GroupSetting, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for GroupSetting
func (r *GroupSettingRequest) Update(ctx context.Context, reqObj *GroupSetting) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for GroupSetting
func (r *GroupSettingRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// GroupSettingTemplateRequestBuilder is request builder for GroupSettingTemplate
type GroupSettingTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns GroupSettingTemplateRequest
func (b *GroupSettingTemplateRequestBuilder) Request() *GroupSettingTemplateRequest {
	return &GroupSettingTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GroupSettingTemplateRequest is request for GroupSettingTemplate
type GroupSettingTemplateRequest struct{ BaseRequest }

// Get performs GET request for GroupSettingTemplate
func (r *GroupSettingTemplateRequest) Get(ctx context.Context) (resObj *GroupSettingTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for GroupSettingTemplate
func (r *GroupSettingTemplateRequest) Update(ctx context.Context, reqObj *GroupSettingTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for GroupSettingTemplate
func (r *GroupSettingTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type GroupAssignLicenseRequestBuilder struct{ BaseRequestBuilder }

// AssignLicense action undocumentedras
func (b *GroupRequestBuilder) AssignLicense(reqObj *GroupAssignLicenseRequestParameter) *GroupAssignLicenseRequestBuilder {
	bb := &GroupAssignLicenseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AssignLicense"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type GroupAssignLicenseRequest struct{ BaseRequest }

func (b *GroupAssignLicenseRequestBuilder) Request() *GroupAssignLicenseRequest {
	return &GroupAssignLicenseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *GroupAssignLicenseRequest) Post(ctx context.Context) (resObj *Group, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type GroupCheckGrantedPermissionsForAppRequestBuilder struct{ BaseRequestBuilder }

// CheckGrantedPermissionsForApp action undocumentedrac
func (b *GroupRequestBuilder) CheckGrantedPermissionsForApp(reqObj *GroupCheckGrantedPermissionsForAppRequestParameter) *GroupCheckGrantedPermissionsForAppRequestBuilder {
	bb := &GroupCheckGrantedPermissionsForAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CheckGrantedPermissionsForApp"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type GroupCheckGrantedPermissionsForAppRequest struct{ BaseRequest }

func (b *GroupCheckGrantedPermissionsForAppRequestBuilder) Request() *GroupCheckGrantedPermissionsForAppRequest {
	return &GroupCheckGrantedPermissionsForAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *GroupCheckGrantedPermissionsForAppRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ResourceSpecificPermissionGrant, error) {
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

func (r *GroupCheckGrantedPermissionsForAppRequest) PostN(ctx context.Context, n int) ([]ResourceSpecificPermissionGrant, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

func (r *GroupCheckGrantedPermissionsForAppRequest) Post(ctx context.Context) ([]ResourceSpecificPermissionGrant, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

type GroupValidatePropertiesRequestBuilder struct{ BaseRequestBuilder }

// ValidateProperties action undocumentedrav
func (b *GroupRequestBuilder) ValidateProperties(reqObj *GroupValidatePropertiesRequestParameter) *GroupValidatePropertiesRequestBuilder {
	bb := &GroupValidatePropertiesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ValidateProperties"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type GroupValidatePropertiesRequest struct{ BaseRequest }

func (b *GroupValidatePropertiesRequestBuilder) Request() *GroupValidatePropertiesRequest {
	return &GroupValidatePropertiesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *GroupValidatePropertiesRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type GroupAddFavoriteRequestBuilder struct{ BaseRequestBuilder }

// AddFavorite action undocumentedrav
func (b *GroupRequestBuilder) AddFavorite(reqObj *GroupAddFavoriteRequestParameter) *GroupAddFavoriteRequestBuilder {
	bb := &GroupAddFavoriteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AddFavorite"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type GroupAddFavoriteRequest struct{ BaseRequest }

func (b *GroupAddFavoriteRequestBuilder) Request() *GroupAddFavoriteRequest {
	return &GroupAddFavoriteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *GroupAddFavoriteRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type GroupRemoveFavoriteRequestBuilder struct{ BaseRequestBuilder }

// RemoveFavorite action undocumentedrav
func (b *GroupRequestBuilder) RemoveFavorite(reqObj *GroupRemoveFavoriteRequestParameter) *GroupRemoveFavoriteRequestBuilder {
	bb := &GroupRemoveFavoriteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RemoveFavorite"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type GroupRemoveFavoriteRequest struct{ BaseRequest }

func (b *GroupRemoveFavoriteRequestBuilder) Request() *GroupRemoveFavoriteRequest {
	return &GroupRemoveFavoriteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *GroupRemoveFavoriteRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type GroupResetUnseenCountRequestBuilder struct{ BaseRequestBuilder }

// ResetUnseenCount action undocumentedrav
func (b *GroupRequestBuilder) ResetUnseenCount(reqObj *GroupResetUnseenCountRequestParameter) *GroupResetUnseenCountRequestBuilder {
	bb := &GroupResetUnseenCountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ResetUnseenCount"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type GroupResetUnseenCountRequest struct{ BaseRequest }

func (b *GroupResetUnseenCountRequestBuilder) Request() *GroupResetUnseenCountRequest {
	return &GroupResetUnseenCountRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *GroupResetUnseenCountRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type GroupSubscribeByMailRequestBuilder struct{ BaseRequestBuilder }

// SubscribeByMail action undocumentedrav
func (b *GroupRequestBuilder) SubscribeByMail(reqObj *GroupSubscribeByMailRequestParameter) *GroupSubscribeByMailRequestBuilder {
	bb := &GroupSubscribeByMailRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/SubscribeByMail"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type GroupSubscribeByMailRequest struct{ BaseRequest }

func (b *GroupSubscribeByMailRequestBuilder) Request() *GroupSubscribeByMailRequest {
	return &GroupSubscribeByMailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *GroupSubscribeByMailRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type GroupUnsubscribeByMailRequestBuilder struct{ BaseRequestBuilder }

// UnsubscribeByMail action undocumentedrav
func (b *GroupRequestBuilder) UnsubscribeByMail(reqObj *GroupUnsubscribeByMailRequestParameter) *GroupUnsubscribeByMailRequestBuilder {
	bb := &GroupUnsubscribeByMailRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/UnsubscribeByMail"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type GroupUnsubscribeByMailRequest struct{ BaseRequest }

func (b *GroupUnsubscribeByMailRequestBuilder) Request() *GroupUnsubscribeByMailRequest {
	return &GroupUnsubscribeByMailRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *GroupUnsubscribeByMailRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type GroupRenewRequestBuilder struct{ BaseRequestBuilder }

// Renew action undocumentedrav
func (b *GroupRequestBuilder) Renew(reqObj *GroupRenewRequestParameter) *GroupRenewRequestBuilder {
	bb := &GroupRenewRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Renew"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type GroupRenewRequest struct{ BaseRequest }

func (b *GroupRenewRequestBuilder) Request() *GroupRenewRequest {
	return &GroupRenewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *GroupRenewRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type GroupLifecyclePolicyAddGroupRequestBuilder struct{ BaseRequestBuilder }

// AddGroup action undocumentedras
func (b *GroupLifecyclePolicyRequestBuilder) AddGroup(reqObj *GroupLifecyclePolicyAddGroupRequestParameter) *GroupLifecyclePolicyAddGroupRequestBuilder {
	bb := &GroupLifecyclePolicyAddGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/AddGroup"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type GroupLifecyclePolicyAddGroupRequest struct{ BaseRequest }

func (b *GroupLifecyclePolicyAddGroupRequestBuilder) Request() *GroupLifecyclePolicyAddGroupRequest {
	return &GroupLifecyclePolicyAddGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *GroupLifecyclePolicyAddGroupRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type GroupLifecyclePolicyRemoveGroupRequestBuilder struct{ BaseRequestBuilder }

// RemoveGroup action undocumentedras
func (b *GroupLifecyclePolicyRequestBuilder) RemoveGroup(reqObj *GroupLifecyclePolicyRemoveGroupRequestParameter) *GroupLifecyclePolicyRemoveGroupRequestBuilder {
	bb := &GroupLifecyclePolicyRemoveGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RemoveGroup"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type GroupLifecyclePolicyRemoveGroupRequest struct{ BaseRequest }

func (b *GroupLifecyclePolicyRemoveGroupRequestBuilder) Request() *GroupLifecyclePolicyRemoveGroupRequest {
	return &GroupLifecyclePolicyRemoveGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *GroupLifecyclePolicyRemoveGroupRequest) Post(ctx context.Context) (resObj *bool, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

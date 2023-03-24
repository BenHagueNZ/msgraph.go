// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// DriveRequestBuilder is request builder for Drive
type DriveRequestBuilder struct{ BaseRequestBuilder }

// Request returns DriveRequest
func (b *DriveRequestBuilder) Request() *DriveRequest {
	return &DriveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DriveRequest is request for Drive
type DriveRequest struct{ BaseRequest }

// Get performs GET request for Drive
func (r *DriveRequest) Get(ctx context.Context) (resObj *Drive, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Drive
func (r *DriveRequest) Update(ctx context.Context, reqObj *Drive) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Drive
func (r *DriveRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DriveItemRequestBuilder is request builder for DriveItem
type DriveItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns DriveItemRequest
func (b *DriveItemRequestBuilder) Request() *DriveItemRequest {
	return &DriveItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DriveItemRequest is request for DriveItem
type DriveItemRequest struct{ BaseRequest }

// Get performs GET request for DriveItem
func (r *DriveItemRequest) Get(ctx context.Context) (resObj *DriveItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DriveItem
func (r *DriveItemRequest) Update(ctx context.Context, reqObj *DriveItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DriveItem
func (r *DriveItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DriveItemUploadablePropertiesRequestBuilder is request builder for DriveItemUploadableProperties
type DriveItemUploadablePropertiesRequestBuilder struct{ BaseRequestBuilder }

// Request returns DriveItemUploadablePropertiesRequest
func (b *DriveItemUploadablePropertiesRequestBuilder) Request() *DriveItemUploadablePropertiesRequest {
	return &DriveItemUploadablePropertiesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DriveItemUploadablePropertiesRequest is request for DriveItemUploadableProperties
type DriveItemUploadablePropertiesRequest struct{ BaseRequest }

// Get performs GET request for DriveItemUploadableProperties
func (r *DriveItemUploadablePropertiesRequest) Get(ctx context.Context) (resObj *DriveItemUploadableProperties, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DriveItemUploadableProperties
func (r *DriveItemUploadablePropertiesRequest) Update(ctx context.Context, reqObj *DriveItemUploadableProperties) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DriveItemUploadableProperties
func (r *DriveItemUploadablePropertiesRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DriveItemVersionRequestBuilder is request builder for DriveItemVersion
type DriveItemVersionRequestBuilder struct{ BaseRequestBuilder }

// Request returns DriveItemVersionRequest
func (b *DriveItemVersionRequestBuilder) Request() *DriveItemVersionRequest {
	return &DriveItemVersionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DriveItemVersionRequest is request for DriveItemVersion
type DriveItemVersionRequest struct{ BaseRequest }

// Get performs GET request for DriveItemVersion
func (r *DriveItemVersionRequest) Get(ctx context.Context) (resObj *DriveItemVersion, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DriveItemVersion
func (r *DriveItemVersionRequest) Update(ctx context.Context, reqObj *DriveItemVersion) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DriveItemVersion
func (r *DriveItemVersionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DriveRecipientRequestBuilder is request builder for DriveRecipient
type DriveRecipientRequestBuilder struct{ BaseRequestBuilder }

// Request returns DriveRecipientRequest
func (b *DriveRecipientRequestBuilder) Request() *DriveRecipientRequest {
	return &DriveRecipientRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DriveRecipientRequest is request for DriveRecipient
type DriveRecipientRequest struct{ BaseRequest }

// Get performs GET request for DriveRecipient
func (r *DriveRecipientRequest) Get(ctx context.Context) (resObj *DriveRecipient, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DriveRecipient
func (r *DriveRecipientRequest) Update(ctx context.Context, reqObj *DriveRecipient) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DriveRecipient
func (r *DriveRecipientRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

type DriveItemRestoreRequestBuilder struct{ BaseRequestBuilder }

// Restore action undocumented
func (b *DriveItemRequestBuilder) Restore(reqObj *DriveItemRestoreRequestParameter) *DriveItemRestoreRequestBuilder {
	bb := &DriveItemRestoreRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Restore"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DriveItemRestoreRequest struct{ BaseRequest }

func (b *DriveItemRestoreRequestBuilder) Request() *DriveItemRestoreRequest {
	return &DriveItemRestoreRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DriveItemRestoreRequest) Post(ctx context.Context) (resObj *DriveItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type DriveItemCopyRequestBuilder struct{ BaseRequestBuilder }

// Copy action undocumented
func (b *DriveItemRequestBuilder) Copy(reqObj *DriveItemCopyRequestParameter) *DriveItemCopyRequestBuilder {
	bb := &DriveItemCopyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Copy"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DriveItemCopyRequest struct{ BaseRequest }

func (b *DriveItemCopyRequestBuilder) Request() *DriveItemCopyRequest {
	return &DriveItemCopyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DriveItemCopyRequest) Post(ctx context.Context) (resObj *DriveItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type DriveItemCreateUploadSessionRequestBuilder struct{ BaseRequestBuilder }

// CreateUploadSession action undocumented
func (b *DriveItemRequestBuilder) CreateUploadSession(reqObj *DriveItemCreateUploadSessionRequestParameter) *DriveItemCreateUploadSessionRequestBuilder {
	bb := &DriveItemCreateUploadSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CreateUploadSession"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DriveItemCreateUploadSessionRequest struct{ BaseRequest }

func (b *DriveItemCreateUploadSessionRequestBuilder) Request() *DriveItemCreateUploadSessionRequest {
	return &DriveItemCreateUploadSessionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DriveItemCreateUploadSessionRequest) Post(ctx context.Context) (resObj *UploadSession, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type DriveItemCheckinRequestBuilder struct{ BaseRequestBuilder }

// Checkin action undocumented
func (b *DriveItemRequestBuilder) Checkin(reqObj *DriveItemCheckinRequestParameter) *DriveItemCheckinRequestBuilder {
	bb := &DriveItemCheckinRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Checkin"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DriveItemCheckinRequest struct{ BaseRequest }

func (b *DriveItemCheckinRequestBuilder) Request() *DriveItemCheckinRequest {
	return &DriveItemCheckinRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DriveItemCheckinRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type DriveItemCheckoutRequestBuilder struct{ BaseRequestBuilder }

// Checkout action undocumented
func (b *DriveItemRequestBuilder) Checkout(reqObj *DriveItemCheckoutRequestParameter) *DriveItemCheckoutRequestBuilder {
	bb := &DriveItemCheckoutRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Checkout"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DriveItemCheckoutRequest struct{ BaseRequest }

func (b *DriveItemCheckoutRequestBuilder) Request() *DriveItemCheckoutRequest {
	return &DriveItemCheckoutRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DriveItemCheckoutRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type DriveItemCreateLinkRequestBuilder struct{ BaseRequestBuilder }

// CreateLink action undocumented
func (b *DriveItemRequestBuilder) CreateLink(reqObj *DriveItemCreateLinkRequestParameter) *DriveItemCreateLinkRequestBuilder {
	bb := &DriveItemCreateLinkRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/CreateLink"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DriveItemCreateLinkRequest struct{ BaseRequest }

func (b *DriveItemCreateLinkRequestBuilder) Request() *DriveItemCreateLinkRequest {
	return &DriveItemCreateLinkRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DriveItemCreateLinkRequest) Post(ctx context.Context) (resObj *Permission, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type DriveItemFollowRequestBuilder struct{ BaseRequestBuilder }

// Follow action undocumented
func (b *DriveItemRequestBuilder) Follow(reqObj *DriveItemFollowRequestParameter) *DriveItemFollowRequestBuilder {
	bb := &DriveItemFollowRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Follow"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DriveItemFollowRequest struct{ BaseRequest }

func (b *DriveItemFollowRequestBuilder) Request() *DriveItemFollowRequest {
	return &DriveItemFollowRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DriveItemFollowRequest) Post(ctx context.Context) (resObj *DriveItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type DriveItemInviteRequestBuilder struct{ BaseRequestBuilder }

// Invite action undocumented
func (b *DriveItemRequestBuilder) Invite(reqObj *DriveItemInviteRequestParameter) *DriveItemInviteRequestBuilder {
	bb := &DriveItemInviteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Invite"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DriveItemInviteRequest struct{ BaseRequest }

func (b *DriveItemInviteRequestBuilder) Request() *DriveItemInviteRequest {
	return &DriveItemInviteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DriveItemInviteRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Permission, error) {
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
	var values []Permission
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
			value  []Permission
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

func (r *DriveItemInviteRequest) PostN(ctx context.Context, n int) ([]Permission, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, n)
}

func (r *DriveItemInviteRequest) Post(ctx context.Context) ([]Permission, error) {
	return r.Paging(ctx, "POST", "", r.requestObject, 0)
}

type DriveItemPreviewRequestBuilder struct{ BaseRequestBuilder }

// Preview action undocumented
func (b *DriveItemRequestBuilder) Preview(reqObj *DriveItemPreviewRequestParameter) *DriveItemPreviewRequestBuilder {
	bb := &DriveItemPreviewRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Preview"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DriveItemPreviewRequest struct{ BaseRequest }

func (b *DriveItemPreviewRequestBuilder) Request() *DriveItemPreviewRequest {
	return &DriveItemPreviewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DriveItemPreviewRequest) Post(ctx context.Context) (resObj *ItemPreviewInfo, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

type DriveItemUnfollowRequestBuilder struct{ BaseRequestBuilder }

// Unfollow action undocumented
func (b *DriveItemRequestBuilder) Unfollow(reqObj *DriveItemUnfollowRequestParameter) *DriveItemUnfollowRequestBuilder {
	bb := &DriveItemUnfollowRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/Unfollow"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DriveItemUnfollowRequest struct{ BaseRequest }

func (b *DriveItemUnfollowRequestBuilder) Request() *DriveItemUnfollowRequest {
	return &DriveItemUnfollowRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DriveItemUnfollowRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type DriveItemValidatePermissionRequestBuilder struct{ BaseRequestBuilder }

// ValidatePermission action undocumented
func (b *DriveItemRequestBuilder) ValidatePermission(reqObj *DriveItemValidatePermissionRequestParameter) *DriveItemValidatePermissionRequestBuilder {
	bb := &DriveItemValidatePermissionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/ValidatePermission"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DriveItemValidatePermissionRequest struct{ BaseRequest }

func (b *DriveItemValidatePermissionRequestBuilder) Request() *DriveItemValidatePermissionRequest {
	return &DriveItemValidatePermissionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DriveItemValidatePermissionRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

type DriveItemVersionRestoreVersionRequestBuilder struct{ BaseRequestBuilder }

// RestoreVersion action undocumented
func (b *DriveItemVersionRequestBuilder) RestoreVersion(reqObj *DriveItemVersionRestoreVersionRequestParameter) *DriveItemVersionRestoreVersionRequestBuilder {
	bb := &DriveItemVersionRestoreVersionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/RestoreVersion"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

type DriveItemVersionRestoreVersionRequest struct{ BaseRequest }

func (b *DriveItemVersionRestoreVersionRequestBuilder) Request() *DriveItemVersionRestoreVersionRequest {
	return &DriveItemVersionRestoreVersionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

func (r *DriveItemVersionRestoreVersionRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

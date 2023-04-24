// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// NotificationMessageTemplateSendTestMessageRequestParameter undocumented
type NotificationMessageTemplateSendTestMessageRequestParameter struct {
}

// LocalizedNotificationMessages returns request builder for LocalizedNotificationMessage collection
func (b *NotificationMessageTemplateRequestBuilder) LocalizedNotificationMessages() *NotificationMessageTemplateLocalizedNotificationMessagesCollectionRequestBuilder {
	bb := &NotificationMessageTemplateLocalizedNotificationMessagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/localizedNotificationMessages"
	return bb
}

// NotificationMessageTemplateLocalizedNotificationMessagesCollectionRequestBuilder is request builder for LocalizedNotificationMessage collection rcn
type NotificationMessageTemplateLocalizedNotificationMessagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for LocalizedNotificationMessage collection
func (b *NotificationMessageTemplateLocalizedNotificationMessagesCollectionRequestBuilder) Request() *NotificationMessageTemplateLocalizedNotificationMessagesCollectionRequest {
	return &NotificationMessageTemplateLocalizedNotificationMessagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for LocalizedNotificationMessage item
func (b *NotificationMessageTemplateLocalizedNotificationMessagesCollectionRequestBuilder) ID(id string) *LocalizedNotificationMessageRequestBuilder {
	bb := &LocalizedNotificationMessageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// NotificationMessageTemplateLocalizedNotificationMessagesCollectionRequest is request for LocalizedNotificationMessage collection
type NotificationMessageTemplateLocalizedNotificationMessagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for LocalizedNotificationMessage collection
func (r *NotificationMessageTemplateLocalizedNotificationMessagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]LocalizedNotificationMessage, error) {
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
	var values []LocalizedNotificationMessage
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
			value  []LocalizedNotificationMessage
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

// GetN performs GET request for LocalizedNotificationMessage collection, max N pages
func (r *NotificationMessageTemplateLocalizedNotificationMessagesCollectionRequest) GetN(ctx context.Context, n int) ([]LocalizedNotificationMessage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for LocalizedNotificationMessage collection
func (r *NotificationMessageTemplateLocalizedNotificationMessagesCollectionRequest) Get(ctx context.Context) ([]LocalizedNotificationMessage, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for LocalizedNotificationMessage collection
func (r *NotificationMessageTemplateLocalizedNotificationMessagesCollectionRequest) Add(ctx context.Context, reqObj *LocalizedNotificationMessage) (resObj *LocalizedNotificationMessage, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *NotificationRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *NotificationMessageTemplateRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

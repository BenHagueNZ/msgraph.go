// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// Lists returns request builder for TodoTaskList collection
func (b *TodoRequestBuilder) Lists() *TodoListsCollectionRequestBuilder {
	bb := &TodoListsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/lists"
	return bb
}

// TodoListsCollectionRequestBuilder is request builder for TodoTaskList collection
type TodoListsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TodoTaskList collection
func (b *TodoListsCollectionRequestBuilder) Request() *TodoListsCollectionRequest {
	return &TodoListsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TodoTaskList item
func (b *TodoListsCollectionRequestBuilder) ID(id string) *TodoTaskListRequestBuilder {
	bb := &TodoTaskListRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TodoListsCollectionRequest is request for TodoTaskList collection
type TodoListsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TodoTaskList collection
func (r *TodoListsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TodoTaskList, error) {
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
	var values []TodoTaskList
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
			value  []TodoTaskList
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

// GetN performs GET request for TodoTaskList collection, max N pages
func (r *TodoListsCollectionRequest) GetN(ctx context.Context, n int) ([]TodoTaskList, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TodoTaskList collection
func (r *TodoListsCollectionRequest) Get(ctx context.Context) ([]TodoTaskList, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TodoTaskList collection
func (r *TodoListsCollectionRequest) Add(ctx context.Context, reqObj *TodoTaskList) (resObj *TodoTaskList, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AttachmentSessions returns request builder for AttachmentSession collection
func (b *TodoTaskRequestBuilder) AttachmentSessions() *TodoTaskAttachmentSessionsCollectionRequestBuilder {
	bb := &TodoTaskAttachmentSessionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/attachmentSessions"
	return bb
}

// TodoTaskAttachmentSessionsCollectionRequestBuilder is request builder for AttachmentSession collection
type TodoTaskAttachmentSessionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AttachmentSession collection
func (b *TodoTaskAttachmentSessionsCollectionRequestBuilder) Request() *TodoTaskAttachmentSessionsCollectionRequest {
	return &TodoTaskAttachmentSessionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AttachmentSession item
func (b *TodoTaskAttachmentSessionsCollectionRequestBuilder) ID(id string) *AttachmentSessionRequestBuilder {
	bb := &AttachmentSessionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TodoTaskAttachmentSessionsCollectionRequest is request for AttachmentSession collection
type TodoTaskAttachmentSessionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AttachmentSession collection
func (r *TodoTaskAttachmentSessionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AttachmentSession, error) {
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
	var values []AttachmentSession
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
			value  []AttachmentSession
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

// GetN performs GET request for AttachmentSession collection, max N pages
func (r *TodoTaskAttachmentSessionsCollectionRequest) GetN(ctx context.Context, n int) ([]AttachmentSession, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AttachmentSession collection
func (r *TodoTaskAttachmentSessionsCollectionRequest) Get(ctx context.Context) ([]AttachmentSession, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AttachmentSession collection
func (r *TodoTaskAttachmentSessionsCollectionRequest) Add(ctx context.Context, reqObj *AttachmentSession) (resObj *AttachmentSession, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Attachments returns request builder for AttachmentBase collection
func (b *TodoTaskRequestBuilder) Attachments() *TodoTaskAttachmentsCollectionRequestBuilder {
	bb := &TodoTaskAttachmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/attachments"
	return bb
}

// TodoTaskAttachmentsCollectionRequestBuilder is request builder for AttachmentBase collection
type TodoTaskAttachmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AttachmentBase collection
func (b *TodoTaskAttachmentsCollectionRequestBuilder) Request() *TodoTaskAttachmentsCollectionRequest {
	return &TodoTaskAttachmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AttachmentBase item
func (b *TodoTaskAttachmentsCollectionRequestBuilder) ID(id string) *AttachmentBaseRequestBuilder {
	bb := &AttachmentBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TodoTaskAttachmentsCollectionRequest is request for AttachmentBase collection
type TodoTaskAttachmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AttachmentBase collection
func (r *TodoTaskAttachmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AttachmentBase, error) {
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
	var values []AttachmentBase
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
			value  []AttachmentBase
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

// GetN performs GET request for AttachmentBase collection, max N pages
func (r *TodoTaskAttachmentsCollectionRequest) GetN(ctx context.Context, n int) ([]AttachmentBase, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AttachmentBase collection
func (r *TodoTaskAttachmentsCollectionRequest) Get(ctx context.Context) ([]AttachmentBase, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AttachmentBase collection
func (r *TodoTaskAttachmentsCollectionRequest) Add(ctx context.Context, reqObj *AttachmentBase) (resObj *AttachmentBase, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ChecklistItems returns request builder for ChecklistItem collection
func (b *TodoTaskRequestBuilder) ChecklistItems() *TodoTaskChecklistItemsCollectionRequestBuilder {
	bb := &TodoTaskChecklistItemsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/checklistItems"
	return bb
}

// TodoTaskChecklistItemsCollectionRequestBuilder is request builder for ChecklistItem collection
type TodoTaskChecklistItemsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ChecklistItem collection
func (b *TodoTaskChecklistItemsCollectionRequestBuilder) Request() *TodoTaskChecklistItemsCollectionRequest {
	return &TodoTaskChecklistItemsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ChecklistItem item
func (b *TodoTaskChecklistItemsCollectionRequestBuilder) ID(id string) *ChecklistItemRequestBuilder {
	bb := &ChecklistItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TodoTaskChecklistItemsCollectionRequest is request for ChecklistItem collection
type TodoTaskChecklistItemsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ChecklistItem collection
func (r *TodoTaskChecklistItemsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ChecklistItem, error) {
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
	var values []ChecklistItem
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
			value  []ChecklistItem
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

// GetN performs GET request for ChecklistItem collection, max N pages
func (r *TodoTaskChecklistItemsCollectionRequest) GetN(ctx context.Context, n int) ([]ChecklistItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ChecklistItem collection
func (r *TodoTaskChecklistItemsCollectionRequest) Get(ctx context.Context) ([]ChecklistItem, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ChecklistItem collection
func (r *TodoTaskChecklistItemsCollectionRequest) Add(ctx context.Context, reqObj *ChecklistItem) (resObj *ChecklistItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Extensions returns request builder for Extension collection
func (b *TodoTaskRequestBuilder) Extensions() *TodoTaskExtensionsCollectionRequestBuilder {
	bb := &TodoTaskExtensionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/extensions"
	return bb
}

// TodoTaskExtensionsCollectionRequestBuilder is request builder for Extension collection
type TodoTaskExtensionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Extension collection
func (b *TodoTaskExtensionsCollectionRequestBuilder) Request() *TodoTaskExtensionsCollectionRequest {
	return &TodoTaskExtensionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Extension item
func (b *TodoTaskExtensionsCollectionRequestBuilder) ID(id string) *ExtensionRequestBuilder {
	bb := &ExtensionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TodoTaskExtensionsCollectionRequest is request for Extension collection
type TodoTaskExtensionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Extension collection
func (r *TodoTaskExtensionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Extension, error) {
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
	var values []Extension
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
			value  []Extension
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

// GetN performs GET request for Extension collection, max N pages
func (r *TodoTaskExtensionsCollectionRequest) GetN(ctx context.Context, n int) ([]Extension, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Extension collection
func (r *TodoTaskExtensionsCollectionRequest) Get(ctx context.Context) ([]Extension, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Extension collection
func (r *TodoTaskExtensionsCollectionRequest) Add(ctx context.Context, reqObj *Extension) (resObj *Extension, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// LinkedResources returns request builder for LinkedResource collection
func (b *TodoTaskRequestBuilder) LinkedResources() *TodoTaskLinkedResourcesCollectionRequestBuilder {
	bb := &TodoTaskLinkedResourcesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/linkedResources"
	return bb
}

// TodoTaskLinkedResourcesCollectionRequestBuilder is request builder for LinkedResource collection
type TodoTaskLinkedResourcesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for LinkedResource collection
func (b *TodoTaskLinkedResourcesCollectionRequestBuilder) Request() *TodoTaskLinkedResourcesCollectionRequest {
	return &TodoTaskLinkedResourcesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for LinkedResource item
func (b *TodoTaskLinkedResourcesCollectionRequestBuilder) ID(id string) *LinkedResourceRequestBuilder {
	bb := &LinkedResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TodoTaskLinkedResourcesCollectionRequest is request for LinkedResource collection
type TodoTaskLinkedResourcesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for LinkedResource collection
func (r *TodoTaskLinkedResourcesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]LinkedResource, error) {
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
	var values []LinkedResource
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
			value  []LinkedResource
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

// GetN performs GET request for LinkedResource collection, max N pages
func (r *TodoTaskLinkedResourcesCollectionRequest) GetN(ctx context.Context, n int) ([]LinkedResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for LinkedResource collection
func (r *TodoTaskLinkedResourcesCollectionRequest) Get(ctx context.Context) ([]LinkedResource, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for LinkedResource collection
func (r *TodoTaskLinkedResourcesCollectionRequest) Add(ctx context.Context, reqObj *LinkedResource) (resObj *LinkedResource, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Extensions returns request builder for Extension collection
func (b *TodoTaskListRequestBuilder) Extensions() *TodoTaskListExtensionsCollectionRequestBuilder {
	bb := &TodoTaskListExtensionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/extensions"
	return bb
}

// TodoTaskListExtensionsCollectionRequestBuilder is request builder for Extension collection
type TodoTaskListExtensionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Extension collection
func (b *TodoTaskListExtensionsCollectionRequestBuilder) Request() *TodoTaskListExtensionsCollectionRequest {
	return &TodoTaskListExtensionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Extension item
func (b *TodoTaskListExtensionsCollectionRequestBuilder) ID(id string) *ExtensionRequestBuilder {
	bb := &ExtensionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TodoTaskListExtensionsCollectionRequest is request for Extension collection
type TodoTaskListExtensionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Extension collection
func (r *TodoTaskListExtensionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Extension, error) {
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
	var values []Extension
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
			value  []Extension
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

// GetN performs GET request for Extension collection, max N pages
func (r *TodoTaskListExtensionsCollectionRequest) GetN(ctx context.Context, n int) ([]Extension, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Extension collection
func (r *TodoTaskListExtensionsCollectionRequest) Get(ctx context.Context) ([]Extension, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Extension collection
func (r *TodoTaskListExtensionsCollectionRequest) Add(ctx context.Context, reqObj *Extension) (resObj *Extension, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Tasks returns request builder for TodoTask collection
func (b *TodoTaskListRequestBuilder) Tasks() *TodoTaskListTasksCollectionRequestBuilder {
	bb := &TodoTaskListTasksCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tasks"
	return bb
}

// TodoTaskListTasksCollectionRequestBuilder is request builder for TodoTask collection
type TodoTaskListTasksCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TodoTask collection
func (b *TodoTaskListTasksCollectionRequestBuilder) Request() *TodoTaskListTasksCollectionRequest {
	return &TodoTaskListTasksCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TodoTask item
func (b *TodoTaskListTasksCollectionRequestBuilder) ID(id string) *TodoTaskRequestBuilder {
	bb := &TodoTaskRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TodoTaskListTasksCollectionRequest is request for TodoTask collection
type TodoTaskListTasksCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TodoTask collection
func (r *TodoTaskListTasksCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TodoTask, error) {
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
	var values []TodoTask
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
			value  []TodoTask
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

// GetN performs GET request for TodoTask collection, max N pages
func (r *TodoTaskListTasksCollectionRequest) GetN(ctx context.Context, n int) ([]TodoTask, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TodoTask collection
func (r *TodoTaskListTasksCollectionRequest) Get(ctx context.Context) ([]TodoTask, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TodoTask collection
func (r *TodoTaskListTasksCollectionRequest) Add(ctx context.Context, reqObj *TodoTask) (resObj *TodoTask, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *TodoRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TodoTaskRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *TodoTaskListRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

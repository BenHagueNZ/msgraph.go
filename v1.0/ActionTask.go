// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// TaskFileAttachment returns request builder for TaskFileAttachment collection
func (b *TodoTaskAttachmentsCollectionRequestBuilder) TaskFileAttachment() *TodoTaskAttachmentsCollectionTaskFileAttachmentCollectionRequestBuilder {
	bb := &TodoTaskAttachmentsCollectionTaskFileAttachmentCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// TodoTaskAttachmentsCollectionTaskFileAttachmentCollectionRequestBuilder is request builder for TaskFileAttachment collection
type TodoTaskAttachmentsCollectionTaskFileAttachmentCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TaskFileAttachment collection
func (b *TodoTaskAttachmentsCollectionTaskFileAttachmentCollectionRequestBuilder) Request() *TodoTaskAttachmentsCollectionTaskFileAttachmentCollectionRequest {
	return &TodoTaskAttachmentsCollectionTaskFileAttachmentCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TaskFileAttachment item
func (b *TodoTaskAttachmentsCollectionTaskFileAttachmentCollectionRequestBuilder) ID(id string) *TaskFileAttachmentRequestBuilder {
	bb := &TaskFileAttachmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TodoTaskAttachmentsCollectionTaskFileAttachmentCollectionRequest is request for TaskFileAttachment collection
type TodoTaskAttachmentsCollectionTaskFileAttachmentCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TaskFileAttachment collection
func (r *TodoTaskAttachmentsCollectionTaskFileAttachmentCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TaskFileAttachment, error) {
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
	var values []TaskFileAttachment
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
			value  []TaskFileAttachment
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

// GetN performs GET request for TaskFileAttachment collection, max N pages
func (r *TodoTaskAttachmentsCollectionTaskFileAttachmentCollectionRequest) GetN(ctx context.Context, n int) ([]TaskFileAttachment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TaskFileAttachment collection
func (r *TodoTaskAttachmentsCollectionTaskFileAttachmentCollectionRequest) Get(ctx context.Context) ([]TaskFileAttachment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TaskFileAttachment collection
func (r *TodoTaskAttachmentsCollectionTaskFileAttachmentCollectionRequest) Add(ctx context.Context, reqObj *TaskFileAttachment) (resObj *TaskFileAttachment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// StandardWebPart returns request builder for StandardWebPart collection
func (b *HorizontalSectionColumnWebpartsCollectionRequestBuilder) StandardWebPart() *HorizontalSectionColumnWebpartsCollectionStandardWebPartCollectionRequestBuilder {
	bb := &HorizontalSectionColumnWebpartsCollectionStandardWebPartCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// HorizontalSectionColumnWebpartsCollectionStandardWebPartCollectionRequestBuilder is request builder for StandardWebPart collection rcn
type HorizontalSectionColumnWebpartsCollectionStandardWebPartCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for StandardWebPart collection
func (b *HorizontalSectionColumnWebpartsCollectionStandardWebPartCollectionRequestBuilder) Request() *HorizontalSectionColumnWebpartsCollectionStandardWebPartCollectionRequest {
	return &HorizontalSectionColumnWebpartsCollectionStandardWebPartCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for StandardWebPart item
func (b *HorizontalSectionColumnWebpartsCollectionStandardWebPartCollectionRequestBuilder) ID(id string) *StandardWebPartRequestBuilder {
	bb := &StandardWebPartRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// HorizontalSectionColumnWebpartsCollectionStandardWebPartCollectionRequest is request for StandardWebPart collection
type HorizontalSectionColumnWebpartsCollectionStandardWebPartCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for StandardWebPart collection
func (r *HorizontalSectionColumnWebpartsCollectionStandardWebPartCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]StandardWebPart, error) {
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
	var values []StandardWebPart
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
			value  []StandardWebPart
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

// GetN performs GET request for StandardWebPart collection, max N pages
func (r *HorizontalSectionColumnWebpartsCollectionStandardWebPartCollectionRequest) GetN(ctx context.Context, n int) ([]StandardWebPart, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for StandardWebPart collection
func (r *HorizontalSectionColumnWebpartsCollectionStandardWebPartCollectionRequest) Get(ctx context.Context) ([]StandardWebPart, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for StandardWebPart collection
func (r *HorizontalSectionColumnWebpartsCollectionStandardWebPartCollectionRequest) Add(ctx context.Context, reqObj *StandardWebPart) (resObj *StandardWebPart, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
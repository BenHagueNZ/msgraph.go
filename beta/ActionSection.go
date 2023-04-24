// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// ParentNotebook is navigation property rn
func (b *SectionGroupRequestBuilder) ParentNotebook() *NotebookRequestBuilder {
	bb := &NotebookRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parentNotebook"
	return bb
}

// ParentSectionGroup is navigation property rn
func (b *SectionGroupRequestBuilder) ParentSectionGroup() *SectionGroupRequestBuilder {
	bb := &SectionGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parentSectionGroup"
	return bb
}

// SectionGroups returns request builder for SectionGroup collection
func (b *SectionGroupRequestBuilder) SectionGroups() *SectionGroupSectionGroupsCollectionRequestBuilder {
	bb := &SectionGroupSectionGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sectionGroups"
	return bb
}

// SectionGroupSectionGroupsCollectionRequestBuilder is request builder for SectionGroup collection rcn
type SectionGroupSectionGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SectionGroup collection
func (b *SectionGroupSectionGroupsCollectionRequestBuilder) Request() *SectionGroupSectionGroupsCollectionRequest {
	return &SectionGroupSectionGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SectionGroup item
func (b *SectionGroupSectionGroupsCollectionRequestBuilder) ID(id string) *SectionGroupRequestBuilder {
	bb := &SectionGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SectionGroupSectionGroupsCollectionRequest is request for SectionGroup collection
type SectionGroupSectionGroupsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SectionGroup collection
func (r *SectionGroupSectionGroupsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SectionGroup, error) {
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
	var values []SectionGroup
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
			value  []SectionGroup
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

// GetN performs GET request for SectionGroup collection, max N pages
func (r *SectionGroupSectionGroupsCollectionRequest) GetN(ctx context.Context, n int) ([]SectionGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SectionGroup collection
func (r *SectionGroupSectionGroupsCollectionRequest) Get(ctx context.Context) ([]SectionGroup, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SectionGroup collection
func (r *SectionGroupSectionGroupsCollectionRequest) Add(ctx context.Context, reqObj *SectionGroup) (resObj *SectionGroup, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Sections returns request builder for OnenoteSection collection
func (b *SectionGroupRequestBuilder) Sections() *SectionGroupSectionsCollectionRequestBuilder {
	bb := &SectionGroupSectionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sections"
	return bb
}

// SectionGroupSectionsCollectionRequestBuilder is request builder for OnenoteSection collection rcn
type SectionGroupSectionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnenoteSection collection
func (b *SectionGroupSectionsCollectionRequestBuilder) Request() *SectionGroupSectionsCollectionRequest {
	return &SectionGroupSectionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnenoteSection item
func (b *SectionGroupSectionsCollectionRequestBuilder) ID(id string) *OnenoteSectionRequestBuilder {
	bb := &OnenoteSectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SectionGroupSectionsCollectionRequest is request for OnenoteSection collection
type SectionGroupSectionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnenoteSection collection
func (r *SectionGroupSectionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]OnenoteSection, error) {
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
	var values []OnenoteSection
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
			value  []OnenoteSection
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

// GetN performs GET request for OnenoteSection collection, max N pages
func (r *SectionGroupSectionsCollectionRequest) GetN(ctx context.Context, n int) ([]OnenoteSection, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for OnenoteSection collection
func (r *SectionGroupSectionsCollectionRequest) Get(ctx context.Context) ([]OnenoteSection, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for OnenoteSection collection
func (r *SectionGroupSectionsCollectionRequest) Add(ctx context.Context, reqObj *OnenoteSection) (resObj *OnenoteSection, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
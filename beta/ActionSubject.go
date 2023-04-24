// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// Approvers returns request builder for User collection
func (b *SubjectRightsRequestObjectRequestBuilder) Approvers() *SubjectRightsRequestObjectApproversCollectionRequestBuilder {
	bb := &SubjectRightsRequestObjectApproversCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/approvers"
	return bb
}

// SubjectRightsRequestObjectApproversCollectionRequestBuilder is request builder for User collection rcn
type SubjectRightsRequestObjectApproversCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for User collection
func (b *SubjectRightsRequestObjectApproversCollectionRequestBuilder) Request() *SubjectRightsRequestObjectApproversCollectionRequest {
	return &SubjectRightsRequestObjectApproversCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for User item
func (b *SubjectRightsRequestObjectApproversCollectionRequestBuilder) ID(id string) *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SubjectRightsRequestObjectApproversCollectionRequest is request for User collection
type SubjectRightsRequestObjectApproversCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for User collection
func (r *SubjectRightsRequestObjectApproversCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]User, error) {
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
func (r *SubjectRightsRequestObjectApproversCollectionRequest) GetN(ctx context.Context, n int) ([]User, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for User collection
func (r *SubjectRightsRequestObjectApproversCollectionRequest) Get(ctx context.Context) ([]User, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for User collection
func (r *SubjectRightsRequestObjectApproversCollectionRequest) Add(ctx context.Context, reqObj *User) (resObj *User, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Collaborators returns request builder for User collection
func (b *SubjectRightsRequestObjectRequestBuilder) Collaborators() *SubjectRightsRequestObjectCollaboratorsCollectionRequestBuilder {
	bb := &SubjectRightsRequestObjectCollaboratorsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/collaborators"
	return bb
}

// SubjectRightsRequestObjectCollaboratorsCollectionRequestBuilder is request builder for User collection rcn
type SubjectRightsRequestObjectCollaboratorsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for User collection
func (b *SubjectRightsRequestObjectCollaboratorsCollectionRequestBuilder) Request() *SubjectRightsRequestObjectCollaboratorsCollectionRequest {
	return &SubjectRightsRequestObjectCollaboratorsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for User item
func (b *SubjectRightsRequestObjectCollaboratorsCollectionRequestBuilder) ID(id string) *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SubjectRightsRequestObjectCollaboratorsCollectionRequest is request for User collection
type SubjectRightsRequestObjectCollaboratorsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for User collection
func (r *SubjectRightsRequestObjectCollaboratorsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]User, error) {
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
func (r *SubjectRightsRequestObjectCollaboratorsCollectionRequest) GetN(ctx context.Context, n int) ([]User, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for User collection
func (r *SubjectRightsRequestObjectCollaboratorsCollectionRequest) Get(ctx context.Context) ([]User, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for User collection
func (r *SubjectRightsRequestObjectCollaboratorsCollectionRequest) Add(ctx context.Context, reqObj *User) (resObj *User, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Notes returns request builder for AuthoredNote collection
func (b *SubjectRightsRequestObjectRequestBuilder) Notes() *SubjectRightsRequestObjectNotesCollectionRequestBuilder {
	bb := &SubjectRightsRequestObjectNotesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/notes"
	return bb
}

// SubjectRightsRequestObjectNotesCollectionRequestBuilder is request builder for AuthoredNote collection rcn
type SubjectRightsRequestObjectNotesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AuthoredNote collection
func (b *SubjectRightsRequestObjectNotesCollectionRequestBuilder) Request() *SubjectRightsRequestObjectNotesCollectionRequest {
	return &SubjectRightsRequestObjectNotesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AuthoredNote item
func (b *SubjectRightsRequestObjectNotesCollectionRequestBuilder) ID(id string) *AuthoredNoteRequestBuilder {
	bb := &AuthoredNoteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SubjectRightsRequestObjectNotesCollectionRequest is request for AuthoredNote collection
type SubjectRightsRequestObjectNotesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AuthoredNote collection
func (r *SubjectRightsRequestObjectNotesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AuthoredNote, error) {
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
	var values []AuthoredNote
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
			value  []AuthoredNote
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

// GetN performs GET request for AuthoredNote collection, max N pages
func (r *SubjectRightsRequestObjectNotesCollectionRequest) GetN(ctx context.Context, n int) ([]AuthoredNote, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AuthoredNote collection
func (r *SubjectRightsRequestObjectNotesCollectionRequest) Get(ctx context.Context) ([]AuthoredNote, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AuthoredNote collection
func (r *SubjectRightsRequestObjectNotesCollectionRequest) Add(ctx context.Context, reqObj *AuthoredNote) (resObj *AuthoredNote, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Team is navigation property rn
func (b *SubjectRightsRequestObjectRequestBuilder) Team() *TeamRequestBuilder {
	bb := &TeamRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/team"
	return bb
}

// Entity is navigation property rn
func (b *SubjectRightsRequestObjectRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
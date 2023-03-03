// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// WindowsInformationProtectionAssignRequestParameter undocumented
type WindowsInformationProtectionAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []TargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}

// Assignments returns request builder for TargetedManagedAppPolicyAssignment collection
func (b *WindowsInformationProtectionRequestBuilder) Assignments() *WindowsInformationProtectionAssignmentsCollectionRequestBuilder {
	bb := &WindowsInformationProtectionAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// WindowsInformationProtectionAssignmentsCollectionRequestBuilder is request builder for TargetedManagedAppPolicyAssignment collection
type WindowsInformationProtectionAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TargetedManagedAppPolicyAssignment collection
func (b *WindowsInformationProtectionAssignmentsCollectionRequestBuilder) Request() *WindowsInformationProtectionAssignmentsCollectionRequest {
	return &WindowsInformationProtectionAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TargetedManagedAppPolicyAssignment item
func (b *WindowsInformationProtectionAssignmentsCollectionRequestBuilder) ID(id string) *TargetedManagedAppPolicyAssignmentRequestBuilder {
	bb := &TargetedManagedAppPolicyAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsInformationProtectionAssignmentsCollectionRequest is request for TargetedManagedAppPolicyAssignment collection
type WindowsInformationProtectionAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TargetedManagedAppPolicyAssignment collection
func (r *WindowsInformationProtectionAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TargetedManagedAppPolicyAssignment, error) {
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
	var values []TargetedManagedAppPolicyAssignment
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
			value  []TargetedManagedAppPolicyAssignment
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

// GetN performs GET request for TargetedManagedAppPolicyAssignment collection, max N pages
func (r *WindowsInformationProtectionAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]TargetedManagedAppPolicyAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TargetedManagedAppPolicyAssignment collection
func (r *WindowsInformationProtectionAssignmentsCollectionRequest) Get(ctx context.Context) ([]TargetedManagedAppPolicyAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TargetedManagedAppPolicyAssignment collection
func (r *WindowsInformationProtectionAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *TargetedManagedAppPolicyAssignment) (resObj *TargetedManagedAppPolicyAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ExemptAppLockerFiles returns request builder for WindowsInformationProtectionAppLockerFile collection
func (b *WindowsInformationProtectionRequestBuilder) ExemptAppLockerFiles() *WindowsInformationProtectionExemptAppLockerFilesCollectionRequestBuilder {
	bb := &WindowsInformationProtectionExemptAppLockerFilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/exemptAppLockerFiles"
	return bb
}

// WindowsInformationProtectionExemptAppLockerFilesCollectionRequestBuilder is request builder for WindowsInformationProtectionAppLockerFile collection
type WindowsInformationProtectionExemptAppLockerFilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WindowsInformationProtectionAppLockerFile collection
func (b *WindowsInformationProtectionExemptAppLockerFilesCollectionRequestBuilder) Request() *WindowsInformationProtectionExemptAppLockerFilesCollectionRequest {
	return &WindowsInformationProtectionExemptAppLockerFilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WindowsInformationProtectionAppLockerFile item
func (b *WindowsInformationProtectionExemptAppLockerFilesCollectionRequestBuilder) ID(id string) *WindowsInformationProtectionAppLockerFileRequestBuilder {
	bb := &WindowsInformationProtectionAppLockerFileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsInformationProtectionExemptAppLockerFilesCollectionRequest is request for WindowsInformationProtectionAppLockerFile collection
type WindowsInformationProtectionExemptAppLockerFilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WindowsInformationProtectionAppLockerFile collection
func (r *WindowsInformationProtectionExemptAppLockerFilesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]WindowsInformationProtectionAppLockerFile, error) {
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
	var values []WindowsInformationProtectionAppLockerFile
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
			value  []WindowsInformationProtectionAppLockerFile
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

// GetN performs GET request for WindowsInformationProtectionAppLockerFile collection, max N pages
func (r *WindowsInformationProtectionExemptAppLockerFilesCollectionRequest) GetN(ctx context.Context, n int) ([]WindowsInformationProtectionAppLockerFile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for WindowsInformationProtectionAppLockerFile collection
func (r *WindowsInformationProtectionExemptAppLockerFilesCollectionRequest) Get(ctx context.Context) ([]WindowsInformationProtectionAppLockerFile, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for WindowsInformationProtectionAppLockerFile collection
func (r *WindowsInformationProtectionExemptAppLockerFilesCollectionRequest) Add(ctx context.Context, reqObj *WindowsInformationProtectionAppLockerFile) (resObj *WindowsInformationProtectionAppLockerFile, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ProtectedAppLockerFiles returns request builder for WindowsInformationProtectionAppLockerFile collection
func (b *WindowsInformationProtectionRequestBuilder) ProtectedAppLockerFiles() *WindowsInformationProtectionProtectedAppLockerFilesCollectionRequestBuilder {
	bb := &WindowsInformationProtectionProtectedAppLockerFilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/protectedAppLockerFiles"
	return bb
}

// WindowsInformationProtectionProtectedAppLockerFilesCollectionRequestBuilder is request builder for WindowsInformationProtectionAppLockerFile collection
type WindowsInformationProtectionProtectedAppLockerFilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WindowsInformationProtectionAppLockerFile collection
func (b *WindowsInformationProtectionProtectedAppLockerFilesCollectionRequestBuilder) Request() *WindowsInformationProtectionProtectedAppLockerFilesCollectionRequest {
	return &WindowsInformationProtectionProtectedAppLockerFilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WindowsInformationProtectionAppLockerFile item
func (b *WindowsInformationProtectionProtectedAppLockerFilesCollectionRequestBuilder) ID(id string) *WindowsInformationProtectionAppLockerFileRequestBuilder {
	bb := &WindowsInformationProtectionAppLockerFileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsInformationProtectionProtectedAppLockerFilesCollectionRequest is request for WindowsInformationProtectionAppLockerFile collection
type WindowsInformationProtectionProtectedAppLockerFilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WindowsInformationProtectionAppLockerFile collection
func (r *WindowsInformationProtectionProtectedAppLockerFilesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]WindowsInformationProtectionAppLockerFile, error) {
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
	var values []WindowsInformationProtectionAppLockerFile
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
			value  []WindowsInformationProtectionAppLockerFile
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

// GetN performs GET request for WindowsInformationProtectionAppLockerFile collection, max N pages
func (r *WindowsInformationProtectionProtectedAppLockerFilesCollectionRequest) GetN(ctx context.Context, n int) ([]WindowsInformationProtectionAppLockerFile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for WindowsInformationProtectionAppLockerFile collection
func (r *WindowsInformationProtectionProtectedAppLockerFilesCollectionRequest) Get(ctx context.Context) ([]WindowsInformationProtectionAppLockerFile, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for WindowsInformationProtectionAppLockerFile collection
func (r *WindowsInformationProtectionProtectedAppLockerFilesCollectionRequest) Add(ctx context.Context, reqObj *WindowsInformationProtectionAppLockerFile) (resObj *WindowsInformationProtectionAppLockerFile, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

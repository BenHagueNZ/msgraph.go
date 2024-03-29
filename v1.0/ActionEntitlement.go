// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// AccessPackageAssignmentApprovals returns request builder for Approval collection
func (b *EntitlementManagementRequestBuilder) AccessPackageAssignmentApprovals() *EntitlementManagementAccessPackageAssignmentApprovalsCollectionRequestBuilder {
	bb := &EntitlementManagementAccessPackageAssignmentApprovalsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageAssignmentApprovals"
	return bb
}

// EntitlementManagementAccessPackageAssignmentApprovalsCollectionRequestBuilder is request builder for Approval collection rcn
type EntitlementManagementAccessPackageAssignmentApprovalsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Approval collection
func (b *EntitlementManagementAccessPackageAssignmentApprovalsCollectionRequestBuilder) Request() *EntitlementManagementAccessPackageAssignmentApprovalsCollectionRequest {
	return &EntitlementManagementAccessPackageAssignmentApprovalsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Approval item
func (b *EntitlementManagementAccessPackageAssignmentApprovalsCollectionRequestBuilder) ID(id string) *ApprovalRequestBuilder {
	bb := &ApprovalRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementAccessPackageAssignmentApprovalsCollectionRequest is request for Approval collection
type EntitlementManagementAccessPackageAssignmentApprovalsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Approval collection
func (r *EntitlementManagementAccessPackageAssignmentApprovalsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Approval, error) {
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
	var values []Approval
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
			value  []Approval
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

// GetN performs GET request for Approval collection, max N pages
func (r *EntitlementManagementAccessPackageAssignmentApprovalsCollectionRequest) GetN(ctx context.Context, n int) ([]Approval, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Approval collection
func (r *EntitlementManagementAccessPackageAssignmentApprovalsCollectionRequest) Get(ctx context.Context) ([]Approval, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Approval collection
func (r *EntitlementManagementAccessPackageAssignmentApprovalsCollectionRequest) Add(ctx context.Context, reqObj *Approval) (resObj *Approval, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AccessPackages returns request builder for AccessPackage collection
func (b *EntitlementManagementRequestBuilder) AccessPackages() *EntitlementManagementAccessPackagesCollectionRequestBuilder {
	bb := &EntitlementManagementAccessPackagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackages"
	return bb
}

// EntitlementManagementAccessPackagesCollectionRequestBuilder is request builder for AccessPackage collection rcn
type EntitlementManagementAccessPackagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackage collection
func (b *EntitlementManagementAccessPackagesCollectionRequestBuilder) Request() *EntitlementManagementAccessPackagesCollectionRequest {
	return &EntitlementManagementAccessPackagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackage item
func (b *EntitlementManagementAccessPackagesCollectionRequestBuilder) ID(id string) *AccessPackageRequestBuilder {
	bb := &AccessPackageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementAccessPackagesCollectionRequest is request for AccessPackage collection
type EntitlementManagementAccessPackagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackage collection
func (r *EntitlementManagementAccessPackagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AccessPackage, error) {
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
	var values []AccessPackage
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
			value  []AccessPackage
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

// GetN performs GET request for AccessPackage collection, max N pages
func (r *EntitlementManagementAccessPackagesCollectionRequest) GetN(ctx context.Context, n int) ([]AccessPackage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AccessPackage collection
func (r *EntitlementManagementAccessPackagesCollectionRequest) Get(ctx context.Context) ([]AccessPackage, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AccessPackage collection
func (r *EntitlementManagementAccessPackagesCollectionRequest) Add(ctx context.Context, reqObj *AccessPackage) (resObj *AccessPackage, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AssignmentPolicies returns request builder for AccessPackageAssignmentPolicy collection
func (b *EntitlementManagementRequestBuilder) AssignmentPolicies() *EntitlementManagementAssignmentPoliciesCollectionRequestBuilder {
	bb := &EntitlementManagementAssignmentPoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignmentPolicies"
	return bb
}

// EntitlementManagementAssignmentPoliciesCollectionRequestBuilder is request builder for AccessPackageAssignmentPolicy collection rcn
type EntitlementManagementAssignmentPoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageAssignmentPolicy collection
func (b *EntitlementManagementAssignmentPoliciesCollectionRequestBuilder) Request() *EntitlementManagementAssignmentPoliciesCollectionRequest {
	return &EntitlementManagementAssignmentPoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageAssignmentPolicy item
func (b *EntitlementManagementAssignmentPoliciesCollectionRequestBuilder) ID(id string) *AccessPackageAssignmentPolicyRequestBuilder {
	bb := &AccessPackageAssignmentPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementAssignmentPoliciesCollectionRequest is request for AccessPackageAssignmentPolicy collection
type EntitlementManagementAssignmentPoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageAssignmentPolicy collection
func (r *EntitlementManagementAssignmentPoliciesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AccessPackageAssignmentPolicy, error) {
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
	var values []AccessPackageAssignmentPolicy
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
			value  []AccessPackageAssignmentPolicy
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

// GetN performs GET request for AccessPackageAssignmentPolicy collection, max N pages
func (r *EntitlementManagementAssignmentPoliciesCollectionRequest) GetN(ctx context.Context, n int) ([]AccessPackageAssignmentPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AccessPackageAssignmentPolicy collection
func (r *EntitlementManagementAssignmentPoliciesCollectionRequest) Get(ctx context.Context) ([]AccessPackageAssignmentPolicy, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AccessPackageAssignmentPolicy collection
func (r *EntitlementManagementAssignmentPoliciesCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageAssignmentPolicy) (resObj *AccessPackageAssignmentPolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AssignmentRequests returns request builder for AccessPackageAssignmentRequestObject collection
func (b *EntitlementManagementRequestBuilder) AssignmentRequests() *EntitlementManagementAssignmentRequestsCollectionRequestBuilder {
	bb := &EntitlementManagementAssignmentRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignmentRequests"
	return bb
}

// EntitlementManagementAssignmentRequestsCollectionRequestBuilder is request builder for AccessPackageAssignmentRequestObject collection rcn
type EntitlementManagementAssignmentRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageAssignmentRequestObject collection
func (b *EntitlementManagementAssignmentRequestsCollectionRequestBuilder) Request() *EntitlementManagementAssignmentRequestsCollectionRequest {
	return &EntitlementManagementAssignmentRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageAssignmentRequestObject item
func (b *EntitlementManagementAssignmentRequestsCollectionRequestBuilder) ID(id string) *AccessPackageAssignmentRequestObjectRequestBuilder {
	bb := &AccessPackageAssignmentRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementAssignmentRequestsCollectionRequest is request for AccessPackageAssignmentRequestObject collection
type EntitlementManagementAssignmentRequestsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageAssignmentRequestObject collection
func (r *EntitlementManagementAssignmentRequestsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AccessPackageAssignmentRequestObject, error) {
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
	var values []AccessPackageAssignmentRequestObject
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
			value  []AccessPackageAssignmentRequestObject
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

// GetN performs GET request for AccessPackageAssignmentRequestObject collection, max N pages
func (r *EntitlementManagementAssignmentRequestsCollectionRequest) GetN(ctx context.Context, n int) ([]AccessPackageAssignmentRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AccessPackageAssignmentRequestObject collection
func (r *EntitlementManagementAssignmentRequestsCollectionRequest) Get(ctx context.Context) ([]AccessPackageAssignmentRequestObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AccessPackageAssignmentRequestObject collection
func (r *EntitlementManagementAssignmentRequestsCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageAssignmentRequestObject) (resObj *AccessPackageAssignmentRequestObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Assignments returns request builder for AccessPackageAssignment collection
func (b *EntitlementManagementRequestBuilder) Assignments() *EntitlementManagementAssignmentsCollectionRequestBuilder {
	bb := &EntitlementManagementAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// EntitlementManagementAssignmentsCollectionRequestBuilder is request builder for AccessPackageAssignment collection rcn
type EntitlementManagementAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageAssignment collection
func (b *EntitlementManagementAssignmentsCollectionRequestBuilder) Request() *EntitlementManagementAssignmentsCollectionRequest {
	return &EntitlementManagementAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageAssignment item
func (b *EntitlementManagementAssignmentsCollectionRequestBuilder) ID(id string) *AccessPackageAssignmentRequestBuilder {
	bb := &AccessPackageAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementAssignmentsCollectionRequest is request for AccessPackageAssignment collection
type EntitlementManagementAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageAssignment collection
func (r *EntitlementManagementAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AccessPackageAssignment, error) {
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
	var values []AccessPackageAssignment
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
			value  []AccessPackageAssignment
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

// GetN performs GET request for AccessPackageAssignment collection, max N pages
func (r *EntitlementManagementAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]AccessPackageAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AccessPackageAssignment collection
func (r *EntitlementManagementAssignmentsCollectionRequest) Get(ctx context.Context) ([]AccessPackageAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AccessPackageAssignment collection
func (r *EntitlementManagementAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageAssignment) (resObj *AccessPackageAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Catalogs returns request builder for AccessPackageCatalog collection
func (b *EntitlementManagementRequestBuilder) Catalogs() *EntitlementManagementCatalogsCollectionRequestBuilder {
	bb := &EntitlementManagementCatalogsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/catalogs"
	return bb
}

// EntitlementManagementCatalogsCollectionRequestBuilder is request builder for AccessPackageCatalog collection rcn
type EntitlementManagementCatalogsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageCatalog collection
func (b *EntitlementManagementCatalogsCollectionRequestBuilder) Request() *EntitlementManagementCatalogsCollectionRequest {
	return &EntitlementManagementCatalogsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageCatalog item
func (b *EntitlementManagementCatalogsCollectionRequestBuilder) ID(id string) *AccessPackageCatalogRequestBuilder {
	bb := &AccessPackageCatalogRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementCatalogsCollectionRequest is request for AccessPackageCatalog collection
type EntitlementManagementCatalogsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageCatalog collection
func (r *EntitlementManagementCatalogsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AccessPackageCatalog, error) {
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
	var values []AccessPackageCatalog
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
			value  []AccessPackageCatalog
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

// GetN performs GET request for AccessPackageCatalog collection, max N pages
func (r *EntitlementManagementCatalogsCollectionRequest) GetN(ctx context.Context, n int) ([]AccessPackageCatalog, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AccessPackageCatalog collection
func (r *EntitlementManagementCatalogsCollectionRequest) Get(ctx context.Context) ([]AccessPackageCatalog, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AccessPackageCatalog collection
func (r *EntitlementManagementCatalogsCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageCatalog) (resObj *AccessPackageCatalog, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ConnectedOrganizations returns request builder for ConnectedOrganization collection
func (b *EntitlementManagementRequestBuilder) ConnectedOrganizations() *EntitlementManagementConnectedOrganizationsCollectionRequestBuilder {
	bb := &EntitlementManagementConnectedOrganizationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/connectedOrganizations"
	return bb
}

// EntitlementManagementConnectedOrganizationsCollectionRequestBuilder is request builder for ConnectedOrganization collection rcn
type EntitlementManagementConnectedOrganizationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ConnectedOrganization collection
func (b *EntitlementManagementConnectedOrganizationsCollectionRequestBuilder) Request() *EntitlementManagementConnectedOrganizationsCollectionRequest {
	return &EntitlementManagementConnectedOrganizationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ConnectedOrganization item
func (b *EntitlementManagementConnectedOrganizationsCollectionRequestBuilder) ID(id string) *ConnectedOrganizationRequestBuilder {
	bb := &ConnectedOrganizationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementConnectedOrganizationsCollectionRequest is request for ConnectedOrganization collection
type EntitlementManagementConnectedOrganizationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ConnectedOrganization collection
func (r *EntitlementManagementConnectedOrganizationsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ConnectedOrganization, error) {
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
	var values []ConnectedOrganization
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
			value  []ConnectedOrganization
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

// GetN performs GET request for ConnectedOrganization collection, max N pages
func (r *EntitlementManagementConnectedOrganizationsCollectionRequest) GetN(ctx context.Context, n int) ([]ConnectedOrganization, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ConnectedOrganization collection
func (r *EntitlementManagementConnectedOrganizationsCollectionRequest) Get(ctx context.Context) ([]ConnectedOrganization, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ConnectedOrganization collection
func (r *EntitlementManagementConnectedOrganizationsCollectionRequest) Add(ctx context.Context, reqObj *ConnectedOrganization) (resObj *ConnectedOrganization, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Settings is navigation property rn
func (b *EntitlementManagementRequestBuilder) Settings() *EntitlementManagementSettingsRequestBuilder {
	bb := &EntitlementManagementSettingsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/settings"
	return bb
}

// Entity is navigation property rn
func (b *EntitlementManagementRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *EntitlementManagementSettingsRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// ApplicationSetVerifiedPublisherRequestParameter undocumented
type ApplicationSetVerifiedPublisherRequestParameter struct {
	// VerifiedPublisherID undocumented
	VerifiedPublisherID *string `json:"verifiedPublisherId,omitempty"`
}

// ApplicationUnsetVerifiedPublisherRequestParameter undocumented
type ApplicationUnsetVerifiedPublisherRequestParameter struct {
}

// ApplicationAddKeyRequestParameter undocumented
type ApplicationAddKeyRequestParameter struct {
	// KeyCredential undocumented
	KeyCredential *KeyCredential `json:"keyCredential,omitempty"`
	// PasswordCredential undocumented
	PasswordCredential *PasswordCredential `json:"passwordCredential,omitempty"`
	// Proof undocumented
	Proof *string `json:"proof,omitempty"`
}

// ApplicationAddPasswordRequestParameter undocumented
type ApplicationAddPasswordRequestParameter struct {
	// PasswordCredential undocumented
	PasswordCredential *PasswordCredential `json:"passwordCredential,omitempty"`
}

// ApplicationRemoveKeyRequestParameter undocumented
type ApplicationRemoveKeyRequestParameter struct {
	// KeyID undocumented
	KeyID *UUID `json:"keyId,omitempty"`
	// Proof undocumented
	Proof *string `json:"proof,omitempty"`
}

// ApplicationRemovePasswordRequestParameter undocumented
type ApplicationRemovePasswordRequestParameter struct {
	// KeyID undocumented
	KeyID *UUID `json:"keyId,omitempty"`
}

// ApplicationTemplateInstantiateRequestParameter undocumented
type ApplicationTemplateInstantiateRequestParameter struct {
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

// AppManagementPolicies returns request builder for AppManagementPolicy collection
func (b *ApplicationRequestBuilder) AppManagementPolicies() *ApplicationAppManagementPoliciesCollectionRequestBuilder {
	bb := &ApplicationAppManagementPoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/appManagementPolicies"
	return bb
}

// ApplicationAppManagementPoliciesCollectionRequestBuilder is request builder for AppManagementPolicy collection rcn
type ApplicationAppManagementPoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AppManagementPolicy collection
func (b *ApplicationAppManagementPoliciesCollectionRequestBuilder) Request() *ApplicationAppManagementPoliciesCollectionRequest {
	return &ApplicationAppManagementPoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AppManagementPolicy item
func (b *ApplicationAppManagementPoliciesCollectionRequestBuilder) ID(id string) *AppManagementPolicyRequestBuilder {
	bb := &AppManagementPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApplicationAppManagementPoliciesCollectionRequest is request for AppManagementPolicy collection
type ApplicationAppManagementPoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AppManagementPolicy collection
func (r *ApplicationAppManagementPoliciesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AppManagementPolicy, error) {
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
	var values []AppManagementPolicy
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
			value  []AppManagementPolicy
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

// GetN performs GET request for AppManagementPolicy collection, max N pages
func (r *ApplicationAppManagementPoliciesCollectionRequest) GetN(ctx context.Context, n int) ([]AppManagementPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AppManagementPolicy collection
func (r *ApplicationAppManagementPoliciesCollectionRequest) Get(ctx context.Context) ([]AppManagementPolicy, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AppManagementPolicy collection
func (r *ApplicationAppManagementPoliciesCollectionRequest) Add(ctx context.Context, reqObj *AppManagementPolicy) (resObj *AppManagementPolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ConnectorGroup is navigation property rn
func (b *ApplicationRequestBuilder) ConnectorGroup() *ConnectorGroupRequestBuilder {
	bb := &ConnectorGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/connectorGroup"
	return bb
}

// CreatedOnBehalfOf is navigation property rn
func (b *ApplicationRequestBuilder) CreatedOnBehalfOf() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/createdOnBehalfOf"
	return bb
}

// ExtensionProperties returns request builder for ExtensionProperty collection
func (b *ApplicationRequestBuilder) ExtensionProperties() *ApplicationExtensionPropertiesCollectionRequestBuilder {
	bb := &ApplicationExtensionPropertiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/extensionProperties"
	return bb
}

// ApplicationExtensionPropertiesCollectionRequestBuilder is request builder for ExtensionProperty collection rcn
type ApplicationExtensionPropertiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ExtensionProperty collection
func (b *ApplicationExtensionPropertiesCollectionRequestBuilder) Request() *ApplicationExtensionPropertiesCollectionRequest {
	return &ApplicationExtensionPropertiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ExtensionProperty item
func (b *ApplicationExtensionPropertiesCollectionRequestBuilder) ID(id string) *ExtensionPropertyRequestBuilder {
	bb := &ExtensionPropertyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApplicationExtensionPropertiesCollectionRequest is request for ExtensionProperty collection
type ApplicationExtensionPropertiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ExtensionProperty collection
func (r *ApplicationExtensionPropertiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ExtensionProperty, error) {
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
	var values []ExtensionProperty
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
			value  []ExtensionProperty
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

// GetN performs GET request for ExtensionProperty collection, max N pages
func (r *ApplicationExtensionPropertiesCollectionRequest) GetN(ctx context.Context, n int) ([]ExtensionProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ExtensionProperty collection
func (r *ApplicationExtensionPropertiesCollectionRequest) Get(ctx context.Context) ([]ExtensionProperty, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ExtensionProperty collection
func (r *ApplicationExtensionPropertiesCollectionRequest) Add(ctx context.Context, reqObj *ExtensionProperty) (resObj *ExtensionProperty, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// FederatedIdentityCredentials returns request builder for FederatedIdentityCredential collection
func (b *ApplicationRequestBuilder) FederatedIdentityCredentials() *ApplicationFederatedIdentityCredentialsCollectionRequestBuilder {
	bb := &ApplicationFederatedIdentityCredentialsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/federatedIdentityCredentials"
	return bb
}

// ApplicationFederatedIdentityCredentialsCollectionRequestBuilder is request builder for FederatedIdentityCredential collection rcn
type ApplicationFederatedIdentityCredentialsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for FederatedIdentityCredential collection
func (b *ApplicationFederatedIdentityCredentialsCollectionRequestBuilder) Request() *ApplicationFederatedIdentityCredentialsCollectionRequest {
	return &ApplicationFederatedIdentityCredentialsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for FederatedIdentityCredential item
func (b *ApplicationFederatedIdentityCredentialsCollectionRequestBuilder) ID(id string) *FederatedIdentityCredentialRequestBuilder {
	bb := &FederatedIdentityCredentialRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApplicationFederatedIdentityCredentialsCollectionRequest is request for FederatedIdentityCredential collection
type ApplicationFederatedIdentityCredentialsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for FederatedIdentityCredential collection
func (r *ApplicationFederatedIdentityCredentialsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]FederatedIdentityCredential, error) {
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
	var values []FederatedIdentityCredential
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
			value  []FederatedIdentityCredential
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

// GetN performs GET request for FederatedIdentityCredential collection, max N pages
func (r *ApplicationFederatedIdentityCredentialsCollectionRequest) GetN(ctx context.Context, n int) ([]FederatedIdentityCredential, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for FederatedIdentityCredential collection
func (r *ApplicationFederatedIdentityCredentialsCollectionRequest) Get(ctx context.Context) ([]FederatedIdentityCredential, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for FederatedIdentityCredential collection
func (r *ApplicationFederatedIdentityCredentialsCollectionRequest) Add(ctx context.Context, reqObj *FederatedIdentityCredential) (resObj *FederatedIdentityCredential, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// HomeRealmDiscoveryPolicies returns request builder for HomeRealmDiscoveryPolicy collection
func (b *ApplicationRequestBuilder) HomeRealmDiscoveryPolicies() *ApplicationHomeRealmDiscoveryPoliciesCollectionRequestBuilder {
	bb := &ApplicationHomeRealmDiscoveryPoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/homeRealmDiscoveryPolicies"
	return bb
}

// ApplicationHomeRealmDiscoveryPoliciesCollectionRequestBuilder is request builder for HomeRealmDiscoveryPolicy collection rcn
type ApplicationHomeRealmDiscoveryPoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for HomeRealmDiscoveryPolicy collection
func (b *ApplicationHomeRealmDiscoveryPoliciesCollectionRequestBuilder) Request() *ApplicationHomeRealmDiscoveryPoliciesCollectionRequest {
	return &ApplicationHomeRealmDiscoveryPoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for HomeRealmDiscoveryPolicy item
func (b *ApplicationHomeRealmDiscoveryPoliciesCollectionRequestBuilder) ID(id string) *HomeRealmDiscoveryPolicyRequestBuilder {
	bb := &HomeRealmDiscoveryPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApplicationHomeRealmDiscoveryPoliciesCollectionRequest is request for HomeRealmDiscoveryPolicy collection
type ApplicationHomeRealmDiscoveryPoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for HomeRealmDiscoveryPolicy collection
func (r *ApplicationHomeRealmDiscoveryPoliciesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]HomeRealmDiscoveryPolicy, error) {
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
	var values []HomeRealmDiscoveryPolicy
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
			value  []HomeRealmDiscoveryPolicy
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

// GetN performs GET request for HomeRealmDiscoveryPolicy collection, max N pages
func (r *ApplicationHomeRealmDiscoveryPoliciesCollectionRequest) GetN(ctx context.Context, n int) ([]HomeRealmDiscoveryPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for HomeRealmDiscoveryPolicy collection
func (r *ApplicationHomeRealmDiscoveryPoliciesCollectionRequest) Get(ctx context.Context) ([]HomeRealmDiscoveryPolicy, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for HomeRealmDiscoveryPolicy collection
func (r *ApplicationHomeRealmDiscoveryPoliciesCollectionRequest) Add(ctx context.Context, reqObj *HomeRealmDiscoveryPolicy) (resObj *HomeRealmDiscoveryPolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Owners returns request builder for DirectoryObject collection
func (b *ApplicationRequestBuilder) Owners() *ApplicationOwnersCollectionRequestBuilder {
	bb := &ApplicationOwnersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/owners"
	return bb
}

// ApplicationOwnersCollectionRequestBuilder is request builder for DirectoryObject collection rcn
type ApplicationOwnersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ApplicationOwnersCollectionRequestBuilder) Request() *ApplicationOwnersCollectionRequest {
	return &ApplicationOwnersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ApplicationOwnersCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApplicationOwnersCollectionRequest is request for DirectoryObject collection
type ApplicationOwnersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ApplicationOwnersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *ApplicationOwnersCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *ApplicationOwnersCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *ApplicationOwnersCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Synchronization is navigation property rn
func (b *ApplicationRequestBuilder) Synchronization() *SynchronizationRequestBuilder {
	bb := &SynchronizationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/synchronization"
	return bb
}

// TokenIssuancePolicies returns request builder for TokenIssuancePolicy collection
func (b *ApplicationRequestBuilder) TokenIssuancePolicies() *ApplicationTokenIssuancePoliciesCollectionRequestBuilder {
	bb := &ApplicationTokenIssuancePoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tokenIssuancePolicies"
	return bb
}

// ApplicationTokenIssuancePoliciesCollectionRequestBuilder is request builder for TokenIssuancePolicy collection rcn
type ApplicationTokenIssuancePoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TokenIssuancePolicy collection
func (b *ApplicationTokenIssuancePoliciesCollectionRequestBuilder) Request() *ApplicationTokenIssuancePoliciesCollectionRequest {
	return &ApplicationTokenIssuancePoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TokenIssuancePolicy item
func (b *ApplicationTokenIssuancePoliciesCollectionRequestBuilder) ID(id string) *TokenIssuancePolicyRequestBuilder {
	bb := &TokenIssuancePolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApplicationTokenIssuancePoliciesCollectionRequest is request for TokenIssuancePolicy collection
type ApplicationTokenIssuancePoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TokenIssuancePolicy collection
func (r *ApplicationTokenIssuancePoliciesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TokenIssuancePolicy, error) {
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
	var values []TokenIssuancePolicy
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
			value  []TokenIssuancePolicy
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

// GetN performs GET request for TokenIssuancePolicy collection, max N pages
func (r *ApplicationTokenIssuancePoliciesCollectionRequest) GetN(ctx context.Context, n int) ([]TokenIssuancePolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TokenIssuancePolicy collection
func (r *ApplicationTokenIssuancePoliciesCollectionRequest) Get(ctx context.Context) ([]TokenIssuancePolicy, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TokenIssuancePolicy collection
func (r *ApplicationTokenIssuancePoliciesCollectionRequest) Add(ctx context.Context, reqObj *TokenIssuancePolicy) (resObj *TokenIssuancePolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// TokenLifetimePolicies returns request builder for TokenLifetimePolicy collection
func (b *ApplicationRequestBuilder) TokenLifetimePolicies() *ApplicationTokenLifetimePoliciesCollectionRequestBuilder {
	bb := &ApplicationTokenLifetimePoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tokenLifetimePolicies"
	return bb
}

// ApplicationTokenLifetimePoliciesCollectionRequestBuilder is request builder for TokenLifetimePolicy collection rcn
type ApplicationTokenLifetimePoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TokenLifetimePolicy collection
func (b *ApplicationTokenLifetimePoliciesCollectionRequestBuilder) Request() *ApplicationTokenLifetimePoliciesCollectionRequest {
	return &ApplicationTokenLifetimePoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TokenLifetimePolicy item
func (b *ApplicationTokenLifetimePoliciesCollectionRequestBuilder) ID(id string) *TokenLifetimePolicyRequestBuilder {
	bb := &TokenLifetimePolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ApplicationTokenLifetimePoliciesCollectionRequest is request for TokenLifetimePolicy collection
type ApplicationTokenLifetimePoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TokenLifetimePolicy collection
func (r *ApplicationTokenLifetimePoliciesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]TokenLifetimePolicy, error) {
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
	var values []TokenLifetimePolicy
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
			value  []TokenLifetimePolicy
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

// GetN performs GET request for TokenLifetimePolicy collection, max N pages
func (r *ApplicationTokenLifetimePoliciesCollectionRequest) GetN(ctx context.Context, n int) ([]TokenLifetimePolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for TokenLifetimePolicy collection
func (r *ApplicationTokenLifetimePoliciesCollectionRequest) Get(ctx context.Context) ([]TokenLifetimePolicy, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for TokenLifetimePolicy collection
func (r *ApplicationTokenLifetimePoliciesCollectionRequest) Add(ctx context.Context, reqObj *TokenLifetimePolicy) (resObj *TokenLifetimePolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Application is navigation property rn
func (b *ApplicationServicePrincipalRequestBuilder) Application() *ApplicationRequestBuilder {
	bb := &ApplicationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/application"
	return bb
}

// ServicePrincipal is navigation property rn
func (b *ApplicationServicePrincipalRequestBuilder) ServicePrincipal() *ServicePrincipalRequestBuilder {
	bb := &ServicePrincipalRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/servicePrincipal"
	return bb
}

// Application returns request builder for Application collection
func (b *AdministrativeUnitMembersCollectionRequestBuilder) Application() *AdministrativeUnitMembersCollectionApplicationCollectionRequestBuilder {
	bb := &AdministrativeUnitMembersCollectionApplicationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// AdministrativeUnitMembersCollectionApplicationCollectionRequestBuilder is request builder for Application collection rcn
type AdministrativeUnitMembersCollectionApplicationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Application collection
func (b *AdministrativeUnitMembersCollectionApplicationCollectionRequestBuilder) Request() *AdministrativeUnitMembersCollectionApplicationCollectionRequest {
	return &AdministrativeUnitMembersCollectionApplicationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Application item
func (b *AdministrativeUnitMembersCollectionApplicationCollectionRequestBuilder) ID(id string) *ApplicationRequestBuilder {
	bb := &ApplicationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AdministrativeUnitMembersCollectionApplicationCollectionRequest is request for Application collection
type AdministrativeUnitMembersCollectionApplicationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Application collection
func (r *AdministrativeUnitMembersCollectionApplicationCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Application, error) {
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
	var values []Application
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
			value  []Application
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

// GetN performs GET request for Application collection, max N pages
func (r *AdministrativeUnitMembersCollectionApplicationCollectionRequest) GetN(ctx context.Context, n int) ([]Application, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Application collection
func (r *AdministrativeUnitMembersCollectionApplicationCollectionRequest) Get(ctx context.Context) ([]Application, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Application collection
func (r *AdministrativeUnitMembersCollectionApplicationCollectionRequest) Add(ctx context.Context, reqObj *Application) (resObj *Application, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *ApplicationSegmentRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *ApplicationSignInDetailedSummaryRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *ApplicationSignInSummaryRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Entity is navigation property rn
func (b *ApplicationTemplateRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}
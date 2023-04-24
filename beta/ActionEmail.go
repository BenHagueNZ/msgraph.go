// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// IncludeTargets returns request builder for AuthenticationMethodTarget collection
func (b *EmailAuthenticationMethodConfigurationRequestBuilder) IncludeTargets() *EmailAuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder {
	bb := &EmailAuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/includeTargets"
	return bb
}

// EmailAuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder is request builder for AuthenticationMethodTarget collection rcn
type EmailAuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AuthenticationMethodTarget collection
func (b *EmailAuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder) Request() *EmailAuthenticationMethodConfigurationIncludeTargetsCollectionRequest {
	return &EmailAuthenticationMethodConfigurationIncludeTargetsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AuthenticationMethodTarget item
func (b *EmailAuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder) ID(id string) *AuthenticationMethodTargetRequestBuilder {
	bb := &AuthenticationMethodTargetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EmailAuthenticationMethodConfigurationIncludeTargetsCollectionRequest is request for AuthenticationMethodTarget collection
type EmailAuthenticationMethodConfigurationIncludeTargetsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AuthenticationMethodTarget collection
func (r *EmailAuthenticationMethodConfigurationIncludeTargetsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AuthenticationMethodTarget, error) {
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
	var values []AuthenticationMethodTarget
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
			value  []AuthenticationMethodTarget
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

// GetN performs GET request for AuthenticationMethodTarget collection, max N pages
func (r *EmailAuthenticationMethodConfigurationIncludeTargetsCollectionRequest) GetN(ctx context.Context, n int) ([]AuthenticationMethodTarget, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AuthenticationMethodTarget collection
func (r *EmailAuthenticationMethodConfigurationIncludeTargetsCollectionRequest) Get(ctx context.Context) ([]AuthenticationMethodTarget, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AuthenticationMethodTarget collection
func (r *EmailAuthenticationMethodConfigurationIncludeTargetsCollectionRequest) Add(ctx context.Context, reqObj *AuthenticationMethodTarget) (resObj *AuthenticationMethodTarget, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// EmailActivityStatistics returns request builder for EmailActivityStatistics collection
func (b *UserAnalyticsActivityStatisticsCollectionRequestBuilder) EmailActivityStatistics() *UserAnalyticsActivityStatisticsCollectionEmailActivityStatisticsCollectionRequestBuilder {
	bb := &UserAnalyticsActivityStatisticsCollectionEmailActivityStatisticsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// UserAnalyticsActivityStatisticsCollectionEmailActivityStatisticsCollectionRequestBuilder is request builder for EmailActivityStatistics collection rcn
type UserAnalyticsActivityStatisticsCollectionEmailActivityStatisticsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EmailActivityStatistics collection
func (b *UserAnalyticsActivityStatisticsCollectionEmailActivityStatisticsCollectionRequestBuilder) Request() *UserAnalyticsActivityStatisticsCollectionEmailActivityStatisticsCollectionRequest {
	return &UserAnalyticsActivityStatisticsCollectionEmailActivityStatisticsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EmailActivityStatistics item
func (b *UserAnalyticsActivityStatisticsCollectionEmailActivityStatisticsCollectionRequestBuilder) ID(id string) *EmailActivityStatisticsRequestBuilder {
	bb := &EmailActivityStatisticsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UserAnalyticsActivityStatisticsCollectionEmailActivityStatisticsCollectionRequest is request for EmailActivityStatistics collection
type UserAnalyticsActivityStatisticsCollectionEmailActivityStatisticsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EmailActivityStatistics collection
func (r *UserAnalyticsActivityStatisticsCollectionEmailActivityStatisticsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EmailActivityStatistics, error) {
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
	var values []EmailActivityStatistics
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
			value  []EmailActivityStatistics
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

// GetN performs GET request for EmailActivityStatistics collection, max N pages
func (r *UserAnalyticsActivityStatisticsCollectionEmailActivityStatisticsCollectionRequest) GetN(ctx context.Context, n int) ([]EmailActivityStatistics, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EmailActivityStatistics collection
func (r *UserAnalyticsActivityStatisticsCollectionEmailActivityStatisticsCollectionRequest) Get(ctx context.Context) ([]EmailActivityStatistics, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EmailActivityStatistics collection
func (r *UserAnalyticsActivityStatisticsCollectionEmailActivityStatisticsCollectionRequest) Add(ctx context.Context, reqObj *EmailActivityStatistics) (resObj *EmailActivityStatistics, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// EmailAuthenticationMethod returns request builder for EmailAuthenticationMethod collection
func (b *AuthenticationMethodsCollectionRequestBuilder) EmailAuthenticationMethod() *AuthenticationMethodsCollectionEmailAuthenticationMethodCollectionRequestBuilder {
	bb := &AuthenticationMethodsCollectionEmailAuthenticationMethodCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// AuthenticationMethodsCollectionEmailAuthenticationMethodCollectionRequestBuilder is request builder for EmailAuthenticationMethod collection rcn
type AuthenticationMethodsCollectionEmailAuthenticationMethodCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EmailAuthenticationMethod collection
func (b *AuthenticationMethodsCollectionEmailAuthenticationMethodCollectionRequestBuilder) Request() *AuthenticationMethodsCollectionEmailAuthenticationMethodCollectionRequest {
	return &AuthenticationMethodsCollectionEmailAuthenticationMethodCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EmailAuthenticationMethod item
func (b *AuthenticationMethodsCollectionEmailAuthenticationMethodCollectionRequestBuilder) ID(id string) *EmailAuthenticationMethodRequestBuilder {
	bb := &EmailAuthenticationMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuthenticationMethodsCollectionEmailAuthenticationMethodCollectionRequest is request for EmailAuthenticationMethod collection
type AuthenticationMethodsCollectionEmailAuthenticationMethodCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EmailAuthenticationMethod collection
func (r *AuthenticationMethodsCollectionEmailAuthenticationMethodCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EmailAuthenticationMethod, error) {
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
	var values []EmailAuthenticationMethod
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
			value  []EmailAuthenticationMethod
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

// GetN performs GET request for EmailAuthenticationMethod collection, max N pages
func (r *AuthenticationMethodsCollectionEmailAuthenticationMethodCollectionRequest) GetN(ctx context.Context, n int) ([]EmailAuthenticationMethod, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EmailAuthenticationMethod collection
func (r *AuthenticationMethodsCollectionEmailAuthenticationMethodCollectionRequest) Get(ctx context.Context) ([]EmailAuthenticationMethod, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EmailAuthenticationMethod collection
func (r *AuthenticationMethodsCollectionEmailAuthenticationMethodCollectionRequest) Add(ctx context.Context, reqObj *EmailAuthenticationMethod) (resObj *EmailAuthenticationMethod, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// EmailAuthenticationMethodConfiguration returns request builder for EmailAuthenticationMethodConfiguration collection
func (b *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionRequestBuilder) EmailAuthenticationMethodConfiguration() *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionEmailAuthenticationMethodConfigurationCollectionRequestBuilder {
	bb := &AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionEmailAuthenticationMethodConfigurationCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionEmailAuthenticationMethodConfigurationCollectionRequestBuilder is request builder for EmailAuthenticationMethodConfiguration collection rcn
type AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionEmailAuthenticationMethodConfigurationCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EmailAuthenticationMethodConfiguration collection
func (b *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionEmailAuthenticationMethodConfigurationCollectionRequestBuilder) Request() *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionEmailAuthenticationMethodConfigurationCollectionRequest {
	return &AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionEmailAuthenticationMethodConfigurationCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EmailAuthenticationMethodConfiguration item
func (b *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionEmailAuthenticationMethodConfigurationCollectionRequestBuilder) ID(id string) *EmailAuthenticationMethodConfigurationRequestBuilder {
	bb := &EmailAuthenticationMethodConfigurationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionEmailAuthenticationMethodConfigurationCollectionRequest is request for EmailAuthenticationMethodConfiguration collection
type AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionEmailAuthenticationMethodConfigurationCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EmailAuthenticationMethodConfiguration collection
func (r *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionEmailAuthenticationMethodConfigurationCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EmailAuthenticationMethodConfiguration, error) {
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
	var values []EmailAuthenticationMethodConfiguration
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
			value  []EmailAuthenticationMethodConfiguration
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

// GetN performs GET request for EmailAuthenticationMethodConfiguration collection, max N pages
func (r *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionEmailAuthenticationMethodConfigurationCollectionRequest) GetN(ctx context.Context, n int) ([]EmailAuthenticationMethodConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EmailAuthenticationMethodConfiguration collection
func (r *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionEmailAuthenticationMethodConfigurationCollectionRequest) Get(ctx context.Context) ([]EmailAuthenticationMethodConfiguration, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EmailAuthenticationMethodConfiguration collection
func (r *AuthenticationMethodsPolicyAuthenticationMethodConfigurationsCollectionEmailAuthenticationMethodConfigurationCollectionRequest) Add(ctx context.Context, reqObj *EmailAuthenticationMethodConfiguration) (resObj *EmailAuthenticationMethodConfiguration, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// EmailFileAssessmentRequestObject returns request builder for EmailFileAssessmentRequestObject collection
func (b *InformationProtectionThreatAssessmentRequestsCollectionRequestBuilder) EmailFileAssessmentRequestObject() *InformationProtectionThreatAssessmentRequestsCollectionEmailFileAssessmentRequestObjectCollectionRequestBuilder {
	bb := &InformationProtectionThreatAssessmentRequestsCollectionEmailFileAssessmentRequestObjectCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// InformationProtectionThreatAssessmentRequestsCollectionEmailFileAssessmentRequestObjectCollectionRequestBuilder is request builder for EmailFileAssessmentRequestObject collection rcn
type InformationProtectionThreatAssessmentRequestsCollectionEmailFileAssessmentRequestObjectCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EmailFileAssessmentRequestObject collection
func (b *InformationProtectionThreatAssessmentRequestsCollectionEmailFileAssessmentRequestObjectCollectionRequestBuilder) Request() *InformationProtectionThreatAssessmentRequestsCollectionEmailFileAssessmentRequestObjectCollectionRequest {
	return &InformationProtectionThreatAssessmentRequestsCollectionEmailFileAssessmentRequestObjectCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EmailFileAssessmentRequestObject item
func (b *InformationProtectionThreatAssessmentRequestsCollectionEmailFileAssessmentRequestObjectCollectionRequestBuilder) ID(id string) *EmailFileAssessmentRequestObjectRequestBuilder {
	bb := &EmailFileAssessmentRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// InformationProtectionThreatAssessmentRequestsCollectionEmailFileAssessmentRequestObjectCollectionRequest is request for EmailFileAssessmentRequestObject collection
type InformationProtectionThreatAssessmentRequestsCollectionEmailFileAssessmentRequestObjectCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EmailFileAssessmentRequestObject collection
func (r *InformationProtectionThreatAssessmentRequestsCollectionEmailFileAssessmentRequestObjectCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]EmailFileAssessmentRequestObject, error) {
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
	var values []EmailFileAssessmentRequestObject
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
			value  []EmailFileAssessmentRequestObject
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

// GetN performs GET request for EmailFileAssessmentRequestObject collection, max N pages
func (r *InformationProtectionThreatAssessmentRequestsCollectionEmailFileAssessmentRequestObjectCollectionRequest) GetN(ctx context.Context, n int) ([]EmailFileAssessmentRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for EmailFileAssessmentRequestObject collection
func (r *InformationProtectionThreatAssessmentRequestsCollectionEmailFileAssessmentRequestObjectCollectionRequest) Get(ctx context.Context) ([]EmailFileAssessmentRequestObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for EmailFileAssessmentRequestObject collection
func (r *InformationProtectionThreatAssessmentRequestsCollectionEmailFileAssessmentRequestObjectCollectionRequest) Add(ctx context.Context, reqObj *EmailFileAssessmentRequestObject) (resObj *EmailFileAssessmentRequestObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

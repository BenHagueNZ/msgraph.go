// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// Alerts returns request builder for Alert collection
func (b *SecurityRequestBuilder) Alerts() *SecurityAlertsCollectionRequestBuilder {
	bb := &SecurityAlertsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/alerts"
	return bb
}

// SecurityAlertsCollectionRequestBuilder is request builder for Alert collection
type SecurityAlertsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Alert collection
func (b *SecurityAlertsCollectionRequestBuilder) Request() *SecurityAlertsCollectionRequest {
	return &SecurityAlertsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Alert item
func (b *SecurityAlertsCollectionRequestBuilder) ID(id string) *AlertRequestBuilder {
	bb := &AlertRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecurityAlertsCollectionRequest is request for Alert collection
type SecurityAlertsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Alert collection
func (r *SecurityAlertsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Alert, error) {
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
	var values []Alert
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
			value  []Alert
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

// GetN performs GET request for Alert collection, max N pages
func (r *SecurityAlertsCollectionRequest) GetN(ctx context.Context, n int) ([]Alert, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Alert collection
func (r *SecurityAlertsCollectionRequest) Get(ctx context.Context) ([]Alert, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Alert collection
func (r *SecurityAlertsCollectionRequest) Add(ctx context.Context, reqObj *Alert) (resObj *Alert, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SecureScoreControlProfiles returns request builder for SecureScoreControlProfile collection
func (b *SecurityRequestBuilder) SecureScoreControlProfiles() *SecuritySecureScoreControlProfilesCollectionRequestBuilder {
	bb := &SecuritySecureScoreControlProfilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/secureScoreControlProfiles"
	return bb
}

// SecuritySecureScoreControlProfilesCollectionRequestBuilder is request builder for SecureScoreControlProfile collection
type SecuritySecureScoreControlProfilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SecureScoreControlProfile collection
func (b *SecuritySecureScoreControlProfilesCollectionRequestBuilder) Request() *SecuritySecureScoreControlProfilesCollectionRequest {
	return &SecuritySecureScoreControlProfilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SecureScoreControlProfile item
func (b *SecuritySecureScoreControlProfilesCollectionRequestBuilder) ID(id string) *SecureScoreControlProfileRequestBuilder {
	bb := &SecureScoreControlProfileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecuritySecureScoreControlProfilesCollectionRequest is request for SecureScoreControlProfile collection
type SecuritySecureScoreControlProfilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SecureScoreControlProfile collection
func (r *SecuritySecureScoreControlProfilesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SecureScoreControlProfile, error) {
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
	var values []SecureScoreControlProfile
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
			value  []SecureScoreControlProfile
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

// GetN performs GET request for SecureScoreControlProfile collection, max N pages
func (r *SecuritySecureScoreControlProfilesCollectionRequest) GetN(ctx context.Context, n int) ([]SecureScoreControlProfile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SecureScoreControlProfile collection
func (r *SecuritySecureScoreControlProfilesCollectionRequest) Get(ctx context.Context) ([]SecureScoreControlProfile, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SecureScoreControlProfile collection
func (r *SecuritySecureScoreControlProfilesCollectionRequest) Add(ctx context.Context, reqObj *SecureScoreControlProfile) (resObj *SecureScoreControlProfile, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SecureScores returns request builder for SecureScore collection
func (b *SecurityRequestBuilder) SecureScores() *SecuritySecureScoresCollectionRequestBuilder {
	bb := &SecuritySecureScoresCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/secureScores"
	return bb
}

// SecuritySecureScoresCollectionRequestBuilder is request builder for SecureScore collection
type SecuritySecureScoresCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SecureScore collection
func (b *SecuritySecureScoresCollectionRequestBuilder) Request() *SecuritySecureScoresCollectionRequest {
	return &SecuritySecureScoresCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SecureScore item
func (b *SecuritySecureScoresCollectionRequestBuilder) ID(id string) *SecureScoreRequestBuilder {
	bb := &SecureScoreRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SecuritySecureScoresCollectionRequest is request for SecureScore collection
type SecuritySecureScoresCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SecureScore collection
func (r *SecuritySecureScoresCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SecureScore, error) {
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
	var values []SecureScore
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
			value  []SecureScore
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

// GetN performs GET request for SecureScore collection, max N pages
func (r *SecuritySecureScoresCollectionRequest) GetN(ctx context.Context, n int) ([]SecureScore, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SecureScore collection
func (r *SecuritySecureScoresCollectionRequest) Get(ctx context.Context) ([]SecureScore, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SecureScore collection
func (r *SecuritySecureScoresCollectionRequest) Add(ctx context.Context, reqObj *SecureScore) (resObj *SecureScore, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

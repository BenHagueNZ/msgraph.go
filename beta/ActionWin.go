// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// WinGetApp returns request builder for WinGetApp collection
func (b *DeviceAppManagementMobileAppsCollectionRequestBuilder) WinGetApp() *DeviceAppManagementMobileAppsCollectionWinGetAppCollectionRequestBuilder {
	bb := &DeviceAppManagementMobileAppsCollectionWinGetAppCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/"
	return bb
}

// DeviceAppManagementMobileAppsCollectionWinGetAppCollectionRequestBuilder is request builder for WinGetApp collection rcn
type DeviceAppManagementMobileAppsCollectionWinGetAppCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WinGetApp collection
func (b *DeviceAppManagementMobileAppsCollectionWinGetAppCollectionRequestBuilder) Request() *DeviceAppManagementMobileAppsCollectionWinGetAppCollectionRequest {
	return &DeviceAppManagementMobileAppsCollectionWinGetAppCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WinGetApp item
func (b *DeviceAppManagementMobileAppsCollectionWinGetAppCollectionRequestBuilder) ID(id string) *WinGetAppRequestBuilder {
	bb := &WinGetAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceAppManagementMobileAppsCollectionWinGetAppCollectionRequest is request for WinGetApp collection
type DeviceAppManagementMobileAppsCollectionWinGetAppCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WinGetApp collection
func (r *DeviceAppManagementMobileAppsCollectionWinGetAppCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]WinGetApp, error) {
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
	var values []WinGetApp
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
			value  []WinGetApp
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

// GetN performs GET request for WinGetApp collection, max N pages
func (r *DeviceAppManagementMobileAppsCollectionWinGetAppCollectionRequest) GetN(ctx context.Context, n int) ([]WinGetApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for WinGetApp collection
func (r *DeviceAppManagementMobileAppsCollectionWinGetAppCollectionRequest) Get(ctx context.Context) ([]WinGetApp, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for WinGetApp collection
func (r *DeviceAppManagementMobileAppsCollectionWinGetAppCollectionRequest) Add(ctx context.Context, reqObj *WinGetApp) (resObj *WinGetApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// ItemCategory is navigation property
func (b *ItemRequestBuilder) ItemCategory() *ItemCategoryRequestBuilder {
	bb := &ItemCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/itemCategory"
	return bb
}

// Picture returns request builder for Picture collection
func (b *ItemRequestBuilder) Picture() *ItemPictureCollectionRequestBuilder {
	bb := &ItemPictureCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/picture"
	return bb
}

// ItemPictureCollectionRequestBuilder is request builder for Picture collection
type ItemPictureCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Picture collection
func (b *ItemPictureCollectionRequestBuilder) Request() *ItemPictureCollectionRequest {
	return &ItemPictureCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Picture item
func (b *ItemPictureCollectionRequestBuilder) ID(id string) *PictureRequestBuilder {
	bb := &PictureRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ItemPictureCollectionRequest is request for Picture collection
type ItemPictureCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Picture collection
func (r *ItemPictureCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Picture, error) {
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
	var values []Picture
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
			value  []Picture
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

// GetN performs GET request for Picture collection, max N pages
func (r *ItemPictureCollectionRequest) GetN(ctx context.Context, n int) ([]Picture, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Picture collection
func (r *ItemPictureCollectionRequest) Get(ctx context.Context) ([]Picture, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Picture collection
func (r *ItemPictureCollectionRequest) Add(ctx context.Context, reqObj *Picture) (resObj *Picture, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DriveItem is navigation property
func (b *ItemActivityRequestBuilder) DriveItem() *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/driveItem"
	return bb
}

// DriveItem is navigation property
func (b *ItemActivityOLDRequestBuilder) DriveItem() *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/driveItem"
	return bb
}

// ListItem is navigation property
func (b *ItemActivityOLDRequestBuilder) ListItem() *ListItemRequestBuilder {
	bb := &ListItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/listItem"
	return bb
}

// Activities returns request builder for ItemActivity collection
func (b *ItemActivityStatRequestBuilder) Activities() *ItemActivityStatActivitiesCollectionRequestBuilder {
	bb := &ItemActivityStatActivitiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activities"
	return bb
}

// ItemActivityStatActivitiesCollectionRequestBuilder is request builder for ItemActivity collection
type ItemActivityStatActivitiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ItemActivity collection
func (b *ItemActivityStatActivitiesCollectionRequestBuilder) Request() *ItemActivityStatActivitiesCollectionRequest {
	return &ItemActivityStatActivitiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ItemActivity item
func (b *ItemActivityStatActivitiesCollectionRequestBuilder) ID(id string) *ItemActivityRequestBuilder {
	bb := &ItemActivityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ItemActivityStatActivitiesCollectionRequest is request for ItemActivity collection
type ItemActivityStatActivitiesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ItemActivity collection
func (r *ItemActivityStatActivitiesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ItemActivity, error) {
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
	var values []ItemActivity
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
			value  []ItemActivity
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

// GetN performs GET request for ItemActivity collection, max N pages
func (r *ItemActivityStatActivitiesCollectionRequest) GetN(ctx context.Context, n int) ([]ItemActivity, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ItemActivity collection
func (r *ItemActivityStatActivitiesCollectionRequest) Get(ctx context.Context) ([]ItemActivity, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ItemActivity collection
func (r *ItemActivityStatActivitiesCollectionRequest) Add(ctx context.Context, reqObj *ItemActivity) (resObj *ItemActivity, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AllTime is navigation property
func (b *ItemAnalyticsRequestBuilder) AllTime() *ItemActivityStatRequestBuilder {
	bb := &ItemActivityStatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/allTime"
	return bb
}

// ItemActivityStats returns request builder for ItemActivityStat collection
func (b *ItemAnalyticsRequestBuilder) ItemActivityStats() *ItemAnalyticsItemActivityStatsCollectionRequestBuilder {
	bb := &ItemAnalyticsItemActivityStatsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/itemActivityStats"
	return bb
}

// ItemAnalyticsItemActivityStatsCollectionRequestBuilder is request builder for ItemActivityStat collection
type ItemAnalyticsItemActivityStatsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ItemActivityStat collection
func (b *ItemAnalyticsItemActivityStatsCollectionRequestBuilder) Request() *ItemAnalyticsItemActivityStatsCollectionRequest {
	return &ItemAnalyticsItemActivityStatsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ItemActivityStat item
func (b *ItemAnalyticsItemActivityStatsCollectionRequestBuilder) ID(id string) *ItemActivityStatRequestBuilder {
	bb := &ItemActivityStatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ItemAnalyticsItemActivityStatsCollectionRequest is request for ItemActivityStat collection
type ItemAnalyticsItemActivityStatsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ItemActivityStat collection
func (r *ItemAnalyticsItemActivityStatsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]ItemActivityStat, error) {
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
	var values []ItemActivityStat
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
			value  []ItemActivityStat
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

// GetN performs GET request for ItemActivityStat collection, max N pages
func (r *ItemAnalyticsItemActivityStatsCollectionRequest) GetN(ctx context.Context, n int) ([]ItemActivityStat, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for ItemActivityStat collection
func (r *ItemAnalyticsItemActivityStatsCollectionRequest) Get(ctx context.Context) ([]ItemActivityStat, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for ItemActivityStat collection
func (r *ItemAnalyticsItemActivityStatsCollectionRequest) Add(ctx context.Context, reqObj *ItemActivityStat) (resObj *ItemActivityStat, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// LastSevenDays is navigation property
func (b *ItemAnalyticsRequestBuilder) LastSevenDays() *ItemActivityStatRequestBuilder {
	bb := &ItemActivityStatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/lastSevenDays"
	return bb
}

// Item is navigation property
func (b *ItemAttachmentRequestBuilder) Item() *OutlookItemRequestBuilder {
	bb := &OutlookItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/item"
	return bb
}

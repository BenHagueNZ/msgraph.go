// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// Channels returns request builder for Channel collection
func (b *DeletedTeamRequestBuilder) Channels() *DeletedTeamChannelsCollectionRequestBuilder {
	bb := &DeletedTeamChannelsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/channels"
	return bb
}

// DeletedTeamChannelsCollectionRequestBuilder is request builder for Channel collection
type DeletedTeamChannelsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Channel collection
func (b *DeletedTeamChannelsCollectionRequestBuilder) Request() *DeletedTeamChannelsCollectionRequest {
	return &DeletedTeamChannelsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Channel item
func (b *DeletedTeamChannelsCollectionRequestBuilder) ID(id string) *ChannelRequestBuilder {
	bb := &ChannelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeletedTeamChannelsCollectionRequest is request for Channel collection
type DeletedTeamChannelsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Channel collection
func (r *DeletedTeamChannelsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Channel, error) {
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
	var values []Channel
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
			value  []Channel
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

// GetN performs GET request for Channel collection, max N pages
func (r *DeletedTeamChannelsCollectionRequest) GetN(ctx context.Context, n int) ([]Channel, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Channel collection
func (r *DeletedTeamChannelsCollectionRequest) Get(ctx context.Context) ([]Channel, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Channel collection
func (r *DeletedTeamChannelsCollectionRequest) Add(ctx context.Context, reqObj *Channel) (resObj *Channel, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Entity is navigation property rn
func (b *DeletedTeamRequestBuilder) Entity() *EntityRequestBuilder {
	bb := &EntityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/Entity"
	return bb
}

// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BenHagueNZ/msgraph.go/jsonx"
)

// Currency is navigation property
func (b *CustomerRequestBuilder) Currency() *CurrencyRequestBuilder {
	bb := &CurrencyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/currency"
	return bb
}

// PaymentMethod is navigation property
func (b *CustomerRequestBuilder) PaymentMethod() *PaymentMethodRequestBuilder {
	bb := &PaymentMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/paymentMethod"
	return bb
}

// PaymentTerm is navigation property
func (b *CustomerRequestBuilder) PaymentTerm() *PaymentTermRequestBuilder {
	bb := &PaymentTermRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/paymentTerm"
	return bb
}

// Picture returns request builder for Picture collection
func (b *CustomerRequestBuilder) Picture() *CustomerPictureCollectionRequestBuilder {
	bb := &CustomerPictureCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/picture"
	return bb
}

// CustomerPictureCollectionRequestBuilder is request builder for Picture collection
type CustomerPictureCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Picture collection
func (b *CustomerPictureCollectionRequestBuilder) Request() *CustomerPictureCollectionRequest {
	return &CustomerPictureCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Picture item
func (b *CustomerPictureCollectionRequestBuilder) ID(id string) *PictureRequestBuilder {
	bb := &PictureRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CustomerPictureCollectionRequest is request for Picture collection
type CustomerPictureCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Picture collection
func (r *CustomerPictureCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]Picture, error) {
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
func (r *CustomerPictureCollectionRequest) GetN(ctx context.Context, n int) ([]Picture, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for Picture collection
func (r *CustomerPictureCollectionRequest) Get(ctx context.Context) ([]Picture, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for Picture collection
func (r *CustomerPictureCollectionRequest) Add(ctx context.Context, reqObj *Picture) (resObj *Picture, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ShipmentMethod is navigation property
func (b *CustomerRequestBuilder) ShipmentMethod() *ShipmentMethodRequestBuilder {
	bb := &ShipmentMethodRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/shipmentMethod"
	return bb
}

// Customer is navigation property
func (b *CustomerPaymentRequestBuilder) Customer() *CustomerRequestBuilder {
	bb := &CustomerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/customer"
	return bb
}

// Account is navigation property
func (b *CustomerPaymentJournalRequestBuilder) Account() *AccountRequestBuilder {
	bb := &AccountRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/account"
	return bb
}

// CustomerPayments returns request builder for CustomerPayment collection
func (b *CustomerPaymentJournalRequestBuilder) CustomerPayments() *CustomerPaymentJournalCustomerPaymentsCollectionRequestBuilder {
	bb := &CustomerPaymentJournalCustomerPaymentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/customerPayments"
	return bb
}

// CustomerPaymentJournalCustomerPaymentsCollectionRequestBuilder is request builder for CustomerPayment collection
type CustomerPaymentJournalCustomerPaymentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CustomerPayment collection
func (b *CustomerPaymentJournalCustomerPaymentsCollectionRequestBuilder) Request() *CustomerPaymentJournalCustomerPaymentsCollectionRequest {
	return &CustomerPaymentJournalCustomerPaymentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CustomerPayment item
func (b *CustomerPaymentJournalCustomerPaymentsCollectionRequestBuilder) ID(id string) *CustomerPaymentRequestBuilder {
	bb := &CustomerPaymentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CustomerPaymentJournalCustomerPaymentsCollectionRequest is request for CustomerPayment collection
type CustomerPaymentJournalCustomerPaymentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CustomerPayment collection
func (r *CustomerPaymentJournalCustomerPaymentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]CustomerPayment, error) {
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
	var values []CustomerPayment
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
			value  []CustomerPayment
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

// GetN performs GET request for CustomerPayment collection, max N pages
func (r *CustomerPaymentJournalCustomerPaymentsCollectionRequest) GetN(ctx context.Context, n int) ([]CustomerPayment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for CustomerPayment collection
func (r *CustomerPaymentJournalCustomerPaymentsCollectionRequest) Get(ctx context.Context) ([]CustomerPayment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for CustomerPayment collection
func (r *CustomerPaymentJournalCustomerPaymentsCollectionRequest) Add(ctx context.Context, reqObj *CustomerPayment) (resObj *CustomerPayment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

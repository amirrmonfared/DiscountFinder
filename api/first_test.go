package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/amirrmonfared/DiscountFinder/db/mock"
	db "github.com/amirrmonfared/DiscountFinder/db/sqlc"
	"github.com/amirrmonfared/DiscountFinder/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func randomFirstProductId() db.First {
	return db.First{
		ID:    util.RandomInt(1, 5),
		Brand: util.RandomString(5),
		Link:  util.RandomLink(),
		Price: util.RandomPriceString(5),
	}
}
func TestGetFirstProductAPI(t *testing.T) {
	product := randomFirstProductId()

	testCases := []struct {
		name          string
		productID     int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			productID: product.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetFirstProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(product, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, product)
			},
		},
		{
			name:      "NotFound",
			productID: product.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetFirstProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(db.First{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "InternalError",
			productID: product.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetFirstProduct(gomock.Any(), gomock.Eq(product.ID)).
					Times(1).
					Return(db.First{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "InvalidID",
			productID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetFirstProduct(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/product/%d", tc.productID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomFirstProduct() db.First {
	return db.First{
		Brand: util.RandomString(5),
		Link:  util.RandomLink(),
		Price: util.RandomPriceString(5),
	}
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.First) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotFirst db.First
	err = json.Unmarshal(data, &gotFirst)
	require.NoError(t, err)
	require.Equal(t, user, gotFirst)
}

func TestCreateFirstProductAPI(t *testing.T) {
	product := randomFirstProduct()

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"brand": product.Brand,
				"link":  product.Link,
				"price": product.Price,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateFirstProductParams{

					Brand: product.Brand,
					Link:  product.Link,
					Price: product.Price,
				}
				store.EXPECT().
					CreateFirstProduct(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(product, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, product)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"brand": product.Brand,
				"link":  product.Link,
				"price": product.Price,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateFirstProduct(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.First{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "DuplicateProduct",
			body: gin.H{
				"brand": product.Brand,
				"link":  product.Link,
				"price": product.Price,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateFirstProduct(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.First{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "InvalidPrice",
			body: gin.H{
				"brand": product.Brand,
				"link":  product.Link,
				"price": 1,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateFirstProduct(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/product"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

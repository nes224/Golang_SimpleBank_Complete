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

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/techschool/simplebank/db/mock"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/util"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	testCases := []struct {
		name          string
		accountID     int64
		buildStubs    func(store *mockdb.MockStore)                           // buildStubs field, which is actually a function that takes a mock store as input.
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder) // function to check the output of the API.
	}{
		{
			name:      "OK",
			accountID: account.ID,
			// build stubs
			// stub for our mock store.
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccountForUpdate(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name:      "NotFound",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccountForUpdate(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:      "InternalError",
			accountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccountForUpdate(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}
	
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) { 
			ctrl := gomock.NewController(t)
			defer ctrl.Finish() // We should defer calling Finish method of this controller.
	
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)
			// start test server and send request
			server := NewServer(store)
			// For testing an HTTP API in Go, we don't have to start a real HTTP server, Instead, we can just use the Recorder feature of the httptest package to record the response of the API request.
			recorder := httptest.NewRecorder() // So here we call httptest.NewRecorder() to create a new ResponseRecorder.
			url := fmt.Sprintf("/accounts/%d", account.ID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t,err)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
// require.NoError(t, err)
// server.router.ServeHTTP(recorder, request) // This will send our API request through the server router and record its response in the recorder.
// // check response
// require.Equal(t, http.StatusOK, recorder.Code)
// requireBodyMatchAccount(t, recorder.Body, account)

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount) // call json.Unmarshal to unmarshal the data to the gotAccount object.
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}

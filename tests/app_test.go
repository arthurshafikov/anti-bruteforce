package tests

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/thewolf27/anti-bruteforce/internal/models"
)

type appSuiteHandler struct {
	AppSuite
}

func TestAppSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip test app")
	}

	appSuiteHandler := &appSuiteHandler{
		AppSuite: AppSuite{},
	}

	suite.Run(t, appSuiteHandler)
}

func (h *appSuiteHandler) TestTryToAuthorizeWithResetBucket() {
	for i := 0; i < limitAttemptsForLogin; i++ {
		statusCode, body := h.tryToAuthorize(authorizeInput)
		require.Equal(h.T(), http.StatusOK, statusCode)
		require.Equal(h.T(), successResponse, body)
	}

	statusCode, body := h.tryToAuthorize(authorizeInput)
	require.Equal(h.T(), http.StatusTooManyRequests, statusCode)
	require.Equal(h.T(), tooManyRequestsResponse, body)

	recorder := h.makeServerHTTPRequest(http.MethodPost, "/bucket/reset", nil)
	statusCode, body = h.getStatusCodeAndBodyFromRecorder(recorder)
	require.Equal(h.T(), http.StatusOK, statusCode)
	require.Equal(h.T(), successResponse, body)

	for i := 0; i < limitAttemptsForLogin; i++ {
		statusCode, body := h.tryToAuthorize(authorizeInput)
		require.Equal(h.T(), http.StatusOK, statusCode)
		require.Equal(h.T(), successResponse, body)
	}
}

func (h *appSuiteHandler) TestAddToWhiteListAndTryToAuthorizeMoreThanLimit() {
	subnet := getJSONBody(h.T(), subnetInput)

	recorder := h.makeServerHTTPRequest(http.MethodPost, "/whitelist/add", bytes.NewBuffer(subnet))
	statusCode, body := h.getStatusCodeAndBodyFromRecorder(recorder)
	require.Equal(h.T(), http.StatusCreated, statusCode)
	require.Equal(h.T(), successResponse, body)

	for i := 0; i < (limitAttemptsForLogin * 2); i++ {
		statusCode, body := h.tryToAuthorize(authorizeInput)
		require.Equal(h.T(), http.StatusOK, statusCode)
		require.Equal(h.T(), successResponse, body)
	}
}

func (h *appSuiteHandler) TestAddToBlackListAndTryToAuthorize() {
	subnet := getJSONBody(h.T(), subnetInput)

	recorder := h.makeServerHTTPRequest(http.MethodPost, "/blacklist/add", bytes.NewBuffer(subnet))
	statusCode, body := h.getStatusCodeAndBodyFromRecorder(recorder)
	require.Equal(h.T(), http.StatusCreated, statusCode)
	require.Equal(h.T(), successResponse, body)

	statusCode, body = h.tryToAuthorize(authorizeInput)
	require.Equal(h.T(), http.StatusTooManyRequests, statusCode)
	require.Equal(h.T(), tooManyRequestsResponse, body)
}

func (h *appSuiteHandler) TestAddToBlackListWrongSubnetFormat() {
	subnetStr := "asfasfafs"
	subnet := getJSONBody(h.T(), models.SubnetInput{
		Subnet: subnetStr,
	})

	recorder := h.makeServerHTTPRequest(http.MethodPost, "/blacklist/add", bytes.NewBuffer(subnet))
	statusCode, body := h.getStatusCodeAndBodyFromRecorder(recorder)
	require.Equal(h.T(), http.StatusUnprocessableEntity, statusCode)
	require.Equal(h.T(), fmt.Sprintf(duplicateSubnetResponseFormat, subnetStr), body)
}

func (h *appSuiteHandler) TestManyIPsTryingToAuthorize() {
	subnet := getJSONBody(h.T(), subnetInput)
	recorder := h.makeServerHTTPRequest(http.MethodPost, "/whitelist/add", bytes.NewBuffer(subnet))
	statusCode, body := h.getStatusCodeAndBodyFromRecorder(recorder)
	require.Equal(h.T(), http.StatusCreated, statusCode)
	require.Equal(h.T(), successResponse, body)

	subnetInput.Subnet = "128.10/16"
	subnet = getJSONBody(h.T(), subnetInput)
	recorder = h.makeServerHTTPRequest(http.MethodPost, "/blacklist/add", bytes.NewBuffer(subnet))
	statusCode, body = h.getStatusCodeAndBodyFromRecorder(recorder)
	require.Equal(h.T(), http.StatusCreated, statusCode)
	require.Equal(h.T(), successResponse, body)

	IPMapWithExpectedAuthorizeOutput := map[string]bool{ // ip:shouldProceed
		"198.24.15.5":  true,  // whitelisted
		"198.24.15.10": true,  // whitelisted
		"128.10.20.5":  false, // blacklisted
		"128.10.40.1":  false, // blacklisted
		"128.10.0.25":  false, // blacklisted
		"130.5.10.1":   true,  // normal
		"226.4.25.16":  true,  // normal
	}

	for ip, shouldAuthorize := range IPMapWithExpectedAuthorizeOutput {
		authorizeInput.IP = ip
		statusCode, body := h.tryToAuthorize(authorizeInput)

		if shouldAuthorize {
			require.Equal(h.T(), http.StatusOK, statusCode)
			require.Equal(h.T(), successResponse, body)
		} else {
			require.Equal(h.T(), http.StatusTooManyRequests, statusCode)
			require.Equal(h.T(), tooManyRequestsResponse, body)
		}
	}
}

func (h *appSuiteHandler) TestAddAndRemoveFromWhiteList() {
	subnet := getJSONBody(h.T(), subnetInput)

	recorder := h.makeServerHTTPRequest(http.MethodPost, "/blacklist/add", bytes.NewBuffer(subnet))
	statusCode, body := h.getStatusCodeAndBodyFromRecorder(recorder)
	require.Equal(h.T(), http.StatusCreated, statusCode)
	require.Equal(h.T(), successResponse, body)

	recorder = h.makeServerHTTPRequest(http.MethodDelete, "/blacklist/remove", bytes.NewBuffer(subnet))
	statusCode, body = h.getStatusCodeAndBodyFromRecorder(recorder)
	require.Equal(h.T(), http.StatusOK, statusCode)
	require.Equal(h.T(), successResponse, body)
}

func (h *appSuiteHandler) tryToAuthorize(authorizeInput models.AuthorizeInput) (int, string) {
	recorder := h.makeServerHTTPRequest(http.MethodPost, "/authorize", bytes.NewBuffer(getJSONBody(h.T(), authorizeInput)))

	return h.getStatusCodeAndBodyFromRecorder(recorder)
}

func (h *appSuiteHandler) makeServerHTTPRequest(method, route string, body io.Reader) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, route, body)
	recorder := httptest.NewRecorder()
	h.ServerEngine.ServeHTTP(recorder, req)

	req.Body.Close()

	return recorder
}

func (h *appSuiteHandler) getStatusCodeAndBodyFromRecorder(recorder *httptest.ResponseRecorder) (int, string) {
	recorderResult := recorder.Result()
	bodyBytes, err := ioutil.ReadAll(recorderResult.Body)
	require.NoError(h.T(), err)

	recorderResult.Body.Close()

	return recorderResult.StatusCode, string(bodyBytes)
}

package api_test

import (
	. "api"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGreetingHandleFunc(t *testing.T) {
	expectedResult := `{"message":"hello world"}`
	requestMessage := `{"message":"world"}`
	request := httptest.NewRequest(http.MethodPost, "/greeting", strings.NewReader(requestMessage))
	responseRecorder := httptest.NewRecorder()

	HelloPost(responseRecorder, request)

	response := responseRecorder.Result()
	responseBody, _ := ioutil.ReadAll(response.Body)
	actualResult := strings.TrimSpace(string(responseBody))

	if expectedResult != actualResult {
		t.Errorf("Expect %v but it got %v", expectedResult, actualResult)
	}
}

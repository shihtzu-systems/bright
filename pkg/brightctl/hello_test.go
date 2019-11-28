package brightctl

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloController_HandleHello(t *testing.T) {
	// setup
	sut := HelloController{}

	sutRequest, err := http.NewRequest("GET", HelloPath(), nil)
	if err != nil {
		t.Fatal(err)
	}
	result := httptest.NewRecorder()

	// test
	sut.HandleRoot(result, sutRequest)

	// verify
	if result.Code != http.StatusOK {
		t.Fail()
	}
}

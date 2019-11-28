package brightctl

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBingoController_HandleBingo(t *testing.T) {
	// setup
	sut := BingoController{}

	sutRequest, err := http.NewRequest("GET", BingoPath(), nil)
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

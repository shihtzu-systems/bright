package brightctl

import (
	"github.com/gorilla/sessions"
	"github.com/shihtzu-systems/bright/pkg/sorry"
	"github.com/shihtzu-systems/bright/pkg/tower"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBnetController_HandleBnet(t *testing.T) {
	// setup
	sut := newTestBnetController()

	sutRequest, err := http.NewRequest("GET", BnetPath(), nil)
	if err != nil {
		t.Fatal(err)
	}
	result := httptest.NewRecorder()

	// test
	sut.HandleRoot(result, sutRequest)

	// verify
	expectedBody := "Hello"

	resultBody := result.Body.String()
	if resultBody != expectedBody {
		sorry.Unexpected(t, resultBody, expectedBody)
		t.Fail()
	}
}

func newTestBnetController() BnetController {
	return BnetController{
		SessionStore: sessions.NewCookieStore([]byte("session-secret")),
		Tower: tower.Tower{
			App: tower.App{},
			System: tower.System{
				Id: "system",
			},
			SessionKey: "session-key",
			Redis: tower.Redis{
				Address:  "localhost:6379",
				Database: 10,
			},
		},
	}
}

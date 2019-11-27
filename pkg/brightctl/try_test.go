package brightctl

import (
	"github.com/gorilla/sessions"
	"github.com/shihtzu-systems/bright/pkg/sorry"
	"github.com/shihtzu-systems/bright/pkg/tower"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTryController_HandleTry(t *testing.T) {
	// setup
	sut := newTestTryController()

	sutRequest, err := http.NewRequest("GET", TryPath(), nil)
	if err != nil {
		t.Fatal(err)
	}
	result := httptest.NewRecorder()

	// test
	sut.HandleTry(result, sutRequest)

	// verify
	resultBody := result.Body.String()
	if resultBody == "" {
		sorry.UnexpectedBlank(t)
		t.Fail()
	}
}

func TestTryController_HandleTryRecycle(t *testing.T) {
	// setup
	sut := newTestTryController()

	sutRequest, err := http.NewRequest("GET", TryPath("recycle"), nil)
	if err != nil {
		t.Fatal(err)
	}
	result := httptest.NewRecorder()

	// test
	sut.HandleTryRecycle(result, sutRequest)

	// verify
	resultLocation := result.Header().Get("Location")
	if resultLocation != TryPath() {
		sorry.Unexpected(t, resultLocation, TryPath())
		t.Fail()
	}
}

func newTestTryController() TryController {
	return TryController{
		SessionStore: sessions.NewCookieStore([]byte("session-secret")),
		Tower: tower.Tower{
			Serial: "v1.2.3",
			System: tower.System{
				Id: "system",
			},
			SessionKey: "session-key",
			Redis: tower.Redis{
				Address:  "localhost",
				Port:     6379,
				Database: 10,
			},
			TemplatesPath: "../../tpl",
		},
	}
}

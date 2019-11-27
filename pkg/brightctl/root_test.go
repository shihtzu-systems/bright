package brightctl

import (
	"github.com/gorilla/sessions"
	"github.com/shihtzu-systems/bright/pkg/sorry"
	"github.com/shihtzu-systems/bright/pkg/tower"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootController_HandleRoot(t *testing.T) {
	// setup
	sut := newTestRootController()

	sutRequest, err := http.NewRequest("GET", RootPath(), nil)
	if err != nil {
		t.Fatal(err)
	}
	result := httptest.NewRecorder()

	// test
	sut.HandleRoot(result, sutRequest)

	// verify
	resultBody := result.Body.String()
	if resultBody == "" {
		sorry.UnexpectedBlank(t)
		t.Fail()
	}
}

func newTestRootController() RootController {
	return RootController{
		SessionStore: sessions.NewCookieStore([]byte("session-secret")),
		Tower: tower.Tower{
			Serial: "v1.2.3",
			System: tower.System{
				Id: "system",
			},
			SessionKey: "session-key",
			Redis: tower.Redis{
				Address:  "localhost:6379",
				Database: 10,
			},
			TemplatesPath: "../../tpl",
		},
	}
}

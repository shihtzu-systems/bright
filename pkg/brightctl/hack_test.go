package brightctl

import (
	"github.com/gorilla/sessions"
	"github.com/shihtzu-systems/bright/pkg/sorry"
	"github.com/shihtzu-systems/bright/pkg/tower"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
)

func TestHackController_HandleHack(t *testing.T) {
	// setup
	sut := newTestHackController()

	sutRequest, err := http.NewRequest("GET", HackPath(), nil)
	if err != nil {
		t.Fatal(err)
	}
	result := httptest.NewRecorder()

	// test
	sut.HandleRoot(result, sutRequest)

	// verify
	expectedLocation := path.Join(bnetBasePath)

	resultLocation := result.Header().Get("Location")
	if resultLocation != expectedLocation {
		sorry.Unexpected(t, resultLocation, expectedLocation)
		t.Fail()
	}
}

func newTestHackController() HackController {
	return HackController{
		SessionStore: sessions.NewCookieStore([]byte("session-secret")),
		Tower: tower.Tower{
			App: tower.App{},
			System: tower.System{
				Id: "system",
			},
			SessionKey: "session-key",
			Redis: tower.Redis{
				Address:  "localhost",
				Port:     6379,
				Database: 10,
			},
		},
		HackToken: tower.HackToken{
			AccessToken:      "access-token",
			TokenType:        "Bearer",
			ExpiresIn:        3600,
			RefreshToken:     "refresh-token",
			RefreshExpiresIn: 7776000,
			MembershipId:     "12345",
		},
	}
}

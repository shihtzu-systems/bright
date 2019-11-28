package brightx

import (
	"context"
	haikunator "github.com/atrox/haikunatorgo/v2"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	. "github.com/shihtzu-systems/bright/pkg/brightctl"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type ServeArgs struct {
	BungieClient *resty.Client
	Tower        tower.Tower
	HackToken    tower.HackToken
}

func Serve(args ServeArgs) {

	r := mux.NewRouter()

	// setup
	if args.Tower.Trace {
		log.SetLevel(log.TraceLevel)
	} else if args.Tower.Debug {
		log.SetLevel(log.DebugLevel)
	}

	// guts
	t := args.Tower
	sessionStore := sessions.NewCookieStore(t.SessionSecret)

	// bingo controller
	bingo := BingoController{
		SessionStore: sessionStore,
		Tower:        t,
	}
	r.HandleFunc(BingoPath(), bingo.HandleRoot)

	// hello controller
	hello := HelloController{}
	r.HandleFunc(HelloPath(), hello.HandleRoot)

	// hack controller
	hack := HackController{
		SessionStore: sessionStore,
		Tower:        t,
		HackToken:    args.HackToken,
	}
	r.HandleFunc(HackPath(), hack.HandleRoot)

	// bnet controller
	bnet := BnetController{
		SessionStore: sessionStore,
		Tower:        t,
		BungieClient: args.BungieClient,
	}
	r.HandleFunc(BnetPath("auth"), bnet.HandleAuth)
	r.HandleFunc(BnetPath("callback"), bnet.HandleCallback).Queries("code", "{code}", "state", "{state}")
	r.HandleFunc(BnetPath(), bnet.HandleRoot)
	r.HandleFunc(BnetPath("guardian", "{membershipType:[a-z0-9-]+}", "{id:[a-z0-9-]+}", "init"), bnet.HandleGuardianInit)
	r.HandleFunc(BnetPath("guardian", "{membershipType:[a-z0-9-]+}", "{id:[a-z0-9-]+}"), bnet.HandleGuardian)
	r.HandleFunc(BnetPath("guardian", "{membershipType:[a-z0-9-]+}", "{id:[a-z0-9-]+}", "swap", "weapon", "{weaponHash:[a-z0-9-]+}"), bnet.HandleSwapWeapon)
	r.HandleFunc(BnetPath("guardian", "{membershipType:[a-z0-9-]+}", "{id:[a-z0-9-]+}", "swap", "armor", "{armorHash:[a-z0-9-]+}"), bnet.HandleSwapArmor)
	r.HandleFunc(BnetPath("guardian", "{membershipType:[a-z0-9-]+}", "{id:[a-z0-9-]+}", "swap", "{swapOut:[a-z0-9-]+}", "{swapIn:[a-z0-9-]+}"), bnet.HandleSwap)

	// try controller
	try := TryController{
		SessionStore: sessionStore,
		Tower:        t,
	}
	r.HandleFunc(TryPath(), try.HandleRoot)
	r.HandleFunc(TryPath("recycle"), try.HandleRecycle)
	r.HandleFunc(TryPath("guardian", "0", "{id:[a-z0-9-]+}"), try.HandleGuardian)
	r.HandleFunc(TryPath("guardian", "0", "{id:[a-z0-9-]+}", "swap", "weapon", "{weaponHash:[a-z0-9-]+}"), try.HandleSwapWeapon)
	r.HandleFunc(TryPath("guardian", "0", "{id:[a-z0-9-]+}", "swap", "armor", "{armorHash:[a-z0-9-]+}"), try.HandleSwapArmor)
	r.HandleFunc(TryPath("guardian", "0", "{id:[a-z0-9-]+}", "swap", "{swapOut:[a-z0-9-]+}", "{swapIn:[a-z0-9-]+}"), try.HandleSwap)

	// root controller
	root := RootController{
		SessionStore: sessionStore,
		Tower:        t,
	}
	r.HandleFunc(RootPath(), root.HandleRoot)

	// static
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static/"))))

	// server startup
	namer := haikunator.New()
	namer.TokenLength = 0
	namer.Delimiter = " "
	name := namer.Haikunate()
	log.Printf("starting v%s as %s", args.Tower.Serial, name)
	log.Printf("listening on localhost:8080")
	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// listen
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	// server teardown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	_ = srv.Shutdown(ctx)
	log.Info("shutting down ", name)
	os.Exit(0)
}

package brightx

import (
	"context"
	haikunator "github.com/atrox/haikunatorgo/v2"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/opentracing-contrib/go-gorilla/gorilla"
	. "github.com/shihtzu-systems/bright/pkg/brightctl"
	"github.com/shihtzu-systems/bright/pkg/tower"
	log "github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go/config"
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
	//TODO contentPath := args.WorkingBasePath
	t := args.Tower
	sessionStore := sessions.NewCookieStore(t.SessionSecret)

	// bingo controller
	bingo := BingoController{
		SessionStore: sessionStore,
		Tower:        t,
	}
	r.HandleFunc(BingoPath(), bingo.HandleBingo)

	// hello controller
	hello := HelloController{}
	r.HandleFunc(HelloPath(), hello.HandleHello)

	// hack controller
	hack := HackController{
		SessionStore: sessionStore,
		Tower:        t,
		HackToken:    args.HackToken,
	}
	r.HandleFunc(HackPath(), hack.HandleHack)

	// bnet controller
	bnet := BnetController{
		SessionStore: sessionStore,
		Tower:        t,
	}
	r.HandleFunc(BnetPath(), bnet.HandleBnet)

	// try controller
	try := TryController{
		SessionStore: sessionStore,
		Tower:        t,
	}
	r.HandleFunc(TryPath(), try.HandleTry)
	r.HandleFunc(TryPath("recycle"), try.HandleTryRecycle)
	r.HandleFunc(TryPath("guardian", "0", "{id:[a-z0-9-]+}"), try.HandleTryGuardian)
	r.HandleFunc(TryPath("guardian", "0", "{id:[a-z0-9-]+}", "swap", "weapon", "{weaponHash:[a-z0-9-]+}"), try.HandleTrySwapWeapon)
	r.HandleFunc(TryPath("guardian", "0", "{id:[a-z0-9-]+}", "swap", "armor", "{armorHash:[a-z0-9-]+}"), try.HandleTrySwapArmor)
	r.HandleFunc(TryPath("guardian", "0", "{id:[a-z0-9-]+}", "swap", "{swapOut:[a-z0-9-]+}", "{swapIn:[a-z0-9-]+}"), try.HandleTrySwap)

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

	// tracing middleware
	jaegerConfig := config.Configuration{
		ServiceName: "bright",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  t.JaegerAgent,
		},
	}
	tracer, closer, err := jaegerConfig.NewTracer()
	defer func() {
		if err := closer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}
	_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		route.Handler(
			gorilla.Middleware(tracer, route.GetHandler()))
		return nil
	})

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

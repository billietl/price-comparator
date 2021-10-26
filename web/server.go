package web

import (
	"fmt"
	"net/http"
	"os"
	"price-comparator/dao"
	"price-comparator/web/api"
	"price-comparator/web/pages"
	"time"

	"github.com/gorilla/handlers"
	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

type Server struct {
	port   int
	router *Router
}

func MakeServer(port int, dao *dao.Bundle) (server *Server) {
	router := NewRouter()

	router.RegisterController(
		api.NewProductController(dao),
		"/api/v1/product",
	)

	router.RegisterController(
		api.NewStoreController(dao),
		"/api/v1/store",
	)

	router.RegisterController(
		pages.NewStoreController(dao),
		"/stores/",
	)

	router.RegisterController(
		pages.NewIndexController(),
		"/",
	)

	return &Server{
		port:   port,
		router: router,
	}
}

func (serv Server) computeMiddlewares(appRouter *Router) (finalRouter http.Handler) {
	chain := alice.New()
	// Logging middleware
	wr := diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
		fmt.Printf("Logger Dropped %d messages", missed)
	})
	logger := zerolog.New(wr).With().
		Timestamp().
		Logger()
	chain = chain.Append(hlog.NewHandler(logger))
	chain = chain.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))
	chain = chain.Append(hlog.RemoteAddrHandler("ip"))
	chain = chain.Append(hlog.UserAgentHandler("user_agent"))
	chain = chain.Append(hlog.RefererHandler("referer"))
	chain = chain.Append(hlog.RequestIDHandler("req_id", "Request-Id"))
	// Performance middleware
	chain = chain.Append(handlers.CompressHandler)
	return chain.Then(serv.router)
}

func (serv Server) Run() error {
	log.Info().Msg(fmt.Sprintf("Starting server on port %d", serv.port))
	log.Fatal().Msgf(
		"%s",
		http.ListenAndServe(
			fmt.Sprintf(":%d", serv.port),
			serv.computeMiddlewares(serv.router),
		),
	)
	log.Info().Msg("Shuting server")

	return nil
}

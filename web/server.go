package web

import (
	"fmt"
	"net/http"
	"price-comparator/dao"
	"price-comparator/logger"
	"price-comparator/web/api"
	"price-comparator/web/pages"
	"time"

	"github.com/gorilla/handlers"
	"github.com/justinas/alice"
	"github.com/rs/zerolog/hlog"
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
	chain = chain.Append(hlog.NewHandler(*logger.GetAccessLogger()))
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
	logger.Info(fmt.Sprintf("Starting server on port %d", serv.port))
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", serv.port),
		serv.computeMiddlewares(serv.router),
	)
	if err != nil {
		return err
	}
	logger.Info("Shuting server")
	return nil
}

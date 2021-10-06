package restserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/josephzxy/timer_apiserver/internal/pkg/util"
	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
	"github.com/josephzxy/timer_apiserver/internal/restserver/controller/v1/timer"
	"github.com/josephzxy/timer_apiserver/internal/restserver/middleware"
	resp "github.com/josephzxy/timer_apiserver/internal/restserver/response"
)

type RESTServer interface {
	Start() error
	Stop() error
}

type restServer struct {
	cfg           Config
	serviceRouter service.ServiceRouter
	*gin.Engine
	insecureServer *http.Server
}

func New(cfg *Config, serviceRouter service.ServiceRouter) RESTServer {
	s := &restServer{
		cfg:           *cfg,
		serviceRouter: serviceRouter,
		Engine:        gin.New(),
	}
	s.init()
	return s
}

func (s *restServer) init() {
	gin.SetMode(s.cfg.Mode)
	s.installMiddlewares()
	s.installRoutes()
}

func (s *restServer) installMiddlewares() {
	for _, name := range s.cfg.Middlewares {
		mw, ok := middleware.Get(name)
		if !ok {
			zap.S().Warnw("middleware not found", "name", name)
		}
		s.Use(mw)
	}
}

func (s *restServer) installRoutes() {
	v1 := s.Group("/v1")
	{
		timers := v1.Group("/timers")
		{
			tc := timer.NewController(s.serviceRouter)
			timers.POST("", tc.Create)
			timers.GET(":name", tc.Get)
			timers.GET("", tc.GetAll)
			timers.DELETE(":name", tc.Delete)
			timers.PUT(":name", tc.Update)
		}
	}

	if s.cfg.UseHealthz {
		s.GET("/healthz", func(c *gin.Context) {
			resp.WriteResponse(c, nil, map[string]string{"status": "ok"})
		})
	}
}

func (s *restServer) startInsecureServing() error {
	addr := s.cfg.InsecureServing.Addr()
	s.insecureServer = &http.Server{
		Addr:    addr,
		Handler: s,
	}
	zap.S().Infow("rest server insecure serving starts", "addr", addr)
	if err := s.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *restServer) Start() error {
	return util.BatchGoOrErr(
		s.startInsecureServing,
	)
}

func (s *restServer) Stop() error {
	zap.L().Info("rest server starts shutting down gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := s.insecureServer.Shutdown(ctx); err != nil {
		msg := "rest server failed to shut down gracefully"
		zap.S().Warnw(msg, "err", err)
		return errors.WithMessage(err, msg)
	}
	return nil
}

package restserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/josephzxy/timer_apiserver/internal/resource/v1/service"
	"github.com/josephzxy/timer_apiserver/internal/restserver/controller/v1/timer"
	"github.com/josephzxy/timer_apiserver/internal/restserver/middleware"
	resp "github.com/josephzxy/timer_apiserver/internal/restserver/response"
)

type RESTServer struct {
	cfg           Config
	serviceRouter service.ServiceRouter
	*gin.Engine
	insecureServer *http.Server
}

func New(cfg *Config, serviceRouter service.ServiceRouter) *RESTServer {
	s := &RESTServer{
		cfg:           *cfg,
		serviceRouter: serviceRouter,
		Engine:        gin.New(),
	}
	s.init()
	return s
}

func (s *RESTServer) init() {
	gin.SetMode(s.cfg.Mode)
	s.installMiddlewares()
	s.installRoutes()
}

func (s *RESTServer) installMiddlewares() {
	for _, name := range s.cfg.Middlewares {
		mw, ok := middleware.Get(name)
		if !ok {
			zap.S().Warnw("middleware not found", "name", name)
		}
		s.Use(mw)
	}
}

func (s *RESTServer) installRoutes() {
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

func (s *RESTServer) startInsecureServing() error {
	s.insecureServer = &http.Server{
		Addr:    s.cfg.InsecureServing.Addr(),
		Handler: s,
	}
	return s.insecureServer.ListenAndServe()
}

func (s *RESTServer) Start() error {
	waitDone := make(chan struct{}, 1)
	var servingErr error
	eg, ctx := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		if err := s.startInsecureServing(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.S().Errorw("error occured during rest server insecure serving", "err", err)
			servingErr = err
			return err
		}
		return nil
	})

	go func() {
		_ = eg.Wait()
		waitDone <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return errors.WithMessage(servingErr, "rest server insecure serving failed")
	case <-waitDone:
		return nil
	}
}

func (s *RESTServer) Stop() error {
	zap.L().Info("gracefully shutting down rest server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := s.insecureServer.Shutdown(ctx); err != nil {
		msg := "failed to shutdown insecure serving rest server"
		zap.S().Warnw(msg, "err", err)
		return errors.WithMessage(err, msg)
	}
	return nil
}

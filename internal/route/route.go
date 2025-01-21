package route

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"pet/configs"
	"pet/internal/route/common"
	"pet/internal/route/middleware"
	"pet/internal/service"
)

func NewGinEngine(conf *configs.Config) *gin.Engine {
	if conf.IsReleaseMode() {
		gin.SetMode(gin.ReleaseMode)
	}
	f, _ := os.OpenFile("./log/gin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
	r := gin.New()
	_ = r.SetTrustedProxies(nil)
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.Cors(true),
	)
	return r
}

type HttpEngine struct {
	log         *zap.SugaredLogger
	conf        *configs.Config
	handler     *gin.Engine
	userService *service.UserService
}

type Registrable interface {
	RegisterRoute(*gin.RouterGroup)
}

func NewHttpEngine(opt *WireOption) *HttpEngine {
	return &HttpEngine{
		log:         opt.Log,
		conf:        opt.Conf,
		handler:     opt.Handler,
		userService: opt.UserService,
	}
}

func (h *HttpEngine) registerRoute() {
	//r := h.handler.Group("/api/v1")
}

func (h *HttpEngine) Run() error {
	common.SetRespLog(h.log)
	h.registerRoute()

	srv := &http.Server{
		Addr:    h.conf.App.Addr,
		Handler: h.handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			h.log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		h.log.Fatal("Server Shutdown:", err)
		return err
	}
	h.log.Infof("server exiting")
	return nil
}

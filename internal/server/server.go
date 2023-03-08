package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"religion/config"
	postDelivery "religion/internal/domain/post/delivery"
	postRepository "religion/internal/domain/post/repository"
	postUsecase "religion/internal/domain/post/usecase"
	"religion/internal/domain/user/delivery"
	"religion/internal/domain/user/repository"
	"religion/internal/domain/user/usecase"
	"religion/internal/server/middleware"
	"religion/pkg/logger"
)

type Server struct {
	echo   *echo.Echo
	logger logger.Logger
	cfg    *config.Config
	db     *gorm.DB
	auth   *auth.Client
}

// NewServer
func NewServer(logger logger.Logger, cfg *config.Config, db *gorm.DB, auth *auth.Client) *Server {
	return &Server{logger: logger, cfg: cfg, db: db, echo: echo.New(), auth: auth}
}

func (s *Server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	validate := validator.New()
	v1 := s.echo.Group("/api/v1")
	userGroup := v1.Group("/users")
	groupGroup := v1.Group("/group")
	postGroup := v1.Group("/posts")

	// auth middleware
	midleWManager := middleware.New(s.auth)
	v1.Use(midleWManager.AuthMiddleware)

	// user
	userRepository := repository.NewUserRepository(s.db)
	userUseCase := usecase.New(userRepository, s.logger)

	userHandlers := delivery.NewUserHandlers(s.cfg, userGroup, groupGroup, userUseCase, s.logger, validate)
	userHandlers.MapUserRoutes()

	// posts
	postRepository := postRepository.NewPostRepository(s.db)
	postUseCase := postUsecase.NewPostUsecase(postRepository, userRepository, s.logger)

	postHandlers := postDelivery.NewPostHandlers(s.cfg, postGroup, postUseCase, s.logger, validate)
	postHandlers.MapPostsRoutes()

	s.echo.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.HttpServer.Port)
		s.echo.Server.ReadTimeout = time.Second * s.cfg.HttpServer.ReadTimeout
		s.echo.Server.WriteTimeout = time.Second * s.cfg.HttpServer.WriteTimeout
		// s.echo.Server.MaxHeaderBytes = maxHeaderBytes
		if err := s.echo.Start(s.cfg.HttpServer.Port); err != nil {
			s.logger.Fatalf("Error starting TLS Server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		s.logger.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		s.logger.Errorf("ctx.Done: %v", done)
	}

	s.logger.Info(ctx, "Server Exited Properly")

	if err := s.echo.Server.Shutdown(ctx); err != nil {
		return fmt.Errorf("echo.Server.Shutdown: %w", err)
	}

	s.logger.Info(ctx, "Server Exited Properly")

	return nil
}

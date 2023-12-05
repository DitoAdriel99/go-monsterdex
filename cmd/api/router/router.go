package router

import (
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/dependency"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/handlers"
	authHandl "github.com/DitoAdriel99/go-monsterdex/cmd/api/handlers/auth"
	monsterHandl "github.com/DitoAdriel99/go-monsterdex/cmd/api/handlers/monster"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/service"
	"github.com/DitoAdriel99/go-monsterdex/config"
	"github.com/DitoAdriel99/go-monsterdex/pkg/tokenizer"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func New() *echo.Echo {
	e := echo.New()
	conf := config.PopulateConfigFromEnv()
	repo := repository.NewRepo()
	dep := dependency.NewDependency(conf)

	serv := service.NewService(conf, repo, dep.Redis, dep.GcsClient, dep.Token)

	// declare handlers
	authHandlers := authHandl.NewHandlers(serv)
	monsterHandlers := monsterHandl.NewHandlers(serv)

	corsConfig := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"*"},
	}

	e.Use(middleware.CORSWithConfig(corsConfig))
	// Health Check
	e.GET("/health-check", handlers.HealthCheckHandler)

	e.POST("/api/v1/login", authHandlers.LoginHandler)
	e.POST("/api/v1/register", authHandlers.RegisterHandler)

	api := e.Group("/api/v1")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &tokenizer.Claims{},
		SigningKey: []byte(conf.JWT.Key),
	}))

	api.POST("/monster", monsterHandlers.CreateMonsterHandler, dep.RBAC.Validate("admin"))
	api.GET("/monsters", monsterHandlers.GetMonstersHandler, dep.RBAC.Validate("admin", "user"))
	api.PUT("/monster/:id", monsterHandlers.UpdateMonsterHandler, dep.RBAC.Validate("admin"))
	api.PUT("/monster/status/:id", monsterHandlers.SetStatusMonsterHandler, dep.RBAC.Validate("admin"))
	api.POST("/monster/catch/:id", monsterHandlers.CatchMonsterHandler, dep.RBAC.Validate("user"))
	// Serve Swagger documentation
	e.GET("/*", echoSwagger.WrapHandler)

	return e
}

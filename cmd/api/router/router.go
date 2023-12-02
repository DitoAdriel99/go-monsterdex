package router

import (
	"github.com/DitoAdriel99/go-monsterdex/bootstrap"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/handlers"
	authHandl "github.com/DitoAdriel99/go-monsterdex/cmd/api/handlers/auth"
	monsterHandl "github.com/DitoAdriel99/go-monsterdex/cmd/api/handlers/monster"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/service"

	// _ "github.com/DitoAdriel99/go-monsterdex/docs/echosimple"
	midd "github.com/DitoAdriel99/go-monsterdex/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func New() *echo.Echo {
	e := echo.New()
	repo := repository.NewRepo()
	rdb := bootstrap.NewRedisClient()

	serv := service.NewService(repo, rdb)

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
		Claims:     &entity.Claims{},
		SigningKey: entity.JWTKEY,
	}))

	api.POST("/monster", monsterHandlers.CreateMonsterHandler, midd.RBAC("admin"))
	api.GET("/monsters", monsterHandlers.GetMonstersHandler, midd.RBAC("admin", "user"))
	api.PUT("/monster/:id", monsterHandlers.UpdateMonsterHandler, midd.RBAC("admin", "user"))
	api.PUT("/monster/status/:id", monsterHandlers.SetStatusMonsterHandler, midd.RBAC("admin"))
	api.POST("/monster/catch/:id", monsterHandlers.CatchMonsterHandler, midd.RBAC("user"))
	// Serve Swagger documentation
	e.GET("/*", echoSwagger.WrapHandler)

	return e
}

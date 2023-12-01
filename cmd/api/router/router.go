package router

import (
	"context"

	"github.com/DitoAdriel99/go-monsterdex/cmd/api/entity"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/handlers"
	authHandl "github.com/DitoAdriel99/go-monsterdex/cmd/api/handlers/auth"
	monsterHandl "github.com/DitoAdriel99/go-monsterdex/cmd/api/handlers/monster"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/repository"
	"github.com/DitoAdriel99/go-monsterdex/cmd/api/service"
	_ "github.com/DitoAdriel99/go-monsterdex/docs/echosimple"
	midd "github.com/DitoAdriel99/go-monsterdex/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:2000
// @BasePath /
// @schemes http

func New() *echo.Echo {
	e := echo.New()
	repo := repository.NewRepo()
	// sdk := sdk.NewSDK(os.Getenv("DOG_URL"))
	// rdb := bootstrap.NewRedisClient(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PASSWORD"))

	serv := service.NewService(repo)

	context.Background()

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
	api.POST("/monster/catch/:id", monsterHandlers.CatchMonsterHandler, midd.RBAC("admin", "user"))
	// Serve Swagger documentation
	e.GET("/*", echoSwagger.WrapHandler)

	return e
}

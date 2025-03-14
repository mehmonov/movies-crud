package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"github.com/mehmonov/movies-crud/config"
	_ "github.com/mehmonov/movies-crud/docs" 
	"github.com/mehmonov/movies-crud/internal/api/routes"
	"github.com/mehmonov/movies-crud/internal/db"
	"github.com/mehmonov/movies-crud/internal/services"
)

// @title           Movies CRUD API
// @version         1.0
// @description     This is a sample movies CRUD API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  mehmonov.husniddin1@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			db.NewDatabase,
			services.NewMovieService,
			services.NewUserService,
			routes.NewRouter,
		),
		fx.Invoke(startServer),
	)

	app.Run()
}

func startServer(router *gin.Engine, cfg *config.Config) {
	log.Printf("Starting server on :%s", cfg.ServerPort)
	router.Run(":" + cfg.ServerPort)
}

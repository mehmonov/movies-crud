package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/mehmonov/movies-crud/config"
	"github.com/mehmonov/movies-crud/internal/api/handlers"
	"github.com/mehmonov/movies-crud/internal/api/middleware"
	"github.com/mehmonov/movies-crud/internal/services"
	"github.com/mehmonov/movies-crud/pkg/auth"
)

func NewRouter(
	cfg *config.Config,
	movieService *services.MovieService,
	userService *services.UserService,
) *gin.Engine {
	router := gin.Default()

	jwtService := auth.NewJWTService(cfg.JWTSecret)

	movieHandler := handlers.NewMovieHandler(movieService)
	userHandler := handlers.NewUserHandler(userService, jwtService)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
		}

		movies := api.Group("/movies")
		{
			movies.GET("", movieHandler.GetAllMovies)     // Public endpoint
			movies.GET("/:id", movieHandler.GetMovieByID) // Public endpoint

			// Protected movie routes (with auth middleware)
			movies.Use(middleware.AuthMiddleware(jwtService))
			{
				movies.POST("", movieHandler.CreateMovie)
				movies.PUT("/:id", movieHandler.UpdateMovie)
				movies.DELETE("/:id", movieHandler.DeleteMovie)
			}
		}
	}

	return router
}

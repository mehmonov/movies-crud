package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/mehmonov/movies-crud/config"
	"github.com/mehmonov/movies-crud/internal/api/controllers"
	"github.com/mehmonov/movies-crud/internal/api/handlers"
	"github.com/mehmonov/movies-crud/internal/api/middleware"
	"github.com/mehmonov/movies-crud/internal/services"
	"github.com/mehmonov/movies-crud/pkg/auth"
	"github.com/mehmonov/movies-crud/pkg/filestore"
)

func NewRouter(
	cfg *config.Config,
	movieService *services.MovieService,
	userService *services.UserService,
) *gin.Engine {
	router := gin.Default()

	jwtService := auth.NewJWTService(cfg.JWTSecret, cfg.JWTSecret)
	fileStore := filestore.NewFileStore("./uploads") // or cfg.UploadPath

	movieController := controllers.NewMovieController(movieService, fileStore)
	movieHandler := handlers.NewMovieHandler(movieController)
	userHandler := handlers.NewUserHandler(userService, jwtService)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
			auth.POST("/refresh", userHandler.RefreshToken)
		}

		movies := api.Group("/movies")
		{
			movies.GET("", movieHandler.GetAllMovies)                 // Public endpoint
			movies.GET("/:id", movieHandler.GetMovieByID)             // Public endpoint
			movies.GET("/:id/media", movieHandler.GetMovieMediaFiles) // Yangi endpoint

			// Protected movie routes (with auth middleware)
			movies.Use(middleware.AuthMiddleware(jwtService))
			{
				movies.POST("", movieHandler.CreateMovie)
				movies.PUT("/:id", movieHandler.UpdateMovie)
				movies.DELETE("/:id", movieHandler.DeleteMovie)
				movies.POST("/:id/file", movieHandler.UploadMovieFile)
				movies.GET("/:id/files/:fileId", movieHandler.GetMovieFile)
			}
		}
	}

	return router
}

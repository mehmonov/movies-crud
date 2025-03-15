package controllers

import (
	"mime/multipart"

	"github.com/mehmonov/movies-crud/internal/models"
	"github.com/mehmonov/movies-crud/internal/services"
	"github.com/mehmonov/movies-crud/pkg/filestore"
)

type MovieController struct {
	movieService *services.MovieService
	fileStore    *filestore.FileStore
}

func NewMovieController(movieService *services.MovieService, fileStore *filestore.FileStore) *MovieController {
	return &MovieController{
		movieService: movieService,
		fileStore:    fileStore,
	}
}

func (c *MovieController) GetAllMovies() ([]models.Movie, error) {
	return c.movieService.GetAllMovies()
}

func (c *MovieController) GetMovieByID(id uint) (*models.Movie, error) {
	return c.movieService.GetMovieByID(id)
}

func (c *MovieController) CreateMovie(req *models.CreateMovieRequest) (*models.Movie, error) {
	return c.movieService.CreateMovie(req)
}

func (c *MovieController) UpdateMovie(id uint, req *models.UpdateMovieRequest) error {
	return c.movieService.UpdateMovie(id, req)
}

func (c *MovieController) DeleteMovie(id uint) error {
	return c.movieService.DeleteMovie(id)
}

func (c *MovieController) GetMovieMediaFiles(movieID uint) ([]models.MovieMedia, error) {
	return c.movieService.GetMovieMediaFiles(movieID)
}

func (c *MovieController) GetMovieMediaFilesByType(movieID uint, mediaType string) ([]models.MovieMedia, error) {
	return c.movieService.GetMovieMediaFilesByType(movieID, mediaType)
}

func (c *MovieController) UploadMovieFile(movieID uint, file *multipart.FileHeader) (*models.MovieFile, error) {
	return c.movieService.UploadMovieFile(movieID, file, c.fileStore)
}

func (c *MovieController) GetMovieFileContent(movieID, fileID uint) (*models.MovieFile, error) {
	return c.movieService.GetMovieFileContent(movieID, fileID)
}

package controllers

import (
	"github.com/mehmonov/movies-crud/internal/models"
	"github.com/mehmonov/movies-crud/internal/services"
)

type MovieController struct {
	movieService *services.MovieService
}

func NewMovieController(movieService *services.MovieService) *MovieController {
	return &MovieController{
		movieService: movieService,
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
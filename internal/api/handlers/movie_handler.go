package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mehmonov/movies-crud/internal/api/controllers"
	"github.com/mehmonov/movies-crud/internal/models"
)

type MovieHandler struct {
	movieController *controllers.MovieController
}

func NewMovieHandler(movieController *controllers.MovieController) *MovieHandler {
	return &MovieHandler{
		movieController: movieController,
	}
}

// @Summary Get all movies
// @Description Get a list of all movies
// @Tags movies
// @Accept json
// @Produce json
// @Success 200 {array} models.Movie
// @Router /movies [get]
func (h *MovieHandler) GetAllMovies(c *gin.Context) {
	movies, err := h.movieController.GetAllMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movies"})
		return
	}

	c.JSON(http.StatusOK, movies)
}

// @Summary Get a movie by ID
// @Description Get details of a specific movie
// @Tags movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} models.Movie
// @Failure 404 {object} object
// @Router /movies/{id} [get]
func (h *MovieHandler) GetMovieByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	movie, err := h.movieController.GetMovieByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movie"})
		return
	}

	if movie == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

// @Summary Create a new movie
// @Description Add a new movie to the database
// @Tags movies
// @Accept json
// @Produce json
// @Param movie body models.CreateMovieRequest true "Movie information"
// @Success 201 {object} models.Movie
// @Failure 400 {object} object
// @Router /movies [post]
func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var req models.CreateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie, err := h.movieController.CreateMovie(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create movie"})
		return
	}

	c.JSON(http.StatusCreated, movie)
}

// @Summary Update a movie
// @Description Update an existing movie's details
// @Tags movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Param movie body models.UpdateMovieRequest true "Movie information"
// @Success 200 {object} models.Movie
// @Failure 400 {object} object
// @Failure 404 {object} object
// @Router /movies/{id} [put]
func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var req models.UpdateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.movieController.UpdateMovie(uint(id), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		return
	}

	movie, _ := h.movieController.GetMovieByID(uint(id))
	c.JSON(http.StatusOK, movie)
}

// @Summary Delete a movie
// @Description Delete a movie from the database
// @Tags movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 204
// @Failure 400 {object} object
// @Router /movies/{id} [delete]
func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.movieController.DeleteMovie(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete movie"})
		return
	}

	c.Status(http.StatusNoContent)
}

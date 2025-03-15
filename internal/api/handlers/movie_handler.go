package handlers

import (
	"fmt"
	"net/http"
	"os"
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

// @Summary Get movie media files
// @Description Get all media files for a specific movie
// @Tags movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Param type query string false "Media type (poster, backdrop, trailer)"
// @Success 200 {array} models.MovieMedia
// @Failure 400 {object} object
// @Failure 404 {object} object
// @Router /movies/{id}/media [get]
func (h *MovieHandler) GetMovieMediaFiles(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Check if movie exists
	movie, err := h.movieController.GetMovieByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movie"})
		return
	}
	if movie == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	// Check if type parameter is provided
	mediaType := c.Query("type")
	var mediaFiles []models.MovieMedia

	if mediaType != "" {
		// Validate media type
		validTypes := map[string]bool{"poster": true, "backdrop": true, "trailer": true}
		if !validTypes[mediaType] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid media type"})
			return
		}

		mediaFiles, err = h.movieController.GetMovieMediaFilesByType(uint(id), mediaType)
	} else {
		mediaFiles, err = h.movieController.GetMovieMediaFiles(uint(id))
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve media files"})
		return
	}

	c.JSON(http.StatusOK, mediaFiles)
}

// @Summary Upload movie file
// @Description Upload a video file for a movie
// @Tags movies
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Movie ID"
// @Param file formData file true "Movie file"
// @Security BearerAuth
// @Success 201 {object} models.MovieFileResponse
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 404 {object} object
// @Router /movies/{id}/file [post]
func (h *MovieHandler) UploadMovieFile(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	movieFile, err := h.movieController.UploadMovieFile(uint(id), file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := models.MovieFileResponse{
		ID:          movieFile.ID,
		MovieID:     movieFile.MovieID,
		FileName:    movieFile.FileName,
		FileSize:    movieFile.FileSize,
		ContentType: movieFile.ContentType,
		CreatedAt:   movieFile.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// @Summary Download movie file
// @Description Download a movie file by ID
// @Tags movies
// @Produce octet-stream
// @Param id path int true "Movie ID"
// @Param fileId path int true "File ID"
// @Security BearerAuth
// @Success 200 {file} binary
// @Failure 400 {object} object
// @Failure 401 {object} object
// @Failure 404 {object} object
// @Router /movies/{id}/files/{fileId} [get]
func (h *MovieHandler) GetMovieFile(c *gin.Context) {
	// Get movie ID
	movieID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID format"})
		return
	}

	// Get file ID
	fileID, err := strconv.ParseUint(c.Param("fileId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID format"})
		return
	}

	// Get file info from database
	movieFile, err := h.movieController.GetMovieFileContent(uint(movieID), uint(fileID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve file info"})
		return
	}
	if movieFile == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Check if file exists
	if _, err := os.Stat(movieFile.FilePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found on disk"})
		return
	}

	// Set headers for file download
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", movieFile.FileName))
	c.Header("Content-Type", movieFile.ContentType)

	// Serve the file
	c.File(movieFile.FilePath)
}

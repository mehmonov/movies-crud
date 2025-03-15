package services

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"github.com/mehmonov/movies-crud/internal/models"
)

type MovieService struct {
	db *gorm.DB
}

func NewMovieService(db *gorm.DB) *MovieService {
	return &MovieService{db: db}
}

func (s *MovieService) GetAllMovies() ([]models.Movie, error) {
	var movies []models.Movie
	result := s.db.Preload("MediaFiles").Preload("Metadata").Find(&movies)
	return movies, result.Error
}

func (s *MovieService) GetMovieByID(id uint) (*models.Movie, error) {
	var movie models.Movie
	result := s.db.Preload("MediaFiles").Preload("Metadata").First(&movie, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &movie, nil
}

func (s *MovieService) CreateMovie(req *models.CreateMovieRequest) (*models.Movie, error) {
	// Validate duration
	if req.Duration > 1000 {
		return nil, errors.New("duration cannot exceed 1000 minutes")
	}

	movie := models.Movie{
		Title:    req.Title,
		Director: req.Director,
		Year:     req.Year,
		Plot:     req.Plot,
		Genre:    req.Genre,
		Rating:   req.Rating,
		Duration: req.Duration,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Create movie
		if err := tx.Create(&movie).Error; err != nil {
			if strings.Contains(err.Error(), "numeric field overflow") {
				return errors.New("invalid numeric value: please check duration, year and rating fields")
			}
			return err
		}

		// Create media files
		for _, mediaReq := range req.MediaFiles {
			media := models.MovieMedia{
				MovieID: movie.ID,
				Type:    mediaReq.Type,
				URL:     mediaReq.URL,
				IsMain:  mediaReq.IsMain,
			}
			if err := tx.Create(&media).Error; err != nil {
				return err
			}
		}

		// Create metadata
		metadata := models.MovieMetadata{
			MovieID:  movie.ID,
			Language: req.Metadata.Language,
			Country:  req.Metadata.Country,
			Awards:   req.Metadata.Awards,
			Cast:     req.Metadata.Cast,
		}
		if err := tx.Create(&metadata).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Load relationships
	s.db.Preload("MediaFiles").Preload("Metadata").First(&movie, movie.ID)

	return &movie, nil
}

func (s *MovieService) UpdateMovie(id uint, req *models.UpdateMovieRequest) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var movie models.Movie
		if err := tx.First(&movie, id).Error; err != nil {
			return err
		}

		// Update movie fields
		if req.Title != "" {
			movie.Title = req.Title
		}
		if req.Director != "" {
			movie.Director = req.Director
		}
		if req.Year != 0 {
			movie.Year = req.Year
		}
		if req.Plot != "" {
			movie.Plot = req.Plot
		}
		if req.Genre != "" {
			movie.Genre = req.Genre
		}
		if req.Rating != 0 {
			movie.Rating = req.Rating
		}
		if req.Duration != 0 {
			movie.Duration = req.Duration
		}

		if err := tx.Save(&movie).Error; err != nil {
			return err
		}

		// Update media files if provided
		if len(req.MediaFiles) > 0 {
			// Delete existing media files
			if err := tx.Where("movie_id = ?", movie.ID).Delete(&models.MovieMedia{}).Error; err != nil {
				return err
			}

			// Create new media files
			for _, mediaReq := range req.MediaFiles {
				media := models.MovieMedia{
					MovieID: movie.ID,
					Type:    mediaReq.Type,
					URL:     mediaReq.URL,
					IsMain:  mediaReq.IsMain,
				}
				if err := tx.Create(&media).Error; err != nil {
					return err
				}
			}
		}

		// Update metadata if provided
		if req.Metadata != (models.MovieMetadataRequest{}) {
			var metadata models.MovieMetadata
			tx.FirstOrCreate(&metadata, models.MovieMetadata{MovieID: movie.ID})

			metadata.Language = req.Metadata.Language
			metadata.Country = req.Metadata.Country
			metadata.Awards = req.Metadata.Awards
			metadata.Cast = req.Metadata.Cast

			if err := tx.Save(&metadata).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *MovieService) DeleteMovie(id uint) error {
	return s.db.Delete(&models.Movie{}, id).Error
}

// GetMovieMediaFiles returns all media files for a specific movie
func (s *MovieService) GetMovieMediaFiles(movieID uint) ([]models.MovieMedia, error) {
	var mediaFiles []models.MovieMedia
	result := s.db.Where("movie_id = ?", movieID).Find(&mediaFiles)
	if result.Error != nil {
		return nil, result.Error
	}
	return mediaFiles, nil
}

// GetMovieMediaFilesByType returns media files of specific type for a movie
func (s *MovieService) GetMovieMediaFilesByType(movieID uint, mediaType string) ([]models.MovieMedia, error) {
	var mediaFiles []models.MovieMedia
	result := s.db.Where("movie_id = ? AND type = ?", movieID, mediaType).Find(&mediaFiles)
	if result.Error != nil {
		return nil, result.Error
	}
	return mediaFiles, nil
}

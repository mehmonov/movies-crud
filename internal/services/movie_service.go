package services

import (
    "errors"
    
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
    result := s.db.Find(&movies)
    return movies, result.Error
}

func (s *MovieService) GetMovieByID(id uint) (*models.Movie, error) {
    var movie models.Movie
    result := s.db.First(&movie, id)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, nil 
        }
        return nil, result.Error
    }
    return &movie, nil
}

func (s *MovieService) CreateMovie(req *models.CreateMovieRequest) (*models.Movie, error) {
    movie := models.Movie{
        Title:    req.Title,
        Director: req.Director,
        Year:     req.Year,
        Plot:     req.Plot,
    }
    
    err := s.db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Create(&movie).Error; err != nil {
            return err
        }
        return nil
    })
    
    if err != nil {
        return nil, err
    }
    
    return &movie, nil
}

func (s *MovieService) UpdateMovie(id uint, req *models.UpdateMovieRequest) error {
    var movie models.Movie
    if err := s.db.First(&movie, id).Error; err != nil {
        return err
    }
    
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
    
    return s.db.Save(&movie).Error
}

func (s *MovieService) DeleteMovie(id uint) error {
    return s.db.Delete(&models.Movie{}, id).Error
}
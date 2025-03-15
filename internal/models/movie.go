package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID         uint           `json:"id" gorm:"primarykey"`
	Title      string         `json:"title" gorm:"size:100;not null"`
	Director   string         `json:"director" gorm:"size:100"`
	Year       int            `json:"year" gorm:"not null;check:year >= 1800 AND year <= 2100"`
	Plot       string         `json:"plot" gorm:"type:text"`
	Genre      string         `json:"genre" gorm:"size:50"`
	Rating     float32        `json:"rating" gorm:"type:decimal(3,1);check:rating >= 0 AND rating <= 10"`
	Duration   int            `json:"duration" gorm:"check:duration >= 0 AND duration <= 1000"` // Minutes
	MediaFiles []MovieMedia   `json:"media_files,omitempty" gorm:"foreignKey:MovieID"`
	Metadata   MovieMetadata  `json:"metadata,omitempty" gorm:"foreignKey:MovieID"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

type MovieMedia struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	MovieID   uint           `json:"movie_id"`
	Type      string         `json:"type" gorm:"size:20"` // poster, backdrop, trailer
	URL       string         `json:"url" gorm:"type:text"`
	IsMain    bool           `json:"is_main"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type MovieMetadata struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	MovieID   uint           `json:"movie_id"`
	Language  string         `json:"language" gorm:"size:10"`
	Country   string         `json:"country" gorm:"size:50"`
	Awards    string         `json:"awards" gorm:"type:text"`
	Cast      string         `json:"cast" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateMovieRequest struct {
	Title      string               `json:"title" binding:"required,max=100"`
	Director   string               `json:"director" binding:"required,max=100"`
	Year       int                  `json:"year" binding:"required,min=1800,max=2100"`
	Plot       string               `json:"plot"`
	Genre      string               `json:"genre" binding:"max=50"`
	Rating     float32              `json:"rating" binding:"min=0,max=10"`
	Duration   int                  `json:"duration" binding:"min=0,max=1000"` // Maximum 1000 minutes (16.6 hours)
	MediaFiles []MovieMediaRequest  `json:"media_files"`
	Metadata   MovieMetadataRequest `json:"metadata"`
}

type MovieMediaRequest struct {
	Type   string `json:"type" binding:"required,oneof=poster backdrop trailer"`
	URL    string `json:"url" binding:"required,url"`
	IsMain bool   `json:"is_main"`
}

type MovieMetadataRequest struct {
	Language string `json:"language"`
	Country  string `json:"country"`
	Awards   string `json:"awards"`
	Cast     string `json:"cast"`
}

type UpdateMovieRequest struct {
	Title      string               `json:"title" binding:"omitempty,max=100"`
	Director   string               `json:"director" binding:"omitempty,max=100"`
	Year       int                  `json:"year" binding:"omitempty,min=1800,max=2100"`
	Plot       string               `json:"plot"`
	Genre      string               `json:"genre" binding:"omitempty,max=50"`
	Rating     float32              `json:"rating" binding:"omitempty,min=0,max=10"`
	Duration   int                  `json:"duration" binding:"omitempty,min=0,max=1000"`
	MediaFiles []MovieMediaRequest  `json:"media_files"`
	Metadata   MovieMetadataRequest `json:"metadata"`
}

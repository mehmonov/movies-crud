package models

import (
    "time"
    
    "gorm.io/gorm"
)

type Movie struct {
    ID        uint           `json:"id" gorm:"primarykey"`
    Title     string         `json:"title" gorm:"size:100;not null"`
    Director  string         `json:"director" gorm:"size:100"`
    Year      int            `json:"year" gorm:"not null"`
    Plot      string         `json:"plot" gorm:"type:text"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateMovieRequest struct {
    Title    string `json:"title" binding:"required"`
    Director string `json:"director" binding:"required"`
    Year     int    `json:"year" binding:"required,min=1800,max=2100"`
    Plot     string `json:"plot"`
}

type UpdateMovieRequest struct {
    Title    string `json:"title"`
    Director string `json:"director"`
    Year     int    `json:"year" binding:"omitempty,min=1800,max=2100"`
    Plot     string `json:"plot"`
}
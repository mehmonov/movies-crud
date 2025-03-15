package models

import (
	"time"

	"gorm.io/gorm"
)

type MovieFile struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	MovieID     uint           `json:"movie_id"`
	FileName    string         `json:"file_name" gorm:"size:255"`
	FileSize    int64          `json:"file_size"`
	ContentType string         `json:"content_type" gorm:"size:100"`
	FilePath    string         `json:"-" gorm:"size:500"` // Path in storage, hidden from JSON
	FileHash    string         `json:"-" gorm:"size:64"`  // SHA-256 hash
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type MovieFileResponse struct {
	ID          uint      `json:"id"`
	MovieID     uint      `json:"movie_id"`
	FileName    string    `json:"file_name"`
	FileSize    int64     `json:"file_size"`
	ContentType string    `json:"content_type"`
	CreatedAt   time.Time `json:"created_at"`
}

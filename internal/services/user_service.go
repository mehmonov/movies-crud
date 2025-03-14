package services

import (
    "errors"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"

    "github.com/mehmonov/movies-crud/internal/models"
)

type UserService struct {
    db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
    return &UserService{db: db}
}

func (s *UserService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {


    var existingUser models.User
    if err := s.db.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
        return nil, errors.New("username already exists")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := models.User{
        Username: req.Username,
        Password: string(hashedPassword),
    }

    if err := s.db.Create(&user).Error; err != nil {
        return nil, err
    }

    // Don't return the password
    user.Password = ""
    return &user, nil
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
    var user models.User
    if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
} 
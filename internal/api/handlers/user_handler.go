package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/mehmonov/movies-crud/internal/models"
	"github.com/mehmonov/movies-crud/internal/services"
	"github.com/mehmonov/movies-crud/pkg/auth"
)

type UserHandler struct {
	userService *services.UserService
	jwtService  *auth.JWTService
}

func NewUserHandler(userService *services.UserService, jwtService *auth.JWTService) *UserHandler {
	return &UserHandler{
		userService: userService,
		jwtService:  jwtService,
	}
}

// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User registration details"
// @Success 201 {object} models.User
// @Failure 400 {object} object
// @Router /auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// @Summary Login user
// @Description Login with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Login credentials"
// @Success 200 {object} models.AuthResponse
// @Failure 401 {object} object
// @Router /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.GetUserByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := h.jwtService.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Create a safe user response without password 
	safeUser := &models.User{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	c.JSON(http.StatusOK, models.AuthResponse{
		Token: token,
		User:  safeUser,
	})
}

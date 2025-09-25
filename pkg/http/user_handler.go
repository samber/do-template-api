package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/samber/do-template-api/pkg/repositories"
	"github.com/samber/do/v2"
)

// UserHandler handles HTTP requests for user operations
// This demonstrates how to organize handlers in a separate struct.
type UserHandler struct {
	userRepo repositories.UserRepository `do:""`
	logger   zerolog.Logger              `do:""`
}

// NewUserHandler creates a new UserHandler with dependency injection
// This demonstrates how to initialize handlers with dependencies.
func NewUserHandler(injector do.Injector) (*UserHandler, error) {
	return do.MustInvokeStruct[*UserHandler](injector), nil
}

// createUser handles user creation requests
// This demonstrates how to implement CREATE endpoint with dependency injection.
func (h *UserHandler) createUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &repositories.User{
		Name:  req.Name,
		Email: req.Email,
	}

	createdUser, err := h.userRepo.CreateUser(c.Request.Context(), user)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to create user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, UserResponse{
		ID:        createdUser.ID,
		Name:      createdUser.Name,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	})
}

// getUser handles user retrieval requests
// This demonstrates how to implement READ endpoint with dependency injection.
func (h *UserHandler) getUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userRepo.GetUserByID(c.Request.Context(), id)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to get user")
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}

// listUsers handles user listing requests
// This demonstrates how to implement LIST endpoint with dependency injection.
func (h *UserHandler) listUsers(c *gin.Context) {
	limit := 20
	offset := 0

	if l, err := strconv.Atoi(c.DefaultQuery("limit", "20")); err == nil && l > 0 {
		limit = l
	}

	if o, err := strconv.Atoi(c.DefaultQuery("offset", "0")); err == nil && o >= 0 {
		offset = o
	}

	users, err := h.userRepo.ListUsers(c.Request.Context(), limit, offset)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to list users")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list users"})
		return
	}

	response := make([]UserResponse, len(users))
	for i, user := range users {
		response[i] = UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"users":  response,
		"limit":  limit,
		"offset": offset,
	})
}

// updateUser handles user update requests
// This demonstrates how to implement UPDATE endpoint with dependency injection.
func (h *UserHandler) updateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &repositories.User{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	}

	updatedUser, err := h.userRepo.UpdateUser(c.Request.Context(), user)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to update user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:        updatedUser.ID,
		Name:      updatedUser.Name,
		Email:     updatedUser.Email,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	})
}

// deleteUser handles user deletion requests
// This demonstrates how to implement DELETE endpoint with dependency injection.
func (h *UserHandler) deleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = h.userRepo.DeleteUser(c.Request.Context(), id)
	if err != nil {
		h.logger.Error().Err(err).Msg("Failed to delete user")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

package handler

import (
	"net/http"
	"strconv"

	"chekisvc/internal/domain/entity"
	"chekisvc/internal/usecase"
	"chekisvc/pkg/utils"

	"github.com/gin-gonic/gin"
)

// UserHandler represents the handler for user-related HTTP requests
type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

// NewUserHandler creates a new user handler instance
func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags Users
// @Accept json
// @Produce json
// @Param user body entity.UserRegisterRequest true "User details"
// @Success 201 {object} entity.UserRegisterRequest
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /users [post]
// CreateUser handles POST /users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user entity.UserRegisterRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		// c.JSON(http.StatusBadRequest, utils.ErrorResponse())
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	checkEmail := utils.ValidateEmail(user.Email)
	if !checkEmail {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid email format", "email tidak valid")
		return
	}

	if err := h.userUsecase.CreateUser(c.Request.Context(), &user); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "email sudah terdaftar", err.Error())
		return
	}
	userRes := entity.UserResponse{
		Name:  user.Name,
		Email: user.Email,
	}
	// c.JSON(http.StatusCreated, user)
	// disini kita bisa mengembalikan response yang lebih lengkap
	utils.SuccessResponse(c, "User created successfully", userRes)
}

// GetUser godoc
// @Summary Get user by ID
// @Description Retrieve a user by their ID
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} entity.UserResponse
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Security BearerAuth
// @Router /users/{id} [get]
// GetUser handles GET /users/:id
func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid user ID", err.Error())
		return
	}

	user, err := h.userUsecase.GetUserByID(c.Request.Context(), uint(id))
	if err != nil {
		// c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		utils.ErrorResponse(c, http.StatusNotFound, "user not found", err.Error())
		return
	}

	// c.JSON(http.StatusOK, user)
	utils.SuccessResponse(c, "User retrieved successfully", user)
}

// UpdateUser godoc
// @Summary Update user by ID
// @Description Update an existing user by their ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body entity.UpdateUserRequest true "User details"
// @Success 200 {object} entity.User
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Security BearerAuth
// @Router /users/{id} [put]
// UpdateUser handles PUT /users/:id
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid user ID", err.Error())
		return
	}

	var user entity.UpdateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	userid, err := h.userUsecase.GetUserByID(c.Request.Context(), uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "user not found", err.Error())
		return
	}
	if userid == nil {
		utils.ErrorResponse(c, http.StatusNotFound, "user not found", "user with this ID does not exist")
		return
	}
	// user.ID = uint(id)
	if err := h.userUsecase.UpdateUser(c.Request.Context(), &user, uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to update user", err.Error())
		return
	}

	// c.JSON(http.StatusOK, user)
	utils.SuccessResponse(c, "User updated successfully", user)
}

// DeleteUser godoc
// @Summary Delete user by ID
// @Description Delete a user by their ID
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} entity.UserResponse
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Security BearerAuth
// @Router /users/{id} [delete]
// DeleteUser handles DELETE /users/:id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid user ID", err.Error())
		return
	}

	if err := h.userUsecase.DeleteUser(c.Request.Context(), uint(id)); err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		utils.ErrorResponse(c, http.StatusBadRequest, "gagal delete user", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

// ListUsers godoc
// @Summary List users with pagination
// @Description Retrieve a list of users with pagination
// @Tags Users
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} []entity.UserResponse
// @Failure 400 {object} utils.Response
// @Security BearerAuth
// @Router /users [get]
// ListUsers handles GET /users
func (h *UserHandler) ListUsers(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		utils.ErrorResponse(c, http.StatusBadRequest, "parameter tidak valid", err.Error())
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
		utils.ErrorResponse(c, http.StatusBadRequest, "parameter tidak valid", err.Error())
		return
	}

	users, err := h.userUsecase.ListUsers(c.Request.Context(), limit, offset)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		utils.ErrorResponse(c, http.StatusInternalServerError, "gagal mengambil daftar pengguna", err.Error())
		return
	}

	// c.JSON(http.StatusOK, gin.H{"users": users})
	utils.SuccessResponse(c, "Users retrieved successfully", users)
}

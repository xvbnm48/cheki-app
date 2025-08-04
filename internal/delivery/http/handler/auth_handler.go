package handler

import (
	"chekisvc/internal/domain/entity"
	"chekisvc/internal/usecase"
	"chekisvc/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewAuthHandler(userUsecase *usecase.UserUsecase) *AuthHandler {
	return &AuthHandler{userUsecase: userUsecase}
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body entity.UserLoginRequest true "Login credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req entity.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	user, err := h.userUsecase.AuthenticateUser(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "invalid email or password", err.Error())
		return
	}
	accessToken, refreshToken, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to generate token", err.Error())
		return
	}
	utils.LoginResponse(c, accessToken, refreshToken)
}

// RefreshToken godoc
// @Summary Refresh JWT token
// @Description Get new access token using refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param refresh_token body map[string]string true "Refresh token"
// @Success 200 {object} map[string]string
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /refresh-token [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}
	token, err := utils.ValidateRefreshJWT(req.RefreshToken)
	if err != nil || !token.Valid {
		utils.ErrorResponse(c, http.StatusUnauthorized, "invalid refresh token", err.Error())
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		utils.ErrorResponse(c, http.StatusUnauthorized, "invalid refresh token claims", "invalid claims")
		return
	}
	userID, _ := claims["user_id"].(float64)
	email, _ := claims["email"].(string)
	accessToken, _, err := utils.GenerateJWT(uint(userID), email)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to generate token", err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

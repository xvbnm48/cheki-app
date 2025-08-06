package handler

import (
	"chekisvc/internal/domain/entity"
	"chekisvc/internal/usecase/product"
	"chekisvc/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase product.ProductUsecase
}

func NewProductHandler(productUsecase product.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the input payload
// @Tags Products
// @Accept  json
// @Produce  json
// @Param product body entity.CreateProductRequest true "Create Product"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /products [post]
// @Security BearerAuth
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req entity.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	if err := h.productUsecase.CreateProduct(c.Request.Context(), &req); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create product", err.Error())
		return
	}

	utils.SuccessResponse(c, "Product created successfully", req)
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Get all products with pagination
// @Tags Products
// @Accept  json
// @Produce  json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /products [get]
// @Security BearerAuth
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid parameter", err.Error())
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset parameter"})
		utils.ErrorResponse(c, http.StatusBadRequest, "invlid parameter", err.Error())
		return
	}
	products, err := h.productUsecase.ListProducts(c.Request.Context(), limit, offset)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to fetch products", err.Error())
		return
	}

	utils.SuccessResponse(c, "Products fetched successfully", products)
}

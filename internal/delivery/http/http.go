package http

import (
	"github.com/kodinggo/product-service-gb1/internal/model"
	"github.com/labstack/echo/v4"
)

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
}

type httpHandler struct {
	categoryUsecase model.CategoryUsecase
	productUsecase  model.ProductUsecase
}

func NewHTTPHandler() *httpHandler {
	return new(httpHandler)
}

func (h *httpHandler) RegisterCategoryUsecase(cu model.CategoryUsecase) {
	h.categoryUsecase = cu
}

func (h *httpHandler) RegisterProductUsecase(pu model.ProductUsecase) {
	h.productUsecase = pu
}

func (h *httpHandler) Routes(route *echo.Echo) {
	v1 := route.Group("/api/v1")

	v1.GET("/categories", h.findCategories)
	v1.GET("/categories/:id", h.findCategoryByID)

	v1.GET("/products", h.findAllProducts)
	v1.GET("/products/:id", h.findProductByID)

	// private routes goes here
	routes := v1.Group("")
	// routes.Use(middleware.Logger())
	// routes.Use(middleware.CORS())
	// routes.Use(echojwt.WithConfig(utils.JWtConfig()))

	routes.POST("/categories", h.createCategory)
	routes.PUT("/categories/:id", h.updateCategory)
	routes.DELETE("/categories/:id", h.deleteCategory)

	routes.POST("/products", h.createProduct)
	routes.PUT("/products/:id", h.updateProduct)
	routes.DELETE("/products/:id", h.deleteProduct)
}

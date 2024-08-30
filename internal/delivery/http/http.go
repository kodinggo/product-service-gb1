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
}

func NewHTTPHandler() *httpHandler {
	return new(httpHandler)
}

func (h *httpHandler) RegisterCategoryUsecase(cu model.CategoryUsecase) {
	h.categoryUsecase = cu
}

func (h *httpHandler) Routes(route *echo.Echo) {
	v1 := route.Group("/api/v1")

	v1.GET("/categories", h.findCategories)
	v1.GET("/categories/:id", h.findCategoryByID)

	// private routes goes here
	routes := v1.Group("")
	// routes.Use(middleware.Logger())
	// routes.Use(middleware.CORS())
	// routes.Use(echojwt.WithConfig(utils.JWtConfig()))

	routes.POST("/categories", h.createCategory)
	routes.PUT("/categories/:id", h.updateCategory)
	routes.DELETE("/categories/:id", h.deleteCategory)
}

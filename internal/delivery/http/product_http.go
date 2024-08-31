package http

import (
	"net/http"
	"strconv"

	"github.com/kodinggo/product-service-gb1/internal/model"
	"github.com/kodinggo/product-service-gb1/internal/utils"
	"github.com/labstack/echo/v4"
)

func (h *httpHandler) findAllProducts(c echo.Context) error {
	var query model.ProductQuery
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	products, err := h.productUsecase.FindAll(c.Request().Context(), query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status: http.StatusOK,
		Data:   products,
	})
}

func (h *httpHandler) findProductByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	product, err := h.productUsecase.FindByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status: http.StatusOK,
		Data:   product,
	})
}

func (h *httpHandler) createProduct(c echo.Context) error {
	var product model.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	session := utils.GetUserSession(c)
	if session == nil {
		return c.JSON(http.StatusUnauthorized, response{
			Status:  http.StatusUnauthorized,
			Message: "error getting user session",
		})
	}

	if session.Role.Name != "admin" {
		return c.JSON(http.StatusForbidden, response{
			Status:  http.StatusForbidden,
			Message: "only admin can create a product",
		})
	}

	product, err := h.productUsecase.Create(c.Request().Context(), product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response{
		Status: http.StatusCreated,
		Data:   product,
	})
}

func (h *httpHandler) updateProduct(c echo.Context) error {
	var product model.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	session := utils.GetUserSession(c)
	if session == nil {
		return c.JSON(http.StatusUnauthorized, response{
			Status:  http.StatusUnauthorized,
			Message: "error getting user session",
		})
	}

	if session.Role.Name != "admin" {
		return c.JSON(http.StatusForbidden, response{
			Status:  http.StatusForbidden,
			Message: "only admin can update a product",
		})
	}

	product, err := h.productUsecase.Update(c.Request().Context(), product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status: http.StatusOK,
		Data:   product,
	})
}

func (h *httpHandler) deleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	session := utils.GetUserSession(c)
	if session == nil {
		return c.JSON(http.StatusUnauthorized, response{
			Status:  http.StatusUnauthorized,
			Message: "error getting user session",
		})
	}

	if session.Role.Name != "admin" {
		return c.JSON(http.StatusForbidden, response{
			Status:  http.StatusForbidden,
			Message: "only admin can delete a product",
		})
	}

	err = h.productUsecase.Delete(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status: http.StatusOK,
	})
}

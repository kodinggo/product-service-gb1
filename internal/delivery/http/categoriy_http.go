package http

import (
	"net/http"
	"strconv"

	"github.com/kodinggo/product-service-gb1/internal/model"
	"github.com/kodinggo/product-service-gb1/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func (h *httpHandler) findCategories(c echo.Context) error {
	logger := logrus.WithContext(c.Request().Context())

	query := model.CategoryQuery{}

	if err := c.Bind(&query); err != nil {
		logger.Error(err)
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	categories, err := h.categoryUsecase.FindAll(c.Request().Context(), query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status: http.StatusOK,
		Data:   categories,
	})
}

func (h *httpHandler) findCategoryByID(c echo.Context) error {
	logger := logrus.WithContext(c.Request().Context())

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	category, err := h.categoryUsecase.FindByID(c.Request().Context(), id)
	if err != nil {
		var status int
		switch err {
		case gorm.ErrRecordNotFound:
			status = http.StatusNotFound
		default:
			status = http.StatusInternalServerError
		}

		return c.JSON(status, response{
			Status:  status,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status: http.StatusOK,
		Data:   category,
	})
}

func (h *httpHandler) createCategory(c echo.Context) error {
	logger := logrus.WithContext(c.Request().Context())

	category := model.Category{}

	if err := c.Bind(&category); err != nil {
		logger.Error(err)
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := category.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user, err := utils.UserClaims(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
	}

	if user.Role != "admin" {
		return c.JSON(http.StatusForbidden, response{
			Status:  http.StatusForbidden,
			Message: "only admin can create a category",
		})
	}

	newCategory, err := h.categoryUsecase.Create(c.Request().Context(), category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response{
		Status: http.StatusCreated,
		Data:   newCategory,
	})
}

func (h *httpHandler) updateCategory(c echo.Context) error {
	logger := logrus.WithContext(c.Request().Context())

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	category := model.Category{}

	if err := c.Bind(&category); err != nil {
		logger.Error(err)
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := category.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user, err := utils.UserClaims(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
	}

	if user.Role != "admin" {
		return c.JSON(http.StatusForbidden, response{
			Status:  http.StatusForbidden,
			Message: "only admin can update a category",
		})
	}

	updatedCategory, err := h.categoryUsecase.Update(c.Request().Context(), id, category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response{
		Status: http.StatusOK,
		Data:   updatedCategory,
	})
}

func (h *httpHandler) deleteCategory(c echo.Context) error {
	logger := logrus.WithContext(c.Request().Context())

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error(err)
		return c.JSON(http.StatusBadRequest, response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user, err := utils.UserClaims(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
	}

	if user.Role != "admin" {
		return c.JSON(http.StatusForbidden, response{
			Status:  http.StatusForbidden,
			Message: "only admin can delete a category",
		})
	}

	if err := h.categoryUsecase.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusNoContent, response{
		Status: http.StatusNoContent,
	})
}

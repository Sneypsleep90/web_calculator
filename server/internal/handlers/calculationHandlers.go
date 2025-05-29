package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"web_calculator/internal/calculationService"
)

type CalculationHandler struct {
	service calculationService.CalculationService
}

func NewCalculationHandler(s calculationService.CalculationService) *CalculationHandler {
	return &CalculationHandler{service: s}
}

func (h *CalculationHandler) GetCalculations(c echo.Context) error {
	calculations, err := h.service.GetAllCalculation()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get calcultion"})
	}

	return c.JSON(http.StatusOK, calculations)

}

func (h *CalculationHandler) PostCalculation(c echo.Context) error {
	var req calculationService.CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})

	}

	calc, err := h.service.CreateCalculation(req.Expression)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "could not create calculation"})
	}

	return c.JSON(http.StatusOK, calc)

}

func (h *CalculationHandler) PatchCalculation(c echo.Context) error {
	id := c.Param("id")
	var req calculationService.CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid path handler request"})
	}

	UpdatedCalculation, err := h.service.UpdateCalculation(id, req.Expression)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "could not update calculation"})
	}

	return c.JSON(http.StatusOK, UpdatedCalculation)
}

func (h *CalculationHandler) DeleteCalculation(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteCalculation(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete calculation"})
	}

	return c.NoContent(http.StatusNoContent)
}

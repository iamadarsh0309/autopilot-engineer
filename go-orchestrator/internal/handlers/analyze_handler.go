package handlers

import (
	"autopilot-engineer/go-orchestrator/internal/services"
	"autopilot-engineer/go-orchestrator/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type AnalyzeHandler struct {
	service *services.AnalyzeService
}

func NewAnalyzeHandler(s *services.AnalyzeService) *AnalyzeHandler {
	return &AnalyzeHandler{service: s}
}

func (h *AnalyzeHandler) Analyze(c *fiber.Ctx) error {
	var body struct {
		Text string `json:"text"`
	}
	if err := c.BodyParser(&body); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid JSON")
	}

	result, err := h.service.AnalyzeText(body.Text)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, fiber.Map{
		"analysis": result,
	})
}

package sampleapps

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"with-sheets/types"
	"with-sheets/utils"

	"github.com/go-playground/validator/v10"
	"google.golang.org/api/sheets/v4"
)

type SampleAppSheets interface {
	AddSampleAppsRequest(payload types.Request) (*sheets.AppendValuesResponse, error)
}

type Handler struct {
	sheets SampleAppSheets
}

func NewHandler(sheets SampleAppSheets) *Handler {
	return &Handler{sheets: sheets}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /sample-apps", h.handleAppRequest)
}

func (h *Handler) handleAppRequest(w http.ResponseWriter, r *http.Request) {
	var payload types.Request
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	utils.SendSampleAppsRequestEmail(payload)

	resp, err := h.sheets.AddSampleAppsRequest(payload)
	if err != nil || resp.HTTPStatusCode != http.StatusOK {
		log.Println(err)
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{
		"status":  "ok",
		"message": fmt.Sprintf("Successfully submitted request for %s sample application(s)", strings.Join(payload.Apps, ", ")),
	})
}

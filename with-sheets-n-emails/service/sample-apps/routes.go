package sampleapps

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"with-sheets-n-emails/types"
	"with-sheets-n-emails/utils"

	"github.com/go-playground/validator/v10"
	"google.golang.org/api/sheets/v4"
)

type SampleAppSheets interface {
	AddSampleAppsRequest(payload types.AppsRequestParams) (*sheets.AppendValuesResponse, error)
}

type SampleAppEmails interface {
	SendConfirmationEmail(payload types.AppsRequestParams) (bool, error)
	SendNotificationEmail(payload types.AppsRequestParams) (bool, error)
}

type Handler struct {
	sheets SampleAppSheets
	emails SampleAppEmails
}

func NewHandler(sheets SampleAppSheets, emails SampleAppEmails) *Handler {
	return &Handler{
		sheets: sheets,
		emails: emails,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /sample-apps", h.handleAppRequest)
}

func (h *Handler) handleAppRequest(w http.ResponseWriter, r *http.Request) {
	var payload types.AppsRequestParams
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	if payload.Subscribe == "on" {
		log.Println(payload.Subscribe)
	}

	if _, err := h.emails.SendConfirmationEmail(payload); err != nil {
		log.Println(err)
	}

	if _, err := h.emails.SendNotificationEmail(payload); err != nil {
		log.Println(err)
	}

	resp, err := h.sheets.AddSampleAppsRequest(payload)
	if err != nil || resp.HTTPStatusCode != http.StatusOK {
		log.Println(err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("there was an unknown problem"))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{
		"status":  "ok",
		"message": fmt.Sprintf("Successfully submitted request for %s sample application(s)", strings.Join(payload.Apps, ", ")),
	})
}

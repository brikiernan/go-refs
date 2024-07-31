package sampleapps

import (
	"context"
	"strings"
	"with-sheets/types"
	"with-sheets/utils"

	"google.golang.org/api/sheets/v4"
)

type Sheets struct {
	ctx *context.Context
	srv *sheets.Service
}

func NewSheets(ctx *context.Context, srv *sheets.Service) *Sheets {
	return &Sheets{
		ctx: ctx,
		srv: srv,
	}
}

func (s *Sheets) AddSampleAppsRequest(payload types.Request) (*sheets.AppendValuesResponse, error) {
	spreadsheetID := utils.Getenv("GOOGLE_SHEET_ID")

	row := &sheets.ValueRange{
		Values: [][]interface{}{{
			payload.Name,
			payload.Email,
			payload.Company,
			strings.Join(payload.Apps, ", "),
		}},
	}

	return s.srv.Spreadsheets.Values.Append(spreadsheetID, "Sheet1", row).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Context(*s.ctx).Do()
}

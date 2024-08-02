package sampleapps

import (
	"context"
	"strings"
	"with-sheets-n-emails/types"
	"with-sheets-n-emails/utils"

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

func (s *Sheets) AddSampleAppsRequest(payload types.AppsRequestParams) (*sheets.AppendValuesResponse, error) {
	spreadsheetID := utils.Getenv("GOOGLE_SHEET_ID")

	row := &sheets.ValueRange{
		Values: [][]interface{}{{
			payload.First,
			payload.Last,
			payload.Email,
			payload.Company,
			payload.Use,
			strings.Join(payload.Apps, ", "),
		}},
	}

	return s.srv.Spreadsheets.Values.Append(spreadsheetID, "Sheet1", row).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Context(*s.ctx).Do()
}

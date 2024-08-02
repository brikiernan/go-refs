package google

import (
	"context"
	"log"
	"with-sheets-n-emails/utils"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func NewGoogleSheetsService() (*context.Context, *sheets.Service) {

	config := &jwt.Config{
		Email:      utils.Getenv("GOOGLE_CLIENT_EMAIL"),
		PrivateKey: []byte(utils.Getenv("GOOGLE_PRIVATE_KEY")),
		Scopes:     []string{sheets.SpreadsheetsScope},
		TokenURL:   google.JWTTokenURL,
	}

	ctx := context.Background()
	client := config.Client(ctx)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatal(err)
	}

	return &ctx, srv
}

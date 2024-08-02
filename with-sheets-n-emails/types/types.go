package types

type AppsRequestParams struct {
	First     string   `json:"first" validate:"required"`
	Last      string   `json:"last" validate:"required"`
	Email     string   `json:"email" validate:"required,email"`
	Company   string   `json:"company"`
	Use       string   `json:"use" validate:"required"`
	Apps      []string `json:"apps" validate:"required,min=1"`
	Subscribe string   `json:"subscribe" validate:"required"`
}

type EmailParams struct {
	FromName  string
	FromEmail string
	To        []string
	ReplyTo   string
	Subject   string
	Body      string
}

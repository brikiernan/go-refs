package types

type Request struct {
	Name    string   `json:"name" validate:"required"`
	Email   string   `json:"email" validate:"required,email"`
	Company string   `json:"company"`
	Apps    []string `json:"apps" validate:"required"`
}

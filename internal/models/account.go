package models

type AccountCreateRequest struct {
	Email      string `json:"email"`
	FullNumber string `json:"full_number"`
	Name       string `json:"name"`
	ZipCode    string `json:"zip_code"`
}

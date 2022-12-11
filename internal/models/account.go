package models

type AccountCreateRequest struct {
	Email      string `json:"email" validate:"required,email"`
	FullNumber string `json:"full_number" validate:"required"`
	Name       string `json:"name" validate:"required"`
	ZipCode    string `json:"zip_code" validate:"required"`
}

type AccountUpdateRequest struct {
	Id         string `param:"id" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	FullNumber string `json:"full_number" validate:"required"`
	Name       string `json:"name" validate:"required"`
	ZipCode    string `json:"zip_code" validate:"required"`
}

type AccountDeleteRequest struct {
	Id string `param:"id" validate:"required"`
}

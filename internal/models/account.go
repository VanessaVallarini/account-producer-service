package models

type AccountCreateRequest struct {
	Email      string `json:"email" validate:"required,email"`
	FullNumber string `json:"full_number" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Status     string `json:"status" validate:"required"`
	ZipCode    string `json:"zip_code" validate:"required"`
}

type AccountUpdateRequest struct {
	Id         string `param:"id" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	FullNumber string `json:"full_number" validate:"required"`
	Name       string `json:"name" validate:"required"`
	ZipCode    string `json:"zip_code" validate:"required"`
	Status     string `json:"status" validate:"required"`
}

type AccountRequestByEmail struct {
	Email string `json:"email" validate:"required,email"`
}

type Account struct {
	Email       string `json:"email"`
	FullNumber  string `json:"full_number"`
	Alias       string `json:"alias"`
	City        string `json:"city"`
	DateTime    string `json:"date_time"`
	District    string `json:"district"`
	Name        string `json:"name"`
	PublicPlace string `json:"public_place"`
	Status      string `json:"status"`
	ZipCode     string `json:"zip_code"`
}

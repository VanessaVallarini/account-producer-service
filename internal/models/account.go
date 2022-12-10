package models

type AccountCreate struct {
	Alias       string `json:"alias"`
	City        string `json:"city"`
	District    string `json:"district"`
	Email       string `json:"email"`
	FullNumber  string `json:"full_number"`
	Name        string `json:"name"`
	PublicPlace string `json:"public_place"`
	ZipCode     string `json:"zip_code"`
}

type Account struct {
	Id          string `json:"id"`
	Alias       string `json:"alias"`
	City        string `json:"city"`
	District    string `json:"district"`
	Email       string `json:"email"`
	FullNumber  string `json:"full_number"`
	Name        string `json:"name"`
	PublicPlace string `json:"public_place"`
	ZipCode     string `json:"zip_code"`
}

type AccountCreateRequest struct {
	Email      string `json:"email"`
	FullNumber string `json:"full_number"`
	Name       string `json:"name"`
	ZipCode    string `json:"zip_code"`
}

type AccountRequestBy struct {
	Id         string `json:"id"`
	Email      string `json:"email"`
	FullNumber string `json:"full_number"`
}

type AccountRequestByEmail struct {
	Email string `json:"email"`
}

type AccountRequestByFullNumber struct {
	FullNumber string `json:"full_number"`
}

type AccountRequestById struct {
	Id string `json:"id"`
}

package repository

import (
	"account-producer-service/internal/models"
	"account-producer-service/internal/pkg/db"
	"account-producer-service/internal/pkg/utils"
	"context"
	"encoding/json"
)

type IAccountRepository interface {
	GetByEmail(ctx context.Context, a models.AccountRequestByEmail) (*models.Account, error)
	List(ctx context.Context) ([]models.Account, error)
}

type AccountRepository struct {
	scylla db.IScylla
}

func NewAccountRepository(s db.IScylla) *AccountRepository {
	return &AccountRepository{
		scylla: s,
	}
}

func (repo *AccountRepository) GetByEmail(ctx context.Context, a models.AccountRequestByEmail) (*models.Account, error) {
	stmt := `SELECT * FROM account WHERE email = ?`
	account := &models.Account{}
	results := map[string]interface{}{
		"email":        &account.Email,
		"full_number":  &account.FullNumber,
		"alias":        &account.Alias,
		"city":         &account.City,
		"date_time":    &account.DateTime,
		"district":     &account.District,
		"name":         &account.Name,
		"public_place": &account.PublicPlace,
		"status":       &account.Status,
		"zip_code":     &account.ZipCode,
	}
	err := repo.scylla.ScanMap(ctx, stmt, results, a.Email)
	if err != nil {
		utils.Logger.Error("account producer failed during query get account by email: %v", err)
		return nil, err
	}

	return account, nil
}

func (repo *AccountRepository) List(ctx context.Context) ([]models.Account, error) {
	stmt := `SELECT * FROM account`

	uList, err := repo.scylla.ScanMapSlice(ctx, stmt)
	if err != nil {
		utils.Logger.Error("account producer failed during query get all accounts: %v", err)
		return nil, err
	}

	convertToUserList := repo.scanAccountList(uList)

	return convertToUserList, nil
}

func (repo *AccountRepository) scanAccountList(results []map[string]interface{}) []models.Account {
	var aList []models.Account

	marshallResult, _ := json.Marshal(results)
	json.Unmarshal(marshallResult, &aList)

	return aList
}

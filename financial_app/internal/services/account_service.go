package services

import (
	"financial_app/internal/models"
	"financial_app/internal/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ValidateAccountId(pool *pgxpool.Pool) {
	Id := repositories.GetLastAccountID(pool)
	idCounter = Id
}

func CreateAccountObject(pool *pgxpool.Pool) {
	accountId := GenerateID()
	account := models.Account{
		AccountID:     accountId,
		CustomerID:    1,
		AccountTypeID: 1,
		Balace:        1000.0,
	}
	repositories.InsertAccountTransactionBody(pool, account)
}

package services

import (
	"financial_app/internal/models"
	"financial_app/internal/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

var (
	idCounterAccount int
)

func validateAccountId(pool *pgxpool.Pool) {
	Id := repositories.GetLastAccountIDObject(pool)
	idCounterAccount = Id
}

func createAccountObject(pool *pgxpool.Pool) error {
	account := models.Account{
		CustomerID:    1,
		AccountTypeID: 1,
		Balace:        1000.0,
	}
	return repositories.InsertAccountObjectPool(pool, account)
}

func getAccountIdPivot(pool *pgxpool.Pool) {
	validateAccountId(pool)
	if idCounterAccount == 0 {
		createAccountObject(pool)
		validateAccountId(pool)
	}
	logrus.Infof("account_id pivot -> %v", idCounterAccount)
}

func getAccountInserts(pool *pgxpool.Pool) {
	Id := repositories.GetLastAccountIDObject(pool)
	result := Id - idCounterAccount
	logrus.Infof("Numero de inserts realizados en la tabla account -> %v", result)
}

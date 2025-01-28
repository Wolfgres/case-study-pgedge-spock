package services

import (
	"financial_app/internal/models"
	"financial_app/internal/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

var (
	idCounterAccount int64
)

func validateAccountId(pool *pgxpool.Pool) {
	Id := repositories.GetLastAccountIDObject(pool)
	idCounterAccount = Id
}

func createAccountObject(pool *pgxpool.Pool, node int) {
	account := models.Account{
		CustomerID:    1,
		AccountTypeID: 1,
		Balace:        1000.0,
	}
	repositories.InsertAccountObjectPool(pool, account)
	getCounter(node)
}

func getAccountIdPivot(pool *pgxpool.Pool, node int) {
	validateAccountId(pool)
	if idCounterAccount == 0 {
		createAccountObject(pool, node)
		validateAccountId(pool)
	}
	logrus.Infof("account_id pivot -> %v", idCounterAccount)
}

func editAccountObject(pool *pgxpool.Pool, node int) {
	account, _ := repositories.GetAccountObject(pool, idCounterAccount)
	account.AccountTypeID = 2
	account.Balace = 2000.0
	repositories.UpdateAccountObject(pool, *account)
	getCounter(node)
}

func getAccountObjectPage(pool *pgxpool.Pool, node int) {
	repositories.GetAccountObjects(pool)
	getCounter(node)
}

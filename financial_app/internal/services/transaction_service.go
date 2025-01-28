package services

import (
	"financial_app/internal/models"
	"financial_app/internal/repositories"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	idCounterTransaction int64 // Contador global de IDs
)

func validateTransactionId(pool *pgxpool.Pool) {
	idCounterTransaction = repositories.GetLastTransactionIDObject(pool)
}

func createTransactionObject(pool *pgxpool.Pool, node int) {
	now := time.Now().UTC()
	transaction := models.Transaction{
		AccountID:   idCounterAccount,
		OperationID: 1,
		Mount:       1000.0,
		Date:        now,
	}
	repositories.InsertTransactionObjectPool(pool, transaction)
	getCounter(node)
}

func getTransactionObjectPage(pool *pgxpool.Pool, node int) {
	repositories.GetTransactionObjects(pool)
	getCounter(node)
}

func editTransactionOnject(pool *pgxpool.Pool, node int) {
	transaction, _ := repositories.GetTransactionObject(pool, idCounterTransaction)
	transaction.Mount = 2300.0
	transaction.Date = time.Now()
	repositories.UpdateTransactionObject(pool, *transaction)
	getCounter(node)
}

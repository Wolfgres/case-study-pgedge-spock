package services

import (
	"financial_app/internal/models"
	"financial_app/internal/repositories"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

var (
	idCounterTransaction int // Contador global de IDs
)

func validateTransactionId(pool *pgxpool.Pool) {
	idCounterTransaction = repositories.GetLastTransactionIDObject(pool)
}

func createTransactionObject(pool *pgxpool.Pool) error {
	now := time.Now().UTC()
	transaction := models.Transaction{
		AccountID:   idCounterAccount,
		OperationID: 1,
		Mount:       1000.0,
		Date:        now,
	}
	return repositories.InsertTransactionObjectPool(pool, transaction)
}

func getTransactionInserts(pool *pgxpool.Pool) {
	Id := repositories.GetLastTransactionIDObject(pool)
	result := Id - idCounterTransaction
	logrus.Infof("Numero de inserts realizados en la tabla transaction -> %v", result)
}

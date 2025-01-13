package services

import (
	"financial_app/internal/models"
	"financial_app/internal/repositories"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ValidateTransactionId(pool *pgxpool.Pool) {
	Id := repositories.GetLastTransactionID(pool)
	idCounter = Id
}

func CreateTransactionObject(pool *pgxpool.Pool) {
	transactionId := GenerateID()
	now := time.Now().UTC()
	transaction := models.Transaction{
		TransactionID: transactionId,
		AccountID:     1,
		OperationID:   1,
		Mount:         1000.0,
		Date:          now,
	}
	repositories.InsertTransactionTransactionBody(pool, transaction)
}

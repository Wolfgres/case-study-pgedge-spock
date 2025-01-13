package repositories

import (
	"context"
	"financial_app/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func GetInsertTransactionQuery() string {
	return "INSERT INTO wfg.transaction (transaction_id, account_id, operation_id, mount, date) VALUES ($1, $2, $3, $4, $5)"
}

func InsertTransaction(pool *pgxpool.Pool, mTransaction models.Transaction) {
	_, err := pool.Exec(
		context.Background(),
		GetInsertTransactionQuery(),
		mTransaction.TransactionID,
		mTransaction.AccountID,
		mTransaction.OperationID,
		mTransaction.Mount,
		mTransaction.Date,
	)
	if err != nil {
		logrus.Fatalf("Error al intentar insertar datos: %v", err)
	}
}

func GetLastTransactionID(pool *pgxpool.Pool) int {
	query := "SELECT COALESCE(MAX(transaction_id), 0) FROM wfg.transaction"
	return GetLastID(pool, query)
}

func InsertTransactionTransactionBody(pool *pgxpool.Pool, mTransaction models.Transaction) int {
	ctx := context.Background()

	// Iniciar la transacción
	tx, err := pool.Begin(ctx)
	if err != nil {
		logrus.Printf("Error al comenzar transacción: %v", err)
		return 0
	}

	// Ejecutar operación de escritura dentro de la transacción
	_, err = tx.Exec(
		ctx,
		GetInsertTransactionQuery(),
		mTransaction.TransactionID,
		mTransaction.AccountID,
		mTransaction.OperationID,
		mTransaction.Mount,
		mTransaction.Date,
	)
	if err != nil {
		logrus.Printf("Error al ejecutar operación en transacción: %v", err)
		tx.Rollback(ctx)
		return 0
	}

	// Confirmar la transacción
	if err := tx.Commit(ctx); err != nil {
		logrus.Printf("Error al hacer commit de la transacción: %v", err)
		return 0
	}

	return mTransaction.TransactionID
}

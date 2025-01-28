package repositories

import (
	"context"
	"financial_app/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

/*
TODO: don't use this method to create PK, instead that use sequence
*/
func GetLastTransactionIDObject(pool *pgxpool.Pool) int64 {
	query := "SELECT COALESCE(MAX(transaction_id), 0) FROM wfg.transaction"
	return GetLastID(pool, query)
}

func InsertTransactionObjectPool(pool *pgxpool.Pool, mTransaction models.Transaction) {
	query := "INSERT INTO wfg.transaction (account_id, operation_id, mount, date) VALUES ($1, $2, $3, $4)"

	// Ejecutar operación de escritura dentro de la transacción
	_, err := pool.Exec(
		context.Background(),
		query,
		mTransaction.AccountID,
		mTransaction.OperationID,
		mTransaction.Mount,
		mTransaction.Date,
	)
	if err != nil {
		logrus.Fatalf("Error executing INSERT in transaction: %v", err)
	}
}

func GetTransactionObjects(pool *pgxpool.Pool) ([]models.Transaction, error) {
	query := "SELECT * FROM wfg.transaction AS t ORDER BY t.transaction_id ASC LIMIT 20"
	// Ejecutar la consulta
	rows, err := pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Lista para almacenar los resultados
	var transactions []models.Transaction

	// Iterar sobre los resultados
	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(
			&transaction.TransactionID,
			&transaction.AccountID,
			&transaction.OperationID,
			&transaction.Mount,
			&transaction.Date)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	// Verificar errores de iteración
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func GetTransactionObject(pool *pgxpool.Pool, transactionID int64) (*models.Transaction, error) {
	var transaction models.Transaction
	query := "SELECT t.transaction_id, t.account_id, t.operation_id, t.mount, t.date FROM wfg.transaction AS t WHERE t.transaction_id = $1"
	err := pool.QueryRow(context.Background(), query, transactionID).Scan(
		&transaction.TransactionID,
		&transaction.AccountID,
		&transaction.OperationID,
		&transaction.Mount,
		&transaction.Date,
	)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func UpdateTransactionObject(pool *pgxpool.Pool, mTransaction models.Transaction) {
	query := "UPDATE wfg.transaction SET account_id = $1, operation_id = $2, mount = $3, date = $4 WHERE transaction_id = $5"
	_, err := pool.Exec(
		context.Background(),
		query,
		mTransaction.AccountID,
		mTransaction.OperationID,
		mTransaction.Mount,
		mTransaction.Date,
		mTransaction.TransactionID,
	)
	if err != nil {
		logrus.Fatalf("Error executing UPDATE in transaction: %v", err)
	}
}

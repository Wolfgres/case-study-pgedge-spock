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
func GetLastTransactionIDObject(pool *pgxpool.Pool) int {
	query := "SELECT COALESCE(MAX(transaction_id), 0) FROM wfg.transaction"
	return GetLastID(pool, query)
}

func InsertTransactionObject(pool *pgxpool.Pool, mTransaction models.Transaction) {
	ctx := context.Background()
	query := "INSERT INTO wfg.transaction (account_id, operation_id, mount, date) VALUES ($1, $2, $3, $4)"
	// Iniciar la transacción
	tx := BeginTransaction(pool)

	// Ejecutar operación de escritura dentro de la transacción
	_, err := tx.Exec(
		ctx,
		query,
		mTransaction.AccountID,
		mTransaction.OperationID,
		mTransaction.Mount,
		mTransaction.Date,
	)
	if err != nil {
		logrus.Fatalf("Error al ejecutar operación en transacción: %v", err)
		tx.Rollback(ctx)
		return
	}
	CommitTransaction(tx)
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

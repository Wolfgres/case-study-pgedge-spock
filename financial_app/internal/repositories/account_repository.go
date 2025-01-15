package repositories

import (
	"context"
	"financial_app/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func InsertAccountObject(pool *pgxpool.Pool, mAccount models.Account) {
	ctx := context.Background()
	query := "INSERT INTO wfg.account (account_id, customer_id, account_type_id, balace) VALUES ($1, $2, $3, $4)"
	// Iniciar la transacción
	tx := BeginTransaction(pool)

	// Ejecutar operación de escritura dentro de la transacción
	_, err := tx.Exec(
		ctx,
		query,
		mAccount.AccountID,
		mAccount.CustomerID,
		mAccount.AccountTypeID,
		mAccount.Balace,
	)
	if err != nil {
		logrus.Fatalf("Error al ejecutar operación en transacción: %v", err)
		tx.Rollback(ctx)
		return
	}

	CommitTransaction(tx)
}

func GetLastAccountIDObject(pool *pgxpool.Pool) int {
	query := "SELECT COALESCE(MAX(account_id), 0) FROM wfg.account"
	return GetLastID(pool, query)
}

func GetAccountObjects(pool *pgxpool.Pool) ([]models.Account, error) {
	query := "SELECT * FROM wfg.account AS t ORDER BY t.account_id ASC LIMIT 20"
	// Ejecutar la consulta
	rows, err := pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Lista para almacenar los resultados
	var accounts []models.Account

	// Iterar sobre los resultados
	for rows.Next() {
		var account models.Account
		err := rows.Scan(
			account.AccountID,
			account.CustomerID,
			account.AccountTypeID,
			account.Balace,
		)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	// Verificar errores de iteración
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func GetAccountObject(pool *pgxpool.Pool, accountID int) (*models.Account, error) {
	var account models.Account
	query := "SELECT t.account_id, t.customer_id, t.account_type_id, t.balace FROM wfg.account AS t WHERE t.account_id = $1"
	err := pool.QueryRow(context.Background(), query, accountID).Scan(
		&account.AccountID,
		&account.CustomerID,
		&account.AccountTypeID,
		&account.Balace)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

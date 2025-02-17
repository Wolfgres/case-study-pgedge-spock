package repositories

import (
	"context"
	"financial_app/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func InsertAccountObjectPool(pool *pgxpool.Pool, mAccount models.Account) error {
	query := "INSERT INTO wfg.account (customer_id, account_type_id, balace) VALUES ($1, $2, $3)"
	// Ejecutar operación de escritura dentro de la transacción
	_, err := pool.Exec(
		context.Background(),
		query,
		mAccount.CustomerID,
		mAccount.AccountTypeID,
		mAccount.Balace,
	)
	if err != nil {
		logrus.Fatalf("Error executing INSERT on account: %v", err)
		return err
	}
	return nil
}

func GetLastAccountIDObject(pool *pgxpool.Pool) int64 {
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

func GetLastAccountID(pool *pgxpool.Pool) int64 {
	query := "SELECT COALESCE(MAX(account_id), 0) FROM wfg.account"
	return GetLastID(pool, query)
}

func GetAccountObject(pool *pgxpool.Pool, accountID int64) (*models.Account, error) {
	var account models.Account
	query := "SELECT a.account_id, a.customer_id, a.account_type_id, a.balace FROM wfg.account AS a WHERE a.account_id = $1"
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

func UpdateAccountObject(pool *pgxpool.Pool, mAccount models.Account) {
	query := "UPDATE wfg.account SET customer_id = $1, account_type_id = $2, balace = $3 WHERE account_id = $4"
	_, err := pool.Exec(
		context.Background(),
		query,
		mAccount.CustomerID,
		mAccount.AccountTypeID,
		mAccount.Balace,
		mAccount.AccountID,
	)
	if err != nil {
		logrus.Fatalf("Error executing UPDATE on account: %v", err)
	}
}

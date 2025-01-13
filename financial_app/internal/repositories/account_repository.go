package repositories

import (
	"context"
	"financial_app/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func GetInsertAccountQuery() string {
	return "INSERT INTO wfg.account (account_id, customer_id, account_type_id, balace) VALUES ($1, $2, $3, $4)"
}

func InsertAccount(pool *pgxpool.Pool, mAccount models.Account) {
	_, err := pool.Exec(
		context.Background(),
		GetInsertAccountQuery(),
		mAccount.AccountID,
		mAccount.CustomerID,
		mAccount.AccountTypeID,
		mAccount.Balace,
	)
	if err != nil {
		logrus.Fatalf("Error al intentar insertar datos: %v", err)
	}
}

func InsertAccountTransactionBody(pool *pgxpool.Pool, mAccount models.Account) int {
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
		GetInsertAccountQuery(),
		mAccount.AccountID,
		mAccount.CustomerID,
		mAccount.AccountTypeID,
		mAccount.Balace,
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

	return mAccount.AccountID
}

func GetLastAccountID(pool *pgxpool.Pool) int {
	query := "SELECT COALESCE(MAX(account_id), 0) FROM wfg.account"
	return GetLastID(pool, query)
}

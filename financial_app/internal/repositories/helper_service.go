package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func GetLastID(pool *pgxpool.Pool, query string) int {
	var lastID int
	err := pool.QueryRow(context.Background(), query).Scan(&lastID)
	if err != nil {
		logrus.Fatalf("Error al obtener el último ID: %v", err)
		return 0
	}
	return lastID
}

func BeginTransaction(pool *pgxpool.Pool) pgx.Tx {
	tx, err := pool.Begin(context.Background())
	if err != nil {
		logrus.Fatalf("Error al comenzar transacción: %v", err)
	}
	return tx
}

func CommitTransaction(tx pgx.Tx) {
	// Confirmar la transacción
	if err := tx.Commit(context.Background()); err != nil {
		logrus.Fatalf("Error al hacer commit de la transacción: %v", err)
	}
}

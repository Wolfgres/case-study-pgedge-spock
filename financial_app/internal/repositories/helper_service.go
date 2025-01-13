package repositories

import (
	"context"

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

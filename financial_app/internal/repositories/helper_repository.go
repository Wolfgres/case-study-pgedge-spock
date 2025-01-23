package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func GetLastID(pool *pgxpool.Pool, query string) int64 {
	var lastID int64
	err := pool.QueryRow(context.Background(), query).Scan(&lastID)
	if err != nil {
		logrus.Fatalf("Error getting last ID: %v", err)
		return 0
	}
	return lastID
}

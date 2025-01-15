package services

import (
	"financial_app/internal/database"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func StressTest(pool *pgxpool.Pool) {
	var wg sync.WaitGroup
	numTransactions, _ := parseAnyToInt(viper.Get("test.transactions")) // Número de transacciones
	// Si max_connections = 100 en postgresql, evita que maxConcurrent supere ese número.
	maxConcurrent, _ := parseAnyToInt(viper.Get("test.maxCurrent")) // Número máximo de transacciones concurrentes
	// Control de concurrencia con un canal
	semaphore := make(chan struct{}, maxConcurrent)

	for i := 0; i < numTransactions; i++ {
		wg.Add(1)
		semaphore <- struct{}{} // Bloquear si hay demasiadas transacciones en progreso

		go func() {
			defer wg.Done()
			CreateAccountObject(pool)
			CreateTransactionObject(pool)
			<-semaphore
		}()
	}

	// Esperar a que todas las goroutines terminen
	wg.Wait()
}

func GenerateStressTest() {
	pool, err := database.PgxPoolConnection()
	if err != nil {
		logrus.Fatal("No se pudo conectar a la base de datos:", err)
	}
	defer pool.Close() // Asegura que el pool se cierre correctamente
	logrus.Info("Generando estrés con transacciones concurrentes...")
	start := time.Now()
	StressTest(pool)
	elapsed := time.Since(start)
	logrus.Infof("Prueba de estrés completada en %v", elapsed)
}

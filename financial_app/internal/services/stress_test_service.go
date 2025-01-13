package services

import (
	"financial_app/internal/database"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func StressTest(pool *pgxpool.Pool) {
	var wg sync.WaitGroup
	numTransactions := 100 // Número de transacciones
	maxConcurrent := 50    // Número máximo de transacciones concurrentes
	// Control de concurrencia con un canal
	semaphore := make(chan struct{}, maxConcurrent)

	for i := 0; i < numTransactions; i++ {
		wg.Add(1)
		semaphore <- struct{}{} // Bloquear si hay demasiadas transacciones en progreso

		go func() {
			defer wg.Done()
			CreateTransactionObject(pool)
			//CreateAccountObject(pool)
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
	ValidateTransactionId(pool)
	//ValidateAccountId(pool)
	logrus.Info("Generando estrés con transacciones concurrentes...")
	StressTest(pool)
	logrus.Info("Prueba de estrés completada.")
}

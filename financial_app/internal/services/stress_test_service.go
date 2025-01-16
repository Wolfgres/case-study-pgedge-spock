package services

import (
	"financial_app/internal/database"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func getNumInserts(pool *pgxpool.Pool) {
	GetAccountInserts(pool)
	GetTransactionInserts(pool)
}

func initCounterInserts(pool *pgxpool.Pool) {
	GetAccountIdPivot(pool)     // Se crea un bjeto account como pivote para transaction.
	ValidateTransactionId(pool) // Se obtinene el Id inicial
}

func StressTest(pool *pgxpool.Pool) {
	var wg sync.WaitGroup
	//numTransactions, _ := parseAnyToInt(viper.Get("test.transactions")) // Número de transacciones
	numGoroutines, _ := parseAnyToInt(viper.Get("test.num_goroutines")) // Número de hilos (workers)
	secondLongValue, _ := parseAnyToInt(viper.Get("test.seconds_long"))
	secondLong := time.Duration(secondLongValue) * time.Second
	stopChan := make(chan struct{}) // Canal para detener las goroutines

	initCounterInserts(pool)
	start := time.Now()
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case <-stopChan:
					return
				default:
					CreateAccountObject(pool)
					CreateTransactionObject(pool)
				}
			}
		}(i)
	}

	// Esperar el tiempo especificado
	time.Sleep(secondLong)

	// Señal para detener las goroutines
	close(stopChan)
	elapsed := time.Since(start)
	logrus.Infof("Prueba de estrés completada en %v", elapsed)
	// Esperar a que todas las goroutines terminen
	logrus.Infof("Esperando a que las goroutines terminen...")
	wg.Wait()
	getNumInserts(pool)
}

func GenerateStressTest() {
	pool, err := database.PgxPoolConnection()
	if err != nil {
		logrus.Fatal("No se pudo conectar a la base de datos:", err)
	}
	defer pool.Close() // Asegura que el pool se cierre correctamente
	logrus.Info("Generando estrés con transacciones concurrentes...")
	StressTest(pool)
}

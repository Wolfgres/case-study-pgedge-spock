package services

import (
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func getNumInserts(pool *pgxpool.Pool) {
	getAccountInserts(pool)
	getTransactionInserts(pool)
}

func initCounterInserts(pool *pgxpool.Pool) {
	getAccountIdPivot(pool)     // Se crea un bjeto account como pivote para transaction.
	validateTransactionId(pool) // Se obtinene el Id inicial
}

func stressTestByTransactions(pool *pgxpool.Pool, numGoroutines int, numTransactions int) {
	var wg sync.WaitGroup
	initCounterInserts(pool)
	start := time.Now()
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < numTransactions/numGoroutines; j++ {
				createAccountObject(pool)
				createTransactionObject(pool)
			}
		}(i)
	}
	elapsed := time.Since(start)
	logrus.Infof("Prueba de estrés completada en %v", elapsed)
	logrus.Infof("Esperando a que las goroutines terminen...")
	wg.Wait()
	getNumInserts(pool)
}

func stressTest(pool *pgxpool.Pool, numGoroutines int, secondDuration int) {
	var wg sync.WaitGroup
	secondLong := time.Duration(secondDuration) * time.Second
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
					createAccountObject(pool)
					createTransactionObject(pool)
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
	logrus.Infof("Esperando a que las goroutines terminen...")
	wg.Wait()
	getNumInserts(pool)
}

func StartStressTest(pool *pgxpool.Pool, numGoroutines int, duration int, numTransactions int) {
	if numGoroutines <= 0 {
		logrus.Fatal("The number of goroutines is not valid")
	}
	if duration == 0 && numTransactions == 0 {
		logrus.Fatal("The duration or transaction parameter must be > 0")
	}
	defer pool.Close() // Asegura que el pool se cierre correctamente
	if numTransactions > 0 && duration == 0 {
		logrus.Info("Ejecutando prueba de estres por transacciiones...")
		stressTestByTransactions(pool, numGoroutines, numTransactions)
	}
	if duration > 10 && numTransactions == 0 {
		logrus.Info("Ejecutando prueba de estres por duración...")
		stressTest(pool, numGoroutines, duration)
	}
}

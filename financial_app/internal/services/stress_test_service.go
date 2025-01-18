package services

import (
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

var idCounter int64 = 0
var idMutex sync.Mutex

// Función para generar IDs incrementales de manera concurrente
func getCounter() {
	idMutex.Lock()
	defer idMutex.Unlock()
	idCounter++
}

func getCounterInserts(option int) {
	switch option {
	case 1:
		logrus.Infof("Number of INSERTS per table %v", idCounter)
	case 2:
		logrus.Infof("Number oF SELECTS per table %v", idCounter)
	case 3:
		logrus.Infof("Number of UPDATE per table %v", idCounter)
	default:
		logrus.Fatal("Invalid option. Check the options.")
	}
}

func initCounterInserts(pool *pgxpool.Pool, option int) {
	if option == 1 || option == 3 {
		getAccountIdPivot(pool)     // Se crea un bjeto account como pivote para transaction.
		validateTransactionId(pool) // Se obtinene el Id inicial
	}
	if option == 3 && idCounterTransaction == 0 {
		logrus.Fatal("There is no complete data available to UPDATE.")
	}
}

func operation(pool *pgxpool.Pool, option int) {
	switch option {
	case 1:
		createAccountObject(pool)
		createTransactionObject(pool)
	case 2:
		getAccountObjectPage(pool)
		getTransactionObjectPage(pool)
	case 3:
		editAccountObject(pool)
		editTransactionOnject(pool)
	default:
		logrus.Fatal("Invalid operation. Check the options.")
	}
}

func captureStressTestEnd(start time.Time) {
	elapsed := time.Since(start)
	logrus.Infof("Stress test completed in %v.", elapsed)
	logrus.Infof("Waiting for goroutines to finish...")
}

func stressTestPerTransactions(pool *pgxpool.Pool, numGoroutines int, numTransactions int, option int) {
	var wg sync.WaitGroup
	initCounterInserts(pool, option)
	start := time.Now()
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < numTransactions/numGoroutines; j++ {
				operation(pool, option)
			}
		}(i)
	}
	captureStressTestEnd(start)
	wg.Wait()
}

func stressTest(pool *pgxpool.Pool, numGoroutines int, secondDuration int, option int) {
	var wg sync.WaitGroup
	secondLong := time.Duration(secondDuration) * time.Second
	stopChan := make(chan struct{}) // Canal para detener las goroutines
	initCounterInserts(pool, option)
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
					operation(pool, option)
					getCounter()
				}
			}
		}(i)
	}

	// Esperar el tiempo especificado
	time.Sleep(secondLong)
	// Señal para detener las goroutines
	close(stopChan)
	captureStressTestEnd(start)
	wg.Wait()
	getCounterInserts(option)
}

func StartStressTest(pool *pgxpool.Pool, numGoroutines int, duration int, numTransactions int, option int) {
	if numGoroutines <= 0 {
		logrus.Fatal("The number of goroutines is not valid.")
	}
	if duration == 0 && numTransactions == 0 {
		logrus.Fatal("The duration or transaction parameter must be > 0.")
	}
	if duration > 0 && numTransactions > 0 {
		logrus.Fatal("Only one parameter (duration or transaction) must be > 0.")
	}
	defer pool.Close() // Asegura que el pool se cierre correctamente
	if duration > 0 && numTransactions == 0 {
		logrus.Info("Running stress test for seconds duration...")
		stressTest(pool, numGoroutines, duration, option)
	}
	if numTransactions%numGoroutines != 0 {
		logrus.Fatal("Please enter only numbers (goroutines and transactions) that are multiples.")
	}
	if numTransactions > 0 && duration == 0 {
		logrus.Info("Running stress test by number of transactions...")
		stressTestPerTransactions(pool, numGoroutines, numTransactions, option)
	}
}

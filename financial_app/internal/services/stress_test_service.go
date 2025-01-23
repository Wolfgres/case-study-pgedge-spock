package services

import (
	"financial_app/internal/database"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

var (
	nodesCounter []int64
	idMutex      sync.Mutex
	nodes        int
)

const firstNode = 0

// Función para generar IDs incrementales de manera concurrente
func getCounter(node int) {
	idMutex.Lock()
	defer idMutex.Unlock()
	nodesCounter[node]++
}

func getCounterResult(option int) {
	for i, value := range nodesCounter {
		getCounterActions(option, int(value), i+1)
	}
}

func getCounterActions(option int, counted int, node int) {
	switch option {
	case 1:
		logrus.Infof("Number of INSERTS per table %v in Node %v", counted, node)
	case 2:
		logrus.Infof("Number of SELECTS per table %v in Node %v", counted, node)
	case 3:
		logrus.Infof("Number of UPDATE per table %v in Node %v", counted, node)
	default:
		logrus.Fatal("Invalid option. Check the options.")
	}
}

func initCounterInserts(pool *pgxpool.Pool, option int, node int) {
	if option == 1 || option == 3 {
		getAccountIdPivot(pool, node) // Se crea un bjeto account como pivote para transaction.
		validateTransactionId(pool)   // Se obtinene el Id inicial
	}
	if option == 3 && idCounterTransaction == 0 {
		logrus.Fatal("There is no complete data available to UPDATE.")
	}
}

func operation(pool *pgxpool.Pool, option int, node int) {
	switch option {
	case 1:
		createAccountObject(pool, node)
		createTransactionObject(pool, node)
	case 2:
		getAccountObjectPage(pool, node)
		getTransactionObjectPage(pool, node)
	case 3:
		editAccountObject(pool, node)
		editTransactionOnject(pool, node)
	default:
		logrus.Fatal("Invalid operation. Check the options.")
	}
}

func captureStressTestEnd(start time.Time) {
	elapsed := time.Since(start)
	logrus.Infof("Stress test completed in %v.", elapsed)
	logrus.Infof("Waiting for goroutines to finish...")
}

func stressTestPerTransactions(pools []*pgxpool.Pool, numGoroutines int, numTransactions int, option int) {
	var wg sync.WaitGroup
	initCounterInserts(pools[firstNode], option, firstNode)
	start := time.Now()
	nodes = len(pools)
	nodesCounter = make([]int64, nodes)
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < numTransactions/numGoroutines; j++ {
				node := RandomNumbersInRange(nodes)
				operation(pools[node], option, node)
			}
		}(i)
	}
	captureStressTestEnd(start)
	wg.Wait()
	getCounterResult(option)
}

func stressTest(pools []*pgxpool.Pool, numGoroutines int, secondDuration int, option int) {
	var wg sync.WaitGroup
	secondLong := time.Duration(secondDuration) * time.Second
	stopChan := make(chan struct{}) // Canal para detener las goroutines
	initCounterInserts(pools[firstNode], option, firstNode)
	start := time.Now()
	nodes = len(pools)
	nodesCounter = make([]int64, nodes)
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case <-stopChan:
					return
				default:
					node := RandomNumbersInRange(nodes)
					operation(pools[node], option, node)
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
	getCounterResult(option)
}

func startStressTest(pools []*pgxpool.Pool, numGoroutines int, duration int, numTransactions int, option int) {
	if numGoroutines <= 0 {
		logrus.Fatal("The number of goroutines is not valid.")
	}
	if duration == 0 && numTransactions == 0 {
		logrus.Fatal("The duration or transaction parameter must be > 0.")
	}
	if duration > 0 && numTransactions > 0 {
		logrus.Fatal("Only one parameter (duration or transaction) must be > 0.")
	}
	if duration > 0 && numTransactions == 0 {
		logrus.Info("Running stress test for seconds duration...")
		stressTest(pools, numGoroutines, duration, option)
	}
	if numTransactions%numGoroutines != 0 {
		logrus.Fatal("Please enter only numbers (goroutines and transactions) that are multiples.")
	}
	if numTransactions > 0 && duration == 0 {
		logrus.Info("Running stress test by number of transactions...")
		stressTestPerTransactions(pools, numGoroutines, numTransactions, option)
	}
}

func StressTestNodes(numGoroutines int, duration int, numTransactions int, option int, maxConns int) {
	connStrNodes := database.GetConnectionBodies()
	//numNodes := len(connStrNodes)
	var pools []*pgxpool.Pool
	for _, value := range connStrNodes {
		connStr := database.GetConnectionString(value)
		pool := database.InitDatabasePool(maxConns, numGoroutines, connStr)
		pools = append(pools, pool)
	}
	startStressTest(pools, numGoroutines, duration, numTransactions, option)
}

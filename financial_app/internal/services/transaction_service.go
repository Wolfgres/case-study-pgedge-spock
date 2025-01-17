package services

import (
	"financial_app/internal/models"
	"financial_app/internal/repositories"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

var (
	idCounterTransaction int        // Contador global de IDs
	idMutexTransaction   sync.Mutex // Mutex global para proteger el contador
)

func validateTransactionId(pool *pgxpool.Pool) {
	idCounterTransaction = repositories.GetLastTransactionIDObject(pool)
}

func createTransactionObject(pool *pgxpool.Pool) error {
	now := time.Now().UTC()
	transaction := models.Transaction{
		AccountID:   idCounterAccount,
		OperationID: 1,
		Mount:       1000.0,
		Date:        now,
	}
	return repositories.InsertTransactionObjectPool(pool, transaction)
}

// Genera un ID único de manera segura
func generateTransactionID() int {
	idMutexTransaction.Lock()         // Bloquea el Mutex para evitar accesos concurrentes
	defer idMutexTransaction.Unlock() // Asegura que el Mutex se libere después de la función
	idCounterTransaction++            // Incrementa el contador global
	return idCounterTransaction       // Retorna el nuevo ID
}

func getTransactionInserts(pool *pgxpool.Pool) {
	Id := repositories.GetLastTransactionIDObject(pool)
	result := Id - idCounterTransaction
	logrus.Infof("Numero de inserts realizados en la tabla transaction -> %v", result)
}

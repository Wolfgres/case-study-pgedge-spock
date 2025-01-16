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

func ValidateTransactionId(pool *pgxpool.Pool) {
	Id := repositories.GetLastTransactionIDObject(pool)
	idCounterTransaction = Id
}

func CreateTransactionObject(pool *pgxpool.Pool) {
	now := time.Now().UTC()
	transaction := models.Transaction{
		AccountID:   idCounterAccount,
		OperationID: 1,
		Mount:       1000.0,
		Date:        now,
	}
	repositories.InsertTransactionObjectPool(pool, transaction)
}

// Genera un ID único de manera segura
func GenerateTransactionID() int {
	idMutexTransaction.Lock()         // Bloquea el Mutex para evitar accesos concurrentes
	defer idMutexTransaction.Unlock() // Asegura que el Mutex se libere después de la función
	idCounterTransaction++            // Incrementa el contador global
	return idCounterTransaction       // Retorna el nuevo ID
}

func GetTransactionInserts(pool *pgxpool.Pool) {
	Id := repositories.GetLastTransactionIDObject(pool)
	result := Id - idCounterTransaction
	logrus.Infof("Numero de inserts realizados en la tabla transaction -> %v", result)
}

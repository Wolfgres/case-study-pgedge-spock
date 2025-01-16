package services

import (
	"financial_app/internal/models"
	"financial_app/internal/repositories"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

var (
	idCounterAccount int
	idMutexAccount   sync.Mutex
)

func ValidateAccountId(pool *pgxpool.Pool) {
	Id := repositories.GetLastAccountIDObject(pool)
	idCounterAccount = Id
}

func CreateAccountObject(pool *pgxpool.Pool) {
	account := models.Account{
		CustomerID:    1,
		AccountTypeID: 1,
		Balace:        1000.0,
	}
	repositories.InsertAccountObjectPool(pool, account)
}

// Genera un ID único de manera segura
func GenerateAccountID() int {
	idMutexAccount.Lock()         // Bloquea el Mutex para evitar accesos concurrentes
	defer idMutexAccount.Unlock() // Asegura que el Mutex se libere después de la función
	idCounterAccount++            // Incrementa el contador global
	return idCounterAccount       // Retorna el nuevo ID
}

func GetAccountIdPivot(pool *pgxpool.Pool) {
	ValidateAccountId(pool)
	if idCounterAccount == 0 {
		CreateAccountObject(pool)
		ValidateAccountId(pool)
	}
	logrus.Infof("account_id pivot -> %v", idCounterAccount)
}

func GetAccountInserts(pool *pgxpool.Pool) {
	Id := repositories.GetLastAccountIDObject(pool)
	result := Id - idCounterAccount
	logrus.Infof("Numero de inserts realizados en la tabla account -> %v", result)
}
